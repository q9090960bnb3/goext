package errs

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/go-multierror"
	"golang.org/x/xerrors"
)

func f1() error {
	return xerrors.New("f1 error")
}

func f2() error {
	return xerrors.New("f2 error")
}

var errSpecial = errors.New("special error")

func specialErr() error {
	return xerrors.Errorf("err: %w", errSpecial)
}

func specialErrLast() error {
	return xerrors.Errorf("err: %w", specialErr())
}

func Test00(t *testing.T) {
	errMulti := multierror.Append(f1(), f2(), specialErrLast())
	fmt.Printf("%#v\n", errMulti)

	err := Append(f1(), f2(), specialErrLast())
	fmt.Printf("%#v\n====\n", err)
	fmt.Printf("%+v\n", err)
}

func f3() error {
	err := Append(f1(), f2())
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

var (
	errf4 = errors.New("f4 error")
	errf5 = errors.New("f5 error")
	errf6 = errors.New("f6 error")
)

func f4() error {
	return xerrors.Errorf("err: %w", errf4)
}

func f5() error {
	return errf5
}

func f6() error {
	return fmt.Errorf("%w", errf6)
}

func Test03(t *testing.T) {
	err := Append(f4(), f5(), f6())

	fmt.Println(errors.Is(err, errf4))
	fmt.Println(errors.Is(err, errf5))
	fmt.Println(errors.Is(err, errf6))

	fmt.Printf("%+v", err)
	fmt.Println(err)
}

func f7() error {
	return errors.New("f7 error")
}

func Test04(t *testing.T) {
	errPre := Append(f4(), f5(), f6())

	err := Append(errPre, f7())
	fmt.Println(err)
}

// one err
func Test05(t *testing.T) {
	err := Append(f4())
	fmt.Println(err)
}
