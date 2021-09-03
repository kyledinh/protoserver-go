package api

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"protoserver-go/pkg/common"
	"protoserver-go/pkg/common/sys"

	"go.uber.org/zap"
)

func logWrapper(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, requestWithLogContext(r))
	}
}

func requestWithLogContext(r *http.Request) *http.Request {
	ctx := r.Context()
	ref := fmt.Sprintf("api%d_%d", rand.Uint32(), rand.Uint32())
	logID := r.Header.Get("Call-Ref")
	if logID == "" {
		id := make([]byte, 16)
		if _, err := rand.Read(id); err != nil {
			common.Logger(ctx).Error("Unable to generate a call ref", zap.Error(err))
		} else {
			logID = fmt.Sprintf("%s_%s", r.RequestURI, hex.EncodeToString(id))
		}
	}

	logger := common.Logger(ctx).With(zap.String(sys.INTERNALREF, ref), zap.String(sys.CALLREF, logID))
	ctx = context.WithValue(ctx, sys.LOG, logger)
	ctx = context.WithValue(ctx, sys.CALLREF, logID)

	return r.WithContext(ctx)
}
