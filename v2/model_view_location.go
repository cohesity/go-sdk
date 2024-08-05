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

// checks if the ViewLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViewLocation{}

// ViewLocation Contains the View info.
type ViewLocation struct {
	// Denotes the name of the view.
	ViewName *string `json:"viewName,omitempty"`
}

// NewViewLocation instantiates a new ViewLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewLocation() *ViewLocation {
	this := ViewLocation{}
	return &this
}

// NewViewLocationWithDefaults instantiates a new ViewLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewLocationWithDefaults() *ViewLocation {
	this := ViewLocation{}
	return &this
}

// GetViewName returns the ViewName field value if set, zero value otherwise.
func (o *ViewLocation) GetViewName() string {
	if o == nil || IsNil(o.ViewName) {
		var ret string
		return ret
	}
	return *o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewLocation) GetViewNameOk() (*string, bool) {
	if o == nil || IsNil(o.ViewName) {
		return nil, false
	}
	return o.ViewName, true
}

// HasViewName returns a boolean if a field has been set.
func (o *ViewLocation) HasViewName() bool {
	if o != nil && !IsNil(o.ViewName) {
		return true
	}

	return false
}

// SetViewName gets a reference to the given string and assigns it to the ViewName field.
func (o *ViewLocation) SetViewName(v string) {
	o.ViewName = &v
}

func (o ViewLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ViewLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ViewName) {
		toSerialize["viewName"] = o.ViewName
	}
	return toSerialize, nil
}

type NullableViewLocation struct {
	value *ViewLocation
	isSet bool
}

func (v NullableViewLocation) Get() *ViewLocation {
	return v.value
}

func (v *NullableViewLocation) Set(val *ViewLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableViewLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableViewLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewLocation(val *ViewLocation) *NullableViewLocation {
	return &NullableViewLocation{value: val, isSet: true}
}

func (v NullableViewLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


