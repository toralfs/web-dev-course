package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "me-cookie",
		Value: "me value",
	})

	incrementCookieCounter(w, req)

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
	fmt.Fprintln(w, "COOKIE WRITTEN - CHEK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome, go to dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	incrementCookieCounter(w, req)
	readCookies(w, req, "me-cookie", "general", "specific", "count")
}

func abundance(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "general value",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "specific value",
	})
	incrementCookieCounter(w, req)
	fmt.Fprintln(w, "MORE COOKIES, OMNOMNOMNOM")
}

func readCookies(w http.ResponseWriter, req *http.Request, cookies ...string) {
	for i, cookie := range cookies {
		c, err := req.Cookie(cookie)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			continue
		}
		fmt.Fprintf(w, "YOUR COOKIE #%d: %v\n", i+1, c)
	}
}

func incrementCookieCounter(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("count")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "count",
			Value: "0",
		}
	}

	i, err := strconv.Atoi(c.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	i++
	c.Value = strconv.Itoa(i)
	http.SetCookie(w, c)
}
