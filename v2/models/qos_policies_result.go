// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QosPoliciesResult Get QoS policies result.
//
// Specifies the list of QoS policies.
//
// swagger:model QosPoliciesResult
type QosPoliciesResult struct {

	// Specifies the list of QoS policies.
	QosPolicies []*QoS `json:"qosPolicies"`
}

// Validate validates this qos policies result
func (m *QosPoliciesResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQosPolicies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QosPoliciesResult) validateQosPolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.QosPolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.QosPolicies); i++ {
		if swag.IsZero(m.QosPolicies[i]) { // not required
			continue
		}

		if m.QosPolicies[i] != nil {
			if err := m.QosPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("qosPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("qosPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this qos policies result based on the context it is used
func (m *QosPoliciesResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQosPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QosPoliciesResult) contextValidateQosPolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.QosPolicies); i++ {

		if m.QosPolicies[i] != nil {

			if swag.IsZero(m.QosPolicies[i]) { // not required
				return nil
			}

			if err := m.QosPolicies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("qosPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("qosPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *QosPoliciesResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QosPoliciesResult) UnmarshalBinary(b []byte) error {
	var res QosPoliciesResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
