// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rockim/api/client/v1/types/group.proto

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

	_ = enums.Group_MemberRole(0)
)

// Validate checks the field values on Group with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Group) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Group with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in GroupMultiError, or nil if none found.
func (m *Group) ValidateAll() error {
	return m.validate(true)
}

func (m *Group) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Category

	// no validation rules for ProductId

	// no validation rules for CustomGroupId

	// no validation rules for Name

	// no validation rules for IconUrl

	// no validation rules for Fields

	// no validation rules for Status

	if len(errors) > 0 {
		return GroupMultiError(errors)
	}

	return nil
}

// GroupMultiError is an error wrapping multiple validation errors returned by
// Group.ValidateAll() if the designated constraints aren't met.
type GroupMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GroupMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GroupMultiError) AllErrors() []error { return m }

// GroupValidationError is the validation error returned by Group.Validate if
// the designated constraints aren't met.
type GroupValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GroupValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GroupValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GroupValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GroupValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GroupValidationError) ErrorName() string { return "GroupValidationError" }

// Error satisfies the builtin error interface
func (e GroupValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGroup.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GroupValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GroupValidationError{}

// Validate checks the field values on GroupMember with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GroupMember) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GroupMember with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GroupMemberMultiError, or
// nil if none found.
func (m *GroupMember) ValidateAll() error {
	return m.validate(true)
}

func (m *GroupMember) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CreateTime

	// no validation rules for ProductId

	// no validation rules for GroupId

	// no validation rules for Account

	// no validation rules for Role

	// no validation rules for Fields

	if len(errors) > 0 {
		return GroupMemberMultiError(errors)
	}

	return nil
}

// GroupMemberMultiError is an error wrapping multiple validation errors
// returned by GroupMember.ValidateAll() if the designated constraints aren't met.
type GroupMemberMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GroupMemberMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GroupMemberMultiError) AllErrors() []error { return m }

// GroupMemberValidationError is the validation error returned by
// GroupMember.Validate if the designated constraints aren't met.
type GroupMemberValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GroupMemberValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GroupMemberValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GroupMemberValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GroupMemberValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GroupMemberValidationError) ErrorName() string { return "GroupMemberValidationError" }

// Error satisfies the builtin error interface
func (e GroupMemberValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGroupMember.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GroupMemberValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GroupMemberValidationError{}
