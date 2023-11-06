package http

const (
	AllowedPaths                  = `^/$|^/echo/.*$|^/user-agent$|^/files/.*$`
	CRLF                          = "\r\n"
	HttpStatusOK                  = 200
	HttpStatusNotFound            = 404
	HttpStatusInternalServerError = 500
	ContentLength                 = "Content-Length"
	ContentType                   = "Content-Type"
	ContentTypeText               = "text/plain"
	ContentTypeApplicationStream  = "application/octet-stream"
	UserAgent                     = "User-Agent"
)
