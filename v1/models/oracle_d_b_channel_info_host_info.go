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

// OracleDBChannelInfoHostInfo The name of this proto message is out-dated. This proto should
// generally be used to represent parameters needed for each Oracle
// 'cluster' node. 'cluster' here is a loose term used to include
// more than Oracle RAC cluster, e.g. 'active-passive' cluster is also
// considered here as 'cluster' and its 'cluster node will also be
// represented by the following proto.
//
// swagger:model OracleDBChannelInfo_HostInfo
type OracleDBChannelInfoHostInfo struct {

	// 'agent_id' of the host from which we are allowed to take the
	// backup/restore.
	Host *string `json:"host,omitempty"`

	// Number of channels we need to create for this host. Default value for
	// num_channels will be calculated as minimum of number of nodes in
	// cohesity cluster, 2 * number of cpu on Oracle host.
	NumChannels *int32 `json:"numChannels,omitempty"`

	// port number where database is listening.
	Portnum *int64 `json:"portnum,omitempty"`

	// The necessury parameters for SBT. This is set only when backup type is
	// kSbt.
	SbtHostParams *OracleSbtHostParams `json:"sbtHostParams,omitempty"`
}

// Validate validates this oracle d b channel info host info
func (m *OracleDBChannelInfoHostInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSbtHostParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OracleDBChannelInfoHostInfo) validateSbtHostParams(formats strfmt.Registry) error {
	if swag.IsZero(m.SbtHostParams) { // not required
		return nil
	}

	if m.SbtHostParams != nil {
		if err := m.SbtHostParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sbtHostParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sbtHostParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this oracle d b channel info host info based on the context it is used
func (m *OracleDBChannelInfoHostInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSbtHostParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OracleDBChannelInfoHostInfo) contextValidateSbtHostParams(ctx context.Context, formats strfmt.Registry) error {

	if m.SbtHostParams != nil {

		if swag.IsZero(m.SbtHostParams) { // not required
			return nil
		}

		if err := m.SbtHostParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sbtHostParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sbtHostParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OracleDBChannelInfoHostInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleDBChannelInfoHostInfo) UnmarshalBinary(b []byte) error {
	var res OracleDBChannelInfoHostInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
