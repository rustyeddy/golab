package jen

import (
	"os"
	"testing"
)

var fileTests = []struct {
	path   string
	exists bool
}{
	{path: ".", exists: true},
	{path: "/etc/passwd", exists: true},
	{path: "/pmt/tmp/xjfjs", exists: false},
	{path: "/tmp", exists: true},
	{path: "/", exists: true},
	{path: "", exists: false},
}

// TestGetFileStat tests file stat
func TestGetFileStat(t *testing.T) {
	if ex, err := DoesFileExist("/etc/passwd"); err != nil || ex == false {
		t.Error("failed to identify passwd file")
	}

	for _, tst := range fileTests {
		if ex, err := DoesFileExist(tst.path); err != nil && !os.IsNotExist(err) {
			t.Errorf("failed determine if file exists %s - %v", tst.path, err)
		} else {
			if ex != tst.exists {
				t.Errorf("expected (%t) - got (%t) ", ex, tst.exists)
			}
		}
	}
}

// TestIsDir tests IsDirs
func TestIsDir(t *testing.T) {
	ok := IsDir("/tmp")
	if !ok {
		t.Error("expected ok got ! ok")
	}
}

// TestMkdir and rmdir
func TestMkdir(t *testing.T) {
	path := "/tmp/pie/raspberry/whipcream"
	if err := Mkdirs(path); err != nil {
		t.Error("failed to mkdirs")
	}

	// Is Dir again
	if ok := IsDir(path); !ok {
		t.Error("path not a dir")
	}

	// Now remove the dirs from /tmp/copied
	if err := Rmdirs("/tmp/pie"); err != nil {
		t.Error("failed removing dirs")
	}

	// Let us see what we end up with the dirs removed
	if ok := IsDir("/tmp/pie"); ok {
		t.Error("we expected the pie to be gone, but seems it is still there")
	}

}

// TestCopy
func TestCopy(t *testing.T) {
	src := "/etc/group"
	dst := "/tmp/groupXXX"
	if err := Copy(src, dst); err != nil {
		t.Errorf("copy %s to %s failed %v", src, dst, err)
	}
	if ok, err := DoesFileExist(dst); err != nil || !ok {
		t.Errorf("expected no err and ok: got ok: %t err: %v", ok, err)
	}
}

// TestCopyNoSrc
func TestCopyNoSrc(t *testing.T) {
	err := Copy("/nowheres/ville", "/tmp/too/tulips")
	if err == nil {
		t.Error("expected an error got none")
	}
}

// TestCantCopy
func TestCantCopy(t *testing.T) {
	err := Copy("/dev", "/tmp/toe")
	if err == nil {
		t.Error("expected an error not none")
	}
}

// TestIsFileMarkdown
func TestIsFileMarkdown(t *testing.T) {
	md := IsFileMarkdown("markdown.md")
	if md == false {
		t.Errorf("expected markdown true got false")
	}
}
