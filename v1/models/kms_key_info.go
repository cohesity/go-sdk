// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KmsKeyInfo Message representing information related to kms key.
//
// swagger:model KmsKeyInfo
type KmsKeyInfo struct {

	// The kms key alias.
	Alias *string `json:"alias,omitempty"`

	// The kms key arn.
	Arn *string `json:"arn,omitempty"`

	// Whether the kms key is enabled or not.
	Enabled *bool `json:"enabled,omitempty"`
}

// Validate validates this kms key info
func (m *KmsKeyInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kms key info based on context it is used
func (m *KmsKeyInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KmsKeyInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KmsKeyInfo) UnmarshalBinary(b []byte) error {
	var res KmsKeyInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
