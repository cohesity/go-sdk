// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DNSDelegationZone Dns delegation zone
//
// swagger:model DnsDelegationZone
type DNSDelegationZone struct {

	// Specifies the dns zone name.
	DNSZoneName *string `json:"dnsZoneName,omitempty"`

	// Specifies list of vips that will be resolved to.
	DNSZoneResolvedVips []string `json:"dnsZoneResolvedVips"`

	// Specifies list of vips part of dns delegation zone.
	DNSZoneVips []string `json:"dnsZoneVips"`
}

// Validate validates this Dns delegation zone
func (m *DNSDelegationZone) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Dns delegation zone based on context it is used
func (m *DNSDelegationZone) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DNSDelegationZone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DNSDelegationZone) UnmarshalBinary(b []byte) error {
	var res DNSDelegationZone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
