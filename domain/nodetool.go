package domain

import "github.com/shshang/gorilla-mux-cassandra/errs"

type NodetoolStatus struct {
	Status  string `json:"status"`
	State   string `json:"state"`
	Address string `json:"address"`
	Load    string `json:"load"`
	Tokens  string `json:"tokens"`
	Owns    string `json:"owns"`
	HostId  string `json:"hostid"`
	Rack    string `json:"rack"`
}

//Nodetool is the primary port on server side
type Nodetool interface {
	RetrieveNodetoolStatus() ([]NodetoolStatus, *errs.AppError)
}
