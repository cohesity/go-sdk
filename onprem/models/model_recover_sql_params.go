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

// RecoverSqlParams Specifies the recovery options specific to Sql environment.
type RecoverSqlParams struct {
	// Specifies the type of recover action to be performed.
	RecoveryAction string `json:"recoveryAction"`
	// Specifies the parameters to recover Sql databases.
	RecoverAppParams []RecoverSqlAppParams `json:"recoverAppParams,omitempty"`
	// Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
	VlanConfig NullableRecoveryVlanConfig `json:"vlanConfig,omitempty"`
}

// NewRecoverSqlParams instantiates a new RecoverSqlParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverSqlParams(recoveryAction string) *RecoverSqlParams {
	this := RecoverSqlParams{}
	this.RecoveryAction = recoveryAction
	return &this
}

// NewRecoverSqlParamsWithDefaults instantiates a new RecoverSqlParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverSqlParamsWithDefaults() *RecoverSqlParams {
	this := RecoverSqlParams{}
	return &this
}

// GetRecoveryAction returns the RecoveryAction field value
func (o *RecoverSqlParams) GetRecoveryAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecoveryAction
}

// GetRecoveryActionOk returns a tuple with the RecoveryAction field value
// and a boolean to check if the value has been set.
func (o *RecoverSqlParams) GetRecoveryActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecoveryAction, true
}

// SetRecoveryAction sets field value
func (o *RecoverSqlParams) SetRecoveryAction(v string) {
	o.RecoveryAction = v
}

// GetRecoverAppParams returns the RecoverAppParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverSqlParams) GetRecoverAppParams() []RecoverSqlAppParams {
	if o == nil  {
		var ret []RecoverSqlAppParams
		return ret
	}
	return o.RecoverAppParams
}

// GetRecoverAppParamsOk returns a tuple with the RecoverAppParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverSqlParams) GetRecoverAppParamsOk() (*[]RecoverSqlAppParams, bool) {
	if o == nil || o.RecoverAppParams == nil {
		return nil, false
	}
	return &o.RecoverAppParams, true
}

// HasRecoverAppParams returns a boolean if a field has been set.
func (o *RecoverSqlParams) HasRecoverAppParams() bool {
	if o != nil && o.RecoverAppParams != nil {
		return true
	}

	return false
}

// SetRecoverAppParams gets a reference to the given []RecoverSqlAppParams and assigns it to the RecoverAppParams field.
func (o *RecoverSqlParams) SetRecoverAppParams(v []RecoverSqlAppParams) {
	o.RecoverAppParams = v
}

// GetVlanConfig returns the VlanConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverSqlParams) GetVlanConfig() RecoveryVlanConfig {
	if o == nil || o.VlanConfig.Get() == nil {
		var ret RecoveryVlanConfig
		return ret
	}
	return *o.VlanConfig.Get()
}

// GetVlanConfigOk returns a tuple with the VlanConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverSqlParams) GetVlanConfigOk() (*RecoveryVlanConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VlanConfig.Get(), o.VlanConfig.IsSet()
}

// HasVlanConfig returns a boolean if a field has been set.
func (o *RecoverSqlParams) HasVlanConfig() bool {
	if o != nil && o.VlanConfig.IsSet() {
		return true
	}

	return false
}

// SetVlanConfig gets a reference to the given NullableRecoveryVlanConfig and assigns it to the VlanConfig field.
func (o *RecoverSqlParams) SetVlanConfig(v RecoveryVlanConfig) {
	o.VlanConfig.Set(&v)
}
// SetVlanConfigNil sets the value for VlanConfig to be an explicit nil
func (o *RecoverSqlParams) SetVlanConfigNil() {
	o.VlanConfig.Set(nil)
}

// UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil
func (o *RecoverSqlParams) UnsetVlanConfig() {
	o.VlanConfig.Unset()
}

func (o RecoverSqlParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recoveryAction"] = o.RecoveryAction
	}
	if o.RecoverAppParams != nil {
		toSerialize["recoverAppParams"] = o.RecoverAppParams
	}
	if o.VlanConfig.IsSet() {
		toSerialize["vlanConfig"] = o.VlanConfig.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverSqlParams struct {
	value *RecoverSqlParams
	isSet bool
}

func (v NullableRecoverSqlParams) Get() *RecoverSqlParams {
	return v.value
}

func (v *NullableRecoverSqlParams) Set(val *RecoverSqlParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverSqlParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverSqlParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverSqlParams(val *RecoverSqlParams) *NullableRecoverSqlParams {
	return &NullableRecoverSqlParams{value: val, isSet: true}
}

func (v NullableRecoverSqlParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverSqlParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RecoverSqlParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}