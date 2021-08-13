/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package helios
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/helios/utils"
)

var _ = NullableBool{}

// SharepointObjectParams Specifies the common parameters for Sharepoint site objects.
type SharepointObjectParams struct {
	// Specifies the web url for the Sharepoint site.
	SiteWebUrl NullableString `json:"siteWebUrl,omitempty"`
}

// NewSharepointObjectParams instantiates a new SharepointObjectParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharepointObjectParams() *SharepointObjectParams {
	this := SharepointObjectParams{}
	return &this
}

// NewSharepointObjectParamsWithDefaults instantiates a new SharepointObjectParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharepointObjectParamsWithDefaults() *SharepointObjectParams {
	this := SharepointObjectParams{}
	return &this
}

// GetSiteWebUrl returns the SiteWebUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharepointObjectParams) GetSiteWebUrl() string {
	if o == nil || o.SiteWebUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.SiteWebUrl.Get()
}

// GetSiteWebUrlOk returns a tuple with the SiteWebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SharepointObjectParams) GetSiteWebUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SiteWebUrl.Get(), o.SiteWebUrl.IsSet()
}

// HasSiteWebUrl returns a boolean if a field has been set.
func (o *SharepointObjectParams) HasSiteWebUrl() bool {
	if o != nil && o.SiteWebUrl.IsSet() {
		return true
	}

	return false
}

// SetSiteWebUrl gets a reference to the given NullableString and assigns it to the SiteWebUrl field.
func (o *SharepointObjectParams) SetSiteWebUrl(v string) {
	o.SiteWebUrl.Set(&v)
}
// SetSiteWebUrlNil sets the value for SiteWebUrl to be an explicit nil
func (o *SharepointObjectParams) SetSiteWebUrlNil() {
	o.SiteWebUrl.Set(nil)
}

// UnsetSiteWebUrl ensures that no value is present for SiteWebUrl, not even an explicit nil
func (o *SharepointObjectParams) UnsetSiteWebUrl() {
	o.SiteWebUrl.Unset()
}

func (o SharepointObjectParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SiteWebUrl.IsSet() {
		toSerialize["siteWebUrl"] = o.SiteWebUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSharepointObjectParams struct {
	value *SharepointObjectParams
	isSet bool
}

func (v NullableSharepointObjectParams) Get() *SharepointObjectParams {
	return v.value
}

func (v *NullableSharepointObjectParams) Set(val *SharepointObjectParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSharepointObjectParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSharepointObjectParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharepointObjectParams(val *SharepointObjectParams) *NullableSharepointObjectParams {
	return &NullableSharepointObjectParams{value: val, isSet: true}
}

func (v NullableSharepointObjectParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharepointObjectParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


