// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rockim/api/client/v1/socket/options.proto

package socket

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on RequestOptions with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RequestOptions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RequestOptions with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RequestOptionsMultiError,
// or nil if none found.
func (m *RequestOptions) ValidateAll() error {
	return m.validate(true)
}

func (m *RequestOptions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Operation

	if len(errors) > 0 {
		return RequestOptionsMultiError(errors)
	}

	return nil
}

// RequestOptionsMultiError is an error wrapping multiple validation errors
// returned by RequestOptions.ValidateAll() if the designated constraints
// aren't met.
type RequestOptionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RequestOptionsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RequestOptionsMultiError) AllErrors() []error { return m }

// RequestOptionsValidationError is the validation error returned by
// RequestOptions.Validate if the designated constraints aren't met.
type RequestOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestOptionsValidationError) ErrorName() string { return "RequestOptionsValidationError" }

// Error satisfies the builtin error interface
func (e RequestOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequestOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestOptionsValidationError{}