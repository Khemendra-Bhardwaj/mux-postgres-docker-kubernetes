package handler

import (
	"backend/db"
	"backend/db/queries"
	"encoding/json"
	"net/http"
)

func GetDepartmentEmployees(writer http.ResponseWriter, reader *http.Request) {
	// Set up the database connection
	dbconn := db.SetupConnectionDB()

	// Start a transaction
	transaction, err := dbconn.Begin()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer transaction.Rollback() // Ensure rollback if anything goes wrong

	// Execute the query
	rows, err := transaction.Query(queries.GetDepartment_EmployeesArr)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close() // Ensure rows are closed when done

	var results []map[string]interface{}

	for rows.Next() {
		var departmentID int
		var departmentName string
		var employees json.RawMessage // Use RawMessage to handle JSON directly

		err := rows.Scan(&departmentID, &departmentName, &employees)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		// Build a map for each row
		result := map[string]interface{}{
			"department_id":   departmentID,
			"department_name": departmentName,
			"employees":       employees, // Already in JSON format
		}

		results = append(results, result)
	}

	if err := transaction.Commit(); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response header and encode the results to JSON
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(results); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
