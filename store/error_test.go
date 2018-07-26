package store

import (
	"strings"
	"testing"
)

func TestErrors(t *testing.T) {
	err := ErrNotFound
	err.Append("rusty")

	errstr := err.Error()
	msgs := []string{"Item not found", "rusty"}
	msgstr := strings.Join(msgs, "\n")

	if strings.Compare(errstr, msgstr) != 0 {
		t.Errorf("expected (%s) got (%s) ", errstr, msgstr)
	}
}
