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

// RecoverKubernetesNamespaceParams Recover Kubernetes Namespace params.
//
// Specifies the parameters to recover Kubernetes Namespaces.
//
// swagger:model RecoverKubernetesNamespaceParams
type RecoverKubernetesNamespaceParams struct {

	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	// Required: true
	// Enum: ["kKubernetes"]
	TargetEnvironment *string `json:"targetEnvironment"`

	// Specifies the params for recovering to a Kubernetes host.
	KubernetesTargetParams *KubernetesTargetParamsForRecoverKubernetesNamespace `json:"kubernetesTargetParams,omitempty"`

	// Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
	VlanConfig *RecoveryVlanConfig `json:"vlanConfig,omitempty"`
}

// Validate validates this recover kubernetes namespace params
func (m *RecoverKubernetesNamespaceParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTargetEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetesTargetParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var recoverKubernetesNamespaceParamsTypeTargetEnvironmentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kKubernetes"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recoverKubernetesNamespaceParamsTypeTargetEnvironmentPropEnum = append(recoverKubernetesNamespaceParamsTypeTargetEnvironmentPropEnum, v)
	}
}

const (

	// RecoverKubernetesNamespaceParamsTargetEnvironmentKKubernetes captures enum value "kKubernetes"
	RecoverKubernetesNamespaceParamsTargetEnvironmentKKubernetes string = "kKubernetes"
)

// prop value enum
func (m *RecoverKubernetesNamespaceParams) validateTargetEnvironmentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recoverKubernetesNamespaceParamsTypeTargetEnvironmentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecoverKubernetesNamespaceParams) validateTargetEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("targetEnvironment", "body", m.TargetEnvironment); err != nil {
		return err
	}

	// value enum
	if err := m.validateTargetEnvironmentEnum("targetEnvironment", "body", *m.TargetEnvironment); err != nil {
		return err
	}

	return nil
}

func (m *RecoverKubernetesNamespaceParams) validateKubernetesTargetParams(formats strfmt.Registry) error {
	if swag.IsZero(m.KubernetesTargetParams) { // not required
		return nil
	}

	if m.KubernetesTargetParams != nil {
		if err := m.KubernetesTargetParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubernetesTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubernetesTargetParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverKubernetesNamespaceParams) validateVlanConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.VlanConfig) { // not required
		return nil
	}

	if m.VlanConfig != nil {
		if err := m.VlanConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlanConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlanConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover kubernetes namespace params based on the context it is used
func (m *RecoverKubernetesNamespaceParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKubernetesTargetParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlanConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverKubernetesNamespaceParams) contextValidateKubernetesTargetParams(ctx context.Context, formats strfmt.Registry) error {

	if m.KubernetesTargetParams != nil {

		if swag.IsZero(m.KubernetesTargetParams) { // not required
			return nil
		}

		if err := m.KubernetesTargetParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubernetesTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubernetesTargetParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverKubernetesNamespaceParams) contextValidateVlanConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.VlanConfig != nil {

		if swag.IsZero(m.VlanConfig) { // not required
			return nil
		}

		if err := m.VlanConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlanConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlanConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverKubernetesNamespaceParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverKubernetesNamespaceParams) UnmarshalBinary(b []byte) error {
	var res RecoverKubernetesNamespaceParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
