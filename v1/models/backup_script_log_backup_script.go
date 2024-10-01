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

// BackupScriptLogBackupScript Log Script.
//
// Specifies the script that should run for the Log backup schedule
// of a Remote Adapter 'kPuppeteer' Job.
// This field is mandatory if the Policy associated with this Job
// has a Log backup schedule and this is Remote Adapter
// 'kPuppeteer' Job.
//
// swagger:model backupScriptLogBackupScript
type BackupScriptLogBackupScript struct {
	RemoteScriptPathAndParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *BackupScriptLogBackupScript) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 RemoteScriptPathAndParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.RemoteScriptPathAndParams = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m BackupScriptLogBackupScript) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.RemoteScriptPathAndParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this backup script log backup script
func (m *BackupScriptLogBackupScript) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with RemoteScriptPathAndParams
	if err := m.RemoteScriptPathAndParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this backup script log backup script based on the context it is used
func (m *BackupScriptLogBackupScript) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with RemoteScriptPathAndParams
	if err := m.RemoteScriptPathAndParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *BackupScriptLogBackupScript) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupScriptLogBackupScript) UnmarshalBinary(b []byte) error {
	var res BackupScriptLogBackupScript
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
