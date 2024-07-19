// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: orders.proto

package orders

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

// Validate checks the field values on AcceptOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AcceptOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AcceptOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AcceptOrderRequestMultiError, or nil if none found.
func (m *AcceptOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AcceptOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetOrderId() <= 0 {
		err := AcceptOrderRequestValidationError{
			field:  "OrderId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUserId() <= 0 {
		err := AcceptOrderRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if t := m.GetExpiry(); t != nil {
		ts, err := t.AsTime(), t.CheckValid()
		if err != nil {
			err = AcceptOrderRequestValidationError{
				field:  "Expiry",
				reason: "value is not a valid timestamp",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gt := time.Unix(0, 0)

			if ts.Sub(gt) <= 0 {
				err := AcceptOrderRequestValidationError{
					field:  "Expiry",
					reason: "value must be greater than 1970-01-01 00:00:00 +0000 UTC",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if m.GetCost() <= 0 {
		err := AcceptOrderRequestValidationError{
			field:  "Cost",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetWeight() <= 0 {
		err := AcceptOrderRequestValidationError{
			field:  "Weight",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _AcceptOrderRequest_Packaging_NotInLookup[m.GetPackaging()]; ok {
		err := AcceptOrderRequestValidationError{
			field:  "Packaging",
			reason: "value must not be in list [PACKAGING_UNSPECIFIED]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := PackagingType_name[int32(m.GetPackaging())]; !ok {
		err := AcceptOrderRequestValidationError{
			field:  "Packaging",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AcceptOrderRequestMultiError(errors)
	}

	return nil
}

// AcceptOrderRequestMultiError is an error wrapping multiple validation errors
// returned by AcceptOrderRequest.ValidateAll() if the designated constraints
// aren't met.
type AcceptOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AcceptOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AcceptOrderRequestMultiError) AllErrors() []error { return m }

// AcceptOrderRequestValidationError is the validation error returned by
// AcceptOrderRequest.Validate if the designated constraints aren't met.
type AcceptOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AcceptOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AcceptOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AcceptOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AcceptOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AcceptOrderRequestValidationError) ErrorName() string {
	return "AcceptOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AcceptOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAcceptOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AcceptOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AcceptOrderRequestValidationError{}

var _AcceptOrderRequest_Packaging_NotInLookup = map[PackagingType]struct{}{
	0: {},
}

// Validate checks the field values on AcceptOrderResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AcceptOrderResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AcceptOrderResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AcceptOrderResponseMultiError, or nil if none found.
func (m *AcceptOrderResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AcceptOrderResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AcceptOrderResponseMultiError(errors)
	}

	return nil
}

// AcceptOrderResponseMultiError is an error wrapping multiple validation
// errors returned by AcceptOrderResponse.ValidateAll() if the designated
// constraints aren't met.
type AcceptOrderResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AcceptOrderResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AcceptOrderResponseMultiError) AllErrors() []error { return m }

// AcceptOrderResponseValidationError is the validation error returned by
// AcceptOrderResponse.Validate if the designated constraints aren't met.
type AcceptOrderResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AcceptOrderResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AcceptOrderResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AcceptOrderResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AcceptOrderResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AcceptOrderResponseValidationError) ErrorName() string {
	return "AcceptOrderResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AcceptOrderResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAcceptOrderResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AcceptOrderResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AcceptOrderResponseValidationError{}

// Validate checks the field values on DeliverOrdersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeliverOrdersRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeliverOrdersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeliverOrdersRequestMultiError, or nil if none found.
func (m *DeliverOrdersRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeliverOrdersRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetOrderIds()) < 1 {
		err := DeliverOrdersRequestValidationError{
			field:  "OrderIds",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeliverOrdersRequestMultiError(errors)
	}

	return nil
}

// DeliverOrdersRequestMultiError is an error wrapping multiple validation
// errors returned by DeliverOrdersRequest.ValidateAll() if the designated
// constraints aren't met.
type DeliverOrdersRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeliverOrdersRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeliverOrdersRequestMultiError) AllErrors() []error { return m }

// DeliverOrdersRequestValidationError is the validation error returned by
// DeliverOrdersRequest.Validate if the designated constraints aren't met.
type DeliverOrdersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeliverOrdersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeliverOrdersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeliverOrdersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeliverOrdersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeliverOrdersRequestValidationError) ErrorName() string {
	return "DeliverOrdersRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeliverOrdersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeliverOrdersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeliverOrdersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeliverOrdersRequestValidationError{}

// Validate checks the field values on DeliverOrdersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeliverOrdersResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeliverOrdersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeliverOrdersResponseMultiError, or nil if none found.
func (m *DeliverOrdersResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeliverOrdersResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeliverOrdersResponseMultiError(errors)
	}

	return nil
}

// DeliverOrdersResponseMultiError is an error wrapping multiple validation
// errors returned by DeliverOrdersResponse.ValidateAll() if the designated
// constraints aren't met.
type DeliverOrdersResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeliverOrdersResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeliverOrdersResponseMultiError) AllErrors() []error { return m }

// DeliverOrdersResponseValidationError is the validation error returned by
// DeliverOrdersResponse.Validate if the designated constraints aren't met.
type DeliverOrdersResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeliverOrdersResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeliverOrdersResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeliverOrdersResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeliverOrdersResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeliverOrdersResponseValidationError) ErrorName() string {
	return "DeliverOrdersResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeliverOrdersResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeliverOrdersResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeliverOrdersResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeliverOrdersResponseValidationError{}

// Validate checks the field values on ListOrdersRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListOrdersRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListOrdersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListOrdersRequestMultiError, or nil if none found.
func (m *ListOrdersRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListOrdersRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() <= 0 {
		err := ListOrdersRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetLimit() < 0 {
		err := ListOrdersRequestValidationError{
			field:  "Limit",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetOffset() < 0 {
		err := ListOrdersRequestValidationError{
			field:  "Offset",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for IsStored

	if len(errors) > 0 {
		return ListOrdersRequestMultiError(errors)
	}

	return nil
}

// ListOrdersRequestMultiError is an error wrapping multiple validation errors
// returned by ListOrdersRequest.ValidateAll() if the designated constraints
// aren't met.
type ListOrdersRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListOrdersRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListOrdersRequestMultiError) AllErrors() []error { return m }

// ListOrdersRequestValidationError is the validation error returned by
// ListOrdersRequest.Validate if the designated constraints aren't met.
type ListOrdersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOrdersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOrdersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOrdersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOrdersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOrdersRequestValidationError) ErrorName() string {
	return "ListOrdersRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListOrdersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOrdersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOrdersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOrdersRequestValidationError{}

// Validate checks the field values on ListOrdersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListOrdersResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListOrdersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListOrdersResponseMultiError, or nil if none found.
func (m *ListOrdersResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListOrdersResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetOrders() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListOrdersResponseValidationError{
						field:  fmt.Sprintf("Orders[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListOrdersResponseValidationError{
						field:  fmt.Sprintf("Orders[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListOrdersResponseValidationError{
					field:  fmt.Sprintf("Orders[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListOrdersResponseMultiError(errors)
	}

	return nil
}

// ListOrdersResponseMultiError is an error wrapping multiple validation errors
// returned by ListOrdersResponse.ValidateAll() if the designated constraints
// aren't met.
type ListOrdersResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListOrdersResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListOrdersResponseMultiError) AllErrors() []error { return m }

// ListOrdersResponseValidationError is the validation error returned by
// ListOrdersResponse.Validate if the designated constraints aren't met.
type ListOrdersResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOrdersResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOrdersResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOrdersResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOrdersResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOrdersResponseValidationError) ErrorName() string {
	return "ListOrdersResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListOrdersResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOrdersResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOrdersResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOrdersResponseValidationError{}

// Validate checks the field values on ReturnOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ReturnOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReturnOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReturnOrderRequestMultiError, or nil if none found.
func (m *ReturnOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ReturnOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetOrderId() <= 0 {
		err := ReturnOrderRequestValidationError{
			field:  "OrderId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ReturnOrderRequestMultiError(errors)
	}

	return nil
}

// ReturnOrderRequestMultiError is an error wrapping multiple validation errors
// returned by ReturnOrderRequest.ValidateAll() if the designated constraints
// aren't met.
type ReturnOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReturnOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReturnOrderRequestMultiError) AllErrors() []error { return m }

// ReturnOrderRequestValidationError is the validation error returned by
// ReturnOrderRequest.Validate if the designated constraints aren't met.
type ReturnOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReturnOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReturnOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReturnOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReturnOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReturnOrderRequestValidationError) ErrorName() string {
	return "ReturnOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ReturnOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReturnOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReturnOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReturnOrderRequestValidationError{}

// Validate checks the field values on ReturnOrderResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ReturnOrderResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReturnOrderResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReturnOrderResponseMultiError, or nil if none found.
func (m *ReturnOrderResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ReturnOrderResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ReturnOrderResponseMultiError(errors)
	}

	return nil
}

// ReturnOrderResponseMultiError is an error wrapping multiple validation
// errors returned by ReturnOrderResponse.ValidateAll() if the designated
// constraints aren't met.
type ReturnOrderResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReturnOrderResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReturnOrderResponseMultiError) AllErrors() []error { return m }

// ReturnOrderResponseValidationError is the validation error returned by
// ReturnOrderResponse.Validate if the designated constraints aren't met.
type ReturnOrderResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReturnOrderResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReturnOrderResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReturnOrderResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReturnOrderResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReturnOrderResponseValidationError) ErrorName() string {
	return "ReturnOrderResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ReturnOrderResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReturnOrderResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReturnOrderResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReturnOrderResponseValidationError{}

// Validate checks the field values on ListOrdersResponse_Order with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListOrdersResponse_Order) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListOrdersResponse_Order with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListOrdersResponse_OrderMultiError, or nil if none found.
func (m *ListOrdersResponse_Order) ValidateAll() error {
	return m.validate(true)
}

func (m *ListOrdersResponse_Order) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	// no validation rules for UserId

	if all {
		switch v := interface{}(m.GetExpiry()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListOrdersResponse_OrderValidationError{
					field:  "Expiry",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListOrdersResponse_OrderValidationError{
					field:  "Expiry",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExpiry()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListOrdersResponse_OrderValidationError{
				field:  "Expiry",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetReturnBy()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListOrdersResponse_OrderValidationError{
					field:  "ReturnBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListOrdersResponse_OrderValidationError{
					field:  "ReturnBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReturnBy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListOrdersResponse_OrderValidationError{
				field:  "ReturnBy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Status

	// no validation rules for Hash

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListOrdersResponse_OrderValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListOrdersResponse_OrderValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListOrdersResponse_OrderValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Cost

	// no validation rules for Weight

	// no validation rules for Packaging

	if len(errors) > 0 {
		return ListOrdersResponse_OrderMultiError(errors)
	}

	return nil
}

// ListOrdersResponse_OrderMultiError is an error wrapping multiple validation
// errors returned by ListOrdersResponse_Order.ValidateAll() if the designated
// constraints aren't met.
type ListOrdersResponse_OrderMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListOrdersResponse_OrderMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListOrdersResponse_OrderMultiError) AllErrors() []error { return m }

// ListOrdersResponse_OrderValidationError is the validation error returned by
// ListOrdersResponse_Order.Validate if the designated constraints aren't met.
type ListOrdersResponse_OrderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOrdersResponse_OrderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOrdersResponse_OrderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOrdersResponse_OrderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOrdersResponse_OrderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOrdersResponse_OrderValidationError) ErrorName() string {
	return "ListOrdersResponse_OrderValidationError"
}

// Error satisfies the builtin error interface
func (e ListOrdersResponse_OrderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOrdersResponse_Order.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOrdersResponse_OrderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOrdersResponse_OrderValidationError{}
