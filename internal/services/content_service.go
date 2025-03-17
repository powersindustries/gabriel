// Most of content service is temp code. Will be updated to pull content from database overtime.

package services

import (
	"email_poc/internal/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
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

func ContentAvailableForSending(contentId string) bool {
	return true
}

// ToDo: Pull raw content from database.
func GetEmailContentById(id string) ([]byte, error) {
	contentObject, err := getContentById(id)
	if err != nil {
		return nil, errors.New("failed to find content by id")
	}

	subject := "Subject: " + contentObject.Title + "\r\n"
	mime := "MIME-version: 1.0;\r\nContent-Type: text/plain; charset=\"UTF-8\";\r\n"
	body := `
		<html>
		<body>
			<h2>Welcome!</h2>
			<p>Thank you for subscribing to our newsletter.</p>
			<p><strong>Enjoy the latest updates!</strong></p>
		</body>
		</html>
	`

	return []byte(subject + mime + "\r\n" + body), nil
}

func GetNewsletterIdByContentId(id string) ([]string, error) {
	contentArraySize := len(contentArray)
	for x := 0; x < contentArraySize; x++ {
		currContent := contentArray[x]
		if id == currContent.Id {
			return currContent.NewsletterId, nil
		}
	}

	return nil, errors.New("failed to find content by id")
}

func getContentById(id string) (*models.Content, error) {
	contentArraySize := len(contentArray)
	for x := 0; x < contentArraySize; x++ {
		currContent := contentArray[x]
		if id == currContent.Id {
			return &currContent, nil
		}
	}

	return nil, errors.New("failed to find content by id")
}
