package app

import (
	"github.com/shshang/gorilla-mux-cassandra/domain"
	"github.com/shshang/gorilla-mux-cassandra/logger"
	"github.com/shshang/gorilla-mux-cassandra/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	//wiring
	nh := NodetoolHandler{service: service.NewNodetoolService(domain.NewNodetool())}
	router.HandleFunc("/nodetool/status", nh.GetNodetoolStatus).Methods("GET")
	logger.Fatal(http.ListenAndServe(":8080", router).Error())
}
