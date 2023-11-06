package http

import (
	"fmt"
	"strings"
)

type HttpRequest struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    []byte
}

func HandleRequest(request []byte) *HttpResponse {
	var statusCode int = HttpStatusOK
	httpRequest := ParseHttpRequest(request)

	if !strings.Contains(AllowedPaths, httpRequest.Path) {
		statusCode = HttpStatusNotFound
	}

	return NewResponse(httpRequest, statusCode)
}

func ParseHttpRequest(request []byte) *HttpRequest {
	var httpRequest *HttpRequest = &HttpRequest{}

	req := string(request)
	lines := strings.Split(req, CRLF)

	for i, line := range lines {
		if i == 0 {
			parts := strings.Split(line, " ")
			httpRequest.Method = parts[0]
			httpRequest.Path = parts[1]
			httpRequest.Version = parts[2]
		} else {
			parts := strings.Split(line, ": ")
			if len(parts) > 1 {
				if httpRequest.Headers == nil {
					httpRequest.Headers = make(map[string]string)
				}
				httpRequest.Headers[parts[0]] = parts[1]
			}
		}
	}

	fmt.Println(httpRequest)
	return httpRequest
}
