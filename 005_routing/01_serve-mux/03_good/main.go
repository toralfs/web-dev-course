package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "it's me dog mate")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "mew mew mew")
}

func main() {
	http.HandleFunc("/dog/", d)

	// Just to show what a "HandlerFunc" is:
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil) // uses the default ServeMux
}
