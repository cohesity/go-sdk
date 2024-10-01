// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HeliosAwsCloudSpinParams AWS Parameters.
//
// Specifies various resources when converting and deploying a VM to AWS.
//
// swagger:model HeliosAwsCloudSpinParams
type HeliosAwsCloudSpinParams struct {

	// Specifies id of the AWS region in which to deploy the VM.
	// Required: true
	Region *int64 `json:"region"`

	// Specifies id of the Virtual Private Cloud to chose for the instance type.
	VpcID *int64 `json:"vpcId,omitempty"`

	// Specifies id of the subnet within above VPC.
	SubnetID *int64 `json:"subnetId,omitempty"`
}

// Validate validates this helios aws cloud spin params
func (m *HeliosAwsCloudSpinParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HeliosAwsCloudSpinParams) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this helios aws cloud spin params based on context it is used
func (m *HeliosAwsCloudSpinParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HeliosAwsCloudSpinParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HeliosAwsCloudSpinParams) UnmarshalBinary(b []byte) error {
	var res HeliosAwsCloudSpinParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
