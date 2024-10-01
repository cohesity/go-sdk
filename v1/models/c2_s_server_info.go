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

// C2SServerInfo Data definition related to C2S Server Info while register AWS source.
//
// C2S Server Info.
//
// Specifies information required to connect to CAP to get AWS credentials.
// C2SAccessPortal(CAP) is AWS commercial cloud service access portal.
//
// swagger:model C2SServerInfo
type C2SServerInfo struct {

	// Specifies the C2S Access Portal (CAP) which is used to get the aws
	// credentials in Amazon Commercial Cloud Service(C2S).
	C2sAccessPortal *C2SAccessPortal `json:"c2sAccessPortal,omitempty"`

	// Specifies the CA (certificate authority) trusted certificate.
	CaTrustedCertificate *string `json:"caTrustedCertificate,omitempty"`

	// Specifies the client CA  certificate. This certificate is in pem format.
	ClientCertificate *string `json:"clientCertificate,omitempty"`

	// Specifies the client private key. This certificate is in pem format.
	ClientPrivateKey *string `json:"clientPrivateKey,omitempty"`
}

// Validate validates this c2 s server info
func (m *C2SServerInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateC2sAccessPortal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *C2SServerInfo) validateC2sAccessPortal(formats strfmt.Registry) error {
	if swag.IsZero(m.C2sAccessPortal) { // not required
		return nil
	}

	if m.C2sAccessPortal != nil {
		if err := m.C2sAccessPortal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("c2sAccessPortal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("c2sAccessPortal")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this c2 s server info based on the context it is used
func (m *C2SServerInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateC2sAccessPortal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *C2SServerInfo) contextValidateC2sAccessPortal(ctx context.Context, formats strfmt.Registry) error {

	if m.C2sAccessPortal != nil {

		if swag.IsZero(m.C2sAccessPortal) { // not required
			return nil
		}

		if err := m.C2sAccessPortal.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("c2sAccessPortal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("c2sAccessPortal")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *C2SServerInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *C2SServerInfo) UnmarshalBinary(b []byte) error {
	var res C2SServerInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
