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

// IsilonNfsMountPoint Specifies NFS Mount Point exposed by Isilon Protection Source.
type IsilonNfsMountPoint struct {
	// Specifies the Access Zone name.
	AccessZoneName NullableString `json:"accessZoneName,omitempty"`
	// Specifies the description of the NFS mount point.
	Description NullableString `json:"description,omitempty"`
	// Specifies the Id of the NFS export.
	Id NullableInt64 `json:"id,omitempty"`
}

// NewIsilonNfsMountPoint instantiates a new IsilonNfsMountPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIsilonNfsMountPoint() *IsilonNfsMountPoint {
	this := IsilonNfsMountPoint{}
	return &this
}

// NewIsilonNfsMountPointWithDefaults instantiates a new IsilonNfsMountPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIsilonNfsMountPointWithDefaults() *IsilonNfsMountPoint {
	this := IsilonNfsMountPoint{}
	return &this
}

// GetAccessZoneName returns the AccessZoneName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IsilonNfsMountPoint) GetAccessZoneName() string {
	if o == nil || o.AccessZoneName.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccessZoneName.Get()
}

// GetAccessZoneNameOk returns a tuple with the AccessZoneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IsilonNfsMountPoint) GetAccessZoneNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccessZoneName.Get(), o.AccessZoneName.IsSet()
}

// HasAccessZoneName returns a boolean if a field has been set.
func (o *IsilonNfsMountPoint) HasAccessZoneName() bool {
	if o != nil && o.AccessZoneName.IsSet() {
		return true
	}

	return false
}

// SetAccessZoneName gets a reference to the given NullableString and assigns it to the AccessZoneName field.
func (o *IsilonNfsMountPoint) SetAccessZoneName(v string) {
	o.AccessZoneName.Set(&v)
}
// SetAccessZoneNameNil sets the value for AccessZoneName to be an explicit nil
func (o *IsilonNfsMountPoint) SetAccessZoneNameNil() {
	o.AccessZoneName.Set(nil)
}

// UnsetAccessZoneName ensures that no value is present for AccessZoneName, not even an explicit nil
func (o *IsilonNfsMountPoint) UnsetAccessZoneName() {
	o.AccessZoneName.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IsilonNfsMountPoint) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IsilonNfsMountPoint) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *IsilonNfsMountPoint) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *IsilonNfsMountPoint) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *IsilonNfsMountPoint) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *IsilonNfsMountPoint) UnsetDescription() {
	o.Description.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IsilonNfsMountPoint) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IsilonNfsMountPoint) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *IsilonNfsMountPoint) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *IsilonNfsMountPoint) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *IsilonNfsMountPoint) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *IsilonNfsMountPoint) UnsetId() {
	o.Id.Unset()
}

func (o IsilonNfsMountPoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessZoneName.IsSet() {
		toSerialize["accessZoneName"] = o.AccessZoneName.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIsilonNfsMountPoint struct {
	value *IsilonNfsMountPoint
	isSet bool
}

func (v NullableIsilonNfsMountPoint) Get() *IsilonNfsMountPoint {
	return v.value
}

func (v *NullableIsilonNfsMountPoint) Set(val *IsilonNfsMountPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableIsilonNfsMountPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableIsilonNfsMountPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsilonNfsMountPoint(val *IsilonNfsMountPoint) *NullableIsilonNfsMountPoint {
	return &NullableIsilonNfsMountPoint{value: val, isSet: true}
}

func (v NullableIsilonNfsMountPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsilonNfsMountPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


