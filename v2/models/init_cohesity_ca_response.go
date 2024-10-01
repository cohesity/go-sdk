// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InitCohesityCaResponse Specifies the response to init cohesity ca request
//
// swagger:model InitCohesityCaResponse
type InitCohesityCaResponse struct {

	// Specifies ca cert in pem format
	CaCertChain []string `json:"caCertChain,omitempty"`
}

// Validate validates this init cohesity ca response
func (m *InitCohesityCaResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this init cohesity ca response based on context it is used
func (m *InitCohesityCaResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InitCohesityCaResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InitCohesityCaResponse) UnmarshalBinary(b []byte) error {
	var res InitCohesityCaResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
