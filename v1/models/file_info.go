// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FileInfo Information about a file.
//
// swagger:model FileInfo
type FileInfo struct {

	// Modification time of file/directory
	Mtime *int64 `json:"mtime,omitempty"`

	// Size of file/directory
	Size *int64 `json:"size,omitempty"`
}

// Validate validates this file info
func (m *FileInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this file info based on context it is used
func (m *FileInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FileInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileInfo) UnmarshalBinary(b []byte) error {
	var res FileInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
