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

	case http.MethodPut:
		// unsupported
	default:
		// unsupported method
	}

	logger.Info("SignupHandler", zap.String("endpoint", endpoint), zap.String("traceUUID", traceUUID))

}
