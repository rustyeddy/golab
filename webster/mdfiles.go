package jen

// Transplanter will scan a source directory tree for input files that will be either copied
// to the destination directory tree.  Certain files will also be transformed along the way,
// for example markdown files will be translated into HTML, we may have had to extract some
// *front matter* and/or merge with a *go template* to produce the resulting HTML.
import (
	"fmt"
	"html/template"
	"io/ioutil"

	"bufio"
	"os"

	"github.com/russross/blackfriday"
)

// =============== MdFile Structure and methods =================================

// Mdfile represents files that are of markdown origin
type Mdfile struct {
	*File
	mdbuf  []byte // direct from source, unprocessed, should be raw markdown/fmatter/whatever
	mdbody string // html snippet (after FMatter has been extracted)
	mdhtml []byte // html after merged with template
	html   string // the final html document

	// destination path
	dstpath string

	// final html for final path
	dsthtml string

	// Interface
	meta  map[string]interface{}
	tmpl8 *Tmpl8
}

// NewMdfile file ready for translation
func (f *File) NewMdfile(path string) (m *Mdfile) {
	m = new(Mdfile)
	m.File = f
	return m
}

// ReadFile if we have not already, .rawbuf will be filled if all goes as planned.
func (m *Mdfile) readMdFile() (err error) {

	// just checkin ..
	if len(m.mdbuf) > 0 {
		return nil
	}
	// get the markdown and possible fmatter from file
	m.mdbuf, err = ioutil.ReadFile(m.Path)
	if err != nil {
		rlog.Translator("ERROR: Opening markdown file for translation: %s %v\n", m.Path, err)
		return err
	}
	return nil
}

// RawBuf return if available otherwise read from the file
func (m *Mdfile) rawBuffer() error {
	if m.mdbuf == nil {
		if err := m.readMdFile(); err != nil {
			return fmt.Errorf("problem reading md file %s - %v", m.Path, err)
		}
	}
	return nil
}

// Extract front matter from this mardown files.
func (m *Mdfile) extractFMatter() (err error) {

	// Extracting fmatter may turn up no fmatter
	fm, body, err := ParseFrontMatter(string(m.mdbuf))
	if err != nil {
		return fmt.Errorf("borked %s: error %v", m.Filename, err)
	}

	// set fields from extracting the meta data
	m.mdbody = body
	m.meta = fm
	return nil
}

// Translate markdown to html.  Makesure fmatter has been extracted first..
func (m *Mdfile) md2html() error {

	bfr := BFRenderer()
	if bfr == nil {
		return fmt.Errorf("failed to create our markdown renderer")
	}

	m.mdhtml = blackfriday.Markdown([]byte(m.mdbody), bfr, bfrExtensions())
	if m.mdhtml == nil {
		return fmt.Errorf("failed attempt to render HTML %s", m.Filename)
	}

	//fmt.Printf("mdhtml: %s", m.mdhtml)
	if options.SanitizeHTML {
		// mdf.Html = bluemonday.UGCPolicy().SanitizeBytes(mdf.Html)
	}
	return nil
}

// executeTmpl8
func (m *Mdfile) executeTmpl8() (err error) {
	var t *Tmpl8
	if t = m.tmpl8; t == nil {
		if t = GetTmpl8("default.html"); t == nil {
			return fmt.Errorf("expected default template got nothing %s", m.Filename)
		}
	}

	rlog.Translator("  creating html file %s", m.dstpath)

	f, err := os.OpenFile(m.dstpath, os.O_CREATE|os.O_WRONLY, 0755) // TODO: fix ext to html ...
	if err != nil {
		return fmt.Errorf("\tfailed to open write file %v", err)
	}
	// close the file and flush buffer(?)
	defer f.Close()
	defer f.Sync()

	iow := bufio.NewWriter(f)
	if iow == nil {
		return fmt.Errorf("\tfailed to create a writer %s", m.dstpath)
	}

	meta := PageMeta{
		"TITLE OF PAGE",
		"This is a page",
		"style.css",
		template.HTML(m.mdhtml),
		"Tis the summary for this article that is a delectable read"}

	err = t.execute(iow, &meta)
	if err != nil {
		return fmt.Errorf("\terror executing template %s", t.Name)
	}
	err = iow.Flush()
	if err != nil {
		return fmt.Errorf("\twriter failed to flush, html file may not have been written")
	}
	return err
}

// !!!!!!!!!!!!!!!!!! GLOBAL VARIABLES !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

var bfrRenderer blackfriday.Renderer

// BFRenderer will convert markdown text to html.  It may need to go through
// some filtering and transformations first, like extracting *front matter*
// and merging with templates, or perhaps scss to css, and so forth ...
func BFRenderer() (bfr blackfriday.Renderer) {
	if bfrRenderer == nil {
		bfrRenderer = blackfriday.HtmlRenderer(bfrFlags(), "TODO TITILE", "TODO.css")
	}
	return bfrRenderer
}

func bfrFlags() (opts int) {
	opts = 0
	opts |= blackfriday.HTML_USE_XHTML
	opts |= blackfriday.HTML_USE_SMARTYPANTS
	opts |= blackfriday.HTML_SMARTYPANTS_FRACTIONS
	opts |= blackfriday.HTML_SMARTYPANTS_LATEX_DASHES
	return opts
}

func bfrExtensions() int {
	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS
	return extensions
}

// output the html from markdown: fmatter should already be extracted
func bfrRender(input []byte) (output []byte) {
	bfr := BFRenderer()
	return blackfriday.Markdown(input, bfr, bfrExtensions())
}
