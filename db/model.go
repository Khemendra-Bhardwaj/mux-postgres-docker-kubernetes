package db

type User struct {
	EmployeeName string `json:"employee_name"`
	DepartmentID int    `json:"department_id"`
}

type Department struct {
	DepartmentName string `json:"department_name"`
}
