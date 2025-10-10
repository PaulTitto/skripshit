package controllers

import (
	errWrap "backend/common/error"
	"backend/common/response"
	"backend/domain/dto"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	service services.IServiceRegistry
}

type IUserController interface {
	Login(*gin.Context)
	Register(*gin.Context)
	Update(*gin.Context)
	GetUserLogin(*gin.Context)
	GetUserByUUID(*gin.Context)
}

func NewUserController(service services.IServiceRegistry) IUserController {
	return &UserController{service: service}
}

func (u UserController) Login(context *gin.Context) {
	request := &dto.LoginRequest{}

	err := context.ShouldBindBodyWithJSON(request)
	if err != nil {
		response.HttpResponse(response.ParamsHTTPRes{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  context,
		})
		return
	}

	validate := validator.New()

	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errWrap.ErrValidation(err)
		response.HttpResponse(response.ParamsHTTPRes{
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errResponse,
			Gin:     context,
		})
	}

	user, err := u.service.GetUser().Login(context, request)
	if err != nil {
		response.HttpResponse(response.ParamsHTTPRes{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  context,
		})
		return
	}

	response.HttpResponse(response.ParamsHTTPRes{
		Code: http.StatusOK,
		Data: user.User,
		Gin:  context,
	})
}

func (u UserController) Register(context *gin.Context) {
	request := &dto.RegisterRequest{}

	err := context.ShouldBindBodyWithJSON(request)
	if err != nil {
		response.HttpResponse(response.ParamsHTTPRes{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  context,
		})
		return
	}

	validate := validator.New()

	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errWrap.ErrValidation(err)
		response.HttpResponse(response.ParamsHTTPRes{
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errResponse,
			Gin:     context,
		})
	}

	user, err := u.service.GetUser().Register(context, request)
	if err != nil {
		response.HttpResponse(response.ParamsHTTPRes{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  context,
		})
		return
	}

	response.HttpResponse(response.ParamsHTTPRes{
		Code: http.StatusOK,
		Data: user.User,
		Gin:  context,
	})
}

func (u UserController) Update(context *gin.Context) {
	request := &dto.UpdateRequest{}
	uuid := context.Param("id")

	err := context.ShouldBindBodyWithJSON(request)
	if err != nil {
		response.HttpResponse(response.ParamsHTTPRes{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  context,
		})
		return
	}

	validate := validator.New()

	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errWrap.ErrValidation(err)
		response.HttpResponse(response.ParamsHTTPRes{
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errResponse,
			Gin:     context,
		})
	}

	user, err := u.service.GetUser().Update(context, request, uuid)
	if err != nil {
		response.HttpResponse(response.ParamsHTTPRes{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  context,
		})
		return
	}

	response.HttpResponse(response.ParamsHTTPRes{
		Code: http.StatusOK,
		Data: user,
		Gin:  context,
	})
}

func (u UserController) GetUserLogin(context *gin.Context) {
	user, err := u.service.GetUser().GetUserLogin(context.Request.Context())
	if err != nil {
		response.HttpResponse(response.ParamsHTTPRes{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  context,
		})
		return
	}

	response.HttpResponse(response.ParamsHTTPRes{
		Code: http.StatusOK,
		Data: user,
		Gin:  context,
	})
}

func (u UserController) GetUserByUUID(context *gin.Context) {
	user, err := u.service.GetUser().GetUserByUUID(context.Request.Context(), context.Param("uuid"))
	if err != nil {
		response.HttpResponse(response.ParamsHTTPRes{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  context,
		})
		return
	}

	response.HttpResponse(response.ParamsHTTPRes{
		Code: http.StatusOK,
		Data: user,
		Gin:  context,
	})
}
