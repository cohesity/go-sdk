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

// ProtectedObjectsTile Protected Objects Tile.
//
// Protected Objects information.
//
// swagger:model ProtectedObjectsTile
type ProtectedObjectsTile struct {

	// Protected Objects breakdown by object type.
	ObjectsProtected []*ProtectedObjectsByEnv `json:"objectsProtected"`

	// Number of Protected Objects.
	ProtectedCount *int32 `json:"protectedCount,omitempty"`

	// Size of Protected Objects.
	ProtectedSizeBytes *int64 `json:"protectedSizeBytes,omitempty"`

	// Number of Unprotected Objects.
	UnprotectedCount *int32 `json:"unprotectedCount,omitempty"`

	// Size of Unprotected Objects.
	UnprotectedSizeBytes *int64 `json:"unprotectedSizeBytes,omitempty"`
}

// Validate validates this protected objects tile
func (m *ProtectedObjectsTile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjectsProtected(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtectedObjectsTile) validateObjectsProtected(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectsProtected) { // not required
		return nil
	}

	for i := 0; i < len(m.ObjectsProtected); i++ {
		if swag.IsZero(m.ObjectsProtected[i]) { // not required
			continue
		}

		if m.ObjectsProtected[i] != nil {
			if err := m.ObjectsProtected[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objectsProtected" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objectsProtected" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this protected objects tile based on the context it is used
func (m *ProtectedObjectsTile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjectsProtected(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtectedObjectsTile) contextValidateObjectsProtected(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ObjectsProtected); i++ {

		if m.ObjectsProtected[i] != nil {

			if swag.IsZero(m.ObjectsProtected[i]) { // not required
				return nil
			}

			if err := m.ObjectsProtected[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objectsProtected" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objectsProtected" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProtectedObjectsTile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectedObjectsTile) UnmarshalBinary(b []byte) error {
	var res ProtectedObjectsTile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
