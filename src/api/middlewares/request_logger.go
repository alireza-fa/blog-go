package middlewares

import (
	"fmt"
	"github.com/alireza-fa/blog-go/src/pkg/logging"
	"net/http"
	"time"
)

func LogMiddleware(next http.Handler) http.Handler {
	logger := logging.NewLogger()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		end := time.Since(start).Milliseconds()

		extra := map[logging.ExtraKey]interface{}{
			logging.ClientIp:  r.RemoteAddr,
			logging.Method:    r.Method,
			logging.Path:      r.URL.Path,
			logging.Timestamp: end,
		}

		logger.Info(logging.RequestResponse, logging.Api, fmt.Sprintf("[%s] %s %dms", r.Method, r.URL.Path, end), extra)
	})
}
