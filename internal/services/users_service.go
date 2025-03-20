// Most of users service is temp code. Will be updated to pull content from database overtime.

package services

import (
	"email_poc/internal/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var usersArray []models.User

// ToDo: Update for pulling data from database.
func InitializeUsers() {
	file, err := os.Open("internal/s3/users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	erro := json.Unmarshal([]byte(fileContent), &usersArray)
	if erro != nil {
		log.Fatal(erro)
	}
}

func GetUserEmailByUUId(uUId string) string {
	userArraySize := len(usersArray)
	for x := 0; x < userArraySize; x++ {
		currUser := usersArray[x]
		if uUId == currUser.UUId {
			return currUser.Email
		}
	}

	return ""
}

func GetAllUsersByNewsletterId(newsletterId string) []string {
	newsLetter, err := GetNewsletterObjectById(newsletterId)
	if err != nil {
		log.Fatal(err)
	}

	var output []string

	newsletterUserSize := len(newsLetter.UserList)
	for x := 0; x < newsletterUserSize; x++ {
		output = append(output, newsLetter.UserList[x])
	}

	return output
}
