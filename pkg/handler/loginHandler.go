package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kyledinh/protoserver-go/internal/database"
	"github.com/kyledinh/protoserver-go/pkg/model"
	"github.com/kyledinh/protoserver-go/pkg/proto"
	"github.com/kyledinh/protoserver-go/pkg/proto/protoerr"
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
	var user model.User

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return user.Email, protoerr.ErrParsingRequest
	}

	if err := json.Unmarshal(body, &user); err != nil {
		return user.Email, protoerr.ErrParsingRequest
	}

	validLogin, _ := database.ValidateLogin(user.Email, user.Password)
	if !validLogin {
		return user.Email, protoerr.ErrFailedLogin
	}

	return createJWT(user.Email)
}

func createJWT(email string) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		ID:        email,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString([]byte(viper.GetString("jwtSecret")))
}
