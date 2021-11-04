package app

import (
	"encoding/json"
	"github.com/shshang/gorilla-mux-cassandra/service"
	"log"
	"net/http"

	_ "github.com/shshang/gorilla-mux-cassandra/service"
)

//NodetoolHandler has dependency on the user-side port, which is NodetoolService interface
type NodetoolHandler struct {
	service service.NodetoolService
}

func (nh NodetoolHandler) GetNodetoolStatus(response http.ResponseWriter, request *http.Request) {
	outputs, err := nh.service.NodetoolStatus()
	if err != nil {
		log.Fatalf("service.NodetoolStatus returned error: %s", err.Error())
		return
	}

	jsonResponse, err := json.Marshal(outputs)

	if err != nil {
		log.Fatalf("failed to marshal go structure %s\n", err)
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusNotFound)
		response.Write(jsonResponse)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}
