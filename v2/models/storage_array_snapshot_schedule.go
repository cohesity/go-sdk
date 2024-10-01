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

// StorageArraySnapshotSchedule Storage Snap Mgmt Schedule
//
// Specifies settings that defines how frequent Storage Snapshot Management backup will be performed for a Protection Group.
//
// swagger:model StorageArraySnapshotSchedule
type StorageArraySnapshotSchedule struct {

	// Specifies how often to start new Protection Group Runs of a Protection Group. <br>'Minutes' specifies that Protection Group run starts periodically after certain number of minutes specified in 'frequency' field. <br>'Hours' specifies that Protection Group run starts periodically after certain number of hours specified in 'frequency' field.
	// Required: true
	// Enum: ["Minutes","Hours","Days","Weeks","Months","Years"]
	Unit *string `json:"unit"`

	// Specifies the days Schedule for Protection Group to start runs after certain number of minutes.
	MinuteSchedule *MinuteSchedule `json:"minuteSchedule,omitempty"`

	// Specifies the days Schedule for Protection Group to start runs after certain number of hours.
	HourSchedule *HourSchedule `json:"hourSchedule,omitempty"`

	// Specifies the days Schedule for Protection Group to start runs after certain number of days.
	DaySchedule *DaySchedule `json:"daySchedule,omitempty"`

	// Specifies the week Schedule for Protection Group to start runs on certain number of days in a week.
	WeekSchedule *WeekSchedule `json:"weekSchedule,omitempty"`

	// Specifies the week Schedule for Protection Group to start runs on specific week in a month and specific days of that week.
	MonthSchedule *MonthSchedule `json:"monthSchedule,omitempty"`

	// Specifies the year Schedule for Protection Group to start runs on specific day of that year.
	YearSchedule *YearSchedule `json:"yearSchedule,omitempty"`
}

// Validate validates this storage array snapshot schedule
func (m *StorageArraySnapshotSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinuteSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHourSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDaySchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeekSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonthSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYearSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var storageArraySnapshotScheduleTypeUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Minutes","Hours","Days","Weeks","Months","Years"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storageArraySnapshotScheduleTypeUnitPropEnum = append(storageArraySnapshotScheduleTypeUnitPropEnum, v)
	}
}

const (

	// StorageArraySnapshotScheduleUnitMinutes captures enum value "Minutes"
	StorageArraySnapshotScheduleUnitMinutes string = "Minutes"

	// StorageArraySnapshotScheduleUnitHours captures enum value "Hours"
	StorageArraySnapshotScheduleUnitHours string = "Hours"

	// StorageArraySnapshotScheduleUnitDays captures enum value "Days"
	StorageArraySnapshotScheduleUnitDays string = "Days"

	// StorageArraySnapshotScheduleUnitWeeks captures enum value "Weeks"
	StorageArraySnapshotScheduleUnitWeeks string = "Weeks"

	// StorageArraySnapshotScheduleUnitMonths captures enum value "Months"
	StorageArraySnapshotScheduleUnitMonths string = "Months"

	// StorageArraySnapshotScheduleUnitYears captures enum value "Years"
	StorageArraySnapshotScheduleUnitYears string = "Years"
)

// prop value enum
func (m *StorageArraySnapshotSchedule) validateUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, storageArraySnapshotScheduleTypeUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StorageArraySnapshotSchedule) validateUnit(formats strfmt.Registry) error {

	if err := validate.Required("unit", "body", m.Unit); err != nil {
		return err
	}

	// value enum
	if err := m.validateUnitEnum("unit", "body", *m.Unit); err != nil {
		return err
	}

	return nil
}

func (m *StorageArraySnapshotSchedule) validateMinuteSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.MinuteSchedule) { // not required
		return nil
	}

	if m.MinuteSchedule != nil {
		if err := m.MinuteSchedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("minuteSchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("minuteSchedule")
			}
			return err
		}
	}

	return nil
}

func (m *StorageArraySnapshotSchedule) validateHourSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.HourSchedule) { // not required
		return nil
	}

	if m.HourSchedule != nil {
		if err := m.HourSchedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hourSchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hourSchedule")
			}
			return err
		}
	}

	return nil
}

func (m *StorageArraySnapshotSchedule) validateDaySchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.DaySchedule) { // not required
		return nil
	}

	if m.DaySchedule != nil {
		if err := m.DaySchedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("daySchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("daySchedule")
			}
			return err
		}
	}

	return nil
}

func (m *StorageArraySnapshotSchedule) validateWeekSchedule(formats strfmt.Registry) error {
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

func (m *StorageArraySnapshotSchedule) validateMonthSchedule(formats strfmt.Registry) error {
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

func (m *StorageArraySnapshotSchedule) validateYearSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.YearSchedule) { // not required
		return nil
	}

	if m.YearSchedule != nil {
		if err := m.YearSchedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("yearSchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("yearSchedule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this storage array snapshot schedule based on the context it is used
func (m *StorageArraySnapshotSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMinuteSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHourSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDaySchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWeekSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMonthSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateYearSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageArraySnapshotSchedule) contextValidateMinuteSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.MinuteSchedule != nil {

		if swag.IsZero(m.MinuteSchedule) { // not required
			return nil
		}

		if err := m.MinuteSchedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("minuteSchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("minuteSchedule")
			}
			return err
		}
	}

	return nil
}

func (m *StorageArraySnapshotSchedule) contextValidateHourSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.HourSchedule != nil {

		if swag.IsZero(m.HourSchedule) { // not required
			return nil
		}

		if err := m.HourSchedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hourSchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hourSchedule")
			}
			return err
		}
	}

	return nil
}

func (m *StorageArraySnapshotSchedule) contextValidateDaySchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.DaySchedule != nil {

		if swag.IsZero(m.DaySchedule) { // not required
			return nil
		}

		if err := m.DaySchedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("daySchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("daySchedule")
			}
			return err
		}
	}

	return nil
}

func (m *StorageArraySnapshotSchedule) contextValidateWeekSchedule(ctx context.Context, formats strfmt.Registry) error {

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

func (m *StorageArraySnapshotSchedule) contextValidateMonthSchedule(ctx context.Context, formats strfmt.Registry) error {

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

func (m *StorageArraySnapshotSchedule) contextValidateYearSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.YearSchedule != nil {

		if swag.IsZero(m.YearSchedule) { // not required
			return nil
		}

		if err := m.YearSchedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("yearSchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("yearSchedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageArraySnapshotSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageArraySnapshotSchedule) UnmarshalBinary(b []byte) error {
	var res StorageArraySnapshotSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
