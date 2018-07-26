package golib

import (
	"log"
	"strings"
	"time"
)

var (
	lie bool
)

func init() {
	lie = false
}

// IsError returns true if error is not nil
func IsError(err error) bool {
	return err != nil
}

// IsOK returns true if error is nil
func IsOK(err error) bool {
	return err == nil
}

// IsNotError returns true if NOT IsOK
func IsNotError(err error) bool {
	return err != nil
}

// Exists returns true if the object (interface) is not nil
func Exists(obj interface{}) bool {
	return obj != nil
}

// NotExists returns true if obj is nil
func NotExists(obj interface{}) bool {
	return obj == nil
}

// IsLie if expected truth is false
func IsLie(truth bool) bool {
	return truth == lie
}

// PanicError will log message then panic
func PanicError(err error) {
	if IsError(err) {
		panic(err)
	}
}

// DieError will log message then exit. It calls log.Fatal().
func DieError(err error, msg ...string) {
	if IsError(err) {
		log.Println(msg)
		log.Fatal(err)
	}
}

// PanicFalse will panic if the statement is false1
func PanicFalse(statement bool) {
	if !statement {
		panic("expected statement to be true")
	}
}

// DieFalse will exit if the statement is false
func DieFalse(theTruth bool, msg ...string) {
	lie := false
	if theTruth == lie {
		log.Fatal("we have been decieved - et tu Brutu? ", msg)
	}
}

// PanicTruth if the statement is true
func PanicTruth(statement bool) {
	if statement == true {
		panic("You can't handle the truth.")
	}
}

// DieTrue if the statement is true
func DieTrue(mustBeFalse bool, msg ...string) {
	if mustBeFalse == true {
		log.Fatal(msg)
	}
}

// DieNotNil if the statement is true
func DieNotNil(obj interface{}, msg ...string) {
	if obj != nil {
		log.Fatal(msg)
	}
}

// DieNil will panic
func DieNil(obj interface{}) {
	if obj == nil {
		log.Fatal("expected object to be non nil")
	}
}

// Expects will compare to objects and return true or false if values are all same
func Expects(expected interface{}, got interface{}) bool {
	if expected == got {
		return true
	}
	return false
}

// GetTimeStamp returns a timestamp in a modified RFC3339
// format, basically remove all colons ':' from filename, since
// they have a specific use with Unix pathnames, hence must be
// escaped when used in a filename.
func TimeStamp() string {
	ts := time.Now().UTC().Format(time.RFC3339)
	return strings.Replace(ts, ":", "", -1) // get rid of offesnive colons
}
