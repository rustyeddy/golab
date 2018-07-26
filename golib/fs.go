package golib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/rustyeddy/logrus"
)

var (
	// Verbose will determine if we print logs
	Verbose bool
)

func init() {
	Verbose = false // Turn verbosity off by default
}

/*
	The FS utilities are all convenience functions that work on path strings.
	Most of these functions simply wrap the underlying go packages calls.
*/

// FileExists makes it simple to get the answer.
func FileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		if Verbose {
			log.Warn("problem accessing file ", path)
		}
		return false
	}
	return true
}

// FileNotExists returns true when the file does not exist
func FileNotExists(path string) bool {
	return !FileExists(path)
}

// Mkdir will create a directory at path, parent directories will be created
// if they do not already exist.  That is, it calls os.MkdirAll(path).
func Mkdir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("error creating directory %s => %s", path, err)
		}
	}
	return nil
}

// ListFileInfo returns a list of *.os.FileInfo
func ListFileInfo(dir string) (files []os.FileInfo) {
	var err error
	if files, err := ioutil.ReadDir(dir); err == nil {
		return files
	}
	log.Errorf("ioutil.ReadDir failed dir %s err {%v}", dir, err)
	return nil
}

// IndexFileInfo returns a map of FileInfo, indexed on the filename (not path)
func IndexFileInfo(dir string) map[string]os.FileInfo {
	flist := ListFileInfo(dir)
	if flist == nil {
		log.Error("No index, no fun, bailing ... ", dir)
	}

	// Loop round the files creating the index
	index := make(map[string]os.FileInfo, len(flist))
	for _, f := range flist {
		index[f.Name()] = f
	}
	if len(index) < 1 {
		return nil
	}
	return index
}

// CWD get the working directory
func CWD() (dir string, err error) {
	return os.Getwd()
}

// ReadJSON from disk, then unmarshal into object.  This will likely
// be wrapped by Storage.JSON() for caching
func ReadJSON(path string, obj interface{}) error {

	// Read whole file []byte buffer
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("ReadJSON expect content from (%s) got error (%v)", path, err)
	}

	// Unravel JSON into a Go thing of some sort
	err = json.Unmarshal(buf, &obj)
	if err != nil {
		return fmt.Errorf("ReadJSON Umarshal path (%s) got error (%v)", path, err)
	}
	return nil
}

// WriteJSON - Turn the object into json then save it to the file ...
func WriteJSON(path string, obj interface{}) error {

	// JSONify
	jbytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("SaveJSON failed (%s) error (%v)", path, err)
	}

	err = ioutil.WriteFile(path, jbytes, 0755)
	if err != nil {
		return fmt.Errorf("SaveJSON write file failed (%s) error (%v)", path, err)
	}

	return nil
}

// ResetDirectory will create a pristine test directory to begin our testing
func ResetDirectory(path string) {
	// If the file does not exist create it and be ready to go.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, 0755); err == nil {
			log.Fatal("failed to create ", path, err)
		}
	}

	// Check again for the basedir, if we don't have it now we never will
	if _, err := os.Stat(path); err != nil {
		log.Fatal("failed to create test directory ", path, err)
	}
}
