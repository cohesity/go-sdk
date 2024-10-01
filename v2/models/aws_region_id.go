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

// AwsRegionID Aws Region Id
//
// Specifies the AWS region id.
//
// swagger:model AwsRegionId
type AwsRegionID struct {

	// Specifies the AWS region id.
	// Enum: ["us-east-1","us-east-2","us-west-1","us-west-2","ca-central-1","ap-northeast-1","ap-south-1","ap-southeast-1","ap-southeast-2","eu-central-1","eu-west-2","me-south-1","eu-west-3"]
	AwsRegionID string `json:"awsRegionId,omitempty"`
}

// Validate validates this aws region Id
func (m *AwsRegionID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsRegionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var awsRegionIdTypeAwsRegionIDPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["us-east-1","us-east-2","us-west-1","us-west-2","ca-central-1","ap-northeast-1","ap-south-1","ap-southeast-1","ap-southeast-2","eu-central-1","eu-west-2","me-south-1","eu-west-3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		awsRegionIdTypeAwsRegionIDPropEnum = append(awsRegionIdTypeAwsRegionIDPropEnum, v)
	}
}

const (

	// AwsRegionIDAwsRegionIDUsDashEastDash1 captures enum value "us-east-1"
	AwsRegionIDAwsRegionIDUsDashEastDash1 string = "us-east-1"

	// AwsRegionIDAwsRegionIDUsDashEastDash2 captures enum value "us-east-2"
	AwsRegionIDAwsRegionIDUsDashEastDash2 string = "us-east-2"

	// AwsRegionIDAwsRegionIDUsDashWestDash1 captures enum value "us-west-1"
	AwsRegionIDAwsRegionIDUsDashWestDash1 string = "us-west-1"

	// AwsRegionIDAwsRegionIDUsDashWestDash2 captures enum value "us-west-2"
	AwsRegionIDAwsRegionIDUsDashWestDash2 string = "us-west-2"

	// AwsRegionIDAwsRegionIDCaDashCentralDash1 captures enum value "ca-central-1"
	AwsRegionIDAwsRegionIDCaDashCentralDash1 string = "ca-central-1"

	// AwsRegionIDAwsRegionIDApDashNortheastDash1 captures enum value "ap-northeast-1"
	AwsRegionIDAwsRegionIDApDashNortheastDash1 string = "ap-northeast-1"

	// AwsRegionIDAwsRegionIDApDashSouthDash1 captures enum value "ap-south-1"
	AwsRegionIDAwsRegionIDApDashSouthDash1 string = "ap-south-1"

	// AwsRegionIDAwsRegionIDApDashSoutheastDash1 captures enum value "ap-southeast-1"
	AwsRegionIDAwsRegionIDApDashSoutheastDash1 string = "ap-southeast-1"

	// AwsRegionIDAwsRegionIDApDashSoutheastDash2 captures enum value "ap-southeast-2"
	AwsRegionIDAwsRegionIDApDashSoutheastDash2 string = "ap-southeast-2"

	// AwsRegionIDAwsRegionIDEuDashCentralDash1 captures enum value "eu-central-1"
	AwsRegionIDAwsRegionIDEuDashCentralDash1 string = "eu-central-1"

	// AwsRegionIDAwsRegionIDEuDashWestDash2 captures enum value "eu-west-2"
	AwsRegionIDAwsRegionIDEuDashWestDash2 string = "eu-west-2"

	// AwsRegionIDAwsRegionIDMeDashSouthDash1 captures enum value "me-south-1"
	AwsRegionIDAwsRegionIDMeDashSouthDash1 string = "me-south-1"

	// AwsRegionIDAwsRegionIDEuDashWestDash3 captures enum value "eu-west-3"
	AwsRegionIDAwsRegionIDEuDashWestDash3 string = "eu-west-3"
)

// prop value enum
func (m *AwsRegionID) validateAwsRegionIDEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, awsRegionIdTypeAwsRegionIDPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AwsRegionID) validateAwsRegionID(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsRegionID) { // not required
		return nil
	}

	// value enum
	if err := m.validateAwsRegionIDEnum("awsRegionId", "body", m.AwsRegionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this aws region Id based on context it is used
func (m *AwsRegionID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AwsRegionID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsRegionID) UnmarshalBinary(b []byte) error {
	var res AwsRegionID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
