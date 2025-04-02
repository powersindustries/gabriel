package repository

import (
	"context"
	"database/sql"
	"email_poc/internal/config"
	"email_poc/internal/models"

	"github.com/lib/pq"
)

type SubscriberRepository interface {
	GetSubscriberEmailByUUId(uUId string) string
}

type subscriberRepository struct {
	sqlDatabase *config.SQLDatabase
}

func CreateNewSubscriberRepository(sqlDatabase *config.SQLDatabase) SubscriberRepository {
	return &subscriberRepository{sqlDatabase: sqlDatabase}
}

func (this *subscriberRepository) GetSubscriberEmailByUUId(uUId string) string {
	sqlQuery := "SELECT uuid, email, newsletter_uuids FROM subscribers WHERE uuid = $1"
	var subscriber models.Subscriber

	database := this.sqlDatabase.GetDatabaseInstance()

	err := database.QueryRowContext(
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
