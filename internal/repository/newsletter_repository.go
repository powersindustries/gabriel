package repository

import (
	"context"
	"database/sql"
	"email_poc/internal/config"
	"email_poc/internal/models"
	"log/slog"

	"github.com/lib/pq"
)

type NewsletterRepository interface {
	GetNewsletterByUUId(newsletterUUId string) (*models.Newsletter, error)
}

type newsletterRepository struct {
	sqlDatabase *config.SQLDatabase
}

func CreateNewNewsletterRepository(sqlDatabase *config.SQLDatabase) NewsletterRepository {
	return &newsletterRepository{sqlDatabase: sqlDatabase}
}

func (this *newsletterRepository) GetNewsletterByUUId(newsletterUUId string) (*models.Newsletter, error) {
	sqlQuery := `
		SELECT uuid, name, description, contact_email, subscriber_list 
		FROM newsletters 
		WHERE uuid = $1
	`

	database := this.sqlDatabase.GetDatabaseInstance()

	var newsletter models.Newsletter
	err := database.QueryRowContext(context.Background(), sqlQuery, newsletterUUId).
		Scan(&newsletter.UUId, &newsletter.Name, &newsletter.Description, &newsletter.ContactEmail, pq.Array(&newsletter.SubscriberList))

	if err != nil || err == sql.ErrNoRows {
		slog.Error("Error getting newsletter: ", err)
		return nil, err
	}

	return &newsletter, nil
}
