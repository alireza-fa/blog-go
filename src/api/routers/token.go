package routers

import (
	"github.com/alireza-fa/blog-go/src/api/handlers"
	"github.com/go-chi/chi"
)

func TokenRouter(router chi.Router) {
	handler := handlers.NewTokenService()

	router.Post("/refresh/", handler.RefreshAccessToken)
}
