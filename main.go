package main

import (
	"github.com/alireza-fa/blog-go/src/constants"
	"github.com/alireza-fa/blog-go/src/data/db"
	"github.com/alireza-fa/blog-go/src/management"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err := godotenv.Load(".env")

	if os.Getenv(constants.DEBUG) == "" {
		err := os.Setenv(constants.BlogDbHost, "localhost")
		if err != nil {
			panic(err)
		}
	}

	if err != nil {
		panic("Error loading .env file")
	}

	err = db.InitDb()
	if err != nil {
		panic("Error connection database: " + err.Error())
	}
}

func main() {
	management.CreateWebServer()
}
