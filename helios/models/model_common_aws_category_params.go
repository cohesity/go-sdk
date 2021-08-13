/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package helios
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/helios/utils"
)

var _ = NullableBool{}

// CommonAWSCategoryParams Specifies the cloud category parameter which are specific to AWS related External Targets.
type CommonAWSCategoryParams struct {
	// Specifies the AWS External Target type.
	CloudType NullableString `json:"cloudType"`
}

// NewCommonAWSCategoryParams instantiates a new CommonAWSCategoryParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonAWSCategoryParams(cloudType NullableString) *CommonAWSCategoryParams {
	this := CommonAWSCategoryParams{}
	this.CloudType = cloudType
	return &this
}

// NewCommonAWSCategoryParamsWithDefaults instantiates a new CommonAWSCategoryParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonAWSCategoryParamsWithDefaults() *CommonAWSCategoryParams {
	this := CommonAWSCategoryParams{}
	return &this
}

// GetCloudType returns the CloudType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonAWSCategoryParams) GetCloudType() string {
	if o == nil || o.CloudType.Get() == nil {
		var ret string
		return ret
	}

	return *o.CloudType.Get()
}

// GetCloudTypeOk returns a tuple with the CloudType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonAWSCategoryParams) GetCloudTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CloudType.Get(), o.CloudType.IsSet()
}

// SetCloudType sets field value
func (o *CommonAWSCategoryParams) SetCloudType(v string) {
	o.CloudType.Set(&v)
}

func (o CommonAWSCategoryParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cloudType"] = o.CloudType.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCommonAWSCategoryParams struct {
	value *CommonAWSCategoryParams
	isSet bool
}

func (v NullableCommonAWSCategoryParams) Get() *CommonAWSCategoryParams {
	return v.value
}

func (v *NullableCommonAWSCategoryParams) Set(val *CommonAWSCategoryParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonAWSCategoryParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonAWSCategoryParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonAWSCategoryParams(val *CommonAWSCategoryParams) *NullableCommonAWSCategoryParams {
	return &NullableCommonAWSCategoryParams{value: val, isSet: true}
}

func (v NullableCommonAWSCategoryParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonAWSCategoryParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


