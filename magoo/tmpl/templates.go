package main

/*
import (
	//"fmt"
	"html/template"
	"log"
	//"net/http"
	"path/filepath"
)

type TmplStorage struct {
	Name string
	*Storage

	compiled *template.Template
}

// NewTmplStorage create a new template storage
func NewTmplStorage(path, name string) *TmplStorage {
	s, err := NewStorage(path+name, name)
	if err != nil {
		log.Fatalf("failed to create template storage, bailing ... %v", err)
	}
	return &TmplStorage{
		Storage: s,
	}
}

// FindTemplates from the given directory
func (ts *TmplStorage) FindTemplates() []string {
	matches, err := filepath.Glob(ts.Path + string(filepath.Separator) + "*.html")
	if err != nil {
		log.Printf("  ## error expecting bad pattern %v", err)
		return nil
	}

	// bail if we don't have anything
	if matches == nil || len(matches) < 1 {
		log.Printf("  ## no templates in %s", ts.Path)
	}
	return matches
}

// getCompiledTemplates from the specified directory
// TODO check the filesystem for valid templates first .?.
func (ts *TmplStorage) getCompiledTmpl() *template.Template {
	var err error

	// Find and compile the templates if this is our first time
	if ts.compiled == nil {
		mstore := TmplStore()
		if tpaths := mstore.FindTemplates(); tpaths != nil {

		// Q: Is it possible to have a len(tpaths) < 1?
		// A: I don't know, but we'll check for it and return
		//    if it does happen..
			if len(tpaths) < 1 {
				log.Printf("  ## template file list is empty, bailing ... ")
				return nil
			}

			// Has to parse files
			ts.compiled, err = template.ParseFiles(tpaths...)
			if err != nil {
				log.Printf("    ## template compilation failed %v ", err)
				return nil
			}
		}
	}
	return ts.compiled
}
*/
