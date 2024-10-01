// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IPConfigParams Specifies the IP config parameters.
//
// swagger:model IPConfigParams
type IPConfigParams struct {

	// Specifies the network interface name. IPs would be assigned to the specified interface.
	Interface *string `json:"interface,omitempty"`

	// Specifies a list of IP addresses to be assigned.
	Ips []string `json:"ips,omitempty"`

	// Specifies the cluster node ids.
	NodeIds []string `json:"nodeIds,omitempty"`

	// Specifies the interface role.
	Role *string `json:"role,omitempty"`

	// Specifies the interface gateway.
	SubnetGateway *string `json:"subnetGateway,omitempty"`

	// Specifies the interface subnet mask bits.
	SubnetMaskBits *int64 `json:"subnetMaskBits,omitempty"`

	// Specifies the IP family of the config.
	IPFamily *int64 `json:"ipFamily,omitempty"`
}

// Validate validates this IP config params
func (m *IPConfigParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this IP config params based on context it is used
func (m *IPConfigParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IPConfigParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPConfigParams) UnmarshalBinary(b []byte) error {
	var res IPConfigParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
