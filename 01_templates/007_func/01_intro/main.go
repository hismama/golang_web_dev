package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}
type car struct {
	Manufacturer string
	Model        string
	Doors        int
}
type item struct {
	Wisdom    []sage
	Transport []car
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	//!!!!  Funcs need to be there before you Parse a template <<<<<<
	//tpl := template.Must(template.New("something").Parse("Here is my template, yo"))
	//tpl.ExecuteTemplate(os.Stdout, "something", nil)
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("tpl.gohtml"))
	//	template.New gives pointer to a template blank
	//  Funcs fills the pointer to the template
	//  ParseFiles gives back pointer to template and error
	//	Wrapped in Must returns pointer to template
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}
	j := sage{
		Name:  "Jesus",
		Motto: "Loves all",
	}
	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        4,
	}
	c := car{
		Manufacturer: "Chevy",
		Model:        "Silverado",
		Doors:        4,
	}
	sages := []sage{b, j}
	cars := []car{f, c}

	data := item{sages, cars}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
