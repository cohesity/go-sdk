// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SmbLeaseInfo SmbLeaseInfo specifies the SMB lease info of the open item.
//
// swagger:model SmbLeaseInfo
type SmbLeaseInfo struct {

	// Whether lease break is in progress
	IsBreaking *bool `json:"isBreaking,omitempty"`

	// Lease type that is attempted to being broken.
	LeaseBreakType []string `json:"leaseBreakType"`

	// Lease type granted for open.
	LeaseTye []string `json:"leaseTye"`
}

// Validate validates this smb lease info
func (m *SmbLeaseInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this smb lease info based on context it is used
func (m *SmbLeaseInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SmbLeaseInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmbLeaseInfo) UnmarshalBinary(b []byte) error {
	var res SmbLeaseInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
