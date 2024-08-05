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

// checks if the WormSpecificTargetParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WormSpecificTargetParams{}

// WormSpecificTargetParams Specifies the parameters which specific to WORM and needs to passed when WORM is enabled for archival External Targets.
type WormSpecificTargetParams struct {
	// Specifies the application id of worm enabled external target.
	ApplicationId NullableString `json:"applicationId,omitempty"`
	// Specifies the encrypted application key of worm enabled external target.
	EncryptedApplicationKey NullableString `json:"encryptedApplicationKey,omitempty"`
	// Specifies the resource group of worm enabled external target.
	ResourceGroup NullableString `json:"resourceGroup,omitempty"`
	// Specifies the subscription id of worm enabled external target.
	SubscriptionId NullableString `json:"subscriptionId,omitempty"`
	// Specifies the tenant id of worm enabled external target.
	TenantId NullableString `json:"tenantId,omitempty"`
}

// NewWormSpecificTargetParams instantiates a new WormSpecificTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWormSpecificTargetParams() *WormSpecificTargetParams {
	this := WormSpecificTargetParams{}
	return &this
}

// NewWormSpecificTargetParamsWithDefaults instantiates a new WormSpecificTargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWormSpecificTargetParamsWithDefaults() *WormSpecificTargetParams {
	this := WormSpecificTargetParams{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WormSpecificTargetParams) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId.Get()) {
		var ret string
		return ret
	}
	return *o.ApplicationId.Get()
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WormSpecificTargetParams) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationId.Get(), o.ApplicationId.IsSet()
}

// HasApplicationId returns a boolean if a field has been set.
func (o *WormSpecificTargetParams) HasApplicationId() bool {
	if o != nil && o.ApplicationId.IsSet() {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given NullableString and assigns it to the ApplicationId field.
func (o *WormSpecificTargetParams) SetApplicationId(v string) {
	o.ApplicationId.Set(&v)
}
// SetApplicationIdNil sets the value for ApplicationId to be an explicit nil
func (o *WormSpecificTargetParams) SetApplicationIdNil() {
	o.ApplicationId.Set(nil)
}

// UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil
func (o *WormSpecificTargetParams) UnsetApplicationId() {
	o.ApplicationId.Unset()
}

// GetEncryptedApplicationKey returns the EncryptedApplicationKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WormSpecificTargetParams) GetEncryptedApplicationKey() string {
	if o == nil || IsNil(o.EncryptedApplicationKey.Get()) {
		var ret string
		return ret
	}
	return *o.EncryptedApplicationKey.Get()
}

// GetEncryptedApplicationKeyOk returns a tuple with the EncryptedApplicationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WormSpecificTargetParams) GetEncryptedApplicationKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EncryptedApplicationKey.Get(), o.EncryptedApplicationKey.IsSet()
}

// HasEncryptedApplicationKey returns a boolean if a field has been set.
func (o *WormSpecificTargetParams) HasEncryptedApplicationKey() bool {
	if o != nil && o.EncryptedApplicationKey.IsSet() {
		return true
	}

	return false
}

// SetEncryptedApplicationKey gets a reference to the given NullableString and assigns it to the EncryptedApplicationKey field.
func (o *WormSpecificTargetParams) SetEncryptedApplicationKey(v string) {
	o.EncryptedApplicationKey.Set(&v)
}
// SetEncryptedApplicationKeyNil sets the value for EncryptedApplicationKey to be an explicit nil
func (o *WormSpecificTargetParams) SetEncryptedApplicationKeyNil() {
	o.EncryptedApplicationKey.Set(nil)
}

// UnsetEncryptedApplicationKey ensures that no value is present for EncryptedApplicationKey, not even an explicit nil
func (o *WormSpecificTargetParams) UnsetEncryptedApplicationKey() {
	o.EncryptedApplicationKey.Unset()
}

// GetResourceGroup returns the ResourceGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WormSpecificTargetParams) GetResourceGroup() string {
	if o == nil || IsNil(o.ResourceGroup.Get()) {
		var ret string
		return ret
	}
	return *o.ResourceGroup.Get()
}

// GetResourceGroupOk returns a tuple with the ResourceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WormSpecificTargetParams) GetResourceGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceGroup.Get(), o.ResourceGroup.IsSet()
}

// HasResourceGroup returns a boolean if a field has been set.
func (o *WormSpecificTargetParams) HasResourceGroup() bool {
	if o != nil && o.ResourceGroup.IsSet() {
		return true
	}

	return false
}

// SetResourceGroup gets a reference to the given NullableString and assigns it to the ResourceGroup field.
func (o *WormSpecificTargetParams) SetResourceGroup(v string) {
	o.ResourceGroup.Set(&v)
}
// SetResourceGroupNil sets the value for ResourceGroup to be an explicit nil
func (o *WormSpecificTargetParams) SetResourceGroupNil() {
	o.ResourceGroup.Set(nil)
}

// UnsetResourceGroup ensures that no value is present for ResourceGroup, not even an explicit nil
func (o *WormSpecificTargetParams) UnsetResourceGroup() {
	o.ResourceGroup.Unset()
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WormSpecificTargetParams) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId.Get()) {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WormSpecificTargetParams) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *WormSpecificTargetParams) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *WormSpecificTargetParams) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *WormSpecificTargetParams) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *WormSpecificTargetParams) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WormSpecificTargetParams) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WormSpecificTargetParams) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *WormSpecificTargetParams) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *WormSpecificTargetParams) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *WormSpecificTargetParams) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *WormSpecificTargetParams) UnsetTenantId() {
	o.TenantId.Unset()
}

func (o WormSpecificTargetParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WormSpecificTargetParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationId.IsSet() {
		toSerialize["applicationId"] = o.ApplicationId.Get()
	}
	if o.EncryptedApplicationKey.IsSet() {
		toSerialize["encryptedApplicationKey"] = o.EncryptedApplicationKey.Get()
	}
	if o.ResourceGroup.IsSet() {
		toSerialize["resourceGroup"] = o.ResourceGroup.Get()
	}
	if o.SubscriptionId.IsSet() {
		toSerialize["subscriptionId"] = o.SubscriptionId.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	return toSerialize, nil
}

type NullableWormSpecificTargetParams struct {
	value *WormSpecificTargetParams
	isSet bool
}

func (v NullableWormSpecificTargetParams) Get() *WormSpecificTargetParams {
	return v.value
}

func (v *NullableWormSpecificTargetParams) Set(val *WormSpecificTargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableWormSpecificTargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableWormSpecificTargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWormSpecificTargetParams(val *WormSpecificTargetParams) *NullableWormSpecificTargetParams {
	return &NullableWormSpecificTargetParams{value: val, isSet: true}
}

func (v NullableWormSpecificTargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWormSpecificTargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


