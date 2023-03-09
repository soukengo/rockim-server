// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rockim/service/user/v1/online.proto

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

// Validate checks the field values on OnlineAddRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *OnlineAddRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OnlineAddRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OnlineAddRequestMultiError, or nil if none found.
func (m *OnlineAddRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OnlineAddRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetBase() == nil {
		err := OnlineAddRequestValidationError{
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
				errors = append(errors, OnlineAddRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OnlineAddRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBase()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OnlineAddRequestValidationError{
				field:  "Base",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetServerId()) < 1 {
		err := OnlineAddRequestValidationError{
			field:  "ServerId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetChannelId()) < 1 {
		err := OnlineAddRequestValidationError{
			field:  "ChannelId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetToken()) < 1 {
		err := OnlineAddRequestValidationError{
			field:  "Token",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return OnlineAddRequestMultiError(errors)
	}

	return nil
}

// OnlineAddRequestMultiError is an error wrapping multiple validation errors
// returned by OnlineAddRequest.ValidateAll() if the designated constraints
// aren't met.
type OnlineAddRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OnlineAddRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OnlineAddRequestMultiError) AllErrors() []error { return m }

// OnlineAddRequestValidationError is the validation error returned by
// OnlineAddRequest.Validate if the designated constraints aren't met.
type OnlineAddRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OnlineAddRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OnlineAddRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OnlineAddRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OnlineAddRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OnlineAddRequestValidationError) ErrorName() string { return "OnlineAddRequestValidationError" }

// Error satisfies the builtin error interface
func (e OnlineAddRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOnlineAddRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OnlineAddRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OnlineAddRequestValidationError{}

// Validate checks the field values on OnlineAddResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *OnlineAddResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OnlineAddResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OnlineAddResponseMultiError, or nil if none found.
func (m *OnlineAddResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *OnlineAddResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uid

	if len(errors) > 0 {
		return OnlineAddResponseMultiError(errors)
	}

	return nil
}

// OnlineAddResponseMultiError is an error wrapping multiple validation errors
// returned by OnlineAddResponse.ValidateAll() if the designated constraints
// aren't met.
type OnlineAddResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OnlineAddResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OnlineAddResponseMultiError) AllErrors() []error { return m }

// OnlineAddResponseValidationError is the validation error returned by
// OnlineAddResponse.Validate if the designated constraints aren't met.
type OnlineAddResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OnlineAddResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OnlineAddResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OnlineAddResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OnlineAddResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OnlineAddResponseValidationError) ErrorName() string {
	return "OnlineAddResponseValidationError"
}

// Error satisfies the builtin error interface
func (e OnlineAddResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOnlineAddResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OnlineAddResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OnlineAddResponseValidationError{}

// Validate checks the field values on OnlineDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OnlineDeleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OnlineDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OnlineDeleteRequestMultiError, or nil if none found.
func (m *OnlineDeleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OnlineDeleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetBase() == nil {
		err := OnlineDeleteRequestValidationError{
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
				errors = append(errors, OnlineDeleteRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OnlineDeleteRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBase()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OnlineDeleteRequestValidationError{
				field:  "Base",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetServerId()) < 1 {
		err := OnlineDeleteRequestValidationError{
			field:  "ServerId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetChannelId()) < 1 {
		err := OnlineDeleteRequestValidationError{
			field:  "ChannelId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return OnlineDeleteRequestMultiError(errors)
	}

	return nil
}

// OnlineDeleteRequestMultiError is an error wrapping multiple validation
// errors returned by OnlineDeleteRequest.ValidateAll() if the designated
// constraints aren't met.
type OnlineDeleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OnlineDeleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OnlineDeleteRequestMultiError) AllErrors() []error { return m }

// OnlineDeleteRequestValidationError is the validation error returned by
// OnlineDeleteRequest.Validate if the designated constraints aren't met.
type OnlineDeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OnlineDeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OnlineDeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OnlineDeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OnlineDeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OnlineDeleteRequestValidationError) ErrorName() string {
	return "OnlineDeleteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e OnlineDeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOnlineDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OnlineDeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OnlineDeleteRequestValidationError{}

// Validate checks the field values on OnlineDeleteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OnlineDeleteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OnlineDeleteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OnlineDeleteResponseMultiError, or nil if none found.
func (m *OnlineDeleteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *OnlineDeleteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return OnlineDeleteResponseMultiError(errors)
	}

	return nil
}

// OnlineDeleteResponseMultiError is an error wrapping multiple validation
// errors returned by OnlineDeleteResponse.ValidateAll() if the designated
// constraints aren't met.
type OnlineDeleteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OnlineDeleteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OnlineDeleteResponseMultiError) AllErrors() []error { return m }

// OnlineDeleteResponseValidationError is the validation error returned by
// OnlineDeleteResponse.Validate if the designated constraints aren't met.
type OnlineDeleteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OnlineDeleteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OnlineDeleteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OnlineDeleteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OnlineDeleteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OnlineDeleteResponseValidationError) ErrorName() string {
	return "OnlineDeleteResponseValidationError"
}

// Error satisfies the builtin error interface
func (e OnlineDeleteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOnlineDeleteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OnlineDeleteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OnlineDeleteResponseValidationError{}

// Validate checks the field values on OnlineRefreshRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OnlineRefreshRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OnlineRefreshRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OnlineRefreshRequestMultiError, or nil if none found.
func (m *OnlineRefreshRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OnlineRefreshRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetBase() == nil {
		err := OnlineRefreshRequestValidationError{
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
				errors = append(errors, OnlineRefreshRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OnlineRefreshRequestValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBase()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OnlineRefreshRequestValidationError{
				field:  "Base",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetServerId()) < 1 {
		err := OnlineRefreshRequestValidationError{
			field:  "ServerId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetChannelId()) < 1 {
		err := OnlineRefreshRequestValidationError{
			field:  "ChannelId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetUid()) < 1 {
		err := OnlineRefreshRequestValidationError{
			field:  "Uid",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return OnlineRefreshRequestMultiError(errors)
	}

	return nil
}

// OnlineRefreshRequestMultiError is an error wrapping multiple validation
// errors returned by OnlineRefreshRequest.ValidateAll() if the designated
// constraints aren't met.
type OnlineRefreshRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OnlineRefreshRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OnlineRefreshRequestMultiError) AllErrors() []error { return m }

// OnlineRefreshRequestValidationError is the validation error returned by
// OnlineRefreshRequest.Validate if the designated constraints aren't met.
type OnlineRefreshRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OnlineRefreshRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OnlineRefreshRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OnlineRefreshRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OnlineRefreshRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OnlineRefreshRequestValidationError) ErrorName() string {
	return "OnlineRefreshRequestValidationError"
}

// Error satisfies the builtin error interface
func (e OnlineRefreshRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOnlineRefreshRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OnlineRefreshRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OnlineRefreshRequestValidationError{}

// Validate checks the field values on OnlineRefreshResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OnlineRefreshResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OnlineRefreshResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OnlineRefreshResponseMultiError, or nil if none found.
func (m *OnlineRefreshResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *OnlineRefreshResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return OnlineRefreshResponseMultiError(errors)
	}

	return nil
}

// OnlineRefreshResponseMultiError is an error wrapping multiple validation
// errors returned by OnlineRefreshResponse.ValidateAll() if the designated
// constraints aren't met.
type OnlineRefreshResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OnlineRefreshResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OnlineRefreshResponseMultiError) AllErrors() []error { return m }

// OnlineRefreshResponseValidationError is the validation error returned by
// OnlineRefreshResponse.Validate if the designated constraints aren't met.
type OnlineRefreshResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OnlineRefreshResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OnlineRefreshResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OnlineRefreshResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OnlineRefreshResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OnlineRefreshResponseValidationError) ErrorName() string {
	return "OnlineRefreshResponseValidationError"
}

// Error satisfies the builtin error interface
func (e OnlineRefreshResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOnlineRefreshResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OnlineRefreshResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OnlineRefreshResponseValidationError{}