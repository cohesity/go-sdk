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

// checks if the RecoverAwsRdsNewSourceConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverAwsRdsNewSourceConfig{}

// RecoverAwsRdsNewSourceConfig Specifies the new destination Source configuration where the RDS instances will be recovered.
type RecoverAwsRdsNewSourceConfig struct {
	NetworkConfig NullableRecoverAwsRdsNewSourceConfigNetworkConfig `json:"networkConfig"`
	Region NullableRecoverAwsRdsNewSourceConfigRegion `json:"region"`
	Source NullableRecoverAwsRdsNewSourceConfigSource `json:"source"`
}

type _RecoverAwsRdsNewSourceConfig RecoverAwsRdsNewSourceConfig

// NewRecoverAwsRdsNewSourceConfig instantiates a new RecoverAwsRdsNewSourceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverAwsRdsNewSourceConfig(networkConfig NullableRecoverAwsRdsNewSourceConfigNetworkConfig, region NullableRecoverAwsRdsNewSourceConfigRegion, source NullableRecoverAwsRdsNewSourceConfigSource) *RecoverAwsRdsNewSourceConfig {
	this := RecoverAwsRdsNewSourceConfig{}
	this.NetworkConfig = networkConfig
	this.Region = region
	this.Source = source
	return &this
}

// NewRecoverAwsRdsNewSourceConfigWithDefaults instantiates a new RecoverAwsRdsNewSourceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverAwsRdsNewSourceConfigWithDefaults() *RecoverAwsRdsNewSourceConfig {
	this := RecoverAwsRdsNewSourceConfig{}
	return &this
}

// GetNetworkConfig returns the NetworkConfig field value
// If the value is explicit nil, the zero value for RecoverAwsRdsNewSourceConfigNetworkConfig will be returned
func (o *RecoverAwsRdsNewSourceConfig) GetNetworkConfig() RecoverAwsRdsNewSourceConfigNetworkConfig {
	if o == nil || o.NetworkConfig.Get() == nil {
		var ret RecoverAwsRdsNewSourceConfigNetworkConfig
		return ret
	}

	return *o.NetworkConfig.Get()
}

// GetNetworkConfigOk returns a tuple with the NetworkConfig field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAwsRdsNewSourceConfig) GetNetworkConfigOk() (*RecoverAwsRdsNewSourceConfigNetworkConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkConfig.Get(), o.NetworkConfig.IsSet()
}

// SetNetworkConfig sets field value
func (o *RecoverAwsRdsNewSourceConfig) SetNetworkConfig(v RecoverAwsRdsNewSourceConfigNetworkConfig) {
	o.NetworkConfig.Set(&v)
}

// GetRegion returns the Region field value
// If the value is explicit nil, the zero value for RecoverAwsRdsNewSourceConfigRegion will be returned
func (o *RecoverAwsRdsNewSourceConfig) GetRegion() RecoverAwsRdsNewSourceConfigRegion {
	if o == nil || o.Region.Get() == nil {
		var ret RecoverAwsRdsNewSourceConfigRegion
		return ret
	}

	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAwsRdsNewSourceConfig) GetRegionOk() (*RecoverAwsRdsNewSourceConfigRegion, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// SetRegion sets field value
func (o *RecoverAwsRdsNewSourceConfig) SetRegion(v RecoverAwsRdsNewSourceConfigRegion) {
	o.Region.Set(&v)
}

// GetSource returns the Source field value
// If the value is explicit nil, the zero value for RecoverAwsRdsNewSourceConfigSource will be returned
func (o *RecoverAwsRdsNewSourceConfig) GetSource() RecoverAwsRdsNewSourceConfigSource {
	if o == nil || o.Source.Get() == nil {
		var ret RecoverAwsRdsNewSourceConfigSource
		return ret
	}

	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverAwsRdsNewSourceConfig) GetSourceOk() (*RecoverAwsRdsNewSourceConfigSource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// SetSource sets field value
func (o *RecoverAwsRdsNewSourceConfig) SetSource(v RecoverAwsRdsNewSourceConfigSource) {
	o.Source.Set(&v)
}

func (o RecoverAwsRdsNewSourceConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverAwsRdsNewSourceConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["networkConfig"] = o.NetworkConfig.Get()
	toSerialize["region"] = o.Region.Get()
	toSerialize["source"] = o.Source.Get()
	return toSerialize, nil
}

func (o *RecoverAwsRdsNewSourceConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"networkConfig",
		"region",
		"source",
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

	varRecoverAwsRdsNewSourceConfig := _RecoverAwsRdsNewSourceConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverAwsRdsNewSourceConfig)

	if err != nil {
		return err
	}

	*o = RecoverAwsRdsNewSourceConfig(varRecoverAwsRdsNewSourceConfig)

	return err
}

type NullableRecoverAwsRdsNewSourceConfig struct {
	value *RecoverAwsRdsNewSourceConfig
	isSet bool
}

func (v NullableRecoverAwsRdsNewSourceConfig) Get() *RecoverAwsRdsNewSourceConfig {
	return v.value
}

func (v *NullableRecoverAwsRdsNewSourceConfig) Set(val *RecoverAwsRdsNewSourceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverAwsRdsNewSourceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverAwsRdsNewSourceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverAwsRdsNewSourceConfig(val *RecoverAwsRdsNewSourceConfig) *NullableRecoverAwsRdsNewSourceConfig {
	return &NullableRecoverAwsRdsNewSourceConfig{value: val, isSet: true}
}

func (v NullableRecoverAwsRdsNewSourceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverAwsRdsNewSourceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


