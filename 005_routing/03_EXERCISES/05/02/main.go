package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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
		defer conn.Close()

		fmt.Fprintln(conn, "\nHave you got something to say?")

		s := bufio.NewScanner(conn)
		for s.Scan() {
			ln := s.Text()
			fmt.Println(ln)
			fmt.Fprintf(conn, "You said: %s\n", ln)
		}

		// will never get here
		// have open stream
		// reader/scanner will never be done
		fmt.Println("code got here")
		io.WriteString(conn, "I see you connected")
	}
}
