package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./index.html"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Your request method at foo: ", r.Method, "\n\n")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Your request method at bar: ", r.Method, "\n\n")
	// could process form data here
	w.Header().Set("Location", "/")
	http.Redirect(w, r, "/", http.StatusSeeOther)
	// Two lines above and below are equivalent
	w.WriteHeader(http.StatusSeeOther)
}

func barred(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at barred: ", r.Method, "\n\n")
	tpl.ExecuteTemplate(w, "index.html", nil)
}
