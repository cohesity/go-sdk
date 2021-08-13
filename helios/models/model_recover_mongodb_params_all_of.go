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

// RecoverMongodbParamsAllOf struct for RecoverMongodbParamsAllOf
type RecoverMongodbParamsAllOf struct {
	// Specifies the local snapshot ids of the Objects to be recovered.
	Snapshots []RecoverMongodbSnapshotParams `json:"snapshots"`
	// A suffix that is to be applied to all recovered objects.
	Suffix NullableString `json:"suffix,omitempty"`
}

// NewRecoverMongodbParamsAllOf instantiates a new RecoverMongodbParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverMongodbParamsAllOf(snapshots []RecoverMongodbSnapshotParams) *RecoverMongodbParamsAllOf {
	this := RecoverMongodbParamsAllOf{}
	this.Snapshots = snapshots
	return &this
}

// NewRecoverMongodbParamsAllOfWithDefaults instantiates a new RecoverMongodbParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverMongodbParamsAllOfWithDefaults() *RecoverMongodbParamsAllOf {
	this := RecoverMongodbParamsAllOf{}
	return &this
}

// GetSnapshots returns the Snapshots field value
// If the value is explicit nil, the zero value for []RecoverMongodbSnapshotParams will be returned
func (o *RecoverMongodbParamsAllOf) GetSnapshots() []RecoverMongodbSnapshotParams {
	if o == nil {
		var ret []RecoverMongodbSnapshotParams
		return ret
	}

	return o.Snapshots
}

// GetSnapshotsOk returns a tuple with the Snapshots field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverMongodbParamsAllOf) GetSnapshotsOk() (*[]RecoverMongodbSnapshotParams, bool) {
	if o == nil || o.Snapshots == nil {
		return nil, false
	}
	return &o.Snapshots, true
}

// SetSnapshots sets field value
func (o *RecoverMongodbParamsAllOf) SetSnapshots(v []RecoverMongodbSnapshotParams) {
	o.Snapshots = v
}

// GetSuffix returns the Suffix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverMongodbParamsAllOf) GetSuffix() string {
	if o == nil || o.Suffix.Get() == nil {
		var ret string
		return ret
	}
	return *o.Suffix.Get()
}

// GetSuffixOk returns a tuple with the Suffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverMongodbParamsAllOf) GetSuffixOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Suffix.Get(), o.Suffix.IsSet()
}

// HasSuffix returns a boolean if a field has been set.
func (o *RecoverMongodbParamsAllOf) HasSuffix() bool {
	if o != nil && o.Suffix.IsSet() {
		return true
	}

	return false
}

// SetSuffix gets a reference to the given NullableString and assigns it to the Suffix field.
func (o *RecoverMongodbParamsAllOf) SetSuffix(v string) {
	o.Suffix.Set(&v)
}
// SetSuffixNil sets the value for Suffix to be an explicit nil
func (o *RecoverMongodbParamsAllOf) SetSuffixNil() {
	o.Suffix.Set(nil)
}

// UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
func (o *RecoverMongodbParamsAllOf) UnsetSuffix() {
	o.Suffix.Unset()
}

func (o RecoverMongodbParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Snapshots != nil {
		toSerialize["snapshots"] = o.Snapshots
	}
	if o.Suffix.IsSet() {
		toSerialize["suffix"] = o.Suffix.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverMongodbParamsAllOf struct {
	value *RecoverMongodbParamsAllOf
	isSet bool
}

func (v NullableRecoverMongodbParamsAllOf) Get() *RecoverMongodbParamsAllOf {
	return v.value
}

func (v *NullableRecoverMongodbParamsAllOf) Set(val *RecoverMongodbParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverMongodbParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverMongodbParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverMongodbParamsAllOf(val *RecoverMongodbParamsAllOf) *NullableRecoverMongodbParamsAllOf {
	return &NullableRecoverMongodbParamsAllOf{value: val, isSet: true}
}

func (v NullableRecoverMongodbParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverMongodbParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


