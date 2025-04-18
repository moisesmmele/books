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
	config   Config
	infoLog  *log.Logger
	errorLog *log.Logger
	models   data.Models
}

func main() {
	var config Config
	config.port = 8082

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	//example DSN: "host=localhost port=5432 user=USER password=SECRET dbname=DBNAME sslmode=disable timezone=utc connect_timeout=5"

	dsn := os.Getenv("DSN")
	db, err := driver.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	defer db.SQL.Close()

	app := &App{
		config:   config,
		infoLog:  infoLog,
		errorLog: errorLog,
		models:   data.New(db.SQL),
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
