package main

import (
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
			continue // remember to continue, if not it will still try to write to connection
		}

		io.WriteString(conn, "\nHello mate\n") // Can use both io.WriteString and fmt.Fprint
		fmt.Fprintln(conn, "\nHope you are doing alright mate")
		conn.Close()
	}
}
