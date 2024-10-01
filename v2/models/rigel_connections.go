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

// RigelConnections Rigel connections.
//
// Specify a list of connection of Rigel.
//
// swagger:model RigelConnections
type RigelConnections struct {

	// Specifies a list of connection of Rigel.
	RigelConnections []*RigelConnection `json:"RigelConnections"`
}

// Validate validates this rigel connections
func (m *RigelConnections) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRigelConnections(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RigelConnections) validateRigelConnections(formats strfmt.Registry) error {
	if swag.IsZero(m.RigelConnections) { // not required
		return nil
	}

	for i := 0; i < len(m.RigelConnections); i++ {
		if swag.IsZero(m.RigelConnections[i]) { // not required
			continue
		}

		if m.RigelConnections[i] != nil {
			if err := m.RigelConnections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RigelConnections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RigelConnections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this rigel connections based on the context it is used
func (m *RigelConnections) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRigelConnections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RigelConnections) contextValidateRigelConnections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RigelConnections); i++ {

		if m.RigelConnections[i] != nil {

			if swag.IsZero(m.RigelConnections[i]) { // not required
				return nil
			}

			if err := m.RigelConnections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RigelConnections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RigelConnections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RigelConnections) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RigelConnections) UnmarshalBinary(b []byte) error {
	var res RigelConnections
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
