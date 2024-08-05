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

// checks if the ObjectSnapshotAwsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectSnapshotAwsParams{}

// ObjectSnapshotAwsParams Specifies the parameters specific to AWS type snapshot.
type ObjectSnapshotAwsParams struct {
	// Specifies the protection type of AWS snapshots.
	ProtectionType NullableString `json:"protectionType,omitempty"`
}

// NewObjectSnapshotAwsParams instantiates a new ObjectSnapshotAwsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectSnapshotAwsParams() *ObjectSnapshotAwsParams {
	this := ObjectSnapshotAwsParams{}
	return &this
}

// NewObjectSnapshotAwsParamsWithDefaults instantiates a new ObjectSnapshotAwsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectSnapshotAwsParamsWithDefaults() *ObjectSnapshotAwsParams {
	this := ObjectSnapshotAwsParams{}
	return &this
}

// GetProtectionType returns the ProtectionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectSnapshotAwsParams) GetProtectionType() string {
	if o == nil || IsNil(o.ProtectionType.Get()) {
		var ret string
		return ret
	}
	return *o.ProtectionType.Get()
}

// GetProtectionTypeOk returns a tuple with the ProtectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectSnapshotAwsParams) GetProtectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProtectionType.Get(), o.ProtectionType.IsSet()
}

// HasProtectionType returns a boolean if a field has been set.
func (o *ObjectSnapshotAwsParams) HasProtectionType() bool {
	if o != nil && o.ProtectionType.IsSet() {
		return true
	}

	return false
}

// SetProtectionType gets a reference to the given NullableString and assigns it to the ProtectionType field.
func (o *ObjectSnapshotAwsParams) SetProtectionType(v string) {
	o.ProtectionType.Set(&v)
}
// SetProtectionTypeNil sets the value for ProtectionType to be an explicit nil
func (o *ObjectSnapshotAwsParams) SetProtectionTypeNil() {
	o.ProtectionType.Set(nil)
}

// UnsetProtectionType ensures that no value is present for ProtectionType, not even an explicit nil
func (o *ObjectSnapshotAwsParams) UnsetProtectionType() {
	o.ProtectionType.Unset()
}

func (o ObjectSnapshotAwsParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectSnapshotAwsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ProtectionType.IsSet() {
		toSerialize["protectionType"] = o.ProtectionType.Get()
	}
	return toSerialize, nil
}

type NullableObjectSnapshotAwsParams struct {
	value *ObjectSnapshotAwsParams
	isSet bool
}

func (v NullableObjectSnapshotAwsParams) Get() *ObjectSnapshotAwsParams {
	return v.value
}

func (v *NullableObjectSnapshotAwsParams) Set(val *ObjectSnapshotAwsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectSnapshotAwsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectSnapshotAwsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectSnapshotAwsParams(val *ObjectSnapshotAwsParams) *NullableObjectSnapshotAwsParams {
	return &NullableObjectSnapshotAwsParams{value: val, isSet: true}
}

func (v NullableObjectSnapshotAwsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectSnapshotAwsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


