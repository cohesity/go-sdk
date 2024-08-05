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

// checks if the ExchangeProtectionGroupParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeProtectionGroupParams{}

// ExchangeProtectionGroupParams Specifies the parameters which are specific to Exchange related Protection Groups.
type ExchangeProtectionGroupParams struct {
	// Specifies whether the backups should be copy-only.
	BackupsCopyOnly NullableBool `json:"backupsCopyOnly,omitempty"`
	// Specifies the list of IDs of the databases to not be protected by this Protection Group. This can be used to ignore specific databases under Exchange Server/DAG which has been included for protection.
	ExcludeDatabaseIds []*int64 `json:"excludeDatabaseIds,omitempty"`
	IndexingPolicy *IndexingPolicy `json:"indexingPolicy,omitempty"`
	// Specifies the list of object ids to be protected.
	Objects []ExchangeProtectionGroupObjectParams `json:"objects"`
}

type _ExchangeProtectionGroupParams ExchangeProtectionGroupParams

// NewExchangeProtectionGroupParams instantiates a new ExchangeProtectionGroupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeProtectionGroupParams(objects []ExchangeProtectionGroupObjectParams) *ExchangeProtectionGroupParams {
	this := ExchangeProtectionGroupParams{}
	this.Objects = objects
	return &this
}

// NewExchangeProtectionGroupParamsWithDefaults instantiates a new ExchangeProtectionGroupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeProtectionGroupParamsWithDefaults() *ExchangeProtectionGroupParams {
	this := ExchangeProtectionGroupParams{}
	return &this
}

// GetBackupsCopyOnly returns the BackupsCopyOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExchangeProtectionGroupParams) GetBackupsCopyOnly() bool {
	if o == nil || IsNil(o.BackupsCopyOnly.Get()) {
		var ret bool
		return ret
	}
	return *o.BackupsCopyOnly.Get()
}

// GetBackupsCopyOnlyOk returns a tuple with the BackupsCopyOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExchangeProtectionGroupParams) GetBackupsCopyOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupsCopyOnly.Get(), o.BackupsCopyOnly.IsSet()
}

// HasBackupsCopyOnly returns a boolean if a field has been set.
func (o *ExchangeProtectionGroupParams) HasBackupsCopyOnly() bool {
	if o != nil && o.BackupsCopyOnly.IsSet() {
		return true
	}

	return false
}

// SetBackupsCopyOnly gets a reference to the given NullableBool and assigns it to the BackupsCopyOnly field.
func (o *ExchangeProtectionGroupParams) SetBackupsCopyOnly(v bool) {
	o.BackupsCopyOnly.Set(&v)
}
// SetBackupsCopyOnlyNil sets the value for BackupsCopyOnly to be an explicit nil
func (o *ExchangeProtectionGroupParams) SetBackupsCopyOnlyNil() {
	o.BackupsCopyOnly.Set(nil)
}

// UnsetBackupsCopyOnly ensures that no value is present for BackupsCopyOnly, not even an explicit nil
func (o *ExchangeProtectionGroupParams) UnsetBackupsCopyOnly() {
	o.BackupsCopyOnly.Unset()
}

// GetExcludeDatabaseIds returns the ExcludeDatabaseIds field value if set, zero value otherwise.
func (o *ExchangeProtectionGroupParams) GetExcludeDatabaseIds() []*int64 {
	if o == nil || IsNil(o.ExcludeDatabaseIds) {
		var ret []*int64
		return ret
	}
	return o.ExcludeDatabaseIds
}

// GetExcludeDatabaseIdsOk returns a tuple with the ExcludeDatabaseIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeProtectionGroupParams) GetExcludeDatabaseIdsOk() ([]*int64, bool) {
	if o == nil || IsNil(o.ExcludeDatabaseIds) {
		return nil, false
	}
	return o.ExcludeDatabaseIds, true
}

// HasExcludeDatabaseIds returns a boolean if a field has been set.
func (o *ExchangeProtectionGroupParams) HasExcludeDatabaseIds() bool {
	if o != nil && !IsNil(o.ExcludeDatabaseIds) {
		return true
	}

	return false
}

// SetExcludeDatabaseIds gets a reference to the given []*int64 and assigns it to the ExcludeDatabaseIds field.
func (o *ExchangeProtectionGroupParams) SetExcludeDatabaseIds(v []*int64) {
	o.ExcludeDatabaseIds = v
}

// GetIndexingPolicy returns the IndexingPolicy field value if set, zero value otherwise.
func (o *ExchangeProtectionGroupParams) GetIndexingPolicy() IndexingPolicy {
	if o == nil || IsNil(o.IndexingPolicy) {
		var ret IndexingPolicy
		return ret
	}
	return *o.IndexingPolicy
}

// GetIndexingPolicyOk returns a tuple with the IndexingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool) {
	if o == nil || IsNil(o.IndexingPolicy) {
		return nil, false
	}
	return o.IndexingPolicy, true
}

// HasIndexingPolicy returns a boolean if a field has been set.
func (o *ExchangeProtectionGroupParams) HasIndexingPolicy() bool {
	if o != nil && !IsNil(o.IndexingPolicy) {
		return true
	}

	return false
}

// SetIndexingPolicy gets a reference to the given IndexingPolicy and assigns it to the IndexingPolicy field.
func (o *ExchangeProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy) {
	o.IndexingPolicy = &v
}

// GetObjects returns the Objects field value
func (o *ExchangeProtectionGroupParams) GetObjects() []ExchangeProtectionGroupObjectParams {
	if o == nil {
		var ret []ExchangeProtectionGroupObjectParams
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *ExchangeProtectionGroupParams) GetObjectsOk() ([]ExchangeProtectionGroupObjectParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.Objects, true
}

// SetObjects sets field value
func (o *ExchangeProtectionGroupParams) SetObjects(v []ExchangeProtectionGroupObjectParams) {
	o.Objects = v
}

func (o ExchangeProtectionGroupParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeProtectionGroupParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BackupsCopyOnly.IsSet() {
		toSerialize["backupsCopyOnly"] = o.BackupsCopyOnly.Get()
	}
	if !IsNil(o.ExcludeDatabaseIds) {
		toSerialize["excludeDatabaseIds"] = o.ExcludeDatabaseIds
	}
	if !IsNil(o.IndexingPolicy) {
		toSerialize["indexingPolicy"] = o.IndexingPolicy
	}
	toSerialize["objects"] = o.Objects
	return toSerialize, nil
}

func (o *ExchangeProtectionGroupParams) UnmarshalJSON(data []byte) (err error) {
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

	varExchangeProtectionGroupParams := _ExchangeProtectionGroupParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExchangeProtectionGroupParams)

	if err != nil {
		return err
	}

	*o = ExchangeProtectionGroupParams(varExchangeProtectionGroupParams)

	return err
}

type NullableExchangeProtectionGroupParams struct {
	value *ExchangeProtectionGroupParams
	isSet bool
}

func (v NullableExchangeProtectionGroupParams) Get() *ExchangeProtectionGroupParams {
	return v.value
}

func (v *NullableExchangeProtectionGroupParams) Set(val *ExchangeProtectionGroupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeProtectionGroupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeProtectionGroupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeProtectionGroupParams(val *ExchangeProtectionGroupParams) *NullableExchangeProtectionGroupParams {
	return &NullableExchangeProtectionGroupParams{value: val, isSet: true}
}

func (v NullableExchangeProtectionGroupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeProtectionGroupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


