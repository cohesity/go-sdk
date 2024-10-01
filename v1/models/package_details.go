// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PackageDetails Package Details
//
// Specifies all of the details of a package that is currently installed
// on the cluster.
//
// swagger:model PackageDetails
type PackageDetails struct {

	// Specifies whether or not downtime is required to update to this package.
	DowntimeRequired *bool `json:"downtimeRequired,omitempty"`

	// Specifies the list of IDs of nodes on the cluster where this package is
	// currently installed.
	InstalledOnNodes []int64 `json:"installedOnNodes"`

	// Specifies the name of the current package.
	PackageName *string `json:"packageName,omitempty"`

	// Specifies the release date of this package.
	ReleaseDate *string `json:"releaseDate,omitempty"`
}

// Validate validates this package details
func (m *PackageDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this package details based on context it is used
func (m *PackageDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PackageDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackageDetails) UnmarshalBinary(b []byte) error {
	var res PackageDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
