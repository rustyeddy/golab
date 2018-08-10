package jen

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// ============ File Implementation ==========================

// File represents all regular files
type File struct {
	Path     string // Full or relative path of the file
	Filename string // Filename including extension
	Fileext  string // The file type "md", "html" or "_tmpl"
	Noext    string // Filename with no extension

	fileinfo *os.FileInfo
}

// NewFileInfo create from a path and file info - from walk
func NewFileInfo(path string, fi os.FileInfo) (f *File) {
	f = new(File)
	f.Init(path, fi)
	return f
}

// NewFile create a struct possibly for non-existant file
func NewFile(path string) (f *File) {
	f = NewFileInfo(path, nil)
	return f
}

// Init - ialize the File structure
func (f *File) Init(path string, fi os.FileInfo) *File {
	f.Path = path
	f.Filename = filepath.Base(path)
	f.Fileext = filepath.Ext(path)
	f.Noext = BasenameNoExt(f.Filename)
	if f.fileinfo == nil {
		fi, err := os.Stat(path)
		if err != nil {
			return nil
		}
		f.fileinfo = &fi
	}
	return f
}

// Size of the file in bytes
func (f *File) Size() int64 {
	if fi := *f.fileinfo; fi != nil {
		return fi.Size()
	}
	return int64(-1)
}

// Exists returns true if file or directory exists, false if not
func (f *File) Exists() bool {
	ex, _ := DoesFileExist(f.Path)
	return ex
}

// ################## DirTree #####################

// Dirtree has a path and a pointer to a Dir
type Dirtree struct {
	path    string
	rootdir *Dir
}

// NewDirtree will create a new directory tree represented by path
func NewDirtree(path string) (dtree *Dirtree, err error) {
	dtree = new(Dirtree)
	dtree.path = path
	dtree.rootdir, err = NewDir(path, nil)
	if err != nil {
		return nil, fmt.Errorf("NewDirtree failed %v", err)
	}
	return dtree, err
}

// Root returns the roots of the source and destination trees
func (dt *Dirtree) Root() (d *Dir) {
	return dt.rootdir
}

// **************** Directories *******************

// Dir structure represents filesystem directories
type Dir struct {
	*File
	path    string
	parent  *Dir
	subdirs map[string]*Dir
	files   map[string]*File

	// Cheat and add specific stuff for us
	mdfiles map[string]*Mdfile
	tmpl8s  map[string]*Tmpl8
}

// NewDir is used to create a new directory (in memory)
// may or may not exist on disk.  May also be used to create
// a directory that does not yet exist
func NewDir(path string, parent *Dir) (d *Dir, err error) {
	d = new(Dir)
	d.path = path
	d.parent = parent

	d.subdirs = make(map[string]*Dir, 10)
	d.files = make(map[string]*File, 100)
	d.mdfiles = make(map[string]*Mdfile, 100)
	d.tmpl8s = make(map[string]*Tmpl8, 100)

	var ex bool
	if ex, err = DoesFileExist(path); err != nil {
		return nil, fmt.Errorf("new dir %s - %v", path, err)
	}

	if ex {
		fi, err := GetFileStat(path)
		if err != nil {
			return nil, fmt.Errorf("NewDir GetFileStat fail %v", err)
		}
		d.File = NewFileInfo(path, fi)
	} else {
		err := Mkdir(path, 0755)
		if err != nil {
			return nil, fmt.Errorf("NewDir Mkdir fail %s err %v", path, err)
		}
		fi, err := GetFileStat(path)
		if err != nil {
			return nil, fmt.Errorf("error getting file info %v", err)
		}
		d.File = NewFileInfo(path, fi)
	}
	return d, nil
}

// NewDir creates a new directory structure (memory) representing
// a directory that may or may not exist.
func (d *Dir) NewDir(path string) (dir *Dir, err error) {
	dir, err = NewDir(path, d)
	d.subdirs[dir.Filename] = dir
	return dir, err
}

// Set the file object for this directory
func (d *Dir) setFile(f *File) {
	d.File = f
}

// ScanSubdirs from the source tree ready it for translation
func (d *Dir) ScanSubdirs() (err error) {

	log.FileSystem("scanning subdirs of %s", d.Path)
	rootpath := d.Path

	// Scan the root directory.  Gather directories, markdown and template files.
	err = filepath.Walk(d.Path, func(path string, fi os.FileInfo, err error) (e error) {

		log.FileSystem("path: %s", path)

		if path == rootpath {
			log.FileSystem("\tskipping root directory %s", "")
			return nil
		}

		// Get basename and ignore hidden files
		basename := filepath.Base(path)
		if basename[0] == '.' {
			log.FileSystem("\tignoring hidden dir and files %s", basename)
			if fi.IsDir() {
				// Must return SkipDir to stop descending into the hidden dir
				return filepath.SkipDir
			}
			// If it is a hidden file returni nil to continue walk.
			// Returning an error would stop the walk, which we do not want to do.
			return nil
		}

		// Get the relative path between src root and current path
		relpath, _ := filepath.Rel(rootpath, path)
		if err != nil || relpath == "" {
			return fmt.Errorf("\texpected relpath got nothing: path %s root: %s", path, rootpath)
		}

		// If this node is a directory we'll create the representative struct
		if fi.IsDir() {
			log.FileSystem("\tdir %s - %s", relpath, path)
			dir, err := NewDir(path, nil)
			if err != nil {
				return fmt.Errorf("\tfailed to create dir %s err %v", path, err)
			}
			d.subdirs[relpath] = dir
			return nil
		}

		// Now we *assume* we have a ordinary file ...
		fp := NewFileInfo(path, fi)
		switch fp.Fileext {
		case ".md":
			log.FileSystem("\tmarkdown rel %s path %s", relpath, path)
			md := fp.NewMdfile(path)
			d.mdfiles[relpath] = md

		case ".tmpl":
			log.FileSystem("\ttemplate rel %s path %s", relpath, path)
			tmpl, err := fp.NewTmpl8(path, fp.Noext)
			if err != nil {
				return fmt.Errorf("%v", err)
			}
			SetTmpl8(tmpl)

		default:
			log.FileSystem("\tfile rel %s path %s", relpath, path)
			d.files[relpath] = fp
		}
		return nil
	})

	// All files have been scanned and placed in their appropriate location(s)
	return err
}
