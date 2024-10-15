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

// BackupTaskStateProtoBackupViewInfo Message containing backup view information about a single entity.
//
// swagger:model BackupTaskStateProto_BackupViewInfo
type BackupTaskStateProtoBackupViewInfo struct {

	// Entity id proto.
	EntityID *EntityIDProto `json:"entityId,omitempty"`

	// Name of the view where the data was copied to while backing up the
	// entity referred to by entity_id above.
	ViewName *string `json:"viewName,omitempty"`
}

// Validate validates this backup task state proto backup view info
func (m *BackupTaskStateProtoBackupViewInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupTaskStateProtoBackupViewInfo) validateEntityID(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityID) { // not required
		return nil
	}

	if m.EntityID != nil {
		if err := m.EntityID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityId")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this backup task state proto backup view info based on the context it is used
func (m *BackupTaskStateProtoBackupViewInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntityID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupTaskStateProtoBackupViewInfo) contextValidateEntityID(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityID != nil {

		if swag.IsZero(m.EntityID) { // not required
			return nil
		}

		if err := m.EntityID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityId")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupTaskStateProtoBackupViewInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupTaskStateProtoBackupViewInfo) UnmarshalBinary(b []byte) error {
	var res BackupTaskStateProtoBackupViewInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
