package repository

import (
	"context"
	"database/sql"
	"email_poc/internal/config"
	"email_poc/internal/models"
	"errors"
	"log"
	"os"
)

func GetContentObjectByUUId(contentUUId string) (*models.Content, error) {
	sqlQuery := `
		SELECT uuid, title, release_date, type, newsletter_uuid 
		FROM content 
		WHERE uuid = $1
	`

	var content models.Content
	err := config.Database.QueryRowContext(context.Background(), sqlQuery, contentUUId).
		Scan(&content.UUId, &content.Title, &content.ReleaseDate, &content.Type, &content.NewsletterUUId)

	if err != nil || err == sql.ErrNoRows {
		log.Println("Error fetching content:", err)
		return nil, err
	}

	return &content, nil
}

// ToDo: Replace with actual call to s3 bucket. Below is temp/debug logic.
func GetRawContentByObject(contentObject *models.Content) ([]byte, error) {

	path := "internal/s3/" + contentObject.NewsletterUUId + "/" + contentObject.UUId + contentObject.Type
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("failed to load raw content by id")
	}

	return data, nil
}
