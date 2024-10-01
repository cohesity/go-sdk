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

// CompressionType Target Compression Type
//
// Specifies the compression type of the external target.
//
// swagger:model CompressionType
type CompressionType struct {

	// Specifies the compression type of the external target.
	// Enum: ["None","Low","High"]
	Enum string `json:"enum,omitempty"`
}

// Validate validates this compression type
func (m *CompressionType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnum(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var compressionTypeTypeEnumPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Low","High"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		compressionTypeTypeEnumPropEnum = append(compressionTypeTypeEnumPropEnum, v)
	}
}

const (

	// CompressionTypeEnumNone captures enum value "None"
	CompressionTypeEnumNone string = "None"

	// CompressionTypeEnumLow captures enum value "Low"
	CompressionTypeEnumLow string = "Low"

	// CompressionTypeEnumHigh captures enum value "High"
	CompressionTypeEnumHigh string = "High"
)

// prop value enum
func (m *CompressionType) validateEnumEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, compressionTypeTypeEnumPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CompressionType) validateEnum(formats strfmt.Registry) error {
	if swag.IsZero(m.Enum) { // not required
		return nil
	}

	// value enum
	if err := m.validateEnumEnum("enum", "body", m.Enum); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this compression type based on context it is used
func (m *CompressionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CompressionType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompressionType) UnmarshalBinary(b []byte) error {
	var res CompressionType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
