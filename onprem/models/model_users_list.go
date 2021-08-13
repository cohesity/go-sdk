/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// UsersList Specifies a list of users.
type UsersList struct {
	// Specifies the list of users.
	Users []UserParams `json:"users,omitempty"`
}

// NewUsersList instantiates a new UsersList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersList() *UsersList {
	this := UsersList{}
	return &this
}

// NewUsersListWithDefaults instantiates a new UsersList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersListWithDefaults() *UsersList {
	this := UsersList{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsersList) GetUsers() []UserParams {
	if o == nil  {
		var ret []UserParams
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsersList) GetUsersOk() (*[]UserParams, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return &o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *UsersList) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []UserParams and assigns it to the Users field.
func (o *UsersList) SetUsers(v []UserParams) {
	o.Users = v
}

func (o UsersList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableUsersList struct {
	value *UsersList
	isSet bool
}

func (v NullableUsersList) Get() *UsersList {
	return v.value
}

func (v *NullableUsersList) Set(val *UsersList) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersList) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersList(val *UsersList) *NullableUsersList {
	return &NullableUsersList{value: val, isSet: true}
}

func (v NullableUsersList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o UsersList) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}