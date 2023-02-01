package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(app.config.Links[0].Url))
}
