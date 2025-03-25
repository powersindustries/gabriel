package services

import "email_poc/internal/repository"

func GetSubscriberEmailByUUId(uUId string) string {
	return repository.GetSubscriberEmailByUUId(uUId)
}
