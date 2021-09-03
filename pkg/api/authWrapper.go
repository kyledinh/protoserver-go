package api

import (
	"net/http"
)

func authWrapper(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, authorizeRequest(r))
	}
}

func authorizeRequest(r *http.Request) *http.Request {
	ctx := r.Context()
	return r.WithContext(ctx)
}
