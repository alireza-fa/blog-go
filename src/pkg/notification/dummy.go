package notification

import (
	"github.com/alireza-fa/blog-go/src/pkg/logging"
)

type DummyNotification struct {
	logger logging.Logger
}

func NewDummyNotification() *DummyNotification {
	return &DummyNotification{
		logger: logging.NewLogger(),
	}
}

func (dummy DummyNotification) Init() {}

func (dummy DummyNotification) Send(receiver string, extra map[string]string) {

	extraLog := map[logging.ExtraKey]interface{}{Receiver: receiver, logging.NotificationMessage: extra[Message]}
	dummy.logger.Info(logging.Notification, logging.SendNotification, extra[Message], extraLog)
}
