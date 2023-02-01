package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/rosswf/dabogo/internal/models"
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
	cfg.Load(f)

	template, err := newTemplate()
	if err != nil {
		panic(err)
	}

	app := application{
		config:   cfg,
		template: template,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)

	http.ListenAndServe(":8080", mux)
}
