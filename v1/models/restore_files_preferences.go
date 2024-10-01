// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestoreFilesPreferences This message captures preferences from the user while restoring the files
// on the target.
//
// swagger:model RestoreFilesPreferences
type RestoreFilesPreferences struct {

	// This must be set to a directory path if restore_to_original_paths is
	// false. All the files and directories restored will be restored under this
	// location.
	AlternateRestoreBaseDirectory *string `json:"alternateRestoreBaseDirectory,omitempty"`

	// Whether to continue with the copy in case of encountering an error.
	ContinueOnError *bool `json:"continueOnError,omitempty"`

	// Whether to enable encryption for NFS and SMB restores.
	EncryptionEnabled *bool `json:"encryptionEnabled,omitempty"`

	// In case of GCP Linux restores, whether to generate ssh keys to connect to
	// the customer's instance.
	GenerateSSHKeys *bool `json:"generateSshKeys,omitempty"`

	// This is relevant only if restore_to_original_paths is true. If this is
	// true, then already existing files will be overridden, otherwise new files
	// will be skipped.
	OverrideOriginals *bool `json:"overrideOriginals,omitempty"`

	// Whether to preserve the ACLs of the original file.
	PreserveAcls *bool `json:"preserveAcls,omitempty"`

	// Whether to preserve the original attributes.
	PreserveAttributes *bool `json:"preserveAttributes,omitempty"`

	// Whether to preserve the original time stamps.
	PreserveTimestamps *bool `json:"preserveTimestamps,omitempty"`

	// Option to select whether to restore everything or ACLs only.
	RestoreEntities *int32 `json:"restoreEntities,omitempty"`

	// If this is true, then files will be restored to original paths.
	RestoreToOriginalPaths *bool `json:"restoreToOriginalPaths,omitempty"`

	// Whether to save success files for FLR.
	SaveSuccessFiles *bool `json:"saveSuccessFiles,omitempty"`

	// Whether to skip the estimation step.
	SkipEstimation *bool `json:"skipEstimation,omitempty"`
}

// Validate validates this restore files preferences
func (m *RestoreFilesPreferences) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this restore files preferences based on context it is used
func (m *RestoreFilesPreferences) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestoreFilesPreferences) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreFilesPreferences) UnmarshalBinary(b []byte) error {
	var res RestoreFilesPreferences
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
