package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	Client *sqlx.DB
}

func NewDatabase() (*Database, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DB"),
		os.Getenv("SSL_MODE"),
	)

	//CI fails to connect to the database, so we retry a few times
	retry := 0
	var err error
	for retry < 5 {
		db, err := sqlx.Connect("postgres", connectionString)
		if err != nil {
			retry++
			time.Sleep(1 * time.Second)
			continue
		}
		return &Database{Client: db}, nil
	}
	return nil, err
}

func (d *Database) Ping(ctx context.Context) error {
	return d.Client.PingContext(ctx)
}
