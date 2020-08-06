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

		fmt.Fprintln(conn, "\nHave you got something to say?")

		s := bufio.NewScanner(conn)
		for s.Scan() {
			ln := s.Text()
			if ln == "" {
				// when ln is empty, header is done
				fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
				break
			}
			fmt.Println(ln)
			fmt.Fprintf(conn, "You said: %s\n", ln)
		}

		fmt.Println("code got here")
		io.WriteString(conn, "I see you connected")
		conn.Close()
	}
}
