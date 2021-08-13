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

// CouchbaseIndexedObjectAllOf struct for CouchbaseIndexedObjectAllOf
type CouchbaseIndexedObjectAllOf struct {
	// Specifies the Couchbase Object Type. For Couchbase this is alywas set to Bucket.
	Type NullableString `json:"type,omitempty"`
	// Specifies the id of the indexed object.
	Id NullableString `json:"id,omitempty"`
}

// NewCouchbaseIndexedObjectAllOf instantiates a new CouchbaseIndexedObjectAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouchbaseIndexedObjectAllOf() *CouchbaseIndexedObjectAllOf {
	this := CouchbaseIndexedObjectAllOf{}
	return &this
}

// NewCouchbaseIndexedObjectAllOfWithDefaults instantiates a new CouchbaseIndexedObjectAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouchbaseIndexedObjectAllOfWithDefaults() *CouchbaseIndexedObjectAllOf {
	this := CouchbaseIndexedObjectAllOf{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CouchbaseIndexedObjectAllOf) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CouchbaseIndexedObjectAllOf) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *CouchbaseIndexedObjectAllOf) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *CouchbaseIndexedObjectAllOf) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *CouchbaseIndexedObjectAllOf) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *CouchbaseIndexedObjectAllOf) UnsetType() {
	o.Type.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CouchbaseIndexedObjectAllOf) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CouchbaseIndexedObjectAllOf) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CouchbaseIndexedObjectAllOf) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *CouchbaseIndexedObjectAllOf) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *CouchbaseIndexedObjectAllOf) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CouchbaseIndexedObjectAllOf) UnsetId() {
	o.Id.Unset()
}

func (o CouchbaseIndexedObjectAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCouchbaseIndexedObjectAllOf struct {
	value *CouchbaseIndexedObjectAllOf
	isSet bool
}

func (v NullableCouchbaseIndexedObjectAllOf) Get() *CouchbaseIndexedObjectAllOf {
	return v.value
}

func (v *NullableCouchbaseIndexedObjectAllOf) Set(val *CouchbaseIndexedObjectAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCouchbaseIndexedObjectAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCouchbaseIndexedObjectAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouchbaseIndexedObjectAllOf(val *CouchbaseIndexedObjectAllOf) *NullableCouchbaseIndexedObjectAllOf {
	return &NullableCouchbaseIndexedObjectAllOf{value: val, isSet: true}
}

func (v NullableCouchbaseIndexedObjectAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouchbaseIndexedObjectAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


