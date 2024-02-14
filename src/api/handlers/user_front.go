package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/alireza-fa/blog-go/src/api/dto"
	"github.com/alireza-fa/blog-go/src/api/helper"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UserFrontHandler struct{}

func (handler UserFrontHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CreateUser(w, r)
	default:
		helper.BaseResponseWithError(w, nil, http.StatusMethodNotAllowed, fmt.Errorf("method %s not allowed", r.Method))
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	helper.BaseResponse(w, user, http.StatusCreated)
}

func UserVerify(w http.ResponseWriter, r *http.Request) {}
