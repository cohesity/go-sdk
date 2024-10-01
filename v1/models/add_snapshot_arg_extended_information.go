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

// AddSnapshotArgExtendedInformation This represents a union for capturing environment specific extended
// information.
// NOTE: This presently contains information for VMware but it can be
// extended to contain information regarding any new environment (such as
// HyperV) as needed. Only one of the fields should be set for a snapshot
// at a time and type determines the field that will be used.
//
// swagger:model AddSnapshotArg_ExtendedInformation
type AddSnapshotArgExtendedInformation struct {

	// O365 specific information.
	O365AddSnapshotArg *O365AddSnapshotArg `json:"o365AddSnapshotArg,omitempty"`

	// record stats
	RecordStats *AddSnapshotArgExtendedInformationRecordStats `json:"recordStats,omitempty"`

	// The type of environment this extended info pertains to.
	Type *int32 `json:"type,omitempty"`

	// UDA specific information.
	UdaAddSnapshotArg *AddSnapshotArgExtendedInformationUdaAddSnapshotArg `json:"udaAddSnapshotArg,omitempty"`

	// VMware specific information.
	VmwareAddSnapshotArg *VMwareAddSnapshotArg `json:"vmwareAddSnapshotArg,omitempty"`
}

// Validate validates this add snapshot arg extended information
func (m *AddSnapshotArgExtendedInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateO365AddSnapshotArg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecordStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUdaAddSnapshotArg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVmwareAddSnapshotArg(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddSnapshotArgExtendedInformation) validateO365AddSnapshotArg(formats strfmt.Registry) error {
	if swag.IsZero(m.O365AddSnapshotArg) { // not required
		return nil
	}

	if m.O365AddSnapshotArg != nil {
		if err := m.O365AddSnapshotArg.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("o365AddSnapshotArg")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("o365AddSnapshotArg")
			}
			return err
		}
	}

	return nil
}

func (m *AddSnapshotArgExtendedInformation) validateRecordStats(formats strfmt.Registry) error {
	if swag.IsZero(m.RecordStats) { // not required
		return nil
	}

	if m.RecordStats != nil {
		if err := m.RecordStats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recordStats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recordStats")
			}
			return err
		}
	}

	return nil
}

func (m *AddSnapshotArgExtendedInformation) validateUdaAddSnapshotArg(formats strfmt.Registry) error {
	if swag.IsZero(m.UdaAddSnapshotArg) { // not required
		return nil
	}

	if m.UdaAddSnapshotArg != nil {
		if err := m.UdaAddSnapshotArg.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("udaAddSnapshotArg")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("udaAddSnapshotArg")
			}
			return err
		}
	}

	return nil
}

func (m *AddSnapshotArgExtendedInformation) validateVmwareAddSnapshotArg(formats strfmt.Registry) error {
	if swag.IsZero(m.VmwareAddSnapshotArg) { // not required
		return nil
	}

	if m.VmwareAddSnapshotArg != nil {
		if err := m.VmwareAddSnapshotArg.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmwareAddSnapshotArg")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmwareAddSnapshotArg")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add snapshot arg extended information based on the context it is used
func (m *AddSnapshotArgExtendedInformation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateO365AddSnapshotArg(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecordStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUdaAddSnapshotArg(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVmwareAddSnapshotArg(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddSnapshotArgExtendedInformation) contextValidateO365AddSnapshotArg(ctx context.Context, formats strfmt.Registry) error {

	if m.O365AddSnapshotArg != nil {

		if swag.IsZero(m.O365AddSnapshotArg) { // not required
			return nil
		}

		if err := m.O365AddSnapshotArg.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("o365AddSnapshotArg")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("o365AddSnapshotArg")
			}
			return err
		}
	}

	return nil
}

func (m *AddSnapshotArgExtendedInformation) contextValidateRecordStats(ctx context.Context, formats strfmt.Registry) error {

	if m.RecordStats != nil {

		if swag.IsZero(m.RecordStats) { // not required
			return nil
		}

		if err := m.RecordStats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recordStats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recordStats")
			}
			return err
		}
	}

	return nil
}

func (m *AddSnapshotArgExtendedInformation) contextValidateUdaAddSnapshotArg(ctx context.Context, formats strfmt.Registry) error {

	if m.UdaAddSnapshotArg != nil {

		if swag.IsZero(m.UdaAddSnapshotArg) { // not required
			return nil
		}

		if err := m.UdaAddSnapshotArg.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("udaAddSnapshotArg")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("udaAddSnapshotArg")
			}
			return err
		}
	}

	return nil
}

func (m *AddSnapshotArgExtendedInformation) contextValidateVmwareAddSnapshotArg(ctx context.Context, formats strfmt.Registry) error {

	if m.VmwareAddSnapshotArg != nil {

		if swag.IsZero(m.VmwareAddSnapshotArg) { // not required
			return nil
		}

		if err := m.VmwareAddSnapshotArg.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmwareAddSnapshotArg")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmwareAddSnapshotArg")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddSnapshotArgExtendedInformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddSnapshotArgExtendedInformation) UnmarshalBinary(b []byte) error {
	var res AddSnapshotArgExtendedInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
