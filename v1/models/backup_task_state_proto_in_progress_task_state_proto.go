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

// BackupTaskStateProtoInProgressTaskStateProto State maintained for the task only as long as the task is pending. Cleared
// when the task is finished or cancelled.
//
// swagger:model BackupTaskStateProto_InProgressTaskStateProto
type BackupTaskStateProtoInProgressTaskStateProto struct {

	// acquired app entity lock vec
	AcquiredAppEntityLockVec []*PrivateEntityProto `json:"acquiredAppEntityLockVec"`
}

// Validate validates this backup task state proto in progress task state proto
func (m *BackupTaskStateProtoInProgressTaskStateProto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcquiredAppEntityLockVec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupTaskStateProtoInProgressTaskStateProto) validateAcquiredAppEntityLockVec(formats strfmt.Registry) error {
	if swag.IsZero(m.AcquiredAppEntityLockVec) { // not required
		return nil
	}

	for i := 0; i < len(m.AcquiredAppEntityLockVec); i++ {
		if swag.IsZero(m.AcquiredAppEntityLockVec[i]) { // not required
			continue
		}

		if m.AcquiredAppEntityLockVec[i] != nil {
			if err := m.AcquiredAppEntityLockVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acquiredAppEntityLockVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("acquiredAppEntityLockVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this backup task state proto in progress task state proto based on the context it is used
func (m *BackupTaskStateProtoInProgressTaskStateProto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAcquiredAppEntityLockVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupTaskStateProtoInProgressTaskStateProto) contextValidateAcquiredAppEntityLockVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AcquiredAppEntityLockVec); i++ {

		if m.AcquiredAppEntityLockVec[i] != nil {

			if swag.IsZero(m.AcquiredAppEntityLockVec[i]) { // not required
				return nil
			}

			if err := m.AcquiredAppEntityLockVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acquiredAppEntityLockVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("acquiredAppEntityLockVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupTaskStateProtoInProgressTaskStateProto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupTaskStateProtoInProgressTaskStateProto) UnmarshalBinary(b []byte) error {
	var res BackupTaskStateProtoInProgressTaskStateProto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
