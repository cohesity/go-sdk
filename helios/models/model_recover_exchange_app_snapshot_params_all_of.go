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

// RecoverExchangeAppSnapshotParamsAllOf struct for RecoverExchangeAppSnapshotParamsAllOf
type RecoverExchangeAppSnapshotParamsAllOf struct {
	// Specifies the list of params for all the databases which have to be recovered.
	AppObjects []ExchangeRecoverDatabaseParams `json:"appObjects,omitempty"`
}

// NewRecoverExchangeAppSnapshotParamsAllOf instantiates a new RecoverExchangeAppSnapshotParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverExchangeAppSnapshotParamsAllOf() *RecoverExchangeAppSnapshotParamsAllOf {
	this := RecoverExchangeAppSnapshotParamsAllOf{}
	return &this
}

// NewRecoverExchangeAppSnapshotParamsAllOfWithDefaults instantiates a new RecoverExchangeAppSnapshotParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverExchangeAppSnapshotParamsAllOfWithDefaults() *RecoverExchangeAppSnapshotParamsAllOf {
	this := RecoverExchangeAppSnapshotParamsAllOf{}
	return &this
}

// GetAppObjects returns the AppObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverExchangeAppSnapshotParamsAllOf) GetAppObjects() []ExchangeRecoverDatabaseParams {
	if o == nil  {
		var ret []ExchangeRecoverDatabaseParams
		return ret
	}
	return o.AppObjects
}

// GetAppObjectsOk returns a tuple with the AppObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverExchangeAppSnapshotParamsAllOf) GetAppObjectsOk() (*[]ExchangeRecoverDatabaseParams, bool) {
	if o == nil || o.AppObjects == nil {
		return nil, false
	}
	return &o.AppObjects, true
}

// HasAppObjects returns a boolean if a field has been set.
func (o *RecoverExchangeAppSnapshotParamsAllOf) HasAppObjects() bool {
	if o != nil && o.AppObjects != nil {
		return true
	}

	return false
}

// SetAppObjects gets a reference to the given []ExchangeRecoverDatabaseParams and assigns it to the AppObjects field.
func (o *RecoverExchangeAppSnapshotParamsAllOf) SetAppObjects(v []ExchangeRecoverDatabaseParams) {
	o.AppObjects = v
}

func (o RecoverExchangeAppSnapshotParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppObjects != nil {
		toSerialize["appObjects"] = o.AppObjects
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverExchangeAppSnapshotParamsAllOf struct {
	value *RecoverExchangeAppSnapshotParamsAllOf
	isSet bool
}

func (v NullableRecoverExchangeAppSnapshotParamsAllOf) Get() *RecoverExchangeAppSnapshotParamsAllOf {
	return v.value
}

func (v *NullableRecoverExchangeAppSnapshotParamsAllOf) Set(val *RecoverExchangeAppSnapshotParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverExchangeAppSnapshotParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverExchangeAppSnapshotParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverExchangeAppSnapshotParamsAllOf(val *RecoverExchangeAppSnapshotParamsAllOf) *NullableRecoverExchangeAppSnapshotParamsAllOf {
	return &NullableRecoverExchangeAppSnapshotParamsAllOf{value: val, isSet: true}
}

func (v NullableRecoverExchangeAppSnapshotParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverExchangeAppSnapshotParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


