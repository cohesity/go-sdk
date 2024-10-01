// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NodeNetworkInterfaces Node Network Interfaces.
//
// Specifies the network interfaces present on a specific Node in a Cluster.
//
// swagger:model NodeNetworkInterfaces
type NodeNetworkInterfaces struct {

	// Specifies the serial number of Chassis.
	ChassisSerial *string `json:"chassisSerial,omitempty"`

	// Specifies the list of network interfaces present on this Node.
	Interfaces []*NetworkInterface `json:"interfaces"`

	// Specifies an optional message describing the result of the request
	// pertaining to this Node.
	Message *string `json:"message,omitempty"`

	// Specifies the ID of the Node.
	NodeID *int64 `json:"nodeId,omitempty"`

	// Specifies the IP of the Node.
	NodeIP *string `json:"nodeIp,omitempty"`

	// Specifies the slot number the Node is located in.
	Slot *int64 `json:"slot,omitempty"`
}

// Validate validates this node network interfaces
func (m *NodeNetworkInterfaces) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInterfaces(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeNetworkInterfaces) validateInterfaces(formats strfmt.Registry) error {
	if swag.IsZero(m.Interfaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Interfaces); i++ {
		if swag.IsZero(m.Interfaces[i]) { // not required
			continue
		}

		if m.Interfaces[i] != nil {
			if err := m.Interfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this node network interfaces based on the context it is used
func (m *NodeNetworkInterfaces) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInterfaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeNetworkInterfaces) contextValidateInterfaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Interfaces); i++ {

		if m.Interfaces[i] != nil {

			if swag.IsZero(m.Interfaces[i]) { // not required
				return nil
			}

			if err := m.Interfaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeNetworkInterfaces) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeNetworkInterfaces) UnmarshalBinary(b []byte) error {
	var res NodeNetworkInterfaces
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
