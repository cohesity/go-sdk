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

// AdDomainIdentity AD domain identity information.
type AdDomainIdentity struct {
	// Specifies distinguished name of the domain.
	Dn NullableString `json:"dn,omitempty"`
	// Specifies Unique objectGUID for an AD domain.
	Guid NullableString `json:"guid,omitempty"`
	// Specifies display name of the domain.
	Name NullableString `json:"name,omitempty"`
	// Specifies domain SID.
	Sid NullableString `json:"sid,omitempty"`
}

// NewAdDomainIdentity instantiates a new AdDomainIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdDomainIdentity() *AdDomainIdentity {
	this := AdDomainIdentity{}
	return &this
}

// NewAdDomainIdentityWithDefaults instantiates a new AdDomainIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdDomainIdentityWithDefaults() *AdDomainIdentity {
	this := AdDomainIdentity{}
	return &this
}

// GetDn returns the Dn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdDomainIdentity) GetDn() string {
	if o == nil || o.Dn.Get() == nil {
		var ret string
		return ret
	}
	return *o.Dn.Get()
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdDomainIdentity) GetDnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Dn.Get(), o.Dn.IsSet()
}

// HasDn returns a boolean if a field has been set.
func (o *AdDomainIdentity) HasDn() bool {
	if o != nil && o.Dn.IsSet() {
		return true
	}

	return false
}

// SetDn gets a reference to the given NullableString and assigns it to the Dn field.
func (o *AdDomainIdentity) SetDn(v string) {
	o.Dn.Set(&v)
}
// SetDnNil sets the value for Dn to be an explicit nil
func (o *AdDomainIdentity) SetDnNil() {
	o.Dn.Set(nil)
}

// UnsetDn ensures that no value is present for Dn, not even an explicit nil
func (o *AdDomainIdentity) UnsetDn() {
	o.Dn.Unset()
}

// GetGuid returns the Guid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdDomainIdentity) GetGuid() string {
	if o == nil || o.Guid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Guid.Get()
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdDomainIdentity) GetGuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Guid.Get(), o.Guid.IsSet()
}

// HasGuid returns a boolean if a field has been set.
func (o *AdDomainIdentity) HasGuid() bool {
	if o != nil && o.Guid.IsSet() {
		return true
	}

	return false
}

// SetGuid gets a reference to the given NullableString and assigns it to the Guid field.
func (o *AdDomainIdentity) SetGuid(v string) {
	o.Guid.Set(&v)
}
// SetGuidNil sets the value for Guid to be an explicit nil
func (o *AdDomainIdentity) SetGuidNil() {
	o.Guid.Set(nil)
}

// UnsetGuid ensures that no value is present for Guid, not even an explicit nil
func (o *AdDomainIdentity) UnsetGuid() {
	o.Guid.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdDomainIdentity) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdDomainIdentity) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AdDomainIdentity) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AdDomainIdentity) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AdDomainIdentity) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AdDomainIdentity) UnsetName() {
	o.Name.Unset()
}

// GetSid returns the Sid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdDomainIdentity) GetSid() string {
	if o == nil || o.Sid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Sid.Get()
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdDomainIdentity) GetSidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Sid.Get(), o.Sid.IsSet()
}

// HasSid returns a boolean if a field has been set.
func (o *AdDomainIdentity) HasSid() bool {
	if o != nil && o.Sid.IsSet() {
		return true
	}

	return false
}

// SetSid gets a reference to the given NullableString and assigns it to the Sid field.
func (o *AdDomainIdentity) SetSid(v string) {
	o.Sid.Set(&v)
}
// SetSidNil sets the value for Sid to be an explicit nil
func (o *AdDomainIdentity) SetSidNil() {
	o.Sid.Set(nil)
}

// UnsetSid ensures that no value is present for Sid, not even an explicit nil
func (o *AdDomainIdentity) UnsetSid() {
	o.Sid.Unset()
}

func (o AdDomainIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dn.IsSet() {
		toSerialize["dn"] = o.Dn.Get()
	}
	if o.Guid.IsSet() {
		toSerialize["guid"] = o.Guid.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Sid.IsSet() {
		toSerialize["sid"] = o.Sid.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAdDomainIdentity struct {
	value *AdDomainIdentity
	isSet bool
}

func (v NullableAdDomainIdentity) Get() *AdDomainIdentity {
	return v.value
}

func (v *NullableAdDomainIdentity) Set(val *AdDomainIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableAdDomainIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableAdDomainIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdDomainIdentity(val *AdDomainIdentity) *NullableAdDomainIdentity {
	return &NullableAdDomainIdentity{value: val, isSet: true}
}

func (v NullableAdDomainIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdDomainIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


