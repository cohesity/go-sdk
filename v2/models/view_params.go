// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ViewParams View
//
// Specifies the details about a view.
//
// swagger:model ViewParams
type ViewParams struct {

	// Specifies the ID of the view.
	ID *string `json:"id,omitempty"`

	// Specifies the name of the view.
	Name *string `json:"name,omitempty"`
}

// Validate validates this view params
func (m *ViewParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this view params based on context it is used
func (m *ViewParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ViewParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ViewParams) UnmarshalBinary(b []byte) error {
	var res ViewParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
