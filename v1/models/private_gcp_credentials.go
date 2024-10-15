// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PrivateGcpCredentials The service account credentials allow us to access the Google account.
//
// swagger:model PrivateGcpCredentials
type PrivateGcpCredentials struct {

	// The url of the public x509 certificate used to verify JSON web tokens
	// signed by the authentication provider.
	AuthProviderX509CertURL *string `json:"authProviderX509CertUrl,omitempty"`

	// The authorization server endpoint URI.
	AuthURI *string `json:"authUri,omitempty"`

	// Client email address associated with the service account.
	ClientEmailAddress *string `json:"clientEmailAddress,omitempty"`

	// Client Id associated with the Google cloud account.
	ClientID *string `json:"clientId,omitempty"`

	// The url of the public x509 certificate used to verify JSON web tokens
	// signed by the client.
	ClientX509CertURL *string `json:"clientX509CertUrl,omitempty"`

	// Client private key generated at the time of creating the service
	// account. This is used for authenticating ourselves with Google cloud.
	// This field only stores the private key in bytes and doesn't encrypt it.
	EncryptedClientPrivateKey []uint8 `json:"encryptedClientPrivateKey"`

	// Encrypted client private key. Unencrypted version is
	// received in "encrypted_client_private_key" from iris.
	EncryptedPrivateKeyID *string `json:"encryptedPrivateKeyId,omitempty"`

	// Private key Id associated with the Google cloud account.
	PrivateKeyID *string `json:"privateKeyId,omitempty"`

	// Id of the project associated with Google cloud account.
	ProjectID *string `json:"projectId,omitempty"`

	// The token server endpoint URI.
	TokenURI *string `json:"tokenUri,omitempty"`

	// The type of our application.
	Type *string `json:"type,omitempty"`
}

// Validate validates this private gcp credentials
func (m *PrivateGcpCredentials) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this private gcp credentials based on context it is used
func (m *PrivateGcpCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrivateGcpCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateGcpCredentials) UnmarshalBinary(b []byte) error {
	var res PrivateGcpCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
