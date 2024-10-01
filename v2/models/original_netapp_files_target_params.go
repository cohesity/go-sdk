// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OriginalNetappFilesTargetParams Recover To Original Netapp Files Target Params.
//
// Specifies the params of the original Netapp recovery target.
//
// swagger:model OriginalNetappFilesTargetParams
type OriginalNetappFilesTargetParams struct {

	// Specifies whether to recover files and folders to the original path location. If false, alternatePath must be specified.
	// Required: true
	RecoverToOriginalPath *bool `json:"recoverToOriginalPath"`

	// Specifies the alternate path location to recover files to.
	AlternatePath *string `json:"alternatePath,omitempty"`

	// Specifies whether to overwrite existing file/folder during recovery.
	OverwriteExistingFile *bool `json:"overwriteExistingFile,omitempty"`

	// Specifies whether to preserve file/folder attributes during recovery.
	PreserveFileAttributes *bool `json:"preserveFileAttributes,omitempty"`

	// Specifies whether to continue recovering other files if one of the files fails to recover. Default value is false.
	ContinueOnError *bool `json:"continueOnError,omitempty"`

	// Specifies whether encryption should be enabled during recovery.
	EncryptionEnabled *bool `json:"encryptionEnabled,omitempty"`

	// Specifies the list of IP addresses that are allowed or denied during recovery. Allowed IPs and Denied IPs cannot be used together.
	FilterIPConfig *FilterIPConfig `json:"filterIpConfig,omitempty"`

	// Specifies VLAN settings associated with the restore. If this is not specified, then the VLAN params will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for restores.
	VlanConfig *RecoveryVlanConfig `json:"vlanConfig,omitempty"`
}

// Validate validates this original netapp files target params
func (m *OriginalNetappFilesTargetParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecoverToOriginalPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilterIPConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OriginalNetappFilesTargetParams) validateRecoverToOriginalPath(formats strfmt.Registry) error {

	if err := validate.Required("recoverToOriginalPath", "body", m.RecoverToOriginalPath); err != nil {
		return err
	}

	return nil
}

func (m *OriginalNetappFilesTargetParams) validateFilterIPConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.FilterIPConfig) { // not required
		return nil
	}

	if m.FilterIPConfig != nil {
		if err := m.FilterIPConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filterIpConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filterIpConfig")
			}
			return err
		}
	}

	return nil
}

func (m *OriginalNetappFilesTargetParams) validateVlanConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.VlanConfig) { // not required
		return nil
	}

	if m.VlanConfig != nil {
		if err := m.VlanConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlanConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlanConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this original netapp files target params based on the context it is used
func (m *OriginalNetappFilesTargetParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilterIPConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlanConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OriginalNetappFilesTargetParams) contextValidateFilterIPConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.FilterIPConfig != nil {

		if swag.IsZero(m.FilterIPConfig) { // not required
			return nil
		}

		if err := m.FilterIPConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filterIpConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filterIpConfig")
			}
			return err
		}
	}

	return nil
}

func (m *OriginalNetappFilesTargetParams) contextValidateVlanConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.VlanConfig != nil {

		if swag.IsZero(m.VlanConfig) { // not required
			return nil
		}

		if err := m.VlanConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlanConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlanConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OriginalNetappFilesTargetParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OriginalNetappFilesTargetParams) UnmarshalBinary(b []byte) error {
	var res OriginalNetappFilesTargetParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
