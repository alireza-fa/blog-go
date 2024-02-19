package routers

import (
	"github.com/alireza-fa/blog-go/src/api/handlers"
	"github.com/alireza-fa/blog-go/src/api/middlewares"
	"github.com/go-chi/chi"
)

func UserRouters(router chi.Router) {
	handler := handlers.NewUserFrontHandler()

	router.Post("/register/", handler.UserRegister)
	router.Post("/verify/", handler.UserVerify)
	router.Post("/login/", handler.UserLogin)

	router.Group(authRouter)
}

func authRouter(router chi.Router) {
	handler := handlers.NewUserFrontHandler()

	router.Use(middlewares.Authentication)

	router.Get("/profile/", middlewares.Authorization(handler.UserProfile, []string{"default"}))
	router.Patch("/profile/update/", middlewares.Authorization(handler.UserProfileUpdate, []string{"default"}))
}
