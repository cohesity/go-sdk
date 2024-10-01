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

// InterfaceGroups Interface groups configured on the cluster.
//
// swagger:model InterfaceGroups
type InterfaceGroups struct {

	// Interface groups configured on the cluster.
	InterfaceGroups []*InterfaceGroup `json:"interfaceGroups"`
}

// Validate validates this interface groups
func (m *InterfaceGroups) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInterfaceGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfaceGroups) validateInterfaceGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.InterfaceGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.InterfaceGroups); i++ {
		if swag.IsZero(m.InterfaceGroups[i]) { // not required
			continue
		}

		if m.InterfaceGroups[i] != nil {
			if err := m.InterfaceGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaceGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("interfaceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this interface groups based on the context it is used
func (m *InterfaceGroups) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInterfaceGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfaceGroups) contextValidateInterfaceGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InterfaceGroups); i++ {

		if m.InterfaceGroups[i] != nil {

			if swag.IsZero(m.InterfaceGroups[i]) { // not required
				return nil
			}

			if err := m.InterfaceGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaceGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("interfaceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceGroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceGroups) UnmarshalBinary(b []byte) error {
	var res InterfaceGroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
