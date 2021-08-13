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

// Office365SharepointSiteObjectProtectionParams Specifies the params to create a Sharepoint site Object Protection.
type Office365SharepointSiteObjectProtectionParams struct {
	// Specifies the objects to be included in the Object Protection.
	Objects []Office365ObjectProtectionObjectParams `json:"objects"`
	IndexingPolicy *IndexingPolicy `json:"indexingPolicy,omitempty"`
	// Specifies the id of the parent of the objects.
	SourceId NullableInt64 `json:"sourceId,omitempty"`
	// Specifies the name of the parent of the objects.
	SourceName NullableString `json:"sourceName,omitempty"`
}

// NewOffice365SharepointSiteObjectProtectionParams instantiates a new Office365SharepointSiteObjectProtectionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOffice365SharepointSiteObjectProtectionParams(objects []Office365ObjectProtectionObjectParams) *Office365SharepointSiteObjectProtectionParams {
	this := Office365SharepointSiteObjectProtectionParams{}
	this.Objects = objects
	return &this
}

// NewOffice365SharepointSiteObjectProtectionParamsWithDefaults instantiates a new Office365SharepointSiteObjectProtectionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOffice365SharepointSiteObjectProtectionParamsWithDefaults() *Office365SharepointSiteObjectProtectionParams {
	this := Office365SharepointSiteObjectProtectionParams{}
	return &this
}

// GetObjects returns the Objects field value
func (o *Office365SharepointSiteObjectProtectionParams) GetObjects() []Office365ObjectProtectionObjectParams {
	if o == nil {
		var ret []Office365ObjectProtectionObjectParams
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *Office365SharepointSiteObjectProtectionParams) GetObjectsOk() (*[]Office365ObjectProtectionObjectParams, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Objects, true
}

// SetObjects sets field value
func (o *Office365SharepointSiteObjectProtectionParams) SetObjects(v []Office365ObjectProtectionObjectParams) {
	o.Objects = v
}

// GetIndexingPolicy returns the IndexingPolicy field value if set, zero value otherwise.
func (o *Office365SharepointSiteObjectProtectionParams) GetIndexingPolicy() IndexingPolicy {
	if o == nil || o.IndexingPolicy == nil {
		var ret IndexingPolicy
		return ret
	}
	return *o.IndexingPolicy
}

// GetIndexingPolicyOk returns a tuple with the IndexingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Office365SharepointSiteObjectProtectionParams) GetIndexingPolicyOk() (*IndexingPolicy, bool) {
	if o == nil || o.IndexingPolicy == nil {
		return nil, false
	}
	return o.IndexingPolicy, true
}

// HasIndexingPolicy returns a boolean if a field has been set.
func (o *Office365SharepointSiteObjectProtectionParams) HasIndexingPolicy() bool {
	if o != nil && o.IndexingPolicy != nil {
		return true
	}

	return false
}

// SetIndexingPolicy gets a reference to the given IndexingPolicy and assigns it to the IndexingPolicy field.
func (o *Office365SharepointSiteObjectProtectionParams) SetIndexingPolicy(v IndexingPolicy) {
	o.IndexingPolicy = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Office365SharepointSiteObjectProtectionParams) GetSourceId() int64 {
	if o == nil || o.SourceId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SourceId.Get()
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Office365SharepointSiteObjectProtectionParams) GetSourceIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceId.Get(), o.SourceId.IsSet()
}

// HasSourceId returns a boolean if a field has been set.
func (o *Office365SharepointSiteObjectProtectionParams) HasSourceId() bool {
	if o != nil && o.SourceId.IsSet() {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given NullableInt64 and assigns it to the SourceId field.
func (o *Office365SharepointSiteObjectProtectionParams) SetSourceId(v int64) {
	o.SourceId.Set(&v)
}
// SetSourceIdNil sets the value for SourceId to be an explicit nil
func (o *Office365SharepointSiteObjectProtectionParams) SetSourceIdNil() {
	o.SourceId.Set(nil)
}

// UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
func (o *Office365SharepointSiteObjectProtectionParams) UnsetSourceId() {
	o.SourceId.Unset()
}

// GetSourceName returns the SourceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Office365SharepointSiteObjectProtectionParams) GetSourceName() string {
	if o == nil || o.SourceName.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceName.Get()
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Office365SharepointSiteObjectProtectionParams) GetSourceNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceName.Get(), o.SourceName.IsSet()
}

// HasSourceName returns a boolean if a field has been set.
func (o *Office365SharepointSiteObjectProtectionParams) HasSourceName() bool {
	if o != nil && o.SourceName.IsSet() {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given NullableString and assigns it to the SourceName field.
func (o *Office365SharepointSiteObjectProtectionParams) SetSourceName(v string) {
	o.SourceName.Set(&v)
}
// SetSourceNameNil sets the value for SourceName to be an explicit nil
func (o *Office365SharepointSiteObjectProtectionParams) SetSourceNameNil() {
	o.SourceName.Set(nil)
}

// UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
func (o *Office365SharepointSiteObjectProtectionParams) UnsetSourceName() {
	o.SourceName.Unset()
}

func (o Office365SharepointSiteObjectProtectionParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objects"] = o.Objects
	}
	if o.IndexingPolicy != nil {
		toSerialize["indexingPolicy"] = o.IndexingPolicy
	}
	if o.SourceId.IsSet() {
		toSerialize["sourceId"] = o.SourceId.Get()
	}
	if o.SourceName.IsSet() {
		toSerialize["sourceName"] = o.SourceName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOffice365SharepointSiteObjectProtectionParams struct {
	value *Office365SharepointSiteObjectProtectionParams
	isSet bool
}

func (v NullableOffice365SharepointSiteObjectProtectionParams) Get() *Office365SharepointSiteObjectProtectionParams {
	return v.value
}

func (v *NullableOffice365SharepointSiteObjectProtectionParams) Set(val *Office365SharepointSiteObjectProtectionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableOffice365SharepointSiteObjectProtectionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableOffice365SharepointSiteObjectProtectionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOffice365SharepointSiteObjectProtectionParams(val *Office365SharepointSiteObjectProtectionParams) *NullableOffice365SharepointSiteObjectProtectionParams {
	return &NullableOffice365SharepointSiteObjectProtectionParams{value: val, isSet: true}
}

func (v NullableOffice365SharepointSiteObjectProtectionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOffice365SharepointSiteObjectProtectionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o Office365SharepointSiteObjectProtectionParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}