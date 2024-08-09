package errs

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
)

type errInfo struct {
	e *multierror.Error
}

func Append(err error, errs ...error) *errInfo {
	return &errInfo{
		e: multierror.Append(err, errs...),
	}
}

func (e *errInfo) Error() string {
	return e.e.Error()
}

func (e *errInfo) Format(s fmt.State, v rune) {

	if v == 'v' && s.Flag('+') {
		for _, err := range e.e.WrappedErrors() {
			fmt.Fprintf(s, "%+v\n", err)
		}
		return
	}

	fmt.Fprintf(s, "%v", e.Error())
}

func (e *errInfo) Unwrap() error {
	return e.e.Unwrap()
}
