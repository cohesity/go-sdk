// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostgresNodeInfo Postgres Node Information.
//
// Specifies the Node Id, IP and port information to access the postgres
// database.
//
// swagger:model PostgresNodeInfo
type PostgresNodeInfo struct {

	// Specifies the default password to access the postgres database.
	DefaultPassword *string `json:"defaultPassword,omitempty"`

	// Specifies the default username to access the postgres database.
	DefaultUsername *string `json:"defaultUsername,omitempty"`

	// Specifies the id of the node where postgres database is running.
	NodeID *int64 `json:"nodeId,omitempty"`

	// Specifies the ip of the node where postgres database is running.
	NodeIP *string `json:"nodeIp,omitempty"`

	// Specifies the information where postgres database is running.
	Port *int32 `json:"port,omitempty"`
}

// Validate validates this postgres node info
func (m *PostgresNodeInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this postgres node info based on context it is used
func (m *PostgresNodeInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostgresNodeInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostgresNodeInfo) UnmarshalBinary(b []byte) error {
	var res PostgresNodeInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
