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

// checks if the S3Config type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &S3Config{}

// S3Config Specifies the S3 config settings for this View.
type S3Config struct {
	AclConfig *S3ConfigAclConfig `json:"aclConfig,omitempty"`
	BucketPolicy *S3ConfigBucketPolicy `json:"bucketPolicy,omitempty"`
	// Specifies if this View has S3 ABAC enabled. This can only be set while creating a view. The ABAC server corresponding the tenant will be used for authentication and authorization checks. 
	EnableAbac NullableBool `json:"enableAbac,omitempty"`
	LifecycleManagement *S3ConfigLifecycleManagement `json:"lifecycleManagement,omitempty"`
	OwnerInfo *S3ConfigOwnerInfo `json:"ownerInfo,omitempty"`
	// Specifies the path to access this View as an S3 share.
	S3AccessPath NullableString `json:"s3AccessPath,omitempty"`
	// Specifies if this View has S3 MPU 2.0 enabled. This can set while editing a view. 
	S3EfficientMpuMaxSubfiles NullableInt32 `json:"s3EfficientMpuMaxSubfiles,omitempty"`
	// Specifies if this View has S3 MPU 2.0 enabled. This can set while editing a view. 
	S3EnableEfficientMpu NullableBool `json:"s3EnableEfficientMpu,omitempty"`
	// Specifies the S3 migration action to be performed on this View. Supported migration actions are: [Enable, Cancel, Pause, Resume].
	S3MigrationAction NullableString `json:"s3MigrationAction,omitempty"`
	// Specifies the current S3 migration state for this View. A View can be under following migration states: [Eligible, Enable, Pause, Complete, UnderMigration].
	S3MigrationState NullableString `json:"s3MigrationState,omitempty"`
	// Specifies the versioning state of S3 bucket. Buckets can be in one of three states: UnVersioned (default), VersioningEnabled, or VersioningSuspended. Once versioning is enabled for a bucket, it can never return to an UnVersioned state. However, versioning on the bucket can be suspended.
	Versioning NullableString `json:"versioning,omitempty"`
}

// NewS3Config instantiates a new S3Config object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3Config() *S3Config {
	this := S3Config{}
	return &this
}

// NewS3ConfigWithDefaults instantiates a new S3Config object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3ConfigWithDefaults() *S3Config {
	this := S3Config{}
	return &this
}

// GetAclConfig returns the AclConfig field value if set, zero value otherwise.
func (o *S3Config) GetAclConfig() S3ConfigAclConfig {
	if o == nil || IsNil(o.AclConfig) {
		var ret S3ConfigAclConfig
		return ret
	}
	return *o.AclConfig
}

// GetAclConfigOk returns a tuple with the AclConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Config) GetAclConfigOk() (*S3ConfigAclConfig, bool) {
	if o == nil || IsNil(o.AclConfig) {
		return nil, false
	}
	return o.AclConfig, true
}

// HasAclConfig returns a boolean if a field has been set.
func (o *S3Config) HasAclConfig() bool {
	if o != nil && !IsNil(o.AclConfig) {
		return true
	}

	return false
}

// SetAclConfig gets a reference to the given S3ConfigAclConfig and assigns it to the AclConfig field.
func (o *S3Config) SetAclConfig(v S3ConfigAclConfig) {
	o.AclConfig = &v
}

// GetBucketPolicy returns the BucketPolicy field value if set, zero value otherwise.
func (o *S3Config) GetBucketPolicy() S3ConfigBucketPolicy {
	if o == nil || IsNil(o.BucketPolicy) {
		var ret S3ConfigBucketPolicy
		return ret
	}
	return *o.BucketPolicy
}

// GetBucketPolicyOk returns a tuple with the BucketPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Config) GetBucketPolicyOk() (*S3ConfigBucketPolicy, bool) {
	if o == nil || IsNil(o.BucketPolicy) {
		return nil, false
	}
	return o.BucketPolicy, true
}

// HasBucketPolicy returns a boolean if a field has been set.
func (o *S3Config) HasBucketPolicy() bool {
	if o != nil && !IsNil(o.BucketPolicy) {
		return true
	}

	return false
}

// SetBucketPolicy gets a reference to the given S3ConfigBucketPolicy and assigns it to the BucketPolicy field.
func (o *S3Config) SetBucketPolicy(v S3ConfigBucketPolicy) {
	o.BucketPolicy = &v
}

// GetEnableAbac returns the EnableAbac field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3Config) GetEnableAbac() bool {
	if o == nil || IsNil(o.EnableAbac.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableAbac.Get()
}

