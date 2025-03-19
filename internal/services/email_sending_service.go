package services

import (
	"email_poc/internal/config"
	"errors"
	"fmt"
	"net/smtp"
)

var smtpHost = "smtp.gmail.com"
var smtpPort = "587"

var fromAddress string
var password string

var authentification smtp.Auth

func InitializeEmailSendingService() {
	var err error

	fromAddress, err = config.GetEnvVariables("user")
	if err != nil {
		println("Failed to get the host from address.")
		return
	}

	password, err = config.GetEnvVariables("password")
	if err != nil {
		println("Failed to get the password from address.")
		return
	}

	authentification = smtp.PlainAuth("", fromAddress, password, smtpHost)
}

// Sends out a newsletter's email by newsletter id.
func SendNewsletterEmail(newsletterId string) error {
	var err error

	// Get newsletter object.
	newsletterObject, err := GetNewsletterObjectById(newsletterId)
	if err != nil {
		println("Failed to get the newsletter object from scheduler service.")
		return errors.New("failed to get the newsletter object from scheduler service")
	}

	// Make sure content is ready to be sent out. Dont sent emails if not ready to send.
	if !ContentShouldSendByNewsletterId(newsletterId) {
		return nil
	}

	// Get email content.
	emailContent, err := GetEmailContentByNewsletterId(newsletterId)

	// Send email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, authentification, fromAddress, newsletterObject.UserList, emailContent)
	if err != nil {
		fmt.Println("Failed to send email:", err)
		return errors.New("failed to send email")
	}
	return nil
}
