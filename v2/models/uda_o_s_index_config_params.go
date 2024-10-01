// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UdaOSIndexConfigParams Specifies the basic configuration for the operating system type.
//
// swagger:model UdaOSIndexConfigParams
type UdaOSIndexConfigParams struct {

	// UI feature flag that controls the visibility of the corresponding operating system type for the UDA connector.
	UIFeatureFlag *string `json:"uiFeatureFlag,omitempty"`
}

// Validate validates this uda o s index config params
func (m *UdaOSIndexConfigParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this uda o s index config params based on context it is used
func (m *UdaOSIndexConfigParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UdaOSIndexConfigParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UdaOSIndexConfigParams) UnmarshalBinary(b []byte) error {
	var res UdaOSIndexConfigParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
