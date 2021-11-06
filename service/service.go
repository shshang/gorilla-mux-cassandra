package service

import (
	"github.com/shshang/gorilla-mux-cassandra/domain"
	"github.com/shshang/gorilla-mux-cassandra/errs"
)

//NodetoolService is the primary port on user side
type NodetoolService interface {
	NodetoolStatusService() ([]domain.NodetoolStatus, *errs.AppError)
}

// DefaultNodetoolService is the implementation of NodetoolService, and has the dependency on server-side port, i.e. Nodetool interface
type DefaultNodetoolService struct {
	nodetool domain.Nodetool
}

func (n DefaultNodetoolService) NodetoolStatusService() ([]domain.NodetoolStatus, *errs.AppError) {
	return n.nodetool.RetrieveNodetoolStatus()
}

func NewNodetoolService(nodetool domain.Nodetool) DefaultNodetoolService {
	return DefaultNodetoolService{nodetool: nodetool}
}
