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

// checks if the FileLevelDataLockConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileLevelDataLockConfig{}

// FileLevelDataLockConfig Specifies a config to lock files in a view - to protect from malicious or an accidental attempt to delete or modify the files in this view.
type FileLevelDataLockConfig struct {
	// Specifies the duration to lock a file that has not been accessed or modified (ie. has been idle) for a certain duration of time in milliseconds. Do not set if it is required to disable auto lock.
	AutoLockAfterDurationIdleMsecs NullableInt32 `json:"autoLockAfterDurationIdleMsecs,omitempty"`
	// Specified if files in the View can be locked in different modes. This property is immutable and can only be set when enabling File level datalock. If this property is set for an S3 View, S3 bucket Versioning should also be enabled.
	CoexistingLockMode NullableBool `json:"coexistingLockMode,omitempty"`
	// Specifies a global default retention duration for files in this view, if file lock is enabled for this view. Also, it is a required field if file lock is enabled. Set to -1 if the required default retention period is forever.
	DefaultRetentionDurationMsecs NullableInt64 `json:"defaultRetentionDurationMsecs,omitempty"`
	// Specifies a global default retention duration in years for files in this view, if file/object lock is enabled for this view.
	DefaultRetentionDurationYears NullableInt64 `json:"defaultRetentionDurationYears,omitempty"`
	// Specifies a definite timestamp in milliseconds for retaining the file.
	ExpiryTimestampMsecs NullableInt32 `json:"expiryTimestampMsecs,omitempty"`
	// Specifies the supported mechanisms to explicity lock a file from NFS/SMB interface. Supported locking protocols: SetReadOnly, SetAtime. 'SetReadOnly' is compatible with Isilon/Netapp behaviour. This locks the file and the retention duration is determined in this order: 1) atime, if set by user/application and within min and max retention duration. 2) Min retention duration, if set. 3) Otherwise, file is switched to expired data automatically. 'SetAtime' is compatible with Data Domain behaviour.
	LockingProtocol NullableString `json:"lockingProtocol,omitempty"`
	// Specifies a maximum duration in milliseconds for which any file in this view can be retained for. Set to -1 if the required retention duration is forever. If set, it should be greater than or equal to the default retention period as well as the min retention period.
	MaxRetentionDurationMsecs NullableInt64 `json:"maxRetentionDurationMsecs,omitempty"`
	// Specifies a minimum retention duration in milliseconds after a file gets locked. The file cannot be modified or deleted during this timeframe. Set to -1 if the required retention duration is forever. This should be set less than or equal to the default retention duration.
	MinRetentionDurationMsecs NullableInt64 `json:"minRetentionDurationMsecs,omitempty"`
	// Specifies the mode of file level datalock. Enterprise mode can be upgraded to Compliance mode, but Compliance mode cannot be downgraded to Enterprise mode. Compliance: This mode would disallow all user to delete/modify file or view under any condition when it 's in locked status except for deleting view when the view is empty. Enterprise: This mode would follow the rules as compliance mode for normal users. But it would allow the storage admin (1) to delete view or file anytime no matter it is in locked status or expired. (2) to rename the view (3) to bring back the retention period when it's in locked mode A lock mode of a file in a view can be in one of the following: 'Compliance': Default mode of datalock, in this mode, Data Security Admin cannot modify/delete this view when datalock is in effect. Data Security Admin can delete this view when datalock is expired. 'kEnterprise' : In this mode, Data Security Admin can change view name or delete view when datalock is in effect. Datalock in this mode can be upgraded to 'Compliance' mode.
	Mode NullableString `json:"mode,omitempty"`
}

// NewFileLevelDataLockConfig instantiates a new FileLevelDataLockConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileLevelDataLockConfig() *FileLevelDataLockConfig {
	this := FileLevelDataLockConfig{}
	return &this
}

// NewFileLevelDataLockConfigWithDefaults instantiates a new FileLevelDataLockConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileLevelDataLockConfigWithDefaults() *FileLevelDataLockConfig {
	this := FileLevelDataLockConfig{}
	return &this
}

