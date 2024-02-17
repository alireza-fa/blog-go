package routers

import (
	"github.com/alireza-fa/blog-go/src/api/handlers"
	"net/http"
)

func UserRouters(mux *http.ServeMux) {
	handler := handlers.NewUserFrontHandler()

	mux.Handle("/users/", handler)

	mux.HandleFunc("/users/verify/", handler.VerifyUser)
}
