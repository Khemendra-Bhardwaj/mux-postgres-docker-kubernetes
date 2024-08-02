package queries

const (
	CreateTableDepartment = `
	CREATE TABLE IF NOT EXISTS departments (
		department_id SERIAL PRIMARY KEY,
		department_name VARCHAR(100) NOT NULL
	);
	`

	CreateTableEmployees = `
	CREATE TABLE IF NOT EXISTS employees (
		employee_id SERIAL PRIMARY KEY,
		employee_name VARCHAR(100) NOT NULL,
		department_id INT,
		FOREIGN KEY (department_id) REFERENCES departments(department_id) ON DELETE SET NULL
	);
	`
)
