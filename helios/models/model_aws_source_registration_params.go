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

// AwsSourceRegistrationParams Specifies the paramaters to register an AWS source.
type AwsSourceRegistrationParams struct {
	// Specifies the AWS Subscription type (Commercial/Gov).
	SubscriptionType NullableString `json:"subscriptionType"`
	StandardParams *StandardParams `json:"standardParams,omitempty"`
}

// NewAwsSourceRegistrationParams instantiates a new AwsSourceRegistrationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsSourceRegistrationParams(subscriptionType NullableString) *AwsSourceRegistrationParams {
	this := AwsSourceRegistrationParams{}
	this.SubscriptionType = subscriptionType
	return &this
}

// NewAwsSourceRegistrationParamsWithDefaults instantiates a new AwsSourceRegistrationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsSourceRegistrationParamsWithDefaults() *AwsSourceRegistrationParams {
	this := AwsSourceRegistrationParams{}
	return &this
}

// GetSubscriptionType returns the SubscriptionType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AwsSourceRegistrationParams) GetSubscriptionType() string {
	if o == nil || o.SubscriptionType.Get() == nil {
		var ret string
		return ret
	}

	return *o.SubscriptionType.Get()
}

// GetSubscriptionTypeOk returns a tuple with the SubscriptionType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsSourceRegistrationParams) GetSubscriptionTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionType.Get(), o.SubscriptionType.IsSet()
}

// SetSubscriptionType sets field value
func (o *AwsSourceRegistrationParams) SetSubscriptionType(v string) {
	o.SubscriptionType.Set(&v)
}

// GetStandardParams returns the StandardParams field value if set, zero value otherwise.
func (o *AwsSourceRegistrationParams) GetStandardParams() StandardParams {
	if o == nil || o.StandardParams == nil {
		var ret StandardParams
		return ret
	}
	return *o.StandardParams
}

// GetStandardParamsOk returns a tuple with the StandardParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsSourceRegistrationParams) GetStandardParamsOk() (*StandardParams, bool) {
	if o == nil || o.StandardParams == nil {
		return nil, false
	}
	return o.StandardParams, true
}

// HasStandardParams returns a boolean if a field has been set.
func (o *AwsSourceRegistrationParams) HasStandardParams() bool {
	if o != nil && o.StandardParams != nil {
		return true
	}

	return false
}

// SetStandardParams gets a reference to the given StandardParams and assigns it to the StandardParams field.
func (o *AwsSourceRegistrationParams) SetStandardParams(v StandardParams) {
	o.StandardParams = &v
}

func (o AwsSourceRegistrationParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subscriptionType"] = o.SubscriptionType.Get()
	}
	if o.StandardParams != nil {
		toSerialize["standardParams"] = o.StandardParams
	}
	return json.Marshal(toSerialize)
}

type NullableAwsSourceRegistrationParams struct {
	value *AwsSourceRegistrationParams
	isSet bool
}

func (v NullableAwsSourceRegistrationParams) Get() *AwsSourceRegistrationParams {
	return v.value
}

func (v *NullableAwsSourceRegistrationParams) Set(val *AwsSourceRegistrationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsSourceRegistrationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsSourceRegistrationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsSourceRegistrationParams(val *AwsSourceRegistrationParams) *NullableAwsSourceRegistrationParams {
	return &NullableAwsSourceRegistrationParams{value: val, isSet: true}
}

func (v NullableAwsSourceRegistrationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsSourceRegistrationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


