package main

import (
	"books-api/internal/data"
	"books-api/internal/driver"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	port int
}

type App struct {
	config      Config
	infoLog     *log.Logger
	errorLog    *log.Logger
	models      data.Models
	environment string
}

func main() {
	var config Config
	config.port = 8082

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	//example DSN: "host=localhost port=5432 user=USER password=SECRET dbname=DBNAME sslmode=disable timezone=utc connect_timeout=5"

	dsn := os.Getenv("DSN")
	env := os.Getenv("ENV")

	db, err := driver.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	defer db.SQL.Close()

	app := &App{
		config:      config,
		infoLog:     infoLog,
		errorLog:    errorLog,
		models:      data.New(db.SQL),
		environment: env,
	}

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) serve() error {
	app.infoLog.Printf("Listening on port %d", app.config.port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}
	return srv.ListenAndServe()
}

func (app *App) EditUser(w http.ResponseWriter, r *http.Request) {
	var user data.User
	err := app.readJSON(w, r, &user)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	if user.Id == 0 {
		if _, err := app.models.User.Insert(user); err != nil {
			app.errorJSON(w, err)
			return
		}
	} else {
		if u, err := app.models.User.GetById(user.Id); err != nil {
			app.errorJSON(w, err)
			return
		} else {
			u.Email = user.Email
			u.FirstName = user.FirstName
			u.LastName = user.LastName
			if err := u.Update(); err != nil {
				app.errorJSON(w, err)
				return
			}
			if user.Password != "" {
				err := u.ResetPassword(user.Password)
				if err != nil {
					app.errorJSON(w, err)
					return
				}
			}
		}
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Changes saved",
	}
	_ = app.writeJSON(w, http.StatusAccepted, payload)
}
