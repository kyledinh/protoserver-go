package api

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/kyledinh/protoserver-go/pkg/config"
	"github.com/kyledinh/protoserver-go/pkg/handler"
	"github.com/kyledinh/protoserver-go/pkg/model"
	"github.com/kyledinh/protoserver-go/pkg/proto"
	"github.com/kyledinh/protoserver-go/pkg/proto/sys"

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
	mux.Handle("/liveness", logWrapper(livenessHandler)) // TODO: add a heathcheck ie. config.Ready()...
	mux.Handle("/version", logWrapper(versionHandler))

	// PROTO X HANDLERS
	mux.Handle("/vx/", logWrapper(handler.VxHandler))

	// LOGIN/SIGNUP
	mux.Handle("/v1/login", logWrapper(handler.LoginHandler))
	mux.Handle("/v1/signup", logWrapper(handler.SignupHandler))

	// ROUTES THAT SECURED BY JWT IN X-Authentication-Header
	mux.Handle("/v1/user", authJWTWrapper(logWrapper(handler.UserHandler)))
	mux.Handle("/v1/users", authJWTWrapper(logWrapper(handler.UsersHandler)))
	mux.Handle("/v1/heartbeat", authJWTWrapper(logWrapper(heartbeatHandler)))

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

func livenessHandler(w http.ResponseWriter, r *http.Request) {
	if config.IsReady() {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("fail"))
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
	logger := proto.Logger(ctx)

	w.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprintln(w, sys.SUCCESS); err != nil {
		logger.Warn("Unable to write response body for heartbeat request", zap.Error(err))
	}
}
