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

// HeliosCdpBackupPolicy Continious Data Protection (CDP) Policy.
//
// Specifies CDP (Continious Data Protection) backup settings for a Protection Group.
//
// swagger:model HeliosCdpBackupPolicy
type HeliosCdpBackupPolicy struct {

	// Specifies the Retention period of a log backup in days, months or years.
	Retention *HeliosCdpRetention `json:"retention,omitempty"`
}

// Validate validates this helios cdp backup policy
func (m *HeliosCdpBackupPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRetention(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HeliosCdpBackupPolicy) validateRetention(formats strfmt.Registry) error {
	if swag.IsZero(m.Retention) { // not required
		return nil
	}

	if m.Retention != nil {
		if err := m.Retention.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retention")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("retention")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this helios cdp backup policy based on the context it is used
func (m *HeliosCdpBackupPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRetention(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HeliosCdpBackupPolicy) contextValidateRetention(ctx context.Context, formats strfmt.Registry) error {

	if m.Retention != nil {

		if swag.IsZero(m.Retention) { // not required
			return nil
		}

		if err := m.Retention.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retention")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("retention")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HeliosCdpBackupPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HeliosCdpBackupPolicy) UnmarshalBinary(b []byte) error {
	var res HeliosCdpBackupPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
