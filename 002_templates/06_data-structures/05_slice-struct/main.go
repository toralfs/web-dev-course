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

	toyota := car{
		Name:    "Toyota",
		Country: "Japan",
	}

	chevrolet := car{
		Name:    "Chevrolet",
		Country: "USA",
	}

	volvo := car{
		Name:    "Volvo",
		Country: "Sweden",
	}

	bmw := car{
		Name:    "BMW",
		Country: "Germany",
	}

	cars := []car{nissan, toyota, chevrolet, volvo, bmw}

	err := tpl.Execute(os.Stdout, cars)
	if err != nil {
		log.Fatalln(err)
	}
}
