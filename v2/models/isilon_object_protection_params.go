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

// IsilonObjectProtectionParams Specifies the parameters which are specific to Isilon object protection.
//
// swagger:model IsilonObjectProtectionParams
type IsilonObjectProtectionParams struct {

	// Specifies the protocol of the NAS device being backed up.
	// Enum: ["kNoProtocol","kNfs3","kNfs4_1","kCifs1","kCifs2","kCifs3"]
	Protocol *string `json:"protocol,omitempty"`

	// Specifies the source snapshots to be taken even if there is a pending run in a protection group.
	ContinuousSnapshots *ContinuousSnapshotParams `json:"continuousSnapshots,omitempty"`

	// Specify whether to use the Isilon Changelist API to directly discover changed files/directories for faster incremental backup. Cohesity will keep an extra snapshot which will be deleted by the next successful backup.
	UseChangelist *bool `json:"useChangelist,omitempty"`
}

// Validate validates this isilon object protection params
func (m *IsilonObjectProtectionParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContinuousSnapshots(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var isilonObjectProtectionParamsTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kNoProtocol","kNfs3","kNfs4_1","kCifs1","kCifs2","kCifs3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		isilonObjectProtectionParamsTypeProtocolPropEnum = append(isilonObjectProtectionParamsTypeProtocolPropEnum, v)
	}
}

const (

	// IsilonObjectProtectionParamsProtocolKNoProtocol captures enum value "kNoProtocol"
	IsilonObjectProtectionParamsProtocolKNoProtocol string = "kNoProtocol"

	// IsilonObjectProtectionParamsProtocolKNfs3 captures enum value "kNfs3"
	IsilonObjectProtectionParamsProtocolKNfs3 string = "kNfs3"

	// IsilonObjectProtectionParamsProtocolKNfs41 captures enum value "kNfs4_1"
	IsilonObjectProtectionParamsProtocolKNfs41 string = "kNfs4_1"

	// IsilonObjectProtectionParamsProtocolKCifs1 captures enum value "kCifs1"
	IsilonObjectProtectionParamsProtocolKCifs1 string = "kCifs1"

	// IsilonObjectProtectionParamsProtocolKCifs2 captures enum value "kCifs2"
	IsilonObjectProtectionParamsProtocolKCifs2 string = "kCifs2"

	// IsilonObjectProtectionParamsProtocolKCifs3 captures enum value "kCifs3"
	IsilonObjectProtectionParamsProtocolKCifs3 string = "kCifs3"
)

// prop value enum
func (m *IsilonObjectProtectionParams) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, isilonObjectProtectionParamsTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IsilonObjectProtectionParams) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *IsilonObjectProtectionParams) validateContinuousSnapshots(formats strfmt.Registry) error {
	if swag.IsZero(m.ContinuousSnapshots) { // not required
		return nil
	}

	if m.ContinuousSnapshots != nil {
		if err := m.ContinuousSnapshots.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("continuousSnapshots")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("continuousSnapshots")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this isilon object protection params based on the context it is used
func (m *IsilonObjectProtectionParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContinuousSnapshots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IsilonObjectProtectionParams) contextValidateContinuousSnapshots(ctx context.Context, formats strfmt.Registry) error {

	if m.ContinuousSnapshots != nil {

		if swag.IsZero(m.ContinuousSnapshots) { // not required
			return nil
		}

		if err := m.ContinuousSnapshots.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("continuousSnapshots")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("continuousSnapshots")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IsilonObjectProtectionParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IsilonObjectProtectionParams) UnmarshalBinary(b []byte) error {
	var res IsilonObjectProtectionParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
