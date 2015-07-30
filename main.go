package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	CONN_HOST = "localhost"
	CONN_TYPE = "tcp"
)

var (
	channels []chan string = []chan string{}
)

func main() {

	//read port
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Port: ")
	port, err := reader.ReadString('\n')
	port = strings.TrimSpace(port)

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
	// listen on given  localhost to given port
	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+port)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	defer listener.Close()
	fmt.Printf("Listening on %s:%s\n", CONN_HOST, port)

	go readInput()
	// handle each incoming request
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting: %s\n", err.Error())
			continue
		}

		go handleRequest(conn)
	}
}

func readInput() {
	for {
		reader := bufio.NewReader(os.Stdin)
		message, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			return
		}
		for _, c := range channels {
			c <- message
		}
	}
}

func handleRequest(conn net.Conn) {
	c := make(chan string)
	channels = append(channels, c)

	for {
		message := <-c
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			break
		}
	}

	i := indexOf(c)
	fmt.Println(i)
	channels = append(channels[:i], channels[i+1:]...)

	conn.Close()
}

func indexOf(val chan string) int {
	for i, el := range channels {
		if el == val {
			return i
		}
	}
	return -1
}
