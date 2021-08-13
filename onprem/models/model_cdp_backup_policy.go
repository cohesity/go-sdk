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

// CdpBackupPolicy Specifies CDP (Continious Data Protection) backup settings for a Protection Group.
type CdpBackupPolicy struct {
	Retention CdpRetention `json:"retention"`
}

// NewCdpBackupPolicy instantiates a new CdpBackupPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdpBackupPolicy(retention CdpRetention) *CdpBackupPolicy {
	this := CdpBackupPolicy{}
	this.Retention = retention
	return &this
}

// NewCdpBackupPolicyWithDefaults instantiates a new CdpBackupPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdpBackupPolicyWithDefaults() *CdpBackupPolicy {
	this := CdpBackupPolicy{}
	return &this
}

// GetRetention returns the Retention field value
func (o *CdpBackupPolicy) GetRetention() CdpRetention {
	if o == nil {
		var ret CdpRetention
		return ret
	}

	return o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value
// and a boolean to check if the value has been set.
func (o *CdpBackupPolicy) GetRetentionOk() (*CdpRetention, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Retention, true
}

// SetRetention sets field value
func (o *CdpBackupPolicy) SetRetention(v CdpRetention) {
	o.Retention = v
}

func (o CdpBackupPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["retention"] = o.Retention
	}
	return json.Marshal(toSerialize)
}

type NullableCdpBackupPolicy struct {
	value *CdpBackupPolicy
	isSet bool
}

func (v NullableCdpBackupPolicy) Get() *CdpBackupPolicy {
	return v.value
}

func (v *NullableCdpBackupPolicy) Set(val *CdpBackupPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableCdpBackupPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableCdpBackupPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdpBackupPolicy(val *CdpBackupPolicy) *NullableCdpBackupPolicy {
	return &NullableCdpBackupPolicy{value: val, isSet: true}
}

func (v NullableCdpBackupPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdpBackupPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o CdpBackupPolicy) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}