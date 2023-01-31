package models

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

// Config contains all the configuration for the application loaded from yaml.
type Config struct {
	Links []Link
}

// NewConfig returns a ptr to a Config with Links initialised to an empty
// slice of Link
func NewConfig() *Config {
	return &Config{Links: []Link{}}
}

// Load populates the Config fields with data in the required yaml format.
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

// CheckLinks ensures that all the required fields in each Link are present.
func (c *Config) checkLinks() error {
	for _, link := range c.Links {
		return link.checkRequired()
	}
	return nil
}
