// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomAttribute Message representing a VMware Custom attribute so that it can be used as an
// attribute of a VMware entity.
//
// swagger:model CustomAttribute
type CustomAttribute struct {

	// The key of the custom attribute.
	Key *int32 `json:"key,omitempty"`

	// The name of the custom attribute. Name is case sensitive.
	// eg. "Last Backup"
	Name *string `json:"name,omitempty"`

	// The ManagedObjectType of the custom attribute.
	Type *string `json:"type,omitempty"`

	// The value of the custom attribute.
	// eg. "10-25-2019 10:00 am"
	Value *string `json:"value,omitempty"`
}

// Validate validates this custom attribute
func (m *CustomAttribute) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this custom attribute based on context it is used
func (m *CustomAttribute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomAttribute) UnmarshalBinary(b []byte) error {
	var res CustomAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
