package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template
var fm = template.FuncMap{
	"fdouble": double,
	"fsquare": square,
	"fsqroot": sqRoot,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*"))
}

func double(x int) int {
	return x * 2
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}
}
