package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm() // Must be called before req.Form
	if err != nil {
		log.Fatalln(err)
	}

	//data := struct {
	//	Method        string
	//	URL           *url.URL
	//	Submissions   url.Values
	//	Header        http.Header
	//	Host          string
	//	ContentLength int64
	//}{
	//	req.Method,
	//	req.URL,
	//	req.Form,
	//	req.Header,
	//	req.Host,
	//	req.ContentLength,
	//}
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
