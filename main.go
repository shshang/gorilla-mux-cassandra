package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/nodetool/status", GetNodetoolStatus).Methods("GET")
	router.HandleFunc("/nodetool/compactionthroughput", GetCompactionThroughput).Methods("GET")
	router.HandleFunc("/nodetool/compactionthroughput/{num}", SetCompactionThroughput).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
