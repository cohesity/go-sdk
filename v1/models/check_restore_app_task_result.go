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

// CheckRestoreAppTaskResult check restore app task result
//
// swagger:model CheckRestoreAppTaskResult
type CheckRestoreAppTaskResult struct {

	// If exactly one entity is being restored, this is populated with the
	// display info for that entity.
	EntityInfo *CheckRestoreAppTaskResultEntityInfo `json:"entityInfo,omitempty"`

	// Error encountered by the RPC.
	Error *PrivateErrorProto `json:"error,omitempty"`

	// User message (if any) associated with the restore request.
	UserMessage *string `json:"userMessage,omitempty"`
}

// Validate validates this check restore app task result
func (m *CheckRestoreAppTaskResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CheckRestoreAppTaskResult) validateEntityInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityInfo) { // not required
		return nil
	}

	if m.EntityInfo != nil {
		if err := m.EntityInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CheckRestoreAppTaskResult) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this check restore app task result based on the context it is used
func (m *CheckRestoreAppTaskResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntityInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CheckRestoreAppTaskResult) contextValidateEntityInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityInfo != nil {

		if swag.IsZero(m.EntityInfo) { // not required
			return nil
		}

		if err := m.EntityInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CheckRestoreAppTaskResult) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {

		if swag.IsZero(m.Error) { // not required
			return nil
		}

		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CheckRestoreAppTaskResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckRestoreAppTaskResult) UnmarshalBinary(b []byte) error {
	var res CheckRestoreAppTaskResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
