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

// VmwareDiskControllerType VMware disk controller type.
//
// Vmware disk controller type.
//
// swagger:model VmwareDiskControllerType
type VmwareDiskControllerType struct {

	// Specifies VMware disk controller type.
	// Enum: ["kScsi","kIde","kSata"]
	VmwareDiskControllerType string `json:"vmwareDiskControllerType,omitempty"`
}

// Validate validates this vmware disk controller type
func (m *VmwareDiskControllerType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVmwareDiskControllerType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var vmwareDiskControllerTypeTypeVmwareDiskControllerTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kScsi","kIde","kSata"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmwareDiskControllerTypeTypeVmwareDiskControllerTypePropEnum = append(vmwareDiskControllerTypeTypeVmwareDiskControllerTypePropEnum, v)
	}
}

const (

	// VmwareDiskControllerTypeVmwareDiskControllerTypeKScsi captures enum value "kScsi"
	VmwareDiskControllerTypeVmwareDiskControllerTypeKScsi string = "kScsi"

	// VmwareDiskControllerTypeVmwareDiskControllerTypeKIde captures enum value "kIde"
	VmwareDiskControllerTypeVmwareDiskControllerTypeKIde string = "kIde"

	// VmwareDiskControllerTypeVmwareDiskControllerTypeKSata captures enum value "kSata"
	VmwareDiskControllerTypeVmwareDiskControllerTypeKSata string = "kSata"
)

// prop value enum
func (m *VmwareDiskControllerType) validateVmwareDiskControllerTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vmwareDiskControllerTypeTypeVmwareDiskControllerTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VmwareDiskControllerType) validateVmwareDiskControllerType(formats strfmt.Registry) error {
	if swag.IsZero(m.VmwareDiskControllerType) { // not required
		return nil
	}

	// value enum
	if err := m.validateVmwareDiskControllerTypeEnum("vmwareDiskControllerType", "body", m.VmwareDiskControllerType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vmware disk controller type based on context it is used
func (m *VmwareDiskControllerType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VmwareDiskControllerType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareDiskControllerType) UnmarshalBinary(b []byte) error {
	var res VmwareDiskControllerType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
