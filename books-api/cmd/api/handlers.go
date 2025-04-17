package main

import (
	"encoding/json"
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

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		//sends error message
		app.errorLog.Println("Invalid JSON")
		payload.Error = true
		payload.Message = "Invalid JSON"

		out, err := json.MarshalIndent(payload, "", "\t")
		if err != nil {
			app.errorLog.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(out)
		return
	}

	//authenticate
	app.infoLog.Println(credentials.Email, credentials.Password)

	//sends response
	payload.Error = false
	payload.Message = "Success"

	out, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		app.errorLog.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
