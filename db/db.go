package db

import (
	"fmt"
	"os"

	migrator "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattes/migrate/source/file"
)

var (
	DB *sqlx.DB
)

func NewDBConnection(connStr string) error {
	conn, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = conn.Ping()
	if err != nil {
		return err
	}

	DB = conn
	return migrateUp(DB)
}

func migrateUp(db *sqlx.DB) error {
	db.Driver()
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrator.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", "./db/migrations"),
		os.Getenv("POSTGRES_DB"), driver)

	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrator.ErrNoChange {
		return err
	}
	return nil
}
