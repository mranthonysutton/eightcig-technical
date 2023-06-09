package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/employees", app.listEmployeesHandler)
	router.HandlerFunc(http.MethodGet, "/v1/employees/:id", app.showEmployeesHandler)
	router.HandlerFunc(http.MethodPost, "/v1/employees", app.createEmployeeHandler)

	return router
}