// GetAutoLockAfterDurationIdleMsecs returns the AutoLockAfterDurationIdleMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileLevelDataLockConfig) GetAutoLockAfterDurationIdleMsecs() int32 {
	if o == nil || IsNil(o.AutoLockAfterDurationIdleMsecs.Get()) {
		var ret int32
		return ret
	}
	return *o.AutoLockAfterDurationIdleMsecs.Get()
}

// GetAutoLockAfterDurationIdleMsecsOk returns a tuple with the AutoLockAfterDurationIdleMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileLevelDataLockConfig) GetAutoLockAfterDurationIdleMsecsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoLockAfterDurationIdleMsecs.Get(), o.AutoLockAfterDurationIdleMsecs.IsSet()
}

// HasAutoLockAfterDurationIdleMsecs returns a boolean if a field has been set.
func (o *FileLevelDataLockConfig) HasAutoLockAfterDurationIdleMsecs() bool {
	if o != nil && o.AutoLockAfterDurationIdleMsecs.IsSet() {
		return true
	}

	return false
}

// SetAutoLockAfterDurationIdleMsecs gets a reference to the given NullableInt32 and assigns it to the AutoLockAfterDurationIdleMsecs field.
func (o *FileLevelDataLockConfig) SetAutoLockAfterDurationIdleMsecs(v int32) {
	o.AutoLockAfterDurationIdleMsecs.Set(&v)
}
// SetAutoLockAfterDurationIdleMsecsNil sets the value for AutoLockAfterDurationIdleMsecs to be an explicit nil
func (o *FileLevelDataLockConfig) SetAutoLockAfterDurationIdleMsecsNil() {
	o.AutoLockAfterDurationIdleMsecs.Set(nil)
}

// UnsetAutoLockAfterDurationIdleMsecs ensures that no value is present for AutoLockAfterDurationIdleMsecs, not even an explicit nil
func (o *FileLevelDataLockConfig) UnsetAutoLockAfterDurationIdleMsecs() {
	o.AutoLockAfterDurationIdleMsecs.Unset()
}

// GetCoexistingLockMode returns the CoexistingLockMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileLevelDataLockConfig) GetCoexistingLockMode() bool {
	if o == nil || IsNil(o.CoexistingLockMode.Get()) {
		var ret bool
		return ret
	}
	return *o.CoexistingLockMode.Get()
}

// GetCoexistingLockModeOk returns a tuple with the CoexistingLockMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileLevelDataLockConfig) GetCoexistingLockModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CoexistingLockMode.Get(), o.CoexistingLockMode.IsSet()
}

// HasCoexistingLockMode returns a boolean if a field has been set.
func (o *FileLevelDataLockConfig) HasCoexistingLockMode() bool {
	if o != nil && o.CoexistingLockMode.IsSet() {
		return true
	}

	return false
}

// SetCoexistingLockMode gets a reference to the given NullableBool and assigns it to the CoexistingLockMode field.
func (o *FileLevelDataLockConfig) SetCoexistingLockMode(v bool) {
	o.CoexistingLockMode.Set(&v)
}
// SetCoexistingLockModeNil sets the value for CoexistingLockMode to be an explicit nil
func (o *FileLevelDataLockConfig) SetCoexistingLockModeNil() {
	o.CoexistingLockMode.Set(nil)
}

// UnsetCoexistingLockMode ensures that no value is present for CoexistingLockMode, not even an explicit nil
func (o *FileLevelDataLockConfig) UnsetCoexistingLockMode() {
	o.CoexistingLockMode.Unset()
}

// GetDefaultRetentionDurationMsecs returns the DefaultRetentionDurationMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileLevelDataLockConfig) GetDefaultRetentionDurationMsecs() int64 {
	if o == nil || IsNil(o.DefaultRetentionDurationMsecs.Get()) {
		var ret int64
		return ret
	}
	return *o.DefaultRetentionDurationMsecs.Get()
}

// GetDefaultRetentionDurationMsecsOk returns a tuple with the DefaultRetentionDurationMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileLevelDataLockConfig) GetDefaultRetentionDurationMsecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultRetentionDurationMsecs.Get(), o.DefaultRetentionDurationMsecs.IsSet()
}

