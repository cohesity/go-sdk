// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ADDomainIdentity AD domain identity. These fields returned by AD powershell. All field values
// are case insensitive. If you change this message, update $ADDomainIdentity
// in magneto/agents/windows/ad/scripts/ad_entityhierarchy.psm1 to maintain
// JSON compatibility.
//
// swagger:model ADDomainIdentity
type ADDomainIdentity struct {

	// Distinguished Name (Suffix) of the domain. Example: DC=tme,DC=local
	Dn *string `json:"dn,omitempty"`

	// Display Name. Example: tme.local
	Name *string `json:"name,omitempty"`

	// Unique objectGUID for an AD domain. This is a unqiue identifier that
	// can be used to compare domains. It would be in the format without {}
	// Example: 299331af-deb3-4cdd-8e2c-d11cde7112ce
	ObjectGUID *string `json:"objectGuid,omitempty"`

	// DomainSID. Example: S-1-5-21-2329549007-3075497753-1539959636
	Sid *string `json:"sid,omitempty"`
}

// Validate validates this a d domain identity
func (m *ADDomainIdentity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this a d domain identity based on context it is used
func (m *ADDomainIdentity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ADDomainIdentity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ADDomainIdentity) UnmarshalBinary(b []byte) error {
	var res ADDomainIdentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
