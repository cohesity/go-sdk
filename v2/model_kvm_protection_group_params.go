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

// checks if the KvmProtectionGroupParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KvmProtectionGroupParams{}

// KvmProtectionGroupParams Specifies the parameters which are specific to Kvm related Protection Groups.
type KvmProtectionGroupParams struct {
	// Specifies whether or not to quiesce apps and the file system in order to take app consistent snapshots. If not specified or false then snapshots will not be app consistent.
	AppConsistentSnapshot NullableBool `json:"appConsistentSnapshot,omitempty"`
	// Specifies the object ids to be excluded in the Protection Group.
	ExcludeObjectIds []int64 `json:"excludeObjectIds,omitempty"`
	IndexingPolicy *IndexingPolicy `json:"indexingPolicy,omitempty"`
	// Specifies the objects to be included in the Protection Group.
	Objects []KvmProtectionGroupObjectParams `json:"objects"`
	// Specifies the id of the parent of the objects.
	SourceId NullableInt64 `json:"sourceId,omitempty"`
	// Specifies the name of the parent of the objects.
	SourceName NullableString `json:"sourceName,omitempty"`
}

type _KvmProtectionGroupParams KvmProtectionGroupParams

// NewKvmProtectionGroupParams instantiates a new KvmProtectionGroupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmProtectionGroupParams(objects []KvmProtectionGroupObjectParams) *KvmProtectionGroupParams {
	this := KvmProtectionGroupParams{}
	this.Objects = objects
	return &this
}

// NewKvmProtectionGroupParamsWithDefaults instantiates a new KvmProtectionGroupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmProtectionGroupParamsWithDefaults() *KvmProtectionGroupParams {
	this := KvmProtectionGroupParams{}
	return &this
}

// GetAppConsistentSnapshot returns the AppConsistentSnapshot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KvmProtectionGroupParams) GetAppConsistentSnapshot() bool {
	if o == nil || IsNil(o.AppConsistentSnapshot.Get()) {
		var ret bool
		return ret
	}
	return *o.AppConsistentSnapshot.Get()
}

// GetAppConsistentSnapshotOk returns a tuple with the AppConsistentSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KvmProtectionGroupParams) GetAppConsistentSnapshotOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppConsistentSnapshot.Get(), o.AppConsistentSnapshot.IsSet()
}

// HasAppConsistentSnapshot returns a boolean if a field has been set.
func (o *KvmProtectionGroupParams) HasAppConsistentSnapshot() bool {
	if o != nil && o.AppConsistentSnapshot.IsSet() {
		return true
	}

	return false
}

// SetAppConsistentSnapshot gets a reference to the given NullableBool and assigns it to the AppConsistentSnapshot field.
func (o *KvmProtectionGroupParams) SetAppConsistentSnapshot(v bool) {
	o.AppConsistentSnapshot.Set(&v)
}
// SetAppConsistentSnapshotNil sets the value for AppConsistentSnapshot to be an explicit nil
func (o *KvmProtectionGroupParams) SetAppConsistentSnapshotNil() {
	o.AppConsistentSnapshot.Set(nil)
}

// UnsetAppConsistentSnapshot ensures that no value is present for AppConsistentSnapshot, not even an explicit nil
func (o *KvmProtectionGroupParams) UnsetAppConsistentSnapshot() {
	o.AppConsistentSnapshot.Unset()
}

// GetExcludeObjectIds returns the ExcludeObjectIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KvmProtectionGroupParams) GetExcludeObjectIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.ExcludeObjectIds
}

// GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KvmProtectionGroupParams) GetExcludeObjectIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.ExcludeObjectIds) {
		return nil, false
	}
	return o.ExcludeObjectIds, true
}

// HasExcludeObjectIds returns a boolean if a field has been set.
func (o *KvmProtectionGroupParams) HasExcludeObjectIds() bool {
	if o != nil && !IsNil(o.ExcludeObjectIds) {
		return true
	}

	return false
}

// SetExcludeObjectIds gets a reference to the given []int64 and assigns it to the ExcludeObjectIds field.
func (o *KvmProtectionGroupParams) SetExcludeObjectIds(v []int64) {
	o.ExcludeObjectIds = v
}

// GetIndexingPolicy returns the IndexingPolicy field value if set, zero value otherwise.
func (o *KvmProtectionGroupParams) GetIndexingPolicy() IndexingPolicy {
	if o == nil || IsNil(o.IndexingPolicy) {
		var ret IndexingPolicy
		return ret
	}
	return *o.IndexingPolicy
}

