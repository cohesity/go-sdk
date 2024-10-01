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

// ProtectedObject Protected Object
//
// Specifies the Protected Object with last Run's snapshot information per Protection Group.
//
// swagger:model ProtectedObject
type ProtectedObject struct {
	Object

	// Specifies the latest snapshot information for every Protection Group for a given object.
	LatestSnapshotsInfo []*ObjectSnapshotsInfo `json:"latestSnapshotsInfo"`

	// Specifies the Source Object information.
	SourceInfo *ObjectSummary `json:"sourceInfo,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ProtectedObject) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Object
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Object = aO0

	// AO1
	var dataAO1 struct {
		LatestSnapshotsInfo []*ObjectSnapshotsInfo `json:"latestSnapshotsInfo"`

		SourceInfo *ObjectSummary `json:"sourceInfo,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.LatestSnapshotsInfo = dataAO1.LatestSnapshotsInfo

	m.SourceInfo = dataAO1.SourceInfo

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ProtectedObject) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Object)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		LatestSnapshotsInfo []*ObjectSnapshotsInfo `json:"latestSnapshotsInfo"`

		SourceInfo *ObjectSummary `json:"sourceInfo,omitempty"`
	}

	dataAO1.LatestSnapshotsInfo = m.LatestSnapshotsInfo

	dataAO1.SourceInfo = m.SourceInfo

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this protected object
func (m *ProtectedObject) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Object
	if err := m.Object.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestSnapshotsInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtectedObject) validateLatestSnapshotsInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.LatestSnapshotsInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.LatestSnapshotsInfo); i++ {
		if swag.IsZero(m.LatestSnapshotsInfo[i]) { // not required
			continue
		}

		if m.LatestSnapshotsInfo[i] != nil {
			if err := m.LatestSnapshotsInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("latestSnapshotsInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("latestSnapshotsInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProtectedObject) validateSourceInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.SourceInfo) { // not required
		return nil
	}

	if m.SourceInfo != nil {
		if err := m.SourceInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourceInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sourceInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this protected object based on the context it is used
func (m *ProtectedObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Object
	if err := m.Object.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatestSnapshotsInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtectedObject) contextValidateLatestSnapshotsInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LatestSnapshotsInfo); i++ {

		if m.LatestSnapshotsInfo[i] != nil {

			if swag.IsZero(m.LatestSnapshotsInfo[i]) { // not required
				return nil
			}

			if err := m.LatestSnapshotsInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("latestSnapshotsInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("latestSnapshotsInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProtectedObject) contextValidateSourceInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.SourceInfo != nil {

		if swag.IsZero(m.SourceInfo) { // not required
			return nil
		}

		if err := m.SourceInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourceInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sourceInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProtectedObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectedObject) UnmarshalBinary(b []byte) error {
	var res ProtectedObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
