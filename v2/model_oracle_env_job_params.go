/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the OracleEnvJobParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OracleEnvJobParams{}

// OracleEnvJobParams Specifies job parameters applicable for all 'kOracle' Environment type Protection Sources in a Protection Job.
type OracleEnvJobParams struct {
	// Specifies whether the mountpoints created while backing up Oracle DBs should be persisted. Note: This parameter is for the entire Job. For overriding persistence of mountpoints for a subset of Oracle hosts within the job, refer OracleSourceParams.
	PersistMountpoints NullableBool `json:"persistMountpoints,omitempty"`
	VlanParams *VlanParams `json:"vlanParams,omitempty"`
}

// NewOracleEnvJobParams instantiates a new OracleEnvJobParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOracleEnvJobParams() *OracleEnvJobParams {
	this := OracleEnvJobParams{}
	return &this
}

// NewOracleEnvJobParamsWithDefaults instantiates a new OracleEnvJobParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOracleEnvJobParamsWithDefaults() *OracleEnvJobParams {
	this := OracleEnvJobParams{}
	return &this
}

// GetPersistMountpoints returns the PersistMountpoints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OracleEnvJobParams) GetPersistMountpoints() bool {
	if o == nil || IsNil(o.PersistMountpoints.Get()) {
		var ret bool
		return ret
	}
	return *o.PersistMountpoints.Get()
}

// GetPersistMountpointsOk returns a tuple with the PersistMountpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OracleEnvJobParams) GetPersistMountpointsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PersistMountpoints.Get(), o.PersistMountpoints.IsSet()
}

// HasPersistMountpoints returns a boolean if a field has been set.
func (o *OracleEnvJobParams) HasPersistMountpoints() bool {
	if o != nil && o.PersistMountpoints.IsSet() {
		return true
	}

	return false
}

// SetPersistMountpoints gets a reference to the given NullableBool and assigns it to the PersistMountpoints field.
func (o *OracleEnvJobParams) SetPersistMountpoints(v bool) {
	o.PersistMountpoints.Set(&v)
}
// SetPersistMountpointsNil sets the value for PersistMountpoints to be an explicit nil
func (o *OracleEnvJobParams) SetPersistMountpointsNil() {
	o.PersistMountpoints.Set(nil)
}

// UnsetPersistMountpoints ensures that no value is present for PersistMountpoints, not even an explicit nil
func (o *OracleEnvJobParams) UnsetPersistMountpoints() {
	o.PersistMountpoints.Unset()
}

// GetVlanParams returns the VlanParams field value if set, zero value otherwise.
func (o *OracleEnvJobParams) GetVlanParams() VlanParams {
	if o == nil || IsNil(o.VlanParams) {
		var ret VlanParams
		return ret
	}
	return *o.VlanParams
}

// GetVlanParamsOk returns a tuple with the VlanParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleEnvJobParams) GetVlanParamsOk() (*VlanParams, bool) {
	if o == nil || IsNil(o.VlanParams) {
		return nil, false
	}
	return o.VlanParams, true
}

// HasVlanParams returns a boolean if a field has been set.
func (o *OracleEnvJobParams) HasVlanParams() bool {
	if o != nil && !IsNil(o.VlanParams) {
		return true
	}

	return false
}

// SetVlanParams gets a reference to the given VlanParams and assigns it to the VlanParams field.
func (o *OracleEnvJobParams) SetVlanParams(v VlanParams) {
	o.VlanParams = &v
}

func (o OracleEnvJobParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OracleEnvJobParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PersistMountpoints.IsSet() {
		toSerialize["persistMountpoints"] = o.PersistMountpoints.Get()
	}
	if !IsNil(o.VlanParams) {
		toSerialize["vlanParams"] = o.VlanParams
	}
	return toSerialize, nil
}

type NullableOracleEnvJobParams struct {
	value *OracleEnvJobParams
	isSet bool
}

func (v NullableOracleEnvJobParams) Get() *OracleEnvJobParams {
	return v.value
}

func (v *NullableOracleEnvJobParams) Set(val *OracleEnvJobParams) {
	v.value = val
	v.isSet = true
}

func (v NullableOracleEnvJobParams) IsSet() bool {
	return v.isSet
}

func (v *NullableOracleEnvJobParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOracleEnvJobParams(val *OracleEnvJobParams) *NullableOracleEnvJobParams {
	return &NullableOracleEnvJobParams{value: val, isSet: true}
}

func (v NullableOracleEnvJobParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOracleEnvJobParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


