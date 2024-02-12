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

	dbClient, err := gorm.Open(postgres.Open(conn))
	if err != nil {
		return err
	}

	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}

	maxIdleConns, _ := strconv.Atoi(os.Getenv(constants.MaxIdleCONNS))
	maxOpenConns, _ := strconv.Atoi(os.Getenv(constants.MaxOpenConns))
	connMaxLifetime, _ := strconv.Atoi(os.Getenv(constants.ConnMaxLifetime))

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
