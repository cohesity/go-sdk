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

// UpdateFleetEnvInfoRequest Update Fleet Env Info Request.
//
// Specifies the params to add fleet env info.
//
// swagger:model UpdateFleetEnvInfoRequest
type UpdateFleetEnvInfoRequest struct {

	// Specifies the IAM role used to create instances.This field is now deprecated, use awsFleetInfo instead.
	IamRole *string `json:"iamRole,omitempty"`

	// Specifies the Region of the CE cluster. This field is now deprecated, use awsFleetInfo instead.
	Region *string `json:"region,omitempty"`

	// Specifies the VPC of the CE cluster.This field is now deprecated, use awsFleetInfo instead.
	VpcID *string `json:"vpcId,omitempty"`

	// Specifies the Subnet of the CE cluster.This field is now deprecated, use awsFleetInfo instead.
	SubnetID *string `json:"subnetId,omitempty"`

	// Specifies the security group Id.This field is now deprecated, use awsFleetInfo instead.
	SecurityGroupID *string `json:"securityGroupId,omitempty"`

	// aws fleet info
	AwsFleetInfo *AwsFleetInfo `json:"awsFleetInfo,omitempty"`

	// azure fleet info
	AzureFleetInfo *AzureFleetInfo `json:"azureFleetInfo,omitempty"`
}

// Validate validates this update fleet env info request
func (m *UpdateFleetEnvInfoRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsFleetInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureFleetInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateFleetEnvInfoRequest) validateAwsFleetInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsFleetInfo) { // not required
		return nil
	}

	if m.AwsFleetInfo != nil {
		if err := m.AwsFleetInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsFleetInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsFleetInfo")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateFleetEnvInfoRequest) validateAzureFleetInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureFleetInfo) { // not required
		return nil
	}

	if m.AzureFleetInfo != nil {
		if err := m.AzureFleetInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureFleetInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureFleetInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update fleet env info request based on the context it is used
func (m *UpdateFleetEnvInfoRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsFleetInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureFleetInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateFleetEnvInfoRequest) contextValidateAwsFleetInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsFleetInfo != nil {

		if swag.IsZero(m.AwsFleetInfo) { // not required
			return nil
		}

		if err := m.AwsFleetInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsFleetInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsFleetInfo")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateFleetEnvInfoRequest) contextValidateAzureFleetInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureFleetInfo != nil {

		if swag.IsZero(m.AzureFleetInfo) { // not required
			return nil
		}

		if err := m.AzureFleetInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureFleetInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureFleetInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateFleetEnvInfoRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateFleetEnvInfoRequest) UnmarshalBinary(b []byte) error {
	var res UpdateFleetEnvInfoRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
