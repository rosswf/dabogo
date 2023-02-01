package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/rosswf/dabogo/ui"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFS(ui.Files, "html/*.tmpl")
	if err != nil {
		log.Print(err)
		return
	}

	err = ts.ExecuteTemplate(w, "home", app.config)
	if err != nil {
		log.Println(err)
		return
	}
}
