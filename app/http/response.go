package http

import (
	"fmt"
	"net"
)

var HttpStatusLines = map[int]string{
	200: "HTTP/1.1 200 OK",
	404: "HTTP/1.1 404 Not Found",
}

type HttpResponse struct {
	StatusCode int
	StatusLine string
	Headers    map[string]string
	Body       []byte
}

func NewResponse(request *HttpRequest, statusCode int) *HttpResponse {
	var response *HttpResponse = &HttpResponse{}
	response.StatusCode = statusCode
	response.StatusLine = HttpStatusLines[statusCode]
	response.Headers = make(map[string]string)

	return response
}

func (response *HttpResponse) WriteResponse(conn net.Conn) {
	fmt.Println("Writing response to socket")
	defer conn.Close()

	_, err := conn.Write([]byte(response.StatusLine + CRLF + CRLF))
	if err != nil {
		fmt.Println("Failed to write to socket:", err)
		return
	}
}
