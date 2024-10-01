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

// RecoverOtherNasToNetappVolumeTargetParams Recover To Netapp Volume Target Params.
//
// Specifies the params of the Netapp recovery target.
//
// swagger:model RecoverOtherNasToNetappVolumeTargetParams
type RecoverOtherNasToNetappVolumeTargetParams struct {

	// Specifies whether to overwrite existing file/folder during recovery.
	OverwriteExistingFile *bool `json:"overwriteExistingFile,omitempty"`

	// Specifies whether to preserve file/folder attributes during recovery.
	PreserveFileAttributes *bool `json:"preserveFileAttributes,omitempty"`

	// Specifies whether to continue recovering other volumes if one of the volumes fails to recover. Default value is false.
	ContinueOnError *bool `json:"continueOnError,omitempty"`

	// Specifies whether encryption should be enabled during recovery.
	EncryptionEnabled *bool `json:"encryptionEnabled,omitempty"`

	// Specifies the list of IP addresses that are allowed or denied during recovery. Allowed IPs and Denied IPs cannot be used together.
	FilterIPConfig *FilterIPConfig `json:"filterIpConfig,omitempty"`

	// Specifies VLAN settings associated with the restore. If this is not specified, then the VLAN params will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for restores.
	VlanConfig *RecoveryVlanConfig `json:"vlanConfig,omitempty"`

	// Specifies the id of the parent source of the recovery target.
	ParentSource *RecoveryObjectIdentifier `json:"parentSource,omitempty"`

	// Specifies the id and name of the volume to recover to. This volume will be the target of the recovery.
	// Required: true
	Volume *RecoveryObjectIdentifier `json:"volume"`
}

// Validate validates this recover other nas to netapp volume target params
func (m *RecoverOtherNasToNetappVolumeTargetParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilterIPConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverOtherNasToNetappVolumeTargetParams) validateFilterIPConfig(formats strfmt.Registry) error {
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

func (m *RecoverOtherNasToNetappVolumeTargetParams) validateVlanConfig(formats strfmt.Registry) error {
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

func (m *RecoverOtherNasToNetappVolumeTargetParams) validateParentSource(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentSource) { // not required
		return nil
	}

	if m.ParentSource != nil {
		if err := m.ParentSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parentSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parentSource")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverOtherNasToNetappVolumeTargetParams) validateVolume(formats strfmt.Registry) error {

	if err := validate.Required("volume", "body", m.Volume); err != nil {
		return err
	}

	if m.Volume != nil {
		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover other nas to netapp volume target params based on the context it is used
func (m *RecoverOtherNasToNetappVolumeTargetParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilterIPConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlanConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParentSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverOtherNasToNetappVolumeTargetParams) contextValidateFilterIPConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecoverOtherNasToNetappVolumeTargetParams) contextValidateVlanConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecoverOtherNasToNetappVolumeTargetParams) contextValidateParentSource(ctx context.Context, formats strfmt.Registry) error {

	if m.ParentSource != nil {

		if swag.IsZero(m.ParentSource) { // not required
			return nil
		}

		if err := m.ParentSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parentSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parentSource")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverOtherNasToNetappVolumeTargetParams) contextValidateVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.Volume != nil {

		if err := m.Volume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverOtherNasToNetappVolumeTargetParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverOtherNasToNetappVolumeTargetParams) UnmarshalBinary(b []byte) error {
	var res RecoverOtherNasToNetappVolumeTargetParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
