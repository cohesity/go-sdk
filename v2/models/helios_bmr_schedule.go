// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HeliosBmrSchedule Bmr Schedule
//
// Specifies settings that defines how frequent bmr backup will be performed for a Protection Group.
//
// swagger:model HeliosBmrSchedule
type HeliosBmrSchedule struct {

	// Specifies how often to start new runs of a Protection Group. <br>'Weeks' specifies that new Protection Group runs start weekly on certain days specified using 'dayOfWeek' field. <br>'Months' specifies that new Protection Group runs start monthly on certain day of specific week.
	// Enum: ["Weeks","Months"]
	Unit *string `json:"unit,omitempty"`

	// Specifies the week Schedule for Protection Group to start runs on certain number of days in a week.
	WeekSchedule *HeliosWeekSchedule `json:"weekSchedule,omitempty"`

	// Specifies the week Schedule for Protection Group to start runs on specific week in a month and specific days of that week.
	MonthSchedule *HeliosMonthSchedule `json:"monthSchedule,omitempty"`
}

// Validate validates this helios bmr schedule
func (m *HeliosBmrSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeekSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonthSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var heliosBmrScheduleTypeUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Weeks","Months"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		heliosBmrScheduleTypeUnitPropEnum = append(heliosBmrScheduleTypeUnitPropEnum, v)
	}
}

const (

	// HeliosBmrScheduleUnitWeeks captures enum value "Weeks"
	HeliosBmrScheduleUnitWeeks string = "Weeks"

	// HeliosBmrScheduleUnitMonths captures enum value "Months"
	HeliosBmrScheduleUnitMonths string = "Months"
)

// prop value enum
func (m *HeliosBmrSchedule) validateUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, heliosBmrScheduleTypeUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HeliosBmrSchedule) validateUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.Unit) { // not required
		return nil
	}

	// value enum
	if err := m.validateUnitEnum("unit", "body", *m.Unit); err != nil {
		return err
	}

	return nil
}

func (m *HeliosBmrSchedule) validateWeekSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.WeekSchedule) { // not required
		return nil
	}

	if m.WeekSchedule != nil {
		if err := m.WeekSchedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weekSchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("weekSchedule")
			}
			return err
		}
	}

	return nil
}

func (m *HeliosBmrSchedule) validateMonthSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.MonthSchedule) { // not required
		return nil
	}

	if m.MonthSchedule != nil {
		if err := m.MonthSchedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monthSchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("monthSchedule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this helios bmr schedule based on the context it is used
func (m *HeliosBmrSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWeekSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMonthSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HeliosBmrSchedule) contextValidateWeekSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.WeekSchedule != nil {

		if swag.IsZero(m.WeekSchedule) { // not required
			return nil
		}

		if err := m.WeekSchedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weekSchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("weekSchedule")
			}
			return err
		}
	}

	return nil
}

func (m *HeliosBmrSchedule) contextValidateMonthSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.MonthSchedule != nil {

		if swag.IsZero(m.MonthSchedule) { // not required
			return nil
		}

		if err := m.MonthSchedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monthSchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("monthSchedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HeliosBmrSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HeliosBmrSchedule) UnmarshalBinary(b []byte) error {
	var res HeliosBmrSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
