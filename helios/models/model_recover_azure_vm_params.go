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

// RecoverAzureVmParams Specifies the parameters to recover Azure VM.
type RecoverAzureVmParams struct {
	// Specifies the Protection Group Runs params to recover. All the VM's that are successfully backed up by specified Runs will be recovered. This can be specified along with individual snapshots of VMs. User has to make sure that specified Object snapshots and Protection Group Runs should not have any intersection. For example, user cannot specify multiple Runs which has same Object or an Object snapshot and a Run which has same Object's snapshot.
	RecoverProtectionGroupRunsParams []RecoverProtectionGroupRunParams `json:"recoverProtectionGroupRunsParams,omitempty"`
	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	TargetEnvironment string `json:"targetEnvironment"`
	// Specifies the params for recovering to an Azure target.
	AzureTargetParams NullableAzureTargetParamsForRecoverVm `json:"azureTargetParams,omitempty"`
}

// NewRecoverAzureVmParams instantiates a new RecoverAzureVmParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverAzureVmParams(targetEnvironment string) *RecoverAzureVmParams {
	this := RecoverAzureVmParams{}
	this.TargetEnvironment = targetEnvironment
	return &this
}

// NewRecoverAzureVmParamsWithDefaults instantiates a new RecoverAzureVmParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverAzureVmParamsWithDefaults() *RecoverAzureVmParams {
	this := RecoverAzureVmParams{}
	return &this
}

// GetRecoverProtectionGroupRunsParams returns the RecoverProtectionGroupRunsParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverAzureVmParams) GetRecoverProtectionGroupRunsParams() []RecoverProtectionGroupRunParams {
	if o == nil  {
		var ret []RecoverProtectionGroupRunParams
		return ret
	}
	return o.RecoverProtectionGroupRunsParams
}

// GetRecoverProtectionGroupRunsParamsOk returns a tuple with the RecoverProtectionGroupRunsParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAzureVmParams) GetRecoverProtectionGroupRunsParamsOk() (*[]RecoverProtectionGroupRunParams, bool) {
	if o == nil || o.RecoverProtectionGroupRunsParams == nil {
		return nil, false
	}
	return &o.RecoverProtectionGroupRunsParams, true
}

// HasRecoverProtectionGroupRunsParams returns a boolean if a field has been set.
func (o *RecoverAzureVmParams) HasRecoverProtectionGroupRunsParams() bool {
	if o != nil && o.RecoverProtectionGroupRunsParams != nil {
		return true
	}

	return false
}

// SetRecoverProtectionGroupRunsParams gets a reference to the given []RecoverProtectionGroupRunParams and assigns it to the RecoverProtectionGroupRunsParams field.
func (o *RecoverAzureVmParams) SetRecoverProtectionGroupRunsParams(v []RecoverProtectionGroupRunParams) {
	o.RecoverProtectionGroupRunsParams = v
}

// GetTargetEnvironment returns the TargetEnvironment field value
func (o *RecoverAzureVmParams) GetTargetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetEnvironment
}

// GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field value
// and a boolean to check if the value has been set.
func (o *RecoverAzureVmParams) GetTargetEnvironmentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetEnvironment, true
}

// SetTargetEnvironment sets field value
func (o *RecoverAzureVmParams) SetTargetEnvironment(v string) {
	o.TargetEnvironment = v
}

// GetAzureTargetParams returns the AzureTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverAzureVmParams) GetAzureTargetParams() AzureTargetParamsForRecoverVm {
	if o == nil || o.AzureTargetParams.Get() == nil {
		var ret AzureTargetParamsForRecoverVm
		return ret
	}
	return *o.AzureTargetParams.Get()
}

// GetAzureTargetParamsOk returns a tuple with the AzureTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAzureVmParams) GetAzureTargetParamsOk() (*AzureTargetParamsForRecoverVm, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AzureTargetParams.Get(), o.AzureTargetParams.IsSet()
}

// HasAzureTargetParams returns a boolean if a field has been set.
func (o *RecoverAzureVmParams) HasAzureTargetParams() bool {
	if o != nil && o.AzureTargetParams.IsSet() {
		return true
	}

	return false
}

// SetAzureTargetParams gets a reference to the given NullableAzureTargetParamsForRecoverVm and assigns it to the AzureTargetParams field.
func (o *RecoverAzureVmParams) SetAzureTargetParams(v AzureTargetParamsForRecoverVm) {
	o.AzureTargetParams.Set(&v)
}
// SetAzureTargetParamsNil sets the value for AzureTargetParams to be an explicit nil
func (o *RecoverAzureVmParams) SetAzureTargetParamsNil() {
	o.AzureTargetParams.Set(nil)
}

// UnsetAzureTargetParams ensures that no value is present for AzureTargetParams, not even an explicit nil
func (o *RecoverAzureVmParams) UnsetAzureTargetParams() {
	o.AzureTargetParams.Unset()
}

func (o RecoverAzureVmParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecoverProtectionGroupRunsParams != nil {
		toSerialize["recoverProtectionGroupRunsParams"] = o.RecoverProtectionGroupRunsParams
	}
	if true {
		toSerialize["targetEnvironment"] = o.TargetEnvironment
	}
	if o.AzureTargetParams.IsSet() {
		toSerialize["azureTargetParams"] = o.AzureTargetParams.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverAzureVmParams struct {
	value *RecoverAzureVmParams
	isSet bool
}

func (v NullableRecoverAzureVmParams) Get() *RecoverAzureVmParams {
	return v.value
}

func (v *NullableRecoverAzureVmParams) Set(val *RecoverAzureVmParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverAzureVmParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverAzureVmParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverAzureVmParams(val *RecoverAzureVmParams) *NullableRecoverAzureVmParams {
	return &NullableRecoverAzureVmParams{value: val, isSet: true}
}

func (v NullableRecoverAzureVmParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverAzureVmParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


