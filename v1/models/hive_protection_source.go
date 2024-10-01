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

// HiveProtectionSource Hive Protection Source.
//
// Specifies an Object representing Hive.
//
// swagger:model HiveProtectionSource
type HiveProtectionSource struct {

	// Specifies the instance name of the Hive entity.
	Name *string `json:"name,omitempty"`

	// Information of a Hive Table, only valid for an entity of type
	// kTable.
	TableInfo *HiveTable `json:"tableInfo,omitempty"`

	// Specifies the type of the managed Object in Hive Protection Source.
	// Specifies the type of an Hive source entity.
	// 'kCluster' indicates a Hive cluster distributed over several physical
	// nodes.
	// 'kDatabase' indicates a Database in the Hive environment.
	// 'kTable' indicates a Table in the Hive environment.
	// Enum: ["kCluster","kDatabase","kTable"]
	Type *string `json:"type,omitempty"`

	// Specifies the UUID for the Hive entity.
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this hive protection source
func (m *HiveProtectionSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTableInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HiveProtectionSource) validateTableInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.TableInfo) { // not required
		return nil
	}

	if m.TableInfo != nil {
		if err := m.TableInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tableInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tableInfo")
			}
			return err
		}
	}

	return nil
}

var hiveProtectionSourceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kCluster","kDatabase","kTable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hiveProtectionSourceTypeTypePropEnum = append(hiveProtectionSourceTypeTypePropEnum, v)
	}
}

const (

	// HiveProtectionSourceTypeKCluster captures enum value "kCluster"
	HiveProtectionSourceTypeKCluster string = "kCluster"

	// HiveProtectionSourceTypeKDatabase captures enum value "kDatabase"
	HiveProtectionSourceTypeKDatabase string = "kDatabase"

	// HiveProtectionSourceTypeKTable captures enum value "kTable"
	HiveProtectionSourceTypeKTable string = "kTable"
)

// prop value enum
func (m *HiveProtectionSource) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, hiveProtectionSourceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HiveProtectionSource) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this hive protection source based on the context it is used
func (m *HiveProtectionSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTableInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HiveProtectionSource) contextValidateTableInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.TableInfo != nil {

		if swag.IsZero(m.TableInfo) { // not required
			return nil
		}

		if err := m.TableInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tableInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tableInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HiveProtectionSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HiveProtectionSource) UnmarshalBinary(b []byte) error {
	var res HiveProtectionSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
