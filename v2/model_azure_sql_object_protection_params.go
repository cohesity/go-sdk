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

// checks if the AzureSqlObjectProtectionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureSqlObjectProtectionParams{}

// AzureSqlObjectProtectionParams Specifies the parameters which are specific to Azure SQL Object Protection Groups using Azure native APIs. Atlease one of objects must be specified.
type AzureSqlObjectProtectionParams struct {
	// If set to true, a copy of the database is created during backup, and the backup is performed from the copied database. This backup will be transactionally consistent. If set to false, the backup is performed from the production database while transactions are in progress. In this case, the backup will be transactionally inconsistent, and recovery can fail or the recovered database may be in an inconsistent state.
	CopyDatabase NullableBool `json:"copyDatabase,omitempty"`
	CopyDatabaseSku *AzureSqlSkuOptions `json:"copyDatabaseSku,omitempty"`
	// Specifies azure managed storage disk to be used for object protection. By default Premium LRS is being used to support Azure SQL workloads.
	DiskType NullableString `json:"diskType,omitempty"`
	// Specifies the ids of the objects to be excluded in the Object Protection. This can be used to ignore specific objects under a parent object which has been included for protection.
	ExcludeObjectIds []int64 `json:"excludeObjectIds,omitempty"`
	// Specifies the objects to be protected.
	Objects []AzureObjectLevelParams `json:"objects,omitempty"`
	SqlPackageOptions *AzureSqlPackageOptions `json:"sqlPackageOptions,omitempty"`
	// Specifies the size of the disk we will attach to rigel to use for exporting this DB(in GB).
	TempDiskSizeGb NullableInt32 `json:"tempDiskSizeGb,omitempty"`
}

// NewAzureSqlObjectProtectionParams instantiates a new AzureSqlObjectProtectionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureSqlObjectProtectionParams() *AzureSqlObjectProtectionParams {
	this := AzureSqlObjectProtectionParams{}
	return &this
}

// NewAzureSqlObjectProtectionParamsWithDefaults instantiates a new AzureSqlObjectProtectionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureSqlObjectProtectionParamsWithDefaults() *AzureSqlObjectProtectionParams {
	this := AzureSqlObjectProtectionParams{}
	return &this
}

// GetCopyDatabase returns the CopyDatabase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureSqlObjectProtectionParams) GetCopyDatabase() bool {
	if o == nil || IsNil(o.CopyDatabase.Get()) {
		var ret bool
		return ret
	}
	return *o.CopyDatabase.Get()
}

// GetCopyDatabaseOk returns a tuple with the CopyDatabase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureSqlObjectProtectionParams) GetCopyDatabaseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CopyDatabase.Get(), o.CopyDatabase.IsSet()
}

// HasCopyDatabase returns a boolean if a field has been set.
func (o *AzureSqlObjectProtectionParams) HasCopyDatabase() bool {
	if o != nil && o.CopyDatabase.IsSet() {
		return true
	}

	return false
}

// SetCopyDatabase gets a reference to the given NullableBool and assigns it to the CopyDatabase field.
func (o *AzureSqlObjectProtectionParams) SetCopyDatabase(v bool) {
	o.CopyDatabase.Set(&v)
}
// SetCopyDatabaseNil sets the value for CopyDatabase to be an explicit nil
func (o *AzureSqlObjectProtectionParams) SetCopyDatabaseNil() {
	o.CopyDatabase.Set(nil)
}

// UnsetCopyDatabase ensures that no value is present for CopyDatabase, not even an explicit nil
func (o *AzureSqlObjectProtectionParams) UnsetCopyDatabase() {
	o.CopyDatabase.Unset()
}

// GetCopyDatabaseSku returns the CopyDatabaseSku field value if set, zero value otherwise.
func (o *AzureSqlObjectProtectionParams) GetCopyDatabaseSku() AzureSqlSkuOptions {
	if o == nil || IsNil(o.CopyDatabaseSku) {
		var ret AzureSqlSkuOptions
		return ret
	}
	return *o.CopyDatabaseSku
}

