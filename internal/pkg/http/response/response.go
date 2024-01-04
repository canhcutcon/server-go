package response

import (
	"encoding/json"
	"net/http"
)

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

// JSONRetry is a struct for retrying the request
// RetryAfter is the time in seconds after which the request should be retried
type JSONRetryResponse struct {
	RetryMin int `json:"retry_min"`
	RetryMax int `json:"retry_max"`
}

// JSON create a new JSON response
func JSON(w http.ResponseWriter) *JSONResponse {
	return &JSONResponse{
		writer: w,
	}
}

func (jresp *JSONResponse) SetHeader(key, value string) {
	jresp.writer.Header().Set(key, value)
}

func (jresp *JSONResponse) Data(data interface{}) *JSONResponse {
	jresp.ResponseData = data
	return jresp
}

func (jresp *JSONResponse) Error(err error, errResp *JSONError) *JSONResponse {
	jresp.error = err
	jresp.ResponseError = errResp
	return jresp
}

func (jresp *JSONResponse) WriteHeader(status int) *JSONResponse {
	if jresp.headerWritten {
		return jresp
	}
	jresp.headerWritten = true
	jresp.writer.WriteHeader(status)
	return jresp
}

func (jresp *JSONResponse) Write() (int, error) {
	jresp.writer.Header().Set("Content-Type", "application/json")

	//process the internal error
	if jresp.error != nil {
		switch jresp.ResponseStatus {
		// case StatusBadRequest:
		// 	jresp.ResponseStatus = StatusBadRequest
		// 	jresp.WriteHeader(http.StatusBadRequest)
		// case ErrUnauthorized:
		// 	jresp.ResponseStatus = StatusUnauthorized
		// 	jresp.WriteHeader(http.StatusUnauthorized)
		// case ErrForbidden:
		// 	jresp.ResponseStatus = StatusForbidden
		// 	jresp.WriteHeader(http.StatusForbidden)
		// case ErrNotFound:
		// 	jresp.ResponseStatus = StatusNotFound
		// 	jresp.WriteHeader(http.StatusNotFound)
		// case ErrConflict:
		// 	jresp.ResponseStatus = StatusConflict
		// 	jresp.WriteHeader(http.StatusConflict)
		// case ErrInternal:
		// 	jresp.ResponseStatus = StatusInternal
		// 	jresp.WriteHeader(http.StatusInternalServerError)
		// case ErrTimeout:
		// 	jresp.ResponseStatus = StatusTimeout
		// 	jresp.WriteHeader(http.StatusGatewayTimeout)
		}

	} else {
		jresp.ResponseStatus = StatusOK
		jresp.WriteHeader(http.StatusOK)
	}

	out, err := json.Marshal(jresp)
	if err != nil {
		return 0, err
	}

	return jresp.writer.Write(out)
}
