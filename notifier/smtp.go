package notifier

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

type SMTPNotifier struct {
	serverAddr string
	fromAddr   string
	toAddr     []string

	auth smtp.Auth
}

func NewSMTPNotifier() (*SMTPNotifier, error) {
	serverAddr := os.Getenv("SMTP_SERVER")
	if serverAddr == "" {
		return nil, fmt.Errorf("env SMTP_SERVER not found")
	}
	fromAddr := os.Getenv("SMTP_FROM")
	if fromAddr == "" {
		return nil, fmt.Errorf("env SMTP_FROM not found")
	}
	toAddr := os.Getenv("SMTP_TO")
	if toAddr == "" {
		return nil, fmt.Errorf("env SMTP_TO not found")
	}

	auth := smtp.PlainAuth("",
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
		"postfix",
	)

	return &SMTPNotifier{
		serverAddr: serverAddr,
		fromAddr:   fromAddr,
		toAddr:     strings.Split(toAddr, ","),
		auth:       auth,
	}, nil
}

func (n *SMTPNotifier) Notify(subject string, text string) error {
	for _, addr := range n.toAddr {
		err := n.SendEmail(addr, subject, text)
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *SMTPNotifier) SendEmail(to string, subject string, body string) error {
	msg := []byte("To: " + to + "\r\n" + "Subject: " + subject + "\r\n" + body)
	return smtp.SendMail(n.serverAddr, n.auth, n.fromAddr, []string{to}, msg)
}
