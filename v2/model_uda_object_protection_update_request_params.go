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

// checks if the UdaObjectProtectionUpdateRequestParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UdaObjectProtectionUpdateRequestParams{}

// UdaObjectProtectionUpdateRequestParams Specifies the update parameters specific to Universal Data Adapter object protection.
type UdaObjectProtectionUpdateRequestParams struct {
	// Specifies the map of custom arguments to be supplied to the various backup scripts.
	BackupJobArguments []KeyValuePair `json:"backupJobArguments,omitempty"`
	// Specifies the maximum number of concurrent IO Streams thatwill be created to exchange data with the cluster. If not specified, the default value is taken as 1.
	Concurrency NullableInt32 `json:"concurrency,omitempty"`
	// Specifies the custom arguments to be supplied to the full backup script when a full backup is enabled in the policy. This field is deprecated. Use backupJobArguments instead.
	FullBackupArgs NullableString `json:"fullBackupArgs,omitempty"`
	// Specifies whether this Protection Group is created from a source having entity support.
	HasEntitySupport NullableBool `json:"hasEntitySupport,omitempty"`
	// Specifies the custom arguments to be supplied to the incremental backup script when an incremental backup is enabled in the policy. This field is deprecated. Use backupJobArguments instead.
	IncrBackupArgs NullableString `json:"incrBackupArgs,omitempty"`
	// Specifies the custom arguments to be supplied to the log backup script when a log backup is enabled in the policy. This field is deprecated. Use backupJobArguments instead.
	LogBackupArgs NullableString `json:"logBackupArgs,omitempty"`
	// Specifies the maximum number of view mounts per host. If not specified, the default value is taken as 1.
	Mounts NullableInt32 `json:"mounts,omitempty"`
	// Specifies the objects to be included in the Object Protection.
	Objects []UdaObjectProtectionObjectParams `json:"objects"`
}

type _UdaObjectProtectionUpdateRequestParams UdaObjectProtectionUpdateRequestParams

// NewUdaObjectProtectionUpdateRequestParams instantiates a new UdaObjectProtectionUpdateRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUdaObjectProtectionUpdateRequestParams(objects []UdaObjectProtectionObjectParams) *UdaObjectProtectionUpdateRequestParams {
	this := UdaObjectProtectionUpdateRequestParams{}
	var concurrency int32 = 1
	this.Concurrency = *NewNullableInt32(&concurrency)
	var mounts int32 = 1
	this.Mounts = *NewNullableInt32(&mounts)
	this.Objects = objects
	return &this
}

// NewUdaObjectProtectionUpdateRequestParamsWithDefaults instantiates a new UdaObjectProtectionUpdateRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUdaObjectProtectionUpdateRequestParamsWithDefaults() *UdaObjectProtectionUpdateRequestParams {
	this := UdaObjectProtectionUpdateRequestParams{}
	var concurrency int32 = 1
	this.Concurrency = *NewNullableInt32(&concurrency)
	var mounts int32 = 1
	this.Mounts = *NewNullableInt32(&mounts)
	return &this
}

// GetBackupJobArguments returns the BackupJobArguments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UdaObjectProtectionUpdateRequestParams) GetBackupJobArguments() []KeyValuePair {
	if o == nil {
		var ret []KeyValuePair
		return ret
	}
	return o.BackupJobArguments
}

// GetBackupJobArgumentsOk returns a tuple with the BackupJobArguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UdaObjectProtectionUpdateRequestParams) GetBackupJobArgumentsOk() ([]KeyValuePair, bool) {
	if o == nil || IsNil(o.BackupJobArguments) {
		return nil, false
	}
	return o.BackupJobArguments, true
}

// HasBackupJobArguments returns a boolean if a field has been set.
func (o *UdaObjectProtectionUpdateRequestParams) HasBackupJobArguments() bool {
	if o != nil && !IsNil(o.BackupJobArguments) {
		return true
	}

	return false
}

