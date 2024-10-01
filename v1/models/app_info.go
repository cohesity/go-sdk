// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppInfo App Information.
//
// Specifies the App information.
//
// swagger:model AppInfo
type AppInfo struct {

	// Specifies the unique instance Id of the App.
	AppInstanceID *int64 `json:"appInstanceId,omitempty"`

	// Specifies the name of the App.
	AppName *string `json:"appName,omitempty"`

	// Specifies the unique id of the App.
	AppUID *int64 `json:"appUid,omitempty"`

	// Specifies the version of the App.
	AppVersion *int64 `json:"appVersion,omitempty"`
}

// Validate validates this app info
func (m *AppInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this app info based on context it is used
func (m *AppInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppInfo) UnmarshalBinary(b []byte) error {
	var res AppInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