// GetCopyDatabaseSkuOk returns a tuple with the CopyDatabaseSku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureSqlObjectProtectionParams) GetCopyDatabaseSkuOk() (*AzureSqlSkuOptions, bool) {
	if o == nil || IsNil(o.CopyDatabaseSku) {
		return nil, false
	}
	return o.CopyDatabaseSku, true
}

// HasCopyDatabaseSku returns a boolean if a field has been set.
func (o *AzureSqlObjectProtectionParams) HasCopyDatabaseSku() bool {
	if o != nil && !IsNil(o.CopyDatabaseSku) {
		return true
	}

	return false
}

// SetCopyDatabaseSku gets a reference to the given AzureSqlSkuOptions and assigns it to the CopyDatabaseSku field.
func (o *AzureSqlObjectProtectionParams) SetCopyDatabaseSku(v AzureSqlSkuOptions) {
	o.CopyDatabaseSku = &v
}

// GetDiskType returns the DiskType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureSqlObjectProtectionParams) GetDiskType() string {
	if o == nil || IsNil(o.DiskType.Get()) {
		var ret string
		return ret
	}
	return *o.DiskType.Get()
}

// GetDiskTypeOk returns a tuple with the DiskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureSqlObjectProtectionParams) GetDiskTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiskType.Get(), o.DiskType.IsSet()
}

// HasDiskType returns a boolean if a field has been set.
func (o *AzureSqlObjectProtectionParams) HasDiskType() bool {
	if o != nil && o.DiskType.IsSet() {
		return true
	}

	return false
}

// SetDiskType gets a reference to the given NullableString and assigns it to the DiskType field.
func (o *AzureSqlObjectProtectionParams) SetDiskType(v string) {
	o.DiskType.Set(&v)
}
// SetDiskTypeNil sets the value for DiskType to be an explicit nil
func (o *AzureSqlObjectProtectionParams) SetDiskTypeNil() {
	o.DiskType.Set(nil)
}

// UnsetDiskType ensures that no value is present for DiskType, not even an explicit nil
func (o *AzureSqlObjectProtectionParams) UnsetDiskType() {
	o.DiskType.Unset()
}

// GetExcludeObjectIds returns the ExcludeObjectIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureSqlObjectProtectionParams) GetExcludeObjectIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.ExcludeObjectIds
}

// GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureSqlObjectProtectionParams) GetExcludeObjectIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.ExcludeObjectIds) {
		return nil, false
	}
	return o.ExcludeObjectIds, true
}

// HasExcludeObjectIds returns a boolean if a field has been set.
func (o *AzureSqlObjectProtectionParams) HasExcludeObjectIds() bool {
	if o != nil && !IsNil(o.ExcludeObjectIds) {
		return true
	}

	return false
}

// SetExcludeObjectIds gets a reference to the given []int64 and assigns it to the ExcludeObjectIds field.
func (o *AzureSqlObjectProtectionParams) SetExcludeObjectIds(v []int64) {
	o.ExcludeObjectIds = v
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *AzureSqlObjectProtectionParams) GetObjects() []AzureObjectLevelParams {
	if o == nil || IsNil(o.Objects) {
		var ret []AzureObjectLevelParams
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureSqlObjectProtectionParams) GetObjectsOk() ([]AzureObjectLevelParams, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *AzureSqlObjectProtectionParams) HasObjects() bool {
	if o != nil && !IsNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []AzureObjectLevelParams and assigns it to the Objects field.
func (o *AzureSqlObjectProtectionParams) SetObjects(v []AzureObjectLevelParams) {
	o.Objects = v
}

// GetSqlPackageOptions returns the SqlPackageOptions field value if set, zero value otherwise.
func (o *AzureSqlObjectProtectionParams) GetSqlPackageOptions() AzureSqlPackageOptions {
	if o == nil || IsNil(o.SqlPackageOptions) {
		var ret AzureSqlPackageOptions
		return ret
	}
	return *o.SqlPackageOptions
}

// GetSqlPackageOptionsOk returns a tuple with the SqlPackageOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureSqlObjectProtectionParams) GetSqlPackageOptionsOk() (*AzureSqlPackageOptions, bool) {
	if o == nil || IsNil(o.SqlPackageOptions) {
		return nil, false
	}
	return o.SqlPackageOptions, true
}

// HasSqlPackageOptions returns a boolean if a field has been set.
func (o *AzureSqlObjectProtectionParams) HasSqlPackageOptions() bool {
	if o != nil && !IsNil(o.SqlPackageOptions) {
		return true
	}

	return false
}

// SetSqlPackageOptions gets a reference to the given AzureSqlPackageOptions and assigns it to the SqlPackageOptions field.
func (o *AzureSqlObjectProtectionParams) SetSqlPackageOptions(v AzureSqlPackageOptions) {
	o.SqlPackageOptions = &v
}

// GetTempDiskSizeGb returns the TempDiskSizeGb field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureSqlObjectProtectionParams) GetTempDiskSizeGb() int32 {
	if o == nil || IsNil(o.TempDiskSizeGb.Get()) {
		var ret int32
		return ret
	}
	return *o.TempDiskSizeGb.Get()
}

