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

// SfdcRestoreObject sfdc restore object
//
// swagger:model SfdcRestoreObject
type SfdcRestoreObject struct {

	// Magneto Entity id.
	// If this is set, object_name/object_uuid should be empty.
	EntityID *int64 `json:"entityId,omitempty"`

	// The original name of the object to be restored.
	ObjectName *string `json:"objectName,omitempty"`

	// The specific params related to the object to be restored.
	RestoreParams *SfdcRestoreObjectParams `json:"restoreParams,omitempty"`
}

// Validate validates this sfdc restore object
func (m *SfdcRestoreObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRestoreParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SfdcRestoreObject) validateRestoreParams(formats strfmt.Registry) error {
	if swag.IsZero(m.RestoreParams) { // not required
		return nil
	}

	if m.RestoreParams != nil {
		if err := m.RestoreParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restoreParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("restoreParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sfdc restore object based on the context it is used
func (m *SfdcRestoreObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRestoreParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SfdcRestoreObject) contextValidateRestoreParams(ctx context.Context, formats strfmt.Registry) error {

	if m.RestoreParams != nil {

		if swag.IsZero(m.RestoreParams) { // not required
			return nil
		}

		if err := m.RestoreParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restoreParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("restoreParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SfdcRestoreObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SfdcRestoreObject) UnmarshalBinary(b []byte) error {
	var res SfdcRestoreObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
