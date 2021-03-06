// +build go1.13

package errors

import (
	stderrors "errors"
	"fmt"
)

// Annotatable represents errors that are unwrappable and formatable.
// Errors that implement special behaviors should implement this interface
// either directly or through embedding in order to avoid breaking Wrap() semantics.
// As a simple example, given the following definitions:
//    type mySpecial struct { errors.Annotatable }
//    func (m mySpecial) Special() bool { return true }
// the following code would add the Special() behavior to a wrapped error
//    specialErr := mySpecial{errors.Wrap(someGenericError, "You are now a special error").(errors.Annotatable)}
type Annotatable interface {
	fmt.Formatter
	Wrapper
}

// Is reports whether any error in err's chain matches target.
//
// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
//
// An error is considered to match a target if it is equal to that target or if
// it implements a method Is(error) bool such that Is(target) returns true.
func Is(err error, target error) bool {
	return stderrors.Is(err, target)
}

// As finds the first error in err's chain that matches target, and if so, sets
// target to that error value and returns true.
//
// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
//
// An error matches target if the error's concrete value is assignable to the value
// pointed to by target, or if the error has a method As(interface{}) bool such that
// As(target) returns true. In the latter case, the As method is responsible for
// setting target.
//
// As will panic if target is not a non-nil pointer to either a type that implements
// error, or to any interface type. As returns false if err is nil.
func As(err error, target interface{}) bool {
	return stderrors.As(err, target)
}

// Unwrap returns the result of calling the Unwrap method on err, if err implements
// Unwrap. Otherwise, Unwrap returns nil.
func Unwrap(err error) error {
	return stderrors.Unwrap(err)
}

// A Wrapper provides context around another error.
type Wrapper interface {
	error
	Unwrap() error
}
