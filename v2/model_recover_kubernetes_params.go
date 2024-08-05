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

// checks if the RecoverKubernetesParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverKubernetesParams{}

// RecoverKubernetesParams Specifies the recovery options specific to Kubernetes environment.
type RecoverKubernetesParams struct {
	DownloadFileAndFolderParams NullableRecoverGcpParamsDownloadFileAndFolderParams `json:"downloadFileAndFolderParams,omitempty"`
	// Specifies the list of objects which need to be recovered.
	Objects []CommonRecoverObjectSnapshotParams `json:"objects,omitempty"`
	RecoverFileAndFolderParams NullableRecoverKubernetesParamsRecoverFileAndFolderParams `json:"recoverFileAndFolderParams,omitempty"`
	RecoverNamespaceParams NullableRecoverKubernetesParamsRecoverNamespaceParams `json:"recoverNamespaceParams,omitempty"`
	// Specifies the type of recover action to be performed.
	RecoveryAction string `json:"recoveryAction"`
}

type _RecoverKubernetesParams RecoverKubernetesParams

// NewRecoverKubernetesParams instantiates a new RecoverKubernetesParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverKubernetesParams(recoveryAction string) *RecoverKubernetesParams {
	this := RecoverKubernetesParams{}
	this.RecoveryAction = recoveryAction
	return &this
}

// NewRecoverKubernetesParamsWithDefaults instantiates a new RecoverKubernetesParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverKubernetesParamsWithDefaults() *RecoverKubernetesParams {
	this := RecoverKubernetesParams{}
	return &this
}

// GetDownloadFileAndFolderParams returns the DownloadFileAndFolderParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverKubernetesParams) GetDownloadFileAndFolderParams() RecoverGcpParamsDownloadFileAndFolderParams {
	if o == nil || IsNil(o.DownloadFileAndFolderParams.Get()) {
		var ret RecoverGcpParamsDownloadFileAndFolderParams
		return ret
	}
	return *o.DownloadFileAndFolderParams.Get()
}

// GetDownloadFileAndFolderParamsOk returns a tuple with the DownloadFileAndFolderParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverKubernetesParams) GetDownloadFileAndFolderParamsOk() (*RecoverGcpParamsDownloadFileAndFolderParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadFileAndFolderParams.Get(), o.DownloadFileAndFolderParams.IsSet()
}

// HasDownloadFileAndFolderParams returns a boolean if a field has been set.
func (o *RecoverKubernetesParams) HasDownloadFileAndFolderParams() bool {
	if o != nil && o.DownloadFileAndFolderParams.IsSet() {
		return true
	}

	return false
}

// SetDownloadFileAndFolderParams gets a reference to the given NullableRecoverGcpParamsDownloadFileAndFolderParams and assigns it to the DownloadFileAndFolderParams field.
func (o *RecoverKubernetesParams) SetDownloadFileAndFolderParams(v RecoverGcpParamsDownloadFileAndFolderParams) {
	o.DownloadFileAndFolderParams.Set(&v)
}
// SetDownloadFileAndFolderParamsNil sets the value for DownloadFileAndFolderParams to be an explicit nil
func (o *RecoverKubernetesParams) SetDownloadFileAndFolderParamsNil() {
	o.DownloadFileAndFolderParams.Set(nil)
}

// UnsetDownloadFileAndFolderParams ensures that no value is present for DownloadFileAndFolderParams, not even an explicit nil
func (o *RecoverKubernetesParams) UnsetDownloadFileAndFolderParams() {
	o.DownloadFileAndFolderParams.Unset()
}

// GetObjects returns the Objects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverKubernetesParams) GetObjects() []CommonRecoverObjectSnapshotParams {
	if o == nil {
		var ret []CommonRecoverObjectSnapshotParams
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverKubernetesParams) GetObjectsOk() ([]CommonRecoverObjectSnapshotParams, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *RecoverKubernetesParams) HasObjects() bool {
	if o != nil && !IsNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []CommonRecoverObjectSnapshotParams and assigns it to the Objects field.
func (o *RecoverKubernetesParams) SetObjects(v []CommonRecoverObjectSnapshotParams) {
	o.Objects = v
}

// GetRecoverFileAndFolderParams returns the RecoverFileAndFolderParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverKubernetesParams) GetRecoverFileAndFolderParams() RecoverKubernetesParamsRecoverFileAndFolderParams {
	if o == nil || IsNil(o.RecoverFileAndFolderParams.Get()) {
		var ret RecoverKubernetesParamsRecoverFileAndFolderParams
		return ret
	}
	return *o.RecoverFileAndFolderParams.Get()
}

// GetRecoverFileAndFolderParamsOk returns a tuple with the RecoverFileAndFolderParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverKubernetesParams) GetRecoverFileAndFolderParamsOk() (*RecoverKubernetesParamsRecoverFileAndFolderParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecoverFileAndFolderParams.Get(), o.RecoverFileAndFolderParams.IsSet()
}

