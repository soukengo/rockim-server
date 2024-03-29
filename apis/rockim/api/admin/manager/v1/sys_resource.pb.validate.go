// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rockim/api/admin/manager/v1/sys_resource.proto

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

	_ = enums.Admin_ResourceCategory(0)
)

// Validate checks the field values on SysResourceOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysResourceOptions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysResourceOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysResourceOptionsMultiError, or nil if none found.
func (m *SysResourceOptions) ValidateAll() error {
	return m.validate(true)
}

func (m *SysResourceOptions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := enums.Admin_ResourceCategory_name[int32(m.GetCategory())]; !ok {
		err := SysResourceOptionsValidationError{
			field:  "Category",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := SysResourceOptionsValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetParentId()) < 1 {
		err := SysResourceOptionsValidationError{
			field:  "ParentId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Url

	// no validation rules for Icon

	// no validation rules for Order

	if len(errors) > 0 {
		return SysResourceOptionsMultiError(errors)
	}

	return nil
}

// SysResourceOptionsMultiError is an error wrapping multiple validation errors
// returned by SysResourceOptions.ValidateAll() if the designated constraints
// aren't met.
type SysResourceOptionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysResourceOptionsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysResourceOptionsMultiError) AllErrors() []error { return m }

// SysResourceOptionsValidationError is the validation error returned by
// SysResourceOptions.Validate if the designated constraints aren't met.
type SysResourceOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysResourceOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysResourceOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysResourceOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysResourceOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysResourceOptionsValidationError) ErrorName() string {
	return "SysResourceOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e SysResourceOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysResourceOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysResourceOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysResourceOptionsValidationError{}

// Validate checks the field values on SysResourceCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysResourceCreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysResourceCreateRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysResourceCreateRequestMultiError, or nil if none found.
func (m *SysResourceCreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SysResourceCreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetOptions() == nil {
		err := SysResourceCreateRequestValidationError{
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
				errors = append(errors, SysResourceCreateRequestValidationError{
					field:  "Options",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SysResourceCreateRequestValidationError{
					field:  "Options",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SysResourceCreateRequestValidationError{
				field:  "Options",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SysResourceCreateRequestMultiError(errors)
	}

	return nil
}

// SysResourceCreateRequestMultiError is an error wrapping multiple validation
// errors returned by SysResourceCreateRequest.ValidateAll() if the designated
// constraints aren't met.
type SysResourceCreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysResourceCreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysResourceCreateRequestMultiError) AllErrors() []error { return m }

// SysResourceCreateRequestValidationError is the validation error returned by
// SysResourceCreateRequest.Validate if the designated constraints aren't met.
type SysResourceCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysResourceCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysResourceCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysResourceCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysResourceCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysResourceCreateRequestValidationError) ErrorName() string {
	return "SysResourceCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SysResourceCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysResourceCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysResourceCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysResourceCreateRequestValidationError{}

// Validate checks the field values on SysResourceCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysResourceCreateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysResourceCreateResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysResourceCreateResponseMultiError, or nil if none found.
func (m *SysResourceCreateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SysResourceCreateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SysResourceCreateResponseMultiError(errors)
	}

	return nil
}

// SysResourceCreateResponseMultiError is an error wrapping multiple validation
// errors returned by SysResourceCreateResponse.ValidateAll() if the
// designated constraints aren't met.
type SysResourceCreateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysResourceCreateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysResourceCreateResponseMultiError) AllErrors() []error { return m }

// SysResourceCreateResponseValidationError is the validation error returned by
// SysResourceCreateResponse.Validate if the designated constraints aren't met.
type SysResourceCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysResourceCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysResourceCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysResourceCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysResourceCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysResourceCreateResponseValidationError) ErrorName() string {
	return "SysResourceCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SysResourceCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysResourceCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysResourceCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysResourceCreateResponseValidationError{}

// Validate checks the field values on SysResourceUpdateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysResourceUpdateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysResourceUpdateRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysResourceUpdateRequestMultiError, or nil if none found.
func (m *SysResourceUpdateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SysResourceUpdateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := SysResourceUpdateRequestValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetOptions() == nil {
		err := SysResourceUpdateRequestValidationError{
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
				errors = append(errors, SysResourceUpdateRequestValidationError{
					field:  "Options",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SysResourceUpdateRequestValidationError{
					field:  "Options",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SysResourceUpdateRequestValidationError{
				field:  "Options",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SysResourceUpdateRequestMultiError(errors)
	}

	return nil
}

// SysResourceUpdateRequestMultiError is an error wrapping multiple validation
// errors returned by SysResourceUpdateRequest.ValidateAll() if the designated
// constraints aren't met.
type SysResourceUpdateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysResourceUpdateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysResourceUpdateRequestMultiError) AllErrors() []error { return m }

// SysResourceUpdateRequestValidationError is the validation error returned by
// SysResourceUpdateRequest.Validate if the designated constraints aren't met.
type SysResourceUpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysResourceUpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysResourceUpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysResourceUpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysResourceUpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysResourceUpdateRequestValidationError) ErrorName() string {
	return "SysResourceUpdateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SysResourceUpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysResourceUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysResourceUpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysResourceUpdateRequestValidationError{}

// Validate checks the field values on SysResourceUpdateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysResourceUpdateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysResourceUpdateResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysResourceUpdateResponseMultiError, or nil if none found.
func (m *SysResourceUpdateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SysResourceUpdateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SysResourceUpdateResponseMultiError(errors)
	}

	return nil
}

// SysResourceUpdateResponseMultiError is an error wrapping multiple validation
// errors returned by SysResourceUpdateResponse.ValidateAll() if the
// designated constraints aren't met.
type SysResourceUpdateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysResourceUpdateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysResourceUpdateResponseMultiError) AllErrors() []error { return m }

// SysResourceUpdateResponseValidationError is the validation error returned by
// SysResourceUpdateResponse.Validate if the designated constraints aren't met.
type SysResourceUpdateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysResourceUpdateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysResourceUpdateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysResourceUpdateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysResourceUpdateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysResourceUpdateResponseValidationError) ErrorName() string {
	return "SysResourceUpdateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SysResourceUpdateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysResourceUpdateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysResourceUpdateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysResourceUpdateResponseValidationError{}

// Validate checks the field values on SysResourceDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysResourceDeleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysResourceDeleteRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysResourceDeleteRequestMultiError, or nil if none found.
func (m *SysResourceDeleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SysResourceDeleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := SysResourceDeleteRequestValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SysResourceDeleteRequestMultiError(errors)
	}

	return nil
}

// SysResourceDeleteRequestMultiError is an error wrapping multiple validation
// errors returned by SysResourceDeleteRequest.ValidateAll() if the designated
// constraints aren't met.
type SysResourceDeleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysResourceDeleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysResourceDeleteRequestMultiError) AllErrors() []error { return m }

// SysResourceDeleteRequestValidationError is the validation error returned by
// SysResourceDeleteRequest.Validate if the designated constraints aren't met.
type SysResourceDeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysResourceDeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysResourceDeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysResourceDeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysResourceDeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysResourceDeleteRequestValidationError) ErrorName() string {
	return "SysResourceDeleteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SysResourceDeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysResourceDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysResourceDeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysResourceDeleteRequestValidationError{}

// Validate checks the field values on SysResourceDeleteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysResourceDeleteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysResourceDeleteResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysResourceDeleteResponseMultiError, or nil if none found.
func (m *SysResourceDeleteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SysResourceDeleteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SysResourceDeleteResponseMultiError(errors)
	}

	return nil
}

// SysResourceDeleteResponseMultiError is an error wrapping multiple validation
// errors returned by SysResourceDeleteResponse.ValidateAll() if the
// designated constraints aren't met.
type SysResourceDeleteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysResourceDeleteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysResourceDeleteResponseMultiError) AllErrors() []error { return m }

// SysResourceDeleteResponseValidationError is the validation error returned by
// SysResourceDeleteResponse.Validate if the designated constraints aren't met.
type SysResourceDeleteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysResourceDeleteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysResourceDeleteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysResourceDeleteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysResourceDeleteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysResourceDeleteResponseValidationError) ErrorName() string {
	return "SysResourceDeleteResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SysResourceDeleteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysResourceDeleteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysResourceDeleteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysResourceDeleteResponseValidationError{}

// Validate checks the field values on SysResourceTreeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysResourceTreeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysResourceTreeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysResourceTreeRequestMultiError, or nil if none found.
func (m *SysResourceTreeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SysResourceTreeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SysResourceTreeRequestMultiError(errors)
	}

	return nil
}

// SysResourceTreeRequestMultiError is an error wrapping multiple validation
// errors returned by SysResourceTreeRequest.ValidateAll() if the designated
// constraints aren't met.
type SysResourceTreeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysResourceTreeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysResourceTreeRequestMultiError) AllErrors() []error { return m }

