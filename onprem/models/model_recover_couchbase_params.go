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

// RecoverCouchbaseParams Specifies the parameters to recover Couchbase objects.
type RecoverCouchbaseParams struct {
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
	Snapshots []RecoverCouchbaseSnapshotParams `json:"snapshots"`
	// A suffix that is to be applied to all recovered objects.
	Suffix NullableString `json:"suffix,omitempty"`
	FilterDocumentsParams FilterDocumentsParams `json:"filterDocumentsParams"`
	// If set to true, docuements from the bucket being recovered will be appended into the bucket at the destination.
	AppendDocuments NullableBool `json:"appendDocuments,omitempty"`
	// Set to true to recover only the bucket configurations. No documents will be recovered.
	DdlOnlyRecovery NullableBool `json:"ddlOnlyRecovery,omitempty"`
	// If set to true existing users will be replaced with users from the bucket being recovered.
	OverwriteUsers NullableBool `json:"overwriteUsers,omitempty"`
}

// NewRecoverCouchbaseParams instantiates a new RecoverCouchbaseParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverCouchbaseParams(snapshots []RecoverCouchbaseSnapshotParams, filterDocumentsParams FilterDocumentsParams) *RecoverCouchbaseParams {
	this := RecoverCouchbaseParams{}
	this.Snapshots = snapshots
	this.FilterDocumentsParams = filterDocumentsParams
	return &this
}

// NewRecoverCouchbaseParamsWithDefaults instantiates a new RecoverCouchbaseParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverCouchbaseParamsWithDefaults() *RecoverCouchbaseParams {
	this := RecoverCouchbaseParams{}
	return &this
}

// GetRecoverTo returns the RecoverTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverCouchbaseParams) GetRecoverTo() int64 {
	if o == nil || o.RecoverTo.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RecoverTo.Get()
}

// GetRecoverToOk returns a tuple with the RecoverTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverCouchbaseParams) GetRecoverToOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverTo.Get(), o.RecoverTo.IsSet()
}

// HasRecoverTo returns a boolean if a field has been set.
func (o *RecoverCouchbaseParams) HasRecoverTo() bool {
	if o != nil && o.RecoverTo.IsSet() {
		return true
	}

	return false
}

// SetRecoverTo gets a reference to the given NullableInt64 and assigns it to the RecoverTo field.
func (o *RecoverCouchbaseParams) SetRecoverTo(v int64) {
	o.RecoverTo.Set(&v)
}
// SetRecoverToNil sets the value for RecoverTo to be an explicit nil
func (o *RecoverCouchbaseParams) SetRecoverToNil() {
	o.RecoverTo.Set(nil)
}

// UnsetRecoverTo ensures that no value is present for RecoverTo, not even an explicit nil
func (o *RecoverCouchbaseParams) UnsetRecoverTo() {
	o.RecoverTo.Unset()
}

// GetOverwrite returns the Overwrite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverCouchbaseParams) GetOverwrite() bool {
	if o == nil || o.Overwrite.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Overwrite.Get()
}

// GetOverwriteOk returns a tuple with the Overwrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverCouchbaseParams) GetOverwriteOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Overwrite.Get(), o.Overwrite.IsSet()
}

// HasOverwrite returns a boolean if a field has been set.
func (o *RecoverCouchbaseParams) HasOverwrite() bool {
	if o != nil && o.Overwrite.IsSet() {
		return true
	}

	return false
}

// SetOverwrite gets a reference to the given NullableBool and assigns it to the Overwrite field.
func (o *RecoverCouchbaseParams) SetOverwrite(v bool) {
	o.Overwrite.Set(&v)
}
// SetOverwriteNil sets the value for Overwrite to be an explicit nil
func (o *RecoverCouchbaseParams) SetOverwriteNil() {
	o.Overwrite.Set(nil)
}

// UnsetOverwrite ensures that no value is present for Overwrite, not even an explicit nil
func (o *RecoverCouchbaseParams) UnsetOverwrite() {
	o.Overwrite.Unset()
}

// GetConcurrency returns the Concurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverCouchbaseParams) GetConcurrency() int32 {
	if o == nil || o.Concurrency.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Concurrency.Get()
}

// GetConcurrencyOk returns a tuple with the Concurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverCouchbaseParams) GetConcurrencyOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Concurrency.Get(), o.Concurrency.IsSet()
}

// HasConcurrency returns a boolean if a field has been set.
func (o *RecoverCouchbaseParams) HasConcurrency() bool {
	if o != nil && o.Concurrency.IsSet() {
		return true
	}

	return false
}

// SetConcurrency gets a reference to the given NullableInt32 and assigns it to the Concurrency field.
func (o *RecoverCouchbaseParams) SetConcurrency(v int32) {
	o.Concurrency.Set(&v)
}
// SetConcurrencyNil sets the value for Concurrency to be an explicit nil
func (o *RecoverCouchbaseParams) SetConcurrencyNil() {
	o.Concurrency.Set(nil)
}

// UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
func (o *RecoverCouchbaseParams) UnsetConcurrency() {
	o.Concurrency.Unset()
}

// GetBandwidthMBPS returns the BandwidthMBPS field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverCouchbaseParams) GetBandwidthMBPS() int64 {
	if o == nil || o.BandwidthMBPS.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BandwidthMBPS.Get()
}

// GetBandwidthMBPSOk returns a tuple with the BandwidthMBPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverCouchbaseParams) GetBandwidthMBPSOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BandwidthMBPS.Get(), o.BandwidthMBPS.IsSet()
}

// HasBandwidthMBPS returns a boolean if a field has been set.
func (o *RecoverCouchbaseParams) HasBandwidthMBPS() bool {
	if o != nil && o.BandwidthMBPS.IsSet() {
		return true
	}

	return false
}

// SetBandwidthMBPS gets a reference to the given NullableInt64 and assigns it to the BandwidthMBPS field.
func (o *RecoverCouchbaseParams) SetBandwidthMBPS(v int64) {
	o.BandwidthMBPS.Set(&v)
}
// SetBandwidthMBPSNil sets the value for BandwidthMBPS to be an explicit nil
func (o *RecoverCouchbaseParams) SetBandwidthMBPSNil() {
	o.BandwidthMBPS.Set(nil)
}

// UnsetBandwidthMBPS ensures that no value is present for BandwidthMBPS, not even an explicit nil
func (o *RecoverCouchbaseParams) UnsetBandwidthMBPS() {
	o.BandwidthMBPS.Unset()
}

// GetWarnings returns the Warnings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverCouchbaseParams) GetWarnings() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverCouchbaseParams) GetWarningsOk() (*[]string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *RecoverCouchbaseParams) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *RecoverCouchbaseParams) SetWarnings(v []string) {
	o.Warnings = v
}

// GetSnapshots returns the Snapshots field value
// If the value is explicit nil, the zero value for []RecoverCouchbaseSnapshotParams will be returned
func (o *RecoverCouchbaseParams) GetSnapshots() []RecoverCouchbaseSnapshotParams {
	if o == nil {
		var ret []RecoverCouchbaseSnapshotParams
		return ret
	}

	return o.Snapshots
}

// GetSnapshotsOk returns a tuple with the Snapshots field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverCouchbaseParams) GetSnapshotsOk() (*[]RecoverCouchbaseSnapshotParams, bool) {
	if o == nil || o.Snapshots == nil {
		return nil, false
	}
	return &o.Snapshots, true
}

// SetSnapshots sets field value
func (o *RecoverCouchbaseParams) SetSnapshots(v []RecoverCouchbaseSnapshotParams) {
	o.Snapshots = v
}

// GetSuffix returns the Suffix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverCouchbaseParams) GetSuffix() string {
	if o == nil || o.Suffix.Get() == nil {
		var ret string
		return ret
	}
	return *o.Suffix.Get()
}

// GetSuffixOk returns a tuple with the Suffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverCouchbaseParams) GetSuffixOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Suffix.Get(), o.Suffix.IsSet()
}

// HasSuffix returns a boolean if a field has been set.
func (o *RecoverCouchbaseParams) HasSuffix() bool {
	if o != nil && o.Suffix.IsSet() {
		return true
	}

	return false
}

// SetSuffix gets a reference to the given NullableString and assigns it to the Suffix field.
func (o *RecoverCouchbaseParams) SetSuffix(v string) {
	o.Suffix.Set(&v)
}
// SetSuffixNil sets the value for Suffix to be an explicit nil
func (o *RecoverCouchbaseParams) SetSuffixNil() {
	o.Suffix.Set(nil)
}

// UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
func (o *RecoverCouchbaseParams) UnsetSuffix() {
	o.Suffix.Unset()
}

// GetFilterDocumentsParams returns the FilterDocumentsParams field value
func (o *RecoverCouchbaseParams) GetFilterDocumentsParams() FilterDocumentsParams {
	if o == nil {
		var ret FilterDocumentsParams
		return ret
	}

	return o.FilterDocumentsParams
}

// GetFilterDocumentsParamsOk returns a tuple with the FilterDocumentsParams field value
// and a boolean to check if the value has been set.
func (o *RecoverCouchbaseParams) GetFilterDocumentsParamsOk() (*FilterDocumentsParams, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FilterDocumentsParams, true
}

// SetFilterDocumentsParams sets field value
func (o *RecoverCouchbaseParams) SetFilterDocumentsParams(v FilterDocumentsParams) {
	o.FilterDocumentsParams = v
}

// GetAppendDocuments returns the AppendDocuments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverCouchbaseParams) GetAppendDocuments() bool {
	if o == nil || o.AppendDocuments.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AppendDocuments.Get()
}

// GetAppendDocumentsOk returns a tuple with the AppendDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverCouchbaseParams) GetAppendDocumentsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AppendDocuments.Get(), o.AppendDocuments.IsSet()
}

