package main

import (
	"fmt"
	"net"
	"os"
)

const (
	StatusLine = "HTTP/1.1 200 OK"
	CRLF       = "\r\n"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Failed to read from socket:", err)
		return
	}

	_, err = conn.Write([]byte(StatusLine + CRLF + CRLF))
	if err != nil {
		fmt.Println("Failed to write to socket:", err)
		return
	}
	fmt.Printf("Received %d bytes: %s\n", n, buf[:n])
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection")
			os.Exit(1)
		}

		go handleConnection(conn)
	}
}
