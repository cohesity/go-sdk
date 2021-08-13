/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// RecoverAwsParams Specifies the recovery options specific to AWS environment.
type RecoverAwsParams struct {
	// Specifies the list of recover Object parameters. This property is mandatory for all recovery action types except recover vms. While recovering VMs, a user can specify snapshots of VM's or a Protection Group Run details to recover all the VM's that are backed up by that Run. For recovering files, specifies the object contains the file to recover.
	Objects []CommonRecoverObjectSnapshotParams `json:"objects,omitempty"`
	// Specifies the type of recover action to be performed.
	RecoveryAction string `json:"recoveryAction"`
	// Specifies the parameters to recover AWS VM.
	RecoverVmParams NullableRecoverAwsVmParams `json:"recoverVmParams,omitempty"`
	// Specifies the parameters to recover files and folders.
	RecoverFileAndFolderParams NullableRecoverAwsFileAndFolderParams `json:"recoverFileAndFolderParams,omitempty"`
	// Specifies the parameters to recover AWS RDS.
	RecoverRdsParams NullableRecoverAwsRdsParams `json:"recoverRdsParams,omitempty"`
	// Specifies the parameters to recover AWS Aurora.
	RecoverAuroraParams NullableRecoverAwsAuroraParams `json:"recoverAuroraParams,omitempty"`
	// Specifies the parameters to download files and folders.
	DownloadFileAndFolderParams NullableCommonDownloadFileAndFolderParams `json:"downloadFileAndFolderParams,omitempty"`
}

// NewRecoverAwsParams instantiates a new RecoverAwsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverAwsParams(recoveryAction string) *RecoverAwsParams {
	this := RecoverAwsParams{}
	this.RecoveryAction = recoveryAction
	return &this
}

// NewRecoverAwsParamsWithDefaults instantiates a new RecoverAwsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverAwsParamsWithDefaults() *RecoverAwsParams {
	this := RecoverAwsParams{}
	return &this
}

// GetObjects returns the Objects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverAwsParams) GetObjects() []CommonRecoverObjectSnapshotParams {
	if o == nil  {
		var ret []CommonRecoverObjectSnapshotParams
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAwsParams) GetObjectsOk() (*[]CommonRecoverObjectSnapshotParams, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return &o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *RecoverAwsParams) HasObjects() bool {
	if o != nil && o.Objects != nil {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []CommonRecoverObjectSnapshotParams and assigns it to the Objects field.
func (o *RecoverAwsParams) SetObjects(v []CommonRecoverObjectSnapshotParams) {
	o.Objects = v
}

// GetRecoveryAction returns the RecoveryAction field value
func (o *RecoverAwsParams) GetRecoveryAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecoveryAction
}

// GetRecoveryActionOk returns a tuple with the RecoveryAction field value
// and a boolean to check if the value has been set.
func (o *RecoverAwsParams) GetRecoveryActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecoveryAction, true
}

// SetRecoveryAction sets field value
func (o *RecoverAwsParams) SetRecoveryAction(v string) {
	o.RecoveryAction = v
}

// GetRecoverVmParams returns the RecoverVmParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverAwsParams) GetRecoverVmParams() RecoverAwsVmParams {
	if o == nil || o.RecoverVmParams.Get() == nil {
		var ret RecoverAwsVmParams
		return ret
	}
	return *o.RecoverVmParams.Get()
}

// GetRecoverVmParamsOk returns a tuple with the RecoverVmParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAwsParams) GetRecoverVmParamsOk() (*RecoverAwsVmParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverVmParams.Get(), o.RecoverVmParams.IsSet()
}

// HasRecoverVmParams returns a boolean if a field has been set.
func (o *RecoverAwsParams) HasRecoverVmParams() bool {
	if o != nil && o.RecoverVmParams.IsSet() {
		return true
	}

	return false
}

// SetRecoverVmParams gets a reference to the given NullableRecoverAwsVmParams and assigns it to the RecoverVmParams field.
func (o *RecoverAwsParams) SetRecoverVmParams(v RecoverAwsVmParams) {
	o.RecoverVmParams.Set(&v)
}
// SetRecoverVmParamsNil sets the value for RecoverVmParams to be an explicit nil
func (o *RecoverAwsParams) SetRecoverVmParamsNil() {
	o.RecoverVmParams.Set(nil)
}

