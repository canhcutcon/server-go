package server

import (
	"server-go/internal/pkg/metrics"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// metrics from server is a
func Prometheus() gin.HandlerFunc { // this is a middleware use to collect metrics from the server
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		c.Next()
		if path == "/metrics" {
			return
		}
		latency := time.Since(start)
		s := c.Writer.Status()
		method := c.Request.Method
		metrics.HttpDuration.WithLabelValues(path, method, strconv.Itoa(s)).
			Observe(latency.Seconds())
	}
}

/*

 */
