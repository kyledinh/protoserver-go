package api

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"protoserver-go/pkg/common"
	"protoserver-go/pkg/common/sys"
	"protoserver-go/pkg/handler"
	"protoserver-go/pkg/model"

	"log"
	"net/http"
	"time"

	// "github.com/prometheus/client_golang/prometheus/promhttp"

	"go.uber.org/zap"
)

func StartRouter(ctx context.Context, port int) {

	log.Printf("... StartRouter on port %s", strconv.Itoa(port))

	mux := http.NewServeMux()
	mux.Handle("/heartbeat", logWrapper(heartbeatHandler))
	mux.Handle("/health", logWrapper(heartbeatHandler)) // TODO: add a heathcheck ie. config.Ready()...
	mux.Handle("/version", logWrapper(versionHandler))

	// This handler will be deprecated for /vx/ handler
	mux.Handle("/v1/mock/", logWrapper(handler.MockHandler))
	mux.Handle("/vx/", logWrapper(handler.VxHandler))
	mux.Handle("/secure/", authWrapper(logWrapper(handler.VxHandler)))

	portStr := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:           portStr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Panic("Failed to start server", zap.Error(err))
	}

}

func heartbeatHandler(w http.ResponseWriter, r *http.Request) {
	payload := model.RespHeartbeat{Status: "OK"}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := common.Logger(ctx)

	w.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprintln(w, sys.SUCCESS); err != nil {
		logger.Warn("Unable to write response body for heartbeat request", zap.Error(err))
	}
}
