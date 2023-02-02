package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/rosswf/dabogo/internal/models"
	"github.com/rosswf/dabogo/ui"
)

type application struct {
	config   *models.Config
	template *template.Template
	infoLog  *log.Logger
}

func main() {
	f, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}

	cfg := models.NewConfig()
	err = cfg.Load(f)
	if err != nil {
		panic(err)
	}

	template, err := newTemplate()
	if err != nil {
		panic(err)
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := application{
		config:   cfg,
		template: template,
		infoLog:  infoLog,
	}

	fileServer := http.FileServer(http.FS(ui.Files))

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.Handle("/static/", fileServer)

	loggedMux := app.logRequest(mux)

	err = http.ListenAndServe(":8080", loggedMux)
	if err != nil {
		panic(err)
	}
}
