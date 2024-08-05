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

// checks if the ServiceGflags type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceGflags{}

// ServiceGflags Specifies the gflags for a service.
type ServiceGflags struct {
	// Specifies a list of gflags for the service.
	Gflags []Gflag `json:"gflags,omitempty"`
	// Specifies the service name.
	ServiceName NullableString `json:"serviceName,omitempty"`
}

// NewServiceGflags instantiates a new ServiceGflags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceGflags() *ServiceGflags {
	this := ServiceGflags{}
	return &this
}

// NewServiceGflagsWithDefaults instantiates a new ServiceGflags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceGflagsWithDefaults() *ServiceGflags {
	this := ServiceGflags{}
	return &this
}

// GetGflags returns the Gflags field value if set, zero value otherwise.
func (o *ServiceGflags) GetGflags() []Gflag {
	if o == nil || IsNil(o.Gflags) {
		var ret []Gflag
		return ret
	}
	return o.Gflags
}

// GetGflagsOk returns a tuple with the Gflags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceGflags) GetGflagsOk() ([]Gflag, bool) {
	if o == nil || IsNil(o.Gflags) {
		return nil, false
	}
	return o.Gflags, true
}

// HasGflags returns a boolean if a field has been set.
func (o *ServiceGflags) HasGflags() bool {
	if o != nil && !IsNil(o.Gflags) {
		return true
	}

	return false
}

// SetGflags gets a reference to the given []Gflag and assigns it to the Gflags field.
func (o *ServiceGflags) SetGflags(v []Gflag) {
	o.Gflags = v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceGflags) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName.Get()) {
		var ret string
		return ret
	}
	return *o.ServiceName.Get()
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceGflags) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceName.Get(), o.ServiceName.IsSet()
}

// HasServiceName returns a boolean if a field has been set.
func (o *ServiceGflags) HasServiceName() bool {
	if o != nil && o.ServiceName.IsSet() {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given NullableString and assigns it to the ServiceName field.
func (o *ServiceGflags) SetServiceName(v string) {
	o.ServiceName.Set(&v)
}
// SetServiceNameNil sets the value for ServiceName to be an explicit nil
func (o *ServiceGflags) SetServiceNameNil() {
	o.ServiceName.Set(nil)
}

// UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil
func (o *ServiceGflags) UnsetServiceName() {
	o.ServiceName.Unset()
}

func (o ServiceGflags) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceGflags) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gflags) {
		toSerialize["gflags"] = o.Gflags
	}
	if o.ServiceName.IsSet() {
		toSerialize["serviceName"] = o.ServiceName.Get()
	}
	return toSerialize, nil
}

type NullableServiceGflags struct {
	value *ServiceGflags
	isSet bool
}

func (v NullableServiceGflags) Get() *ServiceGflags {
	return v.value
}

func (v *NullableServiceGflags) Set(val *ServiceGflags) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceGflags) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceGflags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceGflags(val *ServiceGflags) *NullableServiceGflags {
	return &NullableServiceGflags{value: val, isSet: true}
}

func (v NullableServiceGflags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceGflags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


