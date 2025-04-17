package main

import (
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
}

func main() {
	var config Config
	config.port = 8082

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &App{config, infoLog, errorLog}

	err := app.serve()
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
