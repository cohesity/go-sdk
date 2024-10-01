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

// ArchivalTargetTierInfo Archival Target Tier Info
//
// Specifies the tier info for archival.
//
// swagger:model ArchivalTargetTierInfo
type ArchivalTargetTierInfo struct {
	TierLevelSettings

	// Specifies the type of the current tier where the snapshot resides. This will be specified if the run is a CAD run.
	// Enum: ["kAmazonS3Standard","kAmazonS3StandardIA","kAmazonS3OneZoneIA","kAmazonS3IntelligentTiering","kAmazonS3Glacier","kAmazonS3GlacierDeepArchive","kAzureTierHot","kAzureTierCool","kAzureTierArchive","kGoogleStandard","kGoogleRegional","kGoogleMultiRegional","kGoogleNearline","kGoogleColdline","kOracleTierStandard","kOracleTierArchive"]
	CurrentTierType *string `json:"currentTierType,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ArchivalTargetTierInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TierLevelSettings
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TierLevelSettings = aO0

	// AO1
	var dataAO1 struct {
		CurrentTierType *string `json:"currentTierType,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.CurrentTierType = dataAO1.CurrentTierType

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ArchivalTargetTierInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.TierLevelSettings)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		CurrentTierType *string `json:"currentTierType,omitempty"`
	}

	dataAO1.CurrentTierType = m.CurrentTierType

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this archival target tier info
func (m *ArchivalTargetTierInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TierLevelSettings
	if err := m.TierLevelSettings.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentTierType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var archivalTargetTierInfoTypeCurrentTierTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kAmazonS3Standard","kAmazonS3StandardIA","kAmazonS3OneZoneIA","kAmazonS3IntelligentTiering","kAmazonS3Glacier","kAmazonS3GlacierDeepArchive","kAzureTierHot","kAzureTierCool","kAzureTierArchive","kGoogleStandard","kGoogleRegional","kGoogleMultiRegional","kGoogleNearline","kGoogleColdline","kOracleTierStandard","kOracleTierArchive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		archivalTargetTierInfoTypeCurrentTierTypePropEnum = append(archivalTargetTierInfoTypeCurrentTierTypePropEnum, v)
	}
}

// property enum
func (m *ArchivalTargetTierInfo) validateCurrentTierTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, archivalTargetTierInfoTypeCurrentTierTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ArchivalTargetTierInfo) validateCurrentTierType(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentTierType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCurrentTierTypeEnum("currentTierType", "body", *m.CurrentTierType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this archival target tier info based on the context it is used
func (m *ArchivalTargetTierInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TierLevelSettings
	if err := m.TierLevelSettings.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ArchivalTargetTierInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArchivalTargetTierInfo) UnmarshalBinary(b []byte) error {
	var res ArchivalTargetTierInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
