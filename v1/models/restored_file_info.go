// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestoredFileInfo This message has information on file/directory being restored.
//
// swagger:model RestoredFileInfo
type RestoredFileInfo struct {

	// Full path of the file being restored: the actual file path without the
	// disk. E.g.:
	// \Program Files\App\file.txt
	AbsolutePath *string `json:"absolutePath,omitempty"`

	// Disk information of where the source file is currently located.
	AttachedDiskID *int32 `json:"attachedDiskId,omitempty"`

	// Disk partition to which the file belongs to.
	DiskPartitionID *int32 `json:"diskPartitionId,omitempty"`

	// File system UUID on which file resides.
	FsUUID *string `json:"fsUuid,omitempty"`

	// Inode number of the file.
	// This is needed for snapmirror restore workflow.
	InodeNumber *int64 `json:"inodeNumber,omitempty"`

	// Whether the path points to a directory.
	IsDirectory *bool `json:"isDirectory,omitempty"`

	// This will be set to true for recovery workflows for non-simple volumes
	// on Windows Dynamic Disks. In that case, we will use VolumeInfo instead of
	// some of the details captured here (e.g. virtual_disk_file) for determining
	// disk and volume related details.
	IsNonSimpleLdmVol *bool `json:"isNonSimpleLdmVol,omitempty"`

	// This must be set to a directory path if restore_to_original_paths is
	// false and restore task has multiple files which are not desired to be
	// restore to one common location. If this filed is populated,
	// 'absolute_path' will be restored under this location. If this field is not
	// populated all files in restore task will be restored to location specified
	// in RestoreFilesPreferences.
	RestoreBaseDirectory *string `json:"restoreBaseDirectory,omitempty"`

	// Mount point of the volume on which the file to be restored is located.
	// E.g.:
	// c:\temp\vhd_mount_1234
	RestoreMountPoint *string `json:"restoreMountPoint,omitempty"`

	// Size of the file in bytes. Required in FLR in GCP using Cloud Functions.
	SizeBytes *int64 `json:"sizeBytes,omitempty"`

	// Virtual disk file to which this file belongs to.
	VirtualDiskFile *string `json:"virtualDiskFile,omitempty"`

	// Id of the volume.
	VolumeID *string `json:"volumeId,omitempty"`

	// Original volume name (or drive letter). This is used while performing the
	// copy to the original paths.
	// E.g.:
	// c:
	VolumePath *string `json:"volumePath,omitempty"`
}

// Validate validates this restored file info
func (m *RestoredFileInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this restored file info based on context it is used
func (m *RestoredFileInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestoredFileInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoredFileInfo) UnmarshalBinary(b []byte) error {
	var res RestoredFileInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
