// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProtectionSourceElastifileProtectionSource Specifies details about a Elastifile Protection Source
// when the environment is set to 'kElastifile'.
//
// swagger:model protectionSourceElastifileProtectionSource
type ProtectionSourceElastifileProtectionSource struct {
	ElastifileProtectionSource
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ProtectionSourceElastifileProtectionSource) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ElastifileProtectionSource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ElastifileProtectionSource = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ProtectionSourceElastifileProtectionSource) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ElastifileProtectionSource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this protection source elastifile protection source
func (m *ProtectionSourceElastifileProtectionSource) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ElastifileProtectionSource
	if err := m.ElastifileProtectionSource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this protection source elastifile protection source based on the context it is used
func (m *ProtectionSourceElastifileProtectionSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ElastifileProtectionSource
	if err := m.ElastifileProtectionSource.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ProtectionSourceElastifileProtectionSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectionSourceElastifileProtectionSource) UnmarshalBinary(b []byte) error {
	var res ProtectionSourceElastifileProtectionSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
