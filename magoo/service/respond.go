package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Content can be used to translate data between formats, that
// can then be used to store to disk or send somewhere.
type Content struct {
	Buffer []byte
	JSON   []byte
	HTML   string
	Text   string
}

/* =================== Canned Responses =================== */

// RespondHTML will send a response with text/html as the content type
func RespondHTML(w http.ResponseWriter, text string) {
	o := struct {
		Msg string
	}{
		Msg: "MagOO version 1.2",
	}
	RespondJSON(w, o)
}

// RespondHTMLError will send a response with text/html as the content type
func RespondHTMLError(w http.ResponseWriter, text string) {
	// p := NewPage("response html", []byte(text))
	// RespondTmpl(w, "error.html", p)
	panic("todo HTML Error")
}

// RespondError will send an http error back to the client,
func RespondError(w http.ResponseWriter, status int, errmsg string) {
	http.Error(w, errmsg, http.StatusInternalServerError)
	return
}

// RespondJSON - return a response encoded in JSON
func RespondJSON(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		log.Printf("  ## failed to send json response %v", err)
		return
	}
}

// RespondText - return a response encoded in JSON
func RespondText(w http.ResponseWriter, jstr string) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, jstr)
}

// RespondErrorJSON - return a response encoded in JSON
func RespondErrorJSON(w http.ResponseWriter, status int, message string) {

	// obj just has to be a "JSON-able" go entity
	w.Header().Set("Content-Type", "application/json")
	obj := struct {
		Status  int
		Message string
	}{status, message}
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		log.Printf("  ## failed to send json response %v", err)
		return
	}
}

// RespondTemplate will turn the data provided with the Page
// struct will be placed in the template and given back
// to the caller
/*
func RespondTmpl(w http.ResponseWriter, tmpl string, p *Page) {

		if p == nil {
			RespondError(w, 500, "  ## missing page for rendering")
			return
		}

		ts := TmplStore()
		t := ts.getCompiledTmpl()
		if t == nil {
			RespondError(w, 500, "  ## failed to compile templates")
			return
		}

		if err := t.ExecuteTemplate(w, tmpl, p); err != nil {
			err = fmt.Errorf("  === failed template %s page %s ", tmpl, p.Title)
			RespondError(w, 500, err.Error())
		}
}
*/

// RespondErrorTmpl will create an error page with the following string.
func RespondErrorTmpl(w http.ResponseWriter, errmsg string) {
	/*
		ts := TmplStore()
		t := ts.getCompiledTmpl()
		if t == nil {
			RespondError(w, 500, "  ## failed to compile templates")
			return
		}
		log.Printf("TODO Must implement renderErrorTemplate()")
			if p := GetPage("error"); p != nil {
				if err := t.ExecuteTemplate(w, "error.html", p); err != nil {
					RespondError(w, 500, err.Error())
				}
			}
	*/
}
