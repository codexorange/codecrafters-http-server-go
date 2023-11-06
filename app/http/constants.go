package http

const (
	AllowedPaths       = `^/echo/.*$|^/$|^/user-agent$`
	CRLF               = "\r\n"
	HttpStatusOK       = 200
	HttpStatusNotFound = 404
)
