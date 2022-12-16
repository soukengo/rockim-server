// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/rockim/admin/manager/v1/plat_user.proto

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

// Validate checks the field values on PlatUserOptions with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PlatUserOptions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatUserOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatUserOptionsMultiError, or nil if none found.
func (m *PlatUserOptions) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatUserOptions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := PlatUserOptionsValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PlatUserOptionsMultiError(errors)
	}

	return nil
}

// PlatUserOptionsMultiError is an error wrapping multiple validation errors
// returned by PlatUserOptions.ValidateAll() if the designated constraints
// aren't met.
type PlatUserOptionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatUserOptionsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatUserOptionsMultiError) AllErrors() []error { return m }

// PlatUserOptionsValidationError is the validation error returned by
// PlatUserOptions.Validate if the designated constraints aren't met.
type PlatUserOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatUserOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatUserOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatUserOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatUserOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatUserOptionsValidationError) ErrorName() string { return "PlatUserOptionsValidationError" }

// Error satisfies the builtin error interface
func (e PlatUserOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatUserOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatUserOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatUserOptionsValidationError{}

// Validate checks the field values on PlatUserCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatUserCreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatUserCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatUserCreateRequestMultiError, or nil if none found.
func (m *PlatUserCreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatUserCreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetOptions() == nil {
		err := PlatUserCreateRequestValidationError{
			field:  "Options",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PlatUserCreateRequestValidationError{
					field:  "Options",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PlatUserCreateRequestValidationError{
					field:  "Options",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PlatUserCreateRequestValidationError{
				field:  "Options",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PlatUserCreateRequestMultiError(errors)
	}

	return nil
}

// PlatUserCreateRequestMultiError is an error wrapping multiple validation
// errors returned by PlatUserCreateRequest.ValidateAll() if the designated
// constraints aren't met.
type PlatUserCreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatUserCreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatUserCreateRequestMultiError) AllErrors() []error { return m }

// PlatUserCreateRequestValidationError is the validation error returned by
// PlatUserCreateRequest.Validate if the designated constraints aren't met.
type PlatUserCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatUserCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatUserCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatUserCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatUserCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatUserCreateRequestValidationError) ErrorName() string {
	return "PlatUserCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PlatUserCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatUserCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatUserCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatUserCreateRequestValidationError{}

// Validate checks the field values on PlatUserCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatUserCreateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatUserCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatUserCreateResponseMultiError, or nil if none found.
func (m *PlatUserCreateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatUserCreateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return PlatUserCreateResponseMultiError(errors)
	}

	return nil
}

// PlatUserCreateResponseMultiError is an error wrapping multiple validation
// errors returned by PlatUserCreateResponse.ValidateAll() if the designated
// constraints aren't met.
type PlatUserCreateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatUserCreateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatUserCreateResponseMultiError) AllErrors() []error { return m }

// PlatUserCreateResponseValidationError is the validation error returned by
// PlatUserCreateResponse.Validate if the designated constraints aren't met.
type PlatUserCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatUserCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatUserCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatUserCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatUserCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatUserCreateResponseValidationError) ErrorName() string {
	return "PlatUserCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PlatUserCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatUserCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatUserCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatUserCreateResponseValidationError{}

// Validate checks the field values on PlatUserUpdateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatUserUpdateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatUserUpdateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatUserUpdateRequestMultiError, or nil if none found.
func (m *PlatUserUpdateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatUserUpdateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := PlatUserUpdateRequestValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetOptions() == nil {
		err := PlatUserUpdateRequestValidationError{
			field:  "Options",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PlatUserUpdateRequestValidationError{
					field:  "Options",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PlatUserUpdateRequestValidationError{
					field:  "Options",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PlatUserUpdateRequestValidationError{
				field:  "Options",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PlatUserUpdateRequestMultiError(errors)
	}

	return nil
}

// PlatUserUpdateRequestMultiError is an error wrapping multiple validation
// errors returned by PlatUserUpdateRequest.ValidateAll() if the designated
// constraints aren't met.
type PlatUserUpdateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatUserUpdateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatUserUpdateRequestMultiError) AllErrors() []error { return m }

// PlatUserUpdateRequestValidationError is the validation error returned by
// PlatUserUpdateRequest.Validate if the designated constraints aren't met.
type PlatUserUpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatUserUpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatUserUpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatUserUpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatUserUpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatUserUpdateRequestValidationError) ErrorName() string {
	return "PlatUserUpdateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PlatUserUpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatUserUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatUserUpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatUserUpdateRequestValidationError{}

// Validate checks the field values on PlatUserUpdateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatUserUpdateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatUserUpdateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatUserUpdateResponseMultiError, or nil if none found.
func (m *PlatUserUpdateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatUserUpdateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return PlatUserUpdateResponseMultiError(errors)
	}

	return nil
}

// PlatUserUpdateResponseMultiError is an error wrapping multiple validation
// errors returned by PlatUserUpdateResponse.ValidateAll() if the designated
// constraints aren't met.
type PlatUserUpdateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatUserUpdateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatUserUpdateResponseMultiError) AllErrors() []error { return m }

// PlatUserUpdateResponseValidationError is the validation error returned by
// PlatUserUpdateResponse.Validate if the designated constraints aren't met.
type PlatUserUpdateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatUserUpdateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatUserUpdateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatUserUpdateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatUserUpdateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatUserUpdateResponseValidationError) ErrorName() string {
	return "PlatUserUpdateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PlatUserUpdateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatUserUpdateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatUserUpdateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatUserUpdateResponseValidationError{}

// Validate checks the field values on PlatUserDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatUserDeleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatUserDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatUserDeleteRequestMultiError, or nil if none found.
func (m *PlatUserDeleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatUserDeleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := PlatUserDeleteRequestValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PlatUserDeleteRequestMultiError(errors)
	}

	return nil
}

// PlatUserDeleteRequestMultiError is an error wrapping multiple validation
// errors returned by PlatUserDeleteRequest.ValidateAll() if the designated
// constraints aren't met.
type PlatUserDeleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatUserDeleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatUserDeleteRequestMultiError) AllErrors() []error { return m }

// PlatUserDeleteRequestValidationError is the validation error returned by
// PlatUserDeleteRequest.Validate if the designated constraints aren't met.
type PlatUserDeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatUserDeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatUserDeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatUserDeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatUserDeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatUserDeleteRequestValidationError) ErrorName() string {
	return "PlatUserDeleteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PlatUserDeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatUserDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatUserDeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatUserDeleteRequestValidationError{}

// Validate checks the field values on PlatUserDeleteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatUserDeleteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatUserDeleteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatUserDeleteResponseMultiError, or nil if none found.
func (m *PlatUserDeleteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatUserDeleteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return PlatUserDeleteResponseMultiError(errors)
	}

	return nil
}

