// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/rockim/admin/manager/v1/types/sys.proto

package types

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

	enums "rockim/api/rockim/shared/enums"
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

	_ = enums.AdminResourceCategory(0)
)

// Validate checks the field values on SysUser with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SysUser) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysUser with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SysUserMultiError, or nil if none found.
func (m *SysUser) ValidateAll() error {
	return m.validate(true)
}

func (m *SysUser) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for CreateTime

	// no validation rules for UpdateTime

	// no validation rules for Account

	// no validation rules for Password

	// no validation rules for Name

	// no validation rules for IsAdmin

	if len(errors) > 0 {
		return SysUserMultiError(errors)
	}

	return nil
}

// SysUserMultiError is an error wrapping multiple validation errors returned
// by SysUser.ValidateAll() if the designated constraints aren't met.
type SysUserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysUserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysUserMultiError) AllErrors() []error { return m }

// SysUserValidationError is the validation error returned by SysUser.Validate
// if the designated constraints aren't met.
type SysUserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysUserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysUserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysUserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysUserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysUserValidationError) ErrorName() string { return "SysUserValidationError" }

// Error satisfies the builtin error interface
func (e SysUserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysUserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysUserValidationError{}

// Validate checks the field values on SysRole with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SysRole) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysRole with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SysRoleMultiError, or nil if none found.
func (m *SysRole) ValidateAll() error {
	return m.validate(true)
}

func (m *SysRole) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for CreateTime

	// no validation rules for UpdateTime

	// no validation rules for Name

	if len(errors) > 0 {
		return SysRoleMultiError(errors)
	}

	return nil
}

// SysRoleMultiError is an error wrapping multiple validation errors returned
// by SysRole.ValidateAll() if the designated constraints aren't met.
type SysRoleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysRoleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysRoleMultiError) AllErrors() []error { return m }

// SysRoleValidationError is the validation error returned by SysRole.Validate
// if the designated constraints aren't met.
type SysRoleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysRoleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysRoleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysRoleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysRoleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysRoleValidationError) ErrorName() string { return "SysRoleValidationError" }

// Error satisfies the builtin error interface
func (e SysRoleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysRole.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysRoleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysRoleValidationError{}

// Validate checks the field values on SysResource with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SysResource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysResource with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SysResourceMultiError, or
// nil if none found.
func (m *SysResource) ValidateAll() error {
	return m.validate(true)
}

func (m *SysResource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for CreateTime

	// no validation rules for UpdateTime

	// no validation rules for Category

	// no validation rules for Name

	// no validation rules for ParentId

	// no validation rules for Url

	// no validation rules for Icon

	// no validation rules for Order

	if len(errors) > 0 {
		return SysResourceMultiError(errors)
	}

	return nil
}

// SysResourceMultiError is an error wrapping multiple validation errors
// returned by SysResource.ValidateAll() if the designated constraints aren't met.
type SysResourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysResourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysResourceMultiError) AllErrors() []error { return m }

// SysResourceValidationError is the validation error returned by
// SysResource.Validate if the designated constraints aren't met.
type SysResourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysResourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysResourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysResourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysResourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysResourceValidationError) ErrorName() string { return "SysResourceValidationError" }

// Error satisfies the builtin error interface
func (e SysResourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysResource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysResourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysResourceValidationError{}

// Validate checks the field values on SysResourceTree with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SysResourceTree) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysResourceTree with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysResourceTreeMultiError, or nil if none found.
func (m *SysResourceTree) ValidateAll() error {
	return m.validate(true)
}

func (m *SysResourceTree) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SysResourceTreeValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SysResourceTreeValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SysResourceTreeValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetChildren() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SysResourceTreeValidationError{
						field:  fmt.Sprintf("Children[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SysResourceTreeValidationError{
						field:  fmt.Sprintf("Children[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SysResourceTreeValidationError{
					field:  fmt.Sprintf("Children[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SysResourceTreeMultiError(errors)
	}

	return nil
}

// SysResourceTreeMultiError is an error wrapping multiple validation errors
// returned by SysResourceTree.ValidateAll() if the designated constraints
// aren't met.
type SysResourceTreeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysResourceTreeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysResourceTreeMultiError) AllErrors() []error { return m }

// SysResourceTreeValidationError is the validation error returned by
// SysResourceTree.Validate if the designated constraints aren't met.
type SysResourceTreeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysResourceTreeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysResourceTreeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysResourceTreeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysResourceTreeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysResourceTreeValidationError) ErrorName() string { return "SysResourceTreeValidationError" }

// Error satisfies the builtin error interface
func (e SysResourceTreeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysResourceTree.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysResourceTreeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysResourceTreeValidationError{}

// Validate checks the field values on SysTenantResource with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SysTenantResource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysTenantResource with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysTenantResourceMultiError, or nil if none found.
func (m *SysTenantResource) ValidateAll() error {
	return m.validate(true)
}

func (m *SysTenantResource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for CreateTime

	// no validation rules for Category

	// no validation rules for Name

	// no validation rules for ParentId

	// no validation rules for Url

	// no validation rules for Icon

	// no validation rules for Order

	if len(errors) > 0 {
		return SysTenantResourceMultiError(errors)
	}

	return nil
}

// SysTenantResourceMultiError is an error wrapping multiple validation errors
// returned by SysTenantResource.ValidateAll() if the designated constraints
// aren't met.
type SysTenantResourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysTenantResourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysTenantResourceMultiError) AllErrors() []error { return m }

// SysTenantResourceValidationError is the validation error returned by
// SysTenantResource.Validate if the designated constraints aren't met.
type SysTenantResourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysTenantResourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysTenantResourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysTenantResourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysTenantResourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysTenantResourceValidationError) ErrorName() string {
	return "SysTenantResourceValidationError"
}

// Error satisfies the builtin error interface
func (e SysTenantResourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysTenantResource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysTenantResourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysTenantResourceValidationError{}

// Validate checks the field values on SysTenantResourceTree with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SysTenantResourceTree) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysTenantResourceTree with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SysTenantResourceTreeMultiError, or nil if none found.
func (m *SysTenantResourceTree) ValidateAll() error {
	return m.validate(true)
}

func (m *SysTenantResourceTree) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SysTenantResourceTreeValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SysTenantResourceTreeValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SysTenantResourceTreeValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetChildren() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SysTenantResourceTreeValidationError{
						field:  fmt.Sprintf("Children[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SysTenantResourceTreeValidationError{
						field:  fmt.Sprintf("Children[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SysTenantResourceTreeValidationError{
					field:  fmt.Sprintf("Children[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SysTenantResourceTreeMultiError(errors)
	}

	return nil
}

// SysTenantResourceTreeMultiError is an error wrapping multiple validation
// errors returned by SysTenantResourceTree.ValidateAll() if the designated
// constraints aren't met.
type SysTenantResourceTreeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysTenantResourceTreeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysTenantResourceTreeMultiError) AllErrors() []error { return m }

// SysTenantResourceTreeValidationError is the validation error returned by
// SysTenantResourceTree.Validate if the designated constraints aren't met.
type SysTenantResourceTreeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysTenantResourceTreeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysTenantResourceTreeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysTenantResourceTreeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysTenantResourceTreeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysTenantResourceTreeValidationError) ErrorName() string {
	return "SysTenantResourceTreeValidationError"
}

// Error satisfies the builtin error interface
func (e SysTenantResourceTreeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysTenantResourceTree.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysTenantResourceTreeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysTenantResourceTreeValidationError{}
