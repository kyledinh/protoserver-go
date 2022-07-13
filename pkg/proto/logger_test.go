package proto_test

import (
	"bytes"
	"context"
	"protoserver-go/pkg/proto"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestLogger(t *testing.T) {

	for _, tt := range []struct {
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
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			buf := bytes.Buffer{}
			viper.Set("log.type", "debug")
			viper.Set("log.buffer", &buf)
			viper.Set("log.level", "debug")

			proto.SetupLogger(ctx, tt.servicename)
			logger, err := proto.GetLogger(tt.servicename)
			require.NoError(t, err)

			logger.Info(tt.msg, zap.String(tt.key, tt.value))
			assert.Contains(t, buf.String(), tt.msg)
			assert.Contains(t, buf.String(), string(tt.exp_in_payload))
		})
	}
}
