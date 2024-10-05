// web - contains functions and types for serving basic web page
// designed to serve the story json file
package cyoa

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

type handler struct {
  s Story
}

// init - initialize the template var tpl using the specified html template string
func init() {
  tpl = template.Must(template.New("").Parse(DefaultHandlerTmplt))
}


// NewHandler: returns a http.Handler
func NewHandler(s Story) http.Handler {
  return handler{s}
}


// ServeHTTP - Method to render template (initialized in init function) for the Story passed into the handler
//
// if template fails to be rendered, panics with respective err code
// template is hard coded via the init function
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  err := tpl.Execute(w, h.s["intro"])
  if err != nil {
    panic(err)
  }
}
