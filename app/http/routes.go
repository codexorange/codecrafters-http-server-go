package http

import (
	"flag"
	"io"
	"os"
	"path"
	"strings"
)

type RouteHandler interface {
	Handle(request *HttpRequest) *HttpResponse
}

type EchoHandler struct{}
type RootHandler struct{}
type UserAgentHandler struct{}
type NotFoundHandler struct{}
type FileHandler struct{}

func HandleRoute(request *HttpRequest, statusCode int) *HttpResponse {
	if request.Path == "/" {
		return RootHandler{}.Handle(request)
	} else if strings.Contains(request.Path, "/echo/") {
		return EchoHandler{}.Handle(request)
	} else if request.Path == "/user-agent" {
		return UserAgentHandler{}.Handle(request)
	} else if strings.Contains(request.Path, "/files/") {
		return FileHandler{}.Handle(request)
	} else {
		return NotFoundHandler{}.Handle(request)
	}
}

func (handler RootHandler) Handle(request *HttpRequest) *HttpResponse {
	body := ""
	return NewResponse(request, body, HttpStatusOK)
}

func (handler EchoHandler) Handle(request *HttpRequest) *HttpResponse {
	body := strings.Replace(request.Path, "/echo/", "", 1)
	return NewResponse(request, body, HttpStatusOK)
}

func (handler UserAgentHandler) Handle(request *HttpRequest) *HttpResponse {
	body := request.Headers[UserAgent]
	return NewResponse(request, body, HttpStatusOK)
}

func (handler NotFoundHandler) Handle(request *HttpRequest) *HttpResponse {
	body := "Not Found"
	return NewResponse(request, body, HttpStatusNotFound)
}

func (handler FileHandler) Handle(request *HttpRequest) *HttpResponse {
	filename := strings.Replace(request.Path, "/files/", "", 1)
	dir := flag.String("directory", ".", "Directory to serve")
	flag.Parse()

	path := path.Join(*dir, filename)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		body := "File Not Found"
		return NewResponse(request, body, HttpStatusNotFound)
	}

	file, err := os.Open(path)
	if err != nil {
		body := "Error opening file"
		return NewResponse(request, body, HttpStatusInternalServerError)
	}
	defer file.Close()

	fileContents, err := io.ReadAll(file)
	if err != nil {
		body := "Error reading file"
		return NewResponse(request, body, HttpStatusInternalServerError)
	}

	body := string(fileContents)
	response := NewResponse(request, body, HttpStatusOK)
	response.Headers[ContentType] = ContentTypeApplicationStream

	return response
}
