package services

import (
	"backend/repositories"
	services "backend/services/user"
)

type Registry struct {
	repository repositories.IRepositoryRegistry
}

func (r Registry) GetUser() services.IUserService {
	return services.NewUserService(r.repository)
}

type IServiceRegistry interface {
	GetUser() services.IUserService
}

func NewServiceRegistry(repository repositories.IRepositoryRegistry) IServiceRegistry {
	return &Registry{repository: repository}
}
