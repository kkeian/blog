package main

import (
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Name    string
	Header  template.HTML
	Content template.HTML
	Footer  template.HTML
}

const defaultPage = `<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Keian Kaserman {{if .Name}}- {{ .Name}} {{end}}</title>
	</head>
	<body>
		{{if .Header}} {{.Header}} {{end}}
		{{ .Content }}
		{{if .Footer}} {{ .Footer}} {{end}}
	</body>
</html>
`

func indexHandl(w http.ResponseWriter, r *http.Request) {
	outline := template.Must(template.New("out").Parse(defaultPage)) // build template
	header := `<header><h1>Keian Kaserman</h1></header>`
	content := `<article><p>Testing paragraph content here</p></article>`
	indexPage := Page{Header: template.HTML(header), Content: template.HTML(content)}

	err := outline.Execute(w, indexPage) // insert values into template
	if err != nil {
		panic(err)
	}
	return
}

func main() {

	http.HandleFunc("/", indexHandl) // use function as handler
	// serve site on designated port
	log.Fatal(http.ListenAndServe(":8080", nil))

}
