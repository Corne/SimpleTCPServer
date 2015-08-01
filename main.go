package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/corne/tcpserver/models"
)

func main() {

	//read port
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Port: ")
	readed, err := reader.ReadString('\n')
	readed = strings.TrimSpace(readed)

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	port, err := strconv.ParseInt(readed, 0, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	server := tcpserver.Create("localhost", int(port))
	go server.Start()
	readInput(server)

}

func readInput(server *tcpserver.TCPServer) {
	for {
		reader := bufio.NewReader(os.Stdin)
		message, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			return
		}
		server.Broadcast(message)
	}
}
