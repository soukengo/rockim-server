// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/rockim/service/platform/v1/plat_resource.proto

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

// Validate checks the field values on PlatResourceOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatResourceOptions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatResourceOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatResourceOptionsMultiError, or nil if none found.
func (m *PlatResourceOptions) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatResourceOptions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Category

	// no validation rules for Name

	// no validation rules for ParentId

	// no validation rules for Url

	// no validation rules for Icon

	// no validation rules for Order

	if len(errors) > 0 {
		return PlatResourceOptionsMultiError(errors)
	}

	return nil
}

// PlatResourceOptionsMultiError is an error wrapping multiple validation
// errors returned by PlatResourceOptions.ValidateAll() if the designated
// constraints aren't met.
type PlatResourceOptionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatResourceOptionsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatResourceOptionsMultiError) AllErrors() []error { return m }

// PlatResourceOptionsValidationError is the validation error returned by
// PlatResourceOptions.Validate if the designated constraints aren't met.
type PlatResourceOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatResourceOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatResourceOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatResourceOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatResourceOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatResourceOptionsValidationError) ErrorName() string {
	return "PlatResourceOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e PlatResourceOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatResourceOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatResourceOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatResourceOptionsValidationError{}

// Validate checks the field values on PlatResourceCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatResourceCreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatResourceCreateRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatResourceCreateRequestMultiError, or nil if none found.
func (m *PlatResourceCreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatResourceCreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetOptions() == nil {
		err := PlatResourceCreateRequestValidationError{
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
				errors = append(errors, PlatResourceCreateRequestValidationError{
					field:  "Options",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PlatResourceCreateRequestValidationError{
					field:  "Options",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PlatResourceCreateRequestValidationError{
				field:  "Options",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PlatResourceCreateRequestMultiError(errors)
	}

	return nil
}

// PlatResourceCreateRequestMultiError is an error wrapping multiple validation
// errors returned by PlatResourceCreateRequest.ValidateAll() if the
// designated constraints aren't met.
type PlatResourceCreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatResourceCreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatResourceCreateRequestMultiError) AllErrors() []error { return m }

// PlatResourceCreateRequestValidationError is the validation error returned by
// PlatResourceCreateRequest.Validate if the designated constraints aren't met.
type PlatResourceCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatResourceCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatResourceCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatResourceCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatResourceCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatResourceCreateRequestValidationError) ErrorName() string {
	return "PlatResourceCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PlatResourceCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatResourceCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatResourceCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatResourceCreateRequestValidationError{}

// Validate checks the field values on PlatResourceCreateResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatResourceCreateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatResourceCreateResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatResourceCreateResponseMultiError, or nil if none found.
func (m *PlatResourceCreateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatResourceCreateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return PlatResourceCreateResponseMultiError(errors)
	}

	return nil
}

// PlatResourceCreateResponseMultiError is an error wrapping multiple
// validation errors returned by PlatResourceCreateResponse.ValidateAll() if
// the designated constraints aren't met.
type PlatResourceCreateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatResourceCreateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatResourceCreateResponseMultiError) AllErrors() []error { return m }

// PlatResourceCreateResponseValidationError is the validation error returned
// by PlatResourceCreateResponse.Validate if the designated constraints aren't met.
type PlatResourceCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatResourceCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatResourceCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatResourceCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatResourceCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatResourceCreateResponseValidationError) ErrorName() string {
	return "PlatResourceCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PlatResourceCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatResourceCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatResourceCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatResourceCreateResponseValidationError{}

// Validate checks the field values on PlatResourceUpdateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatResourceUpdateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatResourceUpdateRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatResourceUpdateRequestMultiError, or nil if none found.
func (m *PlatResourceUpdateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatResourceUpdateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.GetOptions() == nil {
		err := PlatResourceUpdateRequestValidationError{
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
				errors = append(errors, PlatResourceUpdateRequestValidationError{
					field:  "Options",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PlatResourceUpdateRequestValidationError{
					field:  "Options",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PlatResourceUpdateRequestValidationError{
				field:  "Options",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PlatResourceUpdateRequestMultiError(errors)
	}

	return nil
}

// PlatResourceUpdateRequestMultiError is an error wrapping multiple validation
// errors returned by PlatResourceUpdateRequest.ValidateAll() if the
// designated constraints aren't met.
type PlatResourceUpdateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatResourceUpdateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatResourceUpdateRequestMultiError) AllErrors() []error { return m }

// PlatResourceUpdateRequestValidationError is the validation error returned by
// PlatResourceUpdateRequest.Validate if the designated constraints aren't met.
type PlatResourceUpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatResourceUpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatResourceUpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatResourceUpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatResourceUpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatResourceUpdateRequestValidationError) ErrorName() string {
	return "PlatResourceUpdateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PlatResourceUpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatResourceUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatResourceUpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatResourceUpdateRequestValidationError{}

// Validate checks the field values on PlatResourceUpdateResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatResourceUpdateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatResourceUpdateResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatResourceUpdateResponseMultiError, or nil if none found.
func (m *PlatResourceUpdateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatResourceUpdateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return PlatResourceUpdateResponseMultiError(errors)
	}

	return nil
}

// PlatResourceUpdateResponseMultiError is an error wrapping multiple
// validation errors returned by PlatResourceUpdateResponse.ValidateAll() if
// the designated constraints aren't met.
type PlatResourceUpdateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatResourceUpdateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatResourceUpdateResponseMultiError) AllErrors() []error { return m }

// PlatResourceUpdateResponseValidationError is the validation error returned
// by PlatResourceUpdateResponse.Validate if the designated constraints aren't met.
type PlatResourceUpdateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatResourceUpdateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatResourceUpdateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatResourceUpdateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatResourceUpdateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatResourceUpdateResponseValidationError) ErrorName() string {
	return "PlatResourceUpdateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PlatResourceUpdateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatResourceUpdateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatResourceUpdateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatResourceUpdateResponseValidationError{}

// Validate checks the field values on PlatResourceDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatResourceDeleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatResourceDeleteRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatResourceDeleteRequestMultiError, or nil if none found.
func (m *PlatResourceDeleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatResourceDeleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return PlatResourceDeleteRequestMultiError(errors)
	}

	return nil
}

// PlatResourceDeleteRequestMultiError is an error wrapping multiple validation
// errors returned by PlatResourceDeleteRequest.ValidateAll() if the
// designated constraints aren't met.
type PlatResourceDeleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatResourceDeleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatResourceDeleteRequestMultiError) AllErrors() []error { return m }

// PlatResourceDeleteRequestValidationError is the validation error returned by
// PlatResourceDeleteRequest.Validate if the designated constraints aren't met.
type PlatResourceDeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatResourceDeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatResourceDeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatResourceDeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatResourceDeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatResourceDeleteRequestValidationError) ErrorName() string {
	return "PlatResourceDeleteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PlatResourceDeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatResourceDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatResourceDeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatResourceDeleteRequestValidationError{}

// Validate checks the field values on PlatResourceDeleteResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatResourceDeleteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatResourceDeleteResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatResourceDeleteResponseMultiError, or nil if none found.
func (m *PlatResourceDeleteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatResourceDeleteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return PlatResourceDeleteResponseMultiError(errors)
	}

	return nil
}

// PlatResourceDeleteResponseMultiError is an error wrapping multiple
// validation errors returned by PlatResourceDeleteResponse.ValidateAll() if
// the designated constraints aren't met.
type PlatResourceDeleteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatResourceDeleteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatResourceDeleteResponseMultiError) AllErrors() []error { return m }

// PlatResourceDeleteResponseValidationError is the validation error returned
// by PlatResourceDeleteResponse.Validate if the designated constraints aren't met.
type PlatResourceDeleteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatResourceDeleteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatResourceDeleteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatResourceDeleteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatResourceDeleteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatResourceDeleteResponseValidationError) ErrorName() string {
	return "PlatResourceDeleteResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PlatResourceDeleteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatResourceDeleteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatResourceDeleteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatResourceDeleteResponseValidationError{}

// Validate checks the field values on PlatResourceListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatResourceListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatResourceListRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatResourceListRequestMultiError, or nil if none found.
func (m *PlatResourceListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatResourceListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return PlatResourceListRequestMultiError(errors)
	}

	return nil
}

