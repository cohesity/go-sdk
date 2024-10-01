// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NisProviders Response of NIS Providers.
//
// swagger:model NisProviders
type NisProviders struct {

	// A list of NIS Providers.
	NisProviders []*NisProvider `json:"nisProviders"`
}

// Validate validates this nis providers
func (m *NisProviders) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNisProviders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NisProviders) validateNisProviders(formats strfmt.Registry) error {
	if swag.IsZero(m.NisProviders) { // not required
		return nil
	}

	for i := 0; i < len(m.NisProviders); i++ {
		if swag.IsZero(m.NisProviders[i]) { // not required
			continue
		}

		if m.NisProviders[i] != nil {
			if err := m.NisProviders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nisProviders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nisProviders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this nis providers based on the context it is used
func (m *NisProviders) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNisProviders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NisProviders) contextValidateNisProviders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NisProviders); i++ {

		if m.NisProviders[i] != nil {

			if swag.IsZero(m.NisProviders[i]) { // not required
				return nil
			}

			if err := m.NisProviders[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nisProviders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nisProviders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NisProviders) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NisProviders) UnmarshalBinary(b []byte) error {
	var res NisProviders
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
