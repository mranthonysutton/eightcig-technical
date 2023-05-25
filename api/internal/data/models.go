package data

import (
	"database/sql"
	"errors"
)

type Models struct {
	Employees EmployeeModel
}

var (
	ErrRecordNotFound = errors.New("record not found")
)

func NewModels(db *sql.DB) Models {
	return Models{
		Employees: EmployeeModel{DB: db},
	}
}
