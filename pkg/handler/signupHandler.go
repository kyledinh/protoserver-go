package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/kyledinh/protoserver-go/internal/database"
	"github.com/kyledinh/protoserver-go/internal/hashing"
	"github.com/kyledinh/protoserver-go/pkg/model"
	"github.com/kyledinh/protoserver-go/pkg/proto"
	"github.com/kyledinh/protoserver-go/pkg/proto/protoerr"

	"github.com/google/uuid"
	"github.com/mrz1836/go-sanitize"

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
			w.WriteHeader(http.StatusBadRequest)
			// find better error response
			// better response messages for err/failure: duplicate | input error | server error
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"created_user":"` + id + `"}`))
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
	var user model.User

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return user.Email, protoerr.ErrParsingRequest
	}

	if err := json.Unmarshal(body, &user); err != nil {
		return user.Email, protoerr.ErrParsingRequest
	}

	user.Email = sanitize.Email(user.Email, false)
	user.Firstname = sanitize.XSS(user.Firstname)
	user.Lastname = sanitize.XSS(user.Lastname)
	user.Password, err = hashing.HashPassword(user.Password)
	if err != nil {
		return user.Email, protoerr.ErrHashPassword
	}

	err = database.InsertNewUser(user)
	if err != nil {
		return user.Email, protoerr.ErrApiRequest
	}

	return user.Email, nil
}
