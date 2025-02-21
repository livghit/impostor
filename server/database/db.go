package database

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// create the database connection as a singleton
type Database struct {
	Conn *gorm.DB
}

func New(path string) *Database {
	connection, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return &Database{
		Conn: connection,
	}
}
