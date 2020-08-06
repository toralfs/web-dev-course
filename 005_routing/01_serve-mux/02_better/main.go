package main

import (
	"io"
	"net/http"
)

type doggo int

func (d doggo) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "it's me dog mate")
}

type catto int

func (c catto) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "mew mew mew")
}

func main() {
	var d doggo
	var c catto
	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat/", c)

	http.ListenAndServe(":8080", mux)
}