// PlatResourceListRequestMultiError is an error wrapping multiple validation
// errors returned by PlatResourceListRequest.ValidateAll() if the designated
// constraints aren't met.
type PlatResourceListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatResourceListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatResourceListRequestMultiError) AllErrors() []error { return m }

// PlatResourceListRequestValidationError is the validation error returned by
// PlatResourceListRequest.Validate if the designated constraints aren't met.
type PlatResourceListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatResourceListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatResourceListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatResourceListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatResourceListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatResourceListRequestValidationError) ErrorName() string {
	return "PlatResourceListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PlatResourceListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatResourceListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatResourceListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatResourceListRequestValidationError{}

// Validate checks the field values on PlatResourceListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatResourceListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatResourceListResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatResourceListResponseMultiError, or nil if none found.
func (m *PlatResourceListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatResourceListResponse) validate(all bool) error {
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
					errors = append(errors, PlatResourceListResponseValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PlatResourceListResponseValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PlatResourceListResponseValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PlatResourceListResponseMultiError(errors)
	}

	return nil
}

// PlatResourceListResponseMultiError is an error wrapping multiple validation
// errors returned by PlatResourceListResponse.ValidateAll() if the designated
// constraints aren't met.
type PlatResourceListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatResourceListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatResourceListResponseMultiError) AllErrors() []error { return m }

