package models

import "testing"

func TestRequiredTag(t *testing.T) {
	l := Link{Name: "Google"}

	err := l.checkRequired()

	if err == nil {
		t.Errorf("Expected an error when required tag is empty")
	}
}
