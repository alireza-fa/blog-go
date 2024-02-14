package routers

import (
	"github.com/alireza-fa/blog-go/src/api/handlers"
	"net/http"
)

func UserRouters(mux *http.ServeMux) {
	mux.Handle("/users/", handlers.UserFrontHandler{})

	mux.HandleFunc("/users/verify/", handlers.UserVerify)
}
