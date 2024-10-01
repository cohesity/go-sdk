// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// O365OneDriveRestoreEntityParamsDriveItem DriveItems are used in the case of partial drive recovery.
//
// swagger:model O365OneDriveRestoreEntityParams_DriveItem
type O365OneDriveRestoreEntityParamsDriveItem struct {

	// The unique identifier of the item within the Drive.
	DriveItemID *string `json:"driveItemId,omitempty"`

	// The path of the drive item relative to root.
	DriveItemPath *string `json:"driveItemPath,omitempty"`

	// Specify if the item is a file or not.
	IsFileItem *bool `json:"isFileItem,omitempty"`
}

// Validate validates this o365 one drive restore entity params drive item
func (m *O365OneDriveRestoreEntityParamsDriveItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this o365 one drive restore entity params drive item based on context it is used
func (m *O365OneDriveRestoreEntityParamsDriveItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *O365OneDriveRestoreEntityParamsDriveItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *O365OneDriveRestoreEntityParamsDriveItem) UnmarshalBinary(b []byte) error {
	var res O365OneDriveRestoreEntityParamsDriveItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
