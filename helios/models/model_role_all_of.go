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

// RoleAllOf struct for RoleAllOf
type RoleAllOf struct {
	// Specifies the Role label.
	Label NullableString `json:"label,omitempty"`
	// Specifies the timestamp when the Role is created in milliseconds.
	CreatedTimestampMsecs NullableInt64 `json:"createdTimestampMsecs,omitempty"`
	// Specifies the timestamp when the Role is last updated in milliseconds.
	LastUpdatedTimestampMsecs NullableInt64 `json:"lastUpdatedTimestampMsecs,omitempty"`
	// Specifies if the Role is created by user.
	IsUserCreatedRole NullableBool `json:"isUserCreatedRole,omitempty"`
	// Specifies the list of tenant ids who have access to this Role.
	TenantIds []string `json:"tenantIds,omitempty"`
}

// NewRoleAllOf instantiates a new RoleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAllOf() *RoleAllOf {
	this := RoleAllOf{}
	return &this
}

// NewRoleAllOfWithDefaults instantiates a new RoleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAllOfWithDefaults() *RoleAllOf {
	this := RoleAllOf{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleAllOf) GetLabel() string {
	if o == nil || o.Label.Get() == nil {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleAllOf) GetLabelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *RoleAllOf) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *RoleAllOf) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *RoleAllOf) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *RoleAllOf) UnsetLabel() {
	o.Label.Unset()
}

// GetCreatedTimestampMsecs returns the CreatedTimestampMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleAllOf) GetCreatedTimestampMsecs() int64 {
	if o == nil || o.CreatedTimestampMsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CreatedTimestampMsecs.Get()
}

// GetCreatedTimestampMsecsOk returns a tuple with the CreatedTimestampMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleAllOf) GetCreatedTimestampMsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedTimestampMsecs.Get(), o.CreatedTimestampMsecs.IsSet()
}

// HasCreatedTimestampMsecs returns a boolean if a field has been set.
func (o *RoleAllOf) HasCreatedTimestampMsecs() bool {
	if o != nil && o.CreatedTimestampMsecs.IsSet() {
		return true
	}

	return false
}

// SetCreatedTimestampMsecs gets a reference to the given NullableInt64 and assigns it to the CreatedTimestampMsecs field.
func (o *RoleAllOf) SetCreatedTimestampMsecs(v int64) {
	o.CreatedTimestampMsecs.Set(&v)
}
// SetCreatedTimestampMsecsNil sets the value for CreatedTimestampMsecs to be an explicit nil
func (o *RoleAllOf) SetCreatedTimestampMsecsNil() {
	o.CreatedTimestampMsecs.Set(nil)
}

// UnsetCreatedTimestampMsecs ensures that no value is present for CreatedTimestampMsecs, not even an explicit nil
func (o *RoleAllOf) UnsetCreatedTimestampMsecs() {
	o.CreatedTimestampMsecs.Unset()
}

// GetLastUpdatedTimestampMsecs returns the LastUpdatedTimestampMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleAllOf) GetLastUpdatedTimestampMsecs() int64 {
	if o == nil || o.LastUpdatedTimestampMsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdatedTimestampMsecs.Get()
}

// GetLastUpdatedTimestampMsecsOk returns a tuple with the LastUpdatedTimestampMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleAllOf) GetLastUpdatedTimestampMsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastUpdatedTimestampMsecs.Get(), o.LastUpdatedTimestampMsecs.IsSet()
}

// HasLastUpdatedTimestampMsecs returns a boolean if a field has been set.
func (o *RoleAllOf) HasLastUpdatedTimestampMsecs() bool {
	if o != nil && o.LastUpdatedTimestampMsecs.IsSet() {
		return true
	}

	return false
}

