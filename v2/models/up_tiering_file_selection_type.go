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

// UpTieringFileSelectionType Data up-tiering file selection type.
//
// Data up-tiering file selection type.
//
// swagger:model UpTieringFileSelectionType
type UpTieringFileSelectionType struct {

	// Specifies the data up-tiering file selection type.
	// Enum: ["LastAccessed","LastModified"]
	Type string `json:"type,omitempty"`
}

// Validate validates this up tiering file selection type
func (m *UpTieringFileSelectionType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var upTieringFileSelectionTypeTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LastAccessed","LastModified"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		upTieringFileSelectionTypeTypeTypePropEnum = append(upTieringFileSelectionTypeTypeTypePropEnum, v)
	}
}

const (

	// UpTieringFileSelectionTypeTypeLastAccessed captures enum value "LastAccessed"
	UpTieringFileSelectionTypeTypeLastAccessed string = "LastAccessed"

	// UpTieringFileSelectionTypeTypeLastModified captures enum value "LastModified"
	UpTieringFileSelectionTypeTypeLastModified string = "LastModified"
)

// prop value enum
func (m *UpTieringFileSelectionType) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, upTieringFileSelectionTypeTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpTieringFileSelectionType) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this up tiering file selection type based on context it is used
func (m *UpTieringFileSelectionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpTieringFileSelectionType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpTieringFileSelectionType) UnmarshalBinary(b []byte) error {
	var res UpTieringFileSelectionType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
