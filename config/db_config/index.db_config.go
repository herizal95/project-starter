package db_config

import "os"

var DB_DRIVER = "psql"
var DB_HOST = "127.0.0.1"
var DB_PORT = "5432"
var DB_NAME = "dbname"
var DB_USER = "username"
var DB_PASSWORD = "!required"

func InitDatabaseConfig() {

	hostEnv := os.Getenv("DB_HOST")
	if hostEnv != "" {
		DB_HOST = hostEnv
	}

	portEnv := os.Getenv("DB_PORT")
	if portEnv != "" {
		DB_PORT = portEnv
	}

	nameEnv := os.Getenv("DB_NAME")
	if nameEnv != "" {
		DB_NAME = nameEnv
	}

	userEnv := os.Getenv("DB_USER")
	if userEnv != "" {
		DB_USER = userEnv
	}

	paswordEnv := os.Getenv("DB_PASSWORD")
	if paswordEnv != "" {
		DB_PASSWORD = paswordEnv
	}
}
