/*
 * Cohesity REST API
 *
 * This API list provides operations for interfacing with the Cohesity Cluster.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesitysdk

import (
	"encoding/json"
)

// UpdateRestoreTaskParams UpdateRestoreTaskParams holds the information to update a Restore Task in Magneto.
type UpdateRestoreTaskParams struct {
	AdOptions *AdRestoreOptions `json:"adOptions,omitempty"`
	// Enables Auto Sync feature for SQL Multi-stage Restore task.
	EnableAutoSync NullableBool `json:"enableAutoSync,omitempty"`
	// Specifies the ID of the existing Restore Task to update.
	RestoreTaskId NullableInt64 `json:"restoreTaskId,omitempty"`
	// Specifies the sql options to update the Restore Task with. Specifies the action type of multi stage SQL restore.  'kCreate' specifies the create action for a restore. 'kUpdate' specifies the user action to update an ongoing restore. 'kFinalize' specifies the user action to finalize a restore.
	SqlOptions NullableString `json:"sqlOptions,omitempty"`
}

// NewUpdateRestoreTaskParams instantiates a new UpdateRestoreTaskParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRestoreTaskParams() *UpdateRestoreTaskParams {
	this := UpdateRestoreTaskParams{}
	return &this
}

// NewUpdateRestoreTaskParamsWithDefaults instantiates a new UpdateRestoreTaskParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRestoreTaskParamsWithDefaults() *UpdateRestoreTaskParams {
	this := UpdateRestoreTaskParams{}
	return &this
}

// GetAdOptions returns the AdOptions field value if set, zero value otherwise.
func (o *UpdateRestoreTaskParams) GetAdOptions() AdRestoreOptions {
	if o == nil || o.AdOptions == nil {
		var ret AdRestoreOptions
		return ret
	}
	return *o.AdOptions
}

// GetAdOptionsOk returns a tuple with the AdOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRestoreTaskParams) GetAdOptionsOk() (*AdRestoreOptions, bool) {
	if o == nil || o.AdOptions == nil {
		return nil, false
	}
	return o.AdOptions, true
}

// HasAdOptions returns a boolean if a field has been set.
func (o *UpdateRestoreTaskParams) HasAdOptions() bool {
	if o != nil && o.AdOptions != nil {
		return true
	}

	return false
}

// SetAdOptions gets a reference to the given AdRestoreOptions and assigns it to the AdOptions field.
func (o *UpdateRestoreTaskParams) SetAdOptions(v AdRestoreOptions) {
	o.AdOptions = &v
}

// GetEnableAutoSync returns the EnableAutoSync field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateRestoreTaskParams) GetEnableAutoSync() bool {
	if o == nil || o.EnableAutoSync.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EnableAutoSync.Get()
}

// GetEnableAutoSyncOk returns a tuple with the EnableAutoSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateRestoreTaskParams) GetEnableAutoSyncOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EnableAutoSync.Get(), o.EnableAutoSync.IsSet()
}

// HasEnableAutoSync returns a boolean if a field has been set.
func (o *UpdateRestoreTaskParams) HasEnableAutoSync() bool {
	if o != nil && o.EnableAutoSync.IsSet() {
		return true
	}

	return false
}

// SetEnableAutoSync gets a reference to the given NullableBool and assigns it to the EnableAutoSync field.
func (o *UpdateRestoreTaskParams) SetEnableAutoSync(v bool) {
	o.EnableAutoSync.Set(&v)
}
// SetEnableAutoSyncNil sets the value for EnableAutoSync to be an explicit nil
func (o *UpdateRestoreTaskParams) SetEnableAutoSyncNil() {
	o.EnableAutoSync.Set(nil)
}

// UnsetEnableAutoSync ensures that no value is present for EnableAutoSync, not even an explicit nil
func (o *UpdateRestoreTaskParams) UnsetEnableAutoSync() {
	o.EnableAutoSync.Unset()
}

// GetRestoreTaskId returns the RestoreTaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateRestoreTaskParams) GetRestoreTaskId() int64 {
	if o == nil || o.RestoreTaskId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RestoreTaskId.Get()
}

// GetRestoreTaskIdOk returns a tuple with the RestoreTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateRestoreTaskParams) GetRestoreTaskIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RestoreTaskId.Get(), o.RestoreTaskId.IsSet()
}

// HasRestoreTaskId returns a boolean if a field has been set.
func (o *UpdateRestoreTaskParams) HasRestoreTaskId() bool {
	if o != nil && o.RestoreTaskId.IsSet() {
		return true
	}

	return false
}

// SetRestoreTaskId gets a reference to the given NullableInt64 and assigns it to the RestoreTaskId field.
func (o *UpdateRestoreTaskParams) SetRestoreTaskId(v int64) {
	o.RestoreTaskId.Set(&v)
}
// SetRestoreTaskIdNil sets the value for RestoreTaskId to be an explicit nil
func (o *UpdateRestoreTaskParams) SetRestoreTaskIdNil() {
	o.RestoreTaskId.Set(nil)
}

// UnsetRestoreTaskId ensures that no value is present for RestoreTaskId, not even an explicit nil
func (o *UpdateRestoreTaskParams) UnsetRestoreTaskId() {
	o.RestoreTaskId.Unset()
}

// GetSqlOptions returns the SqlOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateRestoreTaskParams) GetSqlOptions() string {
	if o == nil || o.SqlOptions.Get() == nil {
		var ret string
		return ret
	}
	return *o.SqlOptions.Get()
}

// GetSqlOptionsOk returns a tuple with the SqlOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateRestoreTaskParams) GetSqlOptionsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SqlOptions.Get(), o.SqlOptions.IsSet()
}

// HasSqlOptions returns a boolean if a field has been set.
func (o *UpdateRestoreTaskParams) HasSqlOptions() bool {
	if o != nil && o.SqlOptions.IsSet() {
		return true
	}

	return false
}

// SetSqlOptions gets a reference to the given NullableString and assigns it to the SqlOptions field.
func (o *UpdateRestoreTaskParams) SetSqlOptions(v string) {
	o.SqlOptions.Set(&v)
}
// SetSqlOptionsNil sets the value for SqlOptions to be an explicit nil
func (o *UpdateRestoreTaskParams) SetSqlOptionsNil() {
	o.SqlOptions.Set(nil)
}

// UnsetSqlOptions ensures that no value is present for SqlOptions, not even an explicit nil
func (o *UpdateRestoreTaskParams) UnsetSqlOptions() {
	o.SqlOptions.Unset()
}

func (o UpdateRestoreTaskParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdOptions != nil {
		toSerialize["adOptions"] = o.AdOptions
	}
	if o.EnableAutoSync.IsSet() {
		toSerialize["enableAutoSync"] = o.EnableAutoSync.Get()
	}
	if o.RestoreTaskId.IsSet() {
		toSerialize["restoreTaskId"] = o.RestoreTaskId.Get()
	}
	if o.SqlOptions.IsSet() {
		toSerialize["sqlOptions"] = o.SqlOptions.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateRestoreTaskParams struct {
	value *UpdateRestoreTaskParams
	isSet bool
}

func (v NullableUpdateRestoreTaskParams) Get() *UpdateRestoreTaskParams {
	return v.value
}

func (v *NullableUpdateRestoreTaskParams) Set(val *UpdateRestoreTaskParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRestoreTaskParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRestoreTaskParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRestoreTaskParams(val *UpdateRestoreTaskParams) *NullableUpdateRestoreTaskParams {
	return &NullableUpdateRestoreTaskParams{value: val, isSet: true}
}

func (v NullableUpdateRestoreTaskParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRestoreTaskParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


