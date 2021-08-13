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

// AwsGlacierParamsAllOf struct for AwsGlacierParamsAllOf
type AwsGlacierParamsAllOf struct {
	AwsCloudStandardParams *AwsCloudStandardParams `json:"awsCloudStandardParams,omitempty"`
	AwsCloudGovParams *AwsCloudGovParams `json:"awsCloudGovParams,omitempty"`
	AwsCloudC2SParams *AwsCloudC2SParams `json:"awsCloudC2SParams,omitempty"`
}

// NewAwsGlacierParamsAllOf instantiates a new AwsGlacierParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsGlacierParamsAllOf() *AwsGlacierParamsAllOf {
	this := AwsGlacierParamsAllOf{}
	return &this
}

// NewAwsGlacierParamsAllOfWithDefaults instantiates a new AwsGlacierParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsGlacierParamsAllOfWithDefaults() *AwsGlacierParamsAllOf {
	this := AwsGlacierParamsAllOf{}
	return &this
}

// GetAwsCloudStandardParams returns the AwsCloudStandardParams field value if set, zero value otherwise.
func (o *AwsGlacierParamsAllOf) GetAwsCloudStandardParams() AwsCloudStandardParams {
	if o == nil || o.AwsCloudStandardParams == nil {
		var ret AwsCloudStandardParams
		return ret
	}
	return *o.AwsCloudStandardParams
}

// GetAwsCloudStandardParamsOk returns a tuple with the AwsCloudStandardParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsGlacierParamsAllOf) GetAwsCloudStandardParamsOk() (*AwsCloudStandardParams, bool) {
	if o == nil || o.AwsCloudStandardParams == nil {
		return nil, false
	}
	return o.AwsCloudStandardParams, true
}

// HasAwsCloudStandardParams returns a boolean if a field has been set.
func (o *AwsGlacierParamsAllOf) HasAwsCloudStandardParams() bool {
	if o != nil && o.AwsCloudStandardParams != nil {
		return true
	}

	return false
}

// SetAwsCloudStandardParams gets a reference to the given AwsCloudStandardParams and assigns it to the AwsCloudStandardParams field.
func (o *AwsGlacierParamsAllOf) SetAwsCloudStandardParams(v AwsCloudStandardParams) {
	o.AwsCloudStandardParams = &v
}

// GetAwsCloudGovParams returns the AwsCloudGovParams field value if set, zero value otherwise.
func (o *AwsGlacierParamsAllOf) GetAwsCloudGovParams() AwsCloudGovParams {
	if o == nil || o.AwsCloudGovParams == nil {
		var ret AwsCloudGovParams
		return ret
	}
	return *o.AwsCloudGovParams
}

// GetAwsCloudGovParamsOk returns a tuple with the AwsCloudGovParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsGlacierParamsAllOf) GetAwsCloudGovParamsOk() (*AwsCloudGovParams, bool) {
	if o == nil || o.AwsCloudGovParams == nil {
		return nil, false
	}
	return o.AwsCloudGovParams, true
}

// HasAwsCloudGovParams returns a boolean if a field has been set.
func (o *AwsGlacierParamsAllOf) HasAwsCloudGovParams() bool {
	if o != nil && o.AwsCloudGovParams != nil {
		return true
	}

	return false
}

// SetAwsCloudGovParams gets a reference to the given AwsCloudGovParams and assigns it to the AwsCloudGovParams field.
func (o *AwsGlacierParamsAllOf) SetAwsCloudGovParams(v AwsCloudGovParams) {
	o.AwsCloudGovParams = &v
}

// GetAwsCloudC2SParams returns the AwsCloudC2SParams field value if set, zero value otherwise.
func (o *AwsGlacierParamsAllOf) GetAwsCloudC2SParams() AwsCloudC2SParams {
	if o == nil || o.AwsCloudC2SParams == nil {
		var ret AwsCloudC2SParams
		return ret
	}
	return *o.AwsCloudC2SParams
}

// GetAwsCloudC2SParamsOk returns a tuple with the AwsCloudC2SParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsGlacierParamsAllOf) GetAwsCloudC2SParamsOk() (*AwsCloudC2SParams, bool) {
	if o == nil || o.AwsCloudC2SParams == nil {
		return nil, false
	}
	return o.AwsCloudC2SParams, true
}

// HasAwsCloudC2SParams returns a boolean if a field has been set.
func (o *AwsGlacierParamsAllOf) HasAwsCloudC2SParams() bool {
	if o != nil && o.AwsCloudC2SParams != nil {
		return true
	}

	return false
}

// SetAwsCloudC2SParams gets a reference to the given AwsCloudC2SParams and assigns it to the AwsCloudC2SParams field.
func (o *AwsGlacierParamsAllOf) SetAwsCloudC2SParams(v AwsCloudC2SParams) {
	o.AwsCloudC2SParams = &v
}

func (o AwsGlacierParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AwsCloudStandardParams != nil {
		toSerialize["awsCloudStandardParams"] = o.AwsCloudStandardParams
	}
	if o.AwsCloudGovParams != nil {
		toSerialize["awsCloudGovParams"] = o.AwsCloudGovParams
	}
	if o.AwsCloudC2SParams != nil {
		toSerialize["awsCloudC2SParams"] = o.AwsCloudC2SParams
	}
	return json.Marshal(toSerialize)
}

type NullableAwsGlacierParamsAllOf struct {
	value *AwsGlacierParamsAllOf
	isSet bool
}

func (v NullableAwsGlacierParamsAllOf) Get() *AwsGlacierParamsAllOf {
	return v.value
}

func (v *NullableAwsGlacierParamsAllOf) Set(val *AwsGlacierParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsGlacierParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsGlacierParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsGlacierParamsAllOf(val *AwsGlacierParamsAllOf) *NullableAwsGlacierParamsAllOf {
	return &NullableAwsGlacierParamsAllOf{value: val, isSet: true}
}

func (v NullableAwsGlacierParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsGlacierParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o AwsGlacierParamsAllOf) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}