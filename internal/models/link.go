package models

import (
	"fmt"
	"reflect"
)

// A Link contains the information to construct a HTML hyperlink.
type Link struct {
	Name string `yaml:"name" dabogo:"required"`
	Url  string `yaml:"url" dabogo:"required"`
}

// checkRequired returns an error containing the name of the missing field and the Link
// if any of the Link fields with the `dabago:"required"` tag are an empty string.
func (l *Link) checkRequired() error {
	v := reflect.ValueOf(*l)

	fields := reflect.VisibleFields(v.Type())
	for i, f := range fields {
		tag, ok := f.Tag.Lookup("dabogo")
		if !ok {
			continue
		}
		if tag == "required" && v.Field(i).Interface() == "" {
			return fmt.Errorf("Missing %s for link, %v", f.Name, l)
		}
	}

	return nil
}
