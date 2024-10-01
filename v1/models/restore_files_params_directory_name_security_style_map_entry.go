// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestoreFilesParamsDirectoryNameSecurityStyleMapEntry restore files params directory name security style map entry
//
// swagger:model RestoreFilesParams_DirectoryNameSecurityStyleMapEntry
type RestoreFilesParamsDirectoryNameSecurityStyleMapEntry struct {

	// key
	Key *string `json:"key,omitempty"`

	// value
	Value *string `json:"value,omitempty"`
}

// Validate validates this restore files params directory name security style map entry
func (m *RestoreFilesParamsDirectoryNameSecurityStyleMapEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this restore files params directory name security style map entry based on context it is used
func (m *RestoreFilesParamsDirectoryNameSecurityStyleMapEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestoreFilesParamsDirectoryNameSecurityStyleMapEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreFilesParamsDirectoryNameSecurityStyleMapEntry) UnmarshalBinary(b []byte) error {
	var res RestoreFilesParamsDirectoryNameSecurityStyleMapEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
