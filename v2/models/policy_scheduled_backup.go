// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PolicyScheduledBackup Policy Scheduled Backup type.
//
// Policy Scheduled Backup type.
//
// swagger:model PolicyScheduledBackup
type PolicyScheduledBackup struct {

	// Specifies Scheduled Backup type.
	// Enum: ["Regular","Full","Log","System"]
	PolicyScheduledBackup string `json:"PolicyScheduledBackup,omitempty"`
}

// Validate validates this policy scheduled backup
func (m *PolicyScheduledBackup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicyScheduledBackup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var policyScheduledBackupTypePolicyScheduledBackupPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Regular","Full","Log","System"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		policyScheduledBackupTypePolicyScheduledBackupPropEnum = append(policyScheduledBackupTypePolicyScheduledBackupPropEnum, v)
	}
}

const (

	// PolicyScheduledBackupPolicyScheduledBackupRegular captures enum value "Regular"
	PolicyScheduledBackupPolicyScheduledBackupRegular string = "Regular"

	// PolicyScheduledBackupPolicyScheduledBackupFull captures enum value "Full"
	PolicyScheduledBackupPolicyScheduledBackupFull string = "Full"

	// PolicyScheduledBackupPolicyScheduledBackupLog captures enum value "Log"
	PolicyScheduledBackupPolicyScheduledBackupLog string = "Log"

	// PolicyScheduledBackupPolicyScheduledBackupSystem captures enum value "System"
	PolicyScheduledBackupPolicyScheduledBackupSystem string = "System"
)

// prop value enum
func (m *PolicyScheduledBackup) validatePolicyScheduledBackupEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, policyScheduledBackupTypePolicyScheduledBackupPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PolicyScheduledBackup) validatePolicyScheduledBackup(formats strfmt.Registry) error {
	if swag.IsZero(m.PolicyScheduledBackup) { // not required
		return nil
	}

	// value enum
	if err := m.validatePolicyScheduledBackupEnum("PolicyScheduledBackup", "body", m.PolicyScheduledBackup); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this policy scheduled backup based on context it is used
func (m *PolicyScheduledBackup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyScheduledBackup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyScheduledBackup) UnmarshalBinary(b []byte) error {
	var res PolicyScheduledBackup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
