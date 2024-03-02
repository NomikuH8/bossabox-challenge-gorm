package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host = "DB_HOST"
	port = "DB_PORT"
	user = "DB_USER"
	pass = "DB_PASS"
	name = "DB_NAME"
	ssl  = "DB_SSL"
)

var database *gorm.DB

// Connects to database and return the Gorm instance
func GetDatabase() *gorm.DB {
	if database != nil {
		return database
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv(host),
		getEnv(port),
		getEnv(user),
		getEnv(pass),
		getEnv(name),
		getEnv(ssl),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	return database
}

// Gets environment variables with a default value
func getEnv(envVar string) string {
	var variable string

	if variable = os.Getenv(envVar); variable != "" {
		return variable
	}

	switch envVar {
	case host:
		variable = "localhost"
	case port:
		variable = "5432"
	case user:
		variable = "postgres"
	case pass:
		variable = "root"
	case name:
		variable = "bossabox"
	case ssl:
		variable = "disable"
	}

	return variable
}
