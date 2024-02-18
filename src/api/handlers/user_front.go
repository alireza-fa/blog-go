package handlers

import (
	"encoding/json"
	"errors"
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

func (handler *UserFrontHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

// UserRegister godoc
// @Summary Send otp to user
// @Description Send otp to user
// @Tags Users
// @Accept json
// @Produce json
// @Param Request body dto.CreateUser true "GetOtpRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /users/register [post]
func (handler *UserFrontHandler) UserRegister(w http.ResponseWriter, r *http.Request) {
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

func (handler *UserFrontHandler) UserVerify(w http.ResponseWriter, r *http.Request) {
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

func (handler UserFrontHandler) UserLogin(w http.ResponseWriter, r *http.Request) {
	var userLogin dto.UserLogin

	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		helper.BaseResponseWithError(w, nil, http.StatusNotAcceptable, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(userLogin); err != nil {
		helper.BaseResponseWithValidationError(w, nil, http.StatusBadRequest, err)
		return
	}

	user, err := handler.service.UserLogin(userLogin)
	if err != nil {
		//helper.BaseResponseWithError(w, nil, http.StatusNotFound, errors.New("user with this information not found"))
		helper.BaseResponseWithError(w, nil, http.StatusNotFound, err)
		return
	}

	helper.BaseResponse(w, user, http.StatusOK)
}
