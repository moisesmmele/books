package main

import (
	"books-api/internal/data"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"time"
)

func (app *App) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(app.authTokenMiddleware)
		mux.Get("/users", func(w http.ResponseWriter, r *http.Request) {
			var users data.User
			all, err := users.GetAll()
			if err != nil {
				app.errorLog.Println(err)
				return
			}

			payload := jsonResponse{
				Error:   false,
				Message: "Success",
				Data:    envelope{"users": all},
			}

			err = app.writeJSON(w, http.StatusOK, payload)
			if err != nil {
				app.errorLog.Println(err)
			}
		})
		mux.Post("/foo", func(w http.ResponseWriter, r *http.Request) {
			payload := jsonResponse{
				Error:   false,
				Message: "bar",
			}
			app.writeJSON(w, http.StatusOK, payload)
		})

	})

	mux.Post("/users/login", app.Login)
	mux.Post("/users/logout", app.Logout)

	mux.Get("/test-add-user", func(w http.ResponseWriter, r *http.Request) {
		var user = data.User{
			Email:     "test@email.com",
			FirstName: "Test",
			LastName:  "Test",
		}

		app.infoLog.Println("Adding user", user)
		id, err := app.models.User.Insert(user)
		if err != nil {
			app.errorLog.Println(err)
			app.errorJSON(w, err, http.StatusForbidden)
			return
		}
		app.infoLog.Println("User added", id)
		newUser, _ := app.models.User.GetById(id)
		app.writeJSON(w, http.StatusOK, newUser)
	})
	mux.Get("/test-generate-token", func(w http.ResponseWriter, r *http.Request) {
		token, err := app.models.Token.GenerateToken(2, 60*time.Minute)
		if err != nil {
			app.errorLog.Println(err)
			return
		}
		token.Email = "admin@example.com"
		token.CreatedAt = time.Now()
		token.UpdatedAt = time.Now()

		payload := jsonResponse{
			Error:   false,
			Message: "Token generated",
			Data:    token,
		}
		app.writeJSON(w, http.StatusOK, payload)
	})
	mux.Get("/test-save-token", func(w http.ResponseWriter, r *http.Request) {
		token, err := app.models.Token.GenerateToken(2, 60*time.Minute)
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		user, err := app.models.User.GetById(token.UserId)
		if err != nil {
			app.errorLog.Println(token.UserId, user, err)
			return
		}

		token.UserId = user.Id
		token.CreatedAt = time.Now()
		token.UpdatedAt = time.Now()

		err = token.Insert(*token, *user)
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		payload := jsonResponse{
			Error:   false,
			Message: "Token generated",
			Data:    token,
		}
		app.writeJSON(w, http.StatusOK, payload)
	})
	mux.Get("/test-validate-token", func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		app.infoLog.Println("Validating token", token)
		isValid, err := app.models.Token.IsValidToken(token)
		if err != nil {
			app.errorJSON(w, err, http.StatusForbidden)
			return
		}
		var payload jsonResponse
		payload.Error = false
		payload.Data = isValid
		payload.Message = "Token is valid"

		app.writeJSON(w, http.StatusOK, payload)
	})

	return mux
}
