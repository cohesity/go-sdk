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

// checks if the RecoverMailboxParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverMailboxParams{}

// RecoverMailboxParams Specifies the parameters to recover an Office 365 Mailbox.
type RecoverMailboxParams struct {
	// Specifies whether to continue recovering other Mailboxes if one of Mailbox failed to recover. Default value is false.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
	// Specifies a list of Mailbox params associated with the objects to recover.
	Objects []ObjectMailboxParam `json:"objects"`
	PstParams *RecoverMailboxParamsPstParams `json:"pstParams,omitempty"`
	// Specifies whether to skip the recovery of the archive mailbox and/or items present in the archive mailbox. Default value is true
	SkipRecoverArchiveMailbox NullableBool `json:"skipRecoverArchiveMailbox,omitempty"`
	// Specifies whether to skip the recovery of the Archive Recoverable Items present in the selected snapshot. Default value is true
	SkipRecoverArchiveRecoverableItems NullableBool `json:"skipRecoverArchiveRecoverableItems,omitempty"`
	// Specifies whether to skip the recovery of the Recoverable Items present in the selected snapshot. Default value is true
	SkipRecoverRecoverableItems NullableBool `json:"skipRecoverRecoverableItems,omitempty"`
	TargetMailbox *RecoverMailboxParamsTargetMailbox `json:"targetMailbox,omitempty"`
}

type _RecoverMailboxParams RecoverMailboxParams

// NewRecoverMailboxParams instantiates a new RecoverMailboxParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverMailboxParams(objects []ObjectMailboxParam) *RecoverMailboxParams {
	this := RecoverMailboxParams{}
	this.Objects = objects
	return &this
}

// NewRecoverMailboxParamsWithDefaults instantiates a new RecoverMailboxParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverMailboxParamsWithDefaults() *RecoverMailboxParams {
	this := RecoverMailboxParams{}
	return &this
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverMailboxParams) GetContinueOnError() bool {
	if o == nil || IsNil(o.ContinueOnError.Get()) {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverMailboxParams) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *RecoverMailboxParams) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *RecoverMailboxParams) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *RecoverMailboxParams) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *RecoverMailboxParams) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

// GetObjects returns the Objects field value
// If the value is explicit nil, the zero value for []ObjectMailboxParam will be returned
func (o *RecoverMailboxParams) GetObjects() []ObjectMailboxParam {
	if o == nil {
		var ret []ObjectMailboxParam
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverMailboxParams) GetObjectsOk() ([]ObjectMailboxParam, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// SetObjects sets field value
func (o *RecoverMailboxParams) SetObjects(v []ObjectMailboxParam) {
	o.Objects = v
}

// GetPstParams returns the PstParams field value if set, zero value otherwise.
func (o *RecoverMailboxParams) GetPstParams() RecoverMailboxParamsPstParams {
	if o == nil || IsNil(o.PstParams) {
		var ret RecoverMailboxParamsPstParams
		return ret
	}
	return *o.PstParams
}

// GetPstParamsOk returns a tuple with the PstParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverMailboxParams) GetPstParamsOk() (*RecoverMailboxParamsPstParams, bool) {
	if o == nil || IsNil(o.PstParams) {
		return nil, false
	}
	return o.PstParams, true
}

// HasPstParams returns a boolean if a field has been set.
func (o *RecoverMailboxParams) HasPstParams() bool {
	if o != nil && !IsNil(o.PstParams) {
		return true
	}

	return false
}

// SetPstParams gets a reference to the given RecoverMailboxParamsPstParams and assigns it to the PstParams field.
func (o *RecoverMailboxParams) SetPstParams(v RecoverMailboxParamsPstParams) {
	o.PstParams = &v
}

// GetSkipRecoverArchiveMailbox returns the SkipRecoverArchiveMailbox field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverMailboxParams) GetSkipRecoverArchiveMailbox() bool {
	if o == nil || IsNil(o.SkipRecoverArchiveMailbox.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipRecoverArchiveMailbox.Get()
}

// GetSkipRecoverArchiveMailboxOk returns a tuple with the SkipRecoverArchiveMailbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverMailboxParams) GetSkipRecoverArchiveMailboxOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipRecoverArchiveMailbox.Get(), o.SkipRecoverArchiveMailbox.IsSet()
}

// HasSkipRecoverArchiveMailbox returns a boolean if a field has been set.
func (o *RecoverMailboxParams) HasSkipRecoverArchiveMailbox() bool {
	if o != nil && o.SkipRecoverArchiveMailbox.IsSet() {
		return true
	}

	return false
}

// SetSkipRecoverArchiveMailbox gets a reference to the given NullableBool and assigns it to the SkipRecoverArchiveMailbox field.
func (o *RecoverMailboxParams) SetSkipRecoverArchiveMailbox(v bool) {
	o.SkipRecoverArchiveMailbox.Set(&v)
}
// SetSkipRecoverArchiveMailboxNil sets the value for SkipRecoverArchiveMailbox to be an explicit nil
func (o *RecoverMailboxParams) SetSkipRecoverArchiveMailboxNil() {
	o.SkipRecoverArchiveMailbox.Set(nil)
}

// UnsetSkipRecoverArchiveMailbox ensures that no value is present for SkipRecoverArchiveMailbox, not even an explicit nil
func (o *RecoverMailboxParams) UnsetSkipRecoverArchiveMailbox() {
	o.SkipRecoverArchiveMailbox.Unset()
}

// GetSkipRecoverArchiveRecoverableItems returns the SkipRecoverArchiveRecoverableItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverMailboxParams) GetSkipRecoverArchiveRecoverableItems() bool {
	if o == nil || IsNil(o.SkipRecoverArchiveRecoverableItems.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipRecoverArchiveRecoverableItems.Get()
}

// GetSkipRecoverArchiveRecoverableItemsOk returns a tuple with the SkipRecoverArchiveRecoverableItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverMailboxParams) GetSkipRecoverArchiveRecoverableItemsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipRecoverArchiveRecoverableItems.Get(), o.SkipRecoverArchiveRecoverableItems.IsSet()
}

// HasSkipRecoverArchiveRecoverableItems returns a boolean if a field has been set.
func (o *RecoverMailboxParams) HasSkipRecoverArchiveRecoverableItems() bool {
	if o != nil && o.SkipRecoverArchiveRecoverableItems.IsSet() {
		return true
	}

	return false
}

// SetSkipRecoverArchiveRecoverableItems gets a reference to the given NullableBool and assigns it to the SkipRecoverArchiveRecoverableItems field.
func (o *RecoverMailboxParams) SetSkipRecoverArchiveRecoverableItems(v bool) {
	o.SkipRecoverArchiveRecoverableItems.Set(&v)
}
// SetSkipRecoverArchiveRecoverableItemsNil sets the value for SkipRecoverArchiveRecoverableItems to be an explicit nil
func (o *RecoverMailboxParams) SetSkipRecoverArchiveRecoverableItemsNil() {
	o.SkipRecoverArchiveRecoverableItems.Set(nil)
}

// UnsetSkipRecoverArchiveRecoverableItems ensures that no value is present for SkipRecoverArchiveRecoverableItems, not even an explicit nil
func (o *RecoverMailboxParams) UnsetSkipRecoverArchiveRecoverableItems() {
	o.SkipRecoverArchiveRecoverableItems.Unset()
}

// GetSkipRecoverRecoverableItems returns the SkipRecoverRecoverableItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverMailboxParams) GetSkipRecoverRecoverableItems() bool {
	if o == nil || IsNil(o.SkipRecoverRecoverableItems.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipRecoverRecoverableItems.Get()
}

// GetSkipRecoverRecoverableItemsOk returns a tuple with the SkipRecoverRecoverableItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverMailboxParams) GetSkipRecoverRecoverableItemsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipRecoverRecoverableItems.Get(), o.SkipRecoverRecoverableItems.IsSet()
}

// HasSkipRecoverRecoverableItems returns a boolean if a field has been set.
func (o *RecoverMailboxParams) HasSkipRecoverRecoverableItems() bool {
	if o != nil && o.SkipRecoverRecoverableItems.IsSet() {
		return true
	}

	return false
}

// SetSkipRecoverRecoverableItems gets a reference to the given NullableBool and assigns it to the SkipRecoverRecoverableItems field.
func (o *RecoverMailboxParams) SetSkipRecoverRecoverableItems(v bool) {
	o.SkipRecoverRecoverableItems.Set(&v)
}
// SetSkipRecoverRecoverableItemsNil sets the value for SkipRecoverRecoverableItems to be an explicit nil
func (o *RecoverMailboxParams) SetSkipRecoverRecoverableItemsNil() {
	o.SkipRecoverRecoverableItems.Set(nil)
}

// UnsetSkipRecoverRecoverableItems ensures that no value is present for SkipRecoverRecoverableItems, not even an explicit nil
func (o *RecoverMailboxParams) UnsetSkipRecoverRecoverableItems() {
	o.SkipRecoverRecoverableItems.Unset()
}

// GetTargetMailbox returns the TargetMailbox field value if set, zero value otherwise.
func (o *RecoverMailboxParams) GetTargetMailbox() RecoverMailboxParamsTargetMailbox {
	if o == nil || IsNil(o.TargetMailbox) {
		var ret RecoverMailboxParamsTargetMailbox
		return ret
	}
	return *o.TargetMailbox
}

// GetTargetMailboxOk returns a tuple with the TargetMailbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverMailboxParams) GetTargetMailboxOk() (*RecoverMailboxParamsTargetMailbox, bool) {
	if o == nil || IsNil(o.TargetMailbox) {
		return nil, false
	}
	return o.TargetMailbox, true
}

// HasTargetMailbox returns a boolean if a field has been set.
func (o *RecoverMailboxParams) HasTargetMailbox() bool {
	if o != nil && !IsNil(o.TargetMailbox) {
		return true
	}

	return false
}

// SetTargetMailbox gets a reference to the given RecoverMailboxParamsTargetMailbox and assigns it to the TargetMailbox field.
func (o *RecoverMailboxParams) SetTargetMailbox(v RecoverMailboxParamsTargetMailbox) {
	o.TargetMailbox = &v
}

func (o RecoverMailboxParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverMailboxParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if !IsNil(o.PstParams) {
		toSerialize["pstParams"] = o.PstParams
	}
	if o.SkipRecoverArchiveMailbox.IsSet() {
		toSerialize["skipRecoverArchiveMailbox"] = o.SkipRecoverArchiveMailbox.Get()
	}
	if o.SkipRecoverArchiveRecoverableItems.IsSet() {
		toSerialize["skipRecoverArchiveRecoverableItems"] = o.SkipRecoverArchiveRecoverableItems.Get()
	}
	if o.SkipRecoverRecoverableItems.IsSet() {
		toSerialize["skipRecoverRecoverableItems"] = o.SkipRecoverRecoverableItems.Get()
	}
	if !IsNil(o.TargetMailbox) {
		toSerialize["targetMailbox"] = o.TargetMailbox
	}
	return toSerialize, nil
}

func (o *RecoverMailboxParams) UnmarshalJSON(data []byte) (err error) {
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

	varRecoverMailboxParams := _RecoverMailboxParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverMailboxParams)

	if err != nil {
		return err
	}

	*o = RecoverMailboxParams(varRecoverMailboxParams)

	return err
}

type NullableRecoverMailboxParams struct {
	value *RecoverMailboxParams
	isSet bool
}

func (v NullableRecoverMailboxParams) Get() *RecoverMailboxParams {
	return v.value
}

func (v *NullableRecoverMailboxParams) Set(val *RecoverMailboxParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverMailboxParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverMailboxParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverMailboxParams(val *RecoverMailboxParams) *NullableRecoverMailboxParams {
	return &NullableRecoverMailboxParams{value: val, isSet: true}
}

func (v NullableRecoverMailboxParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverMailboxParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


