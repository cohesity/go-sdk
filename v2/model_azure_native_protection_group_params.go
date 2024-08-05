/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the AzureNativeProtectionGroupParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureNativeProtectionGroupParams{}

// AzureNativeProtectionGroupParams Specifies the parameters which are specific to Azure related Protection Groups using Azure native snapshot APIs. Objects must be specified.
type AzureNativeProtectionGroupParams struct {
	// Specifies whether or not to move the workload to the cloud.
	CloudMigration NullableBool `json:"cloudMigration,omitempty"`
	CloudPrePostScript *CloudBackupScriptParams `json:"cloudPrePostScript,omitempty"`
	DataTransferInfo *DataTransferInfo `json:"dataTransferInfo,omitempty"`
	// Specifies the objects to be excluded in the Protection Group.
	ExcludeObjectIds []int64 `json:"excludeObjectIds,omitempty"`
	// Array of arrays of VM Tag Ids that Specify VMs to Exclude.
	ExcludeVmTagIds [][]int64 `json:"excludeVmTagIds,omitempty"`
	IndexingPolicy *IndexingPolicy `json:"indexingPolicy,omitempty"`
	// Specifies the objects to be included in the Protection Group.
	Objects []AzureNativeProtectionGroupObjectParams `json:"objects,omitempty"`
	// Specifies the id of the parent of the objects.
	SourceId NullableInt64 `json:"sourceId,omitempty"`
	// Specifies the name of the parent of the objects.
	SourceName NullableString `json:"sourceName,omitempty"`
	// Array of arrays of VM Tag Ids that Specify VMs to Protect.
	VmTagIds [][]int64 `json:"vmTagIds,omitempty"`
}

// NewAzureNativeProtectionGroupParams instantiates a new AzureNativeProtectionGroupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureNativeProtectionGroupParams() *AzureNativeProtectionGroupParams {
	this := AzureNativeProtectionGroupParams{}
	return &this
}

// NewAzureNativeProtectionGroupParamsWithDefaults instantiates a new AzureNativeProtectionGroupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureNativeProtectionGroupParamsWithDefaults() *AzureNativeProtectionGroupParams {
	this := AzureNativeProtectionGroupParams{}
	return &this
}

// GetCloudMigration returns the CloudMigration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureNativeProtectionGroupParams) GetCloudMigration() bool {
	if o == nil || IsNil(o.CloudMigration.Get()) {
		var ret bool
		return ret
	}
	return *o.CloudMigration.Get()
}

// GetCloudMigrationOk returns a tuple with the CloudMigration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureNativeProtectionGroupParams) GetCloudMigrationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudMigration.Get(), o.CloudMigration.IsSet()
}

// HasCloudMigration returns a boolean if a field has been set.
func (o *AzureNativeProtectionGroupParams) HasCloudMigration() bool {
	if o != nil && o.CloudMigration.IsSet() {
		return true
	}

	return false
}

// SetCloudMigration gets a reference to the given NullableBool and assigns it to the CloudMigration field.
func (o *AzureNativeProtectionGroupParams) SetCloudMigration(v bool) {
	o.CloudMigration.Set(&v)
}
// SetCloudMigrationNil sets the value for CloudMigration to be an explicit nil
func (o *AzureNativeProtectionGroupParams) SetCloudMigrationNil() {
	o.CloudMigration.Set(nil)
}

// UnsetCloudMigration ensures that no value is present for CloudMigration, not even an explicit nil
func (o *AzureNativeProtectionGroupParams) UnsetCloudMigration() {
	o.CloudMigration.Unset()
}

// GetCloudPrePostScript returns the CloudPrePostScript field value if set, zero value otherwise.
func (o *AzureNativeProtectionGroupParams) GetCloudPrePostScript() CloudBackupScriptParams {
	if o == nil || IsNil(o.CloudPrePostScript) {
		var ret CloudBackupScriptParams
		return ret
	}
	return *o.CloudPrePostScript
}

// GetCloudPrePostScriptOk returns a tuple with the CloudPrePostScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureNativeProtectionGroupParams) GetCloudPrePostScriptOk() (*CloudBackupScriptParams, bool) {
	if o == nil || IsNil(o.CloudPrePostScript) {
		return nil, false
	}
	return o.CloudPrePostScript, true
}

// HasCloudPrePostScript returns a boolean if a field has been set.
func (o *AzureNativeProtectionGroupParams) HasCloudPrePostScript() bool {
	if o != nil && !IsNil(o.CloudPrePostScript) {
		return true
	}

	return false
}

// SetCloudPrePostScript gets a reference to the given CloudBackupScriptParams and assigns it to the CloudPrePostScript field.
func (o *AzureNativeProtectionGroupParams) SetCloudPrePostScript(v CloudBackupScriptParams) {
	o.CloudPrePostScript = &v
}

