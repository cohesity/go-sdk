// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PasswordFormFieldParams Parameters to specify a password form field.
//
// swagger:model PasswordFormFieldParams
type PasswordFormFieldParams struct {

	// Specifies whether the field is mandatory.
	Required *bool `json:"required,omitempty"`

	// Description for the field to be shown on UI screen.
	Description *string `json:"description,omitempty"`

	// Placeholder value for the form field.
	Placeholder *string `json:"placeholder,omitempty"`
}

// Validate validates this password form field params
func (m *PasswordFormFieldParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this password form field params based on context it is used
func (m *PasswordFormFieldParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PasswordFormFieldParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PasswordFormFieldParams) UnmarshalBinary(b []byte) error {
	var res PasswordFormFieldParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
