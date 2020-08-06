package main

import (
	"fmt"
	"io"
	"net/http"
)

type penguin int

func (m penguin) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
	fmt.Println(req.URL.Path)
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "doggo doggo doggo")
	case "/cat":
		io.WriteString(res, "mew mew mew")
	}
}

func main() {
	var d penguin
	http.ListenAndServe(":8080", d)
}
