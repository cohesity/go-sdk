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

// HeliosAWSTier Specifies the settings for a aws tier.
//
// swagger:model HeliosAWSTier
type HeliosAWSTier struct {
	HeliosCommonCloudTierSettings

	// Specifies the AWS tier types.
	// Required: true
	// Enum: ["kAmazonS3Standard","kAmazonS3StandardIA","kAmazonS3OneZoneIA","kAmazonS3IntelligentTiering","kAmazonS3Glacier","kAmazonS3GlacierDeepArchive"]
	TierType *string `json:"tierType"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HeliosAWSTier) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 HeliosCommonCloudTierSettings
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.HeliosCommonCloudTierSettings = aO0

	// AO1
	var dataAO1 struct {
		TierType *string `json:"tierType"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.TierType = dataAO1.TierType

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HeliosAWSTier) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.HeliosCommonCloudTierSettings)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		TierType *string `json:"tierType"`
	}

	dataAO1.TierType = m.TierType

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this helios a w s tier
func (m *HeliosAWSTier) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with HeliosCommonCloudTierSettings
	if err := m.HeliosCommonCloudTierSettings.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTierType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var heliosAWSTierTypeTierTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kAmazonS3Standard","kAmazonS3StandardIA","kAmazonS3OneZoneIA","kAmazonS3IntelligentTiering","kAmazonS3Glacier","kAmazonS3GlacierDeepArchive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		heliosAWSTierTypeTierTypePropEnum = append(heliosAWSTierTypeTierTypePropEnum, v)
	}
}

// property enum
func (m *HeliosAWSTier) validateTierTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, heliosAWSTierTypeTierTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HeliosAWSTier) validateTierType(formats strfmt.Registry) error {

	if err := validate.Required("tierType", "body", m.TierType); err != nil {
		return err
	}

	// value enum
	if err := m.validateTierTypeEnum("tierType", "body", *m.TierType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this helios a w s tier based on the context it is used
func (m *HeliosAWSTier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with HeliosCommonCloudTierSettings
	if err := m.HeliosCommonCloudTierSettings.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *HeliosAWSTier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HeliosAWSTier) UnmarshalBinary(b []byte) error {
	var res HeliosAWSTier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
