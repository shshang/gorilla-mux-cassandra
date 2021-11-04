package service

import (
	"github.com/shshang/gorilla-mux-cassandra/domain"
)

//NodetoolService is the primary port on user side
type NodetoolService interface {
	NodetoolStatus() ([]domain.NodetoolStatus, error)
}

// DefaultNodetoolService is the implementation of NodetoolService, and has the dependency on server-side port, i.e. Nodetool interface
type DefaultNodetoolService struct {
	nodetool domain.Nodetool
}

func (n DefaultNodetoolService) NodetoolStatus() ([]domain.NodetoolStatus, error) {
	return n.nodetool.RetrieveNodetoolStatus()
}

func NewNodetoolService(nodetool domain.Nodetool) DefaultNodetoolService {
	return DefaultNodetoolService{nodetool: nodetool}
}
