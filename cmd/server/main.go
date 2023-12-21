package server

import (
	"fmt"
	"net/http"

	"github.com/canhcutcon/server-go-postgres/pkg/gee"
)

func main() {
	r := gee.New()

	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello world"))
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":9999")
}
