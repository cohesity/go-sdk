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

// CompressionParams Compression Parameters
//
// Specifies parameters for compression.
//
// swagger:model CompressionParams
type CompressionParams struct {

	// Specifies copmpression type for a Storage Domain.
	// Enum: ["None","Low","High"]
	Type *string `json:"type,omitempty"`

	// Specifies whether inline compression is enabled. This field is appliciable only if inlineDeduplicationEnabled is set to true and compression is enabled.
	InlineEnabled *bool `json:"inlineEnabled,omitempty"`
}

// Validate validates this compression params
func (m *CompressionParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var compressionParamsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Low","High"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		compressionParamsTypeTypePropEnum = append(compressionParamsTypeTypePropEnum, v)
	}
}

const (

	// CompressionParamsTypeNone captures enum value "None"
	CompressionParamsTypeNone string = "None"

	// CompressionParamsTypeLow captures enum value "Low"
	CompressionParamsTypeLow string = "Low"

	// CompressionParamsTypeHigh captures enum value "High"
	CompressionParamsTypeHigh string = "High"
)

// prop value enum
func (m *CompressionParams) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, compressionParamsTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CompressionParams) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this compression params based on context it is used
func (m *CompressionParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CompressionParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompressionParams) UnmarshalBinary(b []byte) error {
	var res CompressionParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
