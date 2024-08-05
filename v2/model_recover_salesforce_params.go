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

// checks if the RecoverSalesforceParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverSalesforceParams{}

// RecoverSalesforceParams Specifies the recovery options specific to Salesforce environment.
type RecoverSalesforceParams struct {
	// Specifies whether to continue recovering other salesforce objects if one of Object failed to recover. Default value is false.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
	// Specifies the list of recover Object parameters.
	Objects []CommonRecoverObjectSnapshotParams `json:"objects"`
	RecoverSfdcObjectParams *RecoverSfdcObjectParams `json:"recoverSfdcObjectParams,omitempty"`
	// Specifies the id of registered source where the objects are to be recovered. If this is not specified, the recovery job will recover to the original location.
	RecoverTo NullableInt64 `json:"recoverTo,omitempty"`
	// Specifies the type of recover action to be performed.
	RecoveryAction string `json:"recoveryAction"`
}

type _RecoverSalesforceParams RecoverSalesforceParams

// NewRecoverSalesforceParams instantiates a new RecoverSalesforceParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverSalesforceParams(objects []CommonRecoverObjectSnapshotParams, recoveryAction string) *RecoverSalesforceParams {
	this := RecoverSalesforceParams{}
	this.Objects = objects
	this.RecoveryAction = recoveryAction
	return &this
}

// NewRecoverSalesforceParamsWithDefaults instantiates a new RecoverSalesforceParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverSalesforceParamsWithDefaults() *RecoverSalesforceParams {
	this := RecoverSalesforceParams{}
	return &this
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverSalesforceParams) GetContinueOnError() bool {
	if o == nil || IsNil(o.ContinueOnError.Get()) {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverSalesforceParams) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *RecoverSalesforceParams) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *RecoverSalesforceParams) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *RecoverSalesforceParams) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *RecoverSalesforceParams) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

// GetObjects returns the Objects field value
// If the value is explicit nil, the zero value for []CommonRecoverObjectSnapshotParams will be returned
func (o *RecoverSalesforceParams) GetObjects() []CommonRecoverObjectSnapshotParams {
	if o == nil {
		var ret []CommonRecoverObjectSnapshotParams
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverSalesforceParams) GetObjectsOk() ([]CommonRecoverObjectSnapshotParams, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// SetObjects sets field value
func (o *RecoverSalesforceParams) SetObjects(v []CommonRecoverObjectSnapshotParams) {
	o.Objects = v
}

// GetRecoverSfdcObjectParams returns the RecoverSfdcObjectParams field value if set, zero value otherwise.
func (o *RecoverSalesforceParams) GetRecoverSfdcObjectParams() RecoverSfdcObjectParams {
	if o == nil || IsNil(o.RecoverSfdcObjectParams) {
		var ret RecoverSfdcObjectParams
		return ret
	}
	return *o.RecoverSfdcObjectParams
}

// GetRecoverSfdcObjectParamsOk returns a tuple with the RecoverSfdcObjectParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverSalesforceParams) GetRecoverSfdcObjectParamsOk() (*RecoverSfdcObjectParams, bool) {
	if o == nil || IsNil(o.RecoverSfdcObjectParams) {
		return nil, false
	}
	return o.RecoverSfdcObjectParams, true
}

// HasRecoverSfdcObjectParams returns a boolean if a field has been set.
func (o *RecoverSalesforceParams) HasRecoverSfdcObjectParams() bool {
	if o != nil && !IsNil(o.RecoverSfdcObjectParams) {
		return true
	}

	return false
}

// SetRecoverSfdcObjectParams gets a reference to the given RecoverSfdcObjectParams and assigns it to the RecoverSfdcObjectParams field.
func (o *RecoverSalesforceParams) SetRecoverSfdcObjectParams(v RecoverSfdcObjectParams) {
	o.RecoverSfdcObjectParams = &v
}

// GetRecoverTo returns the RecoverTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverSalesforceParams) GetRecoverTo() int64 {
	if o == nil || IsNil(o.RecoverTo.Get()) {
		var ret int64
		return ret
	}
	return *o.RecoverTo.Get()
}

// GetRecoverToOk returns a tuple with the RecoverTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverSalesforceParams) GetRecoverToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecoverTo.Get(), o.RecoverTo.IsSet()
}

// HasRecoverTo returns a boolean if a field has been set.
func (o *RecoverSalesforceParams) HasRecoverTo() bool {
	if o != nil && o.RecoverTo.IsSet() {
		return true
	}

	return false
}

// SetRecoverTo gets a reference to the given NullableInt64 and assigns it to the RecoverTo field.
func (o *RecoverSalesforceParams) SetRecoverTo(v int64) {
	o.RecoverTo.Set(&v)
}
// SetRecoverToNil sets the value for RecoverTo to be an explicit nil
func (o *RecoverSalesforceParams) SetRecoverToNil() {
	o.RecoverTo.Set(nil)
}

// UnsetRecoverTo ensures that no value is present for RecoverTo, not even an explicit nil
func (o *RecoverSalesforceParams) UnsetRecoverTo() {
	o.RecoverTo.Unset()
}

// GetRecoveryAction returns the RecoveryAction field value
func (o *RecoverSalesforceParams) GetRecoveryAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecoveryAction
}

// GetRecoveryActionOk returns a tuple with the RecoveryAction field value
// and a boolean to check if the value has been set.
func (o *RecoverSalesforceParams) GetRecoveryActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecoveryAction, true
}

// SetRecoveryAction sets field value
func (o *RecoverSalesforceParams) SetRecoveryAction(v string) {
	o.RecoveryAction = v
}

func (o RecoverSalesforceParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverSalesforceParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if !IsNil(o.RecoverSfdcObjectParams) {
		toSerialize["recoverSfdcObjectParams"] = o.RecoverSfdcObjectParams
	}
	if o.RecoverTo.IsSet() {
		toSerialize["recoverTo"] = o.RecoverTo.Get()
	}
	toSerialize["recoveryAction"] = o.RecoveryAction
	return toSerialize, nil
}

func (o *RecoverSalesforceParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objects",
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

	varRecoverSalesforceParams := _RecoverSalesforceParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverSalesforceParams)

	if err != nil {
		return err
	}

	*o = RecoverSalesforceParams(varRecoverSalesforceParams)

	return err
}

type NullableRecoverSalesforceParams struct {
	value *RecoverSalesforceParams
	isSet bool
}

func (v NullableRecoverSalesforceParams) Get() *RecoverSalesforceParams {
	return v.value
}

func (v *NullableRecoverSalesforceParams) Set(val *RecoverSalesforceParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverSalesforceParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverSalesforceParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverSalesforceParams(val *RecoverSalesforceParams) *NullableRecoverSalesforceParams {
	return &NullableRecoverSalesforceParams{value: val, isSet: true}
}

func (v NullableRecoverSalesforceParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverSalesforceParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


