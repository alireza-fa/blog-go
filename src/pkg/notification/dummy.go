package notification

import (
	"github.com/alireza-fa/blog-go/src/pkg/logging"
)

type DummyNotification struct{}

func NewDummyNotification() *DummyNotification {
	return &DummyNotification{}
}

func (dummy DummyNotification) Init() {}

func (dummy DummyNotification) Send(receiver string, extra map[string]string) {
	logger := logging.NewLogger()

	extraLog := map[logging.ExtraKey]interface{}{Receiver: receiver, logging.NotificationMessage: extra[Message]}
	logger.Info(logging.Notification, logging.SendNotification, extra[Message], extraLog)
}
