package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("penguin.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/penguin", penguin)
	http.HandleFunc("/penguin.jpg", pengPic)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func penguin(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "penguin.gohtml", nil)
}

func pengPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "penguin.jpg")
}
