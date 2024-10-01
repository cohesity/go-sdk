// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SiteProperty Generic site Property structure to store site properties. Some of these
// site properties are not captured by Get-PnpProvisioningTemplate cmdlet. So
// they may need to be set outside of Appy-PnPProvisoningTemplate cmdlet.
//
// swagger:model SiteProperty
type SiteProperty struct {

	// PnP data type of the property ('string', 'mvstring', 'bool',
	// 'guid','<enumname>','int', 'int64', etc.)
	Datatype *string `json:"datatype,omitempty"`

	// Name of the property returned by Get-PnPSite,
	// Get-PnPTenantSite, Get-PnPWeb, or other cmdlets.
	Name *string `json:"name,omitempty"`

	// Property value represented as string.
	Value *string `json:"value,omitempty"`
}

// Validate validates this site property
func (m *SiteProperty) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this site property based on context it is used
func (m *SiteProperty) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SiteProperty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteProperty) UnmarshalBinary(b []byte) error {
	var res SiteProperty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
