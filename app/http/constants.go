package http

const (
	AllowedPaths                  = `^/$|^/echo/.*$|^/user-agent$|^/files/.*$`
	CRLF                          = "\r\n"
	HttpStatusOK                  = 200
	HttpStatusCreated             = 201
	HttpStatusNotFound            = 404
	HttpStatusMethodNotAllowed    = 405
	HttpStatusInternalServerError = 500
	ContentLength                 = "Content-Length"
	ContentType                   = "Content-Type"
	ContentTypeText               = "text/plain"
	ContentTypeApplicationStream  = "application/octet-stream"
	UserAgent                     = "User-Agent"
)
