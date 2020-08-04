package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type car struct {
	Name    string
	Country string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	nissan := car{
		Name:    "Nissan",
		Country: "Japan",
	}

	err := tpl.Execute(os.Stdout, nissan)
	if err != nil {
		log.Fatalln(err)
	}
}
