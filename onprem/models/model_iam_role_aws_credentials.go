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

// IamRoleAwsCredentials Specifies the credentials to register a commercial AWS
type IamRoleAwsCredentials struct {
	// Specifies the IAM role which will be used to access the security credentials required for API calls. This should have all the permissions required for the tenant's use case. In case of DMaaS this will be the Tenant's IAM role ARN. This is assumed only after the cp_iam_role_arn(control plane role) is assumed
	IamRoleArn NullableString `json:"iamRoleArn"`
	// This is only applicable in case of DMaaS. Control plane IAM role ARN, this is first assumed by the dataplane(cluster). Then we assume the iam_role_arn which is tenant's IAM role with all required permissions.
	CpIamRoleArn NullableString `json:"cpIamRoleArn,omitempty"`
}

// NewIamRoleAwsCredentials instantiates a new IamRoleAwsCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamRoleAwsCredentials(iamRoleArn NullableString) *IamRoleAwsCredentials {
	this := IamRoleAwsCredentials{}
	this.IamRoleArn = iamRoleArn
	return &this
}

// NewIamRoleAwsCredentialsWithDefaults instantiates a new IamRoleAwsCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamRoleAwsCredentialsWithDefaults() *IamRoleAwsCredentials {
	this := IamRoleAwsCredentials{}
	return &this
}

// GetIamRoleArn returns the IamRoleArn field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IamRoleAwsCredentials) GetIamRoleArn() string {
	if o == nil || o.IamRoleArn.Get() == nil {
		var ret string
		return ret
	}

	return *o.IamRoleArn.Get()
}

// GetIamRoleArnOk returns a tuple with the IamRoleArn field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamRoleAwsCredentials) GetIamRoleArnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IamRoleArn.Get(), o.IamRoleArn.IsSet()
}

// SetIamRoleArn sets field value
func (o *IamRoleAwsCredentials) SetIamRoleArn(v string) {
	o.IamRoleArn.Set(&v)
}

// GetCpIamRoleArn returns the CpIamRoleArn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamRoleAwsCredentials) GetCpIamRoleArn() string {
	if o == nil || o.CpIamRoleArn.Get() == nil {
		var ret string
		return ret
	}
	return *o.CpIamRoleArn.Get()
}

// GetCpIamRoleArnOk returns a tuple with the CpIamRoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamRoleAwsCredentials) GetCpIamRoleArnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CpIamRoleArn.Get(), o.CpIamRoleArn.IsSet()
}

// HasCpIamRoleArn returns a boolean if a field has been set.
func (o *IamRoleAwsCredentials) HasCpIamRoleArn() bool {
	if o != nil && o.CpIamRoleArn.IsSet() {
		return true
	}

	return false
}

// SetCpIamRoleArn gets a reference to the given NullableString and assigns it to the CpIamRoleArn field.
func (o *IamRoleAwsCredentials) SetCpIamRoleArn(v string) {
	o.CpIamRoleArn.Set(&v)
}
// SetCpIamRoleArnNil sets the value for CpIamRoleArn to be an explicit nil
func (o *IamRoleAwsCredentials) SetCpIamRoleArnNil() {
	o.CpIamRoleArn.Set(nil)
}

// UnsetCpIamRoleArn ensures that no value is present for CpIamRoleArn, not even an explicit nil
func (o *IamRoleAwsCredentials) UnsetCpIamRoleArn() {
	o.CpIamRoleArn.Unset()
}

func (o IamRoleAwsCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["iamRoleArn"] = o.IamRoleArn.Get()
	}
	if o.CpIamRoleArn.IsSet() {
		toSerialize["cpIamRoleArn"] = o.CpIamRoleArn.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIamRoleAwsCredentials struct {
	value *IamRoleAwsCredentials
	isSet bool
}

func (v NullableIamRoleAwsCredentials) Get() *IamRoleAwsCredentials {
	return v.value
}

func (v *NullableIamRoleAwsCredentials) Set(val *IamRoleAwsCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableIamRoleAwsCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableIamRoleAwsCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamRoleAwsCredentials(val *IamRoleAwsCredentials) *NullableIamRoleAwsCredentials {
	return &NullableIamRoleAwsCredentials{value: val, isSet: true}
}

func (v NullableIamRoleAwsCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamRoleAwsCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o IamRoleAwsCredentials) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}