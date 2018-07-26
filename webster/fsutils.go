package jen

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ******************** Some File Utilities ****************************

// GetFileStat does just wat it says.
func GetFileStat(path string) (fi os.FileInfo, err error) {
	return os.Stat(path)
}

// DoesFileExist will return true or false depending on the answer
func DoesFileExist(path string) (ex bool, err error) {

	// ex defaults to false, only set true if all criteria are met
	_, err = GetFileStat(path)
	if err != nil {
		if os.IsNotExist(err) {
			ex = false
			err = nil
		}
	} else {
		ex = true
	}
	return ex, err
}

// IsDir will determine if the path represents a directory
func IsDir(path string) (ok bool) {
	fi, err := GetFileStat(path)
	if err == nil && fi.IsDir() {
		ok = true
	}
	return ok
}

// Mkdir will create a new directory
func Mkdir(path string, perms os.FileMode) (err error) {
	// Create the destination dir and any missing parent directories
	err = os.MkdirAll(path, perms)
	if err != nil {
		rlog.FileSystem("\tfailed to create dir %s because %s\n", path, err)
	}
	return err
}

// Rmdirs recursively removes an entire dir tree
func Rmdirs(path string) error {
	return os.RemoveAll(path)
}

// IsFileMarkdown will determine if the file is a markdown file based on
// it's extension being either .md or .markdown.
//
// TODO: do we want to include .txt files (or make it an option?)
func IsFileMarkdown(path string) (md bool) {
	ext := filepath.Ext(path)
	if ext == ".md" || ext == ".markdown" {
		md = true
	}
	return md
}

// BasenameNoExt returns the filepath.Base() less the extension
func BasenameNoExt(path string) (str string) {
	base := filepath.Base(path)
	ext := filepath.Ext(base)
	str = path[0 : len(path)-len(ext)]
	return str
}

// Copy does just that, copies the file from src to dst
func Copy(src, dst string) (err error) {

	sfi, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sfi.Mode().IsRegular() {
		// Can not copy non-regular files
		return fmt.Errorf("can not copy non-regular file: %s", src)
	}

	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("Failed to open source file %s - %v", src, err)
	}

	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create dst %s - %v", dst, err)
	}

	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}() // parens are necessary - this is an anonymous function
	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	err = out.Sync()
	return err
}

// CountDupLines counts lines from files but only once
func CountDupLines(files []string) (counts map[string]int) {
	counts = make(map[string]int)
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			f.Close()
			continue
		}
		input := bufio.NewScanner(f)
		for input.Scan() {
			counts[input.Text()]++
		}
		f.Close()
	}
	// Ignore potential errors
	return counts
}

// Mkdirs will create all dirs with perms if it can
func Mkdirs(path string) error {
	return os.MkdirAll(path, 0755)
}
