// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rockim/api/openapi/v1/chatroom.proto

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

// Validate checks the field values on ChatRoomCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatRoomCreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatRoomCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatRoomCreateRequestMultiError, or nil if none found.
func (m *ChatRoomCreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatRoomCreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetBase() == nil {
		err := ChatRoomCreateRequestValidationError{
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
				errors = append(errors, ChatRoomCreateRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ChatRoomCreateRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBase()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ChatRoomCreateRequestValidationError{
				field:  "Base",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if !_ChatRoomCreateRequest_CustomGroupId_Pattern.MatchString(m.GetCustomGroupId()) {
		err := ChatRoomCreateRequestValidationError{
			field:  "CustomGroupId",
			reason: "value does not match regex pattern \"^[a-zA-Z0-9_.-]{0,64}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 64 {
		err := ChatRoomCreateRequestValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 64 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetIconUrl()) > 256 {
		err := ChatRoomCreateRequestValidationError{
			field:  "IconUrl",
			reason: "value length must be at most 256 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetFields()) > 100 {
		err := ChatRoomCreateRequestValidationError{
			field:  "Fields",
			reason: "value must contain no more than 100 pair(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	{
		sorted_keys := make([]string, len(m.GetFields()))
		i := 0
		for key := range m.GetFields() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetFields()[key]
			_ = val

			if utf8.RuneCountInString(key) > 64 {
				err := ChatRoomCreateRequestValidationError{
					field:  fmt.Sprintf("Fields[%v]", key),
					reason: "value length must be at most 64 runes",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			if utf8.RuneCountInString(val) > 512 {
				err := ChatRoomCreateRequestValidationError{
					field:  fmt.Sprintf("Fields[%v]", key),
					reason: "value length must be at most 512 runes",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	// no validation rules for OwnerAccount

	if len(errors) > 0 {
		return ChatRoomCreateRequestMultiError(errors)
	}

	return nil
}

// ChatRoomCreateRequestMultiError is an error wrapping multiple validation
// errors returned by ChatRoomCreateRequest.ValidateAll() if the designated
// constraints aren't met.
type ChatRoomCreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatRoomCreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatRoomCreateRequestMultiError) AllErrors() []error { return m }

// ChatRoomCreateRequestValidationError is the validation error returned by
// ChatRoomCreateRequest.Validate if the designated constraints aren't met.
type ChatRoomCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatRoomCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatRoomCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatRoomCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatRoomCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatRoomCreateRequestValidationError) ErrorName() string {
	return "ChatRoomCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChatRoomCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatRoomCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatRoomCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatRoomCreateRequestValidationError{}

var _ChatRoomCreateRequest_CustomGroupId_Pattern = regexp.MustCompile("^[a-zA-Z0-9_.-]{0,64}$")

// Validate checks the field values on ChatRoomCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatRoomCreateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatRoomCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatRoomCreateResponseMultiError, or nil if none found.
func (m *ChatRoomCreateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatRoomCreateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CustomGroupId

	if len(errors) > 0 {
		return ChatRoomCreateResponseMultiError(errors)
	}

	return nil
}

// ChatRoomCreateResponseMultiError is an error wrapping multiple validation
// errors returned by ChatRoomCreateResponse.ValidateAll() if the designated
// constraints aren't met.
type ChatRoomCreateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatRoomCreateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatRoomCreateResponseMultiError) AllErrors() []error { return m }

// ChatRoomCreateResponseValidationError is the validation error returned by
// ChatRoomCreateResponse.Validate if the designated constraints aren't met.
type ChatRoomCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatRoomCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatRoomCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatRoomCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatRoomCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatRoomCreateResponseValidationError) ErrorName() string {
	return "ChatRoomCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ChatRoomCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatRoomCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatRoomCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatRoomCreateResponseValidationError{}
