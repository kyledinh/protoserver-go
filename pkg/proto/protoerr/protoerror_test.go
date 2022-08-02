package protoerr_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/kyledinh/protoserver-go/pkg/proto/protoerr"

	"github.com/stretchr/testify/assert"
)

func Test_(t *testing.T) {
	t.Parallel()

	type Want struct {
		ErrType        error
		ProtoMessage   string
		WrapperMessage string
	}

	tests := []struct {
		name       string
		newError   error
		protoError error
		want       Want
	}{
		{
			name:       "Error with ErrCLIAction",
			newError:   errors.New("failed to connect to db"),
			protoError: protoerr.ErrCLIAction,
			want: Want{
				ErrType:        protoerr.ErrCLIAction,
				ProtoMessage:   "CLI_ACTION_FAILED",
				WrapperMessage: "wrapped message: failed to connect to db",
			},
		},
		{
			name:       "Error with ErrResourceNotFound",
			newError:   errors.New("File not found"),
			protoError: protoerr.ErrNotFound,
			want: Want{
				ErrType:        protoerr.ErrNotFound,
				ProtoMessage:   "NOT_FOUND",
				WrapperMessage: "wrapped message: File not found",
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			result := protoerr.NewWrappedError(tt.newError.Error(), &tt.protoError)
			assert.Equal(t, tt.want.ProtoMessage, error(*result.Err).Error())
			assert.Equal(t, tt.want.WrapperMessage, result.Error())
		})
	}
}
