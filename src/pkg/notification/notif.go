package notification

import (
	"github.com/alireza-fa/blog-go/src/constants"
	"github.com/alireza-fa/blog-go/src/pkg/logging"
	"os"
)

type Notifier interface {
	Init()

	Send(receiver string, extra map[string]string)
}

func NewNotifier() Notifier {
	logger := logging.NewLogger()

	switch os.Getenv(constants.Notification) {
	case "dummy":
		return NewDummyNotification()
	default:
		logger.Fatal(logging.Notification, logging.Startup, "notification service not found", nil)
		return nil
	}
}
