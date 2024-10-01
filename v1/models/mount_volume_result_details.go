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

// MountVolumeResultDetails Mount Volume Result Details.
//
// Specifies the result of mounting an individual mount volume in a
// Restore Task.
//
// swagger:model MountVolumeResultDetails
type MountVolumeResultDetails struct {

	// mount error
	MountError *MountVolumeResultDetailsMountError `json:"mountError,omitempty"`

	// Specifies the mount point where the volume is mounted.
	// NOTE: This field may not be populated for VM environments if the
	// onlining of disks is not requested or there was any issue during onlining.
	MountPoint *string `json:"mountPoint,omitempty"`

	// Specifies the name of the original volume.
	VolumeName *string `json:"volumeName,omitempty"`
}

// Validate validates this mount volume result details
func (m *MountVolumeResultDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMountError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MountVolumeResultDetails) validateMountError(formats strfmt.Registry) error {
	if swag.IsZero(m.MountError) { // not required
		return nil
	}

	if m.MountError != nil {
		if err := m.MountError.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mountError")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mountError")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mount volume result details based on the context it is used
func (m *MountVolumeResultDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMountError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MountVolumeResultDetails) contextValidateMountError(ctx context.Context, formats strfmt.Registry) error {

	if m.MountError != nil {

		if swag.IsZero(m.MountError) { // not required
			return nil
		}

		if err := m.MountError.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mountError")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mountError")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MountVolumeResultDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MountVolumeResultDetails) UnmarshalBinary(b []byte) error {
	var res MountVolumeResultDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
