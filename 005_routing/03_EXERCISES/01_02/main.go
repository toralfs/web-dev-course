package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "dog.gohtml", nil)
}

func me(res http.ResponseWriter, req *http.Request) {
	name := "Toralf"
	tpl.ExecuteTemplate(res, "me.gohtml", name)
	io.WriteString(res, fmt.Sprintf("My name is %s", name))
}
