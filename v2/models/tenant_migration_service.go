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

// TenantMigrationService Tenant
//
// Describes a cluser service participating in tenant migration.
//
// swagger:model TenantMigrationService
type TenantMigrationService struct {

	// Specifies the cluster service on which this action needs to be performed.
	// Enum: ["kMagneto","kYoda","kIcebox"]
	Service *string `json:"service,omitempty"`

	// List of Actions.
	Actions []*TenantMigrationServiceAction `json:"actions"`
}

// Validate validates this tenant migration service
func (m *TenantMigrationService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var tenantMigrationServiceTypeServicePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kMagneto","kYoda","kIcebox"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tenantMigrationServiceTypeServicePropEnum = append(tenantMigrationServiceTypeServicePropEnum, v)
	}
}

const (

	// TenantMigrationServiceServiceKMagneto captures enum value "kMagneto"
	TenantMigrationServiceServiceKMagneto string = "kMagneto"

	// TenantMigrationServiceServiceKYoda captures enum value "kYoda"
	TenantMigrationServiceServiceKYoda string = "kYoda"

	// TenantMigrationServiceServiceKIcebox captures enum value "kIcebox"
	TenantMigrationServiceServiceKIcebox string = "kIcebox"
)

// prop value enum
func (m *TenantMigrationService) validateServiceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tenantMigrationServiceTypeServicePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TenantMigrationService) validateService(formats strfmt.Registry) error {
	if swag.IsZero(m.Service) { // not required
		return nil
	}

	// value enum
	if err := m.validateServiceEnum("service", "body", *m.Service); err != nil {
		return err
	}

	return nil
}

func (m *TenantMigrationService) validateActions(formats strfmt.Registry) error {
	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	for i := 0; i < len(m.Actions); i++ {
		if swag.IsZero(m.Actions[i]) { // not required
			continue
		}

		if m.Actions[i] != nil {
			if err := m.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tenant migration service based on the context it is used
func (m *TenantMigrationService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TenantMigrationService) contextValidateActions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Actions); i++ {

		if m.Actions[i] != nil {

			if swag.IsZero(m.Actions[i]) { // not required
				return nil
			}

			if err := m.Actions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TenantMigrationService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TenantMigrationService) UnmarshalBinary(b []byte) error {
	var res TenantMigrationService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
