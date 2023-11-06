package http

import (
	"fmt"
	"net"
	"strings"
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
	response.Headers["Content-Type"] = "text/plain"

	if strings.Contains(request.Path, "/echo/") {
		body := strings.Replace(request.Path, "/echo/", "", 1)
		response.Headers["Content-Length"] = fmt.Sprintf("%d", len(body))
		response.Body = []byte(body)
	}

	if strings.Contains(request.Path, "/user-agent") {
		body := request.Headers["User-Agent"]
		response.Headers["Content-Length"] = fmt.Sprintf("%d", len(body))
		response.Body = []byte(body)
	}

	return response
}

func (response *HttpResponse) WriteResponse(conn net.Conn) {
	defer conn.Close()

	var out strings.Builder
	out.WriteString(response.StatusLine + CRLF)
	for header, value := range response.Headers {
		out.WriteString(header + ": " + value + CRLF)
	}
	out.WriteString(CRLF)
	out.Write(response.Body)

	_, err := conn.Write([]byte(out.String()))
	if err != nil {
		fmt.Println("Failed to write to socket:", err)
		return
	}
}
