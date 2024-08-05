/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SapHanaObjectProtectionResponseParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SapHanaObjectProtectionResponseParams{}

// SapHanaObjectProtectionResponseParams Specifies the response parameters specific to SAP HANA object protection.
type SapHanaObjectProtectionResponseParams struct {
	// Specifies the maximum number of concurrent IO Streams thatwill be created to exchange data with the cluster. If not specified, the default value is taken as 1.
	Concurrency NullableInt32 `json:"concurrency,omitempty"`
	// Specifies the incremental backup delta (incremental/differential)
	Delta NullableString `json:"delta,omitempty"`
	// Specifies the objects to be included in the Object Protection.
	Objects []UdaObjectProtectionObjectParams `json:"objects"`
}

type _SapHanaObjectProtectionResponseParams SapHanaObjectProtectionResponseParams

// NewSapHanaObjectProtectionResponseParams instantiates a new SapHanaObjectProtectionResponseParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSapHanaObjectProtectionResponseParams(objects []UdaObjectProtectionObjectParams) *SapHanaObjectProtectionResponseParams {
	this := SapHanaObjectProtectionResponseParams{}
	var concurrency int32 = 8
	this.Concurrency = *NewNullableInt32(&concurrency)
	var delta string = "incremental"
	this.Delta = *NewNullableString(&delta)
	this.Objects = objects
	return &this
}

// NewSapHanaObjectProtectionResponseParamsWithDefaults instantiates a new SapHanaObjectProtectionResponseParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSapHanaObjectProtectionResponseParamsWithDefaults() *SapHanaObjectProtectionResponseParams {
	this := SapHanaObjectProtectionResponseParams{}
	var concurrency int32 = 8
	this.Concurrency = *NewNullableInt32(&concurrency)
	var delta string = "incremental"
	this.Delta = *NewNullableString(&delta)
	return &this
}

// GetConcurrency returns the Concurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SapHanaObjectProtectionResponseParams) GetConcurrency() int32 {
	if o == nil || IsNil(o.Concurrency.Get()) {
		var ret int32
		return ret
	}
	return *o.Concurrency.Get()
}

// GetConcurrencyOk returns a tuple with the Concurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SapHanaObjectProtectionResponseParams) GetConcurrencyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Concurrency.Get(), o.Concurrency.IsSet()
}

// HasConcurrency returns a boolean if a field has been set.
func (o *SapHanaObjectProtectionResponseParams) HasConcurrency() bool {
	if o != nil && o.Concurrency.IsSet() {
		return true
	}

	return false
}

// SetConcurrency gets a reference to the given NullableInt32 and assigns it to the Concurrency field.
func (o *SapHanaObjectProtectionResponseParams) SetConcurrency(v int32) {
	o.Concurrency.Set(&v)
}
// SetConcurrencyNil sets the value for Concurrency to be an explicit nil
func (o *SapHanaObjectProtectionResponseParams) SetConcurrencyNil() {
	o.Concurrency.Set(nil)
}

// UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
func (o *SapHanaObjectProtectionResponseParams) UnsetConcurrency() {
	o.Concurrency.Unset()
}

// GetDelta returns the Delta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SapHanaObjectProtectionResponseParams) GetDelta() string {
	if o == nil || IsNil(o.Delta.Get()) {
		var ret string
		return ret
	}
	return *o.Delta.Get()
}

// GetDeltaOk returns a tuple with the Delta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SapHanaObjectProtectionResponseParams) GetDeltaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Delta.Get(), o.Delta.IsSet()
}

// HasDelta returns a boolean if a field has been set.
func (o *SapHanaObjectProtectionResponseParams) HasDelta() bool {
	if o != nil && o.Delta.IsSet() {
		return true
	}

	return false
}

// SetDelta gets a reference to the given NullableString and assigns it to the Delta field.
func (o *SapHanaObjectProtectionResponseParams) SetDelta(v string) {
	o.Delta.Set(&v)
}
// SetDeltaNil sets the value for Delta to be an explicit nil
func (o *SapHanaObjectProtectionResponseParams) SetDeltaNil() {
	o.Delta.Set(nil)
}

// UnsetDelta ensures that no value is present for Delta, not even an explicit nil
func (o *SapHanaObjectProtectionResponseParams) UnsetDelta() {
	o.Delta.Unset()
}

// GetObjects returns the Objects field value
func (o *SapHanaObjectProtectionResponseParams) GetObjects() []UdaObjectProtectionObjectParams {
	if o == nil {
		var ret []UdaObjectProtectionObjectParams
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *SapHanaObjectProtectionResponseParams) GetObjectsOk() ([]UdaObjectProtectionObjectParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.Objects, true
}

// SetObjects sets field value
func (o *SapHanaObjectProtectionResponseParams) SetObjects(v []UdaObjectProtectionObjectParams) {
	o.Objects = v
}

func (o SapHanaObjectProtectionResponseParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SapHanaObjectProtectionResponseParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Concurrency.IsSet() {
		toSerialize["concurrency"] = o.Concurrency.Get()
	}
	if o.Delta.IsSet() {
		toSerialize["delta"] = o.Delta.Get()
	}
	toSerialize["objects"] = o.Objects
	return toSerialize, nil
}

func (o *SapHanaObjectProtectionResponseParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objects",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSapHanaObjectProtectionResponseParams := _SapHanaObjectProtectionResponseParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSapHanaObjectProtectionResponseParams)

	if err != nil {
		return err
	}

	*o = SapHanaObjectProtectionResponseParams(varSapHanaObjectProtectionResponseParams)

	return err
}

type NullableSapHanaObjectProtectionResponseParams struct {
	value *SapHanaObjectProtectionResponseParams
	isSet bool
}

func (v NullableSapHanaObjectProtectionResponseParams) Get() *SapHanaObjectProtectionResponseParams {
	return v.value
}

func (v *NullableSapHanaObjectProtectionResponseParams) Set(val *SapHanaObjectProtectionResponseParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSapHanaObjectProtectionResponseParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSapHanaObjectProtectionResponseParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSapHanaObjectProtectionResponseParams(val *SapHanaObjectProtectionResponseParams) *NullableSapHanaObjectProtectionResponseParams {
	return &NullableSapHanaObjectProtectionResponseParams{value: val, isSet: true}
}

func (v NullableSapHanaObjectProtectionResponseParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSapHanaObjectProtectionResponseParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


