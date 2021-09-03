package common

import (
	"bytes"
	"context"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestLogger(t *testing.T) {

	for _, test := range []struct {
		name           string
		servicename    string
		msg            string
		key            string
		value          string
		exp_in_payload []byte
	}{
		{"testing the serviceName input", "alpha-server", "alpha message log", "userid", "Alpha", []byte(`"userid":"Alpha"`)},
		{"testing the serviceName input", "beta-server", "logging beta", "userid", "Beta", []byte(`"userid":"Beta"`)},
		{"testing the serviceName input", "beta-server", "logging beta", "traceUUID", "ABCDEFG93", []byte(`"traceUUID":"ABCDEFG93"`)},
	} {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			buf := bytes.Buffer{}
			viper.Set("log.type", "debug")
			viper.Set("log.buffer", &buf)
			viper.Set("log.level", "debug")

			SetupLogger(ctx, test.servicename)
			logger, err := GetLogger(test.servicename)
			require.NoError(t, err)

			logger.Info(test.msg, zap.String(test.key, test.value))
			assert.Contains(t, buf.String(), test.msg)
			assert.Contains(t, buf.String(), string(test.exp_in_payload))
		})
	}
}
