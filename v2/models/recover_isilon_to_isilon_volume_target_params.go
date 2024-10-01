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

// RecoverIsilonToIsilonVolumeTargetParams Recover Isilon To Isilon Volume Target Params.
//
// Specifies the params of the Isilon recovery target.
//
// swagger:model RecoverIsilonToIsilonVolumeTargetParams
type RecoverIsilonToIsilonVolumeTargetParams struct {

	// Specifies the parameter whether the recovery should be performed to a new or the original Isilon target.
	// Required: true
	RecoverToNewSource *bool `json:"recoverToNewSource"`

	// Specifies the new destination Source configuration parameters where the volumes will be recovered. This is mandatory if recoverToNewSource is set to true.
	NewSourceConfig *RecoverOtherNasToIsilonVolumeTargetParams `json:"newSourceConfig,omitempty"`

	// Specifies the Source configuration if volumes are being recovered to original Source. If not specified, all the configuration parameters will be retained.
	OriginalSourceConfig *OriginalIsilonTargetParams `json:"originalSourceConfig,omitempty"`
}

// Validate validates this recover isilon to isilon volume target params
func (m *RecoverIsilonToIsilonVolumeTargetParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecoverToNewSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewSourceConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginalSourceConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverIsilonToIsilonVolumeTargetParams) validateRecoverToNewSource(formats strfmt.Registry) error {

	if err := validate.Required("recoverToNewSource", "body", m.RecoverToNewSource); err != nil {
		return err
	}

	return nil
}

func (m *RecoverIsilonToIsilonVolumeTargetParams) validateNewSourceConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.NewSourceConfig) { // not required
		return nil
	}

	if m.NewSourceConfig != nil {
		if err := m.NewSourceConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("newSourceConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("newSourceConfig")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverIsilonToIsilonVolumeTargetParams) validateOriginalSourceConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.OriginalSourceConfig) { // not required
		return nil
	}

	if m.OriginalSourceConfig != nil {
		if err := m.OriginalSourceConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originalSourceConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("originalSourceConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover isilon to isilon volume target params based on the context it is used
func (m *RecoverIsilonToIsilonVolumeTargetParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNewSourceConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginalSourceConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverIsilonToIsilonVolumeTargetParams) contextValidateNewSourceConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.NewSourceConfig != nil {

		if swag.IsZero(m.NewSourceConfig) { // not required
			return nil
		}

		if err := m.NewSourceConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("newSourceConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("newSourceConfig")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverIsilonToIsilonVolumeTargetParams) contextValidateOriginalSourceConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.OriginalSourceConfig != nil {

		if swag.IsZero(m.OriginalSourceConfig) { // not required
			return nil
		}

		if err := m.OriginalSourceConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originalSourceConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("originalSourceConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverIsilonToIsilonVolumeTargetParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverIsilonToIsilonVolumeTargetParams) UnmarshalBinary(b []byte) error {
	var res RecoverIsilonToIsilonVolumeTargetParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
