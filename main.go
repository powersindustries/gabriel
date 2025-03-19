package main

import (
	"email_poc/internal/config"
	"email_poc/internal/models"
	"email_poc/internal/services"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Gabriel.")

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		state := services.Lifecycle()

		switch state {
		case models.Initialing:
			{
				fmt.Println("Gabriel Initializing.")
				config.LoadEnvData()

				services.InitializeNewsletterService()
				services.InitializeUsers()
				services.InitializeContentService()

				services.InitializeEmailSendingService()

				// ToDo: Remove and replace with endpoint.
				services.AddContentToScheduler("test newsletter")

				services.SetLifecycle(models.Running)
			}
		case models.Running:
			{
				select {
				case <-ticker.C:
					services.CycleContentScheduler()
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
