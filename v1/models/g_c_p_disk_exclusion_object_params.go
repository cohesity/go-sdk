// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GCPDiskExclusionObjectParams Message defining the different criteria to exclude GCP disks from
// object-level backup. All the disk name present here will be excluded.
//
// swagger:model GCPDiskExclusionObjectParams
type GCPDiskExclusionObjectParams struct {

	// List of disk name to exclude. Eg - [instance1-disk, instance2-disk]
	DiskNameVec []string `json:"diskNameVec"`
}

// Validate validates this g c p disk exclusion object params
func (m *GCPDiskExclusionObjectParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this g c p disk exclusion object params based on context it is used
func (m *GCPDiskExclusionObjectParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GCPDiskExclusionObjectParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GCPDiskExclusionObjectParams) UnmarshalBinary(b []byte) error {
	var res GCPDiskExclusionObjectParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
