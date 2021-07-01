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

// ProtectedObjectPrivileges ProtectedObjectPrivileges specifies which protected objects are allowed to be accessed by an app instance.
type ProtectedObjectPrivileges struct {
	// Specifies if all, none or specific protected objects are allowed to be accessed. Specifies if all, none or specific protected objects are allowed to be accessed. kNone - None of the protected objects have access. kAll - All the protected objects have access. kSpecific - Only specific protected objects have access.
	ProtectedObjectsprivilegesType NullableString `json:"protectedObjectsprivilegesType,omitempty"`
	// Specifies the ids of the protection sources which are allowed to be accessed in case the privilege type is kSpecific.
	ProtectionSourceIds []int64 `json:"protectionSourceIds,omitempty"`
}

// NewProtectedObjectPrivileges instantiates a new ProtectedObjectPrivileges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtectedObjectPrivileges() *ProtectedObjectPrivileges {
	this := ProtectedObjectPrivileges{}
	return &this
}

// NewProtectedObjectPrivilegesWithDefaults instantiates a new ProtectedObjectPrivileges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtectedObjectPrivilegesWithDefaults() *ProtectedObjectPrivileges {
	this := ProtectedObjectPrivileges{}
	return &this
}

// GetProtectedObjectsprivilegesType returns the ProtectedObjectsprivilegesType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectedObjectPrivileges) GetProtectedObjectsprivilegesType() string {
	if o == nil || o.ProtectedObjectsprivilegesType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProtectedObjectsprivilegesType.Get()
}

// GetProtectedObjectsprivilegesTypeOk returns a tuple with the ProtectedObjectsprivilegesType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectedObjectPrivileges) GetProtectedObjectsprivilegesTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectedObjectsprivilegesType.Get(), o.ProtectedObjectsprivilegesType.IsSet()
}

// HasProtectedObjectsprivilegesType returns a boolean if a field has been set.
func (o *ProtectedObjectPrivileges) HasProtectedObjectsprivilegesType() bool {
	if o != nil && o.ProtectedObjectsprivilegesType.IsSet() {
		return true
	}

	return false
}

// SetProtectedObjectsprivilegesType gets a reference to the given NullableString and assigns it to the ProtectedObjectsprivilegesType field.
func (o *ProtectedObjectPrivileges) SetProtectedObjectsprivilegesType(v string) {
	o.ProtectedObjectsprivilegesType.Set(&v)
}
// SetProtectedObjectsprivilegesTypeNil sets the value for ProtectedObjectsprivilegesType to be an explicit nil
func (o *ProtectedObjectPrivileges) SetProtectedObjectsprivilegesTypeNil() {
	o.ProtectedObjectsprivilegesType.Set(nil)
}

// UnsetProtectedObjectsprivilegesType ensures that no value is present for ProtectedObjectsprivilegesType, not even an explicit nil
func (o *ProtectedObjectPrivileges) UnsetProtectedObjectsprivilegesType() {
	o.ProtectedObjectsprivilegesType.Unset()
}

// GetProtectionSourceIds returns the ProtectionSourceIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProtectedObjectPrivileges) GetProtectionSourceIds() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}
	return o.ProtectionSourceIds
}

// GetProtectionSourceIdsOk returns a tuple with the ProtectionSourceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProtectedObjectPrivileges) GetProtectionSourceIdsOk() (*[]int64, bool) {
	if o == nil || o.ProtectionSourceIds == nil {
		return nil, false
	}
	return &o.ProtectionSourceIds, true
}

// HasProtectionSourceIds returns a boolean if a field has been set.
func (o *ProtectedObjectPrivileges) HasProtectionSourceIds() bool {
	if o != nil && o.ProtectionSourceIds != nil {
		return true
	}

	return false
}

// SetProtectionSourceIds gets a reference to the given []int64 and assigns it to the ProtectionSourceIds field.
func (o *ProtectedObjectPrivileges) SetProtectionSourceIds(v []int64) {
	o.ProtectionSourceIds = v
}

func (o ProtectedObjectPrivileges) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProtectedObjectsprivilegesType.IsSet() {
		toSerialize["protectedObjectsprivilegesType"] = o.ProtectedObjectsprivilegesType.Get()
	}
	if o.ProtectionSourceIds != nil {
		toSerialize["protectionSourceIds"] = o.ProtectionSourceIds
	}
	return json.Marshal(toSerialize)
}

type NullableProtectedObjectPrivileges struct {
	value *ProtectedObjectPrivileges
	isSet bool
}

func (v NullableProtectedObjectPrivileges) Get() *ProtectedObjectPrivileges {
	return v.value
}

func (v *NullableProtectedObjectPrivileges) Set(val *ProtectedObjectPrivileges) {
	v.value = val
	v.isSet = true
}

func (v NullableProtectedObjectPrivileges) IsSet() bool {
	return v.isSet
}

func (v *NullableProtectedObjectPrivileges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtectedObjectPrivileges(val *ProtectedObjectPrivileges) *NullableProtectedObjectPrivileges {
	return &NullableProtectedObjectPrivileges{value: val, isSet: true}
}

func (v NullableProtectedObjectPrivileges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtectedObjectPrivileges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


