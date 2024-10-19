package main

import (
	"log"
	"os"
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

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml")) // Best way to Parse Multiple Files into Container
}

func main() {
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
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
	sages := []sage{b}
	cars := []car{f, c}

	data := item{sages, cars}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
