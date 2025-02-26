// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: authzed/api/v0/acl_service.proto

package v0

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on RelationTupleFilter with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RelationTupleFilter) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetNamespace()) > 128 {
		return RelationTupleFilterValidationError{
			field:  "Namespace",
			reason: "value length must be at most 128 bytes",
		}
	}

	if !_RelationTupleFilter_Namespace_Pattern.MatchString(m.GetNamespace()) {
		return RelationTupleFilterValidationError{
			field:  "Namespace",
			reason: "value does not match regex pattern \"^([a-z][a-z0-9_]{2,61}[a-z0-9]/)?[a-z][a-z0-9_]{2,62}[a-z0-9]$\"",
		}
	}

	if len(m.GetObjectId()) > 128 {
		return RelationTupleFilterValidationError{
			field:  "ObjectId",
			reason: "value length must be at most 128 bytes",
		}
	}

	if !_RelationTupleFilter_ObjectId_Pattern.MatchString(m.GetObjectId()) {
		return RelationTupleFilterValidationError{
			field:  "ObjectId",
			reason: "value does not match regex pattern \"^([a-zA-Z0-9_][a-zA-Z0-9/_-]{0,127})?$\"",
		}
	}

	if len(m.GetRelation()) > 64 {
		return RelationTupleFilterValidationError{
			field:  "Relation",
			reason: "value length must be at most 64 bytes",
		}
	}

	if !_RelationTupleFilter_Relation_Pattern.MatchString(m.GetRelation()) {
		return RelationTupleFilterValidationError{
			field:  "Relation",
			reason: "value does not match regex pattern \"^([a-z][a-z0-9_]{2,62}[a-z0-9])?$\"",
		}
	}

	if v, ok := interface{}(m.GetUserset()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RelationTupleFilterValidationError{
				field:  "Userset",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetFilters() {
		_, _ = idx, item

		if _, ok := RelationTupleFilter_Filter_name[int32(item)]; !ok {
			return RelationTupleFilterValidationError{
				field:  fmt.Sprintf("Filters[%v]", idx),
				reason: "value must be one of the defined enum values",
			}
		}

	}

	return nil
}

// RelationTupleFilterValidationError is the validation error returned by
// RelationTupleFilter.Validate if the designated constraints aren't met.
type RelationTupleFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RelationTupleFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RelationTupleFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RelationTupleFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RelationTupleFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RelationTupleFilterValidationError) ErrorName() string {
	return "RelationTupleFilterValidationError"
}

