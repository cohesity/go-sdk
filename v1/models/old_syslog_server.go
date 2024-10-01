// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OldSyslogServer Data definition related to Syslog Server configuration.
// Refer cluster_config.proto: ClusterConfigProto.TopicToRemoteLogServersMap
// Syslog Server Configuration.
//
// Specifies the syslog servers configuration to upload Cluster
// audit logs and filer audit logs.
//
// swagger:model OldSyslogServer
type OldSyslogServer struct {

	// Specifies the IP address or hostname of the syslog server.
	// Required: true
	Address *string `json:"address"`

	// Specifies if cohesity alert should be sent to syslog server
	// If 'true', alert audting message are sent to the server.
	// If 'false', alert auditng message are not sent to the server.(default)
	IsAlertAuditingEnabled *bool `json:"isAlertAuditingEnabled,omitempty"`

	// Specifies if Cluster audit logs should be sent to this syslog server.
	// If 'true', Cluster audit logs are sent to the syslog server. (default)
	// If 'false', Cluster audit logs are not sent to the syslog server.
	// Either cluster audit logs or filer audit logs should be enabled.
	IsClusterAuditingEnabled *bool `json:"isClusterAuditingEnabled,omitempty"`

	// Specifies if dataprotection  logs should be sent to syslog server
	// If 'true', dataprotection  logs are sent to the server.
	// If 'false', dataprotection  logs are not sent to the server.(default)
	IsDataProtectionEnabled *bool `json:"isDataProtectionEnabled,omitempty"`

	// Specifies if filer audit logs should be sent to this syslog server.
	// If 'true', filer audit logs are sent to the syslog server. (default)
	// If 'false', filer audit logs are not sent to the syslog server.
	// Either cluster audit logs or filer audit logs should be enabled.
	IsFilerAuditingEnabled *bool `json:"isFilerAuditingEnabled,omitempty"`

	// Specifies if ssh login logs should be sent to syslog server
	// If 'true', ssh login logs are sent to the server.
	// If 'false', ssh login logs are not sent to the server.(default)
	IsSSHLogEnabled *bool `json:"isSshLogEnabled,omitempty"`

	// Specifies a unique name for the syslog server on the Cluster.
	Name *string `json:"name,omitempty"`

	// Specifies the port where the syslog server listens.
	// Required: true
	Port *int32 `json:"port"`

	// Specifies the protocol used to send the logs.
	// Required: true
	Protocol *int32 `json:"protocol"`
}

// Validate validates this old syslog server
func (m *OldSyslogServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OldSyslogServer) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *OldSyslogServer) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *OldSyslogServer) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this old syslog server based on context it is used
func (m *OldSyslogServer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OldSyslogServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OldSyslogServer) UnmarshalBinary(b []byte) error {
	var res OldSyslogServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
