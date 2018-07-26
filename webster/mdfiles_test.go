package jen

import "testing"

func TestMdfile(t *testing.T) {

	// Get file
	fname := "./examples/simple.site/home.md"
	f := NewFile(fname)
	if f == nil {
		t.Errorf("creating file %s", fname)
	}

	m := f.NewMdfile(fname)
	if m == nil {
		t.Errorf("problem creating mdfile %s", fname)
	}

	// This md file should exist
	if m.Exists() == false {
		t.Error("says existing file does not exist")
	}

	if m.Size() <= 0 {
		t.Errorf("file size failed for file %s <= 0", fname)
	}

	if m.rawBuffer() != nil {
		t.Error("failed to get raw buffer from md file")
	}
}

// ===================== Test files ============================

// TestNewFile new file creation
func TestNewFile(t *testing.T) {
	newpath := "/tmp/frontier-sucks"
	nf := NewFile(newpath)
	if nf == nil {
		t.Errorf("problem creating file above")
	}
}
