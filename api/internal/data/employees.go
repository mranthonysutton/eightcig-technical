package data

import (
	"context"
	"database/sql"
	"errors"
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

func (emp EmployeeModel) Insert(employee *Employee) error {
	query := `
	INSERT INTO employees (name, performance)
	VALUES ($1, $2) RETURNING id, name, performance, created_at`

	args := []interface{}{employee.Name, employee.Performance}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	return emp.DB.QueryRowContext(ctx, query, args...).Scan(&employee.ID, &employee.Name, &employee.Performance, &employee.CreatedAt)
}

func (emp EmployeeModel) Get(id int64) (*Employee, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `SELECT id, name, performance, created_at FROM employees WHERE id = $1 LIMIT 1`

	var employee Employee

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := emp.DB.QueryRowContext(ctx, query, id).Scan(&employee.ID, &employee.Name, &employee.Performance, &employee.CreatedAt)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &employee, nil
}

func (emp EmployeeModel) GetAll() ([]*Employee, error) {
	query := `SELECT id, name, performance, created_at FROM employees ORDER BY performance DESC`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := emp.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	employees := []*Employee{}

	for rows.Next() {
		var employee Employee

		err := rows.Scan(&employee.ID, &employee.Name, &employee.Performance, &employee.CreatedAt)

		if err != nil {
			return nil, err
		}

		employees = append(employees, &employee)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}
