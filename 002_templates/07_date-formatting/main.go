package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template
var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
	"fdateDMY": dayMonthYear,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func dayMonthYear(t time.Time) string {
	return t.Format("02-01-2006")
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
