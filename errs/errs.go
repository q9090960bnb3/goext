package errs

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
)

type errInfo struct {
	*multierror.Error
}

func Append(err error, errs ...error) *errInfo {
	return &errInfo{
		Error: multierror.Append(err, errs...),
	}
}

func (e *errInfo) Format(s fmt.State, v rune) {

	if v == 'v' && s.Flag('+') {
		for _, err := range e.WrappedErrors() {
			fmt.Fprintf(s, "%+v\n", err)
		}
		return
	}

	fmt.Fprintf(s, "%v", e.Error.Error())
}

func (e *errInfo) Unwrap() error {
	// If we have no errors then we do nothing
	if e == nil || len(e.Errors) == 0 {
		return nil
	}

	// If we have exactly one error, we can just return that directly.
	if len(e.Errors) == 1 {
		return e.Errors[0]
	}

	// Shallow copy the slice
	errs := make([]error, len(e.Errors))
	copy(errs, e.Errors)
	return chain(errs)
}
