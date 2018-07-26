package store

import (
	"mime"
	"net/http"
	"path/filepath"
	"strings"
)

// Object represents a single object contained by Store.  It has a
// name unique to this store.  The object has a File representing the
// objects location on disk and Content representing the actual data
// that belongs to this object, in the format specified by the content
type Object struct {
	name     string // the name as registered by Store
	path     string // full path in filesystem
	contType string // mime/type

	store    *Store // Pointer back to the store
	buffer   []byte // raw data
	accessed int    // the number of times this object has been accessed
}

// ObjectFromPath will create a new object and prefill
// with path, name and mimetype derived from path extension.
func ObjectFromPath(path string) *Object {
	ext := filepath.Ext(path)
	obj := &Object{
		name:     nameFromPath(path),
		path:     filepath.Clean(path),
		contType: mime.TypeByExtension(ext),
		buffer:   nil,
	}
	if obj.contType == "" {
		obj.contType = "application/octet-stream"
	}
	return obj
}

// ContentFromBytes will create a new Content initialized with appropriate
// data.  We will use http.DetectContentType to determine the type of
// content.
func ObjectFromBytes(buf []byte) (obj *Object) {
	obj = &Object{
		buffer:   buf,
		contType: http.DetectContentType(buf),
	}
	return obj
}

// Name returns the name (the index and filename less extension).
func (o *Object) Name() string {
	return o.name
}

// Path to the actual file
func (o *Object) Path() string {
	return o.path
}

// ContentType is the type of content (json, png, pdf, txt ...).
// Defined by RFCxxxx.
func (o *Object) ContentType() string {
	return o.contType
}

// the name of the object is the filename excluding the extension
// and the path.
func nameFromPath(path string) string {
	_, fname := filepath.Split(path)
	flen := len(fname) - len(filepath.Ext(fname))
	return fname[0:flen] // return less the .ext
}

// Compare two Objects, basically, if they point at the same file the
// will be considered equal regardless of the runtime state (if buffer
// is cached or verbosity)
func (o *Object) Compare(obj *Object) bool {
	if strings.Compare(obj.name, o.name) != 0 {
		return false
	}
	if strings.Compare(obj.path, o.path) != 0 {
		return false
	}
	if strings.Compare(obj.contType, o.contType) != 0 {
		return false
	}
	return true
}

/*
// Decode the object and return an instance of otype
func (obj *Object) Decode(otype interface{}) error {
	if err := json.Unmarshal(obj.buffer, otype); err != nil {
		return ErrJSONFail.Append(err.Error())
	}
	return nil
}
*/
