package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"protoserver-go/pkg/common"
	"protoserver-go/pkg/common/sys"
	"protoserver-go/pkg/config"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func MockHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := common.Logger(ctx)
	var dirty bool

	url := r.URL.Path
	slice := strings.Split(url, "/")
	ingress := slice[len(slice)-1]

	logger.Info("MockHandler", zap.String("url", ingress))

	for _, route := range config.RouteConfig.Routes {
		if route.Ingress == ingress {
			switch route.Action {
			case "forward":
				dirty = true
			case "get":
				logger.Info(sys.GET, zap.String("egress", route.Egress))
				resp, err := http.Get(route.Egress)
				if err != nil {
					logger.Error(sys.GET, zap.String("egress", route.Egress))
					break
				}
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Println(ingress, " .... failed to Read resp.Body ", resp.Body)
					break
				}
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				w.Write(body)
				dirty = true
			case "respond":
				logger.Info(sys.RESPOND, zap.String("url", route.Ingress))
				str, _ := strconv.Unquote(string(route.Payload))

				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(str))
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
