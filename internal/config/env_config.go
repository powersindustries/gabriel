package config

import (
	"bufio"
	"email_poc/internal/models"
	"log/slog"
	"os"
	"strings"
)

var environmentVariables models.Env

func setEnvVariables(key string, value string) {
	if len(key) == 0 || len(value) == 0 {
		slog.Info("env key or value is empty: ", key, value)
		return
	}

	switch key {
	case "EMAIL_USER":
		environmentVariables.EmailUser = value
	case "EMAIL_PASS":
		environmentVariables.EmailPassword = value
	case "DB_NAME":
		environmentVariables.DbName = value
	case "DB_USER":
		environmentVariables.DbUser = value
	case "DB_PASS":
		environmentVariables.DbPassword = value
	default:
		slog.Info("env key not found: ", "key", key)
		return
	}
}

func GetEnvVariables(key string) string {
	if len(key) == 0 {
		return ""
	}

	switch key {
	case "email_user":
		return environmentVariables.EmailUser
	case "email_password":
		return environmentVariables.EmailPassword
	case "db_name":
		return environmentVariables.DbName
	case "db_user":
		return environmentVariables.DbUser
	case "db_pass":
		return environmentVariables.DbPassword
	default:
		return ""
	}
}

func LoadEnvData() {
	file, err := os.Open(".env")
	if err != nil {
		slog.Error("Failed to load .env file: ", "error", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		keyValues := strings.SplitN(line, "=", 2)
		if len(keyValues) != 2 {
			slog.Error("env key value not formatted correclty. Check line: ", "line", line)
			return
		}

		key := strings.TrimSpace(keyValues[0])
		value := strings.TrimSpace(keyValues[1])

		setEnvVariables(key, value)
	}
}
