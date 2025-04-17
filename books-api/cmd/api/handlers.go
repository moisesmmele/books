package main

import (
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func (app *App) Login(w http.ResponseWriter, r *http.Request) {
	type Credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var credentials Credentials
	var payload jsonResponse

	err := app.readJSON(w, r, &credentials)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "invalid or missing json"
		_ = app.writeJSON(w, http.StatusBadRequest, payload)
	}

	app.infoLog.Println(credentials.Email, credentials.Password)
	payload.Error = false
	payload.Message = "Success"

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
	}
}
