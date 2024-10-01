// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NodeResetState NodeResetState
//
// # Node reset state information
//
// swagger:model NodeResetState
type NodeResetState struct {

	// Node Id
	NodeID *int64 `json:"nodeId,omitempty"`

	// Node Ip.
	NodeIP *string `json:"nodeIp,omitempty"`

	// Reset state.
	State *string `json:"state,omitempty"`
}

// Validate validates this node reset state
func (m *NodeResetState) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this node reset state based on context it is used
func (m *NodeResetState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NodeResetState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeResetState) UnmarshalBinary(b []byte) error {
	var res NodeResetState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
