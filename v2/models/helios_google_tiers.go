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
	"github.com/go-openapi/validate"
)

// HeliosGoogleTiers Specifies Google tiers.
//
// swagger:model HeliosGoogleTiers
type HeliosGoogleTiers struct {

	// Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
	// Required: true
	Tiers []*HeliosGoogleTier `json:"tiers"`
}

// Validate validates this helios google tiers
func (m *HeliosGoogleTiers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTiers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HeliosGoogleTiers) validateTiers(formats strfmt.Registry) error {

	if err := validate.Required("tiers", "body", m.Tiers); err != nil {
		return err
	}

	for i := 0; i < len(m.Tiers); i++ {
		if swag.IsZero(m.Tiers[i]) { // not required
			continue
		}

		if m.Tiers[i] != nil {
			if err := m.Tiers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this helios google tiers based on the context it is used
func (m *HeliosGoogleTiers) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTiers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HeliosGoogleTiers) contextValidateTiers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tiers); i++ {

		if m.Tiers[i] != nil {

			if swag.IsZero(m.Tiers[i]) { // not required
				return nil
			}

			if err := m.Tiers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HeliosGoogleTiers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HeliosGoogleTiers) UnmarshalBinary(b []byte) error {
	var res HeliosGoogleTiers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
