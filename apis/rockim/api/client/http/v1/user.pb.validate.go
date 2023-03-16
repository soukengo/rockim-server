// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rockim/api/client/http/v1/user.proto

package v1

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

// Validate checks the field values on UserFindRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserFindRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserFindRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserFindRequestMultiError, or nil if none found.
func (m *UserFindRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UserFindRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetBase() == nil {
		err := UserFindRequestValidationError{
			field:  "Base",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetBase()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserFindRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserFindRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBase()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserFindRequestValidationError{
				field:  "Base",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetAccount()) < 1 {
		err := UserFindRequestValidationError{
			field:  "Account",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UserFindRequestMultiError(errors)
	}

	return nil
}

// UserFindRequestMultiError is an error wrapping multiple validation errors
// returned by UserFindRequest.ValidateAll() if the designated constraints
// aren't met.
type UserFindRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserFindRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserFindRequestMultiError) AllErrors() []error { return m }

// UserFindRequestValidationError is the validation error returned by
// UserFindRequest.Validate if the designated constraints aren't met.
type UserFindRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserFindRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserFindRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserFindRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserFindRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserFindRequestValidationError) ErrorName() string { return "UserFindRequestValidationError" }

// Error satisfies the builtin error interface
func (e UserFindRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserFindRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserFindRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserFindRequestValidationError{}

// Validate checks the field values on UserFindResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserFindResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserFindResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserFindResponseMultiError, or nil if none found.
func (m *UserFindResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UserFindResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserFindResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserFindResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserFindResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UserFindResponseMultiError(errors)
	}

	return nil
}

// UserFindResponseMultiError is an error wrapping multiple validation errors
// returned by UserFindResponse.ValidateAll() if the designated constraints
// aren't met.
type UserFindResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserFindResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserFindResponseMultiError) AllErrors() []error { return m }

// UserFindResponseValidationError is the validation error returned by
// UserFindResponse.Validate if the designated constraints aren't met.
type UserFindResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserFindResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserFindResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserFindResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserFindResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserFindResponseValidationError) ErrorName() string { return "UserFindResponseValidationError" }

// Error satisfies the builtin error interface
func (e UserFindResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserFindResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserFindResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserFindResponseValidationError{}