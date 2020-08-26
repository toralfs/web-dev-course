package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar: ", req.Method)

	// Can set header location
	//w.Header().Set("Location", "/")
	//w.WriteHeader(http.StatusSeeOther)

	// Or use http.Redirect
	http.Redirect(w, req, "/", http.StatusSeeOther)
	// http.StatusMovedPermanently to make browser always redirect, without making request to server
	// http.StatusTemporaryRedirect to redirect while keeping Method (like POST)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred: ", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
