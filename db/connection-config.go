package db

import (
	"database/sql"
	"log"
	"sync"
)

var (
	dbConn *sql.DB
	once   sync.Once
	mu     sync.Mutex
)

func SetupConnectionDB() *sql.DB {
	mu.Lock()
	defer mu.Unlock()

	once.Do(func() {
		var err error
		if dbConn == nil {
			dbConn, err = sql.Open("postgres", ConnStr)
			if err != nil {
				log.Fatalf("Error opening database connection: %v", err)
			}

			if err = dbConn.Ping(); err != nil {
				log.Fatalf("Error opening database connection: %v", err)
			}
		}
	})

	return dbConn

}
