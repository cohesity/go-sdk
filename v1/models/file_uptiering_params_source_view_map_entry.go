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

// FileUptieringParamsSourceViewMapEntry file uptiering params source view map entry
//
// swagger:model FileUptieringParams_SourceViewMapEntry
type FileUptieringParamsSourceViewMapEntry struct {

	// key
	Key *int64 `json:"key,omitempty"`

	// value
	Value *FileUptieringParamsSourceViewData `json:"value,omitempty"`
}

// Validate validates this file uptiering params source view map entry
func (m *FileUptieringParamsSourceViewMapEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileUptieringParamsSourceViewMapEntry) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this file uptiering params source view map entry based on the context it is used
func (m *FileUptieringParamsSourceViewMapEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileUptieringParamsSourceViewMapEntry) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if m.Value != nil {

		if swag.IsZero(m.Value) { // not required
			return nil
		}

		if err := m.Value.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileUptieringParamsSourceViewMapEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileUptieringParamsSourceViewMapEntry) UnmarshalBinary(b []byte) error {
	var res FileUptieringParamsSourceViewMapEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
