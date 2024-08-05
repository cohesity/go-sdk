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

// checks if the ClusterOperationTypeAndId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterOperationTypeAndId{}

// ClusterOperationTypeAndId Operation type and id of the cluster operation.
type ClusterOperationTypeAndId struct {
	// Operation Id of cluster operation. 
	OperationId string `json:"operationId"`
	// Operation type of cluster operation created for the request. 
	OperationType string `json:"operationType"`
}

type _ClusterOperationTypeAndId ClusterOperationTypeAndId

// NewClusterOperationTypeAndId instantiates a new ClusterOperationTypeAndId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterOperationTypeAndId(operationId string, operationType string) *ClusterOperationTypeAndId {
	this := ClusterOperationTypeAndId{}
	this.OperationId = operationId
	this.OperationType = operationType
	return &this
}

// NewClusterOperationTypeAndIdWithDefaults instantiates a new ClusterOperationTypeAndId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterOperationTypeAndIdWithDefaults() *ClusterOperationTypeAndId {
	this := ClusterOperationTypeAndId{}
	return &this
}

// GetOperationId returns the OperationId field value
func (o *ClusterOperationTypeAndId) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *ClusterOperationTypeAndId) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *ClusterOperationTypeAndId) SetOperationId(v string) {
	o.OperationId = v
}

// GetOperationType returns the OperationType field value
func (o *ClusterOperationTypeAndId) GetOperationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value
// and a boolean to check if the value has been set.
func (o *ClusterOperationTypeAndId) GetOperationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationType, true
}

// SetOperationType sets field value
func (o *ClusterOperationTypeAndId) SetOperationType(v string) {
	o.OperationType = v
}

func (o ClusterOperationTypeAndId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterOperationTypeAndId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationId"] = o.OperationId
	toSerialize["operationType"] = o.OperationType
	return toSerialize, nil
}

func (o *ClusterOperationTypeAndId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"operationId",
		"operationType",
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

	varClusterOperationTypeAndId := _ClusterOperationTypeAndId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClusterOperationTypeAndId)

	if err != nil {
		return err
	}

	*o = ClusterOperationTypeAndId(varClusterOperationTypeAndId)

	return err
}

type NullableClusterOperationTypeAndId struct {
	value *ClusterOperationTypeAndId
	isSet bool
}

func (v NullableClusterOperationTypeAndId) Get() *ClusterOperationTypeAndId {
	return v.value
}

func (v *NullableClusterOperationTypeAndId) Set(val *ClusterOperationTypeAndId) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterOperationTypeAndId) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterOperationTypeAndId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterOperationTypeAndId(val *ClusterOperationTypeAndId) *NullableClusterOperationTypeAndId {
	return &NullableClusterOperationTypeAndId{value: val, isSet: true}
}

func (v NullableClusterOperationTypeAndId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterOperationTypeAndId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


