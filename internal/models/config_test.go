package models

import (
	"fmt"
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

func TestLoadMalformed(t *testing.T) {
	data := `
    links:
      - name: Google 
      - name: BBC 
        url: https://bbc.co.co.uk
    `

	d := strings.NewReader(data)

	cfg := New()
	err := cfg.Load(d)

	fmt.Println(cfg)

	if err == nil {
		t.Errorf("Expected an error when parsing yaml")
	}

}

func TestLoadEmpty(t *testing.T) {
	data := ""

	d := strings.NewReader(data)

	cfg := New()
	err := cfg.Load(d)

	if err == nil {
		t.Errorf("Expected an error when parsing yaml")
	}

}
