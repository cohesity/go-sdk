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

// checks if the UpdateTenantBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTenantBody{}

// UpdateTenantBody Parameters to be specified for updating a tenant.
type UpdateTenantBody struct {
	// Description about the tenant
	Description NullableString `json:"description,omitempty"`
	// Name of the Tenant.
	Name NullableString `json:"name,omitempty"`
	Network *TenantNetwork `json:"network,omitempty"`
}

// NewUpdateTenantBody instantiates a new UpdateTenantBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTenantBody() *UpdateTenantBody {
	this := UpdateTenantBody{}
	return &this
}

// NewUpdateTenantBodyWithDefaults instantiates a new UpdateTenantBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTenantBodyWithDefaults() *UpdateTenantBody {
	this := UpdateTenantBody{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTenantBody) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTenantBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateTenantBody) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UpdateTenantBody) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UpdateTenantBody) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UpdateTenantBody) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTenantBody) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTenantBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateTenantBody) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateTenantBody) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateTenantBody) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateTenantBody) UnsetName() {
	o.Name.Unset()
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *UpdateTenantBody) GetNetwork() TenantNetwork {
	if o == nil || IsNil(o.Network) {
		var ret TenantNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenantBody) GetNetworkOk() (*TenantNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *UpdateTenantBody) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given TenantNetwork and assigns it to the Network field.
func (o *UpdateTenantBody) SetNetwork(v TenantNetwork) {
	o.Network = &v
}

func (o UpdateTenantBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTenantBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	return toSerialize, nil
}

type NullableUpdateTenantBody struct {
	value *UpdateTenantBody
	isSet bool
}

func (v NullableUpdateTenantBody) Get() *UpdateTenantBody {
	return v.value
}

func (v *NullableUpdateTenantBody) Set(val *UpdateTenantBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTenantBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTenantBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTenantBody(val *UpdateTenantBody) *NullableUpdateTenantBody {
	return &NullableUpdateTenantBody{value: val, isSet: true}
}

func (v NullableUpdateTenantBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTenantBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


