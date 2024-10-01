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

// OneDriveOwner OneDrive Owner.
//
// Specifies OneDrive owner details.
//
// swagger:model OneDriveOwner
type OneDriveOwner struct {

	// Specifies the Drives that a user owns which are to be restored.
	DriveInfoList []*OneDriveInfo `json:"driveInfoList"`

	// Specifies the details about the Office365 user who owns the drive.
	UserDetailObject *RestoreObjectDetails `json:"userDetailObject,omitempty"`
}

// Validate validates this one drive owner
func (m *OneDriveOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDriveInfoList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserDetailObject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OneDriveOwner) validateDriveInfoList(formats strfmt.Registry) error {
	if swag.IsZero(m.DriveInfoList) { // not required
		return nil
	}

	for i := 0; i < len(m.DriveInfoList); i++ {
		if swag.IsZero(m.DriveInfoList[i]) { // not required
			continue
		}

		if m.DriveInfoList[i] != nil {
			if err := m.DriveInfoList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("driveInfoList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("driveInfoList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OneDriveOwner) validateUserDetailObject(formats strfmt.Registry) error {
	if swag.IsZero(m.UserDetailObject) { // not required
		return nil
	}

	if m.UserDetailObject != nil {
		if err := m.UserDetailObject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userDetailObject")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userDetailObject")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this one drive owner based on the context it is used
func (m *OneDriveOwner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDriveInfoList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserDetailObject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OneDriveOwner) contextValidateDriveInfoList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DriveInfoList); i++ {

		if m.DriveInfoList[i] != nil {

			if swag.IsZero(m.DriveInfoList[i]) { // not required
				return nil
			}

			if err := m.DriveInfoList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("driveInfoList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("driveInfoList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OneDriveOwner) contextValidateUserDetailObject(ctx context.Context, formats strfmt.Registry) error {

	if m.UserDetailObject != nil {

		if swag.IsZero(m.UserDetailObject) { // not required
			return nil
		}

		if err := m.UserDetailObject.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userDetailObject")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userDetailObject")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OneDriveOwner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OneDriveOwner) UnmarshalBinary(b []byte) error {
	var res OneDriveOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
