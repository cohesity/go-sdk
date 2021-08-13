/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package helios
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/helios/utils"
)

var _ = NullableBool{}

// RecoverVmwareParams Specifies the recovery options specific to VMware environment.
type RecoverVmwareParams struct {
	// Specifies the list of recover Object parameters. This property is mandatory for all recovery action types except recover vms. While recovering VMs, a user can specify snapshots of VM's or a Protection Group Run details to recover all the VM's that are backed up by that Run. For recovering files, specifies the object contains the file to recover.
	Objects []RecoverVmwareSnapshotParams `json:"objects,omitempty"`
	// Specifies the type of recovery action to be performed.
	RecoveryAction string `json:"recoveryAction"`
	// Specifies the parameters to recover VMware VM.
	RecoverVmParams NullableRecoverVmwareVmParams `json:"recoverVmParams,omitempty"`
	// Specifies the parameters to recover VMware Disks.
	RecoverVmDiskParams NullableRecoverVmwareDiskParams `json:"recoverVmDiskParams,omitempty"`
	// Specifies the parameters to mount VMware Volumes.
	MountVolumeParams NullableMountVmwareVolumeParams `json:"mountVolumeParams,omitempty"`
	// Specifies the parameters to recover a VMware vApp.
	RecoverVAppParams NullableRecoverVmwareVAppParams `json:"recoverVAppParams,omitempty"`
	// Specifies the parameters to recover a VMware vApp template.
	RecoverVAppTemplateParams NullableRecoverVmwareVAppTemplateParams `json:"recoverVAppTemplateParams,omitempty"`
	// Specifies the parameters to recover files and folders.
	RecoverFileAndFolderParams NullableRecoverVMwareFileAndFolderParams `json:"recoverFileAndFolderParams,omitempty"`
	// Specifies the parameters to download files and folders.
	DownloadFileAndFolderParams NullableCommonDownloadFileAndFolderParams `json:"downloadFileAndFolderParams,omitempty"`
}

// NewRecoverVmwareParams instantiates a new RecoverVmwareParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverVmwareParams(recoveryAction string) *RecoverVmwareParams {
	this := RecoverVmwareParams{}
	this.RecoveryAction = recoveryAction
	return &this
}

// NewRecoverVmwareParamsWithDefaults instantiates a new RecoverVmwareParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverVmwareParamsWithDefaults() *RecoverVmwareParams {
	this := RecoverVmwareParams{}
	return &this
}

// GetObjects returns the Objects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverVmwareParams) GetObjects() []RecoverVmwareSnapshotParams {
	if o == nil  {
		var ret []RecoverVmwareSnapshotParams
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverVmwareParams) GetObjectsOk() (*[]RecoverVmwareSnapshotParams, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return &o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *RecoverVmwareParams) HasObjects() bool {
	if o != nil && o.Objects != nil {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []RecoverVmwareSnapshotParams and assigns it to the Objects field.
func (o *RecoverVmwareParams) SetObjects(v []RecoverVmwareSnapshotParams) {
	o.Objects = v
}

// GetRecoveryAction returns the RecoveryAction field value
func (o *RecoverVmwareParams) GetRecoveryAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecoveryAction
}

// GetRecoveryActionOk returns a tuple with the RecoveryAction field value
// and a boolean to check if the value has been set.
func (o *RecoverVmwareParams) GetRecoveryActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecoveryAction, true
}

// SetRecoveryAction sets field value
func (o *RecoverVmwareParams) SetRecoveryAction(v string) {
	o.RecoveryAction = v
}

// GetRecoverVmParams returns the RecoverVmParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverVmwareParams) GetRecoverVmParams() RecoverVmwareVmParams {
	if o == nil || o.RecoverVmParams.Get() == nil {
		var ret RecoverVmwareVmParams
		return ret
	}
	return *o.RecoverVmParams.Get()
}

// GetRecoverVmParamsOk returns a tuple with the RecoverVmParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverVmwareParams) GetRecoverVmParamsOk() (*RecoverVmwareVmParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverVmParams.Get(), o.RecoverVmParams.IsSet()
}

// HasRecoverVmParams returns a boolean if a field has been set.
func (o *RecoverVmwareParams) HasRecoverVmParams() bool {
	if o != nil && o.RecoverVmParams.IsSet() {
		return true
	}

	return false
}

