package error

// HttpErrorInterface HTTP error interface
type HttpErrorInterface interface {
	Code() int
	Message() string
	BadRequest(stacktrace ...string) *HTTP
	Unauthorized(stacktrace ...string) *HTTP
	Forbidden(stacktrace ...string) *HTTP
	NotFound(stacktrace ...string) *HTTP
	Conflict(stacktrace ...string) *HTTP
	InternalServerError(stacktrace ...string) *HTTP
}

// HTTP http error struct
type HTTP struct {
	code       int
	message    string
	stacktrace string
}

// Code return http error status code
func (http *HTTP) Code() int {
	return http.code
}

// Message return http error message
func (http *HTTP) Message() string {
	return http.message
}

// Stacktrace return http error stacktrace
func (http *HTTP) Stacktrace() string {
	return http.stacktrace
}

// BadRequest return http 400 bad request error data
func (http *HTTP) BadRequest(stacktrace ...string) *HTTP {
	return &HTTP{code: 400, message: "Bad request", stacktrace: stacktrace[0]}
}

// Unauthorized return http 401 unauthorized error data
func (http *HTTP) Unauthorized(stacktrace ...string) *HTTP {
	return &HTTP{code: 401, message: "Unauthorized", stacktrace: stacktrace[0]}
}

// Forbidden return http 403 forbidden error data
func (http *HTTP) Forbidden(stacktrace ...string) *HTTP {
	return &HTTP{code: 401, message: "Forbidden", stacktrace: stacktrace[0]}
}

// NotFound return http 404 not found error data
func (http *HTTP) NotFound(stacktrace ...string) *HTTP {
	return &HTTP{code: 404, message: "Not found", stacktrace: stacktrace[0]}
}

// Conflict return http 209 conflict error data
func (http *HTTP) Conflict(stacktrace ...string) *HTTP {
	return &HTTP{code: 409, message: "Conflict", stacktrace: stacktrace[0]}
}

// InternalServerError return http 500 internal server error data
func (http *HTTP) InternalServerError(stacktrace ...string) *HTTP {
	return &HTTP{code: 500, message: "Internal server error", stacktrace: stacktrace[0]}
}
