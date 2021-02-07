package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type Resp struct {
	Errno  int                    `json:"errno"`
	ErrMsg string                 `json:"errmsg"`
	Data   map[string]interface{} `json:"data"`
}

var ErrUnknown = errors.New("unknown support")

func main() {
	err := test1()
	fmt.Printf("main call test1 throw err: %#v\n", err)

	is := errors.Is(err, ErrUnknown)
	fmt.Printf("errors.Is(err, ErrUnknown) result is: %v\n", is)

	unWrap := errors.Unwrap(err)
	fmt.Printf("errors.Unwrap(err) result is: %v\n", unWrap)

	tarAs := new(interface{})
	as := errors.As(err, tarAs)
	fmt.Printf("errors.As(err, ErrUnknown) result is: %v tarAs:%#v\n", as, tarAs)

	err = test4()
	fmt.Printf("main call test4 throw err: %#v\n", err)

	ExampleCause()

}

func fn() error {
	e1 := errors.New("error")
	e2 := errors.Wrap(e1, "inner")
	e3 := errors.Wrap(e2, "middle")
	return errors.Wrap(e3, "outer")
}

func ExampleCause() {
	err := fn()
	fmt.Println(err)
	fmt.Println(errors.Cause(err))

	// Output: outer: middle: inner: error
	// error
}

func test1() error {
	return test2()
}

func test2() error {
	return test3()
}

func test3() error {
	return errors.Wrapf(ErrUnknown, "request failure call test3, %v", 121)
}

func test4() error {
	return errors.WithMessage(ErrUnknown, "request failure call test4")
}
