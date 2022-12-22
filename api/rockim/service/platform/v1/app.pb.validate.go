// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/rockim/service/platform/v1/app.proto

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

// Validate checks the field values on AppCreateRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AppCreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AppCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AppCreateRequestMultiError, or nil if none found.
func (m *AppCreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AppCreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TenantId

	// no validation rules for Name

	if len(errors) > 0 {
		return AppCreateRequestMultiError(errors)
	}

	return nil
}

// AppCreateRequestMultiError is an error wrapping multiple validation errors
// returned by AppCreateRequest.ValidateAll() if the designated constraints
// aren't met.
type AppCreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AppCreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AppCreateRequestMultiError) AllErrors() []error { return m }

// AppCreateRequestValidationError is the validation error returned by
// AppCreateRequest.Validate if the designated constraints aren't met.
type AppCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppCreateRequestValidationError) ErrorName() string { return "AppCreateRequestValidationError" }

// Error satisfies the builtin error interface
func (e AppCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppCreateRequestValidationError{}

// Validate checks the field values on AppCreateResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AppCreateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AppCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AppCreateResponseMultiError, or nil if none found.
func (m *AppCreateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AppCreateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return AppCreateResponseMultiError(errors)
	}

	return nil
}

// AppCreateResponseMultiError is an error wrapping multiple validation errors
// returned by AppCreateResponse.ValidateAll() if the designated constraints
// aren't met.
type AppCreateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AppCreateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AppCreateResponseMultiError) AllErrors() []error { return m }

// AppCreateResponseValidationError is the validation error returned by
// AppCreateResponse.Validate if the designated constraints aren't met.
type AppCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppCreateResponseValidationError) ErrorName() string {
	return "AppCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AppCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppCreateResponseValidationError{}

// Validate checks the field values on AppUpdateRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AppUpdateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AppUpdateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AppUpdateRequestMultiError, or nil if none found.
func (m *AppUpdateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AppUpdateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for AppKey

	if len(errors) > 0 {
		return AppUpdateRequestMultiError(errors)
	}

	return nil
}

// AppUpdateRequestMultiError is an error wrapping multiple validation errors
// returned by AppUpdateRequest.ValidateAll() if the designated constraints
// aren't met.
type AppUpdateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AppUpdateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AppUpdateRequestMultiError) AllErrors() []error { return m }

// AppUpdateRequestValidationError is the validation error returned by
// AppUpdateRequest.Validate if the designated constraints aren't met.
type AppUpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppUpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppUpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppUpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppUpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppUpdateRequestValidationError) ErrorName() string { return "AppUpdateRequestValidationError" }

// Error satisfies the builtin error interface
func (e AppUpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppUpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppUpdateRequestValidationError{}

// Validate checks the field values on AppUpdateResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AppUpdateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AppUpdateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AppUpdateResponseMultiError, or nil if none found.
func (m *AppUpdateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AppUpdateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AppUpdateResponseMultiError(errors)
	}

	return nil
}

// AppUpdateResponseMultiError is an error wrapping multiple validation errors
// returned by AppUpdateResponse.ValidateAll() if the designated constraints
// aren't met.
type AppUpdateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AppUpdateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AppUpdateResponseMultiError) AllErrors() []error { return m }

// AppUpdateResponseValidationError is the validation error returned by
// AppUpdateResponse.Validate if the designated constraints aren't met.
type AppUpdateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppUpdateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppUpdateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppUpdateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppUpdateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppUpdateResponseValidationError) ErrorName() string {
	return "AppUpdateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AppUpdateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppUpdateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppUpdateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppUpdateResponseValidationError{}

// Validate checks the field values on AppFindRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AppFindRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AppFindRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AppFindRequestMultiError,
// or nil if none found.
func (m *AppFindRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AppFindRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return AppFindRequestMultiError(errors)
	}

	return nil
}