// SetRecoverVmParams gets a reference to the given NullableRecoverVmwareVmParams and assigns it to the RecoverVmParams field.
func (o *RecoverVmwareParams) SetRecoverVmParams(v RecoverVmwareVmParams) {
	o.RecoverVmParams.Set(&v)
}
// SetRecoverVmParamsNil sets the value for RecoverVmParams to be an explicit nil
func (o *RecoverVmwareParams) SetRecoverVmParamsNil() {
	o.RecoverVmParams.Set(nil)
}

// UnsetRecoverVmParams ensures that no value is present for RecoverVmParams, not even an explicit nil
func (o *RecoverVmwareParams) UnsetRecoverVmParams() {
	o.RecoverVmParams.Unset()
}

// GetRecoverVmDiskParams returns the RecoverVmDiskParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverVmwareParams) GetRecoverVmDiskParams() RecoverVmwareDiskParams {
	if o == nil || o.RecoverVmDiskParams.Get() == nil {
		var ret RecoverVmwareDiskParams
		return ret
	}
	return *o.RecoverVmDiskParams.Get()
}

// GetRecoverVmDiskParamsOk returns a tuple with the RecoverVmDiskParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverVmwareParams) GetRecoverVmDiskParamsOk() (*RecoverVmwareDiskParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverVmDiskParams.Get(), o.RecoverVmDiskParams.IsSet()
}

// HasRecoverVmDiskParams returns a boolean if a field has been set.
func (o *RecoverVmwareParams) HasRecoverVmDiskParams() bool {
	if o != nil && o.RecoverVmDiskParams.IsSet() {
		return true
	}

	return false
}

// SetRecoverVmDiskParams gets a reference to the given NullableRecoverVmwareDiskParams and assigns it to the RecoverVmDiskParams field.
func (o *RecoverVmwareParams) SetRecoverVmDiskParams(v RecoverVmwareDiskParams) {
	o.RecoverVmDiskParams.Set(&v)
}
// SetRecoverVmDiskParamsNil sets the value for RecoverVmDiskParams to be an explicit nil
func (o *RecoverVmwareParams) SetRecoverVmDiskParamsNil() {
	o.RecoverVmDiskParams.Set(nil)
}

// UnsetRecoverVmDiskParams ensures that no value is present for RecoverVmDiskParams, not even an explicit nil
func (o *RecoverVmwareParams) UnsetRecoverVmDiskParams() {
	o.RecoverVmDiskParams.Unset()
}

// GetMountVolumeParams returns the MountVolumeParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverVmwareParams) GetMountVolumeParams() MountVmwareVolumeParams {
	if o == nil || o.MountVolumeParams.Get() == nil {
		var ret MountVmwareVolumeParams
		return ret
	}
	return *o.MountVolumeParams.Get()
}

// GetMountVolumeParamsOk returns a tuple with the MountVolumeParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverVmwareParams) GetMountVolumeParamsOk() (*MountVmwareVolumeParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MountVolumeParams.Get(), o.MountVolumeParams.IsSet()
}

// HasMountVolumeParams returns a boolean if a field has been set.
func (o *RecoverVmwareParams) HasMountVolumeParams() bool {
	if o != nil && o.MountVolumeParams.IsSet() {
		return true
	}

	return false
}

// SetMountVolumeParams gets a reference to the given NullableMountVmwareVolumeParams and assigns it to the MountVolumeParams field.
func (o *RecoverVmwareParams) SetMountVolumeParams(v MountVmwareVolumeParams) {
	o.MountVolumeParams.Set(&v)
}
// SetMountVolumeParamsNil sets the value for MountVolumeParams to be an explicit nil
func (o *RecoverVmwareParams) SetMountVolumeParamsNil() {
	o.MountVolumeParams.Set(nil)
}

// UnsetMountVolumeParams ensures that no value is present for MountVolumeParams, not even an explicit nil
func (o *RecoverVmwareParams) UnsetMountVolumeParams() {
	o.MountVolumeParams.Unset()
}

// GetRecoverVAppParams returns the RecoverVAppParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverVmwareParams) GetRecoverVAppParams() RecoverVmwareVAppParams {
	if o == nil || o.RecoverVAppParams.Get() == nil {
		var ret RecoverVmwareVAppParams
		return ret
	}
	return *o.RecoverVAppParams.Get()
}

