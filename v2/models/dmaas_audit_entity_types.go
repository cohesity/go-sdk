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

// DmaasAuditEntityTypes Dmaas Audit Entity Types
//
// # Dmaas Audit Entity Types
//
// swagger:model DmaasAuditEntityTypes
type DmaasAuditEntityTypes struct {

	// Dmaas Audit Entity Types
	// Enum: ["Region","SaasConnector","ProtectionPolicy"]
	EntityTypes string `json:"entityTypes,omitempty"`
}

// Validate validates this dmaas audit entity types
func (m *DmaasAuditEntityTypes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var dmaasAuditEntityTypesTypeEntityTypesPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Region","SaasConnector","ProtectionPolicy"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dmaasAuditEntityTypesTypeEntityTypesPropEnum = append(dmaasAuditEntityTypesTypeEntityTypesPropEnum, v)
	}
}

const (

	// DmaasAuditEntityTypesEntityTypesRegion captures enum value "Region"
	DmaasAuditEntityTypesEntityTypesRegion string = "Region"

	// DmaasAuditEntityTypesEntityTypesSaasConnector captures enum value "SaasConnector"
	DmaasAuditEntityTypesEntityTypesSaasConnector string = "SaasConnector"

	// DmaasAuditEntityTypesEntityTypesProtectionPolicy captures enum value "ProtectionPolicy"
	DmaasAuditEntityTypesEntityTypesProtectionPolicy string = "ProtectionPolicy"
)

// prop value enum
func (m *DmaasAuditEntityTypes) validateEntityTypesEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dmaasAuditEntityTypesTypeEntityTypesPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DmaasAuditEntityTypes) validateEntityTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityTypes) { // not required
		return nil
	}

	// value enum
	if err := m.validateEntityTypesEnum("entityTypes", "body", m.EntityTypes); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this dmaas audit entity types based on context it is used
func (m *DmaasAuditEntityTypes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DmaasAuditEntityTypes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DmaasAuditEntityTypes) UnmarshalBinary(b []byte) error {
	var res DmaasAuditEntityTypes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
