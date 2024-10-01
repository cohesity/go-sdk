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

// SfdcRecoverParams sfdc recover params
//
// swagger:model SfdcRecoverParams
type SfdcRecoverParams struct {

	// restore objects
	RestoreObjects []*SfdcRestoreObject `json:"restoreObjects"`
}

// Validate validates this sfdc recover params
func (m *SfdcRecoverParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRestoreObjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SfdcRecoverParams) validateRestoreObjects(formats strfmt.Registry) error {
	if swag.IsZero(m.RestoreObjects) { // not required
		return nil
	}

	for i := 0; i < len(m.RestoreObjects); i++ {
		if swag.IsZero(m.RestoreObjects[i]) { // not required
			continue
		}

		if m.RestoreObjects[i] != nil {
			if err := m.RestoreObjects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("restoreObjects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("restoreObjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this sfdc recover params based on the context it is used
func (m *SfdcRecoverParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRestoreObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SfdcRecoverParams) contextValidateRestoreObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RestoreObjects); i++ {

		if m.RestoreObjects[i] != nil {

			if swag.IsZero(m.RestoreObjects[i]) { // not required
				return nil
			}

			if err := m.RestoreObjects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("restoreObjects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("restoreObjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SfdcRecoverParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SfdcRecoverParams) UnmarshalBinary(b []byte) error {
	var res SfdcRecoverParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