// PlatResourceListResponseValidationError is the validation error returned by
// PlatResourceListResponse.Validate if the designated constraints aren't met.
type PlatResourceListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatResourceListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatResourceListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatResourceListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatResourceListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatResourceListResponseValidationError) ErrorName() string {
	return "PlatResourceListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PlatResourceListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatResourceListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatResourceListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatResourceListResponseValidationError{}

// Validate checks the field values on PlatResourceListByIdsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatResourceListByIdsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatResourceListByIdsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlatResourceListByIdsRequestMultiError, or nil if none found.
func (m *PlatResourceListByIdsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatResourceListByIdsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetIds()) < 1 {
		err := PlatResourceListByIdsRequestValidationError{
			field:  "Ids",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PlatResourceListByIdsRequestMultiError(errors)
	}

	return nil
}

// PlatResourceListByIdsRequestMultiError is an error wrapping multiple
// validation errors returned by PlatResourceListByIdsRequest.ValidateAll() if
// the designated constraints aren't met.
type PlatResourceListByIdsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatResourceListByIdsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatResourceListByIdsRequestMultiError) AllErrors() []error { return m }

// PlatResourceListByIdsRequestValidationError is the validation error returned
// by PlatResourceListByIdsRequest.Validate if the designated constraints
// aren't met.
type PlatResourceListByIdsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatResourceListByIdsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatResourceListByIdsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatResourceListByIdsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatResourceListByIdsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatResourceListByIdsRequestValidationError) ErrorName() string {
	return "PlatResourceListByIdsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PlatResourceListByIdsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatResourceListByIdsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatResourceListByIdsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatResourceListByIdsRequestValidationError{}

// Validate checks the field values on PlatResourceListByIdsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlatResourceListByIdsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatResourceListByIdsResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// PlatResourceListByIdsResponseMultiError, or nil if none found.
func (m *PlatResourceListByIdsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatResourceListByIdsResponse) validate(all bool) error {
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
					errors = append(errors, PlatResourceListByIdsResponseValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PlatResourceListByIdsResponseValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PlatResourceListByIdsResponseValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PlatResourceListByIdsResponseMultiError(errors)
	}

	return nil
}

// PlatResourceListByIdsResponseMultiError is an error wrapping multiple
// validation errors returned by PlatResourceListByIdsResponse.ValidateAll()
// if the designated constraints aren't met.
type PlatResourceListByIdsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatResourceListByIdsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatResourceListByIdsResponseMultiError) AllErrors() []error { return m }

// PlatResourceListByIdsResponseValidationError is the validation error
// returned by PlatResourceListByIdsResponse.Validate if the designated
// constraints aren't met.
type PlatResourceListByIdsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatResourceListByIdsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatResourceListByIdsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatResourceListByIdsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlatResourceListByIdsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatResourceListByIdsResponseValidationError) ErrorName() string {
	return "PlatResourceListByIdsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PlatResourceListByIdsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatResourceListByIdsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatResourceListByIdsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlatResourceListByIdsResponseValidationError{}