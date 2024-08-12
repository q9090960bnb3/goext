package errgroup

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/xerrors"
)

func TestGo(t *testing.T) {
	var g Group

	g.Go(func() error {
		time.Sleep(2 * time.Millisecond)
		fmt.Println("exec 1")
		return nil
	})

	g.Go(func() error {
		time.Sleep(3 * time.Millisecond)
		fmt.Println("exec 2")
		return errors.New("failed to exec 2")
	})

	g.Go(func() error {
		fmt.Println("exec 3")
		return errors.New("failed to exec 3")
	})

	if err := g.Wait(); err == nil {
		fmt.Println("exec done")
	} else {
		fmt.Printf("failed: %+v\n", err)
	}
}

func TestGoLimit(t *testing.T) {
	var g Group

	g.SetLimit(2)

	g.Go(func() error {
		time.Sleep(2 * time.Millisecond)
		fmt.Println("exec 1")
		return nil
	})

	g.Go(func() error {
		time.Sleep(3 * time.Millisecond)
		fmt.Println("exec 2")
		return errors.New("failed to exec 2")
	})

	g.Go(func() error {
		fmt.Println("exec 3")
		return errors.New("failed to exec 3")
	})

	if err := g.Wait(); err == nil {
		fmt.Println("exec done")
	} else {
		fmt.Printf("failed: %+v\n", err)
	}
}

func TestGo_WithStack(t *testing.T) {
	var g Group

	g.Go(func() error {
		time.Sleep(2 * time.Millisecond)
		fmt.Println("exec 1")
		return nil
	})

	g.Go(func() error {
		time.Sleep(3 * time.Millisecond)
		fmt.Println("exec 2")
		return xerrors.New("failed to exec 2")
	})

	g.Go(func() error {
		fmt.Println("exec 3")
		return xerrors.New("failed to exec 3")
	})

	if err := g.Wait(); err == nil {
		fmt.Println("exec done")
	} else {
		fmt.Printf("failed: %+v\n", err)
	}
}

func TestTryGo(t *testing.T) {
	var g Group
	g.SetLimit(5)

	fRun := func(i int) error {
		fmt.Printf("i: %d run", i)
		time.Sleep(1 * time.Millisecond)
		return nil
	}

	for i := 0; i < 10; i++ {
		g.TryGo(func(num int) func() error {
			return func() error {
				return fRun(num)
			}
		}(i))
	}

	if err := g.Wait(); err == nil {
		fmt.Println("exec done")
	} else {
		fmt.Println("failed: ", err)
	}
}

func TestTryGoPanic(t *testing.T) {
	var g Group
	g.SetLimit(-1)
	g.SetLimit(3)

	fRun := func(i int) error {
		fmt.Printf("i: %d run", i)
		time.Sleep(1 * time.Millisecond)
		return nil
	}

	for i := 0; i < 10; i++ {
		g.TryGo(func(num int) func() error {
			return func() error {
				return fRun(num)
			}
		}(i))
	}

	require.Panics(t, func() {
		g.SetLimit(2)
	})
}

func TestCancel(t *testing.T) {

	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Millisecond*5)
	defer cancel1()

	g, ctx2 := WithContext(ctx1)

	g.TryGo(func() error {
		time.Sleep(3 * time.Millisecond)
		return errors.New("finish after 3 second")
	})

	g.Go(func() error {
		time.Sleep(2 * time.Millisecond)
		return errors.New("finish after 2 second")
	})

	if err := g.Wait(); err == nil {
		fmt.Println("exec done")
	} else {
		fmt.Printf("failed: %+v\n", err)
	}

	<-ctx2.Done()
	fmt.Println("ctx2 err:", ctx2.Err())
}
