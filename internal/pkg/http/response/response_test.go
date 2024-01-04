package response_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"server-go/internal/pkg/http/response"
	"testing"
)

func TestWrite(t *testing.T) {
	cases := []struct {
		Name       string
		Headers    map[string]string
		HTTPStatus int
		XErrors    error
	}{
		{
			Name:       "Test Status",
			HTTPStatus: http.StatusOK,
		},
		{
			Name:    "Test XErrors Kind",
			XErrors: errors.New("bad request"),
		},
		{
			Name:       "Test XErrors Kind with Override HTTP Status",
			HTTPStatus: http.StatusOK,
			XErrors:    errors.New("bad request"),
		},
		{
			Name:       "Test Headers",
			Headers:    map[string]string{"asd": "jkl", "sdf": "hjk"},
			HTTPStatus: http.StatusOK,
		},
	}

	for _, c := range cases {
		t.Logf("test number: %s", c.Name)
		handler := func(w http.ResponseWriter, r *http.Request) {
			jsonresp := response.JSON(w)
			for k, v := range c.Headers {
				jsonresp.SetHeader(k, v)
			}
			if c.XErrors == nil {
				jsonresp.WriteHeader(c.HTTPStatus)
			} else {
				jsonresp.Error(c.XErrors, nil)
			}
			jsonresp.Write()
		}

		req := httptest.NewRequest("GET", "http://example.com", nil)
		w := httptest.NewRecorder()
		handler(w, req)

		resp := w.Result()
		statusCode := c.HTTPStatus
		if c.XErrors != nil {
			fmt.Println(c.XErrors)
			// always expect *xerrors.Errors
			statusCode = http.StatusBadRequest
		}
		// check status code
		if statusCode != resp.StatusCode {
			t.Errorf("invalid http status, expect %d but got %d", c.HTTPStatus, resp.StatusCode)
			return
		}
		// check header
		for key, val := range c.Headers {
			hval := resp.Header.Get(key)
			if hval != val {
				t.Errorf("invalid header value, expect %s but got %s", val, hval)
				return
			}
		}
	}
}
