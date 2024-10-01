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

// RecoverPureParams Recover Pure Params.
//
// Specifies the recovery options specific to Pure environment.
//
// swagger:model RecoverPureParams
type RecoverPureParams struct {

	// Specifies the list of recover object parameters.
	// Required: true
	Objects []*CommonRecoverObjectSnapshotParams `json:"objects"`

	// Specifies the type of recovery action to be performed. The corresponding recovery action params must be filled out.
	// Required: true
	// Enum: ["RecoverSanVolumes","RecoverSanGroup"]
	RecoveryAction *string `json:"recoveryAction"`

	// Specifies the parameters to recover SAN Pure Protection Group.
	RecoverSanGroupParams *RecoverPureSanGroupParams `json:"recoverSanGroupParams,omitempty"`

	// Specifies the parameters to recover SAN Volume.
	RecoverSanVolumeParams *RecoverPureSanVolumeParams `json:"recoverSanVolumeParams,omitempty"`
}

// Validate validates this recover pure params
func (m *RecoverPureParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoveryAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoverSanGroupParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoverSanVolumeParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverPureParams) validateObjects(formats strfmt.Registry) error {

	if err := validate.Required("objects", "body", m.Objects); err != nil {
		return err
	}

	for i := 0; i < len(m.Objects); i++ {
		if swag.IsZero(m.Objects[i]) { // not required
			continue
		}

		if m.Objects[i] != nil {
			if err := m.Objects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var recoverPureParamsTypeRecoveryActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RecoverSanVolumes","RecoverSanGroup"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recoverPureParamsTypeRecoveryActionPropEnum = append(recoverPureParamsTypeRecoveryActionPropEnum, v)
	}
}

const (

	// RecoverPureParamsRecoveryActionRecoverSanVolumes captures enum value "RecoverSanVolumes"
	RecoverPureParamsRecoveryActionRecoverSanVolumes string = "RecoverSanVolumes"

	// RecoverPureParamsRecoveryActionRecoverSanGroup captures enum value "RecoverSanGroup"
	RecoverPureParamsRecoveryActionRecoverSanGroup string = "RecoverSanGroup"
)

// prop value enum
func (m *RecoverPureParams) validateRecoveryActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recoverPureParamsTypeRecoveryActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecoverPureParams) validateRecoveryAction(formats strfmt.Registry) error {

	if err := validate.Required("recoveryAction", "body", m.RecoveryAction); err != nil {
		return err
	}

	// value enum
	if err := m.validateRecoveryActionEnum("recoveryAction", "body", *m.RecoveryAction); err != nil {
		return err
	}

	return nil
}

func (m *RecoverPureParams) validateRecoverSanGroupParams(formats strfmt.Registry) error {
	if swag.IsZero(m.RecoverSanGroupParams) { // not required
		return nil
	}

	if m.RecoverSanGroupParams != nil {
		if err := m.RecoverSanGroupParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverSanGroupParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverSanGroupParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverPureParams) validateRecoverSanVolumeParams(formats strfmt.Registry) error {
	if swag.IsZero(m.RecoverSanVolumeParams) { // not required
		return nil
	}

	if m.RecoverSanVolumeParams != nil {
		if err := m.RecoverSanVolumeParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverSanVolumeParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverSanVolumeParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover pure params based on the context it is used
func (m *RecoverPureParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecoverSanGroupParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecoverSanVolumeParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverPureParams) contextValidateObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Objects); i++ {

		if m.Objects[i] != nil {

			if swag.IsZero(m.Objects[i]) { // not required
				return nil
			}

			if err := m.Objects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecoverPureParams) contextValidateRecoverSanGroupParams(ctx context.Context, formats strfmt.Registry) error {

	if m.RecoverSanGroupParams != nil {

		if swag.IsZero(m.RecoverSanGroupParams) { // not required
			return nil
		}

		if err := m.RecoverSanGroupParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverSanGroupParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverSanGroupParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverPureParams) contextValidateRecoverSanVolumeParams(ctx context.Context, formats strfmt.Registry) error {

	if m.RecoverSanVolumeParams != nil {

		if swag.IsZero(m.RecoverSanVolumeParams) { // not required
			return nil
		}

		if err := m.RecoverSanVolumeParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverSanVolumeParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverSanVolumeParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverPureParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverPureParams) UnmarshalBinary(b []byte) error {
	var res RecoverPureParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
