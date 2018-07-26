package store

import (
	"os"
	"strings"
	"testing"
)

var (
	testpath = "./tmp"
)

func init() {
	reset()
}

// Remove and remake test directory
func reset() {
	err := os.RemoveAll(testpath)
	if err != nil {
		panic("failed to remove test directories")
	}

	// create a non-empty directory so we can test it out.
	notempty := testpath + "/notempty"
	err = os.MkdirAll(notempty, 0755)
	if err != nil {
		panic("failed to make dir ./run " + err.Error())
	}

	// Verify that store does not exist
	if _, err := os.Stat(notempty); os.IsExist(err) {
		panic("failed to remove " + notempty + " " + err.Error())
	}
}

// CompareStores will compare some fields and ignore others,
// or just validate there is an entry, though we don't really
// care what its value is.
func compareStores(t *testing.T, expect, got *Store) bool {
	success := true
	e, g := expect.Name, got.Name
	if strings.Compare(e, g) != 0 {
		t.Error("expect name (%s) got (%s) ", g, e)
		success = false
	}

	e, g = expect.Path, got.Path
	if strings.Compare(g, e) != 0 {
		success = false
		t.Errorf("expect path (%s) got (%s) ", g, e)
	}

	// If one of the indexes is nil, we expect them
	if (expect.index == nil && got.index != nil) ||
		(expect.index != nil && got.index == nil) {
		success = false
		t.Errorf("expected index (%v) got (%v) ", expect.index, got.index)
	}

	if expect.verbose != got.verbose {
		success = false
		t.Errorf("verbose expected (%t) got (%t)", expect.verbose, got.verbose)
	}
	if got.created == "" {
		success = false
		t.Error("Created expected a time got () ")
	}
	return success
}
