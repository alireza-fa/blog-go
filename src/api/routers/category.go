package routers

import (
	"github.com/alireza-fa/blog-go/src/api/handlers"
	"github.com/alireza-fa/blog-go/src/api/middlewares"
	"github.com/go-chi/chi"
)

func CategoryRouter(router chi.Router) {
	handler := handlers.NewCategoryHandler()

	router.Use(middlewares.Authentication)

	router.Post("/", middlewares.Authorization(handler.Create, []string{"admin"}))
	router.Patch("/", middlewares.Authorization(handler.Update, []string{"admin"}))
	router.Get("/", handler.GetCategory)
}
