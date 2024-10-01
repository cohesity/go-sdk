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

// RecoverHiveSnapshotParams Recover Hive Snapshot Params.
//
// Specifies the snapshot parameters for a protected Hive object.
//
// swagger:model RecoverHiveSnapshotParams
type RecoverHiveSnapshotParams struct {
	CommonRecoverObjectSnapshotParams

	// Specifies details of objects to be recovered.
	Objects []*RecoverHiveNoSQLObjectParams `json:"objects"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RecoverHiveSnapshotParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonRecoverObjectSnapshotParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonRecoverObjectSnapshotParams = aO0

	// AO1
	var dataAO1 struct {
		Objects []*RecoverHiveNoSQLObjectParams `json:"objects"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Objects = dataAO1.Objects

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RecoverHiveSnapshotParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonRecoverObjectSnapshotParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Objects []*RecoverHiveNoSQLObjectParams `json:"objects"`
	}

	dataAO1.Objects = m.Objects

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this recover hive snapshot params
func (m *RecoverHiveSnapshotParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonRecoverObjectSnapshotParams
	if err := m.CommonRecoverObjectSnapshotParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverHiveSnapshotParams) validateObjects(formats strfmt.Registry) error {

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

// ContextValidate validate this recover hive snapshot params based on the context it is used
func (m *RecoverHiveSnapshotParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonRecoverObjectSnapshotParams
	if err := m.CommonRecoverObjectSnapshotParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverHiveSnapshotParams) contextValidateObjects(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *RecoverHiveSnapshotParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverHiveSnapshotParams) UnmarshalBinary(b []byte) error {
	var res RecoverHiveSnapshotParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
