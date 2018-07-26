package jen

// package gen is a static website generator.  It will search a source directory
// for markdown files that will be converted into html then optionally run through
// one or more go templates, if they exist.  Markdown may also contain "frontmatter"
// consisting of a block of YAML formatted text block delimitted by '---'.
//
// Front matter extraction done by ... front.Matter
// Markdown to HTML conversion by Black Friday: http://github.com/russross/blackfriday
// Templates and the rest thanks to the Go authors and community!

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

/* ********************************** Options *************************************** */

// Options that are used in webster cmdline.
var options = struct {
	SkipHiddenFiles bool
	ClearDst        bool
	SanitizeHTML    bool
}{
	true,
	false,
	true,
}

// ********************************** Webgen *****************************************

// Transplantor is responsible for scanning a source directory, and determining
type Transplantor struct {
	srctree *Dirtree // our tree of sources
	dsttree *Dirtree // we we will move all of our files.

	tmpl8 map[string]*Tmpl8
}

// NewTransplantor will create a transplantor.  A source path is required.
// The destination path will default to the source path prepended with '_'.
func NewTransplantor(src string) (t *Transplantor, err error) {

	t = new(Transplantor)
	if t.srcCheck(src) {
		t.srctree, err = NewDirtree(src)
		if err != nil {
			return nil, fmt.Errorf("trans failed to create dir: %v", err)
		}
	} else {
		return nil, fmt.Errorf("failed source dir check: %v", err)
	}
	t.dsttree = nil
	if t.dsttree != nil {
		return nil, fmt.Errorf("Hmm. dont know how to handle directory trees")
	}

	t.tmpl8 = make(map[string]*Tmpl8)
	return t, nil
}

// Roots will return the src and dst roots. Nil will be returned otherwise
func (t *Transplantor) Roots() (*Dir, *Dir) {
	var st, dt *Dir

	if t.srctree == nil || t.srctree.rootdir == nil {
		st = nil
	} else {
		st = t.srctree.rootdir
	}

	if t.dsttree == nil || t.dsttree.rootdir == nil {
		dt = nil
	} else {
		dt = t.dsttree.rootdir
	}
	return st, dt
}

// srcCheck makes sure the source actually exists
func (t *Transplantor) srcCheck(src string) (ex bool) {
	ex, _ = DoesFileExist(src)
	return ex
}

// Src is the root of the source tree in the filesystem
func (t *Transplantor) Src() string {
	if t.srctree != nil && t.srctree.rootdir != nil {
		return t.srctree.rootdir.Path
	}
	return ""
}

// Dst returns the destination source tree in the filesystem
func (t *Transplantor) Dst() string {
	if t.dsttree != nil && t.dsttree.rootdir != nil {
		return t.dsttree.rootdir.Path
	}
	return ""
}

// DstCheckAndSet determines if the destination exists, if so remove it.
// If the destination does exist we need to decide if we want to clear
// it, delet and recreate or fail.
func (t *Transplantor) DstCheckAndSet(dst string) (err error) {
	var td *Dirtree
	if td, err = NewDirtree(dst); err != nil {
		return fmt.Errorf(" new dirtree failed %v", err)
	}
	t.dsttree = td
	return nil
}

// MakeDirectories will create the directory specified by node and all parent directories
func (t *Transplantor) MakeDirectories(node *string) error {
	return os.MkdirAll(*node, 0755)
}

// BuildDstDir builds the destination directory tree
func (t *Transplantor) BuildDstDir() (err error) {

	var dstpath string
	srcroot, dstroot := t.Roots()
	if dstroot == nil || dstroot.Path == "" {
		return fmt.Errorf("invalid destination director %v", err)
	}

	var exists bool
	if exists, err = DoesFileExist(dstroot.Path); err != nil {
		return fmt.Errorf("puked trying to find dst %s - %v ", dstpath, err)
	}

	// create the directory if it does not already exist
	if exists == false {
		// If not create it (and everything before
		err = os.MkdirAll(dstpath, 0755)
		if err != nil {
			return fmt.Errorf("fail %s, %v", dstpath, err)
		}
		// verify we have created this directory
		if exists, err = DoesFileExist(dstpath); err != nil || exists == false {
			return fmt.Errorf("we tried but failed to create %s - %v ", dstpath, err)
		}
	}

	// We have the destination directory, now we are going to walk the subdir structure
	// of the source dirtree and replicate it on our destination.
	for n, d := range srcroot.subdirs {

		// Do not *copy* directories that begin with an underscore: '_', however
		// we will still scan their contents for required source files.  This can
		// be used for storing files like templates (or scss files) used to generate
		// the final sources, but we don't want those files copied over directly.
		if d.Filename[0] == '_' {
			continue // on to next directory
		}

		// log.Printf("  new directory: %s ", n)
		dstpath := filepath.Join(dstroot.Path, n)
		if exists, err = DoesFileExist(dstpath); err != nil {
			return fmt.Errorf("subdirs: %s - %v", dstpath, err)
		}

		// for now we are just going to destroy whatever exists at the destination
		if exists {
			os.RemoveAll(dstpath)
		}

		// It does not exist, we will create it (at least try to
		err = os.Mkdir(dstpath, 0755)
		if err != nil {
			return fmt.Errorf("failed make dir: %s", d.Path)
		}

		// May as well generate mdf files since we are here ...
		// TODO - create a channel to create a queue of directories that are ready to process
		for n, m := range d.mdfiles {
			n = BasenameNoExt(n) + ".html"
			m.dstpath = filepath.Join(dstroot.Path, n)
		}

	}

	// Walk the files that need to complete their transformation with templates
	for n, m := range srcroot.mdfiles {
		n = BasenameNoExt(n) + ".html"
		rlog.FileSystem("  determine destination new mdfile %s", n)
		m.dstpath = filepath.Join(dstroot.Path, n)
	}
	return err
}

