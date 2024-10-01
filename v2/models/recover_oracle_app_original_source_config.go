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

// RecoverOracleAppOriginalSourceConfig Recover Oracle App Original Source Config.
//
// Specifies the additional Source configuration parameters when databases will be recovered to original location.
//
// swagger:model RecoverOracleAppOriginalSourceConfig
type RecoverOracleAppOriginalSourceConfig struct {
	CommonOracleAppSourceConfig

	// List of archive logs to apply on Database after overwrite restore.
	RollForwardLogPathVec []string `json:"rollForwardLogPathVec"`

	// Whether or not this is a complete recovery attempt.
	AttemptCompleteRecovery *bool `json:"attemptCompleteRecovery,omitempty"`

	// UTC time in msecs till which we have to roll-forward the database.
	RollForwardTimeMsecs *int64 `json:"rollForwardTimeMsecs,omitempty"`

	// Specifies whether allowed to automatically stop active passive resource.
	StopActivePassive *bool `json:"stopActivePassive,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RecoverOracleAppOriginalSourceConfig) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonOracleAppSourceConfig
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonOracleAppSourceConfig = aO0

	// now for regular properties
	var propsRecoverOracleAppOriginalSourceConfig struct {
		RollForwardLogPathVec []string `json:"rollForwardLogPathVec"`

		AttemptCompleteRecovery *bool `json:"attemptCompleteRecovery,omitempty"`

		RollForwardTimeMsecs *int64 `json:"rollForwardTimeMsecs,omitempty"`

		StopActivePassive *bool `json:"stopActivePassive,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsRecoverOracleAppOriginalSourceConfig); err != nil {
		return err
	}
	m.RollForwardLogPathVec = propsRecoverOracleAppOriginalSourceConfig.RollForwardLogPathVec

	m.AttemptCompleteRecovery = propsRecoverOracleAppOriginalSourceConfig.AttemptCompleteRecovery

	m.RollForwardTimeMsecs = propsRecoverOracleAppOriginalSourceConfig.RollForwardTimeMsecs

	m.StopActivePassive = propsRecoverOracleAppOriginalSourceConfig.StopActivePassive

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RecoverOracleAppOriginalSourceConfig) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.CommonOracleAppSourceConfig)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsRecoverOracleAppOriginalSourceConfig struct {
		RollForwardLogPathVec []string `json:"rollForwardLogPathVec"`

		AttemptCompleteRecovery *bool `json:"attemptCompleteRecovery,omitempty"`

		RollForwardTimeMsecs *int64 `json:"rollForwardTimeMsecs,omitempty"`

		StopActivePassive *bool `json:"stopActivePassive,omitempty"`
	}
	propsRecoverOracleAppOriginalSourceConfig.RollForwardLogPathVec = m.RollForwardLogPathVec

	propsRecoverOracleAppOriginalSourceConfig.AttemptCompleteRecovery = m.AttemptCompleteRecovery

	propsRecoverOracleAppOriginalSourceConfig.RollForwardTimeMsecs = m.RollForwardTimeMsecs

	propsRecoverOracleAppOriginalSourceConfig.StopActivePassive = m.StopActivePassive

	jsonDataPropsRecoverOracleAppOriginalSourceConfig, errRecoverOracleAppOriginalSourceConfig := swag.WriteJSON(propsRecoverOracleAppOriginalSourceConfig)
	if errRecoverOracleAppOriginalSourceConfig != nil {
		return nil, errRecoverOracleAppOriginalSourceConfig
	}
	_parts = append(_parts, jsonDataPropsRecoverOracleAppOriginalSourceConfig)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this recover oracle app original source config
func (m *RecoverOracleAppOriginalSourceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonOracleAppSourceConfig
	if err := m.CommonOracleAppSourceConfig.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this recover oracle app original source config based on the context it is used
func (m *RecoverOracleAppOriginalSourceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonOracleAppSourceConfig
	if err := m.CommonOracleAppSourceConfig.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *RecoverOracleAppOriginalSourceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverOracleAppOriginalSourceConfig) UnmarshalBinary(b []byte) error {
	var res RecoverOracleAppOriginalSourceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
