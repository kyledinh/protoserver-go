package handler

import (
	"net/http"

	"github.com/kyledinh/protoserver-go/pkg/proto"

	"github.com/google/uuid"

	"go.uber.org/zap"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := proto.Logger(ctx)
	endpoint := r.URL.Path
	traceUUID := r.Header.Get("traceUUID")

	if traceUUID == "" {
		traceUUID = uuid.New().String()
		r.Header.Set("traceUUID", traceUUID)
		logger.Debug("creating a new traceUUID", zap.String("traceUUID", traceUUID))
	}

	switch r.Method {
	case http.MethodGet:
		// unsupported
	case http.MethodPost:
		id, err := SignupPost(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) // find better error response
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"user":"` + id + `"}`))
		}
	case http.MethodPut:
		// unsupported
	default:
		// unsupported method
	}

	logger.Info("SignupHandler", zap.String("endpoint", endpoint), zap.String("traceUUID", traceUUID))

}

func SignupPost(r *http.Request) (string, error) {
	var err error
	var email string

	return email, err
}
