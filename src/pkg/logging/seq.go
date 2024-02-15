package logging

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/alireza-fa/blog-go/src/constants"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	NewEVENT = "/api/events/raw"
)

var once sync.Once

var seqLogger *SeqLog

type SeqLog struct {
	ApiKey  string
	BaseURl string
	Port    string
	BaseUrl string
}

type EventProperty struct {
	Level           string `json:"Level"`
	MessageTemplate string `json:"MessageTemplate"`
	Timestamp       string `json:"Timestamp"`
	Properties      map[ExtraKey]interface{}
}

func NewSeqLog() *SeqLog {
	once.Do(func() {
		logger := &SeqLog{
			ApiKey:  os.Getenv(constants.SeqApiKey),
			BaseURl: os.Getenv(constants.BaseUrl),
			Port:    os.Getenv(constants.SeqPort),
			BaseUrl: fmt.Sprintf("http://%s:%s", os.Getenv(constants.BaseUrl), os.Getenv(constants.SeqPort)),
		}

		seqLogger = logger
	})
	return seqLogger
}

func (logger SeqLog) Init() {}

func (logger SeqLog) Debug(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{}) {
	go logger.CreateNewEvent(cat, sub, LevelDebug, message, extra)
}

func (logger SeqLog) Info(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{}) {
	go logger.CreateNewEvent(cat, sub, LevelInfo, message, extra)
}

func (logger SeqLog) Warn(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{}) {
	go logger.CreateNewEvent(cat, sub, LevelWarn, message, extra)
}

func (logger SeqLog) Error(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{}) {
	go logger.CreateNewEvent(cat, sub, LevelError, message, extra)
}

func (logger SeqLog) Fatal(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{}) {
	go logger.CreateNewEvent(cat, sub, LevelFatal, message, extra)
	<-time.After(time.Second * 1)
	log.Fatal(message)
}

func (logger SeqLog) CreateNewEvent(cat Category, sub SubCategory, level string, message string, extra map[ExtraKey]interface{}) {
	timestamp := time.Now().Format(time.RFC3339)

	if extra != nil {
		extra["Category"] = cat
		extra["SubCategory"] = sub
	} else {
		extra = map[ExtraKey]interface{}{
			"Category":    cat,
			"SubCategory": sub,
		}
	}

	event := struct {
		Events []EventProperty
	}{
		Events: []EventProperty{
			{
				Level:           level,
				MessageTemplate: message,
				Timestamp:       timestamp,
				Properties:      extra,
			},
		},
	}

	eventJson, err := json.Marshal(event)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest(http.MethodPost, logger.BaseUrl+NewEVENT, bytes.NewReader(eventJson))
	if err != nil {
		panic(err)
	}

	request.Header.Add("Api-Key", logger.ApiKey)

	_, err = http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
}
