package main

import (
	"fmt"
	"os"

	"barterbooks/db"

	"github.com/joho/godotenv"
)

/*
func init() {
	godotenv.Load()


	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		conf.PostgresUser,
		conf.PostgresPassword,
		conf.PostgresHost,
		conf.PostgresDB,
	)

	conf.PostgresURL = connStr
}
*/

func main() {
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
	fmt.Println(connStr)
	err := db.NewDBConnection(connStr)
	if err != nil {
		fmt.Println(err)
	}
}
