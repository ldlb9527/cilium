// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/load_balancing_policies/wrr_locality/v3/wrr_locality.proto

package wrr_localityv3

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

// Validate checks the field values on WrrLocality with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *WrrLocality) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WrrLocality with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in WrrLocalityMultiError, or
// nil if none found.
func (m *WrrLocality) ValidateAll() error {
	return m.validate(true)
}

func (m *WrrLocality) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetEndpointPickingPolicy() == nil {
		err := WrrLocalityValidationError{
			field:  "EndpointPickingPolicy",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetEndpointPickingPolicy()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WrrLocalityValidationError{
					field:  "EndpointPickingPolicy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WrrLocalityValidationError{
					field:  "EndpointPickingPolicy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEndpointPickingPolicy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WrrLocalityValidationError{
				field:  "EndpointPickingPolicy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return WrrLocalityMultiError(errors)
	}
	return nil
}

// WrrLocalityMultiError is an error wrapping multiple validation errors
// returned by WrrLocality.ValidateAll() if the designated constraints aren't met.
type WrrLocalityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WrrLocalityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WrrLocalityMultiError) AllErrors() []error { return m }

// WrrLocalityValidationError is the validation error returned by
// WrrLocality.Validate if the designated constraints aren't met.
type WrrLocalityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WrrLocalityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WrrLocalityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WrrLocalityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WrrLocalityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WrrLocalityValidationError) ErrorName() string { return "WrrLocalityValidationError" }

// Error satisfies the builtin error interface
func (e WrrLocalityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWrrLocality.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WrrLocalityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WrrLocalityValidationError{}
