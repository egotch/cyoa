// houses html, json, and other templates for use in project
package cyoa


var DefaultHandlerTmplt string = `
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
