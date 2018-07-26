package model

import (
	"testing"
)

type KV struct {
	idx string
	val interface{}
}

var (
	defaultVals = Hashmap{}
)

func init() {
	defaultVals["str"] = "this-is-a-string"
	defaultVals["int"] = 88
	defaultVals["nil"] = nil
	defaultVals["bool"] = true
}

func hmapEmpty() Hashmap {
	var h Hashmap = make(Hashmap)
	return h
}

func hmapItems() Hashmap {
	h := hmapEmpty()
	for i, j := range defaultVals {
		h.Set(i, j)
	}
	return h
}

// Test basic hashmap things
func TestEmptyHashmap(t *testing.T) {
	var hmap Hashmap
	if hmap = hmapEmpty(); hmap == nil {
		t.Error("hmap: expected (hmap) got (nil)")
	}

	if len(hmap) != 0 {
		t.Errorf("hmap: expected len (0) got (%d) ", len(hmap))
	}
}

// Test not exists
func TestNotSet(t *testing.T) {
	hmap := hmapEmpty()
	if _, ex := hmap.Fetch("nothing"); ex {
		t.Error("fetch nothing, expeected exists (false) got (true) ")
	}

	if ex := hmap.Exists("nothing"); ex {
		t.Error("exists: nothing expected (false) got (true) ")
	}

	if val := hmap.Get("nothing"); val != nil {
		t.Errorf("val: expected (nil) got (%v) ")
	}
}

// TestSet
func TestSet(t *testing.T) {
	hmap := make(Hashmap)
	for k, v := range defaultVals {
		hmap.Set(k, v)
	}

	if len(hmap) != len(defaultVals) {
		t.Errorf("hmap len: expected (%d) got (%d) ", len(defaultVals), len(hmap))
	}
}

// TestNames
func TestNames(t *testing.T) {
	var (
		allnames []string
		names    []string
	)

	hm := hmapItems()
	if names = hm.Names(); names == nil {
		t.Error("hashmap empty expected 4 names")
	}

	for _, name := range names {
		switch name {
		case "int", "str", "nil", "bool":
			allnames = append(allnames, name)
		default:
			t.Errorf("names expected (name) got (%s) ", name)
		}
	}

	if len(allnames) != len(names) {
		t.Errorf("hashmap expected (%d) names got (%d) ", len(allnames), len(names))
	}
}

// TestNames
func TestValues(t *testing.T) {
	var (
		allvals []interface{}
		vals    []interface{}
	)

	hm := hmapItems()

	if v, e := hm["nil"]; e {
		if v != nil {
			t.Error("expected nil entry to have nil val")
		}
	} else {
		t.Error("entry nil should exist does not")
	}

	vals = hm.Values()
	if vals == nil || len(vals) != len(hm) {
		t.Errorf("expected vals count (%d) got (%d)", len(hm), len(vals))
	}

	for _, val := range vals {
		allvals = append(allvals, val)
	}
	if len(allvals) != len(vals) {
		t.Errorf("expected vals count (%d) got (%d)", len(vals), len(allvals))
	}
}
