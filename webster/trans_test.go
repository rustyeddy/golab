package jen

import (
	"os"
	"testing"
)

var goodpath = "./examples/simple.site"
var badpath = "/some/path/does/not/exist"

// Test creating a new transplantor based on our example
func TestNewTransplanter(t *testing.T) {

	nt, err := NewTransplantor(goodpath)
	if err != nil {
		t.Errorf("failed to create new transplanter %v", err)
	}

	if nt == nil {
		t.Errorf("the transplantor is nil")
	}

	s, d := nt.Roots()
	if s == nil {
		t.Errorf("src root is nil")
	}

	if d != nil {
		t.Error("dst root is nil")
	}
}

func TestNilRoots(t *testing.T) {
	nt := Transplantor{nil, nil, make(map[string]*Tmpl8)}
	if s, d := nt.Roots(); s != nil || d != nil {
		t.Errorf("s expected nil got %v, d expected nil got %v", s, s)
	}

	var dir string
	if dir = nt.Src(); dir != "" {
		t.Errorf("s gave us a directory value (%s) expected (%s)", dir, "")
	}

	if dir = nt.Dst(); dir != "" {
		t.Errorf("s gave us a directory value (%s) expected (%s)", dir, "")
	}
}

func TestBadSrc(t *testing.T) {
	nt, err := NewTransplantor(badpath)
	if err == nil {
		t.Error("expected NIL got Transplantor")
	}
	if nt != nil {
		t.Error("we have a transplantor that we should not have")
	}
}

func TestGoodDst(t *testing.T) {
	var exists bool

	nt, err := NewTransplantor(goodpath)
	if nt == nil || err != nil {
		t.Error("failed to get transplantor")
	}

	// let us test and set a destination
	tmpdir := "/tmp/yfufjjj"
	if exists, err = DoesFileExist(tmpdir); err != nil {
		t.Errorf("Error should be nil from existing file")
	}

	// remove target dir if it exists
	if exists {
		os.RemoveAll(tmpdir)
	}

	// now the old test and set!
	if err = nt.DstCheckAndSet(tmpdir); err != nil {
		t.Errorf("fail - check and set dst %s - %v", tmpdir, err)
	}

	if src := nt.Src(); src == "" {
		t.Errorf("got the source")
	}

	if dst := nt.Dst(); dst == "" {
		t.Errorf("destination did not get set")
	}

	// now lets get the roots, both of which should NOT be ""
	if s, d := nt.Roots(); s == nil || d == nil {
		t.Errorf("a root is nil")
	}

	// Now lets run our destination directory builder
	if err = nt.BuildDstDir(); err != nil {
		t.Errorf("build dst dir %v", err)
	}

}

// TestBadDst is a destination that we really can not get to.
func TestBadDst(t *testing.T) {
	nt, err := NewTransplantor(goodpath)
	if nt == nil && err != nil {
		t.Errorf("failed to get transplantor %v", err)
	}

	// bad destination will be something we don't have access to
	dst := "/foobaz"
	if err = nt.DstCheckAndSet(dst); err == nil {
		t.Errorf("bad set expecting error %v", err)
	}
}

// TestBadDst is a destination that we really can not get to.
func TestMkdirs(t *testing.T) {
	nt, err := NewTransplantor(goodpath)
	if nt == nil && err != nil {
		t.Errorf("failed to get transplantor %v", err)
	}

	// bad destination will be something we don't have access to
	dst := "/foobar"
	if err = nt.DstCheckAndSet(dst); err == nil {
		t.Errorf("bad set expecting error %v", err)
	}
}