// GetRecoverVAppParamsOk returns a tuple with the RecoverVAppParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverVmwareParams) GetRecoverVAppParamsOk() (*RecoverVmwareVAppParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverVAppParams.Get(), o.RecoverVAppParams.IsSet()
}

// HasRecoverVAppParams returns a boolean if a field has been set.
func (o *RecoverVmwareParams) HasRecoverVAppParams() bool {
	if o != nil && o.RecoverVAppParams.IsSet() {
		return true
	}

	return false
}

// SetRecoverVAppParams gets a reference to the given NullableRecoverVmwareVAppParams and assigns it to the RecoverVAppParams field.
func (o *RecoverVmwareParams) SetRecoverVAppParams(v RecoverVmwareVAppParams) {
	o.RecoverVAppParams.Set(&v)
}
// SetRecoverVAppParamsNil sets the value for RecoverVAppParams to be an explicit nil
func (o *RecoverVmwareParams) SetRecoverVAppParamsNil() {
	o.RecoverVAppParams.Set(nil)
}

// UnsetRecoverVAppParams ensures that no value is present for RecoverVAppParams, not even an explicit nil
func (o *RecoverVmwareParams) UnsetRecoverVAppParams() {
	o.RecoverVAppParams.Unset()
}

// GetRecoverVAppTemplateParams returns the RecoverVAppTemplateParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverVmwareParams) GetRecoverVAppTemplateParams() RecoverVmwareVAppTemplateParams {
	if o == nil || o.RecoverVAppTemplateParams.Get() == nil {
		var ret RecoverVmwareVAppTemplateParams
		return ret
	}
	return *o.RecoverVAppTemplateParams.Get()
}

// GetRecoverVAppTemplateParamsOk returns a tuple with the RecoverVAppTemplateParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverVmwareParams) GetRecoverVAppTemplateParamsOk() (*RecoverVmwareVAppTemplateParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverVAppTemplateParams.Get(), o.RecoverVAppTemplateParams.IsSet()
}

// HasRecoverVAppTemplateParams returns a boolean if a field has been set.
func (o *RecoverVmwareParams) HasRecoverVAppTemplateParams() bool {
	if o != nil && o.RecoverVAppTemplateParams.IsSet() {
		return true
	}

	return false
}

// SetRecoverVAppTemplateParams gets a reference to the given NullableRecoverVmwareVAppTemplateParams and assigns it to the RecoverVAppTemplateParams field.
func (o *RecoverVmwareParams) SetRecoverVAppTemplateParams(v RecoverVmwareVAppTemplateParams) {
	o.RecoverVAppTemplateParams.Set(&v)
}
// SetRecoverVAppTemplateParamsNil sets the value for RecoverVAppTemplateParams to be an explicit nil
func (o *RecoverVmwareParams) SetRecoverVAppTemplateParamsNil() {
	o.RecoverVAppTemplateParams.Set(nil)
}

// UnsetRecoverVAppTemplateParams ensures that no value is present for RecoverVAppTemplateParams, not even an explicit nil
func (o *RecoverVmwareParams) UnsetRecoverVAppTemplateParams() {
	o.RecoverVAppTemplateParams.Unset()
}

// GetRecoverFileAndFolderParams returns the RecoverFileAndFolderParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverVmwareParams) GetRecoverFileAndFolderParams() RecoverVMwareFileAndFolderParams {
	if o == nil || o.RecoverFileAndFolderParams.Get() == nil {
		var ret RecoverVMwareFileAndFolderParams
		return ret
	}
	return *o.RecoverFileAndFolderParams.Get()
}

// GetRecoverFileAndFolderParamsOk returns a tuple with the RecoverFileAndFolderParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverVmwareParams) GetRecoverFileAndFolderParamsOk() (*RecoverVMwareFileAndFolderParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverFileAndFolderParams.Get(), o.RecoverFileAndFolderParams.IsSet()
}

// HasRecoverFileAndFolderParams returns a boolean if a field has been set.
func (o *RecoverVmwareParams) HasRecoverFileAndFolderParams() bool {
	if o != nil && o.RecoverFileAndFolderParams.IsSet() {
		return true
	}

	return false
}

