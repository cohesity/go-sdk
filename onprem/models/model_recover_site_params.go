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

// RecoverSiteParams Specifies the parameters to recover Microsoft Office 365 Sharepoint site.
type RecoverSiteParams struct {
	// Specifies a list of site params associated with the objects to recover.
	Objects []ObjectSiteParam `json:"objects"`
	// Specifies the target Site to recover to. If not specified, the objects will be recovered to original location.
	TargetSite *TargetSiteParam `json:"targetSite,omitempty"`
	// Specifies the object id of the target domain in case of full recovery of a site to a target domain.
	TargetDomainObjectId NullableRecoveryObjectIdentifier `json:"targetDomainObjectId,omitempty"`
	// Specifies whether to continue recovering the doc libs of a site, if one or more of doc libs failed to recover. Default value is false.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
}

// NewRecoverSiteParams instantiates a new RecoverSiteParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverSiteParams(objects []ObjectSiteParam) *RecoverSiteParams {
	this := RecoverSiteParams{}
	this.Objects = objects
	return &this
}

// NewRecoverSiteParamsWithDefaults instantiates a new RecoverSiteParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverSiteParamsWithDefaults() *RecoverSiteParams {
	this := RecoverSiteParams{}
	return &this
}

// GetObjects returns the Objects field value
// If the value is explicit nil, the zero value for []ObjectSiteParam will be returned
func (o *RecoverSiteParams) GetObjects() []ObjectSiteParam {
	if o == nil {
		var ret []ObjectSiteParam
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverSiteParams) GetObjectsOk() (*[]ObjectSiteParam, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return &o.Objects, true
}

// SetObjects sets field value
func (o *RecoverSiteParams) SetObjects(v []ObjectSiteParam) {
	o.Objects = v
}

// GetTargetSite returns the TargetSite field value if set, zero value otherwise.
func (o *RecoverSiteParams) GetTargetSite() TargetSiteParam {
	if o == nil || o.TargetSite == nil {
		var ret TargetSiteParam
		return ret
	}
	return *o.TargetSite
}

// GetTargetSiteOk returns a tuple with the TargetSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverSiteParams) GetTargetSiteOk() (*TargetSiteParam, bool) {
	if o == nil || o.TargetSite == nil {
		return nil, false
	}
	return o.TargetSite, true
}

// HasTargetSite returns a boolean if a field has been set.
func (o *RecoverSiteParams) HasTargetSite() bool {
	if o != nil && o.TargetSite != nil {
		return true
	}

	return false
}

// SetTargetSite gets a reference to the given TargetSiteParam and assigns it to the TargetSite field.
func (o *RecoverSiteParams) SetTargetSite(v TargetSiteParam) {
	o.TargetSite = &v
}

// GetTargetDomainObjectId returns the TargetDomainObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverSiteParams) GetTargetDomainObjectId() RecoveryObjectIdentifier {
	if o == nil || o.TargetDomainObjectId.Get() == nil {
		var ret RecoveryObjectIdentifier
		return ret
	}
	return *o.TargetDomainObjectId.Get()
}

// GetTargetDomainObjectIdOk returns a tuple with the TargetDomainObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverSiteParams) GetTargetDomainObjectIdOk() (*RecoveryObjectIdentifier, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetDomainObjectId.Get(), o.TargetDomainObjectId.IsSet()
}

// HasTargetDomainObjectId returns a boolean if a field has been set.
func (o *RecoverSiteParams) HasTargetDomainObjectId() bool {
	if o != nil && o.TargetDomainObjectId.IsSet() {
		return true
	}

	return false
}

// SetTargetDomainObjectId gets a reference to the given NullableRecoveryObjectIdentifier and assigns it to the TargetDomainObjectId field.
func (o *RecoverSiteParams) SetTargetDomainObjectId(v RecoveryObjectIdentifier) {
	o.TargetDomainObjectId.Set(&v)
}
// SetTargetDomainObjectIdNil sets the value for TargetDomainObjectId to be an explicit nil
func (o *RecoverSiteParams) SetTargetDomainObjectIdNil() {
	o.TargetDomainObjectId.Set(nil)
}

// UnsetTargetDomainObjectId ensures that no value is present for TargetDomainObjectId, not even an explicit nil
func (o *RecoverSiteParams) UnsetTargetDomainObjectId() {
	o.TargetDomainObjectId.Unset()
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverSiteParams) GetContinueOnError() bool {
	if o == nil || o.ContinueOnError.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverSiteParams) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *RecoverSiteParams) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *RecoverSiteParams) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *RecoverSiteParams) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *RecoverSiteParams) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

func (o RecoverSiteParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if o.TargetSite != nil {
		toSerialize["targetSite"] = o.TargetSite
	}
	if o.TargetDomainObjectId.IsSet() {
		toSerialize["targetDomainObjectId"] = o.TargetDomainObjectId.Get()
	}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverSiteParams struct {
	value *RecoverSiteParams
	isSet bool
}

func (v NullableRecoverSiteParams) Get() *RecoverSiteParams {
	return v.value
}

func (v *NullableRecoverSiteParams) Set(val *RecoverSiteParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverSiteParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverSiteParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverSiteParams(val *RecoverSiteParams) *NullableRecoverSiteParams {
	return &NullableRecoverSiteParams{value: val, isSet: true}
}

func (v NullableRecoverSiteParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverSiteParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RecoverSiteParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}