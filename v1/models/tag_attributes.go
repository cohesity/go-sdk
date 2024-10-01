// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TagAttributes Message representing a AWS tag so that it can be used as an attribute
// of an AWS entity.
//
// swagger:model TagAttributes
type TagAttributes struct {

	// Entity id corresponding to the tag entity.
	EntityID *int64 `json:"entityId,omitempty"`

	// The name of the tag.
	Name *string `json:"name,omitempty"`

	// The instance UUID of the tag object.
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this tag attributes
func (m *TagAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tag attributes based on context it is used
func (m *TagAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TagAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TagAttributes) UnmarshalBinary(b []byte) error {
	var res TagAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
