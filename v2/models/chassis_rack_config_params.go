// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChassisRackConfigParams Specifies the chassis serial to rack id mapping configuration.
//
// swagger:model ChassisRackConfigParams
type ChassisRackConfigParams struct {

	// Specifies the chassis serial.
	ChassisSerial string `json:"chassisSerial,omitempty"`

	// Specifies the rack id.
	RackID int64 `json:"rackId,omitempty"`
}

// Validate validates this chassis rack config params
func (m *ChassisRackConfigParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this chassis rack config params based on context it is used
func (m *ChassisRackConfigParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChassisRackConfigParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChassisRackConfigParams) UnmarshalBinary(b []byte) error {
	var res ChassisRackConfigParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
