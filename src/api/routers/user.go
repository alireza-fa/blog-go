package routers

import (
	"github.com/alireza-fa/blog-go/src/api/handlers"
	"github.com/alireza-fa/blog-go/src/api/middlewares"
	"github.com/alireza-fa/blog-go/src/pkg/logging"
	"net/http"
)

func UserRouters(mux *http.ServeMux) {
	handler := handlers.NewUserFrontHandler()

	mux.Handle("/users/", middlewares.LogMiddleware(handler, logging.NewLogger()))
}
