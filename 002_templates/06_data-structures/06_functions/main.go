package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

type car struct {
	Name    string
	Country string
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
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

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", cars)
	if err != nil {
		log.Fatalln(err)
	}
}