// AppFindRequestMultiError is an error wrapping multiple validation errors
// returned by AppFindRequest.ValidateAll() if the designated constraints
// aren't met.
type AppFindRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AppFindRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AppFindRequestMultiError) AllErrors() []error { return m }

// AppFindRequestValidationError is the validation error returned by
// AppFindRequest.Validate if the designated constraints aren't met.
type AppFindRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppFindRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppFindRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppFindRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppFindRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppFindRequestValidationError) ErrorName() string { return "AppFindRequestValidationError" }

// Error satisfies the builtin error interface
func (e AppFindRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppFindRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppFindRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppFindRequestValidationError{}

// Validate checks the field values on AppFindResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AppFindResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AppFindResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AppFindResponseMultiError, or nil if none found.
func (m *AppFindResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AppFindResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetApp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AppFindResponseValidationError{
					field:  "App",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AppFindResponseValidationError{
					field:  "App",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetApp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AppFindResponseValidationError{
				field:  "App",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AppFindResponseMultiError(errors)
	}

	return nil
}

// AppFindResponseMultiError is an error wrapping multiple validation errors
// returned by AppFindResponse.ValidateAll() if the designated constraints
// aren't met.
type AppFindResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AppFindResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AppFindResponseMultiError) AllErrors() []error { return m }

// AppFindResponseValidationError is the validation error returned by
// AppFindResponse.Validate if the designated constraints aren't met.
type AppFindResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppFindResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppFindResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppFindResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppFindResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppFindResponseValidationError) ErrorName() string { return "AppFindResponseValidationError" }

// Error satisfies the builtin error interface
func (e AppFindResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppFindResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppFindResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppFindResponseValidationError{}

// Validate checks the field values on AppListByTenantRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AppListByTenantRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AppListByTenantRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AppListByTenantRequestMultiError, or nil if none found.
func (m *AppListByTenantRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AppListByTenantRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTenantId()) < 1 {
		err := AppListByTenantRequestValidationError{
			field:  "TenantId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AppListByTenantRequestMultiError(errors)
	}

	return nil
}

// AppListByTenantRequestMultiError is an error wrapping multiple validation
// errors returned by AppListByTenantRequest.ValidateAll() if the designated
// constraints aren't met.
type AppListByTenantRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AppListByTenantRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AppListByTenantRequestMultiError) AllErrors() []error { return m }

// AppListByTenantRequestValidationError is the validation error returned by
// AppListByTenantRequest.Validate if the designated constraints aren't met.
type AppListByTenantRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppListByTenantRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppListByTenantRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppListByTenantRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppListByTenantRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppListByTenantRequestValidationError) ErrorName() string {
	return "AppListByTenantRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AppListByTenantRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppListByTenantRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppListByTenantRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppListByTenantRequestValidationError{}

// Validate checks the field values on AppListByTenantResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AppListByTenantResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AppListByTenantResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AppListByTenantResponseMultiError, or nil if none found.
func (m *AppListByTenantResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AppListByTenantResponse) validate(all bool) error {
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
					errors = append(errors, AppListByTenantResponseValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AppListByTenantResponseValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AppListByTenantResponseValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return AppListByTenantResponseMultiError(errors)
	}

	return nil
}

// AppListByTenantResponseMultiError is an error wrapping multiple validation
// errors returned by AppListByTenantResponse.ValidateAll() if the designated
// constraints aren't met.
type AppListByTenantResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AppListByTenantResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AppListByTenantResponseMultiError) AllErrors() []error { return m }

// AppListByTenantResponseValidationError is the validation error returned by
// AppListByTenantResponse.Validate if the designated constraints aren't met.
type AppListByTenantResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppListByTenantResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppListByTenantResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppListByTenantResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppListByTenantResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppListByTenantResponseValidationError) ErrorName() string {
	return "AppListByTenantResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AppListByTenantResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppListByTenantResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppListByTenantResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppListByTenantResponseValidationError{}
