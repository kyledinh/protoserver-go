package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/kyledinh/protoserver-go/pkg/config"
	"github.com/kyledinh/protoserver-go/pkg/macro"
	"github.com/kyledinh/protoserver-go/pkg/proto"
	"github.com/kyledinh/protoserver-go/pkg/proto/sys"

	"github.com/google/uuid"

	"go.uber.org/zap"
)

func VxHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := proto.Logger(ctx)
	var dirty bool
	endpoint := r.URL.Path
	traceUUID := r.Header.Get("traceUUID")

	if traceUUID == "" {
		traceUUID = uuid.New().String()
		r.Header.Set("traceUUID", traceUUID)
		logger.Debug("creating a new traceUUID", zap.String("traceUUID", traceUUID))
	}

	logger.Info("VxHandler", zap.String("endpoint", endpoint), zap.String("traceUUID", traceUUID))

	for _, route := range config.RouteConfig.Routes {
		if route.Ingress == endpoint {

			switch route.Action {

			case sys.FORWARD:
				logger.Info(sys.FORWARD, zap.String("egress", route.Egress))
				client := http.Client{
					Timeout: time.Duration(5 * time.Second),
				}
				request, err := http.NewRequest(http.MethodPost, route.Egress, r.Body)
				request.Header.Set("Content-type", "application/json")
				request.Header.Set("traceUUID", traceUUID)
				if err != nil {
					//TODO: Handle error
					break
				}
				resp, err := client.Do(request)
				if err != nil {
					//TODO: Handle error
					break
				}
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					//TODO: Handle error
					break
				}

				// Write the Response
				// Check for Std Response Macro
				body = macro.MacroStdResponse(route.Macros, body, r)

				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				w.Write(body)
				dirty = true

			case sys.GET:
				logger.Info(sys.GET, zap.String("egress", route.Egress))
				client := http.Client{
					Timeout: time.Duration(5 * time.Second),
				}

				request, err := http.NewRequest(http.MethodGet, route.Egress, nil)
				request.Header.Set("Content-type", "application/json")
				request.Header.Set("traceUUID", traceUUID)
				if err != nil {
					//TODO: Handle error
					break
				}
				resp, err := client.Do(request)
				if err != nil {
					//TODO: Handle error
					break
				}
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					logger.Error("Failed to read response body", zap.Error(err))
					break
				}

				// Write the Response
				// Check for Std Response Macro
				body = macro.MacroStdResponse(route.Macros, body, r)
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				w.Write(body)
				dirty = true

			case sys.GETPOST:
				logger.Info(sys.GETPOST, zap.String("egress", route.Egress))
				client := http.Client{
					Timeout: time.Duration(5 * time.Second),
				}
				// GET
				request, err := http.NewRequest(http.MethodGet, route.Egresses[0], nil)
				request.Header.Set("Content-type", "application/json")
				request.Header.Set("traceUUID", traceUUID)
				if err != nil {
					//TODO: Handle error
					break
				}
				resp, err := client.Do(request)
				if err != nil {
					//TODO: Handle error
					break
				}
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					logger.Error("Failed to read response body", zap.Error(err))
					break
				}

				// POST AFTER GET
				str, _ := strconv.Unquote(string(body))
				reqPost, err := http.NewRequest(http.MethodPost, route.Egresses[1], strings.NewReader(str))
				reqPost.Header.Set("Content-type", "application/json")
				reqPost.Header.Set("traceUUID", traceUUID)
				if err != nil {
					//TODO: Handle error
					break
				}
				respPost, err := client.Do(reqPost)
				if err != nil {
					//TODO: Handle error
					break
				}
				defer respPost.Body.Close()

				body, err = ioutil.ReadAll(respPost.Body)
				if err != nil {
					logger.Error("Failed to read response body", zap.Error(err))
					break
				}

				// Write the Response
				// Check for Std Response Macro
				body = macro.MacroStdResponse(route.Macros, body, r)
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				w.Write(body)
				dirty = true

			case sys.POST:
				logger.Info(sys.POST, zap.String("egress", route.Egress))
				str, _ := strconv.Unquote(string(route.Payload))
				client := http.Client{
					Timeout: time.Duration(5 * time.Second),
				}
				req, err := http.NewRequest(http.MethodPost, route.Egress, strings.NewReader(str))
				req.Header.Set("Content-type", "application/json")
				req.Header.Set("traceUUID", traceUUID)
				if err != nil {
					//TODO: Handle error
					break
				}
				resp, err := client.Do(req)
				if err != nil {
					//TODO: Handle error
					break
				}
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					//TODO: Handle error
					break
				}

				// Write the Response
				// Check for Std Response Macro
				body = macro.MacroStdResponse(route.Macros, body, r)

				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				w.Write(body)
				dirty = true

			case sys.RESPOND:
				logger.Info(sys.RESPOND, zap.String("ingress", route.Ingress))
				body := route.Payload

				// Write the Response
				// Check for Std Response Macro
				body = macro.MacroStdResponse(route.Macros, body, r)
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				w.Write(body)
				dirty = true
			}
		}
	}

	// If there wasn't a matching ingress, then send default response
	if !dirty {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := fmt.Fprintln(w, sys.FAILURE); err != nil {
			logger.Warn("Unable to write response body for mock request", zap.Error(err))
		}
	}
}
