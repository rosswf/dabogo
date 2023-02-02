package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rosswf/dabogo/internal/assert"
)

func TestLogging(t *testing.T) {
	var loggerSpy bytes.Buffer

	l := log.New(&loggerSpy, "INFO\t", log.Ldate|log.Ltime)
	app := application{
		infoLog: l,
	}

	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/path", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	app.logRequest(next).ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	got := loggerSpy.Bytes()
	assert.Contains(t, string(got), "INFO\t")
	assert.Contains(t, string(got), "/path")

	// Check that the status code is 200.
	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %d, want: %d", res.StatusCode, http.StatusOK)
	}

}
