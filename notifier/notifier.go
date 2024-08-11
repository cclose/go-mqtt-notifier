package notifier

import (
	"github.com/sirupsen/logrus"
)

type NotificationType string

const (
	FrigateEvent NotificationType = "frigate_event"
)

type NotificationService struct {
	smtp *SMTPNotifier
}

func NewNotificationService() (*NotificationService, error) {
	smtp, err := NewSMTPNotifier()
	if err != nil {
		return nil, err
	}

	return &NotificationService{smtp: smtp}, nil
}

func (ns *NotificationService) SendNotification(notificationType NotificationType, message string) {
	logNotification(message)
	err := ns.smtp.Notify("Event Detected", message)
	if err != nil {
		logrus.Errorf("Failed to send notification to SMTP: %v", err)
	}

	//if dispatcher, found := notificationDispatchTable[notificationType]; found {
	//	dispatcher(message)
	//} else {
	//	logrus.Warnf("No dispatcher found for notification type: %s", notificationType)
	//}
}

func logNotification(message string) {
	logrus.Info("Notification: " + message)
}
