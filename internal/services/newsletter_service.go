package services

import (
	"email_poc/internal/repository"
)

type NewsletterService struct {
	newsletterRepository repository.NewsletterRepository
	subscriberService    *SubscriberService
}

func CreateNewNewsletterService(repository repository.NewsletterRepository, subscriberService *SubscriberService) *NewsletterService {
	return &NewsletterService{
		newsletterRepository: repository,
		subscriberService:    subscriberService,
	}
}

func (this *NewsletterService) GetNewsletterSubscriberEmailsByNewsletterUUId(newsletterUUId string) []string {

	newsletterObject, err := this.newsletterRepository.GetNewsletterByUUId(newsletterUUId)
	if err != nil {
		println("Failed to find the newsletter UUID with the id: " + newsletterUUId)
		return nil
	}

	subscriberListSize := len(newsletterObject.SubscriberList)
	if subscriberListSize == 0 {
		println("Newsletter UUID: " + newsletterUUId + " has 0 subscibers.")
		return nil
	}

	var outputEmail []string

	for x := 0; x < subscriberListSize; x++ {
		currEmail := this.subscriberService.GetSubscriberEmailByUUId(newsletterObject.SubscriberList[x])
		if len(currEmail) > 0 {
			outputEmail = append(outputEmail, currEmail)
		}
	}

	return outputEmail
}