// GetDataTransferInfo returns the DataTransferInfo field value if set, zero value otherwise.
func (o *AzureNativeProtectionGroupParams) GetDataTransferInfo() DataTransferInfo {
	if o == nil || IsNil(o.DataTransferInfo) {
		var ret DataTransferInfo
		return ret
	}
	return *o.DataTransferInfo
}

// GetDataTransferInfoOk returns a tuple with the DataTransferInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureNativeProtectionGroupParams) GetDataTransferInfoOk() (*DataTransferInfo, bool) {
	if o == nil || IsNil(o.DataTransferInfo) {
		return nil, false
	}
	return o.DataTransferInfo, true
}

// HasDataTransferInfo returns a boolean if a field has been set.
func (o *AzureNativeProtectionGroupParams) HasDataTransferInfo() bool {
	if o != nil && !IsNil(o.DataTransferInfo) {
		return true
	}

	return false
}

// SetDataTransferInfo gets a reference to the given DataTransferInfo and assigns it to the DataTransferInfo field.
func (o *AzureNativeProtectionGroupParams) SetDataTransferInfo(v DataTransferInfo) {
	o.DataTransferInfo = &v
}

// GetExcludeObjectIds returns the ExcludeObjectIds field value if set, zero value otherwise.
func (o *AzureNativeProtectionGroupParams) GetExcludeObjectIds() []int64 {
	if o == nil || IsNil(o.ExcludeObjectIds) {
		var ret []int64
		return ret
	}
	return o.ExcludeObjectIds
}

// GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureNativeProtectionGroupParams) GetExcludeObjectIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.ExcludeObjectIds) {
		return nil, false
	}
	return o.ExcludeObjectIds, true
}

// HasExcludeObjectIds returns a boolean if a field has been set.
func (o *AzureNativeProtectionGroupParams) HasExcludeObjectIds() bool {
	if o != nil && !IsNil(o.ExcludeObjectIds) {
		return true
	}

	return false
}

// SetExcludeObjectIds gets a reference to the given []int64 and assigns it to the ExcludeObjectIds field.
func (o *AzureNativeProtectionGroupParams) SetExcludeObjectIds(v []int64) {
	o.ExcludeObjectIds = v
}

// GetExcludeVmTagIds returns the ExcludeVmTagIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureNativeProtectionGroupParams) GetExcludeVmTagIds() [][]int64 {
	if o == nil {
		var ret [][]int64
		return ret
	}
	return o.ExcludeVmTagIds
}

// GetExcludeVmTagIdsOk returns a tuple with the ExcludeVmTagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureNativeProtectionGroupParams) GetExcludeVmTagIdsOk() ([][]int64, bool) {
	if o == nil || IsNil(o.ExcludeVmTagIds) {
		return nil, false
	}
	return o.ExcludeVmTagIds, true
}

// HasExcludeVmTagIds returns a boolean if a field has been set.
func (o *AzureNativeProtectionGroupParams) HasExcludeVmTagIds() bool {
	if o != nil && !IsNil(o.ExcludeVmTagIds) {
		return true
	}

	return false
}

// SetExcludeVmTagIds gets a reference to the given [][]int64 and assigns it to the ExcludeVmTagIds field.
func (o *AzureNativeProtectionGroupParams) SetExcludeVmTagIds(v [][]int64) {
	o.ExcludeVmTagIds = v
}

// GetIndexingPolicy returns the IndexingPolicy field value if set, zero value otherwise.
func (o *AzureNativeProtectionGroupParams) GetIndexingPolicy() IndexingPolicy {
	if o == nil || IsNil(o.IndexingPolicy) {
		var ret IndexingPolicy
		return ret
	}
	return *o.IndexingPolicy
}

// GetIndexingPolicyOk returns a tuple with the IndexingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureNativeProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool) {
	if o == nil || IsNil(o.IndexingPolicy) {
		return nil, false
	}
	return o.IndexingPolicy, true
}

// HasIndexingPolicy returns a boolean if a field has been set.
func (o *AzureNativeProtectionGroupParams) HasIndexingPolicy() bool {
	if o != nil && !IsNil(o.IndexingPolicy) {
		return true
	}

	return false
}