// PlatUserDeleteResponseMultiError is an error wrapping multiple validation
// errors returned by PlatUserDeleteResponse.ValidateAll() if the designated
// constraints aren't met.
type PlatUserDeleteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatUserDeleteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatUserDeleteResponseMultiError) AllErrors() []error { return m }

// PlatUserDeleteResponseValidationError is the validation error returned by
// PlatUserDeleteResponse.Validate if the designated constraints aren't met.
type PlatUserDeleteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatUserDeleteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatUserDeleteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatUserDeleteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatUserDeleteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatUserDeleteResponseValidationError) ErrorName() string {
	return "PlatUserDeleteResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PlatUserDeleteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatUserDeleteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatUserDeleteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatUserDeleteResponseValidationError{}

// Validate checks the field values on PlatUserPagingRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatUserPagingRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatUserPagingRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatUserPagingRequestMultiError, or nil if none found.
func (m *PlatUserPagingRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatUserPagingRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPaginate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PlatUserPagingRequestValidationError{
					field:  "Paginate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PlatUserPagingRequestValidationError{
					field:  "Paginate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPaginate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PlatUserPagingRequestValidationError{
				field:  "Paginate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Keyword

	if len(errors) > 0 {
		return PlatUserPagingRequestMultiError(errors)
	}

	return nil
}

// PlatUserPagingRequestMultiError is an error wrapping multiple validation
// errors returned by PlatUserPagingRequest.ValidateAll() if the designated
// constraints aren't met.
type PlatUserPagingRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatUserPagingRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatUserPagingRequestMultiError) AllErrors() []error { return m }

// PlatUserPagingRequestValidationError is the validation error returned by
// PlatUserPagingRequest.Validate if the designated constraints aren't met.
type PlatUserPagingRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatUserPagingRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatUserPagingRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatUserPagingRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatUserPagingRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatUserPagingRequestValidationError) ErrorName() string {
	return "PlatUserPagingRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PlatUserPagingRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatUserPagingRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatUserPagingRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatUserPagingRequestValidationError{}

// Validate checks the field values on PlatUserPagingResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatUserPagingResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatUserPagingResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatUserPagingResponseMultiError, or nil if none found.
func (m *PlatUserPagingResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatUserPagingResponse) validate(all bool) error {
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
					errors = append(errors, PlatUserPagingResponseValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PlatUserPagingResponseValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PlatUserPagingResponseValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetPaginate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PlatUserPagingResponseValidationError{
					field:  "Paginate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PlatUserPagingResponseValidationError{
					field:  "Paginate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPaginate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PlatUserPagingResponseValidationError{
				field:  "Paginate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PlatUserPagingResponseMultiError(errors)
	}

	return nil
}

// PlatUserPagingResponseMultiError is an error wrapping multiple validation
// errors returned by PlatUserPagingResponse.ValidateAll() if the designated
// constraints aren't met.
type PlatUserPagingResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatUserPagingResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatUserPagingResponseMultiError) AllErrors() []error { return m }

// PlatUserPagingResponseValidationError is the validation error returned by
// PlatUserPagingResponse.Validate if the designated constraints aren't met.
type PlatUserPagingResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatUserPagingResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatUserPagingResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatUserPagingResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatUserPagingResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatUserPagingResponseValidationError) ErrorName() string {
	return "PlatUserPagingResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PlatUserPagingResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatUserPagingResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatUserPagingResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatUserPagingResponseValidationError{}

// Validate checks the field values on PlatUserRoleIdListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatUserRoleIdListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatUserRoleIdListRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatUserRoleIdListRequestMultiError, or nil if none found.
func (m *PlatUserRoleIdListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatUserRoleIdListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUserId()) < 1 {
		err := PlatUserRoleIdListRequestValidationError{
			field:  "UserId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PlatUserRoleIdListRequestMultiError(errors)
	}

	return nil
}

// PlatUserRoleIdListRequestMultiError is an error wrapping multiple validation
// errors returned by PlatUserRoleIdListRequest.ValidateAll() if the
// designated constraints aren't met.
type PlatUserRoleIdListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatUserRoleIdListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatUserRoleIdListRequestMultiError) AllErrors() []error { return m }

