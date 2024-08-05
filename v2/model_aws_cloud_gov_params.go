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

// checks if the AwsCloudGovParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsCloudGovParams{}

// AwsCloudGovParams Specifies the parameters which are specific to AWS related External Targets with Cloud Type Gov.
type AwsCloudGovParams struct {
	AuthenticationMethod *AwsAuthenticationMethodsParams `json:"authenticationMethod,omitempty"`
}

// NewAwsCloudGovParams instantiates a new AwsCloudGovParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsCloudGovParams() *AwsCloudGovParams {
	this := AwsCloudGovParams{}
	return &this
}

// NewAwsCloudGovParamsWithDefaults instantiates a new AwsCloudGovParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsCloudGovParamsWithDefaults() *AwsCloudGovParams {
	this := AwsCloudGovParams{}
	return &this
}

// GetAuthenticationMethod returns the AuthenticationMethod field value if set, zero value otherwise.
func (o *AwsCloudGovParams) GetAuthenticationMethod() AwsAuthenticationMethodsParams {
	if o == nil || IsNil(o.AuthenticationMethod) {
		var ret AwsAuthenticationMethodsParams
		return ret
	}
	return *o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsCloudGovParams) GetAuthenticationMethodOk() (*AwsAuthenticationMethodsParams, bool) {
	if o == nil || IsNil(o.AuthenticationMethod) {
		return nil, false
	}
	return o.AuthenticationMethod, true
}

// HasAuthenticationMethod returns a boolean if a field has been set.
func (o *AwsCloudGovParams) HasAuthenticationMethod() bool {
	if o != nil && !IsNil(o.AuthenticationMethod) {
		return true
	}

	return false
}

// SetAuthenticationMethod gets a reference to the given AwsAuthenticationMethodsParams and assigns it to the AuthenticationMethod field.
func (o *AwsCloudGovParams) SetAuthenticationMethod(v AwsAuthenticationMethodsParams) {
	o.AuthenticationMethod = &v
}

func (o AwsCloudGovParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsCloudGovParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthenticationMethod) {
		toSerialize["authenticationMethod"] = o.AuthenticationMethod
	}
	return toSerialize, nil
}

type NullableAwsCloudGovParams struct {
	value *AwsCloudGovParams
	isSet bool
}

func (v NullableAwsCloudGovParams) Get() *AwsCloudGovParams {
	return v.value
}

func (v *NullableAwsCloudGovParams) Set(val *AwsCloudGovParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsCloudGovParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsCloudGovParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsCloudGovParams(val *AwsCloudGovParams) *NullableAwsCloudGovParams {
	return &NullableAwsCloudGovParams{value: val, isSet: true}
}

func (v NullableAwsCloudGovParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsCloudGovParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


