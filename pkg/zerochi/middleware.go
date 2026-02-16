package zerochi

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

func Logger(zerologger zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			reqID := middleware.GetReqID(r.Context())

			log := zerologger.With().
				Str("req_id", reqID).
				Str("method", r.Method).
				Str("path", r.URL.RequestURI()).
				Str("remote_addr", r.RemoteAddr).
				Str("user_agent", r.UserAgent()).
				Logger()

			log.Debug().Msg("request started")

			defer func() {
				dur := time.Since(start)

				event := log.Info()
				if ww.Status() >= 400 {
					event = log.Warn()
				}
				if ww.Status() >= 500 {
					event = log.Error()
				}

				event.
					Int("status", ww.Status()).
					Int("bytes", ww.BytesWritten()).
					Dur("dur", dur).
					Msg("request done")
			}()

			next.ServeHTTP(ww, r)
		})
	}
}
