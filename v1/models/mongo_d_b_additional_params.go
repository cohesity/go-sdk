// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MongoDBAdditionalParams Contains additional  parameters required for taking backup from
// this Mongo cluster.
//
// swagger:model MongoDBAdditionalParams
type MongoDBAdditionalParams struct {

	// The tag associated with the secondary nodes from which backups
	// should be performed.
	SecondaryNodeTag []string `json:"secondaryNodeTag"`

	// Set to true if this cluster uses secondary nodes for backup.
	UseSecondaryForBackup *bool `json:"useSecondaryForBackup,omitempty"`
}

// Validate validates this mongo d b additional params
func (m *MongoDBAdditionalParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this mongo d b additional params based on context it is used
func (m *MongoDBAdditionalParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MongoDBAdditionalParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MongoDBAdditionalParams) UnmarshalBinary(b []byte) error {
	var res MongoDBAdditionalParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
