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

// KubernetesBackupSourceParams Kubernetes Backup Source Parameters.
//
// Message to capture additional backup params for a Kubernetes type source.
//
// swagger:model KubernetesBackupSourceParams
type KubernetesBackupSourceParams struct {

	// Info about PVC(s) to be excluded from the from the backup job for
	// the source. When set, this will override KubernetesEnvParams label based
	// PVC filter.
	ExcludeParams *K8SFilterParams `json:"excludeParams,omitempty"`

	// Info about PVC(s) to be included from the from the backup job for
	// the source. When set, this will override KubernetesEnvParams label based
	// PVC filter.
	IncludeParams *K8SFilterParams `json:"includeParams,omitempty"`
}

// Validate validates this kubernetes backup source params
func (m *KubernetesBackupSourceParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExcludeParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncludeParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesBackupSourceParams) validateExcludeParams(formats strfmt.Registry) error {
	if swag.IsZero(m.ExcludeParams) { // not required
		return nil
	}

	if m.ExcludeParams != nil {
		if err := m.ExcludeParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("excludeParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("excludeParams")
			}
			return err
		}
	}

	return nil
}

func (m *KubernetesBackupSourceParams) validateIncludeParams(formats strfmt.Registry) error {
	if swag.IsZero(m.IncludeParams) { // not required
		return nil
	}

	if m.IncludeParams != nil {
		if err := m.IncludeParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("includeParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("includeParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kubernetes backup source params based on the context it is used
func (m *KubernetesBackupSourceParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExcludeParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncludeParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesBackupSourceParams) contextValidateExcludeParams(ctx context.Context, formats strfmt.Registry) error {

	if m.ExcludeParams != nil {

		if swag.IsZero(m.ExcludeParams) { // not required
			return nil
		}

		if err := m.ExcludeParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("excludeParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("excludeParams")
			}
			return err
		}
	}

	return nil
}

func (m *KubernetesBackupSourceParams) contextValidateIncludeParams(ctx context.Context, formats strfmt.Registry) error {

	if m.IncludeParams != nil {

		if swag.IsZero(m.IncludeParams) { // not required
			return nil
		}

		if err := m.IncludeParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("includeParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("includeParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesBackupSourceParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesBackupSourceParams) UnmarshalBinary(b []byte) error {
	var res KubernetesBackupSourceParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
