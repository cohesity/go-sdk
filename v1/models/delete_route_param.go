// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteRouteParam Delete Route Parameters.
//
// Specifies the delete param of a Static Route.
//
// swagger:model DeleteRouteParam
type DeleteRouteParam struct {

	// Destination network.
	//
	// Specifies the destination network of the Static Route.
	// overrideDescription: true
	DestNetwork *string `json:"destNetwork,omitempty"`

	// Specifies the network interfaces name to use for communicating with the
	// destination network.
	IfName *string `json:"ifName,omitempty"`

	// Specifies the network interfaces group or vlan interface group to
	// use for communicating with the destination network.
	IfaceGroupName *string `json:"ifaceGroupName,omitempty"`

	// Specifies the network node group to represent a group of nodes.
	NodeGroupName *string `json:"nodeGroupName,omitempty"`
}

// Validate validates this delete route param
func (m *DeleteRouteParam) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete route param based on context it is used
func (m *DeleteRouteParam) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteRouteParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteRouteParam) UnmarshalBinary(b []byte) error {
	var res DeleteRouteParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
