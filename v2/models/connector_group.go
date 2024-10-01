// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConnectorGroup Connector Group
//
// Specify a group of connectors.
//
// swagger:model ConnectorGroup
type ConnectorGroup struct {

	// Specifies the id of the group.
	ID *int64 `json:"id,omitempty"`

	// Specifies the name of the group.
	Name *string `json:"name,omitempty"`

	// Specifies the ids of the connectors in the group.
	Connectors []int64 `json:"connectors"`
}

// Validate validates this connector group
func (m *ConnectorGroup) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this connector group based on context it is used
func (m *ConnectorGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConnectorGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectorGroup) UnmarshalBinary(b []byte) error {
	var res ConnectorGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
