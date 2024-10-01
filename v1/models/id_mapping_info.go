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

// IDMappingInfo Update ID Mapping Information Request.
//
// Specifies the params required to update the user id mapping info of an
// Active Directory.
//
// swagger:model IdMappingInfo
type IDMappingInfo struct {

	// Specifies the fallback id mapping info which is used when an ID mapping
	// for a user is not found via the above IdMappingInfo. Only supported for
	// two types of fallback mapping types - 'kRid' and 'kFixed'.
	FallbackUserIDMappingInfo *UserIDMapping `json:"fallbackUserIdMappingInfo,omitempty"`

	// Specifies the SID of the Active Directory domain user to be mapped to
	// Unix root user.
	UnixRootSid *string `json:"unixRootSid,omitempty"`

	// Specifies the information about how the Unix and Windows users are mapped
	// for this domain.
	UserIDMappingInfo *UserIDMapping `json:"userIdMappingInfo,omitempty"`
}

// Validate validates this Id mapping info
func (m *IDMappingInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFallbackUserIDMappingInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserIDMappingInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IDMappingInfo) validateFallbackUserIDMappingInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.FallbackUserIDMappingInfo) { // not required
		return nil
	}

	if m.FallbackUserIDMappingInfo != nil {
		if err := m.FallbackUserIDMappingInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fallbackUserIdMappingInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fallbackUserIdMappingInfo")
			}
			return err
		}
	}

	return nil
}

func (m *IDMappingInfo) validateUserIDMappingInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.UserIDMappingInfo) { // not required
		return nil
	}

	if m.UserIDMappingInfo != nil {
		if err := m.UserIDMappingInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userIdMappingInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userIdMappingInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Id mapping info based on the context it is used
func (m *IDMappingInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFallbackUserIDMappingInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserIDMappingInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IDMappingInfo) contextValidateFallbackUserIDMappingInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.FallbackUserIDMappingInfo != nil {

		if swag.IsZero(m.FallbackUserIDMappingInfo) { // not required
			return nil
		}

		if err := m.FallbackUserIDMappingInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fallbackUserIdMappingInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fallbackUserIdMappingInfo")
			}
			return err
		}
	}

	return nil
}

func (m *IDMappingInfo) contextValidateUserIDMappingInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.UserIDMappingInfo != nil {

		if swag.IsZero(m.UserIDMappingInfo) { // not required
			return nil
		}

		if err := m.UserIDMappingInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userIdMappingInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userIdMappingInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IDMappingInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IDMappingInfo) UnmarshalBinary(b []byte) error {
	var res IDMappingInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
