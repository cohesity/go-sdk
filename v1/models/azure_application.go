// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureApplication Message that represents an Azure application.
//
// swagger:model AzureApplication
type AzureApplication struct {

	// Id of the Azure application.
	ApplicationID *string `json:"application_id,omitempty"`

	// Object Id of the Azure application.
	ApplicationObjectID *string `json:"application_object_id,omitempty"`

	// Encrypted application key.
	EncryptedApplicationKey *string `json:"encrypted_application_key,omitempty"`
}

// Validate validates this azure application
func (m *AzureApplication) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure application based on context it is used
func (m *AzureApplication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureApplication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureApplication) UnmarshalBinary(b []byte) error {
	var res AzureApplication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
