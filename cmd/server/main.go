package server

import (
	"net/http"

	"github.com/canhcutcon/pkg/gee"
)

func main() {
	r := gee.New()

	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
}
