package models

import (
	"fmt"
	"reflect"
)

type Link struct {
	Name string `yaml:"name" dabogo:"required"`
	Url  string `yaml:"url" dabogo:"required"`
}

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
