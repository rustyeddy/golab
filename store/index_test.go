package store

import (
	"testing"
)

func TestBasicOperations(t *testing.T) {
	st := UseStore(testdir)
	if st == nil {
		t.Error("failed to get testdir")
	}

	nothere := "nothere"
	if st.Exists(nothere) {
		t.Error("expected (false) got (true) ")
	}

	if !st.NotExists(nothere) {
		t.Error("expected (true) got (false) ")
	}

	pattern := "tests/entries/*.json"
	index := st.indexPaths(pattern)
	if index == nil {
		t.Error("idx expected (idx) got () ")
	}

	if len(index) != 4 {
		t.Error("expected (%d) entries got ()", len(index))
	}

	newidx := "e-11"
	if st.NotExists(newidx) {
		t.Errorf("idx %s expected (false) got (true) ", newidx)
	}

	if !st.Exists(newidx) {
		t.Errorf("idx %s expected (true) got (false) ", newidx)
	}

	obj := st.get(newidx)
	if obj == nil {
		t.Error("idx %s expected index got () ", newidx)
	}

	// Copy object and give it a new name
	obj2 := *obj
	newname := "e-22"
	obj2.name = newname

	// Add the object to store with a new name
	st.set(newname, &obj2)

	// now make sure that worked
	obj3 := st.get(newname)
	if obj3 == nil {
		t.Errorf(" failed to get new object just stored ")
	}

	// Are they really the save
	if !obj2.Compare(obj3) {
		t.Error("objects not equal obj2 and obj3 ")
	}

}
