package routers

import (
	"github.com/alireza-fa/blog-go/src/api/handlers"
	"github.com/go-chi/chi"
)

func UserRouters(router chi.Router) {
	handler := handlers.NewUserFrontHandler()

	router.Post("/register/", handler.UserRegister)
	router.Post("/verify/", handler.UserVerify)
	router.Post("/login/", handler.UserLogin)
}
