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

// VmwareRecoverDisksTargetSourceConfig Specifies the configuration for restoring disks to a different VM than the one from which the snapshot was taken.
//
// swagger:model VmwareRecoverDisksTargetSourceConfig
type VmwareRecoverDisksTargetSourceConfig struct {

	// Specifies the source ID of the VM to which the disks will be restored.
	// Required: true
	SourceID *int64 `json:"sourceId"`

	// Specifies the disks to be recovered and the location to which they will be recovered.
	// Required: true
	// Min Items: 1
	Disks []*VmwareRecoverTargetSourceDiskParams `json:"disks"`
}

// Validate validates this vmware recover disks target source config
func (m *VmwareRecoverDisksTargetSourceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmwareRecoverDisksTargetSourceConfig) validateSourceID(formats strfmt.Registry) error {

	if err := validate.Required("sourceId", "body", m.SourceID); err != nil {
		return err
	}

	return nil
}

func (m *VmwareRecoverDisksTargetSourceConfig) validateDisks(formats strfmt.Registry) error {

	if err := validate.Required("disks", "body", m.Disks); err != nil {
		return err
	}

	iDisksSize := int64(len(m.Disks))

	if err := validate.MinItems("disks", "body", iDisksSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Disks); i++ {
		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {
			if err := m.Disks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this vmware recover disks target source config based on the context it is used
func (m *VmwareRecoverDisksTargetSourceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmwareRecoverDisksTargetSourceConfig) contextValidateDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Disks); i++ {

		if m.Disks[i] != nil {

			if swag.IsZero(m.Disks[i]) { // not required
				return nil
			}

			if err := m.Disks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VmwareRecoverDisksTargetSourceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareRecoverDisksTargetSourceConfig) UnmarshalBinary(b []byte) error {
	var res VmwareRecoverDisksTargetSourceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
