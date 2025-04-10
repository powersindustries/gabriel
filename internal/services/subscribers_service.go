package services

import "email_poc/internal/repository"

type SubscriberService struct {
	subscriberRepository repository.SubscriberRepository
}

func CreateNewSubscriberService(repository repository.SubscriberRepository) *SubscriberService {
	return &SubscriberService{subscriberRepository: repository}
}

func (subscriberServiceSelf *SubscriberService) GetSubscriberEmailByUUId(uUId string) string {
	return subscriberServiceSelf.subscriberRepository.GetSubscriberEmailByUUId(uUId)
}
