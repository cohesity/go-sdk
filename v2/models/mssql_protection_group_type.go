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

// MssqlProtectionGroupType MSSQL Protection Group type.
//
// MSSQL Protection Group type.
//
// swagger:model MssqlProtectionGroupType
type MssqlProtectionGroupType struct {

	// Specifies MSSQL Protection Group type.
	// Enum: ["kFile","kVolume","kNative"]
	Environment string `json:"environment,omitempty"`
}

// Validate validates this mssql protection group type
func (m *MssqlProtectionGroupType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var mssqlProtectionGroupTypeTypeEnvironmentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kFile","kVolume","kNative"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mssqlProtectionGroupTypeTypeEnvironmentPropEnum = append(mssqlProtectionGroupTypeTypeEnvironmentPropEnum, v)
	}
}

const (

	// MssqlProtectionGroupTypeEnvironmentKFile captures enum value "kFile"
	MssqlProtectionGroupTypeEnvironmentKFile string = "kFile"

	// MssqlProtectionGroupTypeEnvironmentKVolume captures enum value "kVolume"
	MssqlProtectionGroupTypeEnvironmentKVolume string = "kVolume"

	// MssqlProtectionGroupTypeEnvironmentKNative captures enum value "kNative"
	MssqlProtectionGroupTypeEnvironmentKNative string = "kNative"
)

// prop value enum
func (m *MssqlProtectionGroupType) validateEnvironmentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mssqlProtectionGroupTypeTypeEnvironmentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MssqlProtectionGroupType) validateEnvironment(formats strfmt.Registry) error {
	if swag.IsZero(m.Environment) { // not required
		return nil
	}

	// value enum
	if err := m.validateEnvironmentEnum("environment", "body", m.Environment); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mssql protection group type based on context it is used
func (m *MssqlProtectionGroupType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MssqlProtectionGroupType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MssqlProtectionGroupType) UnmarshalBinary(b []byte) error {
	var res MssqlProtectionGroupType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
