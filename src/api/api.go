package api

import (
	"github.com/alireza-fa/blog-go/src/api/routers"
	"log"
	"net/http"
	"time"
)

func InitialServer() {
	var mux *http.ServeMux = http.NewServeMux()

	Routers(mux)

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	log.Fatal(server.ListenAndServe())
}

func Routers(mux *http.ServeMux) {
	routers.UserRouters(mux)
}
