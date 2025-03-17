// Most of newsletter service is temp code. Will be updated to pull content from database overtime.

package services

import (
	"email_poc/internal/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

var newsletterArray []models.Newsletter

// ToDo: Update for pulling data from database.
func InitializeNewsletterService() {
	file, err := os.Open("internal/dummy/newsletter.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	erro := json.Unmarshal([]byte(fileContent), &newsletterArray)
	if erro != nil {
		log.Fatal(erro)
	}
}

func GetNewsletterObjectById(id string) (*models.Newsletter, error) {
	newsletterArraySize := len(newsletterArray)
	for x := 0; x < newsletterArraySize; x++ {
		currNewsletter := newsletterArray[x]
		if id == currNewsletter.Id {
			return &currNewsletter, nil
		}
	}

	return nil, errors.New("failed to find content by id")
}
