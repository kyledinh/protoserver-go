package api

import (
	"net/http"
)

func validatorWrapper(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, validateRequest(r))
	}
}

// TODO: validate request payload
func validateRequest(r *http.Request) *http.Request {
	ctx := r.Context()
	return r.WithContext(ctx)
}
