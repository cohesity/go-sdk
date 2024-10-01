// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NodeIpmiUser Node level IPMI user parameters.
//
// swagger:model NodeIpmiUser
type NodeIpmiUser struct {

	// IPMI username.
	// Required: true
	Username *string `json:"username"`

	// Node id.
	// Required: true
	NodeID *int64 `json:"nodeId"`
}

// Validate validates this node ipmi user
func (m *NodeIpmiUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeIpmiUser) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

func (m *NodeIpmiUser) validateNodeID(formats strfmt.Registry) error {

	if err := validate.Required("nodeId", "body", m.NodeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this node ipmi user based on context it is used
func (m *NodeIpmiUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NodeIpmiUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeIpmiUser) UnmarshalBinary(b []byte) error {
	var res NodeIpmiUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
