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

// AwsS3IntelligentParams Specifies the parameters which are specific to AWS related External Targets with storage class S3 Intelligent.
type AwsS3IntelligentParams struct {
	// Specifies the AWS External Target type.
	CloudType NullableString `json:"cloudType"`
	AwsCloudStandardParams *AwsCloudStandardParams `json:"awsCloudStandardParams,omitempty"`
	AwsCloudGovParams *AwsCloudGovParams `json:"awsCloudGovParams,omitempty"`
}

// NewAwsS3IntelligentParams instantiates a new AwsS3IntelligentParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsS3IntelligentParams(cloudType NullableString) *AwsS3IntelligentParams {
	this := AwsS3IntelligentParams{}
	this.CloudType = cloudType
	return &this
}

// NewAwsS3IntelligentParamsWithDefaults instantiates a new AwsS3IntelligentParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsS3IntelligentParamsWithDefaults() *AwsS3IntelligentParams {
	this := AwsS3IntelligentParams{}
	return &this
}

// GetCloudType returns the CloudType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AwsS3IntelligentParams) GetCloudType() string {
	if o == nil || o.CloudType.Get() == nil {
		var ret string
		return ret
	}

	return *o.CloudType.Get()
}

// GetCloudTypeOk returns a tuple with the CloudType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsS3IntelligentParams) GetCloudTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CloudType.Get(), o.CloudType.IsSet()
}

// SetCloudType sets field value
func (o *AwsS3IntelligentParams) SetCloudType(v string) {
	o.CloudType.Set(&v)
}

// GetAwsCloudStandardParams returns the AwsCloudStandardParams field value if set, zero value otherwise.
func (o *AwsS3IntelligentParams) GetAwsCloudStandardParams() AwsCloudStandardParams {
	if o == nil || o.AwsCloudStandardParams == nil {
		var ret AwsCloudStandardParams
		return ret
	}
	return *o.AwsCloudStandardParams
}

// GetAwsCloudStandardParamsOk returns a tuple with the AwsCloudStandardParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsS3IntelligentParams) GetAwsCloudStandardParamsOk() (*AwsCloudStandardParams, bool) {
	if o == nil || o.AwsCloudStandardParams == nil {
		return nil, false
	}
	return o.AwsCloudStandardParams, true
}

// HasAwsCloudStandardParams returns a boolean if a field has been set.
func (o *AwsS3IntelligentParams) HasAwsCloudStandardParams() bool {
	if o != nil && o.AwsCloudStandardParams != nil {
		return true
	}

	return false
}

// SetAwsCloudStandardParams gets a reference to the given AwsCloudStandardParams and assigns it to the AwsCloudStandardParams field.
func (o *AwsS3IntelligentParams) SetAwsCloudStandardParams(v AwsCloudStandardParams) {
	o.AwsCloudStandardParams = &v
}

// GetAwsCloudGovParams returns the AwsCloudGovParams field value if set, zero value otherwise.
func (o *AwsS3IntelligentParams) GetAwsCloudGovParams() AwsCloudGovParams {
	if o == nil || o.AwsCloudGovParams == nil {
		var ret AwsCloudGovParams
		return ret
	}
	return *o.AwsCloudGovParams
}

// GetAwsCloudGovParamsOk returns a tuple with the AwsCloudGovParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsS3IntelligentParams) GetAwsCloudGovParamsOk() (*AwsCloudGovParams, bool) {
	if o == nil || o.AwsCloudGovParams == nil {
		return nil, false
	}
	return o.AwsCloudGovParams, true
}

// HasAwsCloudGovParams returns a boolean if a field has been set.
func (o *AwsS3IntelligentParams) HasAwsCloudGovParams() bool {
	if o != nil && o.AwsCloudGovParams != nil {
		return true
	}

	return false
}

// SetAwsCloudGovParams gets a reference to the given AwsCloudGovParams and assigns it to the AwsCloudGovParams field.
func (o *AwsS3IntelligentParams) SetAwsCloudGovParams(v AwsCloudGovParams) {
	o.AwsCloudGovParams = &v
}

func (o AwsS3IntelligentParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cloudType"] = o.CloudType.Get()
	}
	if o.AwsCloudStandardParams != nil {
		toSerialize["awsCloudStandardParams"] = o.AwsCloudStandardParams
	}
	if o.AwsCloudGovParams != nil {
		toSerialize["awsCloudGovParams"] = o.AwsCloudGovParams
	}
	return json.Marshal(toSerialize)
}

type NullableAwsS3IntelligentParams struct {
	value *AwsS3IntelligentParams
	isSet bool
}

func (v NullableAwsS3IntelligentParams) Get() *AwsS3IntelligentParams {
	return v.value
}

func (v *NullableAwsS3IntelligentParams) Set(val *AwsS3IntelligentParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsS3IntelligentParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsS3IntelligentParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsS3IntelligentParams(val *AwsS3IntelligentParams) *NullableAwsS3IntelligentParams {
	return &NullableAwsS3IntelligentParams{value: val, isSet: true}
}

func (v NullableAwsS3IntelligentParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsS3IntelligentParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o AwsS3IntelligentParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}