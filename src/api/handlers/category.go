package handlers

import (
	"encoding/json"
	"github.com/alireza-fa/blog-go/src/api/dto"
	"github.com/alireza-fa/blog-go/src/api/helper"
	"github.com/alireza-fa/blog-go/src/services"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type CategoryHandler struct {
	service  *services.CategoryService
	validate *validator.Validate
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{
		service:  services.NewCategoryService(),
		validate: validator.New(),
	}
}

// Create godoc
// @Summary Create category
// @Description Create category. Only admin can create a new category
// @Tags Categories
// @Accept json
// @Produce json
// @Param Request body dto.CategoryCreate true "category create"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CategoryOutput} "created"
// @Failure 400 {object} helper.BaseHttpResponseWithValidationError "bad request"
// @Failure 406 {object} helper.BaseHttpResponseWithError "error while creating a new category"
// @Router /api/categories/ [post]
// @Security AuthBearer
func (handler *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var categoryCreate dto.CategoryCreate

	err := json.NewDecoder(r.Body).Decode(&categoryCreate)
	if err != nil {
		helper.BaseResponseWithError(w, nil, http.StatusNotAcceptable, err)
		return
	}

	err = handler.validate.Struct(&categoryCreate)
	if err != nil {
		helper.BaseResponseWithValidationError(w, nil, http.StatusBadRequest, err)
		return
	}

	category, err := handler.service.Create(&categoryCreate)
	if err != nil {
		helper.BaseResponseWithError(w, nil, http.StatusNotAcceptable, err)
		return
	}

	helper.BaseResponse(w, category, http.StatusCreated)
}
