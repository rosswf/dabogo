package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rosswf/dabogo/internal/assert"
	"github.com/rosswf/dabogo/internal/models"
)

func TestHomeHandler(t *testing.T) {
	template, err := newTemplate()
	if err != nil {
		t.Fatal(err)
	}

	app := application{
		config: &models.Config{
			Links: []models.Link{
				{
					Name:  "Google",
					Url:   "https://google.co.uk",
					Color: "#ffffff",
				},
				{
					Name: "BBC",
					Url:  "https://bbc.co.uk",
				},
			},
		},
		template: template,
	}

	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	app.home(w, r)

	res := w.Result()
	defer res.Body.Close()

	// Check that the status code is 200.
	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the contents of the response body.
	got := string(body)

	assert.Contains(t, got, "https://google.co.uk")
	assert.Contains(t, got, "https://bbc.co.uk")
	assert.Contains(t, got, "<!DOCTYPE html>")
	assert.Contains(t, got, "#ffffff")
}
