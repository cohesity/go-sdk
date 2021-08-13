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

// RecoverMailboxParams Specifies the parameters to recover an Office 365 Mailbox.
type RecoverMailboxParams struct {
	// Specifies a list of Mailbox params associated with the objects to recover.
	Objects []ObjectMailboxParam `json:"objects"`
	// Specifies the target Mailbox to recover to. If not specified, the objects will be recovered to original location.
	TargetMailbox *TargetMailboxParam `json:"targetMailbox,omitempty"`
	// Specifies whether to continue recovering other Mailboxes if one of Mailbox failed to recover. Default value is false.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
}

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
func (o *RecoverMailboxParams) GetObjectsOk() (*[]ObjectMailboxParam, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return &o.Objects, true
}

// SetObjects sets field value
func (o *RecoverMailboxParams) SetObjects(v []ObjectMailboxParam) {
	o.Objects = v
}

// GetTargetMailbox returns the TargetMailbox field value if set, zero value otherwise.
func (o *RecoverMailboxParams) GetTargetMailbox() TargetMailboxParam {
	if o == nil || o.TargetMailbox == nil {
		var ret TargetMailboxParam
		return ret
	}
	return *o.TargetMailbox
}

// GetTargetMailboxOk returns a tuple with the TargetMailbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverMailboxParams) GetTargetMailboxOk() (*TargetMailboxParam, bool) {
	if o == nil || o.TargetMailbox == nil {
		return nil, false
	}
	return o.TargetMailbox, true
}

// HasTargetMailbox returns a boolean if a field has been set.
func (o *RecoverMailboxParams) HasTargetMailbox() bool {
	if o != nil && o.TargetMailbox != nil {
		return true
	}

	return false
}

// SetTargetMailbox gets a reference to the given TargetMailboxParam and assigns it to the TargetMailbox field.
func (o *RecoverMailboxParams) SetTargetMailbox(v TargetMailboxParam) {
	o.TargetMailbox = &v
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverMailboxParams) GetContinueOnError() bool {
	if o == nil || o.ContinueOnError.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverMailboxParams) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil  {
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

func (o RecoverMailboxParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if o.TargetMailbox != nil {
		toSerialize["targetMailbox"] = o.TargetMailbox
	}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	return json.Marshal(toSerialize)
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




func (o RecoverMailboxParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}