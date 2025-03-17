package config

import (
	"bufio"
	"email_poc/internal/models"
	"errors"
	"fmt"
	"os"
	"strings"
)

var environmentVariables models.Env

func setEnvVariables(key string, value string) {
	if len(key) == 0 || len(value) == 0 {
		fmt.Println("env key or value is empty: ", key, value)
		return
	}

	switch key {
	case "EMAIL_USER":
		environmentVariables.EmailUser = value
	case "EMAIL_PASS":
		environmentVariables.EmailPassword = value
	default:
		fmt.Println("env key not found: ", key)
		return
	}
}

func GetEnvVariables(key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("key is empty")
	}

	switch key {
	case "user":
		return environmentVariables.EmailUser, nil
	case "password":
		return environmentVariables.EmailPassword, nil
	default:
		return "", errors.New("key does not exist in env")
	}
}

func LoadEnvData() {
	rootPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get the root path: ", err)
		return
	}

	file, err := os.Open(rootPath + "/.env")
	if err != nil {
		fmt.Println("Failed to load .env file: ", err)
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
			fmt.Println("env key value not formatted correclty. Check line: ", line)
			return
		}

		key := strings.TrimSpace(keyValues[0])
		value := strings.TrimSpace(keyValues[1])

		setEnvVariables(key, value)
	}
}
