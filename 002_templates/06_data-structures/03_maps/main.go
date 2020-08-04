package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	cars := map[string]string{
		"Toyota":     "Japan",
		"Nissan":     "Japan",
		"Chevrolet":  "USA",
		"Volkswagen": "Germany",
		"Volvo":      "Sweden",
	}

	err := tpl.Execute(os.Stdout, cars)
	if err != nil {
		log.Fatalln(err)
	}
}
