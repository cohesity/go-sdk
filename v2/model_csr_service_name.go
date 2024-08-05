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

// checks if the CsrServiceName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CsrServiceName{}

// CsrServiceName Csr Service Name
type CsrServiceName struct {
	// Specifies the csr service name.
	CsrServiceName *string `json:"csrServiceName,omitempty"`
}

// NewCsrServiceName instantiates a new CsrServiceName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCsrServiceName() *CsrServiceName {
	this := CsrServiceName{}
	return &this
}

// NewCsrServiceNameWithDefaults instantiates a new CsrServiceName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCsrServiceNameWithDefaults() *CsrServiceName {
	this := CsrServiceName{}
	return &this
}

// GetCsrServiceName returns the CsrServiceName field value if set, zero value otherwise.
func (o *CsrServiceName) GetCsrServiceName() string {
	if o == nil || IsNil(o.CsrServiceName) {
		var ret string
		return ret
	}
	return *o.CsrServiceName
}

// GetCsrServiceNameOk returns a tuple with the CsrServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsrServiceName) GetCsrServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.CsrServiceName) {
		return nil, false
	}
	return o.CsrServiceName, true
}

// HasCsrServiceName returns a boolean if a field has been set.
func (o *CsrServiceName) HasCsrServiceName() bool {
	if o != nil && !IsNil(o.CsrServiceName) {
		return true
	}

	return false
}

// SetCsrServiceName gets a reference to the given string and assigns it to the CsrServiceName field.
func (o *CsrServiceName) SetCsrServiceName(v string) {
	o.CsrServiceName = &v
}

func (o CsrServiceName) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CsrServiceName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CsrServiceName) {
		toSerialize["csrServiceName"] = o.CsrServiceName
	}
	return toSerialize, nil
}

type NullableCsrServiceName struct {
	value *CsrServiceName
	isSet bool
}

func (v NullableCsrServiceName) Get() *CsrServiceName {
	return v.value
}

func (v *NullableCsrServiceName) Set(val *CsrServiceName) {
	v.value = val
	v.isSet = true
}

func (v NullableCsrServiceName) IsSet() bool {
	return v.isSet
}

func (v *NullableCsrServiceName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCsrServiceName(val *CsrServiceName) *NullableCsrServiceName {
	return &NullableCsrServiceName{value: val, isSet: true}
}

func (v NullableCsrServiceName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCsrServiceName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


