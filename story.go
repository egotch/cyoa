package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

var tpl *template.Template

func init() {
  tpl = template.Must(template.New("").Parse(defaultHandlerTmplt))
}

var defaultHandlerTmplt string = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Choose Your Own Adventure</title>
    <link href="css/style.css" rel="stylesheet">
  </head>
  <body>
  <h1>{{.Title}}</h1> 
    {{range .Paragraphs}}
      <p>{{.}}</p>
    {{end}}
    <ul>
      {{range .Options}}
      <li><a href="/{{.Arc}}">{{.Text}}</a></li>
      {{end}}
    </ul>
  </body>
</html>
`

func NewHandler(s Story) http.Handler {
  return handler{s}
}

type handler struct {
  s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  err := tpl.Execute(w, h.s["intro"])
  if err != nil {
    panic(err)
  }
}

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	err := d.Decode(&story)

	if err != nil {
		return nil, err
	} else {
		return story, nil
	}

}
