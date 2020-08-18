package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/penguin", penguin)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func penguin(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="pic/penguin.jpg">`)
}
