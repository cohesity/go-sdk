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

// McmTenantType Tenant Status
//
// Describes the status of a Tenant on a Cluster or Helios.
//
// swagger:model McmTenantType
type McmTenantType struct {

	// Specifies the Tenant status
	// Enum: ["Dmaas","Mcm"]
	Enum string `json:"enum,omitempty"`
}

// Validate validates this mcm tenant type
func (m *McmTenantType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnum(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var mcmTenantTypeTypeEnumPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Dmaas","Mcm"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mcmTenantTypeTypeEnumPropEnum = append(mcmTenantTypeTypeEnumPropEnum, v)
	}
}

const (

	// McmTenantTypeEnumDmaas captures enum value "Dmaas"
	McmTenantTypeEnumDmaas string = "Dmaas"

	// McmTenantTypeEnumMcm captures enum value "Mcm"
	McmTenantTypeEnumMcm string = "Mcm"
)

// prop value enum
func (m *McmTenantType) validateEnumEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mcmTenantTypeTypeEnumPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *McmTenantType) validateEnum(formats strfmt.Registry) error {
	if swag.IsZero(m.Enum) { // not required
		return nil
	}

	// value enum
	if err := m.validateEnumEnum("enum", "body", m.Enum); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mcm tenant type based on context it is used
func (m *McmTenantType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *McmTenantType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *McmTenantType) UnmarshalBinary(b []byte) error {
	var res McmTenantType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
