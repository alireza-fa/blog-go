package handlers

import (
	"encoding/json"
	"github.com/alireza-fa/blog-go/src/api/dto"
	"github.com/alireza-fa/blog-go/src/api/helper"
	"github.com/alireza-fa/blog-go/src/services"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type TokenHandler struct {
	service  *services.UserFrontService
	validate *validator.Validate
}

func NewTokenService() *TokenHandler {
	return &TokenHandler{
		service:  services.NewUserFrontService(),
		validate: validator.New(),
	}
}

// RefreshAccessToken godoc
// @Summary refresh access token
// @Description generate and get a new access token
// @Tags Token
// @Accept json
// @Produce json
// @Param Request body dto.RefreshToken true "user profile update"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.TokenDetail} "user profile updated"
// @Failure 400 {object} helper.BaseHttpResponseWithValidationError "bad request"
// @Failure 406 {object} helper.BaseHttpResponseWithError "not acceptable"
// @Router /api/token/refresh/ [post]
func (handler *TokenHandler) RefreshAccessToken(w http.ResponseWriter, r *http.Request) {
	var refreshToken dto.RefreshToken

	err := json.NewDecoder(r.Body).Decode(&refreshToken)
	if err != nil {
		helper.BaseResponseWithError(w, nil, http.StatusNotAcceptable, err)
		return
	}

	err = handler.validate.Struct(refreshToken)
	if err != nil {
		helper.BaseResponseWithValidationError(w, nil, http.StatusBadRequest, err)
		return
	}

	tokenDetail, err := handler.service.UserRefreshToken(refreshToken)
	if err != nil {
		helper.BaseResponseWithError(w, nil, http.StatusNotAcceptable, err)
		return
	}

	helper.BaseResponse(w, tokenDetail, http.StatusOK)
}
