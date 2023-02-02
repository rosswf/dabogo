package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/rosswf/dabogo/internal/models"
	"github.com/rosswf/dabogo/ui"
)

type application struct {
	config   *models.Config
	template *template.Template
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

	app := application{
		config:   cfg,
		template: template,
	}

	fileServer := http.FileServer(http.FS(ui.Files))

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.Handle("/static/", fileServer)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
