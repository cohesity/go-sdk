// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PrivateKeyCredentials Private key credentials
//
// Specifies the credentials for certificate based authentication.
//
// swagger:model PrivateKeyCredentials
type PrivateKeyCredentials struct {

	// Specifies the userId to access target entity.
	// Required: true
	UserID *string `json:"userId"`

	// Specifies the private key to access target entity.
	// Required: true
	PrivateKey *string `json:"privateKey"`

	// Specifies the passphrase to access target entity.
	Passphrase string `json:"passphrase,omitempty"`
}

// Validate validates this private key credentials
func (m *PrivateKeyCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateKeyCredentials) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

func (m *PrivateKeyCredentials) validatePrivateKey(formats strfmt.Registry) error {

	if err := validate.Required("privateKey", "body", m.PrivateKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this private key credentials based on context it is used
func (m *PrivateKeyCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrivateKeyCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateKeyCredentials) UnmarshalBinary(b []byte) error {
	var res PrivateKeyCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
