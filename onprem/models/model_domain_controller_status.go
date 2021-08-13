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

// DomainControllerStatus Connection status of domain controller.
type DomainControllerStatus struct {
	// Specifies the connection status of a domain controller.
	DomainControllerStatus *string `json:"domainControllerStatus,omitempty"`
}

// NewDomainControllerStatus instantiates a new DomainControllerStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainControllerStatus() *DomainControllerStatus {
	this := DomainControllerStatus{}
	return &this
}

// NewDomainControllerStatusWithDefaults instantiates a new DomainControllerStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainControllerStatusWithDefaults() *DomainControllerStatus {
	this := DomainControllerStatus{}
	return &this
}

// GetDomainControllerStatus returns the DomainControllerStatus field value if set, zero value otherwise.
func (o *DomainControllerStatus) GetDomainControllerStatus() string {
	if o == nil || o.DomainControllerStatus == nil {
		var ret string
		return ret
	}
	return *o.DomainControllerStatus
}

// GetDomainControllerStatusOk returns a tuple with the DomainControllerStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainControllerStatus) GetDomainControllerStatusOk() (*string, bool) {
	if o == nil || o.DomainControllerStatus == nil {
		return nil, false
	}
	return o.DomainControllerStatus, true
}

// HasDomainControllerStatus returns a boolean if a field has been set.
func (o *DomainControllerStatus) HasDomainControllerStatus() bool {
	if o != nil && o.DomainControllerStatus != nil {
		return true
	}

	return false
}

// SetDomainControllerStatus gets a reference to the given string and assigns it to the DomainControllerStatus field.
func (o *DomainControllerStatus) SetDomainControllerStatus(v string) {
	o.DomainControllerStatus = &v
}

func (o DomainControllerStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DomainControllerStatus != nil {
		toSerialize["domainControllerStatus"] = o.DomainControllerStatus
	}
	return json.Marshal(toSerialize)
}

type NullableDomainControllerStatus struct {
	value *DomainControllerStatus
	isSet bool
}

func (v NullableDomainControllerStatus) Get() *DomainControllerStatus {
	return v.value
}

func (v *NullableDomainControllerStatus) Set(val *DomainControllerStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainControllerStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainControllerStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainControllerStatus(val *DomainControllerStatus) *NullableDomainControllerStatus {
	return &NullableDomainControllerStatus{value: val, isSet: true}
}

func (v NullableDomainControllerStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainControllerStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o DomainControllerStatus) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}