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

// VolumeNameMap This message store the map, which store the information for each logical
// volume name to volume info.
//
// swagger:model VolumeNameMap
type VolumeNameMap struct {

	// Key for the map is volume name (guest mount point) and value is its
	// corresponding volume information.
	VolumeNameMap []*VolumeNameMapVolumeNameMapEntry `json:"volumeNameMap"`
}

// Validate validates this volume name map
func (m *VolumeNameMap) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVolumeNameMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeNameMap) validateVolumeNameMap(formats strfmt.Registry) error {
	if swag.IsZero(m.VolumeNameMap) { // not required
		return nil
	}

	for i := 0; i < len(m.VolumeNameMap); i++ {
		if swag.IsZero(m.VolumeNameMap[i]) { // not required
			continue
		}

		if m.VolumeNameMap[i] != nil {
			if err := m.VolumeNameMap[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeNameMap" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumeNameMap" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this volume name map based on the context it is used
func (m *VolumeNameMap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVolumeNameMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeNameMap) contextValidateVolumeNameMap(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VolumeNameMap); i++ {

		if m.VolumeNameMap[i] != nil {

			if swag.IsZero(m.VolumeNameMap[i]) { // not required
				return nil
			}

			if err := m.VolumeNameMap[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeNameMap" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumeNameMap" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VolumeNameMap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeNameMap) UnmarshalBinary(b []byte) error {
	var res VolumeNameMap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
