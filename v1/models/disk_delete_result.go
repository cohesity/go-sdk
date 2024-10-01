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

// DiskDeleteResult DiskDeleteResult
//
// swagger:model DiskDeleteResult
type DiskDeleteResult struct {

	// Id of the disk to be marked for deletion.
	ID *int64 `json:"id,omitempty"`

	// MarkDiskForRemoval indicates if the disk is marked for removal
	MarkDiskForRemoval *bool `json:"markDiskForRemoval,omitempty"`

	// TimestampSecs specifies the last run time
	// of the pre-checks execution in Unix epoch timestamp in seconds
	TimestampSecs *int64 `json:"timestampSecs,omitempty"`

	// ValidationChecks specifies list of pre-check validations
	ValidationChecks []*PreCheckValidation `json:"validationChecks"`
}

// Validate validates this disk delete result
func (m *DiskDeleteResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValidationChecks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiskDeleteResult) validateValidationChecks(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidationChecks) { // not required
		return nil
	}

	for i := 0; i < len(m.ValidationChecks); i++ {
		if swag.IsZero(m.ValidationChecks[i]) { // not required
			continue
		}

		if m.ValidationChecks[i] != nil {
			if err := m.ValidationChecks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validationChecks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validationChecks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this disk delete result based on the context it is used
func (m *DiskDeleteResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValidationChecks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiskDeleteResult) contextValidateValidationChecks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ValidationChecks); i++ {

		if m.ValidationChecks[i] != nil {

			if swag.IsZero(m.ValidationChecks[i]) { // not required
				return nil
			}

			if err := m.ValidationChecks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validationChecks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validationChecks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DiskDeleteResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiskDeleteResult) UnmarshalBinary(b []byte) error {
	var res DiskDeleteResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
