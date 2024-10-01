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

// ProgressTask Progress Task
//
// This specifies the details about the Progress Task.
//
// swagger:model ProgressTask
type ProgressTask struct {

	// Specifies the task id of the Progress task.
	ID *string `json:"id,omitempty"`

	ProgressTaskInfo
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ProgressTask) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		ID *string `json:"id,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.ID = dataAO0.ID

	// AO1
	var aO1 ProgressTaskInfo
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ProgressTaskInfo = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ProgressTask) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		ID *string `json:"id,omitempty"`
	}

	dataAO0.ID = m.ID

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	aO1, err := swag.WriteJSON(m.ProgressTaskInfo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this progress task
func (m *ProgressTask) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ProgressTaskInfo
	if err := m.ProgressTaskInfo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this progress task based on the context it is used
func (m *ProgressTask) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ProgressTaskInfo
	if err := m.ProgressTaskInfo.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ProgressTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProgressTask) UnmarshalBinary(b []byte) error {
	var res ProgressTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
