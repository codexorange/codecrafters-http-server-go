package http

import (
	"fmt"
	"net"
)

func ServeHTTP(conn net.Conn, dir string) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Failed to read from socket:", err)
		return
	}

	response := HandleRequest(buf[:n], dir)
	response.WriteResponse(conn)
}
