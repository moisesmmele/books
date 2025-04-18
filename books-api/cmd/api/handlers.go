package main

import (
	"errors"
	"net/http"
	"time"
)

type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type envelope map[string]interface{}

func (app *App) Login(w http.ResponseWriter, r *http.Request) {
	type Credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var credentials Credentials
	var payload jsonResponse

	app.infoLog.Println("reading JSON.")
	err := app.readJSON(w, r, &credentials)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "invalid or missing json"
		_ = app.writeJSON(w, http.StatusBadRequest, payload)
	}
	app.infoLog.Println("JSON decoded. Data:")
	app.infoLog.Println(credentials.Email, credentials.Password)
	user, err := app.models.User.GetByEmail(credentials.Email)
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("invalid username or password"))
		return
	}
	app.infoLog.Println("validating user")
	validatedPass, err := user.VerifyPassword(credentials.Password)
	if err != nil || !validatedPass {
		app.errorLog.Println(err)
		app.errorJSON(w, errors.New("invalid username or password"))
		return
	}
	app.infoLog.Println("user validated. result: ", validatedPass)

	app.infoLog.Println("generating token")
	token, err := app.models.Token.GenerateToken(user.Id, 24*time.Hour)
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, err)
		return
	}
	app.infoLog.Println("generated token. result: ", token)
	app.infoLog.Println("Inserting token into database")
	err = app.models.Token.Insert(*token, *user)
	if err != nil {
		app.errorLog.Println(err)
		app.errorJSON(w, err)
	}
	app.infoLog.Println("writing payload")
	payload = jsonResponse{
		Error:   false,
		Message: "authenticated",
		Data: envelope{
			"token": token,
			"user":  user,
		},
	}
	app.infoLog.Println("responding with payload")
	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
	}
}
func (app *App) Logout(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Token string `json:"token"`
	}
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, errors.New("invalid JSON"))
		return
	}

	err = app.models.Token.DeleteToken(requestPayload.Token)
	if err != nil {
		app.errorJSON(w, errors.New("invalid Token"))
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "deauthenticated",
	}

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
	}
}
