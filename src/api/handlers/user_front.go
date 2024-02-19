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
// @Summary register user
// @Description register user
// @Tags Users
// @Accept json
// @Produce json
// @Param Request body dto.CreateUser true "ًRegister User"
// @Success 201 {object} helper.BaseHttpResponse "register user"
// @Failure 400 {object} helper.BaseHttpResponseWithValidationError "bad request"
// @Failure 406 {object} helper.BaseHttpResponseWithError "error while register user"
// @Router /api/users/register/ [post]
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

// UserVerify godoc
// @Summary user verify account
// @Description user verify account
// @Tags Users
// @Accept json
// @Produce json
// @Param Request body dto.UserVerify true "User verify"
// @Success 200 {object} helper.BaseHttpResponse "user verified"
// @Failure 400 {object} helper.BaseHttpResponseWithValidationError "bad request"
// @Failure 406 {object} helper.BaseHttpResponseWithError "error while verifying user"
// @Router /api/users/verify/ [post]
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

// UserLogin godoc
// @Summary user login
// @Description user login
// @Tags Users
// @Accept json
// @Produce json
// @Param Request body dto.UserLogin true "user login"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.TokenDetail} "user login"
// @Failure 400 {object} helper.BaseHttpResponseWithValidationError "bad request"
// @Failure 406 {object} helper.BaseHttpResponseWithError "error while login user"
// @Router /api/users/login/ [post]
func (handler *UserFrontHandler) UserLogin(w http.ResponseWriter, r *http.Request) {
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
		helper.BaseResponseWithError(w, nil, http.StatusNotFound, err)
		return
	}

	helper.BaseResponse(w, user, http.StatusOK)
}

// UserProfile godoc
// @Summary user profile
// @Description user profile
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} helper.BaseHttpResponse{result=dto.Profile} "user profile info"
// @Failure 401 {object} helper.BaseHttpResponseWithError "UnAuthorization"
// @Router /api/users/profile/ [get]
// @Security AuthBearer
func (handler *UserFrontHandler) UserProfile(w http.ResponseWriter, r *http.Request) {
	var userProfile *dto.Profile = handler.service.UserProfile(r.Context())

	helper.BaseResponse(w, userProfile, http.StatusOK)
}

// UserProfileUpdate godoc
// @Summary user profile update
// @Description user profile update
// @Tags Users
// @Accept json
// @Produce json
// @Param Request body dto.ProfileUpdate true "user profile update"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.Profile} "user profile updated"
// @Failure 400 {object} helper.BaseHttpResponseWithValidationError "bad request"
// @Failure 406 {object} helper.BaseHttpResponseWithError "not acceptable"
// @Router /api/users/profile/update/ [patch]
// @Security AuthBearer
func (handler *UserFrontHandler) UserProfileUpdate(w http.ResponseWriter, r *http.Request) {
	var profileUpdate dto.ProfileUpdate

	err := json.NewDecoder(r.Body).Decode(&profileUpdate)
	if err != nil {
		helper.BaseResponseWithError(w, nil, http.StatusNotAcceptable, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(profileUpdate)
	if err != nil {
		helper.BaseResponseWithValidationError(w, nil, http.StatusBadRequest, err)
		return
	}

	profile, err := handler.service.UserProfileUpdate(r.Context(), profileUpdate)
	if err != nil {
		helper.BaseResponseWithError(w, nil, http.StatusNotAcceptable, err)
		return
	}

	helper.BaseResponse(w, profile, http.StatusOK)
}
