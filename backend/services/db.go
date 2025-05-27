package services

import (
	"database/sql"
	"sync"
)

var (
	dbInstance *sql.DB
	once       sync.Once
)

func InitDB(db *sql.DB) {
	once.Do(func() {
		dbInstance = db
	})
}

func GetDB() *sql.DB {
	return dbInstance
}
