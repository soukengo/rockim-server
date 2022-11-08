package errors

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
)

type Error struct {
	err   *errors.Error
	stack *stack
}

func New(code int, reason, message string) *Error {
	return &Error{err: errors.New(code, reason, message)}
}

func Newf(code int, reason, format string, a ...interface{}) *Error {
	return New(code, reason, fmt.Sprintf(format, a...))
}

func Code(err error) int {
	return errors.Code(err)
}

func Reason(err error) string {
	return errors.Reason(err)
}

func (e *Error) WithStack() *Error {
	e.stack = callers()
	return e
}

func (e *Error) Error() string {
	return fmt.Sprintf("error: %v stack = %v", e.err, e.stack)
}

func (e *Error) Unwrap() error { return e.err }

func (e *Error) Is(err error) bool {
	if e == nil || e.err == nil {
		return false
	}
	return e.err.Is(err)
}
