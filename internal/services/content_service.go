// Most of content service is temp code. Will be updated to pull content from database overtime.

package services

import (
	"bytes"
	"email_poc/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/yuin/goldmark"
)

var contentArray []models.Content

// ToDo: Update for pulling data from database.
func InitializeContentService() {
	file, err := os.Open("internal/s3/content.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	erro := json.Unmarshal([]byte(fileContent), &contentArray)
	if erro != nil {
		log.Fatal(erro)
	}
}

func GetContentObjectByUUId(contentUUId string) (*models.Content, error) {
	contentArraySize := len(contentArray)
	for x := 0; x < contentArraySize; x++ {
		currContent := contentArray[x]

		if currContent.UUId == contentUUId {
			return &currContent, nil
		}
	}

	return nil, errors.New("failed to find content by UUId")
}

func GetEmailContentByContentUUId(contentUUId string) ([]byte, error) {
	// Get content object.
	var contentObject *models.Content

	contentArraySize := len(contentArray)
	for x := 0; x < contentArraySize; x++ {
		if contentArray[x].UUId == contentUUId {
			contentObject = &contentArray[x]
			break
		}
	}

	if contentObject == nil {
		return nil, errors.New("failed to find content by id")
	}

	rawContent, err := getRawContentByObject(contentObject)
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

// ToDo: Replace with actual call to s3 bucket. Below is temp/debug logic.
func getRawContentByObject(contentObject *models.Content) ([]byte, error) {

	path := "internal/s3/" + contentObject.NewsletterUUId + "/" + contentObject.UUId + contentObject.Type
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("failed to load raw content by id")
	}

	return data, nil
}
