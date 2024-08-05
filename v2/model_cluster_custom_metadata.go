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

// checks if the ClusterCustomMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterCustomMetadata{}

// ClusterCustomMetadata Specifies the list of custom properties associated with the cluster. API callers can choose to set the following properties using provided key and value fields. The input values must always be in the string format and each key must be unique. The callers should ensure that no sensitive information such as passwords is sent in the custom metadata.
type ClusterCustomMetadata struct {
	// Specifies the key for the customer cluster metadata.
	Key NullableString `json:"key,omitempty"`
	// Specifies the value for the above key.
	Value NullableString `json:"value,omitempty"`
}

// NewClusterCustomMetadata instantiates a new ClusterCustomMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterCustomMetadata() *ClusterCustomMetadata {
	this := ClusterCustomMetadata{}
	return &this
}

// NewClusterCustomMetadataWithDefaults instantiates a new ClusterCustomMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterCustomMetadataWithDefaults() *ClusterCustomMetadata {
	this := ClusterCustomMetadata{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterCustomMetadata) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterCustomMetadata) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *ClusterCustomMetadata) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *ClusterCustomMetadata) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *ClusterCustomMetadata) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *ClusterCustomMetadata) UnsetKey() {
	o.Key.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterCustomMetadata) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterCustomMetadata) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ClusterCustomMetadata) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *ClusterCustomMetadata) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *ClusterCustomMetadata) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ClusterCustomMetadata) UnsetValue() {
	o.Value.Unset()
}

func (o ClusterCustomMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterCustomMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableClusterCustomMetadata struct {
	value *ClusterCustomMetadata
	isSet bool
}

func (v NullableClusterCustomMetadata) Get() *ClusterCustomMetadata {
	return v.value
}

func (v *NullableClusterCustomMetadata) Set(val *ClusterCustomMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterCustomMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterCustomMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterCustomMetadata(val *ClusterCustomMetadata) *NullableClusterCustomMetadata {
	return &NullableClusterCustomMetadata{value: val, isSet: true}
}

func (v NullableClusterCustomMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterCustomMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


