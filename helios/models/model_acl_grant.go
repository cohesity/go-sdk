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

// AclGrant Specifies an ACL grant.
type AclGrant struct {
	// Specifies the grantee.
	Grantee AclGrantee `json:"grantee"`
	// Specifies a list of permissions granted to the grantees.
	Permissions []string `json:"permissions"`
}

// NewAclGrant instantiates a new AclGrant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAclGrant(grantee AclGrantee, permissions []string) *AclGrant {
	this := AclGrant{}
	this.Grantee = grantee
	this.Permissions = permissions
	return &this
}

// NewAclGrantWithDefaults instantiates a new AclGrant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAclGrantWithDefaults() *AclGrant {
	this := AclGrant{}
	return &this
}

// GetGrantee returns the Grantee field value
func (o *AclGrant) GetGrantee() AclGrantee {
	if o == nil {
		var ret AclGrantee
		return ret
	}

	return o.Grantee
}

// GetGranteeOk returns a tuple with the Grantee field value
// and a boolean to check if the value has been set.
func (o *AclGrant) GetGranteeOk() (*AclGrantee, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Grantee, true
}

// SetGrantee sets field value
func (o *AclGrant) SetGrantee(v AclGrantee) {
	o.Grantee = v
}

// GetPermissions returns the Permissions field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *AclGrant) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AclGrant) GetPermissionsOk() (*[]string, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value
func (o *AclGrant) SetPermissions(v []string) {
	o.Permissions = v
}

func (o AclGrant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["grantee"] = o.Grantee
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableAclGrant struct {
	value *AclGrant
	isSet bool
}

func (v NullableAclGrant) Get() *AclGrant {
	return v.value
}

func (v *NullableAclGrant) Set(val *AclGrant) {
	v.value = val
	v.isSet = true
}

func (v NullableAclGrant) IsSet() bool {
	return v.isSet
}

func (v *NullableAclGrant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAclGrant(val *AclGrant) *NullableAclGrant {
	return &NullableAclGrant{value: val, isSet: true}
}

func (v NullableAclGrant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAclGrant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