// Markdown2HTML - translate our md (with out front matter)
func (t *Transplantor) processMarkdown() (err error) {

	rd := t.srctree.rootdir
	for relpath, mdf := range rd.mdfiles {
		rlog.Translator("  processing markdown file - path: %s relpath %s  ", mdf.Path, relpath)

		// Double check that our markdown file exists
		var exists bool
		if exists, err = DoesFileExist(mdf.Path); err != nil {
			return fmt.Errorf("failed stating mdf path - %s - %v", mdf.Path, err)
		}
		if exists == false {
			return fmt.Errorf("failed mdf path does not exist - %s - %v", mdf.Path, err)
		}

		if err = mdf.rawBuffer(); err != nil {
			return fmt.Errorf("failed reading markdown file %s", mdf.Filename)
		}

		if err = mdf.extractFMatter(); err != nil {
			return fmt.Errorf("failed extracting front matter %s, err %v", mdf.Filename, err)
		}

		if err = mdf.md2html(); err != nil {
			return fmt.Errorf("failed converting markdown to html %s, err %v", mdf.Filename, err)
		}
	}
	return err
}

// Markdown2HTML - translate our md (with out front matter)
func (t *Transplantor) executeTmpl8s() (err error) {

	rootdir := t.srctree.rootdir
	for relpath, mdf := range rootdir.mdfiles {
		rlog.Translator("  processing markdown file - path: %s relpath %s  ", mdf.Filename, relpath)

		if mdf.dstpath == "" {
			rpne := BasenameNoExt(relpath) + ".html"
			mdf.dstpath = filepath.Join(t.dsttree.rootdir.Path, rpne)
			rlog.Translator("    dest path: %s", mdf.dstpath)
		}

		if err = mdf.executeTmpl8(); err != nil {
			return fmt.Errorf("    errored executing template %s", mdf.Path)
		}
	}
	return err
}

// CopySources from d1 to d2
func (t *Transplantor) CopySources() (err error) {
	for _, n := range t.srctree.rootdir.files {
		var rel string
		rel, err = filepath.Rel(t.srctree.path, n.Path)
		if err != nil {
			return fmt.Errorf("copy rel paths %v", err)
		}
		spath := filepath.Join(t.srctree.path, rel)
		dpath := filepath.Join(t.dsttree.path, rel)

		rlog.FileSystem("Copying %s to %s", spath, dpath)
		err = Copy(spath, dpath)
		if err != nil {
			return fmt.Errorf("copy file to %s - %v", dpath, err)
		}
	}
	return err
}

// VerifySite will verify a given site according to our site.map
func (t *Transplantor) VerifySite() (err error) {
	return err
}

// GenerateSite walks through the steps of generating a static site from sources
func (t *Transplantor) GenerateSite(dst string) (err error) {
	// src := "../examples/simple.site"
	srcroot, _ := t.Roots()

	// Scan the source tree
	if err = srcroot.ScanSubdirs(); err != nil {
		log.Fatal("rats scanning the sources failed")
	}

	// Check our destination make sure it can be made and written to
	if err = t.DstCheckAndSet(dst); err != nil {
		log.Fatalf("failed to check and set the dst: %s err %v", dst, err)
	}

	// Build destination directory structure
	if err = t.BuildDstDir(); err != nil {
		log.Fatalf("oh man, failed to build the directory structure... %s", err)
	}

	// Translate markdown files to html snippets
	if err = t.processMarkdown(); err != nil {
		log.Fatalf("failed to translate some md files %s", err)
	}

	// Produce the final HTML and write to destination
	if err = t.executeTmpl8s(); err != nil {
		log.Fatalf("failed to execute tempaltes %s", err)
	}

	// Now move remaining assests (files: .css, .jpg, .pdf, etc.)
	if err = t.CopySources(); err != nil {
		log.Fatal("failed to copy sources")
	}

	if err = t.VerifySite(); err != nil {
		log.Fatal("bummer the destinaiton is broken? ")
	}

	log.Println("We are done!  All is well methinks... ")
	return err
}
