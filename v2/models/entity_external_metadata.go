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

// EntityExternalMetadata Specifies the External metadata of an entity
//
// swagger:model EntityExternalMetadata
type EntityExternalMetadata struct {

	// Specifies the entity metadata for maintenance mode.
	MaintenanceModeConfig *MaintenanceModeConfig `json:"maintenanceModeConfig,omitempty"`
}

// Validate validates this entity external metadata
func (m *EntityExternalMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaintenanceModeConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntityExternalMetadata) validateMaintenanceModeConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.MaintenanceModeConfig) { // not required
		return nil
	}

	if m.MaintenanceModeConfig != nil {
		if err := m.MaintenanceModeConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maintenanceModeConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("maintenanceModeConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this entity external metadata based on the context it is used
func (m *EntityExternalMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMaintenanceModeConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntityExternalMetadata) contextValidateMaintenanceModeConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.MaintenanceModeConfig != nil {

		if swag.IsZero(m.MaintenanceModeConfig) { // not required
			return nil
		}

		if err := m.MaintenanceModeConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maintenanceModeConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("maintenanceModeConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EntityExternalMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntityExternalMetadata) UnmarshalBinary(b []byte) error {
	var res EntityExternalMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
