// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rockim/service/group/v1/chatroom.proto

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

	if uri, err := url.Parse(m.GetIconUrl()); err != nil {
		err = ChatRoomCreateRequestValidationError{
			field:  "IconUrl",
			reason: "value must be a valid URI",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	} else if !uri.IsAbs() {
		err := ChatRoomCreateRequestValidationError{
			field:  "IconUrl",
			reason: "value must be absolute",
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

	// no validation rules for GroupId

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

// Validate checks the field values on ChatRoomDismissRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatRoomDismissRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatRoomDismissRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatRoomDismissRequestMultiError, or nil if none found.
func (m *ChatRoomDismissRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatRoomDismissRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetBase() == nil {
		err := ChatRoomDismissRequestValidationError{
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
				errors = append(errors, ChatRoomDismissRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ChatRoomDismissRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBase()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ChatRoomDismissRequestValidationError{
				field:  "Base",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetGroupId()) > 64 {
		err := ChatRoomDismissRequestValidationError{
			field:  "GroupId",
			reason: "value length must be at most 64 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ChatRoomDismissRequestMultiError(errors)
	}

	return nil
}

// ChatRoomDismissRequestMultiError is an error wrapping multiple validation
// errors returned by ChatRoomDismissRequest.ValidateAll() if the designated
// constraints aren't met.
type ChatRoomDismissRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatRoomDismissRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatRoomDismissRequestMultiError) AllErrors() []error { return m }

// ChatRoomDismissRequestValidationError is the validation error returned by
// ChatRoomDismissRequest.Validate if the designated constraints aren't met.
type ChatRoomDismissRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatRoomDismissRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatRoomDismissRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatRoomDismissRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatRoomDismissRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatRoomDismissRequestValidationError) ErrorName() string {
	return "ChatRoomDismissRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChatRoomDismissRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatRoomDismissRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatRoomDismissRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatRoomDismissRequestValidationError{}

// Validate checks the field values on ChatRoomDismissResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatRoomDismissResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatRoomDismissResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatRoomDismissResponseMultiError, or nil if none found.
func (m *ChatRoomDismissResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatRoomDismissResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ChatRoomDismissResponseMultiError(errors)
	}

	return nil
}

// ChatRoomDismissResponseMultiError is an error wrapping multiple validation
// errors returned by ChatRoomDismissResponse.ValidateAll() if the designated
// constraints aren't met.
type ChatRoomDismissResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatRoomDismissResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatRoomDismissResponseMultiError) AllErrors() []error { return m }

// ChatRoomDismissResponseValidationError is the validation error returned by
// ChatRoomDismissResponse.Validate if the designated constraints aren't met.
type ChatRoomDismissResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatRoomDismissResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatRoomDismissResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatRoomDismissResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatRoomDismissResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatRoomDismissResponseValidationError) ErrorName() string {
	return "ChatRoomDismissResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ChatRoomDismissResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatRoomDismissResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatRoomDismissResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatRoomDismissResponseValidationError{}

// Validate checks the field values on ChatRoomFindRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatRoomFindRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatRoomFindRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatRoomFindRequestMultiError, or nil if none found.
func (m *ChatRoomFindRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatRoomFindRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetBase() == nil {
		err := ChatRoomFindRequestValidationError{
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
				errors = append(errors, ChatRoomFindRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ChatRoomFindRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBase()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ChatRoomFindRequestValidationError{
				field:  "Base",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetCustomGroupId()) > 64 {
		err := ChatRoomFindRequestValidationError{
			field:  "CustomGroupId",
			reason: "value length must be at most 64 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ChatRoomFindRequestMultiError(errors)
	}

	return nil
}

// ChatRoomFindRequestMultiError is an error wrapping multiple validation
// errors returned by ChatRoomFindRequest.ValidateAll() if the designated
// constraints aren't met.
type ChatRoomFindRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatRoomFindRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatRoomFindRequestMultiError) AllErrors() []error { return m }

// ChatRoomFindRequestValidationError is the validation error returned by
// ChatRoomFindRequest.Validate if the designated constraints aren't met.
type ChatRoomFindRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatRoomFindRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatRoomFindRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatRoomFindRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatRoomFindRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatRoomFindRequestValidationError) ErrorName() string {
	return "ChatRoomFindRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChatRoomFindRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatRoomFindRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatRoomFindRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatRoomFindRequestValidationError{}

// Validate checks the field values on ChatRoomFindResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatRoomFindResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatRoomFindResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatRoomFindResponseMultiError, or nil if none found.
func (m *ChatRoomFindResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatRoomFindResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetGroup()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ChatRoomFindResponseValidationError{
					field:  "Group",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ChatRoomFindResponseValidationError{
					field:  "Group",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetGroup()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ChatRoomFindResponseValidationError{
				field:  "Group",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ChatRoomFindResponseMultiError(errors)
	}

	return nil
}

// ChatRoomFindResponseMultiError is an error wrapping multiple validation
// errors returned by ChatRoomFindResponse.ValidateAll() if the designated
// constraints aren't met.
type ChatRoomFindResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatRoomFindResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatRoomFindResponseMultiError) AllErrors() []error { return m }

// ChatRoomFindResponseValidationError is the validation error returned by
// ChatRoomFindResponse.Validate if the designated constraints aren't met.
type ChatRoomFindResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatRoomFindResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatRoomFindResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatRoomFindResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatRoomFindResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatRoomFindResponseValidationError) ErrorName() string {
	return "ChatRoomFindResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ChatRoomFindResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatRoomFindResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatRoomFindResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatRoomFindResponseValidationError{}

// Validate checks the field values on ChatRoomFindByIdRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatRoomFindByIdRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatRoomFindByIdRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatRoomFindByIdRequestMultiError, or nil if none found.
func (m *ChatRoomFindByIdRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatRoomFindByIdRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetBase() == nil {
		err := ChatRoomFindByIdRequestValidationError{
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
				errors = append(errors, ChatRoomFindByIdRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ChatRoomFindByIdRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBase()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ChatRoomFindByIdRequestValidationError{
				field:  "Base",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetGroupId()) > 64 {
		err := ChatRoomFindByIdRequestValidationError{
			field:  "GroupId",
			reason: "value length must be at most 64 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ChatRoomFindByIdRequestMultiError(errors)
	}

	return nil
}

// ChatRoomFindByIdRequestMultiError is an error wrapping multiple validation
// errors returned by ChatRoomFindByIdRequest.ValidateAll() if the designated
// constraints aren't met.
type ChatRoomFindByIdRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatRoomFindByIdRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatRoomFindByIdRequestMultiError) AllErrors() []error { return m }

// ChatRoomFindByIdRequestValidationError is the validation error returned by
// ChatRoomFindByIdRequest.Validate if the designated constraints aren't met.
type ChatRoomFindByIdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatRoomFindByIdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatRoomFindByIdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatRoomFindByIdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatRoomFindByIdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatRoomFindByIdRequestValidationError) ErrorName() string {
	return "ChatRoomFindByIdRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChatRoomFindByIdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatRoomFindByIdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatRoomFindByIdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatRoomFindByIdRequestValidationError{}

// Validate checks the field values on ChatRoomFindByIdResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatRoomFindByIdResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatRoomFindByIdResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatRoomFindByIdResponseMultiError, or nil if none found.
func (m *ChatRoomFindByIdResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatRoomFindByIdResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetGroup()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ChatRoomFindByIdResponseValidationError{
					field:  "Group",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ChatRoomFindByIdResponseValidationError{
					field:  "Group",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetGroup()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ChatRoomFindByIdResponseValidationError{
				field:  "Group",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ChatRoomFindByIdResponseMultiError(errors)
	}

	return nil
}

// ChatRoomFindByIdResponseMultiError is an error wrapping multiple validation
// errors returned by ChatRoomFindByIdResponse.ValidateAll() if the designated
// constraints aren't met.
type ChatRoomFindByIdResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatRoomFindByIdResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatRoomFindByIdResponseMultiError) AllErrors() []error { return m }

// ChatRoomFindByIdResponseValidationError is the validation error returned by
// ChatRoomFindByIdResponse.Validate if the designated constraints aren't met.
type ChatRoomFindByIdResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatRoomFindByIdResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatRoomFindByIdResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatRoomFindByIdResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatRoomFindByIdResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatRoomFindByIdResponseValidationError) ErrorName() string {
	return "ChatRoomFindByIdResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ChatRoomFindByIdResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatRoomFindByIdResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatRoomFindByIdResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatRoomFindByIdResponseValidationError{}
