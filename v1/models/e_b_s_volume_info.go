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

// EBSVolumeInfo Message representing information related to a volume attached to an EC2
// instance.
//
// swagger:model EBSVolumeInfo
type EBSVolumeInfo struct {

	// Name of the device. Eg - /dev/sdb.
	DeviceName *string `json:"deviceName,omitempty"`

	// ID of the volume.
	ID *string `json:"id,omitempty"`

	// True iff the volume is attached as root device for an instance.
	IsRootDevice *bool `json:"isRootDevice,omitempty"`

	// Name of the volume.
	Name *string `json:"name,omitempty"`

	// Size of the volume in bytes.
	SizeBytes *int64 `json:"sizeBytes,omitempty"`

	// tag vec
	TagVec []*EBSVolumeInfoTag `json:"tagVec"`

	// Type of the volume. Eg - gp2, io1.
	Type *string `json:"type,omitempty"`
}

// Validate validates this e b s volume info
func (m *EBSVolumeInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTagVec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EBSVolumeInfo) validateTagVec(formats strfmt.Registry) error {
	if swag.IsZero(m.TagVec) { // not required
		return nil
	}

	for i := 0; i < len(m.TagVec); i++ {
		if swag.IsZero(m.TagVec[i]) { // not required
			continue
		}

		if m.TagVec[i] != nil {
			if err := m.TagVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tagVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tagVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this e b s volume info based on the context it is used
func (m *EBSVolumeInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTagVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EBSVolumeInfo) contextValidateTagVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TagVec); i++ {

		if m.TagVec[i] != nil {

			if swag.IsZero(m.TagVec[i]) { // not required
				return nil
			}

			if err := m.TagVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tagVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tagVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EBSVolumeInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EBSVolumeInfo) UnmarshalBinary(b []byte) error {
	var res EBSVolumeInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
