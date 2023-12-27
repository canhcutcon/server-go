package request

import "net/http"

type HTTPHeaderType struct {
	HTTPHeader http.Header
	headers    []string
}

type HTTPContentType struct {
	header *HTTPHeaderType
	key    string
}

func Header(kv ...string) *HTTPHeaderType {
	return &HTTPHeaderType{
		headers: kv,
	}
} // this func is for setting the header of the response

func (h *HTTPHeaderType) GetHTTPHeaders() http.Header {
	// this func is for getting the header of the response
	// the for loop uses an increase of 2 because the headers are in key value pairs
	// so the first element is the key and the second element is the value
	for i := 0; i < len(h.headers); i += 2 {
		h.HTTPHeader.Set(h.headers[i], h.headers[i+1])
	}
	return h.HTTPHeader
}

func (h *HTTPHeaderType) ContentType(key string) *HTTPContentType {
	// this func is for setting the content type of the response
	// the content type is a header, so we can use the Header func to set the header
	// content-type in headers is a key value pair, so we can use the Header func to set the header
	return &HTTPContentType{
		key: "Content-Type",
	}
}

func (h *HTTPContentType) ApplicationForWWWURLEncode() *HTTPHeaderType {
	// this func is for setting the content type of the response to application/x-www-form-urlencoded
	h.header.HTTPHeader.Add(h.key, "application/x-www-form-urlencoded")
	return h.header
}

func (h *HTTPContentType) ApplicationForJSON() *HTTPHeaderType {
	// this func is for setting the content type of the response to application/json
	h.header.HTTPHeader.Add(h.key, "application/json")
	return h.header
}
