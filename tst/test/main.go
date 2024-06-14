package main

import (
	"errors"
	"fmt"
)

type CustomErr string

func (e CustomErr) Error() string {
	return string(e)
}

var (
	CustomErrFail = CustomErr("Fail")
)

func main() {
	var ch CustomErr
	crr := NewCustomErr()
	fmt.Println(
		string(CustomErrFail),
		errors.Is(
			fmt.Errorf("CustomErrFail %w", CustomErrFail),
			CustomErrFail,
		),
		errors.As(
			crr,
			&ch,
		),
	)
}

func NewCustomErr() error {
	return CustomErr("Error")
}
