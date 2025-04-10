package services

import (
	"email_poc/internal/config"
	"errors"
	"log/slog"
	"net/smtp"
	"time"
)

type EmailSendingService struct {
	contentService    *ContentService
	newsletterService *NewsletterService

	smtpHost         string
	smtpPort         string
	fromAddress      string
	password         string
	authentification smtp.Auth
}

func CreateNewEmailSendingService(contentService *ContentService, newsletterService *NewsletterService) *EmailSendingService {
	outputObject := &EmailSendingService{}

	outputObject.contentService = contentService
	outputObject.newsletterService = newsletterService

	outputObject.smtpHost = "smtp.gmail.com"
	outputObject.smtpPort = "587"
	outputObject.fromAddress = config.GetEnvVariables("email_user")
	if len(outputObject.fromAddress) == 0 {
		slog.Error("Failed to get the host from address.")
		return nil
	}

	outputObject.password = config.GetEnvVariables("email_password")
	if len(outputObject.password) == 0 {
		slog.Error("Failed to get the password from address.")
		return nil
	}

	outputObject.authentification = smtp.PlainAuth("", outputObject.fromAddress, outputObject.password, outputObject.smtpHost)

	return outputObject
}

// Sends out a content's email by content UUId.
func (this *EmailSendingService) SendEmailByContentUUId(contentUUId string) error {
	var err error

	// Get content object.
	contentObject, err := this.contentService.GetContentObjectByUUId(contentUUId)
	if err != nil {
		slog.Error("Failed to get the content object from scheduler service.")
		return errors.New("failed to get the content object from scheduler service")
	}

	// Dont send content emails if not ready to send.
	currUnixTime := time.Now().Unix()
	if currUnixTime >= contentObject.ReleaseDate {

		// Get email content.
		emailContent, err := this.contentService.GetEmailContentByContentUUId(contentUUId)
		if err != nil {
			slog.Error("Failed to get email content by content UUId:", err)
			return errors.New("failed to get email content by content UUId")
		}

		// Get email subscriber list.
		subscriberEmailList := this.newsletterService.GetNewsletterSubscriberEmailsByNewsletterUUId(contentObject.NewsletterUUId)
		if len(subscriberEmailList) > 0 {

			// Send email.
			err = smtp.SendMail(this.smtpHost+":"+this.smtpPort, this.authentification, this.fromAddress, subscriberEmailList, emailContent)
			if err != nil {
				slog.Error("Failed to send email:", err)
				return errors.New("failed to send email")
			}
		}
	}
	return nil
}