// SetLastUpdatedTimestampMsecs gets a reference to the given NullableInt64 and assigns it to the LastUpdatedTimestampMsecs field.
func (o *RoleAllOf) SetLastUpdatedTimestampMsecs(v int64) {
	o.LastUpdatedTimestampMsecs.Set(&v)
}
// SetLastUpdatedTimestampMsecsNil sets the value for LastUpdatedTimestampMsecs to be an explicit nil
func (o *RoleAllOf) SetLastUpdatedTimestampMsecsNil() {
	o.LastUpdatedTimestampMsecs.Set(nil)
}

// UnsetLastUpdatedTimestampMsecs ensures that no value is present for LastUpdatedTimestampMsecs, not even an explicit nil
func (o *RoleAllOf) UnsetLastUpdatedTimestampMsecs() {
	o.LastUpdatedTimestampMsecs.Unset()
}

// GetIsUserCreatedRole returns the IsUserCreatedRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleAllOf) GetIsUserCreatedRole() bool {
	if o == nil || o.IsUserCreatedRole.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsUserCreatedRole.Get()
}

// GetIsUserCreatedRoleOk returns a tuple with the IsUserCreatedRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleAllOf) GetIsUserCreatedRoleOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsUserCreatedRole.Get(), o.IsUserCreatedRole.IsSet()
}

// HasIsUserCreatedRole returns a boolean if a field has been set.
func (o *RoleAllOf) HasIsUserCreatedRole() bool {
	if o != nil && o.IsUserCreatedRole.IsSet() {
		return true
	}

	return false
}

// SetIsUserCreatedRole gets a reference to the given NullableBool and assigns it to the IsUserCreatedRole field.
func (o *RoleAllOf) SetIsUserCreatedRole(v bool) {
	o.IsUserCreatedRole.Set(&v)
}
// SetIsUserCreatedRoleNil sets the value for IsUserCreatedRole to be an explicit nil
func (o *RoleAllOf) SetIsUserCreatedRoleNil() {
	o.IsUserCreatedRole.Set(nil)
}

// UnsetIsUserCreatedRole ensures that no value is present for IsUserCreatedRole, not even an explicit nil
func (o *RoleAllOf) UnsetIsUserCreatedRole() {
	o.IsUserCreatedRole.Unset()
}

// GetTenantIds returns the TenantIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleAllOf) GetTenantIds() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.TenantIds
}

// GetTenantIdsOk returns a tuple with the TenantIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleAllOf) GetTenantIdsOk() (*[]string, bool) {
	if o == nil || o.TenantIds == nil {
		return nil, false
	}
	return &o.TenantIds, true
}

// HasTenantIds returns a boolean if a field has been set.
func (o *RoleAllOf) HasTenantIds() bool {
	if o != nil && o.TenantIds != nil {
		return true
	}

	return false
}

// SetTenantIds gets a reference to the given []string and assigns it to the TenantIds field.
func (o *RoleAllOf) SetTenantIds(v []string) {
	o.TenantIds = v
}

func (o RoleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if o.CreatedTimestampMsecs.IsSet() {
		toSerialize["createdTimestampMsecs"] = o.CreatedTimestampMsecs.Get()
	}
	if o.LastUpdatedTimestampMsecs.IsSet() {
		toSerialize["lastUpdatedTimestampMsecs"] = o.LastUpdatedTimestampMsecs.Get()
	}
	if o.IsUserCreatedRole.IsSet() {
		toSerialize["isUserCreatedRole"] = o.IsUserCreatedRole.Get()
	}
	if o.TenantIds != nil {
		toSerialize["tenantIds"] = o.TenantIds
	}
	return json.Marshal(toSerialize)
}

type NullableRoleAllOf struct {
	value *RoleAllOf
	isSet bool
}

func (v NullableRoleAllOf) Get() *RoleAllOf {
	return v.value
}

func (v *NullableRoleAllOf) Set(val *RoleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAllOf(val *RoleAllOf) *NullableRoleAllOf {
	return &NullableRoleAllOf{value: val, isSet: true}
}

func (v NullableRoleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


