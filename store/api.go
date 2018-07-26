/*
Store is a simple Object Storage library.  Assign a directory
from your filesystem to a store and start using.

Store uses a string for an Index and a customized object that
represents the value of the data stored on disk.

Store handles two types of data: "managed" and "unmanaged", managed
data is basically JSON, YML, .. we will translate to and from
Go objects.

Unmanaged data will be handled as blobs of bytes.  They will have an
index, time, etc.

*/
package store

// ObjectIndex map[string]*Object defined in index.go

// Storage is the interface that determines what we can use
// for an interface.
type Storage interface {
	// Select a storage to use
	UseStorage(dir string) *Store
	Exists(idx string) bool
	Store(idx string, obj interface{}) *StoreError
	Fetch(idx string, obj interface{}) *StoreError
	Delete(idx string) *StoreError

	// Notify when object has changed
	Notify(pattern string, action func(idx, reason string)) *StoreError
}
