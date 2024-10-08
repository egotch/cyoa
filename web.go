// web - contains functions and types for serving basic web page
// designed to serve the story json file
package cyoa

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

// init - initialize the template var tpl using the specified html template string
func init() {
	tpl = template.Must(template.New("").Parse(DefaultHandlerTmplt))
}

var tpl *template.Template

// HandlerOptions - used with NewHandler function to
// configure the http.Handler returned
type HandlerOptions func(h *handler)

// WithTemplate - is an option to provide a custom html
// template to be used in rendering stories
func WithTemplate(t *template.Template) HandlerOptions {
	return func(h *handler) {
		h.t = t
	}
}

// WithPathFunc - option to provide a custom function
// for processing the story chapter from the incomming request
func WithChapterParser(fn func(r *http.Request) string) HandlerOptions {
	return func(h *handler) {
		h.chapterParser = fn
	}
}

// NewHandler: returns a http.Handler
func NewHandler(s Story, opts ...HandlerOptions) http.Handler {

	// define the default handler
	h := handler{s, tpl, defaultPathFn}

	// iterate over options
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type handler struct {
	s             Story
	t             *template.Template
	chapterParser func(r *http.Request) string
}

// defaultPathFn - is the default chapter
// parsing function that expects patterns
// "/intro"
func defaultPathFn(r *http.Request) string {

	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}

	return path[1:]
}

// ServeHTTP - Method to render template (initialized in init function) for the Story passed into the handler
//
// if template fails to be rendered, panics with respective err code
// template is hard coded via the init function
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := h.chapterParser(r)

	if chapter, ok := h.s[path]; ok {
		err := tpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Ruh Roh ...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found", http.StatusNotFound)

}