// SysResourceTreeRequestValidationError is the validation error returned by
// SysResourceTreeRequest.Validate if the designated constraints aren't met.
type SysResourceTreeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysResourceTreeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysResourceTreeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysResourceTreeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysResourceTreeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysResourceTreeRequestValidationError) ErrorName() string {
	return "SysResourceTreeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SysResourceTreeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysResourceTreeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysResourceTreeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysResourceTreeRequestValidationError{}

// Validate checks the field values on SysResourceTreeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysResourceTreeResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysResourceTreeResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysResourceTreeResponseMultiError, or nil if none found.
func (m *SysResourceTreeResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SysResourceTreeResponse) validate(all bool) error {
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
					errors = append(errors, SysResourceTreeResponseValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SysResourceTreeResponseValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SysResourceTreeResponseValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SysResourceTreeResponseMultiError(errors)
	}

	return nil
}

// SysResourceTreeResponseMultiError is an error wrapping multiple validation
// errors returned by SysResourceTreeResponse.ValidateAll() if the designated
// constraints aren't met.
type SysResourceTreeResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysResourceTreeResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysResourceTreeResponseMultiError) AllErrors() []error { return m }

// SysResourceTreeResponseValidationError is the validation error returned by
// SysResourceTreeResponse.Validate if the designated constraints aren't met.
type SysResourceTreeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysResourceTreeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysResourceTreeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysResourceTreeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysResourceTreeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysResourceTreeResponseValidationError) ErrorName() string {
	return "SysResourceTreeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SysResourceTreeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysResourceTreeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysResourceTreeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysResourceTreeResponseValidationError{}
