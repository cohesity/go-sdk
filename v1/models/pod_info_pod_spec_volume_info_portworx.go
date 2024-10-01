// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodInfoPodSpecVolumeInfoPortworx Portworx represents a Portworx volume resource.
//
// swagger:model PodInfo_PodSpec_VolumeInfo_Portworx
type PodInfoPodSpecVolumeInfoPortworx struct {

	// fs type
	FsType *string `json:"fsType,omitempty"`

	// read only
	ReadOnly *bool `json:"readOnly,omitempty"`

	// volume Id
	VolumeID *string `json:"volumeId,omitempty"`
}

// Validate validates this pod info pod spec volume info portworx
func (m *PodInfoPodSpecVolumeInfoPortworx) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod info pod spec volume info portworx based on context it is used
func (m *PodInfoPodSpecVolumeInfoPortworx) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PodInfoPodSpecVolumeInfoPortworx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodInfoPodSpecVolumeInfoPortworx) UnmarshalBinary(b []byte) error {
	var res PodInfoPodSpecVolumeInfoPortworx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
