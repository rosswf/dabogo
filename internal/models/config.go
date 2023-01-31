package models

import (
	"bytes"
	"errors"
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
	n, err := buf.ReadFrom(r)
	if err != nil {
		return fmt.Errorf("Failed to read yaml: %w", err)
	}
	if n == 0 {
		return errors.New("Empty yaml")
	}

	err = yaml.Unmarshal(buf.Bytes(), &c)
	if err != nil {
		return fmt.Errorf("Failed to parse yaml: %w", err)
	}

	err = c.checkLinks()
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) checkLinks() error {
	for _, link := range c.Links {
		if link.Name == "" || link.Url == "" {
			return fmt.Errorf("Missing link information %v", link)
		}
	}
	return nil
}
