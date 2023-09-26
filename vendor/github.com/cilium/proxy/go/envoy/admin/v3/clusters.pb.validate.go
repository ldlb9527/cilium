// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/admin/v3/clusters.proto

package adminv3

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

	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
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

	_ = v3.HealthStatus(0)
)

// Validate checks the field values on Clusters with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Clusters) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Clusters with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ClustersMultiError, or nil
// if none found.
func (m *Clusters) ValidateAll() error {
	return m.validate(true)
}

func (m *Clusters) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetClusterStatuses() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ClustersValidationError{
						field:  fmt.Sprintf("ClusterStatuses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ClustersValidationError{
						field:  fmt.Sprintf("ClusterStatuses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClustersValidationError{
					field:  fmt.Sprintf("ClusterStatuses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ClustersMultiError(errors)
	}
	return nil
}

// ClustersMultiError is an error wrapping multiple validation errors returned
// by Clusters.ValidateAll() if the designated constraints aren't met.
type ClustersMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClustersMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClustersMultiError) AllErrors() []error { return m }

// ClustersValidationError is the validation error returned by
// Clusters.Validate if the designated constraints aren't met.
type ClustersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClustersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClustersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClustersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClustersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClustersValidationError) ErrorName() string { return "ClustersValidationError" }

// Error satisfies the builtin error interface
func (e ClustersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusters.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClustersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClustersValidationError{}

// Validate checks the field values on ClusterStatus with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ClusterStatus) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClusterStatus with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ClusterStatusMultiError, or
// nil if none found.
func (m *ClusterStatus) ValidateAll() error {
	return m.validate(true)
}

func (m *ClusterStatus) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for AddedViaApi

	if all {
		switch v := interface{}(m.GetSuccessRateEjectionThreshold()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ClusterStatusValidationError{
					field:  "SuccessRateEjectionThreshold",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ClusterStatusValidationError{
					field:  "SuccessRateEjectionThreshold",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSuccessRateEjectionThreshold()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClusterStatusValidationError{
				field:  "SuccessRateEjectionThreshold",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetHostStatuses() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ClusterStatusValidationError{
						field:  fmt.Sprintf("HostStatuses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ClusterStatusValidationError{
						field:  fmt.Sprintf("HostStatuses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClusterStatusValidationError{
					field:  fmt.Sprintf("HostStatuses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetLocalOriginSuccessRateEjectionThreshold()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ClusterStatusValidationError{
					field:  "LocalOriginSuccessRateEjectionThreshold",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ClusterStatusValidationError{
					field:  "LocalOriginSuccessRateEjectionThreshold",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocalOriginSuccessRateEjectionThreshold()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClusterStatusValidationError{
				field:  "LocalOriginSuccessRateEjectionThreshold",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCircuitBreakers()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ClusterStatusValidationError{
					field:  "CircuitBreakers",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ClusterStatusValidationError{
					field:  "CircuitBreakers",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCircuitBreakers()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClusterStatusValidationError{
				field:  "CircuitBreakers",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ObservabilityName

	// no validation rules for EdsServiceName

	if len(errors) > 0 {
		return ClusterStatusMultiError(errors)
	}
	return nil
}

// ClusterStatusMultiError is an error wrapping multiple validation errors
// returned by ClusterStatus.ValidateAll() if the designated constraints
// aren't met.
type ClusterStatusMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClusterStatusMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClusterStatusMultiError) AllErrors() []error { return m }

// ClusterStatusValidationError is the validation error returned by
// ClusterStatus.Validate if the designated constraints aren't met.
type ClusterStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterStatusValidationError) ErrorName() string { return "ClusterStatusValidationError" }

// Error satisfies the builtin error interface
func (e ClusterStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterStatusValidationError{}

// Validate checks the field values on HostStatus with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HostStatus) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HostStatus with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HostStatusMultiError, or
// nil if none found.
func (m *HostStatus) ValidateAll() error {
	return m.validate(true)
}

func (m *HostStatus) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HostStatusValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HostStatusValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HostStatusValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetStats() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HostStatusValidationError{
						field:  fmt.Sprintf("Stats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HostStatusValidationError{
						field:  fmt.Sprintf("Stats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HostStatusValidationError{
					field:  fmt.Sprintf("Stats[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetHealthStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HostStatusValidationError{
					field:  "HealthStatus",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HostStatusValidationError{
					field:  "HealthStatus",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHealthStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HostStatusValidationError{
				field:  "HealthStatus",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetSuccessRate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HostStatusValidationError{
					field:  "SuccessRate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HostStatusValidationError{
					field:  "SuccessRate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSuccessRate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HostStatusValidationError{
				field:  "SuccessRate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Weight

	// no validation rules for Hostname

	// no validation rules for Priority

	if all {
		switch v := interface{}(m.GetLocalOriginSuccessRate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HostStatusValidationError{
					field:  "LocalOriginSuccessRate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HostStatusValidationError{
					field:  "LocalOriginSuccessRate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocalOriginSuccessRate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HostStatusValidationError{
				field:  "LocalOriginSuccessRate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetLocality()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HostStatusValidationError{
					field:  "Locality",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HostStatusValidationError{
					field:  "Locality",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocality()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HostStatusValidationError{
				field:  "Locality",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return HostStatusMultiError(errors)
	}
	return nil
}

// HostStatusMultiError is an error wrapping multiple validation errors
// returned by HostStatus.ValidateAll() if the designated constraints aren't met.
type HostStatusMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HostStatusMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HostStatusMultiError) AllErrors() []error { return m }

// HostStatusValidationError is the validation error returned by
// HostStatus.Validate if the designated constraints aren't met.
type HostStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HostStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HostStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HostStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HostStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HostStatusValidationError) ErrorName() string { return "HostStatusValidationError" }

// Error satisfies the builtin error interface
func (e HostStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHostStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HostStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HostStatusValidationError{}

// Validate checks the field values on HostHealthStatus with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *HostHealthStatus) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HostHealthStatus with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HostHealthStatusMultiError, or nil if none found.
func (m *HostHealthStatus) ValidateAll() error {
	return m.validate(true)
}

func (m *HostHealthStatus) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FailedActiveHealthCheck

	// no validation rules for FailedOutlierCheck

	// no validation rules for FailedActiveDegradedCheck

	// no validation rules for PendingDynamicRemoval

	// no validation rules for PendingActiveHc

	// no validation rules for ExcludedViaImmediateHcFail

	// no validation rules for ActiveHcTimeout

	// no validation rules for EdsHealthStatus

	if len(errors) > 0 {
		return HostHealthStatusMultiError(errors)
	}
	return nil
}

// HostHealthStatusMultiError is an error wrapping multiple validation errors
// returned by HostHealthStatus.ValidateAll() if the designated constraints
// aren't met.
type HostHealthStatusMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HostHealthStatusMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HostHealthStatusMultiError) AllErrors() []error { return m }

// HostHealthStatusValidationError is the validation error returned by
// HostHealthStatus.Validate if the designated constraints aren't met.
type HostHealthStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HostHealthStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HostHealthStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HostHealthStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HostHealthStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HostHealthStatusValidationError) ErrorName() string { return "HostHealthStatusValidationError" }

// Error satisfies the builtin error interface
func (e HostHealthStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHostHealthStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HostHealthStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HostHealthStatusValidationError{}