// HasDefaultRetentionDurationMsecs returns a boolean if a field has been set.
func (o *FileLevelDataLockConfig) HasDefaultRetentionDurationMsecs() bool {
	if o != nil && o.DefaultRetentionDurationMsecs.IsSet() {
		return true
	}

	return false
}

// SetDefaultRetentionDurationMsecs gets a reference to the given NullableInt64 and assigns it to the DefaultRetentionDurationMsecs field.
func (o *FileLevelDataLockConfig) SetDefaultRetentionDurationMsecs(v int64) {
	o.DefaultRetentionDurationMsecs.Set(&v)
}
// SetDefaultRetentionDurationMsecsNil sets the value for DefaultRetentionDurationMsecs to be an explicit nil
func (o *FileLevelDataLockConfig) SetDefaultRetentionDurationMsecsNil() {
	o.DefaultRetentionDurationMsecs.Set(nil)
}

// UnsetDefaultRetentionDurationMsecs ensures that no value is present for DefaultRetentionDurationMsecs, not even an explicit nil
func (o *FileLevelDataLockConfig) UnsetDefaultRetentionDurationMsecs() {
	o.DefaultRetentionDurationMsecs.Unset()
}

// GetDefaultRetentionDurationYears returns the DefaultRetentionDurationYears field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileLevelDataLockConfig) GetDefaultRetentionDurationYears() int64 {
	if o == nil || IsNil(o.DefaultRetentionDurationYears.Get()) {
		var ret int64
		return ret
	}
	return *o.DefaultRetentionDurationYears.Get()
}

// GetDefaultRetentionDurationYearsOk returns a tuple with the DefaultRetentionDurationYears field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileLevelDataLockConfig) GetDefaultRetentionDurationYearsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultRetentionDurationYears.Get(), o.DefaultRetentionDurationYears.IsSet()
}

// HasDefaultRetentionDurationYears returns a boolean if a field has been set.
func (o *FileLevelDataLockConfig) HasDefaultRetentionDurationYears() bool {
	if o != nil && o.DefaultRetentionDurationYears.IsSet() {
		return true
	}

	return false
}

// SetDefaultRetentionDurationYears gets a reference to the given NullableInt64 and assigns it to the DefaultRetentionDurationYears field.
func (o *FileLevelDataLockConfig) SetDefaultRetentionDurationYears(v int64) {
	o.DefaultRetentionDurationYears.Set(&v)
}
// SetDefaultRetentionDurationYearsNil sets the value for DefaultRetentionDurationYears to be an explicit nil
func (o *FileLevelDataLockConfig) SetDefaultRetentionDurationYearsNil() {
	o.DefaultRetentionDurationYears.Set(nil)
}

// UnsetDefaultRetentionDurationYears ensures that no value is present for DefaultRetentionDurationYears, not even an explicit nil
func (o *FileLevelDataLockConfig) UnsetDefaultRetentionDurationYears() {
	o.DefaultRetentionDurationYears.Unset()
}

// GetExpiryTimestampMsecs returns the ExpiryTimestampMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileLevelDataLockConfig) GetExpiryTimestampMsecs() int32 {
	if o == nil || IsNil(o.ExpiryTimestampMsecs.Get()) {
		var ret int32
		return ret
	}
	return *o.ExpiryTimestampMsecs.Get()
}

// GetExpiryTimestampMsecsOk returns a tuple with the ExpiryTimestampMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileLevelDataLockConfig) GetExpiryTimestampMsecsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiryTimestampMsecs.Get(), o.ExpiryTimestampMsecs.IsSet()
}

// HasExpiryTimestampMsecs returns a boolean if a field has been set.
func (o *FileLevelDataLockConfig) HasExpiryTimestampMsecs() bool {
	if o != nil && o.ExpiryTimestampMsecs.IsSet() {
		return true
	}

	return false
}

// SetExpiryTimestampMsecs gets a reference to the given NullableInt32 and assigns it to the ExpiryTimestampMsecs field.
func (o *FileLevelDataLockConfig) SetExpiryTimestampMsecs(v int32) {
	o.ExpiryTimestampMsecs.Set(&v)
}
// SetExpiryTimestampMsecsNil sets the value for ExpiryTimestampMsecs to be an explicit nil
func (o *FileLevelDataLockConfig) SetExpiryTimestampMsecsNil() {
	o.ExpiryTimestampMsecs.Set(nil)
}

