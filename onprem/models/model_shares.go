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

// Shares Specifies a list of Shares.
type Shares struct {
	// Specifies the list of shares.
	Shares []Share `json:"shares,omitempty"`
	// Specifies the pagination cookie.
	Cookie NullableString `json:"cookie,omitempty"`
}

// NewShares instantiates a new Shares object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShares() *Shares {
	this := Shares{}
	return &this
}

// NewSharesWithDefaults instantiates a new Shares object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharesWithDefaults() *Shares {
	this := Shares{}
	return &this
}

// GetShares returns the Shares field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Shares) GetShares() []Share {
	if o == nil  {
		var ret []Share
		return ret
	}
	return o.Shares
}

// GetSharesOk returns a tuple with the Shares field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Shares) GetSharesOk() (*[]Share, bool) {
	if o == nil || o.Shares == nil {
		return nil, false
	}
	return &o.Shares, true
}

// HasShares returns a boolean if a field has been set.
func (o *Shares) HasShares() bool {
	if o != nil && o.Shares != nil {
		return true
	}

	return false
}

// SetShares gets a reference to the given []Share and assigns it to the Shares field.
func (o *Shares) SetShares(v []Share) {
	o.Shares = v
}

// GetCookie returns the Cookie field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Shares) GetCookie() string {
	if o == nil || o.Cookie.Get() == nil {
		var ret string
		return ret
	}
	return *o.Cookie.Get()
}

// GetCookieOk returns a tuple with the Cookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Shares) GetCookieOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Cookie.Get(), o.Cookie.IsSet()
}

// HasCookie returns a boolean if a field has been set.
func (o *Shares) HasCookie() bool {
	if o != nil && o.Cookie.IsSet() {
		return true
	}

	return false
}

// SetCookie gets a reference to the given NullableString and assigns it to the Cookie field.
func (o *Shares) SetCookie(v string) {
	o.Cookie.Set(&v)
}
// SetCookieNil sets the value for Cookie to be an explicit nil
func (o *Shares) SetCookieNil() {
	o.Cookie.Set(nil)
}

// UnsetCookie ensures that no value is present for Cookie, not even an explicit nil
func (o *Shares) UnsetCookie() {
	o.Cookie.Unset()
}

func (o Shares) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Shares != nil {
		toSerialize["shares"] = o.Shares
	}
	if o.Cookie.IsSet() {
		toSerialize["cookie"] = o.Cookie.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableShares struct {
	value *Shares
	isSet bool
}

func (v NullableShares) Get() *Shares {
	return v.value
}

func (v *NullableShares) Set(val *Shares) {
	v.value = val
	v.isSet = true
}

func (v NullableShares) IsSet() bool {
	return v.isSet
}

func (v *NullableShares) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShares(val *Shares) *NullableShares {
	return &NullableShares{value: val, isSet: true}
}

func (v NullableShares) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShares) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o Shares) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}