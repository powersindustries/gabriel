package services

import (
	"bytes"
	"email_poc/internal/models"
	"email_poc/internal/repository"
	"errors"
	"fmt"

	"github.com/yuin/goldmark"
)

func GetContentObjectByUUId(contentUUId string) (*models.Content, error) {
	return repository.GetContentObjectByUUId(contentUUId)
}

func GetEmailContentByContentUUId(contentUUId string) ([]byte, error) {
	contentObject, err := GetContentObjectByUUId(contentUUId)
	if contentObject == nil || err != nil {
		return nil, errors.New("failed to find content by id")
	}

	rawContent, err := repository.GetRawContentByObject(contentObject)
	if err != nil {
		println("Failed to get the raw content from the content object.")
		return nil, errors.New("failed to get the raw content from the content object")
	}

	subject := "Subject: " + contentObject.Title + "\r\n"
	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	if contentObject.Type == ".txt" {
		mime = "MIME-version: 1.0;\r\nContent-Type: text/plain; charset=\"UTF-8\";\r\n"
	}

	if contentObject.Type == ".md" {
		var buf bytes.Buffer
		err := goldmark.Convert(rawContent, &buf)
		if err != nil {
			fmt.Println("Error:", err)
			return nil, errors.New("failed to markdown to html")
		}

		return []byte(subject + mime + "\r\n" + buf.String()), nil
	}

	return []byte(subject + mime + "\r\n" + string(rawContent)), nil
}
