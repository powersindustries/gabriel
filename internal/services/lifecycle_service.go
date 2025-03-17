package services

import "email_poc/internal/models"

var lifecycle models.Lifecycle = models.Initialing

func Lifecycle() models.Lifecycle {
	return lifecycle
}

func SetLifecycle(inLifecycle models.Lifecycle) {
	lifecycle = inLifecycle
}
