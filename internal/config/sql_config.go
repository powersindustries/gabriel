package config

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type SQLDatabase struct {
	Database *sql.DB
}

func CreateNewSQLDatabase() *SQLDatabase {
	outputSQLDatabase := &SQLDatabase{}

	dsn := "postgres://" + GetEnvVariables("db_user") + ":" + GetEnvVariables("db_pass") + "@localhost:5432/" + GetEnvVariables("db_name")

	var err error

	outputSQLDatabase.Database, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	outputSQLDatabase.Database.SetMaxOpenConns(10)
	outputSQLDatabase.Database.SetMaxIdleConns(5)
	outputSQLDatabase.Database.SetConnMaxLifetime(time.Hour)

	if err := outputSQLDatabase.Database.PingContext(context.Background()); err != nil {
		log.Fatal("Database ping failed: ", err)
	}

	println("Database successfully connected.")
	return outputSQLDatabase
}
