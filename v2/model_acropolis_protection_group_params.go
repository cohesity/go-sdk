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

// checks if the AcropolisProtectionGroupParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcropolisProtectionGroupParams{}

// AcropolisProtectionGroupParams Specifies the parameters which are related to Acropolis Protection Groups.
type AcropolisProtectionGroupParams struct {
	// Specifies whether or not to quiesce apps and the file system in order to take app consistent snapshots. If not specified or false then snapshots will not be app consistent.
	AppConsistentSnapshot NullableBool `json:"appConsistentSnapshot,omitempty"`
	// Specifies whether to continue backing up on quiesce failure
	ContinueOnQuiesceFailure NullableBool `json:"continueOnQuiesceFailure,omitempty"`
	// Specifies the object ids to be excluded in the Protection Group.
	ExcludeObjectIds []int64 `json:"excludeObjectIds,omitempty"`
	// Specifies a list of disks to exclude from the backup.
	GlobalExcludeDisks []AcropolisDiskInfo `json:"globalExcludeDisks,omitempty"`
	// Specifies a list of disks to include in the backup.
	GlobalIncludeDisks []AcropolisDiskInfo `json:"globalIncludeDisks,omitempty"`
	IndexingPolicy *IndexingPolicy `json:"indexingPolicy,omitempty"`
	// Specifies the objects included in the Protection Group.
	Objects []AcropolisProtectionGroupObjectParams `json:"objects"`
	// Specifies the id of the parent of the objects.
	SourceId NullableInt64 `json:"sourceId,omitempty"`
	// Specifies the name of the parent of the objects.
	SourceName NullableString `json:"sourceName,omitempty"`
}

type _AcropolisProtectionGroupParams AcropolisProtectionGroupParams

// NewAcropolisProtectionGroupParams instantiates a new AcropolisProtectionGroupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcropolisProtectionGroupParams(objects []AcropolisProtectionGroupObjectParams) *AcropolisProtectionGroupParams {
	this := AcropolisProtectionGroupParams{}
	this.Objects = objects
	return &this
}

// NewAcropolisProtectionGroupParamsWithDefaults instantiates a new AcropolisProtectionGroupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcropolisProtectionGroupParamsWithDefaults() *AcropolisProtectionGroupParams {
	this := AcropolisProtectionGroupParams{}
	return &this
}

// GetAppConsistentSnapshot returns the AppConsistentSnapshot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcropolisProtectionGroupParams) GetAppConsistentSnapshot() bool {
	if o == nil || IsNil(o.AppConsistentSnapshot.Get()) {
		var ret bool
		return ret
	}
	return *o.AppConsistentSnapshot.Get()
}

// GetAppConsistentSnapshotOk returns a tuple with the AppConsistentSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcropolisProtectionGroupParams) GetAppConsistentSnapshotOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppConsistentSnapshot.Get(), o.AppConsistentSnapshot.IsSet()
}

// HasAppConsistentSnapshot returns a boolean if a field has been set.
func (o *AcropolisProtectionGroupParams) HasAppConsistentSnapshot() bool {
	if o != nil && o.AppConsistentSnapshot.IsSet() {
		return true
	}

	return false
}

// SetAppConsistentSnapshot gets a reference to the given NullableBool and assigns it to the AppConsistentSnapshot field.
func (o *AcropolisProtectionGroupParams) SetAppConsistentSnapshot(v bool) {
	o.AppConsistentSnapshot.Set(&v)
}
// SetAppConsistentSnapshotNil sets the value for AppConsistentSnapshot to be an explicit nil
func (o *AcropolisProtectionGroupParams) SetAppConsistentSnapshotNil() {
	o.AppConsistentSnapshot.Set(nil)
}

// UnsetAppConsistentSnapshot ensures that no value is present for AppConsistentSnapshot, not even an explicit nil
func (o *AcropolisProtectionGroupParams) UnsetAppConsistentSnapshot() {
	o.AppConsistentSnapshot.Unset()
}

// GetContinueOnQuiesceFailure returns the ContinueOnQuiesceFailure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcropolisProtectionGroupParams) GetContinueOnQuiesceFailure() bool {
	if o == nil || IsNil(o.ContinueOnQuiesceFailure.Get()) {
		var ret bool
		return ret
	}
	return *o.ContinueOnQuiesceFailure.Get()
}

// GetContinueOnQuiesceFailureOk returns a tuple with the ContinueOnQuiesceFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcropolisProtectionGroupParams) GetContinueOnQuiesceFailureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinueOnQuiesceFailure.Get(), o.ContinueOnQuiesceFailure.IsSet()
}

// HasContinueOnQuiesceFailure returns a boolean if a field has been set.
func (o *AcropolisProtectionGroupParams) HasContinueOnQuiesceFailure() bool {
	if o != nil && o.ContinueOnQuiesceFailure.IsSet() {
		return true
	}

	return false
}

