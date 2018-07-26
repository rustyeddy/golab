package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	l "github.com/rustyeddy/logrus"

	"github.com/rustyeddy/store"
)

var (
	path string
	st   *store.Store
)

func init() {
	path = "./tmp/"
	os.RemoveAll(path)
}

type KV struct {
	K   string
	V   string
	err error
}

func (kv *KV) Error() string {
	if kv.err != nil {
		return kv.err.Error()
	}
	return ""
}

// TestSystem will walk through a typical scenario of creating
// a Store, saving files to it, listing contents of the files
// adding, deleting and moving objects, etc.
func TestInitial(t *testing.T) {

	// Create a new store, the directory will need to be created
	st := store.UseStore(path)
	if st == nil {
		t.Errorf("UseStore expected (st) got () ->  ")
	}

	// We expect the store to be empty, let us grab an index
	c := st.Count()
	if c != 0 {
		t.Errorf("expected count (0) got (%d) ", c)
		t.FailNow()
	}

	// Let's create an object that we will store
	kv := KV{K: "rusty", V: "ranger"}
	_, err := st.StoreObject("rusty", &kv)
	if err != nil {
		t.Errorf("failed to get store object %v ", err)
		t.FailNow()
	}

	// Lets try retrieving the same object, and compare.
	// TODO: use reflection to determine the recieving object type
	var kv2 KV
	err = st.FetchObject("rusty", &kv2)
	if err != nil {
		t.Error("expected object (rusty) got () ", err)
		t.FailNow()
	}
	if strings.Compare(kv2.K, "rusty") != 0 {
		t.Errorf("expected (rusty) got (%s)", kv2.K)
	}

	if strings.Compare(kv2.V, "ranger") != 0 {
		t.Errorf("expected (ranger) got (%s) ", kv2.V)
	}
}

// TestExisting store
func TestExisting(t *testing.T) {

	// Create a new store, the directory will need to be created
	st := store.UseStore(path)
	if st == nil {
		t.Error("failed to retrieve store ", path)
	}

	// We expect the store to be empty, let us grab an index
	c := st.Count()
	if c != 1 {
		t.Errorf("  expected count (1) got (%d) ", c)
	}

	// add a few more items, this should create a total of 4
	// items in the store.
	objs := []KV{{"a", "Apple", nil}, {"z", "Zoo", nil}, {"b", "Bird ", nil}}
	for _, o := range objs {
		st.StoreObject(o.K, o)
	}

	// Rewrite the Z to be a list of the last three items.  This will
	// overwrite the existing content in item z.
	_, err := st.StoreObject("z", objs)
	if err != nil {
		t.Errorf("failed to store object Z -> %v ", err)
	}

	// Now we expect the count to increase to 4 items
	c = st.Count()
	if c != 4 {
		t.Errorf("expected index count (4) got (%d) index %+v ", c, st.Index())
		t.FailNow()
	}
}

// TestReadObjects in the current directory
func TestReadObjs(t *testing.T) {

	// Create a new store, the directory will need to be created
	st := store.UseStore(path)
	if st == nil {
		t.Error("failed to retrieve store ", path)
	}

	names := []string{"a", "b", "rusty"}
	for _, name := range names {
		var kv KV
		err := st.FetchObject(name, &kv)
		if err != nil {
			l.Errorf("failed fetching KV name %s", name)
		}
		if strings.Compare(kv.K, name) != 0 {
			t.Errorf("expected (%s) got (%s) ", name, kv.K)
			continue
		}
	}

	// Read the array and make sure we unraveled it correctly
	var kvlist []KV
	err := st.FetchObject("z", &kvlist)
	if err != nil {
		t.Errorf("expect z-list but got err %v", err)
	}

	if len(kvlist) != 3 {
		t.Errorf("expected kvlist of (3) got (%d)", len(kvlist))
	}
}

// TestDeleteObjs
func TestDeleteObjs(t *testing.T) {
	st := store.UseStore(path)
	if st == nil {
		t.Error("failed to find store for path ", path)
	}
	names := []string{"a", "b", "z"}
	for _, n := range names {
		if err := st.RemoveObject(n); err != nil {
			t.Errorf("Error deleting object %s, %v ", n, err)
		}
	}

	// We should now be left with one item
	if st.Count() != 1 {
		t.Errorf("expected objects (1) got (%d)", st.Count())
	}

	// Verify the files are gone from the directory
	paths, err := filepath.Glob("tmp/*")
	if err != nil {
		t.Error("error filepath Glob")
	}

	if len(paths) != 1 {
		t.Errorf("expected len paths (1) got (%d)", len(paths))
	}
	_, f := filepath.Split(paths[0])
	if strings.Compare(f, "rusty.json") != 0 {
		t.Errorf("expected file (rusty.json) got (%s)", f)
	}
}
