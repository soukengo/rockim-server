// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rockim/service/message/v1/message.proto

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

	enums "rockimserver/apis/rockim/shared/enums"
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

	_ = enums.Message_MessageType(0)
)

// Validate checks the field values on MessageSendRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MessageSendRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MessageSendRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MessageSendRequestMultiError, or nil if none found.
func (m *MessageSendRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *MessageSendRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetBase() == nil {
		err := MessageSendRequestValidationError{
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
				errors = append(errors, MessageSendRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MessageSendRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBase()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageSendRequestValidationError{
				field:  "Base",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetUid()) < 1 {
		err := MessageSendRequestValidationError{
			field:  "Uid",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetTarget()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MessageSendRequestValidationError{
					field:  "Target",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MessageSendRequestValidationError{
					field:  "Target",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTarget()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageSendRequestValidationError{
				field:  "Target",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ClientMsgId

	// no validation rules for MessageType

	// no validation rules for Content

	// no validation rules for Payload

	if len(errors) > 0 {
		return MessageSendRequestMultiError(errors)
	}

	return nil
}

// MessageSendRequestMultiError is an error wrapping multiple validation errors
// returned by MessageSendRequest.ValidateAll() if the designated constraints
// aren't met.
type MessageSendRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageSendRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageSendRequestMultiError) AllErrors() []error { return m }

// MessageSendRequestValidationError is the validation error returned by
// MessageSendRequest.Validate if the designated constraints aren't met.
type MessageSendRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageSendRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageSendRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageSendRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageSendRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageSendRequestValidationError) ErrorName() string {
	return "MessageSendRequestValidationError"
}

// Error satisfies the builtin error interface
func (e MessageSendRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessageSendRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageSendRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageSendRequestValidationError{}

// Validate checks the field values on MessageSendResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MessageSendResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MessageSendResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MessageSendResponseMultiError, or nil if none found.
func (m *MessageSendResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *MessageSendResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return MessageSendResponseMultiError(errors)
	}

	return nil
}

// MessageSendResponseMultiError is an error wrapping multiple validation
// errors returned by MessageSendResponse.ValidateAll() if the designated
// constraints aren't met.
type MessageSendResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageSendResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageSendResponseMultiError) AllErrors() []error { return m }

// MessageSendResponseValidationError is the validation error returned by
// MessageSendResponse.Validate if the designated constraints aren't met.
type MessageSendResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageSendResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageSendResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageSendResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageSendResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageSendResponseValidationError) ErrorName() string {
	return "MessageSendResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MessageSendResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessageSendResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageSendResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageSendResponseValidationError{}