// SetIndexingPolicy gets a reference to the given IndexingPolicy and assigns it to the IndexingPolicy field.
func (o *AzureNativeProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy) {
	o.IndexingPolicy = &v
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *AzureNativeProtectionGroupParams) GetObjects() []AzureNativeProtectionGroupObjectParams {
	if o == nil || IsNil(o.Objects) {
		var ret []AzureNativeProtectionGroupObjectParams
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureNativeProtectionGroupParams) GetObjectsOk() ([]AzureNativeProtectionGroupObjectParams, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *AzureNativeProtectionGroupParams) HasObjects() bool {
	if o != nil && !IsNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []AzureNativeProtectionGroupObjectParams and assigns it to the Objects field.
func (o *AzureNativeProtectionGroupParams) SetObjects(v []AzureNativeProtectionGroupObjectParams) {
	o.Objects = v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureNativeProtectionGroupParams) GetSourceId() int64 {
	if o == nil || IsNil(o.SourceId.Get()) {
		var ret int64
		return ret
	}
	return *o.SourceId.Get()
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureNativeProtectionGroupParams) GetSourceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceId.Get(), o.SourceId.IsSet()
}

// HasSourceId returns a boolean if a field has been set.
func (o *AzureNativeProtectionGroupParams) HasSourceId() bool {
	if o != nil && o.SourceId.IsSet() {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given NullableInt64 and assigns it to the SourceId field.
func (o *AzureNativeProtectionGroupParams) SetSourceId(v int64) {
	o.SourceId.Set(&v)
}
// SetSourceIdNil sets the value for SourceId to be an explicit nil
func (o *AzureNativeProtectionGroupParams) SetSourceIdNil() {
	o.SourceId.Set(nil)
}

// UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
func (o *AzureNativeProtectionGroupParams) UnsetSourceId() {
	o.SourceId.Unset()
}

// GetSourceName returns the SourceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureNativeProtectionGroupParams) GetSourceName() string {
	if o == nil || IsNil(o.SourceName.Get()) {
		var ret string
		return ret
	}
	return *o.SourceName.Get()
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureNativeProtectionGroupParams) GetSourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceName.Get(), o.SourceName.IsSet()
}

// HasSourceName returns a boolean if a field has been set.
func (o *AzureNativeProtectionGroupParams) HasSourceName() bool {
	if o != nil && o.SourceName.IsSet() {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given NullableString and assigns it to the SourceName field.
func (o *AzureNativeProtectionGroupParams) SetSourceName(v string) {
	o.SourceName.Set(&v)
}
// SetSourceNameNil sets the value for SourceName to be an explicit nil
func (o *AzureNativeProtectionGroupParams) SetSourceNameNil() {
	o.SourceName.Set(nil)
}

// UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
func (o *AzureNativeProtectionGroupParams) UnsetSourceName() {
	o.SourceName.Unset()
}

// GetVmTagIds returns the VmTagIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureNativeProtectionGroupParams) GetVmTagIds() [][]int64 {
	if o == nil {
		var ret [][]int64
		return ret
	}
	return o.VmTagIds
}

// GetVmTagIdsOk returns a tuple with the VmTagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureNativeProtectionGroupParams) GetVmTagIdsOk() ([][]int64, bool) {
	if o == nil || IsNil(o.VmTagIds) {
		return nil, false
	}
	return o.VmTagIds, true
}

// HasVmTagIds returns a boolean if a field has been set.
func (o *AzureNativeProtectionGroupParams) HasVmTagIds() bool {
	if o != nil && !IsNil(o.VmTagIds) {
		return true
	}

	return false
}

// SetVmTagIds gets a reference to the given [][]int64 and assigns it to the VmTagIds field.
func (o *AzureNativeProtectionGroupParams) SetVmTagIds(v [][]int64) {
	o.VmTagIds = v
}

func (o AzureNativeProtectionGroupParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureNativeProtectionGroupParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CloudMigration.IsSet() {
		toSerialize["cloudMigration"] = o.CloudMigration.Get()
	}
	if !IsNil(o.CloudPrePostScript) {
		toSerialize["cloudPrePostScript"] = o.CloudPrePostScript
	}
	if !IsNil(o.DataTransferInfo) {
		toSerialize["dataTransferInfo"] = o.DataTransferInfo
	}
	if !IsNil(o.ExcludeObjectIds) {
		toSerialize["excludeObjectIds"] = o.ExcludeObjectIds
	}
	if o.ExcludeVmTagIds != nil {
		toSerialize["excludeVmTagIds"] = o.ExcludeVmTagIds
	}
	if !IsNil(o.IndexingPolicy) {
		toSerialize["indexingPolicy"] = o.IndexingPolicy
	}
	if !IsNil(o.Objects) {
		toSerialize["objects"] = o.Objects
	}
	if o.SourceId.IsSet() {
		toSerialize["sourceId"] = o.SourceId.Get()
	}
	if o.SourceName.IsSet() {
		toSerialize["sourceName"] = o.SourceName.Get()
	}
	if o.VmTagIds != nil {
		toSerialize["vmTagIds"] = o.VmTagIds
	}
	return toSerialize, nil
}

type NullableAzureNativeProtectionGroupParams struct {
	value *AzureNativeProtectionGroupParams
	isSet bool
}

func (v NullableAzureNativeProtectionGroupParams) Get() *AzureNativeProtectionGroupParams {
	return v.value
}

func (v *NullableAzureNativeProtectionGroupParams) Set(val *AzureNativeProtectionGroupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureNativeProtectionGroupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureNativeProtectionGroupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureNativeProtectionGroupParams(val *AzureNativeProtectionGroupParams) *NullableAzureNativeProtectionGroupParams {
	return &NullableAzureNativeProtectionGroupParams{value: val, isSet: true}
}

func (v NullableAzureNativeProtectionGroupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureNativeProtectionGroupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


