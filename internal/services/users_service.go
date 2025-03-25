package services

import "email_poc/internal/repository"

func GetUserEmailByUUId(uUId string) string {
	return repository.GetUserEmailByUUId(uUId)
}
