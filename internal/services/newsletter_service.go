// Most of newsletter service is temp code. Will be updated to pull content from database overtime.

package services

import (
	"email_poc/internal/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var newsletterArray []models.Newsletter

// ToDo: Update for pulling data from database.
func InitializeNewsletterService() {
	file, err := os.Open("internal/s3/newsletter.json")
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

func GetNewsletterSubscribersByNewsletterUUId(newsletterUUId string) []string {
	newsletterArraySize := len(newsletterArray)
	for x := 0; x < newsletterArraySize; x++ {
		currNewsletter := newsletterArray[x]
		if newsletterUUId == currNewsletter.UUId {
			var outputSubscriberList []string
			subscriberListSize := len(currNewsletter.SubscriberList)
			for y := 0; y < subscriberListSize; y++ {
				subscriberEmail := GetSubscriberEmailByUUId(currNewsletter.SubscriberList[y])
				if len(subscriberEmail) > 0 {
					outputSubscriberList = append(outputSubscriberList, subscriberEmail)
				}
			}

			return outputSubscriberList
		}
	}

	return nil
}
