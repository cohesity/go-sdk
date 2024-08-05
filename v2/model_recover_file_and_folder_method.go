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

// checks if the RecoverFileAndFolderMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverFileAndFolderMethod{}

// RecoverFileAndFolderMethod Recover File and Folder Method
type RecoverFileAndFolderMethod struct {
	// Recover File and Folder Method.
	RecoverFileAndFolderMethod *string `json:"recoverFileAndFolderMethod,omitempty"`
}

// NewRecoverFileAndFolderMethod instantiates a new RecoverFileAndFolderMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverFileAndFolderMethod() *RecoverFileAndFolderMethod {
	this := RecoverFileAndFolderMethod{}
	return &this
}

// NewRecoverFileAndFolderMethodWithDefaults instantiates a new RecoverFileAndFolderMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverFileAndFolderMethodWithDefaults() *RecoverFileAndFolderMethod {
	this := RecoverFileAndFolderMethod{}
	return &this
}

// GetRecoverFileAndFolderMethod returns the RecoverFileAndFolderMethod field value if set, zero value otherwise.
func (o *RecoverFileAndFolderMethod) GetRecoverFileAndFolderMethod() string {
	if o == nil || IsNil(o.RecoverFileAndFolderMethod) {
		var ret string
		return ret
	}
	return *o.RecoverFileAndFolderMethod
}

// GetRecoverFileAndFolderMethodOk returns a tuple with the RecoverFileAndFolderMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverFileAndFolderMethod) GetRecoverFileAndFolderMethodOk() (*string, bool) {
	if o == nil || IsNil(o.RecoverFileAndFolderMethod) {
		return nil, false
	}
	return o.RecoverFileAndFolderMethod, true
}

// HasRecoverFileAndFolderMethod returns a boolean if a field has been set.
func (o *RecoverFileAndFolderMethod) HasRecoverFileAndFolderMethod() bool {
	if o != nil && !IsNil(o.RecoverFileAndFolderMethod) {
		return true
	}

	return false
}

// SetRecoverFileAndFolderMethod gets a reference to the given string and assigns it to the RecoverFileAndFolderMethod field.
func (o *RecoverFileAndFolderMethod) SetRecoverFileAndFolderMethod(v string) {
	o.RecoverFileAndFolderMethod = &v
}

func (o RecoverFileAndFolderMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverFileAndFolderMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecoverFileAndFolderMethod) {
		toSerialize["recoverFileAndFolderMethod"] = o.RecoverFileAndFolderMethod
	}
	return toSerialize, nil
}

type NullableRecoverFileAndFolderMethod struct {
	value *RecoverFileAndFolderMethod
	isSet bool
}

func (v NullableRecoverFileAndFolderMethod) Get() *RecoverFileAndFolderMethod {
	return v.value
}

func (v *NullableRecoverFileAndFolderMethod) Set(val *RecoverFileAndFolderMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverFileAndFolderMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverFileAndFolderMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverFileAndFolderMethod(val *RecoverFileAndFolderMethod) *NullableRecoverFileAndFolderMethod {
	return &NullableRecoverFileAndFolderMethod{value: val, isSet: true}
}

func (v NullableRecoverFileAndFolderMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverFileAndFolderMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


