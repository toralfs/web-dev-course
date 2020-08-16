package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", peng)
	http.HandleFunc("/penguin.jpg", pengPic)
	http.ListenAndServe(":8080", nil)
}

func peng(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="../penguin.jpg"> `)
}

func pengPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "../penguin.jpg")
}
