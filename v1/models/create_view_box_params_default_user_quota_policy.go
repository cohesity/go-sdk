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

// CreateViewBoxParamsDefaultUserQuotaPolicy Specifies an optional quota policy/limits that are inherited by all users
// within the views in this viewbox.
//
// swagger:model createViewBoxParamsDefaultUserQuotaPolicy
type CreateViewBoxParamsDefaultUserQuotaPolicy struct {
	QuotaPolicy
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CreateViewBoxParamsDefaultUserQuotaPolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 QuotaPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.QuotaPolicy = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CreateViewBoxParamsDefaultUserQuotaPolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.QuotaPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this create view box params default user quota policy
func (m *CreateViewBoxParamsDefaultUserQuotaPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with QuotaPolicy
	if err := m.QuotaPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this create view box params default user quota policy based on the context it is used
func (m *CreateViewBoxParamsDefaultUserQuotaPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with QuotaPolicy
	if err := m.QuotaPolicy.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CreateViewBoxParamsDefaultUserQuotaPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateViewBoxParamsDefaultUserQuotaPolicy) UnmarshalBinary(b []byte) error {
	var res CreateViewBoxParamsDefaultUserQuotaPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
