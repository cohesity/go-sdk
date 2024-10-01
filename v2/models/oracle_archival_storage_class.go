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

// OracleArchivalStorageClass Oracle Archival Storage Class
//
// Specifies the storage class of Oracle External Target of archival purpose type.
//
// swagger:model OracleArchivalStorageClass
type OracleArchivalStorageClass struct {

	// Specifies the storage class of Oracle External Target of archival purpose type.
	// Enum: ["OracleObjectStorage","OracleArchiveStorage"]
	Enum string `json:"enum,omitempty"`
}

// Validate validates this oracle archival storage class
func (m *OracleArchivalStorageClass) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnum(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var oracleArchivalStorageClassTypeEnumPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OracleObjectStorage","OracleArchiveStorage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oracleArchivalStorageClassTypeEnumPropEnum = append(oracleArchivalStorageClassTypeEnumPropEnum, v)
	}
}

const (

	// OracleArchivalStorageClassEnumOracleObjectStorage captures enum value "OracleObjectStorage"
	OracleArchivalStorageClassEnumOracleObjectStorage string = "OracleObjectStorage"

	// OracleArchivalStorageClassEnumOracleArchiveStorage captures enum value "OracleArchiveStorage"
	OracleArchivalStorageClassEnumOracleArchiveStorage string = "OracleArchiveStorage"
)

// prop value enum
func (m *OracleArchivalStorageClass) validateEnumEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, oracleArchivalStorageClassTypeEnumPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OracleArchivalStorageClass) validateEnum(formats strfmt.Registry) error {
	if swag.IsZero(m.Enum) { // not required
		return nil
	}

	// value enum
	if err := m.validateEnumEnum("enum", "body", m.Enum); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this oracle archival storage class based on context it is used
func (m *OracleArchivalStorageClass) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OracleArchivalStorageClass) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleArchivalStorageClass) UnmarshalBinary(b []byte) error {
	var res OracleArchivalStorageClass
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
