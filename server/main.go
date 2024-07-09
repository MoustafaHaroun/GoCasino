package main

import (
	"bufio"
	"log"
	"net"
	"net/http"
)

func main() {
	// Start listing on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	request, _ := http.ReadRequest(reader)

	response, err := handshake(request)
	if err != nil {
		log.Println(err)
		conn.Close()
	}

	conn.Write([]byte(response))
}
