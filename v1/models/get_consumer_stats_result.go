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

// GetConsumerStatsResult GetConsumerStatsResult
//
// GetConsumerStatsResult is the result of get consumerStats api.
//
// swagger:model GetConsumerStatsResult
type GetConsumerStatsResult struct {

	// Specifies an opaque string to pass to get the next set of active opens.
	// If null is returned, this response is the last set of active opens.
	Cookie *string `json:"cookie,omitempty"`

	// Specifies a list of consumer stats.
	StatsList []*ConsumerStats `json:"statsList"`
}

// Validate validates this get consumer stats result
func (m *GetConsumerStatsResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatsList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetConsumerStatsResult) validateStatsList(formats strfmt.Registry) error {
	if swag.IsZero(m.StatsList) { // not required
		return nil
	}

	for i := 0; i < len(m.StatsList); i++ {
		if swag.IsZero(m.StatsList[i]) { // not required
			continue
		}

		if m.StatsList[i] != nil {
			if err := m.StatsList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statsList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("statsList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get consumer stats result based on the context it is used
func (m *GetConsumerStatsResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatsList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetConsumerStatsResult) contextValidateStatsList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatsList); i++ {

		if m.StatsList[i] != nil {

			if swag.IsZero(m.StatsList[i]) { // not required
				return nil
			}

			if err := m.StatsList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statsList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("statsList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetConsumerStatsResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetConsumerStatsResult) UnmarshalBinary(b []byte) error {
	var res GetConsumerStatsResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
