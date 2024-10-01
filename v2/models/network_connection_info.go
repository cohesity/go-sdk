// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NetworkConnectionInfo Network Connection.
//
// Specify the network connection information.
//
// swagger:model NetworkConnectionInfo
type NetworkConnectionInfo struct {

	// Specifies the domain name of the network connection.
	DomainName *string `json:"domainName,omitempty"`

	// Specifies the network Gateway of the network connection.
	NetworkGateway *string `json:"networkGateway,omitempty"`

	// Specifies the DNS Server of the network connection.
	DNS *string `json:"dns,omitempty"`

	// Specifies the NTP Server of the network connection.
	Ntp *string `json:"ntp,omitempty"`
}

// Validate validates this network connection info
func (m *NetworkConnectionInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network connection info based on context it is used
func (m *NetworkConnectionInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkConnectionInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkConnectionInfo) UnmarshalBinary(b []byte) error {
	var res NetworkConnectionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
