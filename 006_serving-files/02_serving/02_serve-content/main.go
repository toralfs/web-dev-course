package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", peng)
	http.HandleFunc("/penguin.jpg", pengPic)
	http.ListenAndServe(":8080", nil)
}

func peng(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="../penguin.jpg">`)
}

func pengPic(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("../penguin.jpg")
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}

	http.ServeContent(res, req, f.Name(), fi.ModTime(), f)
}
