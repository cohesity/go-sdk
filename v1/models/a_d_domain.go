// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ADDomain Represents an AD domain. All fields are case insensitive. If you change this
// message, update $ADDomain in
// magneto/agents/windows/ad/scripts/ad_entityhierarchy.psm1 to maintain JSON
// compatibility.
//
// swagger:model ADDomain
type ADDomain struct {

	// DNS root. Example: tme.local
	DNSRoot *string `json:"dnsRoot,omitempty"`

	// Domain functional level numeric. (Get-ADDomain).DomainMode
	// Windows2012R2Domain
	DomainMode *int32 `json:"domainMode,omitempty"`

	// Forest. Example: qa01.eng.cohesity.com
	Forest *string `json:"forest,omitempty"`

	// Forest functional level. (Get-ADForest).ForestMode
	// Windows2012R2Forest
	ForestMode *int32 `json:"forestMode,omitempty"`

	// Identity of the domain.
	ID *ADDomainIdentity `json:"id,omitempty"`

	// NetBIOS name. Example: TME
	NetbiosName *string `json:"netbiosName,omitempty"`

	// Parent domain name.
	ParentDomain *string `json:"parentDomain,omitempty"`

	// AD Schema version obtained from (Get-ADObject
	// (Get-ADRootDSE).schemaNamingContext -Property objectVersion).
	SchemaVersion *uint32 `json:"schemaVersion,omitempty"`

	// Tombstone lifetime in days. Default is 180 days.
	TombstoneDays *uint32 `json:"tombstoneDays,omitempty"`
}

// Validate validates this a d domain
func (m *ADDomain) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ADDomain) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if m.ID != nil {
		if err := m.ID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("id")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this a d domain based on the context it is used
func (m *ADDomain) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ADDomain) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if m.ID != nil {

		if swag.IsZero(m.ID) { // not required
			return nil
		}

		if err := m.ID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("id")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ADDomain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ADDomain) UnmarshalBinary(b []byte) error {
	var res ADDomain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
