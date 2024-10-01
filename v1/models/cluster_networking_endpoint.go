// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterNetworkingEndpoint Cluster Networking Endpoint.
//
// Specifies information about end point.
//
// swagger:model ClusterNetworkingEndpoint
type ClusterNetworkingEndpoint struct {

	// The Fully Qualified Domain Name.
	Fqdn *string `json:"fqdn,omitempty"`

	// The IPv4 address.
	IPV4Addr *string `json:"ipv4Addr,omitempty"`

	// The IPv6 address.
	IPV6Addr *string `json:"ipv6Addr,omitempty"`
}

// Validate validates this cluster networking endpoint
func (m *ClusterNetworkingEndpoint) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster networking endpoint based on context it is used
func (m *ClusterNetworkingEndpoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNetworkingEndpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNetworkingEndpoint) UnmarshalBinary(b []byte) error {
	var res ClusterNetworkingEndpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
