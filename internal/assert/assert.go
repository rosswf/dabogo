package assert

import (
	"strings"
	"testing"
)

func Contains(t *testing.T, got, want string) {
	t.Helper()
	if !strings.Contains(got, want) {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
