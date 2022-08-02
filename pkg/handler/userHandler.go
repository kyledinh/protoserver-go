package handler

import (
	"net/http"

	"github.com/kyledinh/protoserver-go/pkg/proto"

	"github.com/google/uuid"

	"go.uber.org/zap"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
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
		// unsupported
	case http.MethodPut:

	default:
		// unsupported method
	}

	logger.Info("UserHandler", zap.String("endpoint", endpoint), zap.String("traceUUID", traceUUID))

}
