// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PrivateHdfsConnectParams Contains specific parameters required to connect to an hdfs cluster.
//
// swagger:model PrivateHdfsConnectParams
type PrivateHdfsConnectParams struct {

	// Hdfs connection configuration.
	Configuration []*HdfsConnectParamsConfigurationEntry `json:"configuration"`

	// hadoop distribution
	HadoopDistribution *int32 `json:"hadoopDistribution,omitempty"`

	// Hadoop version
	HadoopVersion *string `json:"hadoopVersion,omitempty"`

	// Authentication type
	HdfsAuthType *int32 `json:"hdfsAuthType,omitempty"`

	// Kerberos principal
	HdfsPrincipal *string `json:"hdfsPrincipal,omitempty"`

	// Namenode host or Nameservice.
	Namenode *string `json:"namenode,omitempty"`

	// Webhdfs Port
	Port *int32 `json:"port,omitempty"`
}

// Validate validates this private hdfs connect params
func (m *PrivateHdfsConnectParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateHdfsConnectParams) validateConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.Configuration) { // not required
		return nil
	}

	for i := 0; i < len(m.Configuration); i++ {
		if swag.IsZero(m.Configuration[i]) { // not required
			continue
		}

		if m.Configuration[i] != nil {
			if err := m.Configuration[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configuration" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("configuration" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this private hdfs connect params based on the context it is used
func (m *PrivateHdfsConnectParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateHdfsConnectParams) contextValidateConfiguration(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Configuration); i++ {

		if m.Configuration[i] != nil {

			if swag.IsZero(m.Configuration[i]) { // not required
				return nil
			}

			if err := m.Configuration[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configuration" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("configuration" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrivateHdfsConnectParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateHdfsConnectParams) UnmarshalBinary(b []byte) error {
	var res PrivateHdfsConnectParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
