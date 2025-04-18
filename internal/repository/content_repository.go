package repository

import (
	"context"
	"database/sql"
	"email_poc/internal/config"
	"email_poc/internal/models"
	"errors"
	"log"
	"log/slog"
	"os"
)

type ContentRepository interface {
	GetContentObjectByUUId(contentUUId string) (*models.Content, error)
	GetRawContentByObject(contentObject *models.Content) ([]byte, error)
}

type contentRepository struct {
	sqlDatabase *config.SQLDatabase
}

func CreateNewContentRepository(sqlDatabase *config.SQLDatabase) ContentRepository {
	return &contentRepository{sqlDatabase: sqlDatabase}
}

func (this *contentRepository) GetContentObjectByUUId(contentUUId string) (*models.Content, error) {
	sqlQuery := `
		SELECT uuid, title, release_date, type, newsletter_uuid 
		FROM content 
		WHERE uuid = $1
	`

	database := this.sqlDatabase.GetDatabaseInstance()

	var content models.Content
	err := database.QueryRowContext(context.Background(), sqlQuery, contentUUId).
		Scan(&content.UUId, &content.Title, &content.ReleaseDate, &content.Type, &content.NewsletterUUId)

	if err != nil || err == sql.ErrNoRows {
		slog.Error("Error fetching content:", err)
		return nil, err
	}

	return &content, nil
}

// ToDo: Replace with actual call to s3 bucket. Below is temp/debug logic.
func (this *contentRepository) GetRawContentByObject(contentObject *models.Content) ([]byte, error) {

	path := "internal/s3/" + contentObject.NewsletterUUId + "/" + contentObject.UUId + contentObject.Type
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("failed to load raw content by id")
	}

	return data, nil
}