// SetBackupJobArguments gets a reference to the given []KeyValuePair and assigns it to the BackupJobArguments field.
func (o *UdaObjectProtectionUpdateRequestParams) SetBackupJobArguments(v []KeyValuePair) {
	o.BackupJobArguments = v
}

// GetConcurrency returns the Concurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UdaObjectProtectionUpdateRequestParams) GetConcurrency() int32 {
	if o == nil || IsNil(o.Concurrency.Get()) {
		var ret int32
		return ret
	}
	return *o.Concurrency.Get()
}

// GetConcurrencyOk returns a tuple with the Concurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UdaObjectProtectionUpdateRequestParams) GetConcurrencyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Concurrency.Get(), o.Concurrency.IsSet()
}

// HasConcurrency returns a boolean if a field has been set.
func (o *UdaObjectProtectionUpdateRequestParams) HasConcurrency() bool {
	if o != nil && o.Concurrency.IsSet() {
		return true
	}

	return false
}

// SetConcurrency gets a reference to the given NullableInt32 and assigns it to the Concurrency field.
func (o *UdaObjectProtectionUpdateRequestParams) SetConcurrency(v int32) {
	o.Concurrency.Set(&v)
}
// SetConcurrencyNil sets the value for Concurrency to be an explicit nil
func (o *UdaObjectProtectionUpdateRequestParams) SetConcurrencyNil() {
	o.Concurrency.Set(nil)
}

// UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
func (o *UdaObjectProtectionUpdateRequestParams) UnsetConcurrency() {
	o.Concurrency.Unset()
}

// GetFullBackupArgs returns the FullBackupArgs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UdaObjectProtectionUpdateRequestParams) GetFullBackupArgs() string {
	if o == nil || IsNil(o.FullBackupArgs.Get()) {
		var ret string
		return ret
	}
	return *o.FullBackupArgs.Get()
}

// GetFullBackupArgsOk returns a tuple with the FullBackupArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UdaObjectProtectionUpdateRequestParams) GetFullBackupArgsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FullBackupArgs.Get(), o.FullBackupArgs.IsSet()
}

// HasFullBackupArgs returns a boolean if a field has been set.
func (o *UdaObjectProtectionUpdateRequestParams) HasFullBackupArgs() bool {
	if o != nil && o.FullBackupArgs.IsSet() {
		return true
	}

	return false
}

// SetFullBackupArgs gets a reference to the given NullableString and assigns it to the FullBackupArgs field.
func (o *UdaObjectProtectionUpdateRequestParams) SetFullBackupArgs(v string) {
	o.FullBackupArgs.Set(&v)
}
// SetFullBackupArgsNil sets the value for FullBackupArgs to be an explicit nil
func (o *UdaObjectProtectionUpdateRequestParams) SetFullBackupArgsNil() {
	o.FullBackupArgs.Set(nil)
}

// UnsetFullBackupArgs ensures that no value is present for FullBackupArgs, not even an explicit nil
func (o *UdaObjectProtectionUpdateRequestParams) UnsetFullBackupArgs() {
	o.FullBackupArgs.Unset()
}

// GetHasEntitySupport returns the HasEntitySupport field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UdaObjectProtectionUpdateRequestParams) GetHasEntitySupport() bool {
	if o == nil || IsNil(o.HasEntitySupport.Get()) {
		var ret bool
		return ret
	}
	return *o.HasEntitySupport.Get()
}

// GetHasEntitySupportOk returns a tuple with the HasEntitySupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UdaObjectProtectionUpdateRequestParams) GetHasEntitySupportOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasEntitySupport.Get(), o.HasEntitySupport.IsSet()
}

// HasHasEntitySupport returns a boolean if a field has been set.
func (o *UdaObjectProtectionUpdateRequestParams) HasHasEntitySupport() bool {
	if o != nil && o.HasEntitySupport.IsSet() {
		return true
	}

	return false
}

