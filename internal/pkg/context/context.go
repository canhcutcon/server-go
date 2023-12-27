package context

import (
	"context"
	"net/http"
)

type RequestContext struct {
	httpResponseWriter http.ResponseWriter // this is the response writer, we can use this to write the response
	httpRequest        *http.Request       // this is the request, we can use this to get the request data, like query params, body, etc
	address            string              // this is the address of the request
	path               string              // this is the path of the request
	method             string
}

type RequestConst struct {
	HTTPResponse http.ResponseWriter
	HTTPRequest  *http.Request
	Address      string
	Path         string
	Method       string
}

func NewRequestContext(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) *RequestContext {
	return &RequestContext{
		httpResponseWriter: httpResponseWriter,
		httpRequest:        httpRequest,
		address:            httpRequest.RemoteAddr,
		path:               httpRequest.URL.Path,
		method:             httpRequest.Method,
	}
}

func (r *RequestContext) GetHTTPResponseWriter() http.ResponseWriter {
	return r.httpResponseWriter
}

func (r *RequestContext) GetHTTPRequest() *http.Request {
	return r.httpRequest
}

func (r *RequestContext) GetAddress() string {
	return r.address
}

func (r *RequestContext) GetPath() string {
	return r.path
}

func (r *RequestContext) GetMethod() string {
	return r.method
}

func (r *RequestContext) GetRequestConst() *RequestConst {
	return &RequestConst{
		HTTPResponse: r.httpResponseWriter,
		HTTPRequest:  r.httpRequest,
		Address:      r.address,
		Path:         r.path,
		Method:       r.method,
	}
}

func (r *RequestContext) Context() context.Context {
	return r.httpRequest.Context()
}

func (r *RequestContext) RequestHeader() http.Header {
	return r.httpRequest.Header
}

func (r *RequestContext) ResponseHeader() http.Header {
	return r.httpResponseWriter.Header()
}

func (r *RequestContext) RequestHandler() string {
	return r.path
}

// JSON to create a json response via http.ResponseWriter
func (r *RequestContext) JSON(statusCode int, data interface{}) {
	r.httpResponseWriter.Header().Set("Content-Type", "application/json")
	r.httpResponseWriter.WriteHeader(statusCode)
	r.httpResponseWriter.Write([]byte(data.(string)))
}

func (r *RequestContext) JSONBytes(statusCode int, data []byte) {
	r.httpResponseWriter.Header().Set("Content-Type", "application/json")
	r.httpResponseWriter.WriteHeader(statusCode)
	r.httpResponseWriter.Write(data)
}

// func (r *RequestContext) GetJSON() *response.JSONResponse {
// 	// return response.NewJSONResponse(r.httpResponseWriter)
// }