// SetRecoverFileAndFolderParams gets a reference to the given NullableRecoverVMwareFileAndFolderParams and assigns it to the RecoverFileAndFolderParams field.
func (o *RecoverVmwareParams) SetRecoverFileAndFolderParams(v RecoverVMwareFileAndFolderParams) {
	o.RecoverFileAndFolderParams.Set(&v)
}
// SetRecoverFileAndFolderParamsNil sets the value for RecoverFileAndFolderParams to be an explicit nil
func (o *RecoverVmwareParams) SetRecoverFileAndFolderParamsNil() {
	o.RecoverFileAndFolderParams.Set(nil)
}

// UnsetRecoverFileAndFolderParams ensures that no value is present for RecoverFileAndFolderParams, not even an explicit nil
func (o *RecoverVmwareParams) UnsetRecoverFileAndFolderParams() {
	o.RecoverFileAndFolderParams.Unset()
}

// GetDownloadFileAndFolderParams returns the DownloadFileAndFolderParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverVmwareParams) GetDownloadFileAndFolderParams() CommonDownloadFileAndFolderParams {
	if o == nil || o.DownloadFileAndFolderParams.Get() == nil {
		var ret CommonDownloadFileAndFolderParams
		return ret
	}
	return *o.DownloadFileAndFolderParams.Get()
}

// GetDownloadFileAndFolderParamsOk returns a tuple with the DownloadFileAndFolderParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverVmwareParams) GetDownloadFileAndFolderParamsOk() (*CommonDownloadFileAndFolderParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DownloadFileAndFolderParams.Get(), o.DownloadFileAndFolderParams.IsSet()
}

// HasDownloadFileAndFolderParams returns a boolean if a field has been set.
func (o *RecoverVmwareParams) HasDownloadFileAndFolderParams() bool {
	if o != nil && o.DownloadFileAndFolderParams.IsSet() {
		return true
	}

	return false
}

// SetDownloadFileAndFolderParams gets a reference to the given NullableCommonDownloadFileAndFolderParams and assigns it to the DownloadFileAndFolderParams field.
func (o *RecoverVmwareParams) SetDownloadFileAndFolderParams(v CommonDownloadFileAndFolderParams) {
	o.DownloadFileAndFolderParams.Set(&v)
}
// SetDownloadFileAndFolderParamsNil sets the value for DownloadFileAndFolderParams to be an explicit nil
func (o *RecoverVmwareParams) SetDownloadFileAndFolderParamsNil() {
	o.DownloadFileAndFolderParams.Set(nil)
}

// UnsetDownloadFileAndFolderParams ensures that no value is present for DownloadFileAndFolderParams, not even an explicit nil
func (o *RecoverVmwareParams) UnsetDownloadFileAndFolderParams() {
	o.DownloadFileAndFolderParams.Unset()
}

func (o RecoverVmwareParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if true {
		toSerialize["recoveryAction"] = o.RecoveryAction
	}
	if o.RecoverVmParams.IsSet() {
		toSerialize["recoverVmParams"] = o.RecoverVmParams.Get()
	}
	if o.RecoverVmDiskParams.IsSet() {
		toSerialize["recoverVmDiskParams"] = o.RecoverVmDiskParams.Get()
	}
	if o.MountVolumeParams.IsSet() {
		toSerialize["mountVolumeParams"] = o.MountVolumeParams.Get()
	}
	if o.RecoverVAppParams.IsSet() {
		toSerialize["recoverVAppParams"] = o.RecoverVAppParams.Get()
	}
	if o.RecoverVAppTemplateParams.IsSet() {
		toSerialize["recoverVAppTemplateParams"] = o.RecoverVAppTemplateParams.Get()
	}
	if o.RecoverFileAndFolderParams.IsSet() {
		toSerialize["recoverFileAndFolderParams"] = o.RecoverFileAndFolderParams.Get()
	}
	if o.DownloadFileAndFolderParams.IsSet() {
		toSerialize["downloadFileAndFolderParams"] = o.DownloadFileAndFolderParams.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverVmwareParams struct {
	value *RecoverVmwareParams
	isSet bool
}

func (v NullableRecoverVmwareParams) Get() *RecoverVmwareParams {
	return v.value
}

func (v *NullableRecoverVmwareParams) Set(val *RecoverVmwareParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverVmwareParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverVmwareParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverVmwareParams(val *RecoverVmwareParams) *NullableRecoverVmwareParams {
	return &NullableRecoverVmwareParams{value: val, isSet: true}
}

func (v NullableRecoverVmwareParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverVmwareParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


