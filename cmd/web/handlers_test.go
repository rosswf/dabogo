package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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
					Name: "Google",
					Url:  "https://google.co.uk",
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

	// Check that the response body contains "https://google.co.uk"
	got := string(body)
	want := "https://google.co.uk"

	if !strings.Contains(got, want) {
		t.Errorf("got: %s, want: %s", got, want)
	}

	// Check that response body contains "https://bbc.co.co.uk"
	want = "https://bbc.co.uk"
	if !strings.Contains(got, want) {
		t.Errorf("got: %s, want: %s", got, want)
	}

	// Check that response body contains "<!DOCTYPE html>"
	want = "<!DOCTYPE html>"
	if !strings.Contains(got, want) {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
