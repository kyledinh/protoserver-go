package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kyledinh/protoserver-go/pkg/proto"
	"github.com/spf13/viper"

	"github.com/google/uuid"

	"go.uber.org/zap"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
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
		token, err := loginPost(ctx, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) // find better error response
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"token":"` + token + `"}`))
		}
	case http.MethodPut:
		// unsupported
	default:
		// unsupported method
	}

	logger.Info("LoginHandler", zap.String("endpoint", endpoint), zap.String("traceUUID", traceUUID))

}

func loginPost(ctx context.Context, r *http.Request) (string, error) {
	var err error
	token := "foo-token"
	email := "kyledinh@email.com"
	validLogin := true

	if validLogin {
		token, err = createJWT(email)
	}

	return token, err
}

func createJWT(email string) (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		Issuer:    email,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString([]byte(viper.GetString("jwtSecret")))
}
