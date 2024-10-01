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

// ExternalClientSubnets Specifies the external Client Subnets that can communicate with this cluster.
//
// swagger:model ExternalClientSubnets
type ExternalClientSubnets struct {

	// Specifies the Client Subnets for the cluster.
	ClientSubnets []*Subnet `json:"clientSubnets"`
}

// Validate validates this external client subnets
func (m *ExternalClientSubnets) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientSubnets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalClientSubnets) validateClientSubnets(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientSubnets) { // not required
		return nil
	}

	for i := 0; i < len(m.ClientSubnets); i++ {
		if swag.IsZero(m.ClientSubnets[i]) { // not required
			continue
		}

		if m.ClientSubnets[i] != nil {
			if err := m.ClientSubnets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clientSubnets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clientSubnets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this external client subnets based on the context it is used
func (m *ExternalClientSubnets) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClientSubnets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalClientSubnets) contextValidateClientSubnets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClientSubnets); i++ {

		if m.ClientSubnets[i] != nil {

			if swag.IsZero(m.ClientSubnets[i]) { // not required
				return nil
			}

			if err := m.ClientSubnets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clientSubnets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clientSubnets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExternalClientSubnets) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalClientSubnets) UnmarshalBinary(b []byte) error {
	var res ExternalClientSubnets
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
