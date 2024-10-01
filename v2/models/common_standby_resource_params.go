// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonStandbyResourceParams Standby resoure Params.
//
// Specifies the params for standby resource.
//
// swagger:model CommonStandbyResourceParams
type CommonStandbyResourceParams struct {

	// Specifies the recovery point objective time user expects for this standby resource.
	RecoveryPointObjectiveSecs *int64 `json:"recoveryPointObjectiveSecs,omitempty"`
}

// Validate validates this common standby resource params
func (m *CommonStandbyResourceParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this common standby resource params based on context it is used
func (m *CommonStandbyResourceParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonStandbyResourceParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonStandbyResourceParams) UnmarshalBinary(b []byte) error {
	var res CommonStandbyResourceParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
