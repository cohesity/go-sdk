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

// NetappClusterInfo Specifies information about a NetApp Cluster Protection Source.
type NetappClusterInfo struct {
	// Specifies information about the contact for the NetApp cluster such as a name, phone number, and email address.
	ContactInfo NullableString `json:"contactInfo,omitempty"`
	// Specifies where this NetApp cluster is located. This location identification string is configured by the NetApp system administrator. This field does not contain the NetApp cluster hostname or IP address.
	Location NullableString `json:"location,omitempty"`
	// Specifies the serial number of the NetApp cluster in the format: x-xx-xxxxxx.
	SerialNumber NullableString `json:"serialNumber,omitempty"`
}

// NewNetappClusterInfo instantiates a new NetappClusterInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetappClusterInfo() *NetappClusterInfo {
	this := NetappClusterInfo{}
	return &this
}

// NewNetappClusterInfoWithDefaults instantiates a new NetappClusterInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetappClusterInfoWithDefaults() *NetappClusterInfo {
	this := NetappClusterInfo{}
	return &this
}

// GetContactInfo returns the ContactInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetappClusterInfo) GetContactInfo() string {
	if o == nil || o.ContactInfo.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContactInfo.Get()
}

// GetContactInfoOk returns a tuple with the ContactInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetappClusterInfo) GetContactInfoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContactInfo.Get(), o.ContactInfo.IsSet()
}

// HasContactInfo returns a boolean if a field has been set.
func (o *NetappClusterInfo) HasContactInfo() bool {
	if o != nil && o.ContactInfo.IsSet() {
		return true
	}

	return false
}

// SetContactInfo gets a reference to the given NullableString and assigns it to the ContactInfo field.
func (o *NetappClusterInfo) SetContactInfo(v string) {
	o.ContactInfo.Set(&v)
}
// SetContactInfoNil sets the value for ContactInfo to be an explicit nil
func (o *NetappClusterInfo) SetContactInfoNil() {
	o.ContactInfo.Set(nil)
}

// UnsetContactInfo ensures that no value is present for ContactInfo, not even an explicit nil
func (o *NetappClusterInfo) UnsetContactInfo() {
	o.ContactInfo.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetappClusterInfo) GetLocation() string {
	if o == nil || o.Location.Get() == nil {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetappClusterInfo) GetLocationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *NetappClusterInfo) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *NetappClusterInfo) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *NetappClusterInfo) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *NetappClusterInfo) UnsetLocation() {
	o.Location.Unset()
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetappClusterInfo) GetSerialNumber() string {
	if o == nil || o.SerialNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber.Get()
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetappClusterInfo) GetSerialNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SerialNumber.Get(), o.SerialNumber.IsSet()
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *NetappClusterInfo) HasSerialNumber() bool {
	if o != nil && o.SerialNumber.IsSet() {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given NullableString and assigns it to the SerialNumber field.
func (o *NetappClusterInfo) SetSerialNumber(v string) {
	o.SerialNumber.Set(&v)
}
// SetSerialNumberNil sets the value for SerialNumber to be an explicit nil
func (o *NetappClusterInfo) SetSerialNumberNil() {
	o.SerialNumber.Set(nil)
}

// UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
func (o *NetappClusterInfo) UnsetSerialNumber() {
	o.SerialNumber.Unset()
}

func (o NetappClusterInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContactInfo.IsSet() {
		toSerialize["contactInfo"] = o.ContactInfo.Get()
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.SerialNumber.IsSet() {
		toSerialize["serialNumber"] = o.SerialNumber.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableNetappClusterInfo struct {
	value *NetappClusterInfo
	isSet bool
}

func (v NullableNetappClusterInfo) Get() *NetappClusterInfo {
	return v.value
}

func (v *NullableNetappClusterInfo) Set(val *NetappClusterInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNetappClusterInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNetappClusterInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetappClusterInfo(val *NetappClusterInfo) *NullableNetappClusterInfo {
	return &NullableNetappClusterInfo{value: val, isSet: true}
}

func (v NullableNetappClusterInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetappClusterInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


