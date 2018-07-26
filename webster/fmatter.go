package jen

import (
	"fmt"
	"strings"

	"github.com/gernest/front"
)

var fmx *front.Matter

// GetFrontMatter extracts front matter from MD files
// if they exist
func GetFrontMatter() *front.Matter {
	if fmx == nil {
		fmx = front.NewMatter()
		fmx.Handle("---", front.YAMLHandler)
	}
	return fmx
}

// ParseFrontMatter and return front matter, string and error
func ParseFrontMatter(txt string) (f map[string]interface{}, b string, err error) {
	fmx := GetFrontMatter()
	if fmx == nil {
		return nil, "", fmt.Errorf("Failed to get front matter harness")
	}

	f, body, err := fmx.Parse(strings.NewReader(txt))
	if err != nil {
		if err != front.ErrUnknownDelim {
			return nil, "", err
		}
		err = nil
		f = nil
		b = txt
	} else {
		b = body
	}
	return f, b, err
}

// DoFrontMatter will create a front handler and try the
// extraction itself
func DoFrontMatter(txt string) (m map[string]interface{}, b string, e error) {
	m, b, e = ParseFrontMatter(txt)
	if e != nil {
		if e != front.ErrUnknownDelim {
			return nil, "", fmt.Errorf("Do Front Matter %v", e)
		}
		e = nil
		m = nil
		b = txt

	}
	return m, b, e
}
