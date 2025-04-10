package services

import (
	"email_poc/internal/models"
	"log/slog"
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

func (schedulerServiceSelf *SchedulerService) CycleContentScheduler() {
	slog.Info("Running content scheduler...")

	scheduledContentUUIdsLength := len(schedulerServiceSelf.scheduledContentUUIds)
	for x := 0; x < scheduledContentUUIdsLength; x++ {
		currContentUUId := schedulerServiceSelf.scheduledContentUUIds[x]

		err := schedulerServiceSelf.emailSendingService.SendEmailByContentUUId(currContentUUId)
		if err != nil {
			slog.Error("Failed to send email: ", err)
			return
		}
	}

	// ToDo: Replace functionality with DEBUG flag.
	if schedulerServiceSelf.runCount == 2 {
		SetLifecycle(models.Stopping)
		return
	}
	schedulerServiceSelf.runCount++
}

func (schedulerServiceSelf *SchedulerService) AddContentToScheduler(contentId string) {
	schedulerServiceSelf.scheduledContentUUIds = append(schedulerServiceSelf.scheduledContentUUIds, contentId)
}