// UnsetRecoverVmParams ensures that no value is present for RecoverVmParams, not even an explicit nil
func (o *RecoverAwsParams) UnsetRecoverVmParams() {
	o.RecoverVmParams.Unset()
}

// GetRecoverFileAndFolderParams returns the RecoverFileAndFolderParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverAwsParams) GetRecoverFileAndFolderParams() RecoverAwsFileAndFolderParams {
	if o == nil || o.RecoverFileAndFolderParams.Get() == nil {
		var ret RecoverAwsFileAndFolderParams
		return ret
	}
	return *o.RecoverFileAndFolderParams.Get()
}

// GetRecoverFileAndFolderParamsOk returns a tuple with the RecoverFileAndFolderParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAwsParams) GetRecoverFileAndFolderParamsOk() (*RecoverAwsFileAndFolderParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverFileAndFolderParams.Get(), o.RecoverFileAndFolderParams.IsSet()
}

// HasRecoverFileAndFolderParams returns a boolean if a field has been set.
func (o *RecoverAwsParams) HasRecoverFileAndFolderParams() bool {
	if o != nil && o.RecoverFileAndFolderParams.IsSet() {
		return true
	}

	return false
}

// SetRecoverFileAndFolderParams gets a reference to the given NullableRecoverAwsFileAndFolderParams and assigns it to the RecoverFileAndFolderParams field.
func (o *RecoverAwsParams) SetRecoverFileAndFolderParams(v RecoverAwsFileAndFolderParams) {
	o.RecoverFileAndFolderParams.Set(&v)
}
// SetRecoverFileAndFolderParamsNil sets the value for RecoverFileAndFolderParams to be an explicit nil
func (o *RecoverAwsParams) SetRecoverFileAndFolderParamsNil() {
	o.RecoverFileAndFolderParams.Set(nil)
}

// UnsetRecoverFileAndFolderParams ensures that no value is present for RecoverFileAndFolderParams, not even an explicit nil
func (o *RecoverAwsParams) UnsetRecoverFileAndFolderParams() {
	o.RecoverFileAndFolderParams.Unset()
}

// GetRecoverRdsParams returns the RecoverRdsParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverAwsParams) GetRecoverRdsParams() RecoverAwsRdsParams {
	if o == nil || o.RecoverRdsParams.Get() == nil {
		var ret RecoverAwsRdsParams
		return ret
	}
	return *o.RecoverRdsParams.Get()
}

// GetRecoverRdsParamsOk returns a tuple with the RecoverRdsParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAwsParams) GetRecoverRdsParamsOk() (*RecoverAwsRdsParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverRdsParams.Get(), o.RecoverRdsParams.IsSet()
}

// HasRecoverRdsParams returns a boolean if a field has been set.
func (o *RecoverAwsParams) HasRecoverRdsParams() bool {
	if o != nil && o.RecoverRdsParams.IsSet() {
		return true
	}

	return false
}

// SetRecoverRdsParams gets a reference to the given NullableRecoverAwsRdsParams and assigns it to the RecoverRdsParams field.
func (o *RecoverAwsParams) SetRecoverRdsParams(v RecoverAwsRdsParams) {
	o.RecoverRdsParams.Set(&v)
}
// SetRecoverRdsParamsNil sets the value for RecoverRdsParams to be an explicit nil
func (o *RecoverAwsParams) SetRecoverRdsParamsNil() {
	o.RecoverRdsParams.Set(nil)
}

// UnsetRecoverRdsParams ensures that no value is present for RecoverRdsParams, not even an explicit nil
func (o *RecoverAwsParams) UnsetRecoverRdsParams() {
	o.RecoverRdsParams.Unset()
}

// GetRecoverAuroraParams returns the RecoverAuroraParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverAwsParams) GetRecoverAuroraParams() RecoverAwsAuroraParams {
	if o == nil || o.RecoverAuroraParams.Get() == nil {
		var ret RecoverAwsAuroraParams
		return ret
	}
	return *o.RecoverAuroraParams.Get()
}

// GetRecoverAuroraParamsOk returns a tuple with the RecoverAuroraParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAwsParams) GetRecoverAuroraParamsOk() (*RecoverAwsAuroraParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverAuroraParams.Get(), o.RecoverAuroraParams.IsSet()
}

