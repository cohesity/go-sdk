// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClearNlmLocksParameters Clear NLM Locks Parameters.
//
// Specifies parameters required to force clear NLM Locks.
//
// swagger:model ClearNlmLocksParameters
type ClearNlmLocksParameters struct {

	// Specifies the id of the client, related NLM locks needs to be clear.
	ClientID *string `json:"clientId,omitempty"`

	// Specifies the filepath in the view relative to the root filesystem.
	// If this field is specified, viewName field must also be specified.
	FilePath *string `json:"filePath,omitempty"`

	// Specifies the name of the View in which to search. If a view name is not
	// specified, all the views in the Cluster is searched.
	// This field is mandatory if filePath field is specified.
	ViewName *string `json:"viewName,omitempty"`
}

// Validate validates this clear nlm locks parameters
func (m *ClearNlmLocksParameters) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this clear nlm locks parameters based on context it is used
func (m *ClearNlmLocksParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClearNlmLocksParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClearNlmLocksParameters) UnmarshalBinary(b []byte) error {
	var res ClearNlmLocksParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