// SetHasEntitySupport gets a reference to the given NullableBool and assigns it to the HasEntitySupport field.
func (o *UdaObjectProtectionUpdateRequestParams) SetHasEntitySupport(v bool) {
	o.HasEntitySupport.Set(&v)
}
// SetHasEntitySupportNil sets the value for HasEntitySupport to be an explicit nil
func (o *UdaObjectProtectionUpdateRequestParams) SetHasEntitySupportNil() {
	o.HasEntitySupport.Set(nil)
}

// UnsetHasEntitySupport ensures that no value is present for HasEntitySupport, not even an explicit nil
func (o *UdaObjectProtectionUpdateRequestParams) UnsetHasEntitySupport() {
	o.HasEntitySupport.Unset()
}

// GetIncrBackupArgs returns the IncrBackupArgs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UdaObjectProtectionUpdateRequestParams) GetIncrBackupArgs() string {
	if o == nil || IsNil(o.IncrBackupArgs.Get()) {
		var ret string
		return ret
	}
	return *o.IncrBackupArgs.Get()
}

// GetIncrBackupArgsOk returns a tuple with the IncrBackupArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UdaObjectProtectionUpdateRequestParams) GetIncrBackupArgsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncrBackupArgs.Get(), o.IncrBackupArgs.IsSet()
}

// HasIncrBackupArgs returns a boolean if a field has been set.
func (o *UdaObjectProtectionUpdateRequestParams) HasIncrBackupArgs() bool {
	if o != nil && o.IncrBackupArgs.IsSet() {
		return true
	}

	return false
}

// SetIncrBackupArgs gets a reference to the given NullableString and assigns it to the IncrBackupArgs field.
func (o *UdaObjectProtectionUpdateRequestParams) SetIncrBackupArgs(v string) {
	o.IncrBackupArgs.Set(&v)
}
// SetIncrBackupArgsNil sets the value for IncrBackupArgs to be an explicit nil
func (o *UdaObjectProtectionUpdateRequestParams) SetIncrBackupArgsNil() {
	o.IncrBackupArgs.Set(nil)
}

// UnsetIncrBackupArgs ensures that no value is present for IncrBackupArgs, not even an explicit nil
func (o *UdaObjectProtectionUpdateRequestParams) UnsetIncrBackupArgs() {
	o.IncrBackupArgs.Unset()
}

// GetLogBackupArgs returns the LogBackupArgs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UdaObjectProtectionUpdateRequestParams) GetLogBackupArgs() string {
	if o == nil || IsNil(o.LogBackupArgs.Get()) {
		var ret string
		return ret
	}
	return *o.LogBackupArgs.Get()
}

// GetLogBackupArgsOk returns a tuple with the LogBackupArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UdaObjectProtectionUpdateRequestParams) GetLogBackupArgsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogBackupArgs.Get(), o.LogBackupArgs.IsSet()
}

// HasLogBackupArgs returns a boolean if a field has been set.
func (o *UdaObjectProtectionUpdateRequestParams) HasLogBackupArgs() bool {
	if o != nil && o.LogBackupArgs.IsSet() {
		return true
	}

	return false
}

// SetLogBackupArgs gets a reference to the given NullableString and assigns it to the LogBackupArgs field.
func (o *UdaObjectProtectionUpdateRequestParams) SetLogBackupArgs(v string) {
	o.LogBackupArgs.Set(&v)
}
// SetLogBackupArgsNil sets the value for LogBackupArgs to be an explicit nil
func (o *UdaObjectProtectionUpdateRequestParams) SetLogBackupArgsNil() {
	o.LogBackupArgs.Set(nil)
}

// UnsetLogBackupArgs ensures that no value is present for LogBackupArgs, not even an explicit nil
func (o *UdaObjectProtectionUpdateRequestParams) UnsetLogBackupArgs() {
	o.LogBackupArgs.Unset()
}

// GetMounts returns the Mounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UdaObjectProtectionUpdateRequestParams) GetMounts() int32 {
	if o == nil || IsNil(o.Mounts.Get()) {
		var ret int32
		return ret
	}
	return *o.Mounts.Get()
}

