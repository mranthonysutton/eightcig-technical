package data

import (
	"database/sql"
	"time"
)

type Employee struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	Name        string    `json:"name"`
	Performance int64     `json:"performance"`
	Date        time.Time `json:"date"`
}

type EmployeeModel struct {
	DB *sql.DB
}
