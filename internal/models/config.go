package models

import (
	"bytes"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Links []Link
}

func New() *Config {
	return &Config{Links: []Link{}}
}

func (c *Config) Load(r io.Reader) error {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		return fmt.Errorf("Failed to read yaml: %w", err)
	}

	yaml.Unmarshal(buf.Bytes(), &c)

	return nil
}
