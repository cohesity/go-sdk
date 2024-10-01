// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// S3ACLPermission S3 ACL Permission
//
// Specifies S3 ACL permission type.
//
// swagger:model S3AclPermission
type S3ACLPermission struct {

	// Specifies S3 ACL permission type.
	// Enum: ["Read","Write","ReadACP","WriteACP","FullControl"]
	Enum string `json:"enum,omitempty"`
}

// Validate validates this s3 Acl permission
func (m *S3ACLPermission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnum(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var s3AclPermissionTypeEnumPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Read","Write","ReadACP","WriteACP","FullControl"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		s3AclPermissionTypeEnumPropEnum = append(s3AclPermissionTypeEnumPropEnum, v)
	}
}

const (

	// S3ACLPermissionEnumRead captures enum value "Read"
	S3ACLPermissionEnumRead string = "Read"

	// S3ACLPermissionEnumWrite captures enum value "Write"
	S3ACLPermissionEnumWrite string = "Write"

	// S3ACLPermissionEnumReadACP captures enum value "ReadACP"
	S3ACLPermissionEnumReadACP string = "ReadACP"

	// S3ACLPermissionEnumWriteACP captures enum value "WriteACP"
	S3ACLPermissionEnumWriteACP string = "WriteACP"

	// S3ACLPermissionEnumFullControl captures enum value "FullControl"
	S3ACLPermissionEnumFullControl string = "FullControl"
)

// prop value enum
func (m *S3ACLPermission) validateEnumEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, s3AclPermissionTypeEnumPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *S3ACLPermission) validateEnum(formats strfmt.Registry) error {
	if swag.IsZero(m.Enum) { // not required
		return nil
	}

	// value enum
	if err := m.validateEnumEnum("enum", "body", m.Enum); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this s3 Acl permission based on context it is used
func (m *S3ACLPermission) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *S3ACLPermission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3ACLPermission) UnmarshalBinary(b []byte) error {
	var res S3ACLPermission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
