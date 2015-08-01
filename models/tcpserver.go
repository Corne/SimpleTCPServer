package tcpserver

import (
	"fmt"
	"net"
	"os"
)

type TCPServer struct {
	Host     string
	Port     int
	channels []chan string
}

func Create(host string, port int) *TCPServer {
	return &TCPServer{
		Host:     host,
		Port:     port,
		channels: []chan string{},
	}
}

func (server *TCPServer) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Host, server.Port))
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("Listening on %s:%d\n", server.Host, server.Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting: %s\n", err.Error())
			continue
		}

		go server.handleRequest(conn)
	}
}

func (server TCPServer) Broadcast(message string) {
	for _, c := range server.channels {
		c <- message
	}
}

func (server *TCPServer) handleRequest(conn net.Conn) {
	fmt.Println("New Connection!")
	c := make(chan string)
	server.channels = append(server.channels, c)

	for {
		message := <-c
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			break
		}
	}

	i := server.indexOf(c)
	server.channels = append(server.channels[:i], server.channels[i+1:]...)

	conn.Close()
	close(c)
	fmt.Println("Connection closed!")
}

func (server TCPServer) indexOf(val chan string) int {
	for i, el := range server.channels {
		if el == val {
			return i
		}
	}
	return -1
}
