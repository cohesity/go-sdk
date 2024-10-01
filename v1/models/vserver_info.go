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
)

// VserverInfo Proto that contains specific information about a vserver in a Netapp
// cluster.
//
// swagger:model VserverInfo
type VserverInfo struct {

	// The data protocols supported by this vserver. kManagement protocol does
	// not show in this vector.
	DataProtocolVec []int32 `json:"dataProtocolVec"`

	// Information about all interfaces on this vserver.
	InterfaceVec []*VserverInfoNetworkInterfaceInfo `json:"interfaceVec"`

	// Whether the vserver has enabled NFSv3 protocol.
	// TODO(akarsh) : Remove this after enough soak time for NFSv4.1 workflow.
	Nfsv3Enabled *bool `json:"nfsv3Enabled,omitempty"`

	// Whether the vserver has enabled NFSv4.1 protocol.
	// TODO(akarsh) : Remove this after enough soak time for NFSv4.1 workflow.
	Nfsv41Enabled *bool `json:"nfsv41Enabled,omitempty"`

	// The root 'c$' cifs share of this vserver. If exists it can be used to
	// mount all CIFS volumes that are junctioned under '/' on this vserver.
	RootCifsShare *PrivateCifsShareInfo `json:"rootCifsShare,omitempty"`

	// The smb krb5 hostname.
	SmbKrb5Hostname *string `json:"smbKrb5Hostname,omitempty"`

	// The type of this vserver.
	Type *int32 `json:"type,omitempty"`

	// Volumes managed by this VServer. Only populated for virtualized storage
	// snapshots backup.
	VolumeInfoVec []*PrivateNetappEntity `json:"volumeInfoVec"`
}

// Validate validates this vserver info
func (m *VserverInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInterfaceVec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootCifsShare(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeInfoVec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VserverInfo) validateInterfaceVec(formats strfmt.Registry) error {
	if swag.IsZero(m.InterfaceVec) { // not required
		return nil
	}

	for i := 0; i < len(m.InterfaceVec); i++ {
		if swag.IsZero(m.InterfaceVec[i]) { // not required
			continue
		}

		if m.InterfaceVec[i] != nil {
			if err := m.InterfaceVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaceVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("interfaceVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VserverInfo) validateRootCifsShare(formats strfmt.Registry) error {
	if swag.IsZero(m.RootCifsShare) { // not required
		return nil
	}

	if m.RootCifsShare != nil {
		if err := m.RootCifsShare.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rootCifsShare")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rootCifsShare")
			}
			return err
		}
	}

	return nil
}

func (m *VserverInfo) validateVolumeInfoVec(formats strfmt.Registry) error {
	if swag.IsZero(m.VolumeInfoVec) { // not required
		return nil
	}

	for i := 0; i < len(m.VolumeInfoVec); i++ {
		if swag.IsZero(m.VolumeInfoVec[i]) { // not required
			continue
		}

		if m.VolumeInfoVec[i] != nil {
			if err := m.VolumeInfoVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeInfoVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumeInfoVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this vserver info based on the context it is used
func (m *VserverInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInterfaceVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRootCifsShare(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumeInfoVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VserverInfo) contextValidateInterfaceVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InterfaceVec); i++ {

		if m.InterfaceVec[i] != nil {

			if swag.IsZero(m.InterfaceVec[i]) { // not required
				return nil
			}

			if err := m.InterfaceVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaceVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("interfaceVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VserverInfo) contextValidateRootCifsShare(ctx context.Context, formats strfmt.Registry) error {

	if m.RootCifsShare != nil {

		if swag.IsZero(m.RootCifsShare) { // not required
			return nil
		}

		if err := m.RootCifsShare.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rootCifsShare")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rootCifsShare")
			}
			return err
		}
	}

	return nil
}

func (m *VserverInfo) contextValidateVolumeInfoVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VolumeInfoVec); i++ {

		if m.VolumeInfoVec[i] != nil {

			if swag.IsZero(m.VolumeInfoVec[i]) { // not required
				return nil
			}

			if err := m.VolumeInfoVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeInfoVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumeInfoVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VserverInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VserverInfo) UnmarshalBinary(b []byte) error {
	var res VserverInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
