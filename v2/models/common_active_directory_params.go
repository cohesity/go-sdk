// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CommonActiveDirectoryParams Specifies the params of Active Directory which are used across creating and updating.
//
// swagger:model CommonActiveDirectoryParams
type CommonActiveDirectoryParams struct {

	// Specifies a list of computer names used to identify the Cohesity Cluster on the Active Directory domain. The first machine account is used as primary machine account and it can not be modified.
	// Required: true
	// Min Items: 1
	// Unique: true
	MachineAccounts []*MachineAccount `json:"machineAccounts"`

	// Specifies the id of the Active Directory.
	// Read Only: true
	ID *int64 `json:"id,omitempty"`

	// Specifies an optional organizational unit name.
	OrganizationalUnitName *string `json:"organizationalUnitName,omitempty"`

	// Specifies a work group name.
	WorkGroupName *string `json:"workGroupName,omitempty"`

	// Specifies a list of preferred domain controllers of this Active Directory.
	PreferredDomainControllers []*DomainController `json:"preferredDomainControllers"`

	// Specifies a list of denied domain controllers of this Active Directory Domain.
	DomainControllersDenyList []*string `json:"domainControllersDenyList"`

	// Specifies the LDAP provider id which is mapped to this Active Directory
	LdapProviderID *int64 `json:"ldapProviderId,omitempty"`

	// Specifies the name of the NIS Provider which is mapped to this Active Directory.
	NisProviderDomainName *string `json:"nisProviderDomainName,omitempty"`

	// Specifies the id of the connection.
	ConnectionID *int64 `json:"connectionId,omitempty"`

	// Specifies the params of trusted domain info of an Active Directory.
	TrustedDomainParams *TrustedDomainParams `json:"trustedDomainParams,omitempty"`
}

// Validate validates this common active directory params
func (m *CommonActiveDirectoryParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMachineAccounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreferredDomainControllers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrustedDomainParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonActiveDirectoryParams) validateMachineAccounts(formats strfmt.Registry) error {

	if err := validate.Required("machineAccounts", "body", m.MachineAccounts); err != nil {
		return err
	}

	iMachineAccountsSize := int64(len(m.MachineAccounts))

	if err := validate.MinItems("machineAccounts", "body", iMachineAccountsSize, 1); err != nil {
		return err
	}

	if err := validate.UniqueItems("machineAccounts", "body", m.MachineAccounts); err != nil {
		return err
	}

	for i := 0; i < len(m.MachineAccounts); i++ {
		if swag.IsZero(m.MachineAccounts[i]) { // not required
			continue
		}

		if m.MachineAccounts[i] != nil {
			if err := m.MachineAccounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("machineAccounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("machineAccounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonActiveDirectoryParams) validatePreferredDomainControllers(formats strfmt.Registry) error {
	if swag.IsZero(m.PreferredDomainControllers) { // not required
		return nil
	}

	for i := 0; i < len(m.PreferredDomainControllers); i++ {
		if swag.IsZero(m.PreferredDomainControllers[i]) { // not required
			continue
		}

		if m.PreferredDomainControllers[i] != nil {
			if err := m.PreferredDomainControllers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("preferredDomainControllers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("preferredDomainControllers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonActiveDirectoryParams) validateTrustedDomainParams(formats strfmt.Registry) error {
	if swag.IsZero(m.TrustedDomainParams) { // not required
		return nil
	}

	if m.TrustedDomainParams != nil {
		if err := m.TrustedDomainParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trustedDomainParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trustedDomainParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this common active directory params based on the context it is used
func (m *CommonActiveDirectoryParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMachineAccounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreferredDomainControllers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrustedDomainParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonActiveDirectoryParams) contextValidateMachineAccounts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MachineAccounts); i++ {

		if m.MachineAccounts[i] != nil {

			if swag.IsZero(m.MachineAccounts[i]) { // not required
				return nil
			}

			if err := m.MachineAccounts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("machineAccounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("machineAccounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonActiveDirectoryParams) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *CommonActiveDirectoryParams) contextValidatePreferredDomainControllers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PreferredDomainControllers); i++ {

		if m.PreferredDomainControllers[i] != nil {

			if swag.IsZero(m.PreferredDomainControllers[i]) { // not required
				return nil
			}

			if err := m.PreferredDomainControllers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("preferredDomainControllers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("preferredDomainControllers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonActiveDirectoryParams) contextValidateTrustedDomainParams(ctx context.Context, formats strfmt.Registry) error {

	if m.TrustedDomainParams != nil {

		if swag.IsZero(m.TrustedDomainParams) { // not required
			return nil
		}

		if err := m.TrustedDomainParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trustedDomainParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trustedDomainParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonActiveDirectoryParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonActiveDirectoryParams) UnmarshalBinary(b []byte) error {
	var res CommonActiveDirectoryParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
