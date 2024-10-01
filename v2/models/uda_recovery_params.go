// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UdaRecoveryParams Parameters to customize recovery workflow.
//
// swagger:model UdaRecoveryParams
type UdaRecoveryParams struct {

	// dynamic form
	DynamicForm *UdaDynamicFormParams `json:"dynamicForm,omitempty"`

	// primary fields
	PrimaryFields *UdaRecoveryParamsPrimaryFields `json:"primaryFields,omitempty"`
}

// Validate validates this uda recovery params
func (m *UdaRecoveryParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDynamicForm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryFields(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UdaRecoveryParams) validateDynamicForm(formats strfmt.Registry) error {
	if swag.IsZero(m.DynamicForm) { // not required
		return nil
	}

	if m.DynamicForm != nil {
		if err := m.DynamicForm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dynamicForm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dynamicForm")
			}
			return err
		}
	}

	return nil
}

func (m *UdaRecoveryParams) validatePrimaryFields(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimaryFields) { // not required
		return nil
	}

	if m.PrimaryFields != nil {
		if err := m.PrimaryFields.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primaryFields")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primaryFields")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this uda recovery params based on the context it is used
func (m *UdaRecoveryParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDynamicForm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UdaRecoveryParams) contextValidateDynamicForm(ctx context.Context, formats strfmt.Registry) error {

	if m.DynamicForm != nil {

		if swag.IsZero(m.DynamicForm) { // not required
			return nil
		}

		if err := m.DynamicForm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dynamicForm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dynamicForm")
			}
			return err
		}
	}

	return nil
}

func (m *UdaRecoveryParams) contextValidatePrimaryFields(ctx context.Context, formats strfmt.Registry) error {

	if m.PrimaryFields != nil {

		if swag.IsZero(m.PrimaryFields) { // not required
			return nil
		}

		if err := m.PrimaryFields.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primaryFields")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primaryFields")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UdaRecoveryParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UdaRecoveryParams) UnmarshalBinary(b []byte) error {
	var res UdaRecoveryParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UdaRecoveryParamsPrimaryFields Parameters to customize existing/ default form fields.
//
// swagger:model UdaRecoveryParamsPrimaryFields
type UdaRecoveryParamsPrimaryFields struct {

	// recovery args
	RecoveryArgs *UdaRecoveryParamsPrimaryFieldsRecoveryArgs `json:"recoveryArgs,omitempty"`
}

// Validate validates this uda recovery params primary fields
func (m *UdaRecoveryParamsPrimaryFields) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecoveryArgs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UdaRecoveryParamsPrimaryFields) validateRecoveryArgs(formats strfmt.Registry) error {
	if swag.IsZero(m.RecoveryArgs) { // not required
		return nil
	}

	if m.RecoveryArgs != nil {
		if err := m.RecoveryArgs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primaryFields" + "." + "recoveryArgs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primaryFields" + "." + "recoveryArgs")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this uda recovery params primary fields based on the context it is used
func (m *UdaRecoveryParamsPrimaryFields) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRecoveryArgs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UdaRecoveryParamsPrimaryFields) contextValidateRecoveryArgs(ctx context.Context, formats strfmt.Registry) error {

	if m.RecoveryArgs != nil {

		if swag.IsZero(m.RecoveryArgs) { // not required
			return nil
		}

		if err := m.RecoveryArgs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primaryFields" + "." + "recoveryArgs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primaryFields" + "." + "recoveryArgs")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UdaRecoveryParamsPrimaryFields) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UdaRecoveryParamsPrimaryFields) UnmarshalBinary(b []byte) error {
	var res UdaRecoveryParamsPrimaryFields
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UdaRecoveryParamsPrimaryFieldsRecoveryArgs Specifies custom arguments to be supplied to the source recovery.
//
// swagger:model UdaRecoveryParamsPrimaryFieldsRecoveryArgs
type UdaRecoveryParamsPrimaryFieldsRecoveryArgs struct {

	// Default value for the field.
	DefaultValue *string `json:"defaultValue,omitempty"`

	// Whether the field will be only a readonly field on the UI.
	Readonly *bool `json:"readonly,omitempty"`

	// Whether the field is mandatory.
	Required *bool `json:"required,omitempty"`

	// Specifies if the field will be hidden on the UI screen or not.
	Hidden *bool `json:"hidden,omitempty"`
}

// Validate validates this uda recovery params primary fields recovery args
func (m *UdaRecoveryParamsPrimaryFieldsRecoveryArgs) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this uda recovery params primary fields recovery args based on context it is used
func (m *UdaRecoveryParamsPrimaryFieldsRecoveryArgs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UdaRecoveryParamsPrimaryFieldsRecoveryArgs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UdaRecoveryParamsPrimaryFieldsRecoveryArgs) UnmarshalBinary(b []byte) error {
	var res UdaRecoveryParamsPrimaryFieldsRecoveryArgs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
