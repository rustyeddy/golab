package main

import "io/ioutil"

// Write JSON string to specified File
func WriteJSON(fname string, data []byte) error {
	if err = ioutil.WriteFile(fname, data, 0644); err != nil {
		return err
	}
	return nil
}
