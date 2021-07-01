/*
 * Cohesity REST API
 *
 * This API list provides operations for interfacing with the Cohesity Cluster.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesitysdk

import (
	"encoding/json"
)

// TenantViewBoxUpdateParameters Specifies view box update details about a tenant.
type TenantViewBoxUpdateParameters struct {
	// Specifies the unique id of the tenant.
	TenantId NullableString `json:"tenantId,omitempty"`
	// Specifies the ViewBoxIds for respective tenant.
	ViewBoxIds []int64 `json:"viewBoxIds,omitempty"`
}

// NewTenantViewBoxUpdateParameters instantiates a new TenantViewBoxUpdateParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantViewBoxUpdateParameters() *TenantViewBoxUpdateParameters {
	this := TenantViewBoxUpdateParameters{}
	return &this
}

// NewTenantViewBoxUpdateParametersWithDefaults instantiates a new TenantViewBoxUpdateParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantViewBoxUpdateParametersWithDefaults() *TenantViewBoxUpdateParameters {
	this := TenantViewBoxUpdateParameters{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantViewBoxUpdateParameters) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantViewBoxUpdateParameters) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *TenantViewBoxUpdateParameters) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *TenantViewBoxUpdateParameters) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *TenantViewBoxUpdateParameters) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *TenantViewBoxUpdateParameters) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetViewBoxIds returns the ViewBoxIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantViewBoxUpdateParameters) GetViewBoxIds() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}
	return o.ViewBoxIds
}

// GetViewBoxIdsOk returns a tuple with the ViewBoxIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantViewBoxUpdateParameters) GetViewBoxIdsOk() (*[]int64, bool) {
	if o == nil || o.ViewBoxIds == nil {
		return nil, false
	}
	return &o.ViewBoxIds, true
}

// HasViewBoxIds returns a boolean if a field has been set.
func (o *TenantViewBoxUpdateParameters) HasViewBoxIds() bool {
	if o != nil && o.ViewBoxIds != nil {
		return true
	}

	return false
}

// SetViewBoxIds gets a reference to the given []int64 and assigns it to the ViewBoxIds field.
func (o *TenantViewBoxUpdateParameters) SetViewBoxIds(v []int64) {
	o.ViewBoxIds = v
}

func (o TenantViewBoxUpdateParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.ViewBoxIds != nil {
		toSerialize["viewBoxIds"] = o.ViewBoxIds
	}
	return json.Marshal(toSerialize)
}

type NullableTenantViewBoxUpdateParameters struct {
	value *TenantViewBoxUpdateParameters
	isSet bool
}

func (v NullableTenantViewBoxUpdateParameters) Get() *TenantViewBoxUpdateParameters {
	return v.value
}

func (v *NullableTenantViewBoxUpdateParameters) Set(val *TenantViewBoxUpdateParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantViewBoxUpdateParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantViewBoxUpdateParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantViewBoxUpdateParameters(val *TenantViewBoxUpdateParameters) *NullableTenantViewBoxUpdateParameters {
	return &NullableTenantViewBoxUpdateParameters{value: val, isSet: true}
}

func (v NullableTenantViewBoxUpdateParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantViewBoxUpdateParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


