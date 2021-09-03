package api

import (
	"net/http"
)

func secWrapper(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, secCheckRequest(r))
	}
}

// TODO: add security checks to block foul requests
func secCheckRequest(r *http.Request) *http.Request {
	ctx := r.Context()
	return r.WithContext(ctx)
}
