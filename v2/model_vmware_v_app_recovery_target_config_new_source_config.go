/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the VmwareVAppRecoveryTargetConfigNewSourceConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VmwareVAppRecoveryTargetConfigNewSourceConfig{}

// VmwareVAppRecoveryTargetConfigNewSourceConfig Specifies the new destination Source configuration parameters where the VMs will be recovered. This is mandatory if recoverToNewSource is set to true.
type VmwareVAppRecoveryTargetConfigNewSourceConfig struct {
	// Specifies the type of VMware source to which the VMs are being restored.
	SourceType NullableString `json:"sourceType"`
	VCloudDirectorParams *RecoverVmwareVAppVCDSourceConfig `json:"vCloudDirectorParams,omitempty"`
}

type _VmwareVAppRecoveryTargetConfigNewSourceConfig VmwareVAppRecoveryTargetConfigNewSourceConfig

// NewVmwareVAppRecoveryTargetConfigNewSourceConfig instantiates a new VmwareVAppRecoveryTargetConfigNewSourceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareVAppRecoveryTargetConfigNewSourceConfig(sourceType NullableString) *VmwareVAppRecoveryTargetConfigNewSourceConfig {
	this := VmwareVAppRecoveryTargetConfigNewSourceConfig{}
	this.SourceType = sourceType
	return &this
}

// NewVmwareVAppRecoveryTargetConfigNewSourceConfigWithDefaults instantiates a new VmwareVAppRecoveryTargetConfigNewSourceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareVAppRecoveryTargetConfigNewSourceConfigWithDefaults() *VmwareVAppRecoveryTargetConfigNewSourceConfig {
	this := VmwareVAppRecoveryTargetConfigNewSourceConfig{}
	return &this
}

// GetSourceType returns the SourceType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VmwareVAppRecoveryTargetConfigNewSourceConfig) GetSourceType() string {
	if o == nil || o.SourceType.Get() == nil {
		var ret string
		return ret
	}

	return *o.SourceType.Get()
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareVAppRecoveryTargetConfigNewSourceConfig) GetSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceType.Get(), o.SourceType.IsSet()
}

// SetSourceType sets field value
func (o *VmwareVAppRecoveryTargetConfigNewSourceConfig) SetSourceType(v string) {
	o.SourceType.Set(&v)
}

// GetVCloudDirectorParams returns the VCloudDirectorParams field value if set, zero value otherwise.
func (o *VmwareVAppRecoveryTargetConfigNewSourceConfig) GetVCloudDirectorParams() RecoverVmwareVAppVCDSourceConfig {
	if o == nil || IsNil(o.VCloudDirectorParams) {
		var ret RecoverVmwareVAppVCDSourceConfig
		return ret
	}
	return *o.VCloudDirectorParams
}

// GetVCloudDirectorParamsOk returns a tuple with the VCloudDirectorParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareVAppRecoveryTargetConfigNewSourceConfig) GetVCloudDirectorParamsOk() (*RecoverVmwareVAppVCDSourceConfig, bool) {
	if o == nil || IsNil(o.VCloudDirectorParams) {
		return nil, false
	}
	return o.VCloudDirectorParams, true
}

// HasVCloudDirectorParams returns a boolean if a field has been set.
func (o *VmwareVAppRecoveryTargetConfigNewSourceConfig) HasVCloudDirectorParams() bool {
	if o != nil && !IsNil(o.VCloudDirectorParams) {
		return true
	}

	return false
}

// SetVCloudDirectorParams gets a reference to the given RecoverVmwareVAppVCDSourceConfig and assigns it to the VCloudDirectorParams field.
func (o *VmwareVAppRecoveryTargetConfigNewSourceConfig) SetVCloudDirectorParams(v RecoverVmwareVAppVCDSourceConfig) {
	o.VCloudDirectorParams = &v
}

func (o VmwareVAppRecoveryTargetConfigNewSourceConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VmwareVAppRecoveryTargetConfigNewSourceConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourceType"] = o.SourceType.Get()
	if !IsNil(o.VCloudDirectorParams) {
		toSerialize["vCloudDirectorParams"] = o.VCloudDirectorParams
	}
	return toSerialize, nil
}

func (o *VmwareVAppRecoveryTargetConfigNewSourceConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sourceType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varVmwareVAppRecoveryTargetConfigNewSourceConfig := _VmwareVAppRecoveryTargetConfigNewSourceConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVmwareVAppRecoveryTargetConfigNewSourceConfig)

	if err != nil {
		return err
	}

	*o = VmwareVAppRecoveryTargetConfigNewSourceConfig(varVmwareVAppRecoveryTargetConfigNewSourceConfig)

	return err
}

type NullableVmwareVAppRecoveryTargetConfigNewSourceConfig struct {
	value *VmwareVAppRecoveryTargetConfigNewSourceConfig
	isSet bool
}

func (v NullableVmwareVAppRecoveryTargetConfigNewSourceConfig) Get() *VmwareVAppRecoveryTargetConfigNewSourceConfig {
	return v.value
}

func (v *NullableVmwareVAppRecoveryTargetConfigNewSourceConfig) Set(val *VmwareVAppRecoveryTargetConfigNewSourceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareVAppRecoveryTargetConfigNewSourceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareVAppRecoveryTargetConfigNewSourceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareVAppRecoveryTargetConfigNewSourceConfig(val *VmwareVAppRecoveryTargetConfigNewSourceConfig) *NullableVmwareVAppRecoveryTargetConfigNewSourceConfig {
	return &NullableVmwareVAppRecoveryTargetConfigNewSourceConfig{value: val, isSet: true}
}

func (v NullableVmwareVAppRecoveryTargetConfigNewSourceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareVAppRecoveryTargetConfigNewSourceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


