package api

import (
	"net/http"
)

func metricsWrapper(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, metricsRequest(r))
	}
}

//TODO: Add Prometheus
func metricsRequest(r *http.Request) *http.Request {
	ctx := r.Context()
	return r.WithContext(ctx)
}
