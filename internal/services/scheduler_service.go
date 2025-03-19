package services

import (
	"email_poc/internal/models"
)

// ToDo: Replace functionality with DEBUG flag. Used for only running scheduler 3 times, then stopping application.
var runCount int = 0

// Array of Newsletter Ids.
var scheduledNewsletterIds []string

func CycleContentScheduler() {
	println("Running content scheduler...")

	scheduledNewsletterIdsLength := len(scheduledNewsletterIds)
	for x := 0; x < scheduledNewsletterIdsLength; x++ {
		currNewsletterId := scheduledNewsletterIds[x]

		err := SendNewsletterEmail(currNewsletterId)
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
	scheduledNewsletterIds = append(scheduledNewsletterIds, contentId)
}
