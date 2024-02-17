package main

import (
	"github.com/alireza-fa/blog-go/src/api"
	"github.com/alireza-fa/blog-go/src/data/cache"
	"github.com/alireza-fa/blog-go/src/data/db"
	"github.com/alireza-fa/blog-go/src/data/db/migrations"
	"github.com/alireza-fa/blog-go/src/pkg/logging"
	"github.com/joho/godotenv"
)

var logger logging.Logger

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	logger = logging.NewLogger()
}

func main() {
	err := db.InitDb()
	if err != nil {
		panic("connection to postgres failed: " + err.Error())
	}
	logger.Info(logging.Postgres, logging.Startup, "connection to postgres", nil)
	defer db.CloseDb()

	migrations.Up1()

	err = cache.InitRedis()
	if err != nil {
		panic("connection to redis failed: " + err.Error())
	}
	defer cache.CloseRedis()

	api.InitialServer(logger)
}
