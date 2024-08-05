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

// checks if the UdaRegistrationParamsPrimaryFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UdaRegistrationParamsPrimaryFields{}

// UdaRegistrationParamsPrimaryFields Parameters to customize existing/ default form fields.
type UdaRegistrationParamsPrimaryFields struct {
	AppAuthentication NullableUdaRegistrationParamsPrimaryFieldsAppAuthentication `json:"appAuthentication,omitempty"`
	MountView NullableUdaRegistrationParamsPrimaryFieldsMountView `json:"mountView,omitempty"`
	ScriptDir NullableUdaRegistrationParamsPrimaryFieldsScriptDir `json:"scriptDir,omitempty"`
	SourceRegistrationArgs NullableUdaRegistrationParamsPrimaryFieldsSourceRegistrationArgs `json:"sourceRegistrationArgs,omitempty"`
}

// NewUdaRegistrationParamsPrimaryFields instantiates a new UdaRegistrationParamsPrimaryFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUdaRegistrationParamsPrimaryFields() *UdaRegistrationParamsPrimaryFields {
	this := UdaRegistrationParamsPrimaryFields{}
	return &this
}

// NewUdaRegistrationParamsPrimaryFieldsWithDefaults instantiates a new UdaRegistrationParamsPrimaryFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUdaRegistrationParamsPrimaryFieldsWithDefaults() *UdaRegistrationParamsPrimaryFields {
	this := UdaRegistrationParamsPrimaryFields{}
	return &this
}

// GetAppAuthentication returns the AppAuthentication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UdaRegistrationParamsPrimaryFields) GetAppAuthentication() UdaRegistrationParamsPrimaryFieldsAppAuthentication {
	if o == nil || IsNil(o.AppAuthentication.Get()) {
		var ret UdaRegistrationParamsPrimaryFieldsAppAuthentication
		return ret
	}
	return *o.AppAuthentication.Get()
}

// GetAppAuthenticationOk returns a tuple with the AppAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UdaRegistrationParamsPrimaryFields) GetAppAuthenticationOk() (*UdaRegistrationParamsPrimaryFieldsAppAuthentication, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppAuthentication.Get(), o.AppAuthentication.IsSet()
}

// HasAppAuthentication returns a boolean if a field has been set.
func (o *UdaRegistrationParamsPrimaryFields) HasAppAuthentication() bool {
	if o != nil && o.AppAuthentication.IsSet() {
		return true
	}

	return false
}

// SetAppAuthentication gets a reference to the given NullableUdaRegistrationParamsPrimaryFieldsAppAuthentication and assigns it to the AppAuthentication field.
func (o *UdaRegistrationParamsPrimaryFields) SetAppAuthentication(v UdaRegistrationParamsPrimaryFieldsAppAuthentication) {
	o.AppAuthentication.Set(&v)
}
// SetAppAuthenticationNil sets the value for AppAuthentication to be an explicit nil
func (o *UdaRegistrationParamsPrimaryFields) SetAppAuthenticationNil() {
	o.AppAuthentication.Set(nil)
}

// UnsetAppAuthentication ensures that no value is present for AppAuthentication, not even an explicit nil
func (o *UdaRegistrationParamsPrimaryFields) UnsetAppAuthentication() {
	o.AppAuthentication.Unset()
}

// GetMountView returns the MountView field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UdaRegistrationParamsPrimaryFields) GetMountView() UdaRegistrationParamsPrimaryFieldsMountView {
	if o == nil || IsNil(o.MountView.Get()) {
		var ret UdaRegistrationParamsPrimaryFieldsMountView
		return ret
	}
	return *o.MountView.Get()
}

// GetMountViewOk returns a tuple with the MountView field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UdaRegistrationParamsPrimaryFields) GetMountViewOk() (*UdaRegistrationParamsPrimaryFieldsMountView, bool) {
	if o == nil {
		return nil, false
	}
	return o.MountView.Get(), o.MountView.IsSet()
}

// HasMountView returns a boolean if a field has been set.
func (o *UdaRegistrationParamsPrimaryFields) HasMountView() bool {
	if o != nil && o.MountView.IsSet() {
		return true
	}

	return false
}

// SetMountView gets a reference to the given NullableUdaRegistrationParamsPrimaryFieldsMountView and assigns it to the MountView field.
func (o *UdaRegistrationParamsPrimaryFields) SetMountView(v UdaRegistrationParamsPrimaryFieldsMountView) {
	o.MountView.Set(&v)
}
// SetMountViewNil sets the value for MountView to be an explicit nil
func (o *UdaRegistrationParamsPrimaryFields) SetMountViewNil() {
	o.MountView.Set(nil)
}

// UnsetMountView ensures that no value is present for MountView, not even an explicit nil
func (o *UdaRegistrationParamsPrimaryFields) UnsetMountView() {
	o.MountView.Unset()
}

// GetScriptDir returns the ScriptDir field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UdaRegistrationParamsPrimaryFields) GetScriptDir() UdaRegistrationParamsPrimaryFieldsScriptDir {
	if o == nil || IsNil(o.ScriptDir.Get()) {
		var ret UdaRegistrationParamsPrimaryFieldsScriptDir
		return ret
	}
	return *o.ScriptDir.Get()
}

