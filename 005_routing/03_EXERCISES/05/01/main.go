package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// This itteration can only accept 1 connection at a time
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
			fmt.Fprintf(conn, "You said: %s\n", ln)
		}
	}
}
