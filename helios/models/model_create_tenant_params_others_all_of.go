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

// CreateTenantParamsOthersAllOf struct for CreateTenantParamsOthersAllOf
type CreateTenantParamsOthersAllOf struct {
	Network *TenantNetwork `json:"network,omitempty"`
}

// NewCreateTenantParamsOthersAllOf instantiates a new CreateTenantParamsOthersAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTenantParamsOthersAllOf() *CreateTenantParamsOthersAllOf {
	this := CreateTenantParamsOthersAllOf{}
	return &this
}

// NewCreateTenantParamsOthersAllOfWithDefaults instantiates a new CreateTenantParamsOthersAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTenantParamsOthersAllOfWithDefaults() *CreateTenantParamsOthersAllOf {
	this := CreateTenantParamsOthersAllOf{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *CreateTenantParamsOthersAllOf) GetNetwork() TenantNetwork {
	if o == nil || o.Network == nil {
		var ret TenantNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenantParamsOthersAllOf) GetNetworkOk() (*TenantNetwork, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *CreateTenantParamsOthersAllOf) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given TenantNetwork and assigns it to the Network field.
func (o *CreateTenantParamsOthersAllOf) SetNetwork(v TenantNetwork) {
	o.Network = &v
}

func (o CreateTenantParamsOthersAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTenantParamsOthersAllOf struct {
	value *CreateTenantParamsOthersAllOf
	isSet bool
}

func (v NullableCreateTenantParamsOthersAllOf) Get() *CreateTenantParamsOthersAllOf {
	return v.value
}

func (v *NullableCreateTenantParamsOthersAllOf) Set(val *CreateTenantParamsOthersAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTenantParamsOthersAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTenantParamsOthersAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTenantParamsOthersAllOf(val *CreateTenantParamsOthersAllOf) *NullableCreateTenantParamsOthersAllOf {
	return &NullableCreateTenantParamsOthersAllOf{value: val, isSet: true}
}

func (v NullableCreateTenantParamsOthersAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTenantParamsOthersAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


