package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml")) // Best way to Parse Multiple Files into Container
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Release self-focus; embrace other-focus.")
	if err != nil {
		log.Fatalln(err)
	}

}
