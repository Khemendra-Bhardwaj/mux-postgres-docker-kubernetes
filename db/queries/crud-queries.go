package queries

const CreateEmployee string = `INSERT INTO employees (employee_name, department_id) VALUES ($1, $2) RETURNING employee_id`

const CreateDepartment = `
    INSERT INTO departments (department_name) 
    VALUES ($1) 
    RETURNING department_id;
`

const GetEmployees = `
    SELECT employee_id, employee_name, department_id 
    FROM employees;
`
const GetDepartments = `
    SELECT department_id, department_name 
    FROM departments;
`

const GetEmployeesWithDepartments = `
    SELECT e.employee_id, e.employee_name, d.department_id, d.department_name 
    FROM employees e
    LEFT JOIN departments d ON e.department_id = d.department_id;
`

const GetDepartment_EmployeesArr = `
    SELECT
    d.department_id,
    d.department_name,
    json_agg(
        json_build_object(
            'employee_id', e.employee_id,
            'employee_name', e.employee_name
        )
    ) AS employees
FROM departments d
LEFT JOIN employees e ON d.department_id = e.department_id
GROUP BY d.department_id, d.department_name;
`
