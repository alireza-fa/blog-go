package management

import (
	"log"
	"net/http"
	"time"
)

func CreateWebServer() {
	mux := http.NewServeMux()

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	log.Fatal(server.ListenAndServe())
}
