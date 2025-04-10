package config

import (
	"context"
	"database/sql"
	"log"
	"log/slog"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type SQLDatabase struct {
	database *sql.DB
}

func CreateNewSQLDatabase() *SQLDatabase {
	outputSQLDatabase := &SQLDatabase{}

	dsn := "postgres://" + GetEnvVariables("db_user") + ":" + GetEnvVariables("db_pass") + "@localhost:5432/" + GetEnvVariables("db_name")

	var err error

	outputSQLDatabase.database, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	outputSQLDatabase.database.SetMaxOpenConns(10)
	outputSQLDatabase.database.SetMaxIdleConns(5)
	outputSQLDatabase.database.SetConnMaxLifetime(time.Hour)

	if err := outputSQLDatabase.database.PingContext(context.Background()); err != nil {
		log.Fatal("Database ping failed: ", err)
	}

	slog.Info("Database successfully connected.")
	return outputSQLDatabase
}

func (this *SQLDatabase) GetDatabaseInstance() *sql.DB {
	return this.database
}
