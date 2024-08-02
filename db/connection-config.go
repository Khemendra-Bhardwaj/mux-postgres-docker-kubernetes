package db

import (
	"database/sql"
	"log"
)

var dbConn *sql.DB

func SetupConnectionDB() (*sql.DB, error) {
	var err error
	if dbConn == nil {
		dbConn, err = sql.Open("postgres", ConnStr)
		if err != nil {
			log.Printf("Error opening database connection: %v", err)
			return nil, err
		}

		if err = dbConn.Ping(); err != nil {
			log.Printf("Error pinging database: %v", err)
			return nil, err
		}
	}
	return dbConn, nil
}
