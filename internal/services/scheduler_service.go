package services

import (
	"email_poc/internal/models"
)

type SchedulerService struct {
	emailSendingService *EmailSendingService

	scheduledContentUUIds []string
	// ToDo: Replace functionality with DEBUG flag. Used for only running scheduler 3 times, then stopping application.
	runCount int
}

func CreateNewSchedulerService(emailSendingService *EmailSendingService) *SchedulerService {
	outputObject := &SchedulerService{}

	outputObject.emailSendingService = emailSendingService

	// ToDo: Debug.
	outputObject.runCount = 0

	return outputObject
}

func (this *SchedulerService) CycleContentScheduler() {
	println("Running content scheduler...")

	scheduledContentUUIdsLength := len(this.scheduledContentUUIds)
	for x := 0; x < scheduledContentUUIdsLength; x++ {
		currContentUUId := this.scheduledContentUUIds[x]

		err := this.emailSendingService.SendEmailByContentUUId(currContentUUId)
		if err != nil {
			println("Failed to send email: ", err)
			return
		}
	}

	// ToDo: Replace functionality with DEBUG flag.
	if this.runCount == 2 {
		SetLifecycle(models.Stopping)
		return
	}
	this.runCount++
}

func (this *SchedulerService) AddContentToScheduler(contentId string) {
	this.scheduledContentUUIds = append(this.scheduledContentUUIds, contentId)
}
