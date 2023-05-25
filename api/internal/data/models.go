package data

import "database/sql"

type Models struct {
	Employees EmployeeModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Employees: EmployeeModel{DB: db},
	}
}
