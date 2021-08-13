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

// CommonSourceRegistrationParamsAllOf struct for CommonSourceRegistrationParamsAllOf
type CommonSourceRegistrationParamsAllOf struct {
	// Specifies the endpoint IPaddress, URL or hostname of the host.
	Endpoint string `json:"endpoint"`
	// Specifies the description of the source being registered.
	Description NullableString `json:"description,omitempty"`
}

// NewCommonSourceRegistrationParamsAllOf instantiates a new CommonSourceRegistrationParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonSourceRegistrationParamsAllOf(endpoint string) *CommonSourceRegistrationParamsAllOf {
	this := CommonSourceRegistrationParamsAllOf{}
	this.Endpoint = endpoint
	return &this
}

// NewCommonSourceRegistrationParamsAllOfWithDefaults instantiates a new CommonSourceRegistrationParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonSourceRegistrationParamsAllOfWithDefaults() *CommonSourceRegistrationParamsAllOf {
	this := CommonSourceRegistrationParamsAllOf{}
	return &this
}

// GetEndpoint returns the Endpoint field value
func (o *CommonSourceRegistrationParamsAllOf) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *CommonSourceRegistrationParamsAllOf) GetEndpointOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *CommonSourceRegistrationParamsAllOf) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonSourceRegistrationParamsAllOf) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonSourceRegistrationParamsAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CommonSourceRegistrationParamsAllOf) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CommonSourceRegistrationParamsAllOf) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CommonSourceRegistrationParamsAllOf) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CommonSourceRegistrationParamsAllOf) UnsetDescription() {
	o.Description.Unset()
}

func (o CommonSourceRegistrationParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCommonSourceRegistrationParamsAllOf struct {
	value *CommonSourceRegistrationParamsAllOf
	isSet bool
}

func (v NullableCommonSourceRegistrationParamsAllOf) Get() *CommonSourceRegistrationParamsAllOf {
	return v.value
}

func (v *NullableCommonSourceRegistrationParamsAllOf) Set(val *CommonSourceRegistrationParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonSourceRegistrationParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonSourceRegistrationParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonSourceRegistrationParamsAllOf(val *CommonSourceRegistrationParamsAllOf) *NullableCommonSourceRegistrationParamsAllOf {
	return &NullableCommonSourceRegistrationParamsAllOf{value: val, isSet: true}
}

func (v NullableCommonSourceRegistrationParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonSourceRegistrationParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


