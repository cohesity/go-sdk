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

// RecoverOracleAppParams Specifies the parameters to recover Oracle databases.
type RecoverOracleAppParams struct {
	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	TargetEnvironment string `json:"targetEnvironment"`
	// Specifies the params for recovering to a oracle host. Provided oracle backup should be recovered to same type of target host. For Example: If you have oracle backup taken from a physical host then that should be recovered to physical host only.
	OracleTargetParams NullableOracleTargetParamsForRecoverOracleApp `json:"oracleTargetParams,omitempty"`
	// Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
	VlanConfig NullableRecoveryVlanConfig `json:"vlanConfig,omitempty"`
}

// NewRecoverOracleAppParams instantiates a new RecoverOracleAppParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverOracleAppParams(targetEnvironment string) *RecoverOracleAppParams {
	this := RecoverOracleAppParams{}
	this.TargetEnvironment = targetEnvironment
	return &this
}

// NewRecoverOracleAppParamsWithDefaults instantiates a new RecoverOracleAppParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverOracleAppParamsWithDefaults() *RecoverOracleAppParams {
	this := RecoverOracleAppParams{}
	return &this
}

// GetTargetEnvironment returns the TargetEnvironment field value
func (o *RecoverOracleAppParams) GetTargetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetEnvironment
}

// GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field value
// and a boolean to check if the value has been set.
func (o *RecoverOracleAppParams) GetTargetEnvironmentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetEnvironment, true
}

// SetTargetEnvironment sets field value
func (o *RecoverOracleAppParams) SetTargetEnvironment(v string) {
	o.TargetEnvironment = v
}

// GetOracleTargetParams returns the OracleTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverOracleAppParams) GetOracleTargetParams() OracleTargetParamsForRecoverOracleApp {
	if o == nil || o.OracleTargetParams.Get() == nil {
		var ret OracleTargetParamsForRecoverOracleApp
		return ret
	}
	return *o.OracleTargetParams.Get()
}

// GetOracleTargetParamsOk returns a tuple with the OracleTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverOracleAppParams) GetOracleTargetParamsOk() (*OracleTargetParamsForRecoverOracleApp, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OracleTargetParams.Get(), o.OracleTargetParams.IsSet()
}

// HasOracleTargetParams returns a boolean if a field has been set.
func (o *RecoverOracleAppParams) HasOracleTargetParams() bool {
	if o != nil && o.OracleTargetParams.IsSet() {
		return true
	}

	return false
}

// SetOracleTargetParams gets a reference to the given NullableOracleTargetParamsForRecoverOracleApp and assigns it to the OracleTargetParams field.
func (o *RecoverOracleAppParams) SetOracleTargetParams(v OracleTargetParamsForRecoverOracleApp) {
	o.OracleTargetParams.Set(&v)
}
// SetOracleTargetParamsNil sets the value for OracleTargetParams to be an explicit nil
func (o *RecoverOracleAppParams) SetOracleTargetParamsNil() {
	o.OracleTargetParams.Set(nil)
}

// UnsetOracleTargetParams ensures that no value is present for OracleTargetParams, not even an explicit nil
func (o *RecoverOracleAppParams) UnsetOracleTargetParams() {
	o.OracleTargetParams.Unset()
}

// GetVlanConfig returns the VlanConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverOracleAppParams) GetVlanConfig() RecoveryVlanConfig {
	if o == nil || o.VlanConfig.Get() == nil {
		var ret RecoveryVlanConfig
		return ret
	}
	return *o.VlanConfig.Get()
}

// GetVlanConfigOk returns a tuple with the VlanConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverOracleAppParams) GetVlanConfigOk() (*RecoveryVlanConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VlanConfig.Get(), o.VlanConfig.IsSet()
}

// HasVlanConfig returns a boolean if a field has been set.
func (o *RecoverOracleAppParams) HasVlanConfig() bool {
	if o != nil && o.VlanConfig.IsSet() {
		return true
	}

	return false
}

// SetVlanConfig gets a reference to the given NullableRecoveryVlanConfig and assigns it to the VlanConfig field.
func (o *RecoverOracleAppParams) SetVlanConfig(v RecoveryVlanConfig) {
	o.VlanConfig.Set(&v)
}
// SetVlanConfigNil sets the value for VlanConfig to be an explicit nil
func (o *RecoverOracleAppParams) SetVlanConfigNil() {
	o.VlanConfig.Set(nil)
}

// UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil
func (o *RecoverOracleAppParams) UnsetVlanConfig() {
	o.VlanConfig.Unset()
}

func (o RecoverOracleAppParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["targetEnvironment"] = o.TargetEnvironment
	}
	if o.OracleTargetParams.IsSet() {
		toSerialize["oracleTargetParams"] = o.OracleTargetParams.Get()
	}
	if o.VlanConfig.IsSet() {
		toSerialize["vlanConfig"] = o.VlanConfig.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverOracleAppParams struct {
	value *RecoverOracleAppParams
	isSet bool
}

func (v NullableRecoverOracleAppParams) Get() *RecoverOracleAppParams {
	return v.value
}

func (v *NullableRecoverOracleAppParams) Set(val *RecoverOracleAppParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverOracleAppParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverOracleAppParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverOracleAppParams(val *RecoverOracleAppParams) *NullableRecoverOracleAppParams {
	return &NullableRecoverOracleAppParams{value: val, isSet: true}
}

func (v NullableRecoverOracleAppParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverOracleAppParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RecoverOracleAppParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}