package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func Connect(dbUrl string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dbUrl)
	if err != nil {
		log.Printf("%v", err)

		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)

	return db, nil
}
