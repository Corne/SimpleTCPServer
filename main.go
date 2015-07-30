package main

import (
	"fmt"
	"net"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	defer l.Close()
	fmt.Printf("Listening on %s:%s\n", CONN_HOST, CONN_PORT)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("Error accepting: %s\n", err.Error())
			continue
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	conn.Write([]byte("Hello."))
	conn.Close()
}
