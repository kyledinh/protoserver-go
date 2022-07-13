package protoerr

import (
	"fmt"

	"github.com/pkg/errors"
)

type WrappedError struct {
	WrapType error
	Err      error
}

func (w *WrappedError) Error() string {
	return fmt.Sprintf("%s: %v", w.WrapType.Error(), w.Err)
}

func Wrap(err error, wrapType error) *WrappedError {
	return &WrappedError{
		WrapType: wrapType,
		Err:      err,
	}
}

var (
	ErrNotFound = errors.New("NOT_FOUND")
)
