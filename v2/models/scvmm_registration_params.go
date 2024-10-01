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

// ScvmmRegistrationParams Register HyperV SCVMM request parameters.
//
// Specifies parameters to register HyperV SCVMM.
//
// swagger:model ScvmmRegistrationParams
type ScvmmRegistrationParams struct {
	CommonSourceRegistrationParams

	// Specifies the agent endpoint if it is different from the source endpoint.
	AgentEndpoint *string `json:"agentEndpoint,omitempty"`

	// Specifies the throttling params.
	ThrottlingParams *ThrottlingParams `json:"throttlingParams,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ScvmmRegistrationParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonSourceRegistrationParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonSourceRegistrationParams = aO0

	// AO1
	var dataAO1 struct {
		AgentEndpoint *string `json:"agentEndpoint,omitempty"`

		ThrottlingParams *ThrottlingParams `json:"throttlingParams,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AgentEndpoint = dataAO1.AgentEndpoint

	m.ThrottlingParams = dataAO1.ThrottlingParams

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ScvmmRegistrationParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonSourceRegistrationParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		AgentEndpoint *string `json:"agentEndpoint,omitempty"`

		ThrottlingParams *ThrottlingParams `json:"throttlingParams,omitempty"`
	}

	dataAO1.AgentEndpoint = m.AgentEndpoint

	dataAO1.ThrottlingParams = m.ThrottlingParams

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this scvmm registration params
func (m *ScvmmRegistrationParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonSourceRegistrationParams
	if err := m.CommonSourceRegistrationParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThrottlingParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScvmmRegistrationParams) validateThrottlingParams(formats strfmt.Registry) error {

	if swag.IsZero(m.ThrottlingParams) { // not required
		return nil
	}

	if m.ThrottlingParams != nil {
		if err := m.ThrottlingParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throttlingParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throttlingParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this scvmm registration params based on the context it is used
func (m *ScvmmRegistrationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonSourceRegistrationParams
	if err := m.CommonSourceRegistrationParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThrottlingParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScvmmRegistrationParams) contextValidateThrottlingParams(ctx context.Context, formats strfmt.Registry) error {

	if m.ThrottlingParams != nil {

		if swag.IsZero(m.ThrottlingParams) { // not required
			return nil
		}

		if err := m.ThrottlingParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throttlingParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throttlingParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScvmmRegistrationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScvmmRegistrationParams) UnmarshalBinary(b []byte) error {
	var res ScvmmRegistrationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
