package store

import (
	"strings"
	"testing"
)

// TestObjectFromPath
func TestObjectFromPath(t *testing.T) {

	tcases := []struct {
		name        string // test name
		path        string // complete path of the object
		contentType string //
	}{
		{"e1", "/srv/storefs/barney/e1.json", "application/json"},
		{"noext", "/tmp/noext", "application/octet-stream"},
		{"nosuchplace", "/nosuchplace", "application/octet-stream"},
	}

	for _, tc := range tcases {
		obj := ObjectFromPath(tc.path)
		if obj == nil {
			t.Error("expected object got nil")
		} else {
			if strings.Compare(obj.name, tc.name) != 0 {
				t.Errorf("obj name expected (%s) got (%s) ", tc.name, obj.Name)
			}
			if strings.Compare(obj.path, tc.path) != 0 {
				t.Errorf("obj path expected (%s) got (%s) ", tc.path, obj.Path)
			}
			if strings.Compare(obj.contType, tc.contentType) != 0 {
				t.Errorf("obj type expected (%s) got (%s) ", tc.contentType, obj.contType)
			}
		}
	}
}