// SetContinueOnQuiesceFailure gets a reference to the given NullableBool and assigns it to the ContinueOnQuiesceFailure field.
func (o *AcropolisProtectionGroupParams) SetContinueOnQuiesceFailure(v bool) {
	o.ContinueOnQuiesceFailure.Set(&v)
}
// SetContinueOnQuiesceFailureNil sets the value for ContinueOnQuiesceFailure to be an explicit nil
func (o *AcropolisProtectionGroupParams) SetContinueOnQuiesceFailureNil() {
	o.ContinueOnQuiesceFailure.Set(nil)
}

// UnsetContinueOnQuiesceFailure ensures that no value is present for ContinueOnQuiesceFailure, not even an explicit nil
func (o *AcropolisProtectionGroupParams) UnsetContinueOnQuiesceFailure() {
	o.ContinueOnQuiesceFailure.Unset()
}

// GetExcludeObjectIds returns the ExcludeObjectIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcropolisProtectionGroupParams) GetExcludeObjectIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.ExcludeObjectIds
}

// GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcropolisProtectionGroupParams) GetExcludeObjectIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.ExcludeObjectIds) {
		return nil, false
	}
	return o.ExcludeObjectIds, true
}

// HasExcludeObjectIds returns a boolean if a field has been set.
func (o *AcropolisProtectionGroupParams) HasExcludeObjectIds() bool {
	if o != nil && !IsNil(o.ExcludeObjectIds) {
		return true
	}

	return false
}

// SetExcludeObjectIds gets a reference to the given []int64 and assigns it to the ExcludeObjectIds field.
func (o *AcropolisProtectionGroupParams) SetExcludeObjectIds(v []int64) {
	o.ExcludeObjectIds = v
}

// GetGlobalExcludeDisks returns the GlobalExcludeDisks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcropolisProtectionGroupParams) GetGlobalExcludeDisks() []AcropolisDiskInfo {
	if o == nil {
		var ret []AcropolisDiskInfo
		return ret
	}
	return o.GlobalExcludeDisks
}

// GetGlobalExcludeDisksOk returns a tuple with the GlobalExcludeDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcropolisProtectionGroupParams) GetGlobalExcludeDisksOk() ([]AcropolisDiskInfo, bool) {
	if o == nil || IsNil(o.GlobalExcludeDisks) {
		return nil, false
	}
	return o.GlobalExcludeDisks, true
}

// HasGlobalExcludeDisks returns a boolean if a field has been set.
func (o *AcropolisProtectionGroupParams) HasGlobalExcludeDisks() bool {
	if o != nil && !IsNil(o.GlobalExcludeDisks) {
		return true
	}

	return false
}

// SetGlobalExcludeDisks gets a reference to the given []AcropolisDiskInfo and assigns it to the GlobalExcludeDisks field.
func (o *AcropolisProtectionGroupParams) SetGlobalExcludeDisks(v []AcropolisDiskInfo) {
	o.GlobalExcludeDisks = v
}

// GetGlobalIncludeDisks returns the GlobalIncludeDisks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcropolisProtectionGroupParams) GetGlobalIncludeDisks() []AcropolisDiskInfo {
	if o == nil {
		var ret []AcropolisDiskInfo
		return ret
	}
	return o.GlobalIncludeDisks
}

// GetGlobalIncludeDisksOk returns a tuple with the GlobalIncludeDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcropolisProtectionGroupParams) GetGlobalIncludeDisksOk() ([]AcropolisDiskInfo, bool) {
	if o == nil || IsNil(o.GlobalIncludeDisks) {
		return nil, false
	}
	return o.GlobalIncludeDisks, true
}

// HasGlobalIncludeDisks returns a boolean if a field has been set.
func (o *AcropolisProtectionGroupParams) HasGlobalIncludeDisks() bool {
	if o != nil && !IsNil(o.GlobalIncludeDisks) {
		return true
	}

	return false
}

// SetGlobalIncludeDisks gets a reference to the given []AcropolisDiskInfo and assigns it to the GlobalIncludeDisks field.
func (o *AcropolisProtectionGroupParams) SetGlobalIncludeDisks(v []AcropolisDiskInfo) {
	o.GlobalIncludeDisks = v
}

// GetIndexingPolicy returns the IndexingPolicy field value if set, zero value otherwise.
func (o *AcropolisProtectionGroupParams) GetIndexingPolicy() IndexingPolicy {
	if o == nil || IsNil(o.IndexingPolicy) {
		var ret IndexingPolicy
		return ret
	}
	return *o.IndexingPolicy
}

// GetIndexingPolicyOk returns a tuple with the IndexingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcropolisProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool) {
	if o == nil || IsNil(o.IndexingPolicy) {
		return nil, false
	}
	return o.IndexingPolicy, true
}

// HasIndexingPolicy returns a boolean if a field has been set.
func (o *AcropolisProtectionGroupParams) HasIndexingPolicy() bool {
	if o != nil && !IsNil(o.IndexingPolicy) {
		return true
	}

	return false
}

