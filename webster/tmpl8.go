package jen

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
)

// ================= Page Templates ============================

// PageMeta is standard page meta data that we can use
type PageMeta struct {
	Title       string
	Description template.HTML
	Cssfile     string
	Content     template.HTML
	Summary     template.HTML
}

// ============== Tmpl8 Template Wrapper =======================

var tmpl8 = make(map[string]*Tmpl8, 100)

// GetTmpl8 from our map
func GetTmpl8(name string) *Tmpl8 {
	v, o := tmpl8[name]
	if o {
		return v
	}
	return nil
}

// SetTmpl8 in our template map
func SetTmpl8(t *Tmpl8) {
	tmpl8[t.Name] = t
}

// Tmpl8 (Temple 8) are wrappers around go templates
type Tmpl8 struct {
	*File
	Name     string
	template *template.Template
}

// NewTmpl8 will create a new template
func (f *File) NewTmpl8(path string, name string) (t *Tmpl8, err error) {
	t = new(Tmpl8)
	t.Name = name
	t.File = f

	// read the template from the file
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("expected to read file %s but failed %v", path, err)
	}

	// parse and compile the template, if no error template will be ready for input
	tn := template.New(name)
	if tn == nil {
		return nil, fmt.Errorf("expected a new template %s but got none %v", path, err)
	}
	t2, err := tn.Parse(string(content))
	if err != nil {
		return nil, fmt.Errorf("expected a parsed %s file %v", path, err)
	}
	t.template = t2

	// template is ready to go
	return t, nil
}

func (t *Tmpl8) execute(wr io.Writer, meta *PageMeta) (err error) {

	rlog.Translator("  executing template ")

	// TODO: get an io.Writer from the output file directory
	err = t.template.Execute(wr, meta)
	if err != nil {
		return fmt.Errorf("\ttemplate failed to execute %v", err)
	}
	return nil
}
