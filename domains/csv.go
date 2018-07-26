package domains

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ReadCSVFile given a string representing a full path to a csv file
func ReadCSVFile(fname *string) (csvrecs [][]string, err error) {

	// Open the file and get a io.Writer()
	f, err := os.Open(*fname)
	if err != nil {
		return nil, fmt.Errorf("expected to open %s but failed %v", err)
	}

	// Now get a csv reader and get to readin
	r := csv.NewReader(f)
	if r == nil {
		return nil, fmt.Errorf("failed to create csv.reader for %s", fname)
	}
	r.Comment = '#'
	csvrecs, err = r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to ReadAll() csv data %v", err)
	}
	return csvrecs, err
}

// WriteCSVFile takes the given 2d array of strings and formats the
// data as CSV and writes to fname.
// TODO - Turn this into a coroutine to be called by domains
func WriteCSVFile(fname, string, csvrecs [][]string) (err error) {

	return err
}
