package services

import (
	"email_poc/internal/models"
)

// ToDo: Replace functionality with DEBUG flag. Used for only running scheduler 3 times, then stopping application.
var runCount int = 0

var scheduledContent []string

func RunContentScheduler() {
	println("Running content scheduler...")

	scheduledContentLength := len(scheduledContent)
	for x := 0; x < scheduledContentLength; x++ {
		currContentId := scheduledContent[x]
		if ContentAvailableForSending(currContentId) {
			err := SendEmailByContentId(currContentId)
			if err != nil {
				println("Failed to send email: ", err)
				return
			}
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
	scheduledContent = append(scheduledContent, contentId)
}
