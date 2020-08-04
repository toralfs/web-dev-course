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
	comicChars := []string{"Donald Duck", "Dolly Duck", "Scrooge McDuck", "Grandma Duck"}

	err := tpl.Execute(os.Stdout, comicChars)
	if err != nil {
		log.Fatalln(err)
	}
}
