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

// LocalCopyInfoBase local copy info base
//
// swagger:model LocalCopyInfoBase
type LocalCopyInfoBase struct {

	// If the flag magneto_master_notify_yoda_task_wise is set, magneto starts
	// notifying yoda whenever a task completes. This field is used for storing
	// task ids where were already yoda notified.
	YodaNotifiedTaskIds []*LocalCopyInfoBaseYodaNotifiedTaskIdsEntry `json:"yodaNotifiedTaskIds"`

	// This contains the value of flag magneto_master_notify_yoda_task_wise set
	// at the beginning of this run.
	YodaObjectLevelNotifyEnabled *bool `json:"yodaObjectLevelNotifyEnabled,omitempty"`
}

// Validate validates this local copy info base
func (m *LocalCopyInfoBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateYodaNotifiedTaskIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalCopyInfoBase) validateYodaNotifiedTaskIds(formats strfmt.Registry) error {
	if swag.IsZero(m.YodaNotifiedTaskIds) { // not required
		return nil
	}

	for i := 0; i < len(m.YodaNotifiedTaskIds); i++ {
		if swag.IsZero(m.YodaNotifiedTaskIds[i]) { // not required
			continue
		}

		if m.YodaNotifiedTaskIds[i] != nil {
			if err := m.YodaNotifiedTaskIds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("yodaNotifiedTaskIds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("yodaNotifiedTaskIds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this local copy info base based on the context it is used
func (m *LocalCopyInfoBase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateYodaNotifiedTaskIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalCopyInfoBase) contextValidateYodaNotifiedTaskIds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.YodaNotifiedTaskIds); i++ {

		if m.YodaNotifiedTaskIds[i] != nil {

			if swag.IsZero(m.YodaNotifiedTaskIds[i]) { // not required
				return nil
			}

			if err := m.YodaNotifiedTaskIds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("yodaNotifiedTaskIds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("yodaNotifiedTaskIds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LocalCopyInfoBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocalCopyInfoBase) UnmarshalBinary(b []byte) error {
	var res LocalCopyInfoBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
