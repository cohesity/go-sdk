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

// UdaHostOsSpecificParams Connector parameters specific to an operating system.
//
// swagger:model UdaHostOsSpecificParams
type UdaHostOsSpecificParams struct {

	// Type of the host operating system.
	HostOsType *string `json:"hostOsType,omitempty"`

	// index
	Index *UdaOSIndexConfigParams `json:"index,omitempty"`

	// registration
	Registration *UdaRegistrationParams `json:"registration,omitempty"`

	// protection
	Protection *UdaProtectionParams `json:"protection,omitempty"`

	// recovery
	Recovery *UdaRecoveryParams `json:"recovery,omitempty"`
}

// Validate validates this uda host os specific params
func (m *UdaHostOsSpecificParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecovery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UdaHostOsSpecificParams) validateIndex(formats strfmt.Registry) error {
	if swag.IsZero(m.Index) { // not required
		return nil
	}

	if m.Index != nil {
		if err := m.Index.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("index")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("index")
			}
			return err
		}
	}

	return nil
}

func (m *UdaHostOsSpecificParams) validateRegistration(formats strfmt.Registry) error {
	if swag.IsZero(m.Registration) { // not required
		return nil
	}

	if m.Registration != nil {
		if err := m.Registration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registration")
			}
			return err
		}
	}

	return nil
}

func (m *UdaHostOsSpecificParams) validateProtection(formats strfmt.Registry) error {
	if swag.IsZero(m.Protection) { // not required
		return nil
	}

	if m.Protection != nil {
		if err := m.Protection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protection")
			}
			return err
		}
	}

	return nil
}

func (m *UdaHostOsSpecificParams) validateRecovery(formats strfmt.Registry) error {
	if swag.IsZero(m.Recovery) { // not required
		return nil
	}

	if m.Recovery != nil {
		if err := m.Recovery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recovery")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recovery")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this uda host os specific params based on the context it is used
func (m *UdaHostOsSpecificParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIndex(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegistration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecovery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UdaHostOsSpecificParams) contextValidateIndex(ctx context.Context, formats strfmt.Registry) error {

	if m.Index != nil {

		if swag.IsZero(m.Index) { // not required
			return nil
		}

		if err := m.Index.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("index")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("index")
			}
			return err
		}
	}

	return nil
}

func (m *UdaHostOsSpecificParams) contextValidateRegistration(ctx context.Context, formats strfmt.Registry) error {

	if m.Registration != nil {

		if swag.IsZero(m.Registration) { // not required
			return nil
		}

		if err := m.Registration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registration")
			}
			return err
		}
	}

	return nil
}

func (m *UdaHostOsSpecificParams) contextValidateProtection(ctx context.Context, formats strfmt.Registry) error {

	if m.Protection != nil {

		if swag.IsZero(m.Protection) { // not required
			return nil
		}

		if err := m.Protection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protection")
			}
			return err
		}
	}

	return nil
}

func (m *UdaHostOsSpecificParams) contextValidateRecovery(ctx context.Context, formats strfmt.Registry) error {

	if m.Recovery != nil {

		if swag.IsZero(m.Recovery) { // not required
			return nil
		}

		if err := m.Recovery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recovery")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recovery")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UdaHostOsSpecificParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UdaHostOsSpecificParams) UnmarshalBinary(b []byte) error {
	var res UdaHostOsSpecificParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
