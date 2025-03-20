package services

import (
	"email_poc/internal/config"
	"errors"
	"fmt"
	"net/smtp"
	"time"
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

// Sends out a content's email by content UUId.
func SendEmailByContentUUId(contentUUId string) error {
	var err error

	// Get content object.
	contentObject, err := GetContentObjectByUUId(contentUUId)
	if err != nil {
		println("Failed to get the content object from scheduler service.")
		return errors.New("failed to get the content object from scheduler service")
	}

	// Dont send content emails if not ready to send.
	currUnixTime := time.Now().Unix()
	if currUnixTime >= contentObject.ReleaseDate {

		// Get email content.
		emailContent, err := GetEmailContentByContentUUId(contentUUId)
		if err != nil {
			fmt.Println("Failed to get email content by content UUId:", err)
			return errors.New("failed to get email content by content UUId")
		}

		// Get email userlist.
		userList := GetNewsletterUserlistByNewsletterUUId(contentObject.NewsletterUUId)
		if len(userList) > 0 {

			// Send email.
			err = smtp.SendMail(smtpHost+":"+smtpPort, authentification, fromAddress, userList, emailContent)
			if err != nil {
				fmt.Println("Failed to send email:", err)
				return errors.New("failed to send email")
			}
		}
	}
	return nil
}
