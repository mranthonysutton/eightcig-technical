package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/mranthonysutton/eightcig-technical/api/internal/data"
)

func (app *application) createEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name        string `json:"name"`
		Performance int64  `json:"performance"`
		Date        string `json:"date"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	parsedTime, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	employee := &data.Employee{
		Name:        input.Name,
		Performance: input.Performance,
		Date:        parsedTime,
	}

	err = app.models.Employees.Insert(employee)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/employees/%d", employee.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"employee": employee}, headers)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	employees, err := app.models.Employees.GetAll()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"employees": employees}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	employee, err := app.models.Employees.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
			return
		default:
			app.serverErrorResponse(w, r, err)
		}

		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"employee": employee}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
