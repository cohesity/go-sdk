/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// CloudSpinConfigAllOf struct for CloudSpinConfigAllOf
type CloudSpinConfigAllOf struct {
	Target CloudSpinTarget `json:"target"`
}

// NewCloudSpinConfigAllOf instantiates a new CloudSpinConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudSpinConfigAllOf(target CloudSpinTarget) *CloudSpinConfigAllOf {
	this := CloudSpinConfigAllOf{}
	this.Target = target
	return &this
}

// NewCloudSpinConfigAllOfWithDefaults instantiates a new CloudSpinConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudSpinConfigAllOfWithDefaults() *CloudSpinConfigAllOf {
	this := CloudSpinConfigAllOf{}
	return &this
}

// GetTarget returns the Target field value
func (o *CloudSpinConfigAllOf) GetTarget() CloudSpinTarget {
	if o == nil {
		var ret CloudSpinTarget
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *CloudSpinConfigAllOf) GetTargetOk() (*CloudSpinTarget, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *CloudSpinConfigAllOf) SetTarget(v CloudSpinTarget) {
	o.Target = v
}

func (o CloudSpinConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableCloudSpinConfigAllOf struct {
	value *CloudSpinConfigAllOf
	isSet bool
}

func (v NullableCloudSpinConfigAllOf) Get() *CloudSpinConfigAllOf {
	return v.value
}

func (v *NullableCloudSpinConfigAllOf) Set(val *CloudSpinConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudSpinConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudSpinConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudSpinConfigAllOf(val *CloudSpinConfigAllOf) *NullableCloudSpinConfigAllOf {
	return &NullableCloudSpinConfigAllOf{value: val, isSet: true}
}

func (v NullableCloudSpinConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudSpinConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o CloudSpinConfigAllOf) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}