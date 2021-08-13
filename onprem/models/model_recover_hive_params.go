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

// RecoverHiveParams Specifies the type of recover action to be performed.
type RecoverHiveParams struct {
	// Specifies the 'Source Registration ID' of the source where the objects are to be recovered. If this is not specified, the recovery job will recover to the original location.
	RecoverTo NullableInt64 `json:"recoverTo,omitempty"`
	// Set to true to overwrite an existing object at the destination. If set to false, and the same object exists at the destination, then recovery will fail for that object.
	Overwrite NullableBool `json:"overwrite,omitempty"`
	// Specifies the maximum number of concurrent IO Streams that will be created to exchange data with the cluster.
	Concurrency NullableInt32 `json:"concurrency,omitempty"`
	// Specifies the maximum network bandwidth that each concurrent IO Stream can use for exchanging data with the cluster.
	BandwidthMBPS NullableInt64 `json:"bandwidthMBPS,omitempty"`
	// This field will hold the warnings in cases where the job status is SucceededWithWarnings.
	Warnings []string `json:"warnings,omitempty"`
	// Specifies the local snapshot ids of the Objects to be recovered.
	Snapshots []RecoverHiveSnapshotParams `json:"snapshots"`
	// A suffix that is to be applied to all recovered objects.
	Suffix NullableString `json:"suffix,omitempty"`
}

// NewRecoverHiveParams instantiates a new RecoverHiveParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverHiveParams(snapshots []RecoverHiveSnapshotParams) *RecoverHiveParams {
	this := RecoverHiveParams{}
	this.Snapshots = snapshots
	return &this
}

// NewRecoverHiveParamsWithDefaults instantiates a new RecoverHiveParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverHiveParamsWithDefaults() *RecoverHiveParams {
	this := RecoverHiveParams{}
	return &this
}

// GetRecoverTo returns the RecoverTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHiveParams) GetRecoverTo() int64 {
	if o == nil || o.RecoverTo.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RecoverTo.Get()
}

// GetRecoverToOk returns a tuple with the RecoverTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHiveParams) GetRecoverToOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverTo.Get(), o.RecoverTo.IsSet()
}

// HasRecoverTo returns a boolean if a field has been set.
func (o *RecoverHiveParams) HasRecoverTo() bool {
	if o != nil && o.RecoverTo.IsSet() {
		return true
	}

	return false
}

// SetRecoverTo gets a reference to the given NullableInt64 and assigns it to the RecoverTo field.
func (o *RecoverHiveParams) SetRecoverTo(v int64) {
	o.RecoverTo.Set(&v)
}
// SetRecoverToNil sets the value for RecoverTo to be an explicit nil
func (o *RecoverHiveParams) SetRecoverToNil() {
	o.RecoverTo.Set(nil)
}

// UnsetRecoverTo ensures that no value is present for RecoverTo, not even an explicit nil
func (o *RecoverHiveParams) UnsetRecoverTo() {
	o.RecoverTo.Unset()
}

// GetOverwrite returns the Overwrite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHiveParams) GetOverwrite() bool {
	if o == nil || o.Overwrite.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Overwrite.Get()
}

// GetOverwriteOk returns a tuple with the Overwrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHiveParams) GetOverwriteOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Overwrite.Get(), o.Overwrite.IsSet()
}

// HasOverwrite returns a boolean if a field has been set.
func (o *RecoverHiveParams) HasOverwrite() bool {
	if o != nil && o.Overwrite.IsSet() {
		return true
	}

	return false
}

// SetOverwrite gets a reference to the given NullableBool and assigns it to the Overwrite field.
func (o *RecoverHiveParams) SetOverwrite(v bool) {
	o.Overwrite.Set(&v)
}
// SetOverwriteNil sets the value for Overwrite to be an explicit nil
func (o *RecoverHiveParams) SetOverwriteNil() {
	o.Overwrite.Set(nil)
}

// UnsetOverwrite ensures that no value is present for Overwrite, not even an explicit nil
func (o *RecoverHiveParams) UnsetOverwrite() {
	o.Overwrite.Unset()
}

// GetConcurrency returns the Concurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHiveParams) GetConcurrency() int32 {
	if o == nil || o.Concurrency.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Concurrency.Get()
}

// GetConcurrencyOk returns a tuple with the Concurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHiveParams) GetConcurrencyOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Concurrency.Get(), o.Concurrency.IsSet()
}

// HasConcurrency returns a boolean if a field has been set.
func (o *RecoverHiveParams) HasConcurrency() bool {
	if o != nil && o.Concurrency.IsSet() {
		return true
	}

	return false
}

// SetConcurrency gets a reference to the given NullableInt32 and assigns it to the Concurrency field.
func (o *RecoverHiveParams) SetConcurrency(v int32) {
	o.Concurrency.Set(&v)
}
// SetConcurrencyNil sets the value for Concurrency to be an explicit nil
func (o *RecoverHiveParams) SetConcurrencyNil() {
	o.Concurrency.Set(nil)
}

// UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
func (o *RecoverHiveParams) UnsetConcurrency() {
	o.Concurrency.Unset()
}

// GetBandwidthMBPS returns the BandwidthMBPS field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHiveParams) GetBandwidthMBPS() int64 {
	if o == nil || o.BandwidthMBPS.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BandwidthMBPS.Get()
}

// GetBandwidthMBPSOk returns a tuple with the BandwidthMBPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHiveParams) GetBandwidthMBPSOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BandwidthMBPS.Get(), o.BandwidthMBPS.IsSet()
}

// HasBandwidthMBPS returns a boolean if a field has been set.
func (o *RecoverHiveParams) HasBandwidthMBPS() bool {
	if o != nil && o.BandwidthMBPS.IsSet() {
		return true
	}

	return false
}

// SetBandwidthMBPS gets a reference to the given NullableInt64 and assigns it to the BandwidthMBPS field.
func (o *RecoverHiveParams) SetBandwidthMBPS(v int64) {
	o.BandwidthMBPS.Set(&v)
}
// SetBandwidthMBPSNil sets the value for BandwidthMBPS to be an explicit nil
func (o *RecoverHiveParams) SetBandwidthMBPSNil() {
	o.BandwidthMBPS.Set(nil)
}

// UnsetBandwidthMBPS ensures that no value is present for BandwidthMBPS, not even an explicit nil
func (o *RecoverHiveParams) UnsetBandwidthMBPS() {
	o.BandwidthMBPS.Unset()
}

// GetWarnings returns the Warnings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHiveParams) GetWarnings() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHiveParams) GetWarningsOk() (*[]string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *RecoverHiveParams) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *RecoverHiveParams) SetWarnings(v []string) {
	o.Warnings = v
}

// GetSnapshots returns the Snapshots field value
// If the value is explicit nil, the zero value for []RecoverHiveSnapshotParams will be returned
func (o *RecoverHiveParams) GetSnapshots() []RecoverHiveSnapshotParams {
	if o == nil {
		var ret []RecoverHiveSnapshotParams
		return ret
	}

	return o.Snapshots
}

// GetSnapshotsOk returns a tuple with the Snapshots field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHiveParams) GetSnapshotsOk() (*[]RecoverHiveSnapshotParams, bool) {
	if o == nil || o.Snapshots == nil {
		return nil, false
	}
	return &o.Snapshots, true
}

// SetSnapshots sets field value
func (o *RecoverHiveParams) SetSnapshots(v []RecoverHiveSnapshotParams) {
	o.Snapshots = v
}

// GetSuffix returns the Suffix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHiveParams) GetSuffix() string {
	if o == nil || o.Suffix.Get() == nil {
		var ret string
		return ret
	}
	return *o.Suffix.Get()
}

// GetSuffixOk returns a tuple with the Suffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHiveParams) GetSuffixOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Suffix.Get(), o.Suffix.IsSet()
}

// HasSuffix returns a boolean if a field has been set.
func (o *RecoverHiveParams) HasSuffix() bool {
	if o != nil && o.Suffix.IsSet() {
		return true
	}

	return false
}

// SetSuffix gets a reference to the given NullableString and assigns it to the Suffix field.
func (o *RecoverHiveParams) SetSuffix(v string) {
	o.Suffix.Set(&v)
}
// SetSuffixNil sets the value for Suffix to be an explicit nil
func (o *RecoverHiveParams) SetSuffixNil() {
	o.Suffix.Set(nil)
}

// UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
func (o *RecoverHiveParams) UnsetSuffix() {
	o.Suffix.Unset()
}

func (o RecoverHiveParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecoverTo.IsSet() {
		toSerialize["recoverTo"] = o.RecoverTo.Get()
	}
	if o.Overwrite.IsSet() {
		toSerialize["overwrite"] = o.Overwrite.Get()
	}
	if o.Concurrency.IsSet() {
		toSerialize["concurrency"] = o.Concurrency.Get()
	}
	if o.BandwidthMBPS.IsSet() {
		toSerialize["bandwidthMBPS"] = o.BandwidthMBPS.Get()
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	if o.Snapshots != nil {
		toSerialize["snapshots"] = o.Snapshots
	}
	if o.Suffix.IsSet() {
		toSerialize["suffix"] = o.Suffix.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverHiveParams struct {
	value *RecoverHiveParams
	isSet bool
}

func (v NullableRecoverHiveParams) Get() *RecoverHiveParams {
	return v.value
}

func (v *NullableRecoverHiveParams) Set(val *RecoverHiveParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverHiveParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverHiveParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverHiveParams(val *RecoverHiveParams) *NullableRecoverHiveParams {
	return &NullableRecoverHiveParams{value: val, isSet: true}
}

func (v NullableRecoverHiveParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverHiveParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RecoverHiveParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}