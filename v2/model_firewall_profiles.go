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

// checks if the FirewallProfiles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirewallProfiles{}

// FirewallProfiles Specifies the firewall profile & their attachments.
type FirewallProfiles struct {
	Profiles []FirewallProfile `json:"profiles,omitempty"`
}

// NewFirewallProfiles instantiates a new FirewallProfiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallProfiles() *FirewallProfiles {
	this := FirewallProfiles{}
	return &this
}

// NewFirewallProfilesWithDefaults instantiates a new FirewallProfiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallProfilesWithDefaults() *FirewallProfiles {
	this := FirewallProfiles{}
	return &this
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *FirewallProfiles) GetProfiles() []FirewallProfile {
	if o == nil || IsNil(o.Profiles) {
		var ret []FirewallProfile
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallProfiles) GetProfilesOk() ([]FirewallProfile, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *FirewallProfiles) HasProfiles() bool {
	if o != nil && !IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []FirewallProfile and assigns it to the Profiles field.
func (o *FirewallProfiles) SetProfiles(v []FirewallProfile) {
	o.Profiles = v
}

func (o FirewallProfiles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirewallProfiles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Profiles) {
		toSerialize["profiles"] = o.Profiles
	}
	return toSerialize, nil
}

type NullableFirewallProfiles struct {
	value *FirewallProfiles
	isSet bool
}

func (v NullableFirewallProfiles) Get() *FirewallProfiles {
	return v.value
}

func (v *NullableFirewallProfiles) Set(val *FirewallProfiles) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallProfiles) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallProfiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallProfiles(val *FirewallProfiles) *NullableFirewallProfiles {
	return &NullableFirewallProfiles{value: val, isSet: true}
}

func (v NullableFirewallProfiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallProfiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