// GetEnableAbacOk returns a tuple with the EnableAbac field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3Config) GetEnableAbacOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableAbac.Get(), o.EnableAbac.IsSet()
}

// HasEnableAbac returns a boolean if a field has been set.
func (o *S3Config) HasEnableAbac() bool {
	if o != nil && o.EnableAbac.IsSet() {
		return true
	}

	return false
}

// SetEnableAbac gets a reference to the given NullableBool and assigns it to the EnableAbac field.
func (o *S3Config) SetEnableAbac(v bool) {
	o.EnableAbac.Set(&v)
}
// SetEnableAbacNil sets the value for EnableAbac to be an explicit nil
func (o *S3Config) SetEnableAbacNil() {
	o.EnableAbac.Set(nil)
}

// UnsetEnableAbac ensures that no value is present for EnableAbac, not even an explicit nil
func (o *S3Config) UnsetEnableAbac() {
	o.EnableAbac.Unset()
}

// GetLifecycleManagement returns the LifecycleManagement field value if set, zero value otherwise.
func (o *S3Config) GetLifecycleManagement() S3ConfigLifecycleManagement {
	if o == nil || IsNil(o.LifecycleManagement) {
		var ret S3ConfigLifecycleManagement
		return ret
	}
	return *o.LifecycleManagement
}

// GetLifecycleManagementOk returns a tuple with the LifecycleManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Config) GetLifecycleManagementOk() (*S3ConfigLifecycleManagement, bool) {
	if o == nil || IsNil(o.LifecycleManagement) {
		return nil, false
	}
	return o.LifecycleManagement, true
}

// HasLifecycleManagement returns a boolean if a field has been set.
func (o *S3Config) HasLifecycleManagement() bool {
	if o != nil && !IsNil(o.LifecycleManagement) {
		return true
	}

	return false
}

// SetLifecycleManagement gets a reference to the given S3ConfigLifecycleManagement and assigns it to the LifecycleManagement field.
func (o *S3Config) SetLifecycleManagement(v S3ConfigLifecycleManagement) {
	o.LifecycleManagement = &v
}

// GetOwnerInfo returns the OwnerInfo field value if set, zero value otherwise.
func (o *S3Config) GetOwnerInfo() S3ConfigOwnerInfo {
	if o == nil || IsNil(o.OwnerInfo) {
		var ret S3ConfigOwnerInfo
		return ret
	}
	return *o.OwnerInfo
}

// GetOwnerInfoOk returns a tuple with the OwnerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3Config) GetOwnerInfoOk() (*S3ConfigOwnerInfo, bool) {
	if o == nil || IsNil(o.OwnerInfo) {
		return nil, false
	}
	return o.OwnerInfo, true
}

// HasOwnerInfo returns a boolean if a field has been set.
func (o *S3Config) HasOwnerInfo() bool {
	if o != nil && !IsNil(o.OwnerInfo) {
		return true
	}

	return false
}

// SetOwnerInfo gets a reference to the given S3ConfigOwnerInfo and assigns it to the OwnerInfo field.
func (o *S3Config) SetOwnerInfo(v S3ConfigOwnerInfo) {
	o.OwnerInfo = &v
}

// GetS3AccessPath returns the S3AccessPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3Config) GetS3AccessPath() string {
	if o == nil || IsNil(o.S3AccessPath.Get()) {
		var ret string
		return ret
	}
	return *o.S3AccessPath.Get()
}

// GetS3AccessPathOk returns a tuple with the S3AccessPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3Config) GetS3AccessPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.S3AccessPath.Get(), o.S3AccessPath.IsSet()
}

// HasS3AccessPath returns a boolean if a field has been set.
func (o *S3Config) HasS3AccessPath() bool {
	if o != nil && o.S3AccessPath.IsSet() {
		return true
	}

	return false
}

// SetS3AccessPath gets a reference to the given NullableString and assigns it to the S3AccessPath field.
func (o *S3Config) SetS3AccessPath(v string) {
	o.S3AccessPath.Set(&v)
}
// SetS3AccessPathNil sets the value for S3AccessPath to be an explicit nil
func (o *S3Config) SetS3AccessPathNil() {
	o.S3AccessPath.Set(nil)
}

// UnsetS3AccessPath ensures that no value is present for S3AccessPath, not even an explicit nil
func (o *S3Config) UnsetS3AccessPath() {
	o.S3AccessPath.Unset()
}

