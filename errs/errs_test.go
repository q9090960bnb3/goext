package errs

import (
	"errors"
	"fmt"
	"testing"

	"golang.org/x/xerrors"
)

func f1() error {
	return xerrors.New("f1 error")
}

func f2() error {
	return xerrors.New("f2 error")
}

func Test00(t *testing.T) {

	err := Append(f1(), f2())
	fmt.Printf("%#v\n", err)
	fmt.Printf("%+v\n", err)
}

func f3() error {

	err := Append(f1(), f2())
	err.Unwrap()
	return xerrors.Errorf("err: %w", err)
}

func Test01(t *testing.T) {
	err := f3()

	fmt.Printf("%+v\n", err)
}

func Test02(t *testing.T) {

	err := xerrors.Errorf("err: %w", errors.New("wrong"))
	fmt.Printf("%+v\n", err)
}
