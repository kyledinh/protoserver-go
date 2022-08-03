package protoerr

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrResourceNotFound = errors.New("RESOURCE_NOT_FOUND")
	ErrResourceRead     = errors.New("RESOURCE_NOT_READ")
	ErrConversionFormat = errors.New("CONVERSION_FORMAT_FAILED")
	ErrCLIAction        = errors.New("CLI_ACTION_FAILED")
	ErrNotFound         = errors.New("NOT_FOUND")
	ErrWriteFile        = errors.New("WRITE_FILE_FAILED")
	ErrHashPassword     = errors.New("HASH_PASSWD_FAILED")
	ErrParsingRequest   = errors.New("PARSING_REQUEST_FAILED")
	ErrApiRequest       = errors.New("API_REQUEST_FAILED")
)

type WrappedError struct {
	Message string
	Err     *error
}

func (w *WrappedError) Error() string {
	return fmt.Sprintf("wrapped message: %s", w.Message)
}

func NewWrappedError(message string, err *error) WrappedError {
	wrappedError := WrappedError{
		Message: message,
		Err:     err,
	}
	return wrappedError
}
