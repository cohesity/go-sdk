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

// checks if the LockFileParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LockFileParams{}

// LockFileParams Specifies the parameters to lock a file in a view.
type LockFileParams struct {
	// Specifies a definite timestamp in milliseconds for retaining the file or to extend it's expiry timestamp.
	ExpiryTimestampMsecs NullableInt32 `json:"expiryTimestampMsecs"`
	// Specifies the file path that needs to be locked.
	FilePath NullableString `json:"filePath"`
}

type _LockFileParams LockFileParams

// NewLockFileParams instantiates a new LockFileParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLockFileParams(expiryTimestampMsecs NullableInt32, filePath NullableString) *LockFileParams {
	this := LockFileParams{}
	this.ExpiryTimestampMsecs = expiryTimestampMsecs
	this.FilePath = filePath
	return &this
}

// NewLockFileParamsWithDefaults instantiates a new LockFileParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLockFileParamsWithDefaults() *LockFileParams {
	this := LockFileParams{}
	return &this
}

// GetExpiryTimestampMsecs returns the ExpiryTimestampMsecs field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *LockFileParams) GetExpiryTimestampMsecs() int32 {
	if o == nil || o.ExpiryTimestampMsecs.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ExpiryTimestampMsecs.Get()
}

// GetExpiryTimestampMsecsOk returns a tuple with the ExpiryTimestampMsecs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LockFileParams) GetExpiryTimestampMsecsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiryTimestampMsecs.Get(), o.ExpiryTimestampMsecs.IsSet()
}

// SetExpiryTimestampMsecs sets field value
func (o *LockFileParams) SetExpiryTimestampMsecs(v int32) {
	o.ExpiryTimestampMsecs.Set(&v)
}

// GetFilePath returns the FilePath field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LockFileParams) GetFilePath() string {
	if o == nil || o.FilePath.Get() == nil {
		var ret string
		return ret
	}

	return *o.FilePath.Get()
}

// GetFilePathOk returns a tuple with the FilePath field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LockFileParams) GetFilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilePath.Get(), o.FilePath.IsSet()
}

// SetFilePath sets field value
func (o *LockFileParams) SetFilePath(v string) {
	o.FilePath.Set(&v)
}

func (o LockFileParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LockFileParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expiryTimestampMsecs"] = o.ExpiryTimestampMsecs.Get()
	toSerialize["filePath"] = o.FilePath.Get()
	return toSerialize, nil
}

func (o *LockFileParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"expiryTimestampMsecs",
		"filePath",
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

	varLockFileParams := _LockFileParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLockFileParams)

	if err != nil {
		return err
	}

	*o = LockFileParams(varLockFileParams)

	return err
}

type NullableLockFileParams struct {
	value *LockFileParams
	isSet bool
}

func (v NullableLockFileParams) Get() *LockFileParams {
	return v.value
}

func (v *NullableLockFileParams) Set(val *LockFileParams) {
	v.value = val
	v.isSet = true
}

func (v NullableLockFileParams) IsSet() bool {
	return v.isSet
}

func (v *NullableLockFileParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLockFileParams(val *LockFileParams) *NullableLockFileParams {
	return &NullableLockFileParams{value: val, isSet: true}
}

func (v NullableLockFileParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLockFileParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


