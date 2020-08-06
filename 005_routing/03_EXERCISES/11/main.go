package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		go serve(conn)
	}
}

func serve(c net.Conn) {
	defer c.Close()
	s := bufio.NewScanner(c)
	var i int
	for s.Scan() {
		ln := s.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(c, ln)
		}
		if ln == "" {
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
		i++
	}
}

func mux(c net.Conn, ln string) {
	// request line
	xs := strings.Fields(ln)
	m := xs[0]
	u := xs[1]
	fmt.Println("METHOD:", m)
	fmt.Println("URI:", u)

	// multiplexer
	switch {
	case m == "GET" && u == "/":
		index(c)
	case m == "GET" && u == "/apply":
		apply(c)
	case m == "POST" && u == "/apply":
		applyProcess(c)
	default:
		notFound(c)
	}

}

func index(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Index</title>
		</head>
		<body>
		<h1>Welcome to index</h1>
		<h2>Select from these wonderful choises</h2>
		<a href="/">index</a><br>
		<a href="/apply">apply</a><br>
		</body>
		</html>
	`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Lenght: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}

func apply(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Apply</title>
		</head>
		<body>
		<a href="/">index</a><br>
		<a href="/apply">apply</a><br>
		<h1>Psst.. You want good deal?</h1>
		<h2>Just hit apply</h2>
		<form method="POST" action="/apply">
			<input type="submit" value="apply">
		</form>
		</body>
		</html>
	`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Lenght: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}

func applyProcess(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Applied</title>
		</head>
		<body>
		<a href="/">index</a><br>
		<a href="/apply">apply</a><br>
		<h1>Good boi</h1>
		<h2>You made good deal</h2>
		</body>
		</html>
	`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Lenght: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}

func notFound(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Applied</title>
		</head>
		<body>
		<p><strong>404 Not Found</strong></p>
		</body>
		</html>
	`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Lenght: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}
