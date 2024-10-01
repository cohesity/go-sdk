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

// CreateActiveDirectoryEntryParams Specifies the params to join a Microsoft Active Directory domain.
//
// swagger:model CreateActiveDirectoryEntryParams
type CreateActiveDirectoryEntryParams struct {

	// Specifies the fully qualified domain name (FQDN) of an Active Directory.
	DomainName *string `json:"domainName,omitempty"`

	// Specifies the fallback id mapping info which is used when an ID mapping
	// for a user is not found via the above IdMappingInfo. Only supported for
	// two types of fallback mapping types - 'kRid' and 'kFixed'.
	FallbackUserIDMappingInfo *UserIDMapping `json:"fallbackUserIdMappingInfo,omitempty"`

	// Specifies if Active Directory entry should be force removed from cluster.
	ForceRemove *bool `json:"forceRemove,omitempty"`

	// Specifies the list of trusted domains that were set by the user to be
	// ignored during trusted domain discovery.
	IgnoredTrustedDomains []string `json:"ignoredTrustedDomains"`

	// Specifies the LDAP provider id which is map to this Active Directory
	LdapProviderID *int64 `json:"ldapProviderId,omitempty"`

	// Array of Machine Accounts.
	//
	// Specifies an array of computer names used to identify the Cohesity
	// Cluster on the domain.
	MachineAccounts []string `json:"machineAccounts"`

	// Specifies whether to use 'whitelistedDomains' only for authentication.
	OnlyUseWhitelistedDomains *bool `json:"onlyUseWhitelistedDomains,omitempty"`

	// Specifies an optional Organizational Unit name.
	OuName *string `json:"ouName,omitempty"`

	// Specifies whether the specified machine accounts should overwrite the
	// existing machine accounts in this domain.
	OverwriteExistingAccounts *bool `json:"overwriteExistingAccounts,omitempty"`

	// Specifies the password for the specified userName.
	Password *string `json:"password,omitempty"`

	// Specifies Map of Active Directory domain names to its preferred domain
	// controllers.
	PreferredDomainControllers []*PreferredDomainController `json:"preferredDomainControllers"`

	// Specifies the task path for AD joining task.
	TaskPath *string `json:"taskPath,omitempty"`

	// Specifies the unique id of the tenant.
	TenantID *string `json:"tenantId,omitempty"`

	// Specifies the trusted domains of the Active Directory domain.
	// Read Only: true
	TrustedDomains []string `json:"trustedDomains"`

	// Specifies whether Trusted Domain discovery is disabled.
	TrustedDomainsEnabled *bool `json:"trustedDomainsEnabled,omitempty"`

	// Specifies the SID of the Active Directory domain user to be mapped to
	// Unix root user.
	UnixRootSid *string `json:"unixRootSid,omitempty"`

	// Specifies the information about how the Unix and Windows users are mapped
	// for this domain.
	UserIDMappingInfo *UserIDMapping `json:"userIdMappingInfo,omitempty"`

	// Specifies a userName that has administrative privileges in the domain.
	UserName *string `json:"userName,omitempty"`

	// Specifies the Whitelisted Domains of the Active Directory domain.
	WhitelistedDomains []string `json:"whitelistedDomains"`

	// Specifies an optional Workgroup name.
	Workgroup *string `json:"workgroup,omitempty"`
}

// Validate validates this create active directory entry params
func (m *CreateActiveDirectoryEntryParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFallbackUserIDMappingInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreferredDomainControllers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserIDMappingInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateActiveDirectoryEntryParams) validateFallbackUserIDMappingInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.FallbackUserIDMappingInfo) { // not required
		return nil
	}

	if m.FallbackUserIDMappingInfo != nil {
		if err := m.FallbackUserIDMappingInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fallbackUserIdMappingInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fallbackUserIdMappingInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CreateActiveDirectoryEntryParams) validatePreferredDomainControllers(formats strfmt.Registry) error {
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

func (m *CreateActiveDirectoryEntryParams) validateUserIDMappingInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.UserIDMappingInfo) { // not required
		return nil
	}

	if m.UserIDMappingInfo != nil {
		if err := m.UserIDMappingInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userIdMappingInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userIdMappingInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create active directory entry params based on the context it is used
func (m *CreateActiveDirectoryEntryParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFallbackUserIDMappingInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreferredDomainControllers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrustedDomains(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserIDMappingInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateActiveDirectoryEntryParams) contextValidateFallbackUserIDMappingInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.FallbackUserIDMappingInfo != nil {

		if swag.IsZero(m.FallbackUserIDMappingInfo) { // not required
			return nil
		}

		if err := m.FallbackUserIDMappingInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fallbackUserIdMappingInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fallbackUserIdMappingInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CreateActiveDirectoryEntryParams) contextValidatePreferredDomainControllers(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CreateActiveDirectoryEntryParams) contextValidateTrustedDomains(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "trustedDomains", "body", []string(m.TrustedDomains)); err != nil {
		return err
	}

	return nil
}

func (m *CreateActiveDirectoryEntryParams) contextValidateUserIDMappingInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.UserIDMappingInfo != nil {

		if swag.IsZero(m.UserIDMappingInfo) { // not required
			return nil
		}

		if err := m.UserIDMappingInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userIdMappingInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userIdMappingInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateActiveDirectoryEntryParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateActiveDirectoryEntryParams) UnmarshalBinary(b []byte) error {
	var res CreateActiveDirectoryEntryParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
