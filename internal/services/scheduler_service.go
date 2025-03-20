package services

import (
	"email_poc/internal/models"
)

// ToDo: Replace functionality with DEBUG flag. Used for only running scheduler 3 times, then stopping application.
var runCount int = 0

// Array of content UUIds.
var scheduledContentUUIds []string

func CycleContentScheduler() {
	println("Running content scheduler...")

	scheduledContentUUIdsLength := len(scheduledContentUUIds)
	for x := 0; x < scheduledContentUUIdsLength; x++ {
		currContentUUId := scheduledContentUUIds[x]

		err := SendEmailByContentUUId(currContentUUId)
		if err != nil {
			println("Failed to send email: ", err)
			return
		}
	}

	// ToDo: Replace functionality with DEBUG flag.
	if runCount == 2 {
		SetLifecycle(models.Stopping)
		return
	}
	runCount++
}

func AddContentToScheduler(contentId string) {
	scheduledContentUUIds = append(scheduledContentUUIds, contentId)
}