// UnsetExpiryTimestampMsecs ensures that no value is present for ExpiryTimestampMsecs, not even an explicit nil
func (o *FileLevelDataLockConfig) UnsetExpiryTimestampMsecs() {
	o.ExpiryTimestampMsecs.Unset()
}

// GetLockingProtocol returns the LockingProtocol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileLevelDataLockConfig) GetLockingProtocol() string {
	if o == nil || IsNil(o.LockingProtocol.Get()) {
		var ret string
		return ret
	}
	return *o.LockingProtocol.Get()
}

// GetLockingProtocolOk returns a tuple with the LockingProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileLevelDataLockConfig) GetLockingProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LockingProtocol.Get(), o.LockingProtocol.IsSet()
}

// HasLockingProtocol returns a boolean if a field has been set.
func (o *FileLevelDataLockConfig) HasLockingProtocol() bool {
	if o != nil && o.LockingProtocol.IsSet() {
		return true
	}

	return false
}

// SetLockingProtocol gets a reference to the given NullableString and assigns it to the LockingProtocol field.
func (o *FileLevelDataLockConfig) SetLockingProtocol(v string) {
	o.LockingProtocol.Set(&v)
}
// SetLockingProtocolNil sets the value for LockingProtocol to be an explicit nil
func (o *FileLevelDataLockConfig) SetLockingProtocolNil() {
	o.LockingProtocol.Set(nil)
}

// UnsetLockingProtocol ensures that no value is present for LockingProtocol, not even an explicit nil
func (o *FileLevelDataLockConfig) UnsetLockingProtocol() {
	o.LockingProtocol.Unset()
}

// GetMaxRetentionDurationMsecs returns the MaxRetentionDurationMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileLevelDataLockConfig) GetMaxRetentionDurationMsecs() int64 {
	if o == nil || IsNil(o.MaxRetentionDurationMsecs.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxRetentionDurationMsecs.Get()
}

// GetMaxRetentionDurationMsecsOk returns a tuple with the MaxRetentionDurationMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileLevelDataLockConfig) GetMaxRetentionDurationMsecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxRetentionDurationMsecs.Get(), o.MaxRetentionDurationMsecs.IsSet()
}

// HasMaxRetentionDurationMsecs returns a boolean if a field has been set.
func (o *FileLevelDataLockConfig) HasMaxRetentionDurationMsecs() bool {
	if o != nil && o.MaxRetentionDurationMsecs.IsSet() {
		return true
	}

	return false
}

// SetMaxRetentionDurationMsecs gets a reference to the given NullableInt64 and assigns it to the MaxRetentionDurationMsecs field.
func (o *FileLevelDataLockConfig) SetMaxRetentionDurationMsecs(v int64) {
	o.MaxRetentionDurationMsecs.Set(&v)
}
// SetMaxRetentionDurationMsecsNil sets the value for MaxRetentionDurationMsecs to be an explicit nil
func (o *FileLevelDataLockConfig) SetMaxRetentionDurationMsecsNil() {
	o.MaxRetentionDurationMsecs.Set(nil)
}

// UnsetMaxRetentionDurationMsecs ensures that no value is present for MaxRetentionDurationMsecs, not even an explicit nil
func (o *FileLevelDataLockConfig) UnsetMaxRetentionDurationMsecs() {
	o.MaxRetentionDurationMsecs.Unset()
}

// GetMinRetentionDurationMsecs returns the MinRetentionDurationMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileLevelDataLockConfig) GetMinRetentionDurationMsecs() int64 {
	if o == nil || IsNil(o.MinRetentionDurationMsecs.Get()) {
		var ret int64
		return ret
	}
	return *o.MinRetentionDurationMsecs.Get()
}

// GetMinRetentionDurationMsecsOk returns a tuple with the MinRetentionDurationMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileLevelDataLockConfig) GetMinRetentionDurationMsecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinRetentionDurationMsecs.Get(), o.MinRetentionDurationMsecs.IsSet()
}

