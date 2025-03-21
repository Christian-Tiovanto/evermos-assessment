// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: protobuf/api/shop.proto

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

// Validate checks the field values on CreateShopRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateShopRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateShopRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateShopRequestMultiError, or nil if none found.
func (m *CreateShopRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateShopRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Address

	if len(errors) > 0 {
		return CreateShopRequestMultiError(errors)
	}

	return nil
}

// CreateShopRequestMultiError is an error wrapping multiple validation errors
// returned by CreateShopRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateShopRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateShopRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateShopRequestMultiError) AllErrors() []error { return m }

// CreateShopRequestValidationError is the validation error returned by
// CreateShopRequest.Validate if the designated constraints aren't met.
type CreateShopRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateShopRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateShopRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateShopRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateShopRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateShopRequestValidationError) ErrorName() string {
	return "CreateShopRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateShopRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateShopRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateShopRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateShopRequestValidationError{}

// Validate checks the field values on CreateWarehouseRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateWarehouseRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateWarehouseRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateWarehouseRequestMultiError, or nil if none found.
func (m *CreateWarehouseRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateWarehouseRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ShopId

	// no validation rules for Name

	if len(errors) > 0 {
		return CreateWarehouseRequestMultiError(errors)
	}

	return nil
}

// CreateWarehouseRequestMultiError is an error wrapping multiple validation
// errors returned by CreateWarehouseRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateWarehouseRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateWarehouseRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateWarehouseRequestMultiError) AllErrors() []error { return m }

// CreateWarehouseRequestValidationError is the validation error returned by
// CreateWarehouseRequest.Validate if the designated constraints aren't met.
type CreateWarehouseRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateWarehouseRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateWarehouseRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateWarehouseRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateWarehouseRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateWarehouseRequestValidationError) ErrorName() string {
	return "CreateWarehouseRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateWarehouseRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateWarehouseRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateWarehouseRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateWarehouseRequestValidationError{}

// Validate checks the field values on ShopIDRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ShopIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ShopIDRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ShopIDRequestMultiError, or
// nil if none found.
func (m *ShopIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ShopIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ShopId

	if len(errors) > 0 {
		return ShopIDRequestMultiError(errors)
	}

	return nil
}

// ShopIDRequestMultiError is an error wrapping multiple validation errors
// returned by ShopIDRequest.ValidateAll() if the designated constraints
// aren't met.
type ShopIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ShopIDRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ShopIDRequestMultiError) AllErrors() []error { return m }

// ShopIDRequestValidationError is the validation error returned by
// ShopIDRequest.Validate if the designated constraints aren't met.
type ShopIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ShopIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ShopIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ShopIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ShopIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ShopIDRequestValidationError) ErrorName() string { return "ShopIDRequestValidationError" }

// Error satisfies the builtin error interface
func (e ShopIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sShopIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ShopIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ShopIDRequestValidationError{}

// Validate checks the field values on Warehouse with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Warehouse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Warehouse with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in WarehouseMultiError, or nil
// if none found.
func (m *Warehouse) ValidateAll() error {
	return m.validate(true)
}

func (m *Warehouse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Status

	// no validation rules for ShopId

	if len(errors) > 0 {
		return WarehouseMultiError(errors)
	}

	return nil
}

// WarehouseMultiError is an error wrapping multiple validation errors returned
// by Warehouse.ValidateAll() if the designated constraints aren't met.
type WarehouseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WarehouseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WarehouseMultiError) AllErrors() []error { return m }

// WarehouseValidationError is the validation error returned by
// Warehouse.Validate if the designated constraints aren't met.
type WarehouseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WarehouseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WarehouseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WarehouseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WarehouseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WarehouseValidationError) ErrorName() string { return "WarehouseValidationError" }

// Error satisfies the builtin error interface
func (e WarehouseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWarehouse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WarehouseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WarehouseValidationError{}

// Validate checks the field values on GetWarehousesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetWarehousesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetWarehousesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetWarehousesResponseMultiError, or nil if none found.
func (m *GetWarehousesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetWarehousesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetWarehouses() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetWarehousesResponseValidationError{
						field:  fmt.Sprintf("Warehouses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetWarehousesResponseValidationError{
						field:  fmt.Sprintf("Warehouses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetWarehousesResponseValidationError{
					field:  fmt.Sprintf("Warehouses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetWarehousesResponseMultiError(errors)
	}

	return nil
}

// GetWarehousesResponseMultiError is an error wrapping multiple validation
// errors returned by GetWarehousesResponse.ValidateAll() if the designated
// constraints aren't met.
type GetWarehousesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetWarehousesResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetWarehousesResponseMultiError) AllErrors() []error { return m }

// GetWarehousesResponseValidationError is the validation error returned by
// GetWarehousesResponse.Validate if the designated constraints aren't met.
type GetWarehousesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetWarehousesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetWarehousesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetWarehousesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetWarehousesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetWarehousesResponseValidationError) ErrorName() string {
	return "GetWarehousesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetWarehousesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetWarehousesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetWarehousesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetWarehousesResponseValidationError{}

// Validate checks the field values on AddProductRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddProductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddProductRequestMultiError, or nil if none found.
func (m *AddProductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddProductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ShopId

	// no validation rules for WarehouseId

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for Price

	// no validation rules for Quantity

	if len(errors) > 0 {
		return AddProductRequestMultiError(errors)
	}

	return nil
}

// AddProductRequestMultiError is an error wrapping multiple validation errors
// returned by AddProductRequest.ValidateAll() if the designated constraints
// aren't met.
type AddProductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddProductRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddProductRequestMultiError) AllErrors() []error { return m }

// AddProductRequestValidationError is the validation error returned by
// AddProductRequest.Validate if the designated constraints aren't met.
type AddProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddProductRequestValidationError) ErrorName() string {
	return "AddProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddProductRequestValidationError{}