// HasRecoverAuroraParams returns a boolean if a field has been set.
func (o *RecoverAwsParams) HasRecoverAuroraParams() bool {
	if o != nil && o.RecoverAuroraParams.IsSet() {
		return true
	}

	return false
}

// SetRecoverAuroraParams gets a reference to the given NullableRecoverAwsAuroraParams and assigns it to the RecoverAuroraParams field.
func (o *RecoverAwsParams) SetRecoverAuroraParams(v RecoverAwsAuroraParams) {
	o.RecoverAuroraParams.Set(&v)
}
// SetRecoverAuroraParamsNil sets the value for RecoverAuroraParams to be an explicit nil
func (o *RecoverAwsParams) SetRecoverAuroraParamsNil() {
	o.RecoverAuroraParams.Set(nil)
}

// UnsetRecoverAuroraParams ensures that no value is present for RecoverAuroraParams, not even an explicit nil
func (o *RecoverAwsParams) UnsetRecoverAuroraParams() {
	o.RecoverAuroraParams.Unset()
}

// GetDownloadFileAndFolderParams returns the DownloadFileAndFolderParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverAwsParams) GetDownloadFileAndFolderParams() CommonDownloadFileAndFolderParams {
	if o == nil || o.DownloadFileAndFolderParams.Get() == nil {
		var ret CommonDownloadFileAndFolderParams
		return ret
	}
	return *o.DownloadFileAndFolderParams.Get()
}

// GetDownloadFileAndFolderParamsOk returns a tuple with the DownloadFileAndFolderParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAwsParams) GetDownloadFileAndFolderParamsOk() (*CommonDownloadFileAndFolderParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DownloadFileAndFolderParams.Get(), o.DownloadFileAndFolderParams.IsSet()
}

// HasDownloadFileAndFolderParams returns a boolean if a field has been set.
func (o *RecoverAwsParams) HasDownloadFileAndFolderParams() bool {
	if o != nil && o.DownloadFileAndFolderParams.IsSet() {
		return true
	}

	return false
}

// SetDownloadFileAndFolderParams gets a reference to the given NullableCommonDownloadFileAndFolderParams and assigns it to the DownloadFileAndFolderParams field.
func (o *RecoverAwsParams) SetDownloadFileAndFolderParams(v CommonDownloadFileAndFolderParams) {
	o.DownloadFileAndFolderParams.Set(&v)
}
// SetDownloadFileAndFolderParamsNil sets the value for DownloadFileAndFolderParams to be an explicit nil
func (o *RecoverAwsParams) SetDownloadFileAndFolderParamsNil() {
	o.DownloadFileAndFolderParams.Set(nil)
}

// UnsetDownloadFileAndFolderParams ensures that no value is present for DownloadFileAndFolderParams, not even an explicit nil
func (o *RecoverAwsParams) UnsetDownloadFileAndFolderParams() {
	o.DownloadFileAndFolderParams.Unset()
}

func (o RecoverAwsParams) MarshalJSON() ([]byte, error) {
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
	if o.RecoverFileAndFolderParams.IsSet() {
		toSerialize["recoverFileAndFolderParams"] = o.RecoverFileAndFolderParams.Get()
	}
	if o.RecoverRdsParams.IsSet() {
		toSerialize["recoverRdsParams"] = o.RecoverRdsParams.Get()
	}
	if o.RecoverAuroraParams.IsSet() {
		toSerialize["recoverAuroraParams"] = o.RecoverAuroraParams.Get()
	}
	if o.DownloadFileAndFolderParams.IsSet() {
		toSerialize["downloadFileAndFolderParams"] = o.DownloadFileAndFolderParams.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverAwsParams struct {
	value *RecoverAwsParams
	isSet bool
}

func (v NullableRecoverAwsParams) Get() *RecoverAwsParams {
	return v.value
}

func (v *NullableRecoverAwsParams) Set(val *RecoverAwsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverAwsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverAwsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverAwsParams(val *RecoverAwsParams) *NullableRecoverAwsParams {
	return &NullableRecoverAwsParams{value: val, isSet: true}
}

func (v NullableRecoverAwsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverAwsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RecoverAwsParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}