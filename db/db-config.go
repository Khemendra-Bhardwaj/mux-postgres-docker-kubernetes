package db

import (
	"backend/db/queries"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Dbconn *sql.DB
var ConnStr = "user=postgres password=postgres123 dbname=mydatabase host=postgres sslmode=disable"

func SetupDatabase() {
	Dbconn, err := sql.Open("postgres", ConnStr)

	if err != nil {
		log.Fatalf("Error Connecting to database %v", err.Error())
		return
	}

	if err = Dbconn.Ping(); err != nil {
		log.Fatalf("Error Ping DB: %v", err)
	}

	defer func() {
		err := Dbconn.Close()
		if err != nil {
			log.Fatalf("Error Closing DB %v", err.Error())
			return
		}
	}()

	_, err = Dbconn.Exec(queries.CreateTableDepartment)
	if err != nil {
		log.Fatalf("Error Creating Department Table %v ", err.Error())
		return
	}

	_, err = Dbconn.Exec(queries.CreateTableEmployees)
	if err != nil {
		log.Fatalf("Error Creating Employee Table %v ", err.Error())
		return
	}

	log.Println("SuccessFully created Both Tables")

}
