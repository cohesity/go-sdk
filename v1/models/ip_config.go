// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IPConfig Config Ip Parameters.
//
// Specifies the configuration of an IP.
//
// swagger:model IpConfig
type IPConfig struct {

	// The interface name.
	//
	// Specifies which interface to assign IP to.
	InterfaceName *string `json:"interfaceName,omitempty"`

	// IpFamily of this config.
	IPFamily *int32 `json:"ipFamily,omitempty"`

	// The interface ips.
	Ips []string `json:"ips"`

	// Node ids.
	NodeIds []int64 `json:"nodeIds"`

	// The interface role.
	Role *string `json:"role,omitempty"`

	// The interface gateway.
	SubnetGateway *string `json:"subnetGateway,omitempty"`

	// The interface subnet mask bits.
	SubnetMaskBits *int32 `json:"subnetMaskBits,omitempty"`
}

// Validate validates this Ip config
func (m *IPConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Ip config based on context it is used
func (m *IPConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IPConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPConfig) UnmarshalBinary(b []byte) error {
	var res IPConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
