// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateSupportUserParams Specifies the support user params.
//
// swagger:model UpdateSupportUserParams
type UpdateSupportUserParams struct {

	// Specifies the current password of the user. This is required when trying to update the current user's password.
	CurrentPassword *string `json:"currentPassword,omitempty"`

	// Specifies the new password for the support user.
	NewPassword *string `json:"newPassword,omitempty"`

	// If set to true, sudo access will be enabled for the user. If null, the endpoint will not attempt to alter sudo access privilege for the support user.
	EnableSudoAccess *bool `json:"enableSudoAccess,omitempty"`
}

// Validate validates this update support user params
func (m *UpdateSupportUserParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update support user params based on context it is used
func (m *UpdateSupportUserParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateSupportUserParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateSupportUserParams) UnmarshalBinary(b []byte) error {
	var res UpdateSupportUserParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
