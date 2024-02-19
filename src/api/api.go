package api

import (
	"fmt"
	"github.com/alireza-fa/blog-go/docs"
	"github.com/alireza-fa/blog-go/src/api/middlewares"
	"github.com/alireza-fa/blog-go/src/api/routers"
	"github.com/alireza-fa/blog-go/src/constants"
	"github.com/alireza-fa/blog-go/src/pkg/logging"
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"time"
)

func InitialServer(logger logging.Logger) {
	var router chi.Router = chi.NewRouter()

	router.Use(middlewares.LogMiddleware)

	router.Route("/api/", Routers)

	RegisterSwagger(router, logger)

	server := http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv(constants.ServerPort)),
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	log.Fatal(server.ListenAndServe())
}

func Routers(router chi.Router) {
	// User
	router.Route("/users/", routers.UserRouters)
	router.Route("/token/", routers.TokenRouter)

	// Post
	router.Route("/categories/", routers.CategoryRouter)
}

func RegisterSwagger(router chi.Router, logger logging.Logger) {
	docs.SwaggerInfo.Title = "blog go"
	docs.SwaggerInfo.Description = "blog og web server"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", os.Getenv(constants.ServerPort))
	docs.SwaggerInfo.Schemes = []string{"http"}

	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", os.Getenv(constants.ServerPort)))))
}