// GetMountsOk returns a tuple with the Mounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UdaObjectProtectionUpdateRequestParams) GetMountsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mounts.Get(), o.Mounts.IsSet()
}

// HasMounts returns a boolean if a field has been set.
func (o *UdaObjectProtectionUpdateRequestParams) HasMounts() bool {
	if o != nil && o.Mounts.IsSet() {
		return true
	}

	return false
}

// SetMounts gets a reference to the given NullableInt32 and assigns it to the Mounts field.
func (o *UdaObjectProtectionUpdateRequestParams) SetMounts(v int32) {
	o.Mounts.Set(&v)
}
// SetMountsNil sets the value for Mounts to be an explicit nil
func (o *UdaObjectProtectionUpdateRequestParams) SetMountsNil() {
	o.Mounts.Set(nil)
}

// UnsetMounts ensures that no value is present for Mounts, not even an explicit nil
func (o *UdaObjectProtectionUpdateRequestParams) UnsetMounts() {
	o.Mounts.Unset()
}

// GetObjects returns the Objects field value
func (o *UdaObjectProtectionUpdateRequestParams) GetObjects() []UdaObjectProtectionObjectParams {
	if o == nil {
		var ret []UdaObjectProtectionObjectParams
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *UdaObjectProtectionUpdateRequestParams) GetObjectsOk() ([]UdaObjectProtectionObjectParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.Objects, true
}

// SetObjects sets field value
func (o *UdaObjectProtectionUpdateRequestParams) SetObjects(v []UdaObjectProtectionObjectParams) {
	o.Objects = v
}

func (o UdaObjectProtectionUpdateRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UdaObjectProtectionUpdateRequestParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BackupJobArguments != nil {
		toSerialize["backupJobArguments"] = o.BackupJobArguments
	}
	if o.Concurrency.IsSet() {
		toSerialize["concurrency"] = o.Concurrency.Get()
	}
	if o.FullBackupArgs.IsSet() {
		toSerialize["fullBackupArgs"] = o.FullBackupArgs.Get()
	}
	if o.HasEntitySupport.IsSet() {
		toSerialize["hasEntitySupport"] = o.HasEntitySupport.Get()
	}
	if o.IncrBackupArgs.IsSet() {
		toSerialize["incrBackupArgs"] = o.IncrBackupArgs.Get()
	}
	if o.LogBackupArgs.IsSet() {
		toSerialize["logBackupArgs"] = o.LogBackupArgs.Get()
	}
	if o.Mounts.IsSet() {
		toSerialize["mounts"] = o.Mounts.Get()
	}
	toSerialize["objects"] = o.Objects
	return toSerialize, nil
}

func (o *UdaObjectProtectionUpdateRequestParams) UnmarshalJSON(data []byte) (err error) {
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

	varUdaObjectProtectionUpdateRequestParams := _UdaObjectProtectionUpdateRequestParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUdaObjectProtectionUpdateRequestParams)

	if err != nil {
		return err
	}

	*o = UdaObjectProtectionUpdateRequestParams(varUdaObjectProtectionUpdateRequestParams)

	return err
}

type NullableUdaObjectProtectionUpdateRequestParams struct {
	value *UdaObjectProtectionUpdateRequestParams
	isSet bool
}

func (v NullableUdaObjectProtectionUpdateRequestParams) Get() *UdaObjectProtectionUpdateRequestParams {
	return v.value
}

func (v *NullableUdaObjectProtectionUpdateRequestParams) Set(val *UdaObjectProtectionUpdateRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUdaObjectProtectionUpdateRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUdaObjectProtectionUpdateRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUdaObjectProtectionUpdateRequestParams(val *UdaObjectProtectionUpdateRequestParams) *NullableUdaObjectProtectionUpdateRequestParams {
	return &NullableUdaObjectProtectionUpdateRequestParams{value: val, isSet: true}
}

func (v NullableUdaObjectProtectionUpdateRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUdaObjectProtectionUpdateRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


