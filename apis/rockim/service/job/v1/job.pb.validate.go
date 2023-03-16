// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rockim/service/job/v1/job.proto

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

// Validate checks the field values on MessagePushRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MessagePushRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MessagePushRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MessagePushRequestMultiError, or nil if none found.
func (m *MessagePushRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *MessagePushRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MessagePushRequestValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MessagePushRequestValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MessagePushRequestValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return MessagePushRequestMultiError(errors)
	}

	return nil
}

// MessagePushRequestMultiError is an error wrapping multiple validation errors
// returned by MessagePushRequest.ValidateAll() if the designated constraints
// aren't met.
type MessagePushRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessagePushRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessagePushRequestMultiError) AllErrors() []error { return m }

// MessagePushRequestValidationError is the validation error returned by
// MessagePushRequest.Validate if the designated constraints aren't met.
type MessagePushRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessagePushRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessagePushRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessagePushRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessagePushRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessagePushRequestValidationError) ErrorName() string {
	return "MessagePushRequestValidationError"
}

// Error satisfies the builtin error interface
func (e MessagePushRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessagePushRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessagePushRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessagePushRequestValidationError{}
