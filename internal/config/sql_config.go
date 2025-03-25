package config

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var Database *sql.DB

func InitializeDatabase() {
	dsn := "postgres://" + GetEnvVariables("db_user") + ":" + GetEnvVariables("db_pass") + "@localhost:5432/" + GetEnvVariables("db_name")

	var err error

	Database, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	Database.SetMaxOpenConns(10)
	Database.SetMaxIdleConns(5)
	Database.SetConnMaxLifetime(time.Hour)

	if err := Database.PingContext(context.Background()); err != nil {
		log.Fatal("Database ping failed: ", err)
	}

	println("Database successfully connected.")
}
