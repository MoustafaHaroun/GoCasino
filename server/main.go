package main

import (
	"bufio"
	"log"
	"net"
	"net/http"
)

func main() {
	// Start listing on port 8080.
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// Clean up the port listener.
	defer func() {
		if err := listener.Close(); err != nil {
			log.Println("Error closing listener:", err)
		}
	}()

	// Start listing for connections.
	log.Println("Listening on :8080")
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
	// Clean connection.
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("Error closing connection:", err)
		}
	}()

	reader := bufio.NewReader(conn)
	request, _ := http.ReadRequest(reader)

	response, err := handshake(request)
	if err != nil {
		log.Println("Error in handshake:", err)
		if err := conn.Close(); err != nil {
			log.Println("Error closing connection after handshake error:", err)
		}
		return
	}

	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Println("Error writing response:", err)
	}
}
