// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExchangeServerVersion Exchange Server Version.
//
// swagger:model ExchangeServerVersion
type ExchangeServerVersion struct {

	// build ver
	BuildVer *int32 `json:"buildVer,omitempty"`

	// Display version string. From Get-exchangeserver | fl AdminDisplayVersion
	// Example: Version 15.0 (Build 1365.1)
	// Mapping located at https://goo.gl/9JQNfT
	// Microsot does not provide any mechanism to get nice display name like
	// 'Exchange Server 2013 CU19'
	DisplayVersion *string `json:"displayVersion,omitempty"`

	// Parsed file version of store service number from
	// Get-Command Microsoft.Exchange.Store.Service.exe |
	// ForEach-Object {$_.FileVersionInfo.ProductVersion} = 15.00.1365.000
	// If this is set to -1, that means the version number is unknown.
	// It follows Microsoft convention https://goo.gl/536Ztg
	MajorVer *int32 `json:"majorVer,omitempty"`

	// minor ver
	MinorVer *int32 `json:"minorVer,omitempty"`

	// revision num
	RevisionNum *int32 `json:"revisionNum,omitempty"`
}

// Validate validates this exchange server version
func (m *ExchangeServerVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this exchange server version based on context it is used
func (m *ExchangeServerVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExchangeServerVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExchangeServerVersion) UnmarshalBinary(b []byte) error {
	var res ExchangeServerVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
