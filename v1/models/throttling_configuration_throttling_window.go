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

// ThrottlingConfigurationThrottlingWindow Message encapsulating a DayTimeWindow and the corresponding threshold
// value applicable in that window.
//
// swagger:model ThrottlingConfiguration_ThrottlingWindow
type ThrottlingConfigurationThrottlingWindow struct {

	// Day time window. This captures day and time window on that day.
	DayTimeWindow *PrivateDayTimeWindow `json:"dayTimeWindow,omitempty"`

	// Throttling threshold applicable in the window.
	Threshold *int64 `json:"threshold,omitempty"`
}

// Validate validates this throttling configuration throttling window
func (m *ThrottlingConfigurationThrottlingWindow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDayTimeWindow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThrottlingConfigurationThrottlingWindow) validateDayTimeWindow(formats strfmt.Registry) error {
	if swag.IsZero(m.DayTimeWindow) { // not required
		return nil
	}

	if m.DayTimeWindow != nil {
		if err := m.DayTimeWindow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dayTimeWindow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dayTimeWindow")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this throttling configuration throttling window based on the context it is used
func (m *ThrottlingConfigurationThrottlingWindow) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDayTimeWindow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThrottlingConfigurationThrottlingWindow) contextValidateDayTimeWindow(ctx context.Context, formats strfmt.Registry) error {

	if m.DayTimeWindow != nil {

		if swag.IsZero(m.DayTimeWindow) { // not required
			return nil
		}

		if err := m.DayTimeWindow.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dayTimeWindow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dayTimeWindow")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThrottlingConfigurationThrottlingWindow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThrottlingConfigurationThrottlingWindow) UnmarshalBinary(b []byte) error {
	var res ThrottlingConfigurationThrottlingWindow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
