package main

import (
	"log"
	"net/http"
)

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Println(err)

	message := "the server encountered a problem and could not parse the request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}
