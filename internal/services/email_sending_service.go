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

func SendEmailByContentId(content string) error {
	emailContent, err := GetEmailContentById(content)
	if err != nil {
		return errors.New("failed to find content by id")
	}

	if emailContent != nil {
		// Get users.
		newsletterIds, err := GetNewsletterIdByContentId(content)
		if err != nil {
			return errors.New("failed to get newsletter ids")
		}

		var toEmails []string

		newsletterIdsSize := len(newsletterIds)
		for x := 0; x < newsletterIdsSize; x++ {
			emails := GetAllUsersByNewsletterId(newsletterIds[x])
			for y := 0; y < len(emails); y++ {
				toEmails = append(toEmails, emails[y])
			}
		}

		erro := smtp.SendMail(smtpHost+":"+smtpPort, authentification, fromAddress, toEmails, emailContent)
		if erro != nil {
			fmt.Println("Failed to send email:", erro)
			return errors.New("failed to send email")
		}
	}

	return nil
}
