// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExchangeServerIdentity Exchange Server Identity. Fields returned by PowerShell: Get-MailboxServer
// Status. All fields are case insensitive.
//
// swagger:model ExchangeServerIdentity
type ExchangeServerIdentity struct {

	// Fully Qualified Domain Name (FQDN) of the Exchange sever. Obtained from
	// (Get-ExchangeServer).NetWorkAddress protocol type:'ncacn_ip_tcp'.
	Fqdn *string `json:"fqdn,omitempty"`

	// unique guid for the server in AD. This is a unqiue identifier that can be
	// used to compare servers.
	GUID *string `json:"guid,omitempty"`

	// Display Name. Example: EXCH12-MAILB01. This is also the identity, which is
	// unique within an AD forest.
	Name *string `json:"name,omitempty"`
}

// Validate validates this exchange server identity
func (m *ExchangeServerIdentity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this exchange server identity based on context it is used
func (m *ExchangeServerIdentity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExchangeServerIdentity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExchangeServerIdentity) UnmarshalBinary(b []byte) error {
	var res ExchangeServerIdentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
