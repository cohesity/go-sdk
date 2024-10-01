// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeviceTreeNode Specifies the tree structure of a logical volume. The leaves are slices of partitions and the other nodes are assemled by combining nodes in some mode.
//
// swagger:model DeviceTreeNode
type DeviceTreeNode struct {

	// Specifies if the node is a leaf node.
	IsLeaf *bool `json:"isLeaf,omitempty"`

	// Specifies the parameters for a leaf node.
	LeafNodeParams *DeviceTreeLeafNode `json:"leafNodeParams,omitempty"`

	// Specifies the parameters for a non leaf node.
	NonLeafNodeParams *DeviceTreeNonLeafNode `json:"nonLeafNodeParams,omitempty"`
}

// Validate validates this device tree node
func (m *DeviceTreeNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLeafNodeParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNonLeafNodeParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceTreeNode) validateLeafNodeParams(formats strfmt.Registry) error {
	if swag.IsZero(m.LeafNodeParams) { // not required
		return nil
	}

	if m.LeafNodeParams != nil {
		if err := m.LeafNodeParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("leafNodeParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("leafNodeParams")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceTreeNode) validateNonLeafNodeParams(formats strfmt.Registry) error {
	if swag.IsZero(m.NonLeafNodeParams) { // not required
		return nil
	}

	if m.NonLeafNodeParams != nil {
		if err := m.NonLeafNodeParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nonLeafNodeParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nonLeafNodeParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this device tree node based on the context it is used
func (m *DeviceTreeNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLeafNodeParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNonLeafNodeParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceTreeNode) contextValidateLeafNodeParams(ctx context.Context, formats strfmt.Registry) error {

	if m.LeafNodeParams != nil {

		if swag.IsZero(m.LeafNodeParams) { // not required
			return nil
		}

		if err := m.LeafNodeParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("leafNodeParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("leafNodeParams")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceTreeNode) contextValidateNonLeafNodeParams(ctx context.Context, formats strfmt.Registry) error {

	if m.NonLeafNodeParams != nil {

		if swag.IsZero(m.NonLeafNodeParams) { // not required
			return nil
		}

		if err := m.NonLeafNodeParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nonLeafNodeParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nonLeafNodeParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceTreeNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceTreeNode) UnmarshalBinary(b []byte) error {
	var res DeviceTreeNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
