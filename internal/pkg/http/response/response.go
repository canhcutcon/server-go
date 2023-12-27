package response

import "net/http"

type Status string

const (
	StatusOK           Status = "OK"
	StatusBadRequest   Status = "BAD_REQUEST"
	StatusUnauthorized Status = "UNAUTHORIZED"
	StatusForbidden    Status = "FORBIDDEN"
	StatusNotFound     Status = "NOT_FOUND"
	StatusConflict     Status = "CONFLICT"
	StatusInternal     Status = "INTERNAL"
	StatusTimeout      Status = "TIMEOUT"
	StatusUnknown      Status = "UNKNOWN"
	StatusRetry        Status = "RETRY"
)

type JSONError struct {
	Title   string   `json:"title"`
	Detail  string   `json:"detail"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

type JSONRetry struct {
	RetryAfter int `json:"retry_after"`
	RetryMin   int `json:"retry_min"`
	RetryMax   int `json:"retry_max"`
}

type JSONResponse struct {
	writer        http.ResponseWriter
	error         error
	headerWritten bool

	ResponseStatus Status      `json:"status"`
	ResponseData   interface{} `json:"data"`
	ResponseError  *JSONError  `json:"error, omitempty"`
	ResponseRetry  *JSONRetry  `json:"retry, omitempty"`
}
