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

// PhysicalFileBackupParams Physical File Backup Parameters.
//
// Message to capture params when backing up files on a Physical source.
//
// swagger:model PhysicalFileBackupParams
type PhysicalFileBackupParams struct {

	// Specifies the paths to backup on the Physical source.
	BackupPathInfoVec []*PhysicalFileBackupParamsBackupPathInfo `json:"backupPathInfoVec"`

	// Specifies metadata path on source. This file contains absolute paths of
	// files that needs to be backed up on the same source. If this field is
	// set, backup_path_info_vec will be ignored.
	MetadataFilePath *string `json:"metadataFilePath,omitempty"`

	// Mount types of nested volumes to be skipped.
	SkipNestedVolumesVec []string `json:"skipNestedVolumesVec"`

	// Specifies whether to follow nas target pointed by symlink.
	// Set to true only for windows physical file based job.
	SymlinkFollowNasTarget *bool `json:"symlinkFollowNasTarget,omitempty"`

	// Specifies whether to use skip_nested_volumes_vec to skip nested mounts.
	// Before 6.4, BackupPathInfo.skip_nested_volumes boolean was used to skip
	// nested volumes. So we use this boolean to support older jobs.
	UsesSkipNestedVolumesVec *bool `json:"usesSkipNestedVolumesVec,omitempty"`
}

// Validate validates this physical file backup params
func (m *PhysicalFileBackupParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupPathInfoVec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PhysicalFileBackupParams) validateBackupPathInfoVec(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupPathInfoVec) { // not required
		return nil
	}

	for i := 0; i < len(m.BackupPathInfoVec); i++ {
		if swag.IsZero(m.BackupPathInfoVec[i]) { // not required
			continue
		}

		if m.BackupPathInfoVec[i] != nil {
			if err := m.BackupPathInfoVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backupPathInfoVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backupPathInfoVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this physical file backup params based on the context it is used
func (m *PhysicalFileBackupParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupPathInfoVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PhysicalFileBackupParams) contextValidateBackupPathInfoVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BackupPathInfoVec); i++ {

		if m.BackupPathInfoVec[i] != nil {

			if swag.IsZero(m.BackupPathInfoVec[i]) { // not required
				return nil
			}

			if err := m.BackupPathInfoVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backupPathInfoVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backupPathInfoVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PhysicalFileBackupParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PhysicalFileBackupParams) UnmarshalBinary(b []byte) error {
	var res PhysicalFileBackupParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
