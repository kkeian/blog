package main

import (
	"html/template"
	"log"
	"net/http"
)

type page struct {
	Header  string
	Content string
}

// ServeHTTP implements http.Handler.
func (p *page) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	page := `<!DOCTYPE html><html>
	<head>
		<meta charset="UTF-8">
		<title>Keian Kaserman</title>
	</head>
	<body>
		{{.Header}}
		{{.Content}}
	</body>
	</html`
	outline := template.Must(template.New("out").Parse(page))
	outline.Execute(w, p)
	return
}

func genBlogHeader() string {
	header := "<header><h1>Keian Kaserman</h1></header>"
	return header
}

// use raw function handler when we don't have any state to track
// on each handler invocation
func createIndexPage() page {
	const cont = `<p>Welcome!</p>`
	indexPage := page{genBlogHeader(), cont}

	return indexPage
}

func main() {
	indexP := createIndexPage()

	http.Handle("/", indexP)
	// http.HandleFunc("/") // use function as handler
	// serve site on designated port
	log.Fatal(http.ListenAndServe(":8080", nil))
}
