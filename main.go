package main

import (
	"email_poc/internal/config"
	"email_poc/internal/models"
	"email_poc/internal/repository"
	"email_poc/internal/services"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Gabriel.")

	config.LoadEnvData()
	config.InitializeDatabase()

	contentRepository := repository.CreateNewContentRepository()
	newsletterRepository := repository.CreateNewNewsletterRepository()
	subscriberRepository := repository.CreateNewSubscriberRepository()

	contentService := services.CreateNewContentService(contentRepository)
	subscriberService := services.CreateNewSubscriberService(subscriberRepository)
	newsletterService := services.CreateNewNewsletterService(newsletterRepository, subscriberService)
	emailSendingService := services.CreateNewEmailSendingService(contentService, newsletterService)
	schedulerService := services.CreateNewSchedulerService(emailSendingService)

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		state := services.Lifecycle()

		switch state {
		case models.Initialing:
			{
				fmt.Println("Gabriel Initializing.")

				// ToDo: Debug - Remove and replace with endpoint.
				schedulerService.AddContentToScheduler("ea36aeeb-f1d4-49cf-9f1d-34bb47d928d7")

				services.SetLifecycle(models.Running)
			}
		case models.Running:
			{
				select {
				case <-ticker.C:
					schedulerService.CycleContentScheduler()
				default:
					// Yield CPU time to avoid busy waiting
					// time.Sleep(100 * time.Millisecond)
				}
			}
		case models.Stopping:
			{
				fmt.Println("Gabriel Stopping.")
				return
			}
		default:
			{
				fmt.Println("Reached unreachable state. Exiting program.")

				services.SetLifecycle(models.Stopping)
			}
		}
	}
}
