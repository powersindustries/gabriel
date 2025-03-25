package repository

import (
	"context"
	"database/sql"
	"email_poc/internal/config"
	"email_poc/internal/models"

	"github.com/lib/pq"
)

func GetUserEmailByUUId(uUId string) string {
	sqlQuery := "SELECT uuid, email, newsletter_uuids FROM users WHERE uuid = $1"
	var user models.User

	err := config.Database.QueryRowContext(
		context.Background(), sqlQuery, uUId,
	).Scan(&user.UUId, &user.Email, pq.Array(&user.NewsletterUUIds))

	if err != nil || err == sql.ErrNoRows {
		return ""
	}

	if len(user.Email) > 0 {
		return user.Email
	}

	return user.Email
}