// GetS3EfficientMpuMaxSubfiles returns the S3EfficientMpuMaxSubfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3Config) GetS3EfficientMpuMaxSubfiles() int32 {
	if o == nil || IsNil(o.S3EfficientMpuMaxSubfiles.Get()) {
		var ret int32
		return ret
	}
	return *o.S3EfficientMpuMaxSubfiles.Get()
}

// GetS3EfficientMpuMaxSubfilesOk returns a tuple with the S3EfficientMpuMaxSubfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3Config) GetS3EfficientMpuMaxSubfilesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.S3EfficientMpuMaxSubfiles.Get(), o.S3EfficientMpuMaxSubfiles.IsSet()
}

// HasS3EfficientMpuMaxSubfiles returns a boolean if a field has been set.
func (o *S3Config) HasS3EfficientMpuMaxSubfiles() bool {
	if o != nil && o.S3EfficientMpuMaxSubfiles.IsSet() {
		return true
	}

	return false
}

// SetS3EfficientMpuMaxSubfiles gets a reference to the given NullableInt32 and assigns it to the S3EfficientMpuMaxSubfiles field.
func (o *S3Config) SetS3EfficientMpuMaxSubfiles(v int32) {
	o.S3EfficientMpuMaxSubfiles.Set(&v)
}
// SetS3EfficientMpuMaxSubfilesNil sets the value for S3EfficientMpuMaxSubfiles to be an explicit nil
func (o *S3Config) SetS3EfficientMpuMaxSubfilesNil() {
	o.S3EfficientMpuMaxSubfiles.Set(nil)
}

// UnsetS3EfficientMpuMaxSubfiles ensures that no value is present for S3EfficientMpuMaxSubfiles, not even an explicit nil
func (o *S3Config) UnsetS3EfficientMpuMaxSubfiles() {
	o.S3EfficientMpuMaxSubfiles.Unset()
}

// GetS3EnableEfficientMpu returns the S3EnableEfficientMpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3Config) GetS3EnableEfficientMpu() bool {
	if o == nil || IsNil(o.S3EnableEfficientMpu.Get()) {
		var ret bool
		return ret
	}
	return *o.S3EnableEfficientMpu.Get()
}

// GetS3EnableEfficientMpuOk returns a tuple with the S3EnableEfficientMpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3Config) GetS3EnableEfficientMpuOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.S3EnableEfficientMpu.Get(), o.S3EnableEfficientMpu.IsSet()
}

// HasS3EnableEfficientMpu returns a boolean if a field has been set.
func (o *S3Config) HasS3EnableEfficientMpu() bool {
	if o != nil && o.S3EnableEfficientMpu.IsSet() {
		return true
	}

	return false
}

// SetS3EnableEfficientMpu gets a reference to the given NullableBool and assigns it to the S3EnableEfficientMpu field.
func (o *S3Config) SetS3EnableEfficientMpu(v bool) {
	o.S3EnableEfficientMpu.Set(&v)
}
// SetS3EnableEfficientMpuNil sets the value for S3EnableEfficientMpu to be an explicit nil
func (o *S3Config) SetS3EnableEfficientMpuNil() {
	o.S3EnableEfficientMpu.Set(nil)
}

// UnsetS3EnableEfficientMpu ensures that no value is present for S3EnableEfficientMpu, not even an explicit nil
func (o *S3Config) UnsetS3EnableEfficientMpu() {
	o.S3EnableEfficientMpu.Unset()
}

// GetS3MigrationAction returns the S3MigrationAction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3Config) GetS3MigrationAction() string {
	if o == nil || IsNil(o.S3MigrationAction.Get()) {
		var ret string
		return ret
	}
	return *o.S3MigrationAction.Get()
}

// GetS3MigrationActionOk returns a tuple with the S3MigrationAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3Config) GetS3MigrationActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.S3MigrationAction.Get(), o.S3MigrationAction.IsSet()
}

// HasS3MigrationAction returns a boolean if a field has been set.
func (o *S3Config) HasS3MigrationAction() bool {
	if o != nil && o.S3MigrationAction.IsSet() {
		return true
	}

	return false
}

// SetS3MigrationAction gets a reference to the given NullableString and assigns it to the S3MigrationAction field.
func (o *S3Config) SetS3MigrationAction(v string) {
	o.S3MigrationAction.Set(&v)
}
// SetS3MigrationActionNil sets the value for S3MigrationAction to be an explicit nil
func (o *S3Config) SetS3MigrationActionNil() {
	o.S3MigrationAction.Set(nil)
}

