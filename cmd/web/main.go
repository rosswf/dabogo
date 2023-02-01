package main

import (
	"net/http"
	"os"

	"github.com/rosswf/dabogo/internal/models"
)

type application struct {
	config *models.Config
}

func main() {
	f, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}

	cfg := models.NewConfig()
	cfg.Load(f)

	app := application{
		config: cfg,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)

	http.ListenAndServe(":8080", mux)
}
