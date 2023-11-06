package http

import (
	"fmt"
	"net"
	"strings"
)

var HttpStatusLines = map[int]string{
	HttpStatusOK:                  "HTTP/1.1 200 OK",
	HttpStatusNotFound:            "HTTP/1.1 404 Not Found",
	HttpStatusInternalServerError: "HTTP/1.1 500 Internal Server Error",
	HttpStatusCreated:             "HTTP/1.1 201 Created",
}

type HttpResponse struct {
	StatusCode int
	StatusLine string
	Headers    map[string]string
	Body       []byte
}

func NewResponse(request *HttpRequest, body string, statusCode int) *HttpResponse {
	var response *HttpResponse = &HttpResponse{
		StatusCode: statusCode,
		StatusLine: HttpStatusLines[statusCode],
		Headers:    make(map[string]string),
	}

	response.Headers[ContentType] = ContentTypeText
	if body != "" {
		response.Headers[ContentLength] = fmt.Sprintf("%d", len(body))
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
