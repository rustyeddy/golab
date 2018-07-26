package store

import "strings"

/*
Some Custom Errors defined just for this Store package.  These were
created to hold a little bit more information than the light wieght
error package provided by Go.  See documentation for specific Errors.

At some point it may be wise to combine the Error and logging packages
since they are so heavily connected.

That is, an Error should NEVER be encountered without being
logged. Likewise, almost all logging messages (unless we are
debugging) are generated due to Errors, Warnings or some anamolie that
we almost certainly should be at least notified about.
*/

// StoreError is an Error type specific to this package.  TODO - take
// advantage of logrus fields to log certain fields are are almost
// always going to be interested in dumping.
type StoreError struct {
	name string   // Error Identity, may act as short description.
	msgs []string // optional more messages
	err  error    // passed by caller (optional)
}

var (
	// General Errors that correspond to errors we know from Unix.
	ErrJSONFail  = &StoreError{name: "JSONFailed", msgs: []string{"JSON failed"}}
	ErrNotFound  = &StoreError{name: "NotFound", msgs: []string{"Item not found"}}
	ErrNotExists = &StoreError{name: "NotExists", msgs: []string{"Item does not already exists"}}
	ErrExists    = &StoreError{name: "Exists", msgs: []string{"Item already exists"}}
	ErrWriteFail = &StoreError{name: "WriteFailed", msgs: []string{"Write failed"}}
	ErrReadFail  = &StoreError{name: "ReadFailed", msgs: []string{"Read failed"}}
)

// Error returns the error string, and satisfies the Error() interface.
func (es *StoreError) Error() string {
	// This is were we'd add the fields if we included them?
	return strings.Join(es.msgs, "\n")
}

// Append more messages to the error. This is optional.  The must always
// have some text for err.msg.
func (es *StoreError) Append(msg ...string) *StoreError {
	es.msgs = append(es.msgs, msg...)
	return es
}
