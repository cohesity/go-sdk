// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SharepointBackupSourceParams Message defining the object-level params for Sharepoint site backup.
//
// swagger:model SharepointBackupSourceParams
type SharepointBackupSourceParams struct {

	// Specifies to whether autoprotect the site entity  or not.
	// If this is set to true, the site entity and subsites of it are protected.
	// If this is set to false, only the site entity is protected.
	AutoprotectEntity *bool `json:"autoprotectEntity,omitempty"`
}

// Validate validates this sharepoint backup source params
func (m *SharepointBackupSourceParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sharepoint backup source params based on context it is used
func (m *SharepointBackupSourceParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SharepointBackupSourceParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SharepointBackupSourceParams) UnmarshalBinary(b []byte) error {
	var res SharepointBackupSourceParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
