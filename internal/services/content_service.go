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
	"time"

	"github.com/yuin/goldmark"
)

var contentArray []models.Content

// ToDo: Update for pulling data from database.
func InitializeContentService() {
	file, err := os.Open("internal/dummy/content.json")
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

// Look at the content in a newsletter and check that it is ready to be sent out.
func ContentShouldSendByNewsletterId(newsletterId string) bool {
	currUnixTime := time.Now().Unix()

	contentArraySize := len(contentArray)
	for x := 0; x < contentArraySize; x++ {
		currContent := contentArray[x]

		// Iterate through newsletters and find content for newsletter.
		newslettersArraySize := len(currContent.NewsletterId)
		for y := 0; y < newslettersArraySize; y++ {
			if newsletterId == currContent.NewsletterId[y] {
				if currUnixTime >= currContent.ReleaseDate {
					return true
				}
			}
		}
	}

	return false
}

func GetEmailContentByNewsletterId(newsletterId string) ([]byte, error) {
	contentArraySize := len(contentArray)
	for x := 0; x < contentArraySize; x++ {
		currContent := contentArray[x]

		// Iterate through newsletters and find content for newsletter.
		newslettersArraySize := len(currContent.NewsletterId)
		for y := 0; y < newslettersArraySize; y++ {
			if newsletterId == currContent.NewsletterId[y] {
				// Content found. Now get raw.
				rawContent, err := getRawContentByObject(&currContent)
				if err != nil {
					println("Failed to get the raw content from the content object.")
					return nil, errors.New("failed to get the raw content from the content object")
				}

				subject := "Subject: " + currContent.Title + "\r\n"
				mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
				if currContent.ContentType == ".txt" {
					mime = "MIME-version: 1.0;\r\nContent-Type: text/plain; charset=\"UTF-8\";\r\n"
				}

				if currContent.ContentType == ".md" {
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
		}
	}

	return nil, errors.New("failed to find content by id")
}

// ToDo: Replace with actual call to database. Below is temp/debug logic.
func getRawContentByObject(contentObject *models.Content) ([]byte, error) {

	path := "internal/dummy/content/" + contentObject.ContentId + contentObject.ContentType
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("failed to load raw content by id")
	}

	return data, nil
}
