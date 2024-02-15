package logging

import (
	"github.com/alireza-fa/blog-go/src/constants"
	"log"
	"os"
)

type Logger interface {
	Init()

	Debug(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{})

	Info(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{})

	Warn(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{})

	Error(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{})

	Fatal(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{})
}

func NewLogger() Logger {
	log.Println("logger name:", os.Getenv(constants.Logger))

	switch os.Getenv(constants.Logger) {
	case "seq":
		return NewSeqLog()
	default:
		panic("setting LOGGER in .env")
	}
}
