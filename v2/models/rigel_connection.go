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

// RigelConnection Rigel connection.
//
// Specify a connection of Rigel.
//
// swagger:model RigelConnection
type RigelConnection struct {

	// Specifies the id of the connection.
	ID *int64 `json:"id,omitempty"`

	// Specifies the name of the connection.
	Name *string `json:"name,omitempty"`

	// Flag to specify if the connection is scalable.
	Scalable *bool `json:"scalable,omitempty"`

	// Specifies the ids of the connectors which are not grouped in this connection
	UngroupedConnectors []int64 `json:"ungroupedConnectors"`

	// Specifies the connector groups in the connection.
	ConnectorGroups []*ConnectorGroup `json:"connectorGroups"`
}

// Validate validates this rigel connection
func (m *RigelConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectorGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RigelConnection) validateConnectorGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectorGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.ConnectorGroups); i++ {
		if swag.IsZero(m.ConnectorGroups[i]) { // not required
			continue
		}

		if m.ConnectorGroups[i] != nil {
			if err := m.ConnectorGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connectorGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connectorGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this rigel connection based on the context it is used
func (m *RigelConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnectorGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RigelConnection) contextValidateConnectorGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConnectorGroups); i++ {

		if m.ConnectorGroups[i] != nil {

			if swag.IsZero(m.ConnectorGroups[i]) { // not required
				return nil
			}

			if err := m.ConnectorGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connectorGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connectorGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RigelConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RigelConnection) UnmarshalBinary(b []byte) error {
	var res RigelConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