// SetIndexingPolicy gets a reference to the given IndexingPolicy and assigns it to the IndexingPolicy field.
func (o *AcropolisProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy) {
	o.IndexingPolicy = &v
}

// GetObjects returns the Objects field value
func (o *AcropolisProtectionGroupParams) GetObjects() []AcropolisProtectionGroupObjectParams {
	if o == nil {
		var ret []AcropolisProtectionGroupObjectParams
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *AcropolisProtectionGroupParams) GetObjectsOk() ([]AcropolisProtectionGroupObjectParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.Objects, true
}

// SetObjects sets field value
func (o *AcropolisProtectionGroupParams) SetObjects(v []AcropolisProtectionGroupObjectParams) {
	o.Objects = v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcropolisProtectionGroupParams) GetSourceId() int64 {
	if o == nil || IsNil(o.SourceId.Get()) {
		var ret int64
		return ret
	}
	return *o.SourceId.Get()
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcropolisProtectionGroupParams) GetSourceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceId.Get(), o.SourceId.IsSet()
}

// HasSourceId returns a boolean if a field has been set.
func (o *AcropolisProtectionGroupParams) HasSourceId() bool {
	if o != nil && o.SourceId.IsSet() {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given NullableInt64 and assigns it to the SourceId field.
func (o *AcropolisProtectionGroupParams) SetSourceId(v int64) {
	o.SourceId.Set(&v)
}
// SetSourceIdNil sets the value for SourceId to be an explicit nil
func (o *AcropolisProtectionGroupParams) SetSourceIdNil() {
	o.SourceId.Set(nil)
}

// UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
func (o *AcropolisProtectionGroupParams) UnsetSourceId() {
	o.SourceId.Unset()
}

// GetSourceName returns the SourceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcropolisProtectionGroupParams) GetSourceName() string {
	if o == nil || IsNil(o.SourceName.Get()) {
		var ret string
		return ret
	}
	return *o.SourceName.Get()
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcropolisProtectionGroupParams) GetSourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceName.Get(), o.SourceName.IsSet()
}

// HasSourceName returns a boolean if a field has been set.
func (o *AcropolisProtectionGroupParams) HasSourceName() bool {
	if o != nil && o.SourceName.IsSet() {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given NullableString and assigns it to the SourceName field.
func (o *AcropolisProtectionGroupParams) SetSourceName(v string) {
	o.SourceName.Set(&v)
}
// SetSourceNameNil sets the value for SourceName to be an explicit nil
func (o *AcropolisProtectionGroupParams) SetSourceNameNil() {
	o.SourceName.Set(nil)
}

// UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
func (o *AcropolisProtectionGroupParams) UnsetSourceName() {
	o.SourceName.Unset()
}

func (o AcropolisProtectionGroupParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcropolisProtectionGroupParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AppConsistentSnapshot.IsSet() {
		toSerialize["appConsistentSnapshot"] = o.AppConsistentSnapshot.Get()
	}
	if o.ContinueOnQuiesceFailure.IsSet() {
		toSerialize["continueOnQuiesceFailure"] = o.ContinueOnQuiesceFailure.Get()
	}
	if o.ExcludeObjectIds != nil {
		toSerialize["excludeObjectIds"] = o.ExcludeObjectIds
	}
	if o.GlobalExcludeDisks != nil {
		toSerialize["globalExcludeDisks"] = o.GlobalExcludeDisks
	}
	if o.GlobalIncludeDisks != nil {
		toSerialize["globalIncludeDisks"] = o.GlobalIncludeDisks
	}
	if !IsNil(o.IndexingPolicy) {
		toSerialize["indexingPolicy"] = o.IndexingPolicy
	}
	toSerialize["objects"] = o.Objects
	if o.SourceId.IsSet() {
		toSerialize["sourceId"] = o.SourceId.Get()
	}
	if o.SourceName.IsSet() {
		toSerialize["sourceName"] = o.SourceName.Get()
	}
	return toSerialize, nil
}

func (o *AcropolisProtectionGroupParams) UnmarshalJSON(data []byte) (err error) {
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

	varAcropolisProtectionGroupParams := _AcropolisProtectionGroupParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAcropolisProtectionGroupParams)

	if err != nil {
		return err
	}

	*o = AcropolisProtectionGroupParams(varAcropolisProtectionGroupParams)

	return err
}

type NullableAcropolisProtectionGroupParams struct {
	value *AcropolisProtectionGroupParams
	isSet bool
}

func (v NullableAcropolisProtectionGroupParams) Get() *AcropolisProtectionGroupParams {
	return v.value
}

func (v *NullableAcropolisProtectionGroupParams) Set(val *AcropolisProtectionGroupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAcropolisProtectionGroupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAcropolisProtectionGroupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcropolisProtectionGroupParams(val *AcropolisProtectionGroupParams) *NullableAcropolisProtectionGroupParams {
	return &NullableAcropolisProtectionGroupParams{value: val, isSet: true}
}

func (v NullableAcropolisProtectionGroupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcropolisProtectionGroupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


