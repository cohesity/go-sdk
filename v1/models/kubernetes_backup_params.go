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

// KubernetesBackupParams This message contains additional params for Kubernetes backups.
//
// swagger:model KubernetesBackupParams
type KubernetesBackupParams struct {

	// Cluster software version.
	ClusterSoftwareVersion *string `json:"clusterSoftwareVersion,omitempty"`

	// Indicates the kubernetes service type to use.
	DatamoverServiceType *int32 `json:"datamoverServiceType,omitempty"`

	// Container image used to mounting PVCs in temp pods.
	InitContainerImage *string `json:"initContainerImage,omitempty"`

	// This indicates if magneto_kubernetes_enable_protection_using_datamover
	// is true and the flag is enabled in the feature enabler.
	IsProtectionUsingDatamoverEnabled *bool `json:"isProtectionUsingDatamoverEnabled,omitempty"`

	// Indicates if volume inclusion/exclusion feature is enabled in the feature
	// enabler.
	IsVolumeExclusionEnabled *bool `json:"isVolumeExclusionEnabled,omitempty"`

	// Managment namespace is used to create all the backup jobs. This namespace
	// is the umbrella for all backup operations. We backup applications in their
	// respective namespace.
	ManagementNamespace *string `json:"managementNamespace,omitempty"`

	// S3 account ID of the user using which backup S3 buckets will be created.
	S3AccountID *string `json:"s3AccountId,omitempty"`

	// Indicates whether to use annotated volume snapshot class while taking
	// volume snapshot.
	UseAnnotatedVsc *bool `json:"useAnnotatedVsc,omitempty"`

	// Indicates the VLAN information to use for backup.
	VlanParams *VlanParams `json:"vlanParams,omitempty"`
}

// Validate validates this kubernetes backup params
func (m *KubernetesBackupParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVlanParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesBackupParams) validateVlanParams(formats strfmt.Registry) error {
	if swag.IsZero(m.VlanParams) { // not required
		return nil
	}

	if m.VlanParams != nil {
		if err := m.VlanParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlanParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlanParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kubernetes backup params based on the context it is used
func (m *KubernetesBackupParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVlanParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesBackupParams) contextValidateVlanParams(ctx context.Context, formats strfmt.Registry) error {

	if m.VlanParams != nil {

		if swag.IsZero(m.VlanParams) { // not required
			return nil
		}

		if err := m.VlanParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlanParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlanParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesBackupParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesBackupParams) UnmarshalBinary(b []byte) error {
	var res KubernetesBackupParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
