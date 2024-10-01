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

// MountVolumesState Mount Volumes States.
//
// Specifies the states of mounting all the volumes onto a mount target
// for a 'kRecoverVMs' Restore Task.
//
// swagger:model MountVolumesState
type MountVolumesState struct {

	// Optional setting that determines if the volumes are brought
	// online on the mount target after attaching the disks.
	// This option is only significant for VMs.
	BringDisksOnline *bool `json:"bringDisksOnline,omitempty"`

	// Array of Mount Volume Results.
	//
	// Specifies the results of mounting each specified volume.
	MountVolumeResults []*MountVolumeResultDetails `json:"mountVolumeResults"`

	// other error
	OtherError *MountVolumesStateOtherError `json:"otherError,omitempty"`

	// Specifies the target Protection Source Id where the volumes will be
	// mounted.
	// NOTE: The source that was backed up and the mount target must be the
	// same type, for example if the source is a VMware VM, then the mount
	// target must also be a VMware VM.
	// The mount target must be registered on the Cohesity Cluster.
	TargetSourceID *int64 `json:"targetSourceId,omitempty"`

	// Specifies the username to access the mount target.
	Username *string `json:"username,omitempty"`
}

// Validate validates this mount volumes state
func (m *MountVolumesState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMountVolumeResults(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOtherError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MountVolumesState) validateMountVolumeResults(formats strfmt.Registry) error {
	if swag.IsZero(m.MountVolumeResults) { // not required
		return nil
	}

	for i := 0; i < len(m.MountVolumeResults); i++ {
		if swag.IsZero(m.MountVolumeResults[i]) { // not required
			continue
		}

		if m.MountVolumeResults[i] != nil {
			if err := m.MountVolumeResults[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mountVolumeResults" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mountVolumeResults" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MountVolumesState) validateOtherError(formats strfmt.Registry) error {
	if swag.IsZero(m.OtherError) { // not required
		return nil
	}

	if m.OtherError != nil {
		if err := m.OtherError.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("otherError")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("otherError")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mount volumes state based on the context it is used
func (m *MountVolumesState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMountVolumeResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOtherError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MountVolumesState) contextValidateMountVolumeResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MountVolumeResults); i++ {

		if m.MountVolumeResults[i] != nil {

			if swag.IsZero(m.MountVolumeResults[i]) { // not required
				return nil
			}

			if err := m.MountVolumeResults[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mountVolumeResults" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mountVolumeResults" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MountVolumesState) contextValidateOtherError(ctx context.Context, formats strfmt.Registry) error {

	if m.OtherError != nil {

		if swag.IsZero(m.OtherError) { // not required
			return nil
		}

		if err := m.OtherError.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("otherError")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("otherError")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MountVolumesState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MountVolumesState) UnmarshalBinary(b []byte) error {
	var res MountVolumesState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