// GetScriptDirOk returns a tuple with the ScriptDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UdaRegistrationParamsPrimaryFields) GetScriptDirOk() (*UdaRegistrationParamsPrimaryFieldsScriptDir, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScriptDir.Get(), o.ScriptDir.IsSet()
}

// HasScriptDir returns a boolean if a field has been set.
func (o *UdaRegistrationParamsPrimaryFields) HasScriptDir() bool {
	if o != nil && o.ScriptDir.IsSet() {
		return true
	}

	return false
}

// SetScriptDir gets a reference to the given NullableUdaRegistrationParamsPrimaryFieldsScriptDir and assigns it to the ScriptDir field.
func (o *UdaRegistrationParamsPrimaryFields) SetScriptDir(v UdaRegistrationParamsPrimaryFieldsScriptDir) {
	o.ScriptDir.Set(&v)
}
// SetScriptDirNil sets the value for ScriptDir to be an explicit nil
func (o *UdaRegistrationParamsPrimaryFields) SetScriptDirNil() {
	o.ScriptDir.Set(nil)
}

// UnsetScriptDir ensures that no value is present for ScriptDir, not even an explicit nil
func (o *UdaRegistrationParamsPrimaryFields) UnsetScriptDir() {
	o.ScriptDir.Unset()
}

// GetSourceRegistrationArgs returns the SourceRegistrationArgs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UdaRegistrationParamsPrimaryFields) GetSourceRegistrationArgs() UdaRegistrationParamsPrimaryFieldsSourceRegistrationArgs {
	if o == nil || IsNil(o.SourceRegistrationArgs.Get()) {
		var ret UdaRegistrationParamsPrimaryFieldsSourceRegistrationArgs
		return ret
	}
	return *o.SourceRegistrationArgs.Get()
}

// GetSourceRegistrationArgsOk returns a tuple with the SourceRegistrationArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UdaRegistrationParamsPrimaryFields) GetSourceRegistrationArgsOk() (*UdaRegistrationParamsPrimaryFieldsSourceRegistrationArgs, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceRegistrationArgs.Get(), o.SourceRegistrationArgs.IsSet()
}

// HasSourceRegistrationArgs returns a boolean if a field has been set.
func (o *UdaRegistrationParamsPrimaryFields) HasSourceRegistrationArgs() bool {
	if o != nil && o.SourceRegistrationArgs.IsSet() {
		return true
	}

	return false
}

// SetSourceRegistrationArgs gets a reference to the given NullableUdaRegistrationParamsPrimaryFieldsSourceRegistrationArgs and assigns it to the SourceRegistrationArgs field.
func (o *UdaRegistrationParamsPrimaryFields) SetSourceRegistrationArgs(v UdaRegistrationParamsPrimaryFieldsSourceRegistrationArgs) {
	o.SourceRegistrationArgs.Set(&v)
}
// SetSourceRegistrationArgsNil sets the value for SourceRegistrationArgs to be an explicit nil
func (o *UdaRegistrationParamsPrimaryFields) SetSourceRegistrationArgsNil() {
	o.SourceRegistrationArgs.Set(nil)
}

// UnsetSourceRegistrationArgs ensures that no value is present for SourceRegistrationArgs, not even an explicit nil
func (o *UdaRegistrationParamsPrimaryFields) UnsetSourceRegistrationArgs() {
	o.SourceRegistrationArgs.Unset()
}

func (o UdaRegistrationParamsPrimaryFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UdaRegistrationParamsPrimaryFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AppAuthentication.IsSet() {
		toSerialize["appAuthentication"] = o.AppAuthentication.Get()
	}
	if o.MountView.IsSet() {
		toSerialize["mountView"] = o.MountView.Get()
	}
	if o.ScriptDir.IsSet() {
		toSerialize["scriptDir"] = o.ScriptDir.Get()
	}
	if o.SourceRegistrationArgs.IsSet() {
		toSerialize["sourceRegistrationArgs"] = o.SourceRegistrationArgs.Get()
	}
	return toSerialize, nil
}

type NullableUdaRegistrationParamsPrimaryFields struct {
	value *UdaRegistrationParamsPrimaryFields
	isSet bool
}

func (v NullableUdaRegistrationParamsPrimaryFields) Get() *UdaRegistrationParamsPrimaryFields {
	return v.value
}

func (v *NullableUdaRegistrationParamsPrimaryFields) Set(val *UdaRegistrationParamsPrimaryFields) {
	v.value = val
	v.isSet = true
}

func (v NullableUdaRegistrationParamsPrimaryFields) IsSet() bool {
	return v.isSet
}

func (v *NullableUdaRegistrationParamsPrimaryFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUdaRegistrationParamsPrimaryFields(val *UdaRegistrationParamsPrimaryFields) *NullableUdaRegistrationParamsPrimaryFields {
	return &NullableUdaRegistrationParamsPrimaryFields{value: val, isSet: true}
}

func (v NullableUdaRegistrationParamsPrimaryFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUdaRegistrationParamsPrimaryFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


