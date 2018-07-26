package jen

import "testing"

// TestNewDirtree will test the NewDirtree function
func TestNewDirtree(t *testing.T) {
	dpath := "/tmp/fusjf"
	nd, err := NewDirtree(dpath)
	if err != nil {
		t.Errorf("failed to create new dirtree %s", dpath)
	}

	ndroot := nd.Root()
	if ndroot.Path != dpath {
		t.Errorf("expected %s got %v", dpath, nd.Root())
	}
}

// TestNewdir gegging created
func TestNewdir(t *testing.T) {
	nd, err := NewDir("/tmp", nil)
	if err == nil {
		t.Errorf("failed to create new directory")
	}

	if nd.Path != "/tmp" {
		t.Errorf("failed to create the dir")
	}
}
