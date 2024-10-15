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

// FortKnoxAzureInfo FortKnoxAzureInfo holds information about the Fortknox
// Azure or Azure FreeTrial subscription such as if it is active.
//
// swagger:model FortKnoxAzureInfo
type FortKnoxAzureInfo struct {

	// Specifies the information regarding banner to display.
	Banner *EntitlementBannerInfo `json:"banner,omitempty"`

	// Specifies the end date of the subscription.
	EndDate *string `json:"endDate,omitempty"`

	// Specifies whether the subscription is active.
	IsActive *bool `json:"isActive,omitempty"`

	// Specifies whether the subscription is free trial.
	IsFreeTrial *bool `json:"isFreeTrial,omitempty"`

	// Display name of the Product
	ProductDisplayName *string `json:"productDisplayName,omitempty"`

	// Specifies the quantity of the subscription.
	Quantity *int64 `json:"quantity,omitempty"`

	// Specifies the start date of the subscription.
	StartDate *string `json:"startDate,omitempty"`
}

// Validate validates this fort knox azure info
func (m *FortKnoxAzureInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBanner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FortKnoxAzureInfo) validateBanner(formats strfmt.Registry) error {
	if swag.IsZero(m.Banner) { // not required
		return nil
	}

	if m.Banner != nil {
		if err := m.Banner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("banner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("banner")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fort knox azure info based on the context it is used
func (m *FortKnoxAzureInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBanner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FortKnoxAzureInfo) contextValidateBanner(ctx context.Context, formats strfmt.Registry) error {

	if m.Banner != nil {

		if swag.IsZero(m.Banner) { // not required
			return nil
		}

		if err := m.Banner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("banner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("banner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FortKnoxAzureInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FortKnoxAzureInfo) UnmarshalBinary(b []byte) error {
	var res FortKnoxAzureInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
