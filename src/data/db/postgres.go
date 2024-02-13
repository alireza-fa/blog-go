package db

import (
	"fmt"
	"github.com/alireza-fa/blog-go/src/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

var dbClient *gorm.DB

func InitDb() error {
	dbPort, err := strconv.Atoi(os.Getenv(constants.BlogDbPort))

	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		os.Getenv(constants.BlogDbHost), dbPort, os.Getenv(constants.BlogDbUser),
		os.Getenv(constants.BlogDbPassword), os.Getenv(constants.BlogDbName), os.Getenv(constants.BlogDbSslMode))

	dbClient, err = gorm.Open(postgres.Open(conn))
	if err != nil {
		return err
	}

	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}

	maxIdleConns, err := strconv.Atoi(os.Getenv(constants.MaxIdleCONNS))
	if err != nil {
		panic("error while set maxIdleConns")
	}

	maxOpenConns, err := strconv.Atoi(os.Getenv(constants.MaxOpenConns))
	if err != nil {
		panic("error while set maxOpenConns")
	}

	connMaxLifetime, err := strconv.Atoi(os.Getenv(constants.ConnMaxLifetime))
	if err != nil {
		panic("error while set connMaxLifetime")
	}

	sqlDb.SetMaxIdleConns(maxIdleConns)
	sqlDb.SetMaxOpenConns(maxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Minute)

	return nil
}

func GetDb() *gorm.DB {
	return dbClient
}

func CloseDb() {
	conn, _ := dbClient.DB()
	conn.Close()
}
