package app

import (
	"fmt"
	"github.com/shshang/gorilla-mux-cassandra/domain"
	"github.com/shshang/gorilla-mux-cassandra/logger"
	"github.com/shshang/gorilla-mux-cassandra/service"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	nh := NodetoolHandler{service: service.NewNodetoolService(domain.NewNodetool())}
	router.HandleFunc("/nodetool/status", nh.GetNodetoolStatus).Methods("GET")

	address, ok := os.LookupEnv("API_SERVER_ADDRESS")
	if !ok {
		address = "localhost"
	}

	port, ok := os.LookupEnv("API_SERVER_PORT")
	if !ok {
		port = "8080"
	}

	logger.Info(fmt.Sprintf("API server binds to %s:%s", address, port))
	logger.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router).Error())
}
