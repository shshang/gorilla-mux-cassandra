package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os/exec"

	"github.com/draios/scratch/shshang/gorilla-mux-cassandra/domain"
)

type NodetoolService interface {
	GetNodetoolStatus() []Nod
}

func GetNodetoolStatus(response http.ResponseWriter, request *http.Request) {
	outputs := RetrieveStatusFromNodetool()
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

func GetCompactionThroughput(response http.ResponseWriter, request *http.Request) {

	var args = []string{"-n", "sysdigcloud", "exec", "sysdigcloud-cassandra-0", "--", "nodetool", "getcompactionthroughput"}
	cmd := exec.Command("oc", args...)
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	fmt.Fprintf(response, string(out))
}

func SetCompactionThroughput(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var args = []string{"-n", "sysdigcloud", "exec", "sysdigcloud-cassandra-0", "--", "nodetool", "setcompactionthroughput", vars["num"]}
	cmd := exec.Command("oc", args...)

	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	fmt.Fprintf(w, string(out))
}
