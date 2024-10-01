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

// SchedulingPolicyRpoSchedule Specifies an RPO backup schedule.
// Set if periodicity is kContinuousRPO.
//
// swagger:model schedulingPolicyRpoSchedule
type SchedulingPolicyRpoSchedule struct {
	RpoSchedule
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SchedulingPolicyRpoSchedule) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 RpoSchedule
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.RpoSchedule = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SchedulingPolicyRpoSchedule) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.RpoSchedule)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this scheduling policy rpo schedule
func (m *SchedulingPolicyRpoSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with RpoSchedule
	if err := m.RpoSchedule.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this scheduling policy rpo schedule based on the context it is used
func (m *SchedulingPolicyRpoSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with RpoSchedule
	if err := m.RpoSchedule.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SchedulingPolicyRpoSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulingPolicyRpoSchedule) UnmarshalBinary(b []byte) error {
	var res SchedulingPolicyRpoSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
