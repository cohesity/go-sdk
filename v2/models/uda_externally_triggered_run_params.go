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

// UdaExternallyTriggeredRunParams Specifies the parameters for an externally triggered run.
//
// swagger:model UdaExternallyTriggeredRunParams
type UdaExternallyTriggeredRunParams struct {

	// Specifies the IP or FQDN of the source host where this backup will run.
	ControlNode *string `json:"controlNode,omitempty"`

	// Specifies a map of custom arguments to be supplied to the plugin.
	BackupArgs []*KeyValuePair `json:"backupArgs"`
}

// Validate validates this uda externally triggered run params
func (m *UdaExternallyTriggeredRunParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupArgs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UdaExternallyTriggeredRunParams) validateBackupArgs(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupArgs) { // not required
		return nil
	}

	for i := 0; i < len(m.BackupArgs); i++ {
		if swag.IsZero(m.BackupArgs[i]) { // not required
			continue
		}

		if m.BackupArgs[i] != nil {
			if err := m.BackupArgs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backupArgs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backupArgs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this uda externally triggered run params based on the context it is used
func (m *UdaExternallyTriggeredRunParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupArgs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UdaExternallyTriggeredRunParams) contextValidateBackupArgs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BackupArgs); i++ {

		if m.BackupArgs[i] != nil {

			if swag.IsZero(m.BackupArgs[i]) { // not required
				return nil
			}

			if err := m.BackupArgs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backupArgs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backupArgs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UdaExternallyTriggeredRunParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UdaExternallyTriggeredRunParams) UnmarshalBinary(b []byte) error {
	var res UdaExternallyTriggeredRunParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
