// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rockim/api/client/v1/protocol/http/chatroom_member.proto

package http

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

// Validate checks the field values on ChatRoomJoinRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatRoomJoinRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatRoomJoinRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatRoomJoinRequestMultiError, or nil if none found.
func (m *ChatRoomJoinRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatRoomJoinRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetBase() == nil {
		err := ChatRoomJoinRequestValidationError{
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
				errors = append(errors, ChatRoomJoinRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ChatRoomJoinRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBase()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ChatRoomJoinRequestValidationError{
				field:  "Base",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetCustomGroupId()) < 1 {
		err := ChatRoomJoinRequestValidationError{
			field:  "CustomGroupId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ChatRoomJoinRequestMultiError(errors)
	}

	return nil
}

// ChatRoomJoinRequestMultiError is an error wrapping multiple validation
// errors returned by ChatRoomJoinRequest.ValidateAll() if the designated
// constraints aren't met.
type ChatRoomJoinRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatRoomJoinRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatRoomJoinRequestMultiError) AllErrors() []error { return m }

// ChatRoomJoinRequestValidationError is the validation error returned by
// ChatRoomJoinRequest.Validate if the designated constraints aren't met.
type ChatRoomJoinRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatRoomJoinRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatRoomJoinRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatRoomJoinRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatRoomJoinRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatRoomJoinRequestValidationError) ErrorName() string {
	return "ChatRoomJoinRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChatRoomJoinRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatRoomJoinRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatRoomJoinRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatRoomJoinRequestValidationError{}

// Validate checks the field values on ChatRoomJoinResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatRoomJoinResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatRoomJoinResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatRoomJoinResponseMultiError, or nil if none found.
func (m *ChatRoomJoinResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatRoomJoinResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ChatRoomJoinResponseMultiError(errors)
	}

	return nil
}

// ChatRoomJoinResponseMultiError is an error wrapping multiple validation
// errors returned by ChatRoomJoinResponse.ValidateAll() if the designated
// constraints aren't met.
type ChatRoomJoinResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatRoomJoinResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatRoomJoinResponseMultiError) AllErrors() []error { return m }

// ChatRoomJoinResponseValidationError is the validation error returned by
// ChatRoomJoinResponse.Validate if the designated constraints aren't met.
type ChatRoomJoinResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatRoomJoinResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatRoomJoinResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatRoomJoinResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatRoomJoinResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatRoomJoinResponseValidationError) ErrorName() string {
	return "ChatRoomJoinResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ChatRoomJoinResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatRoomJoinResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatRoomJoinResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatRoomJoinResponseValidationError{}

// Validate checks the field values on ChatRoomQuitRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatRoomQuitRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatRoomQuitRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatRoomQuitRequestMultiError, or nil if none found.
func (m *ChatRoomQuitRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatRoomQuitRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetBase() == nil {
		err := ChatRoomQuitRequestValidationError{
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
				errors = append(errors, ChatRoomQuitRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ChatRoomQuitRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBase()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ChatRoomQuitRequestValidationError{
				field:  "Base",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetCustomGroupId()) < 1 {
		err := ChatRoomQuitRequestValidationError{
			field:  "CustomGroupId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ChatRoomQuitRequestMultiError(errors)
	}

	return nil
}

// ChatRoomQuitRequestMultiError is an error wrapping multiple validation
// errors returned by ChatRoomQuitRequest.ValidateAll() if the designated
// constraints aren't met.
type ChatRoomQuitRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatRoomQuitRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatRoomQuitRequestMultiError) AllErrors() []error { return m }

// ChatRoomQuitRequestValidationError is the validation error returned by
// ChatRoomQuitRequest.Validate if the designated constraints aren't met.
type ChatRoomQuitRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatRoomQuitRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatRoomQuitRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatRoomQuitRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatRoomQuitRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatRoomQuitRequestValidationError) ErrorName() string {
	return "ChatRoomQuitRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChatRoomQuitRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatRoomQuitRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatRoomQuitRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatRoomQuitRequestValidationError{}

// Validate checks the field values on ChatRoomQuitResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatRoomQuitResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatRoomQuitResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatRoomQuitResponseMultiError, or nil if none found.
func (m *ChatRoomQuitResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatRoomQuitResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ChatRoomQuitResponseMultiError(errors)
	}

	return nil
}

// ChatRoomQuitResponseMultiError is an error wrapping multiple validation
// errors returned by ChatRoomQuitResponse.ValidateAll() if the designated
// constraints aren't met.
type ChatRoomQuitResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatRoomQuitResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatRoomQuitResponseMultiError) AllErrors() []error { return m }

// ChatRoomQuitResponseValidationError is the validation error returned by
// ChatRoomQuitResponse.Validate if the designated constraints aren't met.
type ChatRoomQuitResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatRoomQuitResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatRoomQuitResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatRoomQuitResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatRoomQuitResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatRoomQuitResponseValidationError) ErrorName() string {
	return "ChatRoomQuitResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ChatRoomQuitResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatRoomQuitResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatRoomQuitResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatRoomQuitResponseValidationError{}