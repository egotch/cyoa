// web - contains functions and types for serving basic web page
// designed to serve the story json file
package cyoa

import (
  "html/template"
  "log"
  "net/http"
  "strings"
)

var tpl *template.Template

type HandlerOptions func(h *handler)

type handler struct {
  s Story
  t *template.Template
}

// init - initialize the template var tpl using the specified html template string
func init() {
  tpl = template.Must(template.New("").Parse(DefaultHandlerTmplt))
}

// WithTemplate - Closure that sets the given template as the
// template on the handler
func WithTemplate(t *template.Template) HandlerOptions {
  return func(h *handler) {
    h.t = t
  }
}

// NewHandler: returns a http.Handler
func NewHandler(s Story, opts ...HandlerOptions) http.Handler {

  // define the default handler
  h := handler{s, tpl}

  // iterate over options
  for _, opt := range opts {
    opt(&h)
  }
  return h
}


// ServeHTTP - Method to render template (initialized in init function) for the Story passed into the handler
//
// if template fails to be rendered, panics with respective err code
// template is hard coded via the init function
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  path := strings.TrimSpace(r.URL.Path)
  if path == "" || path == "/" {
    path = "/intro"
  }

  // "/intro" => "intro"
  path = path[1:]

  // Check if chapter is present in the stories
  // if it is, render the template
  // if it is not, report not found error
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
