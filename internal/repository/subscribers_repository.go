package repository

import (
	"context"
	"database/sql"
	"email_poc/internal/config"
	"email_poc/internal/models"

	"github.com/lib/pq"
)

func GetSubscriberEmailByUUId(uUId string) string {
	sqlQuery := "SELECT uuid, email, newsletter_uuids FROM subscribers WHERE uuid = $1"
	var subscriber models.Subscriber

	err := config.Database.QueryRowContext(
		context.Background(), sqlQuery, uUId,
	).Scan(&subscriber.UUId, &subscriber.Email, pq.Array(&subscriber.NewsletterUUIds))

	if err != nil || err == sql.ErrNoRows {
		return ""
	}

	if len(subscriber.Email) > 0 {
		return subscriber.Email
	}

	return subscriber.Email
}
