package models

import (
	"reflect"
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	data := `
    links:
      - name: Google
        url: https://google.co.uk
      - name: BBC
        url: https://bbc.co.uk
    `
	d := strings.NewReader(data)

	cfg := New()
	cfg.Load(d)

	want := &Config{
		Links: []Link{
			{Name: "Google", Url: "https://google.co.uk"},
			{Name: "BBC", Url: "https://bbc.co.uk"},
		},
	}

	if !reflect.DeepEqual(cfg, want) {
		t.Errorf("got: %v, want: %v", cfg, want)
	}

}