// GetTempDiskSizeGbOk returns a tuple with the TempDiskSizeGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureSqlObjectProtectionParams) GetTempDiskSizeGbOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TempDiskSizeGb.Get(), o.TempDiskSizeGb.IsSet()
}

// HasTempDiskSizeGb returns a boolean if a field has been set.
func (o *AzureSqlObjectProtectionParams) HasTempDiskSizeGb() bool {
	if o != nil && o.TempDiskSizeGb.IsSet() {
		return true
	}

	return false
}

// SetTempDiskSizeGb gets a reference to the given NullableInt32 and assigns it to the TempDiskSizeGb field.
func (o *AzureSqlObjectProtectionParams) SetTempDiskSizeGb(v int32) {
	o.TempDiskSizeGb.Set(&v)
}
// SetTempDiskSizeGbNil sets the value for TempDiskSizeGb to be an explicit nil
func (o *AzureSqlObjectProtectionParams) SetTempDiskSizeGbNil() {
	o.TempDiskSizeGb.Set(nil)
}

// UnsetTempDiskSizeGb ensures that no value is present for TempDiskSizeGb, not even an explicit nil
func (o *AzureSqlObjectProtectionParams) UnsetTempDiskSizeGb() {
	o.TempDiskSizeGb.Unset()
}

func (o AzureSqlObjectProtectionParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureSqlObjectProtectionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CopyDatabase.IsSet() {
		toSerialize["copyDatabase"] = o.CopyDatabase.Get()
	}
	if !IsNil(o.CopyDatabaseSku) {
		toSerialize["copyDatabaseSku"] = o.CopyDatabaseSku
	}
	if o.DiskType.IsSet() {
		toSerialize["diskType"] = o.DiskType.Get()
	}
	if o.ExcludeObjectIds != nil {
		toSerialize["excludeObjectIds"] = o.ExcludeObjectIds
	}
	if !IsNil(o.Objects) {
		toSerialize["objects"] = o.Objects
	}
	if !IsNil(o.SqlPackageOptions) {
		toSerialize["sqlPackageOptions"] = o.SqlPackageOptions
	}
	if o.TempDiskSizeGb.IsSet() {
		toSerialize["tempDiskSizeGb"] = o.TempDiskSizeGb.Get()
	}
	return toSerialize, nil
}

type NullableAzureSqlObjectProtectionParams struct {
	value *AzureSqlObjectProtectionParams
	isSet bool
}

func (v NullableAzureSqlObjectProtectionParams) Get() *AzureSqlObjectProtectionParams {
	return v.value
}

func (v *NullableAzureSqlObjectProtectionParams) Set(val *AzureSqlObjectProtectionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureSqlObjectProtectionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureSqlObjectProtectionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureSqlObjectProtectionParams(val *AzureSqlObjectProtectionParams) *NullableAzureSqlObjectProtectionParams {
	return &NullableAzureSqlObjectProtectionParams{value: val, isSet: true}
}

func (v NullableAzureSqlObjectProtectionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureSqlObjectProtectionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


