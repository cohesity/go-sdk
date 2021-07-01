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

// UserId Specifies the mapping between an Unix and an SMB SID.
type UserId struct {
	// If interested in a user via smb_client, include SID. Otherwise, If a valid unix-id to SID mappings are available (i.e., when mixed mode is enabled) the server will perform the necessary id mapping and return the correct usage irrespective of whether the unix id / SID is provided. The string is of following format - S-1-IdentifierAuthority-SubAuthority1-SubAuthority2-...-SubAuthorityn.
	Sid NullableString `json:"sid,omitempty"`
	// If interested in a user via unix-identifier, include UnixUid. Otherwise, If a valid unix-id to SID mappings are available (i.e., when mixed mode is enabled) the server will perform the necessary id mapping and return the correct usage irrespective of whether the unix id / SID is provided.
	UnixUid NullableInt32 `json:"unixUid,omitempty"`
}

// NewUserId instantiates a new UserId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserId() *UserId {
	this := UserId{}
	return &this
}

// NewUserIdWithDefaults instantiates a new UserId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdWithDefaults() *UserId {
	this := UserId{}
	return &this
}

// GetSid returns the Sid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserId) GetSid() string {
	if o == nil || o.Sid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Sid.Get()
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserId) GetSidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Sid.Get(), o.Sid.IsSet()
}

// HasSid returns a boolean if a field has been set.
func (o *UserId) HasSid() bool {
	if o != nil && o.Sid.IsSet() {
		return true
	}

	return false
}

// SetSid gets a reference to the given NullableString and assigns it to the Sid field.
func (o *UserId) SetSid(v string) {
	o.Sid.Set(&v)
}
// SetSidNil sets the value for Sid to be an explicit nil
func (o *UserId) SetSidNil() {
	o.Sid.Set(nil)
}

// UnsetSid ensures that no value is present for Sid, not even an explicit nil
func (o *UserId) UnsetSid() {
	o.Sid.Unset()
}

// GetUnixUid returns the UnixUid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserId) GetUnixUid() int32 {
	if o == nil || o.UnixUid.Get() == nil {
		var ret int32
		return ret
	}
	return *o.UnixUid.Get()
}

// GetUnixUidOk returns a tuple with the UnixUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserId) GetUnixUidOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnixUid.Get(), o.UnixUid.IsSet()
}

// HasUnixUid returns a boolean if a field has been set.
func (o *UserId) HasUnixUid() bool {
	if o != nil && o.UnixUid.IsSet() {
		return true
	}

	return false
}

// SetUnixUid gets a reference to the given NullableInt32 and assigns it to the UnixUid field.
func (o *UserId) SetUnixUid(v int32) {
	o.UnixUid.Set(&v)
}
// SetUnixUidNil sets the value for UnixUid to be an explicit nil
func (o *UserId) SetUnixUidNil() {
	o.UnixUid.Set(nil)
}

// UnsetUnixUid ensures that no value is present for UnixUid, not even an explicit nil
func (o *UserId) UnsetUnixUid() {
	o.UnixUid.Unset()
}

func (o UserId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Sid.IsSet() {
		toSerialize["sid"] = o.Sid.Get()
	}
	if o.UnixUid.IsSet() {
		toSerialize["unixUid"] = o.UnixUid.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUserId struct {
	value *UserId
	isSet bool
}

func (v NullableUserId) Get() *UserId {
	return v.value
}

func (v *NullableUserId) Set(val *UserId) {
	v.value = val
	v.isSet = true
}

func (v NullableUserId) IsSet() bool {
	return v.isSet
}

func (v *NullableUserId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserId(val *UserId) *NullableUserId {
	return &NullableUserId{value: val, isSet: true}
}

func (v NullableUserId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