// GetIndexingPolicyOk returns a tuple with the IndexingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool) {
	if o == nil || IsNil(o.IndexingPolicy) {
		return nil, false
	}
	return o.IndexingPolicy, true
}

// HasIndexingPolicy returns a boolean if a field has been set.
func (o *KvmProtectionGroupParams) HasIndexingPolicy() bool {
	if o != nil && !IsNil(o.IndexingPolicy) {
		return true
	}

	return false
}

// SetIndexingPolicy gets a reference to the given IndexingPolicy and assigns it to the IndexingPolicy field.
func (o *KvmProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy) {
	o.IndexingPolicy = &v
}

// GetObjects returns the Objects field value
func (o *KvmProtectionGroupParams) GetObjects() []KvmProtectionGroupObjectParams {
	if o == nil {
		var ret []KvmProtectionGroupObjectParams
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *KvmProtectionGroupParams) GetObjectsOk() ([]KvmProtectionGroupObjectParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.Objects, true
}

// SetObjects sets field value
func (o *KvmProtectionGroupParams) SetObjects(v []KvmProtectionGroupObjectParams) {
	o.Objects = v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KvmProtectionGroupParams) GetSourceId() int64 {
	if o == nil || IsNil(o.SourceId.Get()) {
		var ret int64
		return ret
	}
	return *o.SourceId.Get()
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KvmProtectionGroupParams) GetSourceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceId.Get(), o.SourceId.IsSet()
}

// HasSourceId returns a boolean if a field has been set.
func (o *KvmProtectionGroupParams) HasSourceId() bool {
	if o != nil && o.SourceId.IsSet() {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given NullableInt64 and assigns it to the SourceId field.
func (o *KvmProtectionGroupParams) SetSourceId(v int64) {
	o.SourceId.Set(&v)
}
// SetSourceIdNil sets the value for SourceId to be an explicit nil
func (o *KvmProtectionGroupParams) SetSourceIdNil() {
	o.SourceId.Set(nil)
}

// UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
func (o *KvmProtectionGroupParams) UnsetSourceId() {
	o.SourceId.Unset()
}

// GetSourceName returns the SourceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KvmProtectionGroupParams) GetSourceName() string {
	if o == nil || IsNil(o.SourceName.Get()) {
		var ret string
		return ret
	}
	return *o.SourceName.Get()
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KvmProtectionGroupParams) GetSourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceName.Get(), o.SourceName.IsSet()
}

// HasSourceName returns a boolean if a field has been set.
func (o *KvmProtectionGroupParams) HasSourceName() bool {
	if o != nil && o.SourceName.IsSet() {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given NullableString and assigns it to the SourceName field.
func (o *KvmProtectionGroupParams) SetSourceName(v string) {
	o.SourceName.Set(&v)
}
// SetSourceNameNil sets the value for SourceName to be an explicit nil
func (o *KvmProtectionGroupParams) SetSourceNameNil() {
	o.SourceName.Set(nil)
}

// UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
func (o *KvmProtectionGroupParams) UnsetSourceName() {
	o.SourceName.Unset()
}

func (o KvmProtectionGroupParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KvmProtectionGroupParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AppConsistentSnapshot.IsSet() {
		toSerialize["appConsistentSnapshot"] = o.AppConsistentSnapshot.Get()
	}
	if o.ExcludeObjectIds != nil {
		toSerialize["excludeObjectIds"] = o.ExcludeObjectIds
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

func (o *KvmProtectionGroupParams) UnmarshalJSON(data []byte) (err error) {
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

	varKvmProtectionGroupParams := _KvmProtectionGroupParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKvmProtectionGroupParams)

	if err != nil {
		return err
	}

	*o = KvmProtectionGroupParams(varKvmProtectionGroupParams)

	return err
}

type NullableKvmProtectionGroupParams struct {
	value *KvmProtectionGroupParams
	isSet bool
}

func (v NullableKvmProtectionGroupParams) Get() *KvmProtectionGroupParams {
	return v.value
}

func (v *NullableKvmProtectionGroupParams) Set(val *KvmProtectionGroupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmProtectionGroupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmProtectionGroupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmProtectionGroupParams(val *KvmProtectionGroupParams) *NullableKvmProtectionGroupParams {
	return &NullableKvmProtectionGroupParams{value: val, isSet: true}
}

func (v NullableKvmProtectionGroupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmProtectionGroupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