// HasMinRetentionDurationMsecs returns a boolean if a field has been set.
func (o *FileLevelDataLockConfig) HasMinRetentionDurationMsecs() bool {
	if o != nil && o.MinRetentionDurationMsecs.IsSet() {
		return true
	}

	return false
}

// SetMinRetentionDurationMsecs gets a reference to the given NullableInt64 and assigns it to the MinRetentionDurationMsecs field.
func (o *FileLevelDataLockConfig) SetMinRetentionDurationMsecs(v int64) {
	o.MinRetentionDurationMsecs.Set(&v)
}
// SetMinRetentionDurationMsecsNil sets the value for MinRetentionDurationMsecs to be an explicit nil
func (o *FileLevelDataLockConfig) SetMinRetentionDurationMsecsNil() {
	o.MinRetentionDurationMsecs.Set(nil)
}

// UnsetMinRetentionDurationMsecs ensures that no value is present for MinRetentionDurationMsecs, not even an explicit nil
func (o *FileLevelDataLockConfig) UnsetMinRetentionDurationMsecs() {
	o.MinRetentionDurationMsecs.Unset()
}

// GetMode returns the Mode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileLevelDataLockConfig) GetMode() string {
	if o == nil || IsNil(o.Mode.Get()) {
		var ret string
		return ret
	}
	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileLevelDataLockConfig) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// HasMode returns a boolean if a field has been set.
func (o *FileLevelDataLockConfig) HasMode() bool {
	if o != nil && o.Mode.IsSet() {
		return true
	}

	return false
}

// SetMode gets a reference to the given NullableString and assigns it to the Mode field.
func (o *FileLevelDataLockConfig) SetMode(v string) {
	o.Mode.Set(&v)
}
// SetModeNil sets the value for Mode to be an explicit nil
func (o *FileLevelDataLockConfig) SetModeNil() {
	o.Mode.Set(nil)
}

// UnsetMode ensures that no value is present for Mode, not even an explicit nil
func (o *FileLevelDataLockConfig) UnsetMode() {
	o.Mode.Unset()
}

func (o FileLevelDataLockConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileLevelDataLockConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoLockAfterDurationIdleMsecs.IsSet() {
		toSerialize["autoLockAfterDurationIdleMsecs"] = o.AutoLockAfterDurationIdleMsecs.Get()
	}
	if o.CoexistingLockMode.IsSet() {
		toSerialize["coexistingLockMode"] = o.CoexistingLockMode.Get()
	}
	if o.DefaultRetentionDurationMsecs.IsSet() {
		toSerialize["defaultRetentionDurationMsecs"] = o.DefaultRetentionDurationMsecs.Get()
	}
	if o.DefaultRetentionDurationYears.IsSet() {
		toSerialize["defaultRetentionDurationYears"] = o.DefaultRetentionDurationYears.Get()
	}
	if o.ExpiryTimestampMsecs.IsSet() {
		toSerialize["expiryTimestampMsecs"] = o.ExpiryTimestampMsecs.Get()
	}
	if o.LockingProtocol.IsSet() {
		toSerialize["lockingProtocol"] = o.LockingProtocol.Get()
	}
	if o.MaxRetentionDurationMsecs.IsSet() {
		toSerialize["maxRetentionDurationMsecs"] = o.MaxRetentionDurationMsecs.Get()
	}
	if o.MinRetentionDurationMsecs.IsSet() {
		toSerialize["minRetentionDurationMsecs"] = o.MinRetentionDurationMsecs.Get()
	}
	if o.Mode.IsSet() {
		toSerialize["mode"] = o.Mode.Get()
	}
	return toSerialize, nil
}

type NullableFileLevelDataLockConfig struct {
	value *FileLevelDataLockConfig
	isSet bool
}

func (v NullableFileLevelDataLockConfig) Get() *FileLevelDataLockConfig {
	return v.value
}

func (v *NullableFileLevelDataLockConfig) Set(val *FileLevelDataLockConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableFileLevelDataLockConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableFileLevelDataLockConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileLevelDataLockConfig(val *FileLevelDataLockConfig) *NullableFileLevelDataLockConfig {
	return &NullableFileLevelDataLockConfig{value: val, isSet: true}
}

func (v NullableFileLevelDataLockConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileLevelDataLockConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


