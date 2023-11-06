package http

import (
	"regexp"
	"strings"
)

type HttpRequest struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    []byte
}

func HandleRequest(request []byte, dir string) *HttpResponse {
	httpRequest := ParseHttpRequest(request)
	statusCode := ValidatePath(httpRequest.Path)

	return HandleRoute(httpRequest, statusCode, dir)
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

	return httpRequest
}

func ValidatePath(path string) int {
	validPathRegex := regexp.MustCompile(AllowedPaths)
	if validPathRegex.MatchString(path) {
		return HttpStatusOK
	} else {
		return HttpStatusNotFound
	}
}