// PlatUserRoleIdListRequestValidationError is the validation error returned by
// PlatUserRoleIdListRequest.Validate if the designated constraints aren't met.
type PlatUserRoleIdListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatUserRoleIdListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatUserRoleIdListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatUserRoleIdListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatUserRoleIdListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatUserRoleIdListRequestValidationError) ErrorName() string {
	return "PlatUserRoleIdListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PlatUserRoleIdListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatUserRoleIdListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatUserRoleIdListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatUserRoleIdListRequestValidationError{}

// Validate checks the field values on PlatUserRoleIdListResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatUserRoleIdListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatUserRoleIdListResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatUserRoleIdListResponseMultiError, or nil if none found.
func (m *PlatUserRoleIdListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatUserRoleIdListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return PlatUserRoleIdListResponseMultiError(errors)
	}

	return nil
}

// PlatUserRoleIdListResponseMultiError is an error wrapping multiple
// validation errors returned by PlatUserRoleIdListResponse.ValidateAll() if
// the designated constraints aren't met.
type PlatUserRoleIdListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatUserRoleIdListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatUserRoleIdListResponseMultiError) AllErrors() []error { return m }

// PlatUserRoleIdListResponseValidationError is the validation error returned
// by PlatUserRoleIdListResponse.Validate if the designated constraints aren't met.
type PlatUserRoleIdListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatUserRoleIdListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatUserRoleIdListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatUserRoleIdListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatUserRoleIdListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatUserRoleIdListResponseValidationError) ErrorName() string {
	return "PlatUserRoleIdListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PlatUserRoleIdListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatUserRoleIdListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatUserRoleIdListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatUserRoleIdListResponseValidationError{}
