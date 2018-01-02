package model

import (
	"github.com/SlyMarbo/gmail"
	"os"
)

func SendMail(address string,subject string, body string)(error){
	email := gmail.Compose(subject, body)
	email.From = os.Getenv("GMAIL_ADDRESS")
	email.Password = os.Getenv("GMAIL_PASSWORD")

	// Defaults to "text/plain; charset=utf-8" if unset.
	email.ContentType = "text/html; charset=utf-8"

	// Normally you'll only need one of these, but I thought I'd show both.
	email.AddRecipient(address)

	err := email.Send()
	if err != nil {
		return err
	}
	return nil
}