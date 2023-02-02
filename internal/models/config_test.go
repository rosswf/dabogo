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
        color: "#ffffff"
      - name: BBC
        url: https://bbc.co.uk
    `
	d := strings.NewReader(data)

	cfg := NewConfig()
	cfg.Load(d)

	want := &Config{
		Links: []Link{
			{Name: "Google", Url: "https://google.co.uk", Color: "#ffffff"},
			{Name: "BBC", Url: "https://bbc.co.uk"},
		},
	}

	if !reflect.DeepEqual(cfg, want) {
		t.Errorf("got: %v, want: %v", cfg, want)
	}

}

func TestLoadMalformed(t *testing.T) {
	data := `
    links:
      - name: Google 
      - name: BBC 
        url: https://bbc.co.co.uk
    `

	d := strings.NewReader(data)

	cfg := NewConfig()
	err := cfg.Load(d)

	if err == nil {
		t.Errorf("Expected an error when parsing yaml")
	}
}

func TestLoadEmpty(t *testing.T) {
	data := ""

	d := strings.NewReader(data)

	cfg := NewConfig()
	err := cfg.Load(d)

	if err == nil {
		t.Errorf("Expected an error when parsing yaml")
	}
}
