package main

import (
	"fmt"
	"net/http"
)

type penguin int

func (m penguin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var d penguin
	http.ListenAndServe(":8080", d)
}