// HasRecoverFileAndFolderParams returns a boolean if a field has been set.
func (o *RecoverKubernetesParams) HasRecoverFileAndFolderParams() bool {
	if o != nil && o.RecoverFileAndFolderParams.IsSet() {
		return true
	}

	return false
}

// SetRecoverFileAndFolderParams gets a reference to the given NullableRecoverKubernetesParamsRecoverFileAndFolderParams and assigns it to the RecoverFileAndFolderParams field.
func (o *RecoverKubernetesParams) SetRecoverFileAndFolderParams(v RecoverKubernetesParamsRecoverFileAndFolderParams) {
	o.RecoverFileAndFolderParams.Set(&v)
}
// SetRecoverFileAndFolderParamsNil sets the value for RecoverFileAndFolderParams to be an explicit nil
func (o *RecoverKubernetesParams) SetRecoverFileAndFolderParamsNil() {
	o.RecoverFileAndFolderParams.Set(nil)
}

// UnsetRecoverFileAndFolderParams ensures that no value is present for RecoverFileAndFolderParams, not even an explicit nil
func (o *RecoverKubernetesParams) UnsetRecoverFileAndFolderParams() {
	o.RecoverFileAndFolderParams.Unset()
}

// GetRecoverNamespaceParams returns the RecoverNamespaceParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverKubernetesParams) GetRecoverNamespaceParams() RecoverKubernetesParamsRecoverNamespaceParams {
	if o == nil || IsNil(o.RecoverNamespaceParams.Get()) {
		var ret RecoverKubernetesParamsRecoverNamespaceParams
		return ret
	}
	return *o.RecoverNamespaceParams.Get()
}

// GetRecoverNamespaceParamsOk returns a tuple with the RecoverNamespaceParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverKubernetesParams) GetRecoverNamespaceParamsOk() (*RecoverKubernetesParamsRecoverNamespaceParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecoverNamespaceParams.Get(), o.RecoverNamespaceParams.IsSet()
}

// HasRecoverNamespaceParams returns a boolean if a field has been set.
func (o *RecoverKubernetesParams) HasRecoverNamespaceParams() bool {
	if o != nil && o.RecoverNamespaceParams.IsSet() {
		return true
	}

	return false
}

// SetRecoverNamespaceParams gets a reference to the given NullableRecoverKubernetesParamsRecoverNamespaceParams and assigns it to the RecoverNamespaceParams field.
func (o *RecoverKubernetesParams) SetRecoverNamespaceParams(v RecoverKubernetesParamsRecoverNamespaceParams) {
	o.RecoverNamespaceParams.Set(&v)
}
// SetRecoverNamespaceParamsNil sets the value for RecoverNamespaceParams to be an explicit nil
func (o *RecoverKubernetesParams) SetRecoverNamespaceParamsNil() {
	o.RecoverNamespaceParams.Set(nil)
}

// UnsetRecoverNamespaceParams ensures that no value is present for RecoverNamespaceParams, not even an explicit nil
func (o *RecoverKubernetesParams) UnsetRecoverNamespaceParams() {
	o.RecoverNamespaceParams.Unset()
}

// GetRecoveryAction returns the RecoveryAction field value
func (o *RecoverKubernetesParams) GetRecoveryAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecoveryAction
}

// GetRecoveryActionOk returns a tuple with the RecoveryAction field value
// and a boolean to check if the value has been set.
func (o *RecoverKubernetesParams) GetRecoveryActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecoveryAction, true
}

// SetRecoveryAction sets field value
func (o *RecoverKubernetesParams) SetRecoveryAction(v string) {
	o.RecoveryAction = v
}

func (o RecoverKubernetesParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverKubernetesParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DownloadFileAndFolderParams.IsSet() {
		toSerialize["downloadFileAndFolderParams"] = o.DownloadFileAndFolderParams.Get()
	}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if o.RecoverFileAndFolderParams.IsSet() {
		toSerialize["recoverFileAndFolderParams"] = o.RecoverFileAndFolderParams.Get()
	}
	if o.RecoverNamespaceParams.IsSet() {
		toSerialize["recoverNamespaceParams"] = o.RecoverNamespaceParams.Get()
	}
	toSerialize["recoveryAction"] = o.RecoveryAction
	return toSerialize, nil
}

func (o *RecoverKubernetesParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"recoveryAction",
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

	varRecoverKubernetesParams := _RecoverKubernetesParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverKubernetesParams)

	if err != nil {
		return err
	}

	*o = RecoverKubernetesParams(varRecoverKubernetesParams)

	return err
}

type NullableRecoverKubernetesParams struct {
	value *RecoverKubernetesParams
	isSet bool
}

func (v NullableRecoverKubernetesParams) Get() *RecoverKubernetesParams {
	return v.value
}

func (v *NullableRecoverKubernetesParams) Set(val *RecoverKubernetesParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverKubernetesParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverKubernetesParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverKubernetesParams(val *RecoverKubernetesParams) *NullableRecoverKubernetesParams {
	return &NullableRecoverKubernetesParams{value: val, isSet: true}
}

func (v NullableRecoverKubernetesParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverKubernetesParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


