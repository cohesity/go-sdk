// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TotpKeyInfo Specifies the Time based OTP key info
//
// swagger:model TotpKeyInfo
type TotpKeyInfo struct {

	// Specifies the Time based OTP key name.
	TotpKeyName string `json:"totpKeyName,omitempty"`

	// Specifies the Time based OTP key ID.
	TotpKeyID string `json:"totpKeyId,omitempty"`

	// Specifies the Time based OTP secret key.
	TotpSecretKey string `json:"totpSecretKey,omitempty"`

	// Specifies the Time based OTP key URI for generating MFA QR code.
	TotpURI string `json:"totpUri,omitempty"`
}

// Validate validates this totp key info
func (m *TotpKeyInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this totp key info based on context it is used
func (m *TotpKeyInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TotpKeyInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TotpKeyInfo) UnmarshalBinary(b []byte) error {
	var res TotpKeyInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
