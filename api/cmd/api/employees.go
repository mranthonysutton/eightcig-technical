package main

import (
	"errors"
	"net/http"

	"github.com/mranthonysutton/eightcig-technical/api/internal/data"
)

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
