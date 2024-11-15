package main

import (
	"fmt"
	"os"

	"user/db"
	"user/router"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	// database connection
	godotenv.Load()
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("SSLMODE"),
	)

	dbConn, err := db.NewDBConnection(connStr)
	if err != nil {
		log.Error(err)
	}

	r := router.SetupRouter(dbConn)
	// Listen and Server in 0.0.0.0:8081
	r.Run(":8081")
}
