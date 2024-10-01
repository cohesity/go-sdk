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

// FullScheduleUnit FullScheduleUnit
//
// # Full Schedule Units
//
// swagger:model FullScheduleUnit
type FullScheduleUnit struct {

	// Specifies the full schedule unit (including ProtectOnce policy).
	// Enum: ["Days","Weeks","Months","Years","ProtectOnce"]
	FullScheduleUnit string `json:"fullScheduleUnit,omitempty"`
}

// Validate validates this full schedule unit
func (m *FullScheduleUnit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFullScheduleUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var fullScheduleUnitTypeFullScheduleUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Days","Weeks","Months","Years","ProtectOnce"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fullScheduleUnitTypeFullScheduleUnitPropEnum = append(fullScheduleUnitTypeFullScheduleUnitPropEnum, v)
	}
}

const (

	// FullScheduleUnitFullScheduleUnitDays captures enum value "Days"
	FullScheduleUnitFullScheduleUnitDays string = "Days"

	// FullScheduleUnitFullScheduleUnitWeeks captures enum value "Weeks"
	FullScheduleUnitFullScheduleUnitWeeks string = "Weeks"

	// FullScheduleUnitFullScheduleUnitMonths captures enum value "Months"
	FullScheduleUnitFullScheduleUnitMonths string = "Months"

	// FullScheduleUnitFullScheduleUnitYears captures enum value "Years"
	FullScheduleUnitFullScheduleUnitYears string = "Years"

	// FullScheduleUnitFullScheduleUnitProtectOnce captures enum value "ProtectOnce"
	FullScheduleUnitFullScheduleUnitProtectOnce string = "ProtectOnce"
)

// prop value enum
func (m *FullScheduleUnit) validateFullScheduleUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fullScheduleUnitTypeFullScheduleUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FullScheduleUnit) validateFullScheduleUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.FullScheduleUnit) { // not required
		return nil
	}

	// value enum
	if err := m.validateFullScheduleUnitEnum("fullScheduleUnit", "body", m.FullScheduleUnit); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this full schedule unit based on context it is used
func (m *FullScheduleUnit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FullScheduleUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FullScheduleUnit) UnmarshalBinary(b []byte) error {
	var res FullScheduleUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
