package app

import (
	"encoding/json"
	"github.com/shshang/gorilla-mux-cassandra/service"
	"log"
	"net/http"
)

type ErrorReturnedToUser struct {
	Message string
}

//NodetoolHandler has dependency on the user-side port, which is NodetoolService interface
type NodetoolHandler struct {
	service service.NodetoolService
}

func (nh NodetoolHandler) GetNodetoolStatus(w http.ResponseWriter, r *http.Request) {
	outputs, err := nh.service.NodetoolStatusService()
	if err != nil {
		log.Printf("NodetoolStatusService returned error: %s\n", err.Message)
		writeResponse(w, err.Code, err)
	} else {
		jsonResponse, err := json.Marshal(outputs)

		if err != nil {
			log.Printf("failed to marshal go structure %s\n", err.Error())
			writeResponse(w, http.StatusInternalServerError, "malformed response")
		} else {
			writeResponse(w, http.StatusOK, jsonResponse)
		}
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
