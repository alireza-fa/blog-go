package routers

import (
	"github.com/alireza-fa/blog-go/src/api/handlers"
	"net/http"
)

func UserRouters(mux *http.ServeMux) {
	handler := handlers.NewUserFrontHandler()

	mux.HandleFunc("/users/register/", handler.UserRegister)
	mux.HandleFunc("/users/verify/", handler.UserVerify)
	mux.HandleFunc("/users/login/", handler.UserLogin)
}
