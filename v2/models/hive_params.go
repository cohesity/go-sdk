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

// HiveParams Recover Hive environment params.
//
// Specifies the recovery options specific to Hive environment.
//
// swagger:model HiveParams
type HiveParams struct {

	// Specifies the type of recover action to be performed.
	// Required: true
	// Enum: ["RecoverObjects"]
	RecoveryAction *string `json:"recoveryAction"`

	// Specifies the parameters to recover Hive objects.
	// Required: true
	RecoverHiveParams *RecoverHiveParams `json:"recoverHiveParams"`
}

// Validate validates this hive params
func (m *HiveParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecoveryAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoverHiveParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var hiveParamsTypeRecoveryActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RecoverObjects"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hiveParamsTypeRecoveryActionPropEnum = append(hiveParamsTypeRecoveryActionPropEnum, v)
	}
}

const (

	// HiveParamsRecoveryActionRecoverObjects captures enum value "RecoverObjects"
	HiveParamsRecoveryActionRecoverObjects string = "RecoverObjects"
)

// prop value enum
func (m *HiveParams) validateRecoveryActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, hiveParamsTypeRecoveryActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HiveParams) validateRecoveryAction(formats strfmt.Registry) error {

	if err := validate.Required("recoveryAction", "body", m.RecoveryAction); err != nil {
		return err
	}

	// value enum
	if err := m.validateRecoveryActionEnum("recoveryAction", "body", *m.RecoveryAction); err != nil {
		return err
	}

	return nil
}

func (m *HiveParams) validateRecoverHiveParams(formats strfmt.Registry) error {

	if err := validate.Required("recoverHiveParams", "body", m.RecoverHiveParams); err != nil {
		return err
	}

	if m.RecoverHiveParams != nil {
		if err := m.RecoverHiveParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverHiveParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverHiveParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hive params based on the context it is used
func (m *HiveParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRecoverHiveParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HiveParams) contextValidateRecoverHiveParams(ctx context.Context, formats strfmt.Registry) error {

	if m.RecoverHiveParams != nil {

		if err := m.RecoverHiveParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverHiveParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverHiveParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HiveParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HiveParams) UnmarshalBinary(b []byte) error {
	var res HiveParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
