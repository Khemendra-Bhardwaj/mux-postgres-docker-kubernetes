package routes

import (
	"backend/db"
	"backend/db/queries"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func GetEmployees(writer http.ResponseWriter, reader *http.Request) {
	var dbconn *sql.DB
	dbconn, err := sql.Open("postgres", db.ConnStr)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dbconn.Close()

	rows, err := dbconn.Query(queries.GetEmployees)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(rows)
}

func CreateEmployee(writer http.ResponseWriter, reader *http.Request) {

}

func GetDepartments(writer http.ResponseWriter, reader *http.Request) {
	dbconn, err := sql.Open("postgres", db.ConnStr)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dbconn.Close()

	rows, err := dbconn.Query(queries.GetDepartments)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var departments []map[string]interface{}
	for rows.Next() {
		var departmentID int
		var departmentName string
		if err := rows.Scan(&departmentID, &departmentName); err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
			return
		}
		departments = append(departments, map[string]interface{}{
			"department_id":   departmentID,
			"department_name": departmentName,
		})
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(departments); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func CreateDepartment(writer http.ResponseWriter, reader *http.Request) {

}