// HasAppendDocuments returns a boolean if a field has been set.
func (o *RecoverCouchbaseParams) HasAppendDocuments() bool {
	if o != nil && o.AppendDocuments.IsSet() {
		return true
	}

	return false
}

// SetAppendDocuments gets a reference to the given NullableBool and assigns it to the AppendDocuments field.
func (o *RecoverCouchbaseParams) SetAppendDocuments(v bool) {
	o.AppendDocuments.Set(&v)
}
// SetAppendDocumentsNil sets the value for AppendDocuments to be an explicit nil
func (o *RecoverCouchbaseParams) SetAppendDocumentsNil() {
	o.AppendDocuments.Set(nil)
}

// UnsetAppendDocuments ensures that no value is present for AppendDocuments, not even an explicit nil
func (o *RecoverCouchbaseParams) UnsetAppendDocuments() {
	o.AppendDocuments.Unset()
}

// GetDdlOnlyRecovery returns the DdlOnlyRecovery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverCouchbaseParams) GetDdlOnlyRecovery() bool {
	if o == nil || o.DdlOnlyRecovery.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DdlOnlyRecovery.Get()
}

// GetDdlOnlyRecoveryOk returns a tuple with the DdlOnlyRecovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverCouchbaseParams) GetDdlOnlyRecoveryOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DdlOnlyRecovery.Get(), o.DdlOnlyRecovery.IsSet()
}

// HasDdlOnlyRecovery returns a boolean if a field has been set.
func (o *RecoverCouchbaseParams) HasDdlOnlyRecovery() bool {
	if o != nil && o.DdlOnlyRecovery.IsSet() {
		return true
	}

	return false
}

// SetDdlOnlyRecovery gets a reference to the given NullableBool and assigns it to the DdlOnlyRecovery field.
func (o *RecoverCouchbaseParams) SetDdlOnlyRecovery(v bool) {
	o.DdlOnlyRecovery.Set(&v)
}
// SetDdlOnlyRecoveryNil sets the value for DdlOnlyRecovery to be an explicit nil
func (o *RecoverCouchbaseParams) SetDdlOnlyRecoveryNil() {
	o.DdlOnlyRecovery.Set(nil)
}

// UnsetDdlOnlyRecovery ensures that no value is present for DdlOnlyRecovery, not even an explicit nil
func (o *RecoverCouchbaseParams) UnsetDdlOnlyRecovery() {
	o.DdlOnlyRecovery.Unset()
}

// GetOverwriteUsers returns the OverwriteUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverCouchbaseParams) GetOverwriteUsers() bool {
	if o == nil || o.OverwriteUsers.Get() == nil {
		var ret bool
		return ret
	}
	return *o.OverwriteUsers.Get()
}

// GetOverwriteUsersOk returns a tuple with the OverwriteUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverCouchbaseParams) GetOverwriteUsersOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OverwriteUsers.Get(), o.OverwriteUsers.IsSet()
}

// HasOverwriteUsers returns a boolean if a field has been set.
func (o *RecoverCouchbaseParams) HasOverwriteUsers() bool {
	if o != nil && o.OverwriteUsers.IsSet() {
		return true
	}

	return false
}

// SetOverwriteUsers gets a reference to the given NullableBool and assigns it to the OverwriteUsers field.
func (o *RecoverCouchbaseParams) SetOverwriteUsers(v bool) {
	o.OverwriteUsers.Set(&v)
}
// SetOverwriteUsersNil sets the value for OverwriteUsers to be an explicit nil
func (o *RecoverCouchbaseParams) SetOverwriteUsersNil() {
	o.OverwriteUsers.Set(nil)
}

// UnsetOverwriteUsers ensures that no value is present for OverwriteUsers, not even an explicit nil
func (o *RecoverCouchbaseParams) UnsetOverwriteUsers() {
	o.OverwriteUsers.Unset()
}

func (o RecoverCouchbaseParams) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["filterDocumentsParams"] = o.FilterDocumentsParams
	}
	if o.AppendDocuments.IsSet() {
		toSerialize["appendDocuments"] = o.AppendDocuments.Get()
	}
	if o.DdlOnlyRecovery.IsSet() {
		toSerialize["ddlOnlyRecovery"] = o.DdlOnlyRecovery.Get()
	}
	if o.OverwriteUsers.IsSet() {
		toSerialize["overwriteUsers"] = o.OverwriteUsers.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverCouchbaseParams struct {
	value *RecoverCouchbaseParams
	isSet bool
}

func (v NullableRecoverCouchbaseParams) Get() *RecoverCouchbaseParams {
	return v.value
}

func (v *NullableRecoverCouchbaseParams) Set(val *RecoverCouchbaseParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverCouchbaseParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverCouchbaseParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverCouchbaseParams(val *RecoverCouchbaseParams) *NullableRecoverCouchbaseParams {
	return &NullableRecoverCouchbaseParams{value: val, isSet: true}
}

func (v NullableRecoverCouchbaseParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverCouchbaseParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RecoverCouchbaseParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}