// UnsetS3MigrationAction ensures that no value is present for S3MigrationAction, not even an explicit nil
func (o *S3Config) UnsetS3MigrationAction() {
	o.S3MigrationAction.Unset()
}

// GetS3MigrationState returns the S3MigrationState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3Config) GetS3MigrationState() string {
	if o == nil || IsNil(o.S3MigrationState.Get()) {
		var ret string
		return ret
	}
	return *o.S3MigrationState.Get()
}

// GetS3MigrationStateOk returns a tuple with the S3MigrationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3Config) GetS3MigrationStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.S3MigrationState.Get(), o.S3MigrationState.IsSet()
}

// HasS3MigrationState returns a boolean if a field has been set.
func (o *S3Config) HasS3MigrationState() bool {
	if o != nil && o.S3MigrationState.IsSet() {
		return true
	}

	return false
}

// SetS3MigrationState gets a reference to the given NullableString and assigns it to the S3MigrationState field.
func (o *S3Config) SetS3MigrationState(v string) {
	o.S3MigrationState.Set(&v)
}
// SetS3MigrationStateNil sets the value for S3MigrationState to be an explicit nil
func (o *S3Config) SetS3MigrationStateNil() {
	o.S3MigrationState.Set(nil)
}

// UnsetS3MigrationState ensures that no value is present for S3MigrationState, not even an explicit nil
func (o *S3Config) UnsetS3MigrationState() {
	o.S3MigrationState.Unset()
}

// GetVersioning returns the Versioning field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3Config) GetVersioning() string {
	if o == nil || IsNil(o.Versioning.Get()) {
		var ret string
		return ret
	}
	return *o.Versioning.Get()
}

// GetVersioningOk returns a tuple with the Versioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3Config) GetVersioningOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Versioning.Get(), o.Versioning.IsSet()
}

// HasVersioning returns a boolean if a field has been set.
func (o *S3Config) HasVersioning() bool {
	if o != nil && o.Versioning.IsSet() {
		return true
	}

	return false
}

// SetVersioning gets a reference to the given NullableString and assigns it to the Versioning field.
func (o *S3Config) SetVersioning(v string) {
	o.Versioning.Set(&v)
}
// SetVersioningNil sets the value for Versioning to be an explicit nil
func (o *S3Config) SetVersioningNil() {
	o.Versioning.Set(nil)
}

// UnsetVersioning ensures that no value is present for Versioning, not even an explicit nil
func (o *S3Config) UnsetVersioning() {
	o.Versioning.Unset()
}

func (o S3Config) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o S3Config) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AclConfig) {
		toSerialize["aclConfig"] = o.AclConfig
	}
	if !IsNil(o.BucketPolicy) {
		toSerialize["bucketPolicy"] = o.BucketPolicy
	}
	if o.EnableAbac.IsSet() {
		toSerialize["enableAbac"] = o.EnableAbac.Get()
	}
	if !IsNil(o.LifecycleManagement) {
		toSerialize["lifecycleManagement"] = o.LifecycleManagement
	}
	if !IsNil(o.OwnerInfo) {
		toSerialize["ownerInfo"] = o.OwnerInfo
	}
	if o.S3AccessPath.IsSet() {
		toSerialize["s3AccessPath"] = o.S3AccessPath.Get()
	}
	if o.S3EfficientMpuMaxSubfiles.IsSet() {
		toSerialize["s3EfficientMpuMaxSubfiles"] = o.S3EfficientMpuMaxSubfiles.Get()
	}
	if o.S3EnableEfficientMpu.IsSet() {
		toSerialize["s3EnableEfficientMpu"] = o.S3EnableEfficientMpu.Get()
	}
	if o.S3MigrationAction.IsSet() {
		toSerialize["s3MigrationAction"] = o.S3MigrationAction.Get()
	}
	if o.S3MigrationState.IsSet() {
		toSerialize["s3MigrationState"] = o.S3MigrationState.Get()
	}
	if o.Versioning.IsSet() {
		toSerialize["versioning"] = o.Versioning.Get()
	}
	return toSerialize, nil
}

type NullableS3Config struct {
	value *S3Config
	isSet bool
}

func (v NullableS3Config) Get() *S3Config {
	return v.value
}

func (v *NullableS3Config) Set(val *S3Config) {
	v.value = val
	v.isSet = true
}

func (v NullableS3Config) IsSet() bool {
	return v.isSet
}

func (v *NullableS3Config) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3Config(val *S3Config) *NullableS3Config {
	return &NullableS3Config{value: val, isSet: true}
}

func (v NullableS3Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3Config) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


