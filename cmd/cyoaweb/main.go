package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/egotch/cyoa"
)

func main() {

	//flags
	filePtr := flag.String("story", "gopher.json", "path + file name of json file containing story data")
	portPtr := flag.Int("port", 3000, "the port to start the CYOA web app on")

	flag.Parse()

	// open the json file and read it in
	fmt.Printf("\nUsing story file - %s\n", *filePtr)
	file, err := os.Open(*filePtr)
	if err != nil {
		panic(err)
	}

	// parse the json file >> story struct
	story, err := cyoa.JsonStory(file)

	// create our http handler
	tpl := template.Must(template.New("").Parse(storyTmpl))
	h := cyoa.NewHandler(
		story,
		cyoa.WithTemplate(tpl),
		cyoa.WithChapterParser(cstmChapterParser),
	)

	// Create a ServeMux to route our requests
	mux := http.NewServeMux()
	// This story handler is using a custom function and template
	// Because we use /story/ (trailing slash) all web requests
	// whose path has the /story/ prefix will be routed here.
	mux.Handle("/story/", h)
	// This story handler is using the default functions and templates
	// Because we use / (base path) all incoming requests not
	// mapped elsewhere will be sent here.
	mux.Handle("/", cyoa.NewHandler(story))
	// Start the server using our ServeMux
	fmt.Printf("Starting the server on port: %d\n", *portPtr)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *portPtr), mux))

}

// Updated chapter parsing function.
func cstmChapterParser(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

// Slightly altered tempalte to show how this feature works
var storyTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      <ul>
      {{range .Options}}
        <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
      {{end}}
      </ul>
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FCF6FC;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #797;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: underline;
        color: #555;
      }
      a:active,
      a:hover {
        color: #222;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`
