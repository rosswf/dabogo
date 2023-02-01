package main

import (
	"log"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	err := app.template.ExecuteTemplate(w, "home", app.config)
	if err != nil {
		log.Println(err)
		return
	}
}
