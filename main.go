package main

import (
	"email_poc/internal/config"
	"email_poc/internal/models"
	"email_poc/internal/repository"
	"email_poc/internal/services"
	"log/slog"
	"time"
)

func main() {
	slog.Info("Gabriel.")

	config.LoadEnvData()

	sqlDatabase := config.CreateNewSQLDatabase()

	contentRepository := repository.CreateNewContentRepository(sqlDatabase)
	newsletterRepository := repository.CreateNewNewsletterRepository(sqlDatabase)
	subscriberRepository := repository.CreateNewSubscriberRepository(sqlDatabase)

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
				slog.Info("Gabriel Initializing.")

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
				slog.Info("Gabriel Stopping.")
				return
			}
		default:
			{
				slog.Error("Reached unreachable state. Exiting program.")

				services.SetLifecycle(models.Stopping)
			}
		}
	}
}
