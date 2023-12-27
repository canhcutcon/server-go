package request

// this package is a wrapper for http request, such as method, url, query, body, etc.
import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var (
	// flag to activate request url and forwarded headers for specific version
	_requestVersionMatching = true

	// RoutingContext is a key name for version selection routing
	RoutingContext = "REQUEST_ROUTING_HEADER"
)

// the main type , it holds all the information for the http request.
// It also have context.Context field, which can be uses to carry deadline, cancellation signal, etc.
// it also have a http.request field which can be used to send the request to the target url and get the response back

type RequestType struct {
	method           string
	url              string
	query            string
	header           http.Header
	additionalHeader []string

	body      io.Reader
	vBody     interface{}
	bodyJSON  bool
	noVersion bool

	ctx context.Context
	r   *http.Request
}

func NewRequest(ctx context.Context) *RequestType {
	return &RequestType{ctx: ctx}
}

func (r *RequestType) SetHeader(header http.Header) *RequestType {
	r.header = header
	return r
}

func (r *RequestType) SetNoVersion() *RequestType {
	r.noVersion = true
	return r
}

func (r *RequestType) SetMethod(method string) *RequestType {
	r.method = method
	return r
}

func (r *RequestType) SetURL(url string) *RequestType {
	r.url = url
	return r
}

func (r *RequestType) SetQuery(query string) *RequestType {
	data := url.Values{}
	data.Set("query", query)
	r.query = data.Encode()
	return r
}

func (r *RequestType) SetBody(body io.Reader) *RequestType {
	r.body = body
	return r
}

func (r *RequestType) SetBodyJSON(body interface{}) *RequestType {
	r.bodyJSON = true
	r.vBody = body
	r.additionalHeader = append(r.additionalHeader, "Content-Type", "application/json") //
	return r
}

//Compile the http request
// version selection header spectification
/*
1. if no version selection header is set, the request will be sent to the latest version
routers-version-select
propagated header for selecting version in maching url

*/
func (r *RequestType) Compile() (*http.Request, error) {
	u, err := url.Parse(r.url)
	if err != nil {
		return nil, err
	}

	finalUrl := u.String()
	if r.query != "" {
		finalUrl += "?" + r.query
	}

	req, err := http.NewRequestWithContext(r.ctx, r.method, finalUrl, r.body)
	if err != nil {
		return nil, err
	}
	req.Header = r.header

	// flag the version matching
	// this logic might be moved to infrastructure instead of here
	if _requestVersionMatching && !r.noVersion {
		// set the version selection header
		rvcHeader := ""
		vHeader, vRouters := getRoutingHeader(r.ctx)
		for k, v := range vRouters {
			if k == u.Hostname() {
				req.Header.Set("routers-version-select", vHeader)
			} else {
				rvcHeader += k + "|" + v + ","
			}
		}
		rvcHeader = rvcHeader[:len(rvcHeader)-1]
		req.Header.Set("route-version-select", rvcHeader)
		req.Header.Set("routes-version-select", vHeader)
	}

	return req, nil
}

func getRoutingHeader(ctx context.Context) (string, map[string]string) {
	v := ctx.Value(RoutingContext)
	header, ok := v.(string)

	if !ok || header == "" {
		return "", nil
	}

	routings := strings.Split(header, ",")
	vRouting := make(map[string]string)

	for i, r := range routings {
		selections := strings.Split(r, "|")

		if len(selections) < 2 {
			continue
		}
		if i > 1 {
			i += 2
		}

		vRouting[strings.TrimSpace(selections[0])] = strings.TrimSpace(selections[1])
	}

	return header, vRouting
}

// Get function builds the http request and send it to the target url
func (r *RequestType) Get(url string) *RequestType {
	r.method = http.MethodGet
	r.url = url
	return r
}

func (r *RequestType) Post(url string) *RequestType {
	r.method = http.MethodPost
	r.url = url
	return r
}

func (r *RequestType) Put(url string) *RequestType {
	r.method = http.MethodPut
	r.url = url
	return r
}

func (r *RequestType) Delete(url string) *RequestType {
	r.method = http.MethodDelete
	r.url = url
	return r
}

func (r *RequestType) Patch(url string) *RequestType {
	r.method = http.MethodPatch
	r.url = url
	return r
}

func (r *RequestType) PostForm(kv ...string) *RequestType {
	data := url.Values{}
	for i := range kv {
		if i > 0 {
			i++
			if i == len(kv)-1 {
				break

			}
			data.Add(kv[i], kv[i+1])
		}
	}

	r.body = strings.NewReader(data.Encode())
	r.header.Set("Content-Type", "application/x-www-form-urlencoded")
	return r
}

func (r *RequestType) Query(kv ...string) *RequestType {
	data := url.Values{}
	for i := range kv {
		if i > 0 {
			i++
			if i == len(kv)-1 {
				break

			}
			data.Add(kv[i], kv[i+1])
		}
	}

	r.query = data.Encode()
	return r
}
