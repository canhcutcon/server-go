package server

import (
	"context"
	"server-go/internal/pkg/http/misc"
	"server-go/internal/pkg/http/monitoring"
	"server-go/internal/pkg/router"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	requestContext "server-go/internal/pkg/context"
)

type Server struct {
	runners []Runner   // List of runners
	errChan chan error // chan is used to send error from runner to server

	// prometheus metrics is used to collect metrics from runners
	counter           *prometheus.CounterVec
	durationHistogram *prometheus.HistogramVec
	requestSize       *prometheus.HistogramVec
}

// Runner is the interface that wraps the Run method.
type Runner interface {
	Run(middleware ...router.MiddlewareFunc) error
	Shutdown(ctx context.Context) error
}

func (s *Server) Metrics(next router.HandlerFunc) router.HandlerFunc {
	return func(rctx *requestContext.RequestContext) error {
		now := time.Now()
		err := next(rctx)
		httpStatus := 0
		duration := time.Since(now).Seconds()
		requestSize := misc.ComputeApproximateRequestSize(rctx.GetHTTPRequest())
		requestMethod := rctx.GetHTTPRequest().Method
		handlerName := rctx.RequestHandler()
		address := rctx.GetAddress()

		w := rctx.GetHTTPResponseWriter()
		d, ok := w.(monitoring.Delegator)
		if ok {
			httpStatus = d.Status()
		}

		s.counter.WithLabelValues(address, misc.SanitizeCode(httpStatus), requestMethod, handlerName).Inc()
		s.durationHistogram.WithLabelValues(address, misc.SanitizeCode(httpStatus), requestMethod, handlerName).Observe(duration)
		s.requestSize.WithLabelValues(address, misc.SanitizeCode(httpStatus), requestMethod, handlerName).Observe(float64(requestSize))
		return err
	}
}

// Run the server
func (s *Server) Run() chan error {
	for _, r := range s.runners {
		go func(r Runner) {
			if err := r.Run(s.Metrics); err != nil {
				s.errChan <- err
			}
		}(r)
	}
	return s.errChan
}
