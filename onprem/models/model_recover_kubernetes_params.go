/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// RecoverKubernetesParams Specifies the recovery options specific to Kubernetes environment.
type RecoverKubernetesParams struct {
	// Specifies the type of recover action to be performed.
	RecoveryAction string `json:"recoveryAction"`
	// Specifies the parameters to recover Kubernetes Namespaces.
	RecoverNamespaceParams NullableRecoverKubernetesNamespaceParams `json:"recoverNamespaceParams,omitempty"`
}

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
	if o == nil  {
		return nil, false
	}
	return &o.RecoveryAction, true
}

// SetRecoveryAction sets field value
func (o *RecoverKubernetesParams) SetRecoveryAction(v string) {
	o.RecoveryAction = v
}

// GetRecoverNamespaceParams returns the RecoverNamespaceParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverKubernetesParams) GetRecoverNamespaceParams() RecoverKubernetesNamespaceParams {
	if o == nil || o.RecoverNamespaceParams.Get() == nil {
		var ret RecoverKubernetesNamespaceParams
		return ret
	}
	return *o.RecoverNamespaceParams.Get()
}

// GetRecoverNamespaceParamsOk returns a tuple with the RecoverNamespaceParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverKubernetesParams) GetRecoverNamespaceParamsOk() (*RecoverKubernetesNamespaceParams, bool) {
	if o == nil  {
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

// SetRecoverNamespaceParams gets a reference to the given NullableRecoverKubernetesNamespaceParams and assigns it to the RecoverNamespaceParams field.
func (o *RecoverKubernetesParams) SetRecoverNamespaceParams(v RecoverKubernetesNamespaceParams) {
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

func (o RecoverKubernetesParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recoveryAction"] = o.RecoveryAction
	}
	if o.RecoverNamespaceParams.IsSet() {
		toSerialize["recoverNamespaceParams"] = o.RecoverNamespaceParams.Get()
	}
	return json.Marshal(toSerialize)
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




func (o RecoverKubernetesParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}