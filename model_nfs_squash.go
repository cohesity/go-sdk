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

// NfsSquash struct for NfsSquash
type NfsSquash struct {
	// GID mapped for all clients.
	Gid NullableInt32 `json:"gid,omitempty"`
	// UID mapped for all clients.
	Uid NullableInt32 `json:"uid,omitempty"`
}

// NewNfsSquash instantiates a new NfsSquash object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfsSquash() *NfsSquash {
	this := NfsSquash{}
	return &this
}

// NewNfsSquashWithDefaults instantiates a new NfsSquash object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfsSquashWithDefaults() *NfsSquash {
	this := NfsSquash{}
	return &this
}

// GetGid returns the Gid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NfsSquash) GetGid() int32 {
	if o == nil || o.Gid.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Gid.Get()
}

// GetGidOk returns a tuple with the Gid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NfsSquash) GetGidOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Gid.Get(), o.Gid.IsSet()
}

// HasGid returns a boolean if a field has been set.
func (o *NfsSquash) HasGid() bool {
	if o != nil && o.Gid.IsSet() {
		return true
	}

	return false
}

// SetGid gets a reference to the given NullableInt32 and assigns it to the Gid field.
func (o *NfsSquash) SetGid(v int32) {
	o.Gid.Set(&v)
}
// SetGidNil sets the value for Gid to be an explicit nil
func (o *NfsSquash) SetGidNil() {
	o.Gid.Set(nil)
}

// UnsetGid ensures that no value is present for Gid, not even an explicit nil
func (o *NfsSquash) UnsetGid() {
	o.Gid.Unset()
}

// GetUid returns the Uid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NfsSquash) GetUid() int32 {
	if o == nil || o.Uid.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Uid.Get()
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NfsSquash) GetUidOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Uid.Get(), o.Uid.IsSet()
}

// HasUid returns a boolean if a field has been set.
func (o *NfsSquash) HasUid() bool {
	if o != nil && o.Uid.IsSet() {
		return true
	}

	return false
}

// SetUid gets a reference to the given NullableInt32 and assigns it to the Uid field.
func (o *NfsSquash) SetUid(v int32) {
	o.Uid.Set(&v)
}
// SetUidNil sets the value for Uid to be an explicit nil
func (o *NfsSquash) SetUidNil() {
	o.Uid.Set(nil)
}

// UnsetUid ensures that no value is present for Uid, not even an explicit nil
func (o *NfsSquash) UnsetUid() {
	o.Uid.Unset()
}

func (o NfsSquash) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Gid.IsSet() {
		toSerialize["gid"] = o.Gid.Get()
	}
	if o.Uid.IsSet() {
		toSerialize["uid"] = o.Uid.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableNfsSquash struct {
	value *NfsSquash
	isSet bool
}

func (v NullableNfsSquash) Get() *NfsSquash {
	return v.value
}

func (v *NullableNfsSquash) Set(val *NfsSquash) {
	v.value = val
	v.isSet = true
}

func (v NullableNfsSquash) IsSet() bool {
	return v.isSet
}

func (v *NullableNfsSquash) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfsSquash(val *NfsSquash) *NullableNfsSquash {
	return &NullableNfsSquash{value: val, isSet: true}
}

func (v NullableNfsSquash) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfsSquash) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


