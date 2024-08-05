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

// checks if the HeliosRemoteTargetConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeliosRemoteTargetConfig{}

// HeliosRemoteTargetConfig Specifies the configuration for adding remote cluster as repilcation target
type HeliosRemoteTargetConfig struct {
	// Specifies the cluster id of the target replication cluster.
	ClusterId NullableInt64 `json:"clusterId"`
	// Specifies the cluster name of the target replication cluster.
	ClusterName NullableString `json:"clusterName,omitempty"`
}

type _HeliosRemoteTargetConfig HeliosRemoteTargetConfig

// NewHeliosRemoteTargetConfig instantiates a new HeliosRemoteTargetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeliosRemoteTargetConfig(clusterId NullableInt64) *HeliosRemoteTargetConfig {
	this := HeliosRemoteTargetConfig{}
	this.ClusterId = clusterId
	return &this
}

// NewHeliosRemoteTargetConfigWithDefaults instantiates a new HeliosRemoteTargetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeliosRemoteTargetConfigWithDefaults() *HeliosRemoteTargetConfig {
	this := HeliosRemoteTargetConfig{}
	return &this
}

// GetClusterId returns the ClusterId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *HeliosRemoteTargetConfig) GetClusterId() int64 {
	if o == nil || o.ClusterId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.ClusterId.Get()
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosRemoteTargetConfig) GetClusterIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterId.Get(), o.ClusterId.IsSet()
}

// SetClusterId sets field value
func (o *HeliosRemoteTargetConfig) SetClusterId(v int64) {
	o.ClusterId.Set(&v)
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosRemoteTargetConfig) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName.Get()) {
		var ret string
		return ret
	}
	return *o.ClusterName.Get()
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosRemoteTargetConfig) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterName.Get(), o.ClusterName.IsSet()
}

// HasClusterName returns a boolean if a field has been set.
func (o *HeliosRemoteTargetConfig) HasClusterName() bool {
	if o != nil && o.ClusterName.IsSet() {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given NullableString and assigns it to the ClusterName field.
func (o *HeliosRemoteTargetConfig) SetClusterName(v string) {
	o.ClusterName.Set(&v)
}
// SetClusterNameNil sets the value for ClusterName to be an explicit nil
func (o *HeliosRemoteTargetConfig) SetClusterNameNil() {
	o.ClusterName.Set(nil)
}

// UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
func (o *HeliosRemoteTargetConfig) UnsetClusterName() {
	o.ClusterName.Unset()
}

func (o HeliosRemoteTargetConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HeliosRemoteTargetConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clusterId"] = o.ClusterId.Get()
	if o.ClusterName.IsSet() {
		toSerialize["clusterName"] = o.ClusterName.Get()
	}
	return toSerialize, nil
}

func (o *HeliosRemoteTargetConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clusterId",
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

	varHeliosRemoteTargetConfig := _HeliosRemoteTargetConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHeliosRemoteTargetConfig)

	if err != nil {
		return err
	}

	*o = HeliosRemoteTargetConfig(varHeliosRemoteTargetConfig)

	return err
}

type NullableHeliosRemoteTargetConfig struct {
	value *HeliosRemoteTargetConfig
	isSet bool
}

func (v NullableHeliosRemoteTargetConfig) Get() *HeliosRemoteTargetConfig {
	return v.value
}

func (v *NullableHeliosRemoteTargetConfig) Set(val *HeliosRemoteTargetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableHeliosRemoteTargetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableHeliosRemoteTargetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeliosRemoteTargetConfig(val *HeliosRemoteTargetConfig) *NullableHeliosRemoteTargetConfig {
	return &NullableHeliosRemoteTargetConfig{value: val, isSet: true}
}

func (v NullableHeliosRemoteTargetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeliosRemoteTargetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


