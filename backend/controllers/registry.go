package controllers

import (
	controllers "backend/controllers/user"
	"backend/services"
)

type Registry struct {
	service services.IServiceRegistry
}

func (r *Registry) GetUserController() controllers.IUserController {
	return controllers.NewUserController(r.service)
}

type IControllerRegistry interface {
	GetUserController() controllers.IUserController
}

func NewControllerRegistry(service services.IServiceRegistry) IControllerRegistry {
	return &Registry{service: service}
}
