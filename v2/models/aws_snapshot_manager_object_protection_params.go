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

// AwsSnapshotManagerObjectProtectionParams Create AWS Snapshot Manager Object Protection Request Body
//
// Specifies the parameters which are specific to AWS Object Protection using AWS native snapshot orchestration with snapshot manager. Atlease one of tags or objects must be specified.
//
// swagger:model AwsSnapshotManagerObjectProtectionParams
type AwsSnapshotManagerObjectProtectionParams struct {

	// Specifies the objects to be protected.
	Objects []*AwsObjectLevelParams `json:"objects"`

	// Specifies the paramaters to exclude volumes attached to EC2 instances at global level.
	VolumeExclusionParams *EbsVolumeExclusionParams `json:"volumeExclusionParams,omitempty"`

	// Specifies whether AMI should be created after taking snapshots of the instance.
	CreateAmi *bool `json:"createAmi,omitempty"`

	// The frequency of AMI creation. This should be set if the option to create AMI is set. A value of n creates an AMI from the snapshots after every n runs. eg. n = 2 implies every alternate backup run starting from the first will create an AMI.
	AmiCreationFrequency *int32 `json:"amiCreationFrequency,omitempty"`
}

// Validate validates this aws snapshot manager object protection params
func (m *AwsSnapshotManagerObjectProtectionParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeExclusionParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsSnapshotManagerObjectProtectionParams) validateObjects(formats strfmt.Registry) error {
	if swag.IsZero(m.Objects) { // not required
		return nil
	}

	for i := 0; i < len(m.Objects); i++ {
		if swag.IsZero(m.Objects[i]) { // not required
			continue
		}

		if m.Objects[i] != nil {
			if err := m.Objects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AwsSnapshotManagerObjectProtectionParams) validateVolumeExclusionParams(formats strfmt.Registry) error {
	if swag.IsZero(m.VolumeExclusionParams) { // not required
		return nil
	}

	if m.VolumeExclusionParams != nil {
		if err := m.VolumeExclusionParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volumeExclusionParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volumeExclusionParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this aws snapshot manager object protection params based on the context it is used
func (m *AwsSnapshotManagerObjectProtectionParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumeExclusionParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsSnapshotManagerObjectProtectionParams) contextValidateObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Objects); i++ {

		if m.Objects[i] != nil {

			if swag.IsZero(m.Objects[i]) { // not required
				return nil
			}

			if err := m.Objects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AwsSnapshotManagerObjectProtectionParams) contextValidateVolumeExclusionParams(ctx context.Context, formats strfmt.Registry) error {

	if m.VolumeExclusionParams != nil {

		if swag.IsZero(m.VolumeExclusionParams) { // not required
			return nil
		}

		if err := m.VolumeExclusionParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volumeExclusionParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volumeExclusionParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsSnapshotManagerObjectProtectionParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsSnapshotManagerObjectProtectionParams) UnmarshalBinary(b []byte) error {
	var res AwsSnapshotManagerObjectProtectionParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
