// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AdRfc2307TypeParams Specifies the properties associated to a Rfc2307 type user id mapping.
//
// swagger:model AdRfc2307TypeParams
type AdRfc2307TypeParams struct {

	// Specifies a fallback user id mapping param in case the primary config does not work.
	// Required: true
	FallbackOption *FallbackUserIDMappingParams `json:"fallbackOption"`
}

// Validate validates this ad rfc2307 type params
func (m *AdRfc2307TypeParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFallbackOption(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdRfc2307TypeParams) validateFallbackOption(formats strfmt.Registry) error {

	if err := validate.Required("fallbackOption", "body", m.FallbackOption); err != nil {
		return err
	}

	if m.FallbackOption != nil {
		if err := m.FallbackOption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fallbackOption")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fallbackOption")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ad rfc2307 type params based on the context it is used
func (m *AdRfc2307TypeParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFallbackOption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdRfc2307TypeParams) contextValidateFallbackOption(ctx context.Context, formats strfmt.Registry) error {

	if m.FallbackOption != nil {

		if err := m.FallbackOption.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fallbackOption")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fallbackOption")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdRfc2307TypeParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdRfc2307TypeParams) UnmarshalBinary(b []byte) error {
	var res AdRfc2307TypeParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
