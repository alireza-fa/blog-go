package main

import (
	"github.com/alireza-fa/blog-go/src/api"
	"github.com/alireza-fa/blog-go/src/data/cache"
	"github.com/alireza-fa/blog-go/src/data/db"
	"github.com/alireza-fa/blog-go/src/data/db/migrations"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	err = db.InitDb()
	if err != nil {
		panic("connection to postgres failed: " + err.Error())
	}
	defer db.CloseDb()

	migrations.Up1()

	err = cache.InitRedis()
	if err != nil {
		panic("connection to redis failed: " + err.Error())
	}
	defer cache.CloseRedis()
}

func main() {
	api.InitialServer()
}
