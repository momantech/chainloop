// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: controlplane/v1/cas_backends.proto

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

// define the regex for a UUID once up-front
var _cas_backends_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CASBackendServiceListRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CASBackendServiceListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CASBackendServiceListRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CASBackendServiceListRequestMultiError, or nil if none found.
func (m *CASBackendServiceListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CASBackendServiceListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CASBackendServiceListRequestMultiError(errors)
	}

	return nil
}

// CASBackendServiceListRequestMultiError is an error wrapping multiple
// validation errors returned by CASBackendServiceListRequest.ValidateAll() if
// the designated constraints aren't met.
type CASBackendServiceListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CASBackendServiceListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CASBackendServiceListRequestMultiError) AllErrors() []error { return m }

// CASBackendServiceListRequestValidationError is the validation error returned
// by CASBackendServiceListRequest.Validate if the designated constraints
// aren't met.
type CASBackendServiceListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CASBackendServiceListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CASBackendServiceListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CASBackendServiceListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CASBackendServiceListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CASBackendServiceListRequestValidationError) ErrorName() string {
	return "CASBackendServiceListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CASBackendServiceListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCASBackendServiceListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CASBackendServiceListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CASBackendServiceListRequestValidationError{}

// Validate checks the field values on CASBackendServiceListResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CASBackendServiceListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CASBackendServiceListResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CASBackendServiceListResponseMultiError, or nil if none found.
func (m *CASBackendServiceListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CASBackendServiceListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResult() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CASBackendServiceListResponseValidationError{
						field:  fmt.Sprintf("Result[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CASBackendServiceListResponseValidationError{
						field:  fmt.Sprintf("Result[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CASBackendServiceListResponseValidationError{
					field:  fmt.Sprintf("Result[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CASBackendServiceListResponseMultiError(errors)
	}

	return nil
}

// CASBackendServiceListResponseMultiError is an error wrapping multiple
// validation errors returned by CASBackendServiceListResponse.ValidateAll()
// if the designated constraints aren't met.
type CASBackendServiceListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CASBackendServiceListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CASBackendServiceListResponseMultiError) AllErrors() []error { return m }

// CASBackendServiceListResponseValidationError is the validation error
// returned by CASBackendServiceListResponse.Validate if the designated
// constraints aren't met.
type CASBackendServiceListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CASBackendServiceListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CASBackendServiceListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CASBackendServiceListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CASBackendServiceListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CASBackendServiceListResponseValidationError) ErrorName() string {
	return "CASBackendServiceListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CASBackendServiceListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCASBackendServiceListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CASBackendServiceListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CASBackendServiceListResponseValidationError{}

// Validate checks the field values on CASBackendServiceCreateRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CASBackendServiceCreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CASBackendServiceCreateRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CASBackendServiceCreateRequestMultiError, or nil if none found.
func (m *CASBackendServiceCreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CASBackendServiceCreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetLocation()) < 1 {
		err := CASBackendServiceCreateRequestValidationError{
			field:  "Location",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetProvider()) < 1 {
		err := CASBackendServiceCreateRequestValidationError{
			field:  "Provider",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Description

	// no validation rules for Default

	if m.GetCredentials() == nil {
		err := CASBackendServiceCreateRequestValidationError{
			field:  "Credentials",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetCredentials()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CASBackendServiceCreateRequestValidationError{
					field:  "Credentials",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CASBackendServiceCreateRequestValidationError{
					field:  "Credentials",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCredentials()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CASBackendServiceCreateRequestValidationError{
				field:  "Credentials",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := CASBackendServiceCreateRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CASBackendServiceCreateRequestMultiError(errors)
	}

	return nil
}

// CASBackendServiceCreateRequestMultiError is an error wrapping multiple
// validation errors returned by CASBackendServiceCreateRequest.ValidateAll()
// if the designated constraints aren't met.
type CASBackendServiceCreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CASBackendServiceCreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CASBackendServiceCreateRequestMultiError) AllErrors() []error { return m }

// CASBackendServiceCreateRequestValidationError is the validation error
// returned by CASBackendServiceCreateRequest.Validate if the designated
// constraints aren't met.
type CASBackendServiceCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CASBackendServiceCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CASBackendServiceCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CASBackendServiceCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CASBackendServiceCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CASBackendServiceCreateRequestValidationError) ErrorName() string {
	return "CASBackendServiceCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CASBackendServiceCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCASBackendServiceCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CASBackendServiceCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CASBackendServiceCreateRequestValidationError{}

// Validate checks the field values on CASBackendServiceCreateResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CASBackendServiceCreateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CASBackendServiceCreateResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CASBackendServiceCreateResponseMultiError, or nil if none found.
func (m *CASBackendServiceCreateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CASBackendServiceCreateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CASBackendServiceCreateResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CASBackendServiceCreateResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CASBackendServiceCreateResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CASBackendServiceCreateResponseMultiError(errors)
	}

	return nil
}

// CASBackendServiceCreateResponseMultiError is an error wrapping multiple
// validation errors returned by CASBackendServiceCreateResponse.ValidateAll()
// if the designated constraints aren't met.
type CASBackendServiceCreateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CASBackendServiceCreateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CASBackendServiceCreateResponseMultiError) AllErrors() []error { return m }

// CASBackendServiceCreateResponseValidationError is the validation error
// returned by CASBackendServiceCreateResponse.Validate if the designated
// constraints aren't met.
type CASBackendServiceCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CASBackendServiceCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CASBackendServiceCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CASBackendServiceCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CASBackendServiceCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CASBackendServiceCreateResponseValidationError) ErrorName() string {
	return "CASBackendServiceCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CASBackendServiceCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCASBackendServiceCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CASBackendServiceCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CASBackendServiceCreateResponseValidationError{}

// Validate checks the field values on CASBackendServiceUpdateRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CASBackendServiceUpdateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CASBackendServiceUpdateRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CASBackendServiceUpdateRequestMultiError, or nil if none found.
func (m *CASBackendServiceUpdateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CASBackendServiceUpdateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = CASBackendServiceUpdateRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for Default

	if all {
		switch v := interface{}(m.GetCredentials()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CASBackendServiceUpdateRequestValidationError{
					field:  "Credentials",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CASBackendServiceUpdateRequestValidationError{
					field:  "Credentials",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCredentials()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CASBackendServiceUpdateRequestValidationError{
				field:  "Credentials",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CASBackendServiceUpdateRequestMultiError(errors)
	}

	return nil
}

func (m *CASBackendServiceUpdateRequest) _validateUuid(uuid string) error {
	if matched := _cas_backends_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// CASBackendServiceUpdateRequestMultiError is an error wrapping multiple
// validation errors returned by CASBackendServiceUpdateRequest.ValidateAll()
// if the designated constraints aren't met.
type CASBackendServiceUpdateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CASBackendServiceUpdateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CASBackendServiceUpdateRequestMultiError) AllErrors() []error { return m }

// CASBackendServiceUpdateRequestValidationError is the validation error
// returned by CASBackendServiceUpdateRequest.Validate if the designated
// constraints aren't met.
type CASBackendServiceUpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CASBackendServiceUpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CASBackendServiceUpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CASBackendServiceUpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CASBackendServiceUpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CASBackendServiceUpdateRequestValidationError) ErrorName() string {
	return "CASBackendServiceUpdateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CASBackendServiceUpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCASBackendServiceUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CASBackendServiceUpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CASBackendServiceUpdateRequestValidationError{}

// Validate checks the field values on CASBackendServiceUpdateResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CASBackendServiceUpdateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CASBackendServiceUpdateResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CASBackendServiceUpdateResponseMultiError, or nil if none found.
func (m *CASBackendServiceUpdateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CASBackendServiceUpdateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CASBackendServiceUpdateResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CASBackendServiceUpdateResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CASBackendServiceUpdateResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CASBackendServiceUpdateResponseMultiError(errors)
	}

	return nil
}

// CASBackendServiceUpdateResponseMultiError is an error wrapping multiple
// validation errors returned by CASBackendServiceUpdateResponse.ValidateAll()
// if the designated constraints aren't met.
type CASBackendServiceUpdateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CASBackendServiceUpdateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CASBackendServiceUpdateResponseMultiError) AllErrors() []error { return m }

// CASBackendServiceUpdateResponseValidationError is the validation error
// returned by CASBackendServiceUpdateResponse.Validate if the designated
// constraints aren't met.
type CASBackendServiceUpdateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CASBackendServiceUpdateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CASBackendServiceUpdateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CASBackendServiceUpdateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CASBackendServiceUpdateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CASBackendServiceUpdateResponseValidationError) ErrorName() string {
	return "CASBackendServiceUpdateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CASBackendServiceUpdateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCASBackendServiceUpdateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CASBackendServiceUpdateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CASBackendServiceUpdateResponseValidationError{}

// Validate checks the field values on CASBackendServiceDeleteRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CASBackendServiceDeleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CASBackendServiceDeleteRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CASBackendServiceDeleteRequestMultiError, or nil if none found.
func (m *CASBackendServiceDeleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CASBackendServiceDeleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = CASBackendServiceDeleteRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CASBackendServiceDeleteRequestMultiError(errors)
	}

	return nil
}

func (m *CASBackendServiceDeleteRequest) _validateUuid(uuid string) error {
	if matched := _cas_backends_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// CASBackendServiceDeleteRequestMultiError is an error wrapping multiple
// validation errors returned by CASBackendServiceDeleteRequest.ValidateAll()
// if the designated constraints aren't met.
type CASBackendServiceDeleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CASBackendServiceDeleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CASBackendServiceDeleteRequestMultiError) AllErrors() []error { return m }

// CASBackendServiceDeleteRequestValidationError is the validation error
// returned by CASBackendServiceDeleteRequest.Validate if the designated
// constraints aren't met.
type CASBackendServiceDeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CASBackendServiceDeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CASBackendServiceDeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CASBackendServiceDeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CASBackendServiceDeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CASBackendServiceDeleteRequestValidationError) ErrorName() string {
	return "CASBackendServiceDeleteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CASBackendServiceDeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCASBackendServiceDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CASBackendServiceDeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CASBackendServiceDeleteRequestValidationError{}

// Validate checks the field values on CASBackendServiceDeleteResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CASBackendServiceDeleteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CASBackendServiceDeleteResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CASBackendServiceDeleteResponseMultiError, or nil if none found.
func (m *CASBackendServiceDeleteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CASBackendServiceDeleteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CASBackendServiceDeleteResponseMultiError(errors)
	}

	return nil
}

// CASBackendServiceDeleteResponseMultiError is an error wrapping multiple
// validation errors returned by CASBackendServiceDeleteResponse.ValidateAll()
// if the designated constraints aren't met.
type CASBackendServiceDeleteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CASBackendServiceDeleteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CASBackendServiceDeleteResponseMultiError) AllErrors() []error { return m }

// CASBackendServiceDeleteResponseValidationError is the validation error
// returned by CASBackendServiceDeleteResponse.Validate if the designated
// constraints aren't met.
type CASBackendServiceDeleteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CASBackendServiceDeleteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CASBackendServiceDeleteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CASBackendServiceDeleteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CASBackendServiceDeleteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CASBackendServiceDeleteResponseValidationError) ErrorName() string {
	return "CASBackendServiceDeleteResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CASBackendServiceDeleteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCASBackendServiceDeleteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CASBackendServiceDeleteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CASBackendServiceDeleteResponseValidationError{}
