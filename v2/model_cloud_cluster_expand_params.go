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

// checks if the CloudClusterExpandParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudClusterExpandParams{}

// CloudClusterExpandParams Parameters to expand cloud edition cluster.
type CloudClusterExpandParams struct {
	// List of IPs of the new nodes.
	NodeIps []string `json:"nodeIps"`
}

type _CloudClusterExpandParams CloudClusterExpandParams

// NewCloudClusterExpandParams instantiates a new CloudClusterExpandParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudClusterExpandParams(nodeIps []string) *CloudClusterExpandParams {
	this := CloudClusterExpandParams{}
	this.NodeIps = nodeIps
	return &this
}

// NewCloudClusterExpandParamsWithDefaults instantiates a new CloudClusterExpandParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudClusterExpandParamsWithDefaults() *CloudClusterExpandParams {
	this := CloudClusterExpandParams{}
	return &this
}

// GetNodeIps returns the NodeIps field value
func (o *CloudClusterExpandParams) GetNodeIps() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.NodeIps
}

// GetNodeIpsOk returns a tuple with the NodeIps field value
// and a boolean to check if the value has been set.
func (o *CloudClusterExpandParams) GetNodeIpsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeIps, true
}

// SetNodeIps sets field value
func (o *CloudClusterExpandParams) SetNodeIps(v []string) {
	o.NodeIps = v
}

func (o CloudClusterExpandParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudClusterExpandParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nodeIps"] = o.NodeIps
	return toSerialize, nil
}

func (o *CloudClusterExpandParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nodeIps",
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

	varCloudClusterExpandParams := _CloudClusterExpandParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudClusterExpandParams)

	if err != nil {
		return err
	}

	*o = CloudClusterExpandParams(varCloudClusterExpandParams)

	return err
}

type NullableCloudClusterExpandParams struct {
	value *CloudClusterExpandParams
	isSet bool
}

func (v NullableCloudClusterExpandParams) Get() *CloudClusterExpandParams {
	return v.value
}

func (v *NullableCloudClusterExpandParams) Set(val *CloudClusterExpandParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudClusterExpandParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudClusterExpandParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudClusterExpandParams(val *CloudClusterExpandParams) *NullableCloudClusterExpandParams {
	return &NullableCloudClusterExpandParams{value: val, isSet: true}
}

func (v NullableCloudClusterExpandParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudClusterExpandParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