// Error satisfies the builtin error interface
func (e RelationTupleFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRelationTupleFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RelationTupleFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RelationTupleFilterValidationError{}

var _RelationTupleFilter_Namespace_Pattern = regexp.MustCompile("^([a-z][a-z0-9_]{2,61}[a-z0-9]/)?[a-z][a-z0-9_]{2,62}[a-z0-9]$")

var _RelationTupleFilter_ObjectId_Pattern = regexp.MustCompile("^([a-zA-Z0-9_][a-zA-Z0-9/_-]{0,127})?$")

var _RelationTupleFilter_Relation_Pattern = regexp.MustCompile("^([a-z][a-z0-9_]{2,62}[a-z0-9])?$")

// Validate checks the field values on ReadRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ReadRequest) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetTuplesets()) < 1 {
		return ReadRequestValidationError{
			field:  "Tuplesets",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetTuplesets() {
		_, _ = idx, item

		if item == nil {
			return ReadRequestValidationError{
				field:  fmt.Sprintf("Tuplesets[%v]", idx),
				reason: "value is required",
			}
		}

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ReadRequestValidationError{
					field:  fmt.Sprintf("Tuplesets[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetAtRevision()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReadRequestValidationError{
				field:  "AtRevision",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ReadRequestValidationError is the validation error returned by
// ReadRequest.Validate if the designated constraints aren't met.
type ReadRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadRequestValidationError) ErrorName() string { return "ReadRequestValidationError" }

// Error satisfies the builtin error interface
func (e ReadRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadRequestValidationError{}

// Validate checks the field values on ReadResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ReadResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetTuplesets() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ReadResponseValidationError{
					field:  fmt.Sprintf("Tuplesets[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetRevision()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReadResponseValidationError{
				field:  "Revision",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ReadResponseValidationError is the validation error returned by
// ReadResponse.Validate if the designated constraints aren't met.
type ReadResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadResponseValidationError) ErrorName() string { return "ReadResponseValidationError" }

// Error satisfies the builtin error interface
func (e ReadResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadResponseValidationError{}

// Validate checks the field values on WriteRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *WriteRequest) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetWriteConditions() {
		_, _ = idx, item

		if item == nil {
			return WriteRequestValidationError{
				field:  fmt.Sprintf("WriteConditions[%v]", idx),
				reason: "value is required",
			}
		}

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WriteRequestValidationError{
					field:  fmt.Sprintf("WriteConditions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetUpdates() {
		_, _ = idx, item

		if item == nil {
			return WriteRequestValidationError{
				field:  fmt.Sprintf("Updates[%v]", idx),
				reason: "value is required",
			}
		}

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WriteRequestValidationError{
					field:  fmt.Sprintf("Updates[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// WriteRequestValidationError is the validation error returned by
// WriteRequest.Validate if the designated constraints aren't met.
type WriteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WriteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WriteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WriteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WriteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WriteRequestValidationError) ErrorName() string { return "WriteRequestValidationError" }

// Error satisfies the builtin error interface
func (e WriteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWriteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WriteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WriteRequestValidationError{}

// Validate checks the field values on WriteResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *WriteResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRevision()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WriteResponseValidationError{
				field:  "Revision",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// WriteResponseValidationError is the validation error returned by
// WriteResponse.Validate if the designated constraints aren't met.
type WriteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WriteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WriteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WriteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WriteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WriteResponseValidationError) ErrorName() string { return "WriteResponseValidationError" }

// Error satisfies the builtin error interface
func (e WriteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWriteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WriteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WriteResponseValidationError{}

// Validate checks the field values on CheckRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CheckRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetTestUserset() == nil {
		return CheckRequestValidationError{
			field:  "TestUserset",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetTestUserset()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckRequestValidationError{
				field:  "TestUserset",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetUser() == nil {
		return CheckRequestValidationError{
			field:  "User",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckRequestValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAtRevision()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckRequestValidationError{
				field:  "AtRevision",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CheckRequestValidationError is the validation error returned by
// CheckRequest.Validate if the designated constraints aren't met.
type CheckRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckRequestValidationError) ErrorName() string { return "CheckRequestValidationError" }

// Error satisfies the builtin error interface
func (e CheckRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckRequestValidationError{}

// Validate checks the field values on ContentChangeCheckRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ContentChangeCheckRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetTestUserset() == nil {
		return ContentChangeCheckRequestValidationError{
			field:  "TestUserset",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetTestUserset()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ContentChangeCheckRequestValidationError{
				field:  "TestUserset",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetUser() == nil {
		return ContentChangeCheckRequestValidationError{
			field:  "User",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ContentChangeCheckRequestValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ContentChangeCheckRequestValidationError is the validation error returned by
// ContentChangeCheckRequest.Validate if the designated constraints aren't met.
type ContentChangeCheckRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ContentChangeCheckRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ContentChangeCheckRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ContentChangeCheckRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ContentChangeCheckRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ContentChangeCheckRequestValidationError) ErrorName() string {
	return "ContentChangeCheckRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ContentChangeCheckRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContentChangeCheckRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ContentChangeCheckRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ContentChangeCheckRequestValidationError{}

// Validate checks the field values on CheckResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CheckResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for IsMember

	if v, ok := interface{}(m.GetRevision()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckResponseValidationError{
				field:  "Revision",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Membership

	return nil
}

// CheckResponseValidationError is the validation error returned by
// CheckResponse.Validate if the designated constraints aren't met.
type CheckResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResponseValidationError) ErrorName() string { return "CheckResponseValidationError" }

// Error satisfies the builtin error interface
func (e CheckResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResponseValidationError{}

// Validate checks the field values on ExpandRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ExpandRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUserset() == nil {
		return ExpandRequestValidationError{
			field:  "Userset",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetUserset()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExpandRequestValidationError{
				field:  "Userset",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAtRevision()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExpandRequestValidationError{
				field:  "AtRevision",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ExpandRequestValidationError is the validation error returned by
// ExpandRequest.Validate if the designated constraints aren't met.
type ExpandRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExpandRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExpandRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExpandRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExpandRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExpandRequestValidationError) ErrorName() string { return "ExpandRequestValidationError" }

// Error satisfies the builtin error interface
func (e ExpandRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExpandRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExpandRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExpandRequestValidationError{}

// Validate checks the field values on ExpandResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ExpandResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTreeNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExpandResponseValidationError{
				field:  "TreeNode",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetRevision()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExpandResponseValidationError{
				field:  "Revision",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ExpandResponseValidationError is the validation error returned by
// ExpandResponse.Validate if the designated constraints aren't met.
type ExpandResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExpandResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExpandResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExpandResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExpandResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExpandResponseValidationError) ErrorName() string { return "ExpandResponseValidationError" }

// Error satisfies the builtin error interface
func (e ExpandResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExpandResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExpandResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExpandResponseValidationError{}

// Validate checks the field values on LookupRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LookupRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetObjectRelation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LookupRequestValidationError{
				field:  "ObjectRelation",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LookupRequestValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAtRevision()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LookupRequestValidationError{
				field:  "AtRevision",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for PageReference

	// no validation rules for Limit

	return nil
}

// LookupRequestValidationError is the validation error returned by
// LookupRequest.Validate if the designated constraints aren't met.
type LookupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LookupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LookupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LookupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LookupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LookupRequestValidationError) ErrorName() string { return "LookupRequestValidationError" }

// Error satisfies the builtin error interface
func (e LookupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLookupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LookupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LookupRequestValidationError{}

// Validate checks the field values on LookupResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LookupResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for NextPageReference

	if v, ok := interface{}(m.GetRevision()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LookupResponseValidationError{
				field:  "Revision",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// LookupResponseValidationError is the validation error returned by
// LookupResponse.Validate if the designated constraints aren't met.
type LookupResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LookupResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LookupResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LookupResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LookupResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LookupResponseValidationError) ErrorName() string { return "LookupResponseValidationError" }

// Error satisfies the builtin error interface
func (e LookupResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLookupResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LookupResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LookupResponseValidationError{}

// Validate checks the field values on ReadResponse_Tupleset with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ReadResponse_Tupleset) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetTuples() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ReadResponse_TuplesetValidationError{
					field:  fmt.Sprintf("Tuples[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ReadResponse_TuplesetValidationError is the validation error returned by
// ReadResponse_Tupleset.Validate if the designated constraints aren't met.
type ReadResponse_TuplesetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadResponse_TuplesetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadResponse_TuplesetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadResponse_TuplesetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadResponse_TuplesetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadResponse_TuplesetValidationError) ErrorName() string {
	return "ReadResponse_TuplesetValidationError"
}

// Error satisfies the builtin error interface
func (e ReadResponse_TuplesetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadResponse_Tupleset.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadResponse_TuplesetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadResponse_TuplesetValidationError{}
