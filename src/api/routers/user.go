package routers

import (
	"github.com/alireza-fa/blog-go/src/api/handlers"
	"github.com/go-chi/chi"
)

func UserRouters(router chi.Router) {
	handler := handlers.NewUserFrontHandler()

	router.Post("/users/register/", handler.UserRegister)
	router.Post("/users/verify/", handler.UserVerify)
	router.Post("/users/login/", handler.UserLogin)
}
