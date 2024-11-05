// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: protobuf/api/order.proto

package commerce_api

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

// Validate checks the field values on CheckoutOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CheckoutOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckoutOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CheckoutOrderRequestMultiError, or nil if none found.
func (m *CheckoutOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckoutOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ShopId

	// no validation rules for WarehouseId

	// no validation rules for ProductId

	// no validation rules for Quantity

	if len(errors) > 0 {
		return CheckoutOrderRequestMultiError(errors)
	}

	return nil
}

// CheckoutOrderRequestMultiError is an error wrapping multiple validation
// errors returned by CheckoutOrderRequest.ValidateAll() if the designated
// constraints aren't met.
type CheckoutOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckoutOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckoutOrderRequestMultiError) AllErrors() []error { return m }

// CheckoutOrderRequestValidationError is the validation error returned by
// CheckoutOrderRequest.Validate if the designated constraints aren't met.
type CheckoutOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckoutOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckoutOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckoutOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckoutOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckoutOrderRequestValidationError) ErrorName() string {
	return "CheckoutOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CheckoutOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckoutOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckoutOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckoutOrderRequestValidationError{}

// Validate checks the field values on OrderIDRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrderIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderIDRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrderIDRequestMultiError,
// or nil if none found.
func (m *OrderIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	if len(errors) > 0 {
		return OrderIDRequestMultiError(errors)
	}

	return nil
}

// OrderIDRequestMultiError is an error wrapping multiple validation errors
// returned by OrderIDRequest.ValidateAll() if the designated constraints
// aren't met.
type OrderIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderIDRequestMultiError) AllErrors() []error { return m }

// OrderIDRequestValidationError is the validation error returned by
// OrderIDRequest.Validate if the designated constraints aren't met.
type OrderIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderIDRequestValidationError) ErrorName() string { return "OrderIDRequestValidationError" }

// Error satisfies the builtin error interface
func (e OrderIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderIDRequestValidationError{}
