package errs

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
)

type errInfo struct {
	e *multierror.Error
}

func Append(err error, errs ...error) *errInfo {
	res := &errInfo{
		e: multierror.Append(err, errs...),
	}
	res.e.ErrorFormat = ListFormatFunc
	return res
}

func (e *errInfo) Error() string {
	return e.e.Error()
}

func (e *errInfo) Format(s fmt.State, v rune) {

	if v == 'v' && s.Flag('+') {
		wrappedErrs := e.e.WrappedErrors()
		lenErr := len(wrappedErrs)
		for i, err := range wrappedErrs {
			if i != lenErr-1 {
				fmt.Fprintf(s, "%+v\n", err)
			} else {
				fmt.Fprintf(s, "%+v", err)
			}
		}
		return
	}

	fmt.Fprintf(s, "%v", e.Error())
}

func (e *errInfo) Unwrap() error {
	return e.e.Unwrap()
}
