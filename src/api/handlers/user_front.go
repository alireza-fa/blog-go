package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alireza-fa/blog-go/src/api/dto"
	"github.com/alireza-fa/blog-go/src/api/helper"
	"github.com/alireza-fa/blog-go/src/services"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UserFrontHandler struct {
	service *services.UserFrontService
}

func NewUserFrontHandler() *UserFrontHandler {
	return &UserFrontHandler{
		service: services.NewUserFrontService(),
	}
}

func (handler *UserFrontHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handler.CreateUser(w, r)
	default:
		helper.BaseResponseWithError(w, nil, http.StatusMethodNotAllowed, fmt.Errorf("method %s not allowed", r.Method))
	}
}

func (handler *UserFrontHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUser

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.BaseResponseWithError(w, nil, http.StatusNotAcceptable, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		helper.BaseResponseWithValidationError(w, nil, http.StatusBadRequest, err)
		return
	}

	err = handler.service.CreateUser(user)
	if err != nil {
		helper.BaseResponseWithError(w, nil, http.StatusNotAcceptable, err)
		return
	}

	helper.BaseResponse(w, nil, http.StatusCreated)
}

func (handler *UserFrontHandler) VerifyUser(w http.ResponseWriter, r *http.Request) {
	var userCreate dto.UserVerify

	err := json.NewDecoder(r.Body).Decode(&userCreate)
	if err != nil {
		helper.BaseResponseWithError(w, nil, http.StatusNotAcceptable, err)
		return
	}

	validate := validator.New()
	if err = validate.Struct(userCreate); err != nil {
		helper.BaseResponseWithValidationError(w, nil, http.StatusBadRequest, err)
		return
	}

	user, err := handler.service.VerifyUser(&userCreate)
	if err != nil {
		helper.BaseResponseWithError(w, nil, http.StatusNotAcceptable, errors.New("invalid code"))
		return
	}

	helper.BaseResponse(w, user, http.StatusOK)
}
