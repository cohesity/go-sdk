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

// VcenterRegistrationParams Register VMware vCenter request parameters.
//
// Specifies parameters to register VMware vCenter.
//
// swagger:model VcenterRegistrationParams
type VcenterRegistrationParams struct {
	CommonSourceRegistrationParams

	// Specifies the CA certificate to enable SSL communication between host and cluster.
	CaCert *string `json:"caCert,omitempty"`

	// Specifies to use VM BIOS UUID to track virtual machines in the host.
	UseVMBiosUUID *bool `json:"useVmBiosUuid,omitempty"`

	// Specifies the minimum free space (in GB) expected to be available in the datastore where the virtual disks of the VM being backed up reside. If the space available is lower than the specified value, backup will be aborted.
	MinFreeDatastoreSpaceForBackupGb *int64 `json:"minFreeDatastoreSpaceForBackupGb,omitempty"`

	// Specifies the minimum free space (in percentage) expected to be available in the datastore where the virtual disks of the VM being backed up reside. If the space available is lower than the specified value, backup will be aborted.
	MinFreeDatastoreSpaceForBackupPercentage *int64 `json:"minFreeDatastoreSpaceForBackupPercentage,omitempty"`

	// Specifies if the VM linking feature is enabled for the VCenter. If enabled, migrated VMs present in the VCenter which earlier belonged to some other VCenter will be linked during EH refresh.
	LinkVmsAcrossVcenter *bool `json:"linkVmsAcrossVcenter,omitempty"`

	// Specifies the throttling params.
	ThrottlingParams *VmwareThrottlingParams `json:"throttlingParams,omitempty"`

	// Specifies datastore specific parameters.
	DataStoreParams []*DatastoreParams `json:"dataStoreParams"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VcenterRegistrationParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonSourceRegistrationParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonSourceRegistrationParams = aO0

	// AO1
	var dataAO1 struct {
		CaCert *string `json:"caCert,omitempty"`

		UseVMBiosUUID *bool `json:"useVmBiosUuid,omitempty"`

		MinFreeDatastoreSpaceForBackupGb *int64 `json:"minFreeDatastoreSpaceForBackupGb,omitempty"`

		MinFreeDatastoreSpaceForBackupPercentage *int64 `json:"minFreeDatastoreSpaceForBackupPercentage,omitempty"`

		LinkVmsAcrossVcenter *bool `json:"linkVmsAcrossVcenter,omitempty"`

		ThrottlingParams *VmwareThrottlingParams `json:"throttlingParams,omitempty"`

		DataStoreParams []*DatastoreParams `json:"dataStoreParams"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.CaCert = dataAO1.CaCert

	m.UseVMBiosUUID = dataAO1.UseVMBiosUUID

	m.MinFreeDatastoreSpaceForBackupGb = dataAO1.MinFreeDatastoreSpaceForBackupGb

	m.MinFreeDatastoreSpaceForBackupPercentage = dataAO1.MinFreeDatastoreSpaceForBackupPercentage

	m.LinkVmsAcrossVcenter = dataAO1.LinkVmsAcrossVcenter

	m.ThrottlingParams = dataAO1.ThrottlingParams

	m.DataStoreParams = dataAO1.DataStoreParams

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VcenterRegistrationParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonSourceRegistrationParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		CaCert *string `json:"caCert,omitempty"`

		UseVMBiosUUID *bool `json:"useVmBiosUuid,omitempty"`

		MinFreeDatastoreSpaceForBackupGb *int64 `json:"minFreeDatastoreSpaceForBackupGb,omitempty"`

		MinFreeDatastoreSpaceForBackupPercentage *int64 `json:"minFreeDatastoreSpaceForBackupPercentage,omitempty"`

		LinkVmsAcrossVcenter *bool `json:"linkVmsAcrossVcenter,omitempty"`

		ThrottlingParams *VmwareThrottlingParams `json:"throttlingParams,omitempty"`

		DataStoreParams []*DatastoreParams `json:"dataStoreParams"`
	}

	dataAO1.CaCert = m.CaCert

	dataAO1.UseVMBiosUUID = m.UseVMBiosUUID

	dataAO1.MinFreeDatastoreSpaceForBackupGb = m.MinFreeDatastoreSpaceForBackupGb

	dataAO1.MinFreeDatastoreSpaceForBackupPercentage = m.MinFreeDatastoreSpaceForBackupPercentage

	dataAO1.LinkVmsAcrossVcenter = m.LinkVmsAcrossVcenter

	dataAO1.ThrottlingParams = m.ThrottlingParams

	dataAO1.DataStoreParams = m.DataStoreParams

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vcenter registration params
func (m *VcenterRegistrationParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonSourceRegistrationParams
	if err := m.CommonSourceRegistrationParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThrottlingParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataStoreParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcenterRegistrationParams) validateThrottlingParams(formats strfmt.Registry) error {

	if swag.IsZero(m.ThrottlingParams) { // not required
		return nil
	}

	if m.ThrottlingParams != nil {
		if err := m.ThrottlingParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throttlingParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throttlingParams")
			}
			return err
		}
	}

	return nil
}

func (m *VcenterRegistrationParams) validateDataStoreParams(formats strfmt.Registry) error {

	if swag.IsZero(m.DataStoreParams) { // not required
		return nil
	}

	for i := 0; i < len(m.DataStoreParams); i++ {
		if swag.IsZero(m.DataStoreParams[i]) { // not required
			continue
		}

		if m.DataStoreParams[i] != nil {
			if err := m.DataStoreParams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataStoreParams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dataStoreParams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this vcenter registration params based on the context it is used
func (m *VcenterRegistrationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonSourceRegistrationParams
	if err := m.CommonSourceRegistrationParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThrottlingParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataStoreParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcenterRegistrationParams) contextValidateThrottlingParams(ctx context.Context, formats strfmt.Registry) error {

	if m.ThrottlingParams != nil {

		if swag.IsZero(m.ThrottlingParams) { // not required
			return nil
		}

		if err := m.ThrottlingParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throttlingParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throttlingParams")
			}
			return err
		}
	}

	return nil
}

func (m *VcenterRegistrationParams) contextValidateDataStoreParams(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DataStoreParams); i++ {

		if m.DataStoreParams[i] != nil {

			if swag.IsZero(m.DataStoreParams[i]) { // not required
				return nil
			}

			if err := m.DataStoreParams[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataStoreParams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dataStoreParams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VcenterRegistrationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VcenterRegistrationParams) UnmarshalBinary(b []byte) error {
	var res VcenterRegistrationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
