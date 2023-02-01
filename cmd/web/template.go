package main

import (
	"html/template"

	"github.com/rosswf/dabogo/ui"
)

func newTemplate() (*template.Template, error) {
	ts, err := template.ParseFS(ui.Files, "html/*.tmpl")
	if err != nil {
		return nil, err
	}
	return ts, nil
}
