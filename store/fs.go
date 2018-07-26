package store

import (
	"fmt"
	"os"

	"github.com/rustyeddy/golib"
)

// =======================================================================
//                  Disk / Container Management
// =======================================================================

// CreateDisk will create the directory Store is using for the container.
func CreateDirectory(dir string) error {

	// Create the directory.  Die if directory exits -- should check perms
	// before attempting to write?  or just go for it..?
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("dir %s: %v", dir, err)
	}

	// Verify that file exists, return error otherwise
	if !golib.FileExists(dir) {
		return fmt.Errorf("expected (%s) got ()", dir)
	}
	return nil
}

// MustCreateDisk just makes sure CreateDisk succeeds or this program
// will go belly up.  There is no point continueing if we ca not save
// our data
func MustCreateDir(dir string) {
	golib.DieError(CreateDirectory(dir))
	golib.DieFalse(golib.FileExists(dir), "Disk Create: path EXISTS ", dir)
}

// Clear out an existing store.  Actually, we completely delete the
// directory and everything that is in it, then we create it again.
// Everything in the directories will be lost.
func Clear(path string) (err error) {
	err = os.RemoveAll(path)
	if err == nil {
		err = os.MkdirAll(path, 0755)
	}
	return err
}
