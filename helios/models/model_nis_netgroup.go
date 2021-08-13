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

// NisNetgroup Specifies an NIS netgroup.
type NisNetgroup struct {
	// Specifies the netgroup name.
	Name NullableString `json:"name"`
	// Specifies the domain name for the netgroup.
	Domain NullableString `json:"domain"`
	// Specifies NFS protocol acess level for clients from the netgroup.
	NfsAccess NullableString `json:"nfsAccess,omitempty"`
	// Specifies which nfsSquash Mounted. 'kNone' mounts none. 'kRootSquash' mounts nfsRootSquash. Whether clients from this subnet can mount as root on NFS. 'kAllSquash' mounts nfsAllSquash. Whether all clients from this subnet can map view with view_all_squash_uid/view_all_squash_gid configured in the view.
	NfsSquash NullableString `json:"nfsSquash,omitempty"`
}

// NewNisNetgroup instantiates a new NisNetgroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNisNetgroup(name NullableString, domain NullableString) *NisNetgroup {
	this := NisNetgroup{}
	this.Name = name
	this.Domain = domain
	return &this
}

// NewNisNetgroupWithDefaults instantiates a new NisNetgroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNisNetgroupWithDefaults() *NisNetgroup {
	this := NisNetgroup{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NisNetgroup) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NisNetgroup) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *NisNetgroup) SetName(v string) {
	o.Name.Set(&v)
}

// GetDomain returns the Domain field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NisNetgroup) GetDomain() string {
	if o == nil || o.Domain.Get() == nil {
		var ret string
		return ret
	}

	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NisNetgroup) GetDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// SetDomain sets field value
func (o *NisNetgroup) SetDomain(v string) {
	o.Domain.Set(&v)
}

// GetNfsAccess returns the NfsAccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NisNetgroup) GetNfsAccess() string {
	if o == nil || o.NfsAccess.Get() == nil {
		var ret string
		return ret
	}
	return *o.NfsAccess.Get()
}

// GetNfsAccessOk returns a tuple with the NfsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NisNetgroup) GetNfsAccessOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NfsAccess.Get(), o.NfsAccess.IsSet()
}

// HasNfsAccess returns a boolean if a field has been set.
func (o *NisNetgroup) HasNfsAccess() bool {
	if o != nil && o.NfsAccess.IsSet() {
		return true
	}

	return false
}

// SetNfsAccess gets a reference to the given NullableString and assigns it to the NfsAccess field.
func (o *NisNetgroup) SetNfsAccess(v string) {
	o.NfsAccess.Set(&v)
}
// SetNfsAccessNil sets the value for NfsAccess to be an explicit nil
func (o *NisNetgroup) SetNfsAccessNil() {
	o.NfsAccess.Set(nil)
}

// UnsetNfsAccess ensures that no value is present for NfsAccess, not even an explicit nil
func (o *NisNetgroup) UnsetNfsAccess() {
	o.NfsAccess.Unset()
}

// GetNfsSquash returns the NfsSquash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NisNetgroup) GetNfsSquash() string {
	if o == nil || o.NfsSquash.Get() == nil {
		var ret string
		return ret
	}
	return *o.NfsSquash.Get()
}

// GetNfsSquashOk returns a tuple with the NfsSquash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NisNetgroup) GetNfsSquashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NfsSquash.Get(), o.NfsSquash.IsSet()
}

// HasNfsSquash returns a boolean if a field has been set.
func (o *NisNetgroup) HasNfsSquash() bool {
	if o != nil && o.NfsSquash.IsSet() {
		return true
	}

	return false
}

// SetNfsSquash gets a reference to the given NullableString and assigns it to the NfsSquash field.
func (o *NisNetgroup) SetNfsSquash(v string) {
	o.NfsSquash.Set(&v)
}
// SetNfsSquashNil sets the value for NfsSquash to be an explicit nil
func (o *NisNetgroup) SetNfsSquashNil() {
	o.NfsSquash.Set(nil)
}

// UnsetNfsSquash ensures that no value is present for NfsSquash, not even an explicit nil
func (o *NisNetgroup) UnsetNfsSquash() {
	o.NfsSquash.Unset()
}

func (o NisNetgroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["domain"] = o.Domain.Get()
	}
	if o.NfsAccess.IsSet() {
		toSerialize["nfsAccess"] = o.NfsAccess.Get()
	}
	if o.NfsSquash.IsSet() {
		toSerialize["nfsSquash"] = o.NfsSquash.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableNisNetgroup struct {
	value *NisNetgroup
	isSet bool
}

func (v NullableNisNetgroup) Get() *NisNetgroup {
	return v.value
}

func (v *NullableNisNetgroup) Set(val *NisNetgroup) {
	v.value = val
	v.isSet = true
}

func (v NullableNisNetgroup) IsSet() bool {
	return v.isSet
}

func (v *NullableNisNetgroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNisNetgroup(val *NisNetgroup) *NullableNisNetgroup {
	return &NullableNisNetgroup{value: val, isSet: true}
}

func (v NullableNisNetgroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNisNetgroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


