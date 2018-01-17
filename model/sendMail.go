package model

import (
	"os"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
)

func SendMail(address string, subject string, body string) error {
	from := mail.NewEmail("MultiLogin", "azusa@azunyan1111.com")
	to := mail.NewEmail("New User", address)
	plainTextContent := body
	htmlContent := body
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)
	if err != nil {
		return err
	}
	return nil
}
