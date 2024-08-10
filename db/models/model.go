package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	DepartmentName string
}

// Employee represents the employees table
type Employee struct {
	gorm.Model
	EmployeeName string
	DepartmentID uint
	Department   Department `gorm:"foreignKey:DepartmentID"` // Foreign key association
}
