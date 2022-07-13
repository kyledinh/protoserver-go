package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_heartbeatHandler(t *testing.T) {
	t.Parallel()

	r := httptest.NewRequest("GET", "/heartbeat", strings.NewReader(""))
	w := httptest.NewRecorder()
	heartbeatHandler(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "{\"status\":\"OK\"}")
}

// func Test_livenessHandler(t *testing.T) {
// 	r := httptest.NewRequest("GET", "/liveness", strings.NewReader(""))
// 	w2 := httptest.NewRecorder()
// 	heartbeatHandler(w2, r)

// 	assert.Equal(t, http.StatusOK, w2.Code)
// 	assert.Contains(t, w2.Body.String(), "ok")
// }
