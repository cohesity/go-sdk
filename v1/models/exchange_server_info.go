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

// ExchangeServerInfo Information about an Exchange Server. Obtained from fields returned by
// Get-MailboxServer -Status
//
// swagger:model ExchangeServerInfo
type ExchangeServerInfo struct {

	// Whether backing up this Server is supported or not due to system or
	// implementation limitations. If the value is false, reason string below can
	// be used to say why its not supported.
	BackupSupported *bool `json:"backupSupported,omitempty"`

	// Free form text for the reasons why backup is not supported for this
	// Server. This is valid only if backup_supported = false.
	BackupUnsupportedReasonVec []string `json:"backupUnsupportedReasonVec"`

	// Default dir path for creating future mailbox databases. This path does not
	// correspond to current database directory.
	DefaultDataDirPath *string `json:"defaultDataDirPath,omitempty"`

	// Identity of the server.
	ID *ExchangeServerIdentity `json:"id,omitempty"`

	// Exchange organization distinguished name. This may be used in database
	// restore cases to match backed up org vs restoring org. Exchange allows
	// only recovery DB to be created on foreign organizations. Example: CN=First
	// Organization,CN=Microsoft
	// Exchange,CN=Services,CN=Configuration,DC=Exch16-OnPrem,DC=com
	OrganizationDn *string `json:"organizationDn,omitempty"`

	// Role of the exchange server.
	Role *int32 `json:"role,omitempty"`

	// Offset from UTC in minutes for this Exchange server machine. This number
	// can be positive or negative. For example -480 is Pacific Time zone.
	UtcOffsetmin *int32 `json:"utcOffsetmin,omitempty"`

	// Server version.
	Version *ExchangeServerVersion `json:"version,omitempty"`

	// VSS writers that should be included during backup. For Exchange 2013
	// onwards, it will be a single 'Microsoft Exchange Writer'. For Exchange
	// 2010 DAG, it will also include 'Microsoft Exchange Replica Writer'.
	VssBackupWritersVec []string `json:"vssBackupWritersVec"`

	// VSS writer that should be included during restore. Its value will be
	// 'Microsoft Exchange Writer'. Replica writer should not be included during
	// VSS restore.
	VssRestoreWriter *string `json:"vssRestoreWriter,omitempty"`
}

// Validate validates this exchange server info
func (m *ExchangeServerInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExchangeServerInfo) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if m.ID != nil {
		if err := m.ID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("id")
			}
			return err
		}
	}

	return nil
}

func (m *ExchangeServerInfo) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if m.Version != nil {
		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this exchange server info based on the context it is used
func (m *ExchangeServerInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExchangeServerInfo) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if m.ID != nil {

		if swag.IsZero(m.ID) { // not required
			return nil
		}

		if err := m.ID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("id")
			}
			return err
		}
	}

	return nil
}

func (m *ExchangeServerInfo) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.Version != nil {

		if swag.IsZero(m.Version) { // not required
			return nil
		}

		if err := m.Version.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExchangeServerInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExchangeServerInfo) UnmarshalBinary(b []byte) error {
	var res ExchangeServerInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
