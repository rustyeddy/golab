package store

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/rustyeddy/golib"
	log "github.com/rustyeddy/logrus"
)

var (
	testdir string // shared with all _test.go files
)

// Create the testdir
func init() {
	var err error
	if testdir, err = ioutil.TempDir("/tmp", "store"); err != nil {
		log.Fatal(err)
	}
}

// TestCreateDir will verify that we will create the directory
// when it does not already exist.
func TestCreateDir(t *testing.T) {

	// The temp file should not exist
	if !golib.FileExists(testdir) {
		t.Errorf("expected () got (%s) ", testdir)
	}

	// Get a Store then start storing stuff
	var st *Store
	if st = UseStore(testdir); st == nil {
		t.Errorf("expected store (%s) got ()", testdir)
	}

	// make sure the directory actually exists
	if _, err := os.Stat(testdir); os.IsNotExist(err) {
		t.Errorf("expect (%s) got ()", testdir)
	}
}

// TestCleanup test dir
func TestCleanup(t *testing.T) {
	if true {
		os.RemoveAll(testdir)
	}
	if _, err := os.Stat(testdir); os.IsExist(err) {
		t.Errorf("expect (%s) got ()", testdir)
	}
}
