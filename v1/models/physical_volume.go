// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PhysicalVolume Physical Volume.
//
// Specifies volume information about a Physical Protection Source.
//
// swagger:model PhysicalVolume
type PhysicalVolume struct {

	// Specifies the path to the device that hosts the volume locally.
	DevicePath *string `json:"devicePath,omitempty"`

	// Specifies an id for the Physical Volume.
	GUID *string `json:"guid,omitempty"`

	// Specifies whether the volume is boot volume.
	IsBootVolume *bool `json:"isBootVolume,omitempty"`

	// Specifies whether this volume supports extended attributes (like ACLs)
	// when performing file backups.
	IsExtendedAttributesSupported *bool `json:"isExtendedAttributesSupported,omitempty"`

	// Specifies if a volume is protected by a Job.
	IsProtected *bool `json:"isProtected,omitempty"`

	// Specifies whether the volume is shared volume.
	IsSharedVolume *bool `json:"isSharedVolume,omitempty"`

	// Specifies a volume label that can be used for displaying additional
	// identifying information about a volume.
	Label *string `json:"label,omitempty"`

	// Specifies the logical size of the volume in bytes that is
	// not reduced by change-block tracking, compression and deduplication.
	LogicalSizeBytes *uint64 `json:"logicalSizeBytes,omitempty"`

	// Array of Mount Points.
	//
	// Specifies the mount points where the volume is mounted,
	// for example: 'C:\', '/mnt/foo' etc.
	MountPoints []string `json:"mountPoints"`

	// Specifies mount type of volume e.g. nfs, autofs, ext4 etc.
	MountType *string `json:"mountType,omitempty"`

	// Specifies the full path to connect to the network attached volume.
	// For example, (IP or hostname):/path/to/share for NFS volumes).
	NetworkPath *string `json:"networkPath,omitempty"`

	// Specifies the size used by the volume in bytes.
	UsedSizeBytes *uint64 `json:"usedSizeBytes,omitempty"`
}

// Validate validates this physical volume
func (m *PhysicalVolume) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this physical volume based on context it is used
func (m *PhysicalVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PhysicalVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PhysicalVolume) UnmarshalBinary(b []byte) error {
	var res PhysicalVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
