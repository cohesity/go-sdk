// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HeliosWeekSchedule Week Schedule
//
// Specifies settings that define a schedule for a Protection Group runs to start on certain days of week.
//
// swagger:model HeliosWeekSchedule
type HeliosWeekSchedule struct {

	// Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday']
	DayOfWeek []string `json:"dayOfWeek"`
}

// Validate validates this helios week schedule
func (m *HeliosWeekSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDayOfWeek(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var heliosWeekScheduleDayOfWeekItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Sunday","Monday","Tuesday","Wednesday","Thursday","Friday","Saturday"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		heliosWeekScheduleDayOfWeekItemsEnum = append(heliosWeekScheduleDayOfWeekItemsEnum, v)
	}
}

func (m *HeliosWeekSchedule) validateDayOfWeekItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, heliosWeekScheduleDayOfWeekItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HeliosWeekSchedule) validateDayOfWeek(formats strfmt.Registry) error {
	if swag.IsZero(m.DayOfWeek) { // not required
		return nil
	}

	for i := 0; i < len(m.DayOfWeek); i++ {

		// value enum
		if err := m.validateDayOfWeekItemsEnum("dayOfWeek"+"."+strconv.Itoa(i), "body", m.DayOfWeek[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this helios week schedule based on context it is used
func (m *HeliosWeekSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HeliosWeekSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HeliosWeekSchedule) UnmarshalBinary(b []byte) error {
	var res HeliosWeekSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
