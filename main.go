package main

import (
	"fmt"
	"os"
	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	PostgresDB			string
	PostgresUser		string
	PostgresPassword	string
	PostgresHost		string
	PostgresURL			string
}

var conf Config

func init() {
	godotenv.Load()

	conf.PostgresDB = os.Getenv("POSTGRES_DB")
	conf.PostgresUser = os.Getenv("POSTGRES_USER")
	conf.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	conf.PostgresHost = os.Getenv("POSTGRES_HOST")

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		conf.PostgresUser,
		conf.PostgresPassword,
		conf.PostgresHost,
		conf.PostgresDB,
	)

	conf.PostgresURL = connStr
}

func main() {
	conn, err := sql.Open("postgres", conf.PostgresURL)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
}
