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

// CassandraSourceRegistrationParamsAllOf struct for CassandraSourceRegistrationParamsAllOf
type CassandraSourceRegistrationParamsAllOf struct {
	JmxCredentials NullableCassandraSourceRegistrationParamsAllOfJmxCredentials `json:"jmxCredentials,omitempty"`
	CassandraCredentials NullableCassandraSourceRegistrationParamsAllOfCassandraCredentials `json:"cassandraCredentials,omitempty"`
	// Data centers for this cluster.
	DataCenterNames *[]string `json:"dataCenterNames,omitempty"`
	// Principal for the kerberos connection. (This is required only if your Cassandra has Kerberos authentication. Please refer to the user guide.)
	KerberosPrincipal NullableString `json:"kerberosPrincipal,omitempty"`
	DseSolrInfo *DSESolrInfo `json:"dseSolrInfo,omitempty"`
}

// NewCassandraSourceRegistrationParamsAllOf instantiates a new CassandraSourceRegistrationParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCassandraSourceRegistrationParamsAllOf() *CassandraSourceRegistrationParamsAllOf {
	this := CassandraSourceRegistrationParamsAllOf{}
	return &this
}

// NewCassandraSourceRegistrationParamsAllOfWithDefaults instantiates a new CassandraSourceRegistrationParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCassandraSourceRegistrationParamsAllOfWithDefaults() *CassandraSourceRegistrationParamsAllOf {
	this := CassandraSourceRegistrationParamsAllOf{}
	return &this
}

// GetJmxCredentials returns the JmxCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceRegistrationParamsAllOf) GetJmxCredentials() CassandraSourceRegistrationParamsAllOfJmxCredentials {
	if o == nil || o.JmxCredentials.Get() == nil {
		var ret CassandraSourceRegistrationParamsAllOfJmxCredentials
		return ret
	}
	return *o.JmxCredentials.Get()
}

// GetJmxCredentialsOk returns a tuple with the JmxCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceRegistrationParamsAllOf) GetJmxCredentialsOk() (*CassandraSourceRegistrationParamsAllOfJmxCredentials, bool) {
	if o == nil  {
		return nil, false
	}
	return o.JmxCredentials.Get(), o.JmxCredentials.IsSet()
}

// HasJmxCredentials returns a boolean if a field has been set.
func (o *CassandraSourceRegistrationParamsAllOf) HasJmxCredentials() bool {
	if o != nil && o.JmxCredentials.IsSet() {
		return true
	}

	return false
}

// SetJmxCredentials gets a reference to the given NullableCassandraSourceRegistrationParamsAllOfJmxCredentials and assigns it to the JmxCredentials field.
func (o *CassandraSourceRegistrationParamsAllOf) SetJmxCredentials(v CassandraSourceRegistrationParamsAllOfJmxCredentials) {
	o.JmxCredentials.Set(&v)
}
// SetJmxCredentialsNil sets the value for JmxCredentials to be an explicit nil
func (o *CassandraSourceRegistrationParamsAllOf) SetJmxCredentialsNil() {
	o.JmxCredentials.Set(nil)
}

// UnsetJmxCredentials ensures that no value is present for JmxCredentials, not even an explicit nil
func (o *CassandraSourceRegistrationParamsAllOf) UnsetJmxCredentials() {
	o.JmxCredentials.Unset()
}

// GetCassandraCredentials returns the CassandraCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceRegistrationParamsAllOf) GetCassandraCredentials() CassandraSourceRegistrationParamsAllOfCassandraCredentials {
	if o == nil || o.CassandraCredentials.Get() == nil {
		var ret CassandraSourceRegistrationParamsAllOfCassandraCredentials
		return ret
	}
	return *o.CassandraCredentials.Get()
}

// GetCassandraCredentialsOk returns a tuple with the CassandraCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceRegistrationParamsAllOf) GetCassandraCredentialsOk() (*CassandraSourceRegistrationParamsAllOfCassandraCredentials, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CassandraCredentials.Get(), o.CassandraCredentials.IsSet()
}

// HasCassandraCredentials returns a boolean if a field has been set.
func (o *CassandraSourceRegistrationParamsAllOf) HasCassandraCredentials() bool {
	if o != nil && o.CassandraCredentials.IsSet() {
		return true
	}

	return false
}

// SetCassandraCredentials gets a reference to the given NullableCassandraSourceRegistrationParamsAllOfCassandraCredentials and assigns it to the CassandraCredentials field.
func (o *CassandraSourceRegistrationParamsAllOf) SetCassandraCredentials(v CassandraSourceRegistrationParamsAllOfCassandraCredentials) {
	o.CassandraCredentials.Set(&v)
}
// SetCassandraCredentialsNil sets the value for CassandraCredentials to be an explicit nil
func (o *CassandraSourceRegistrationParamsAllOf) SetCassandraCredentialsNil() {
	o.CassandraCredentials.Set(nil)
}

// UnsetCassandraCredentials ensures that no value is present for CassandraCredentials, not even an explicit nil
func (o *CassandraSourceRegistrationParamsAllOf) UnsetCassandraCredentials() {
	o.CassandraCredentials.Unset()
}

// GetDataCenterNames returns the DataCenterNames field value if set, zero value otherwise.
func (o *CassandraSourceRegistrationParamsAllOf) GetDataCenterNames() []string {
	if o == nil || o.DataCenterNames == nil {
		var ret []string
		return ret
	}
	return *o.DataCenterNames
}

// GetDataCenterNamesOk returns a tuple with the DataCenterNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CassandraSourceRegistrationParamsAllOf) GetDataCenterNamesOk() (*[]string, bool) {
	if o == nil || o.DataCenterNames == nil {
		return nil, false
	}
	return o.DataCenterNames, true
}

// HasDataCenterNames returns a boolean if a field has been set.
func (o *CassandraSourceRegistrationParamsAllOf) HasDataCenterNames() bool {
	if o != nil && o.DataCenterNames != nil {
		return true
	}

	return false
}

// SetDataCenterNames gets a reference to the given []string and assigns it to the DataCenterNames field.
func (o *CassandraSourceRegistrationParamsAllOf) SetDataCenterNames(v []string) {
	o.DataCenterNames = &v
}

// GetKerberosPrincipal returns the KerberosPrincipal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceRegistrationParamsAllOf) GetKerberosPrincipal() string {
	if o == nil || o.KerberosPrincipal.Get() == nil {
		var ret string
		return ret
	}
	return *o.KerberosPrincipal.Get()
}

// GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceRegistrationParamsAllOf) GetKerberosPrincipalOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.KerberosPrincipal.Get(), o.KerberosPrincipal.IsSet()
}

// HasKerberosPrincipal returns a boolean if a field has been set.
func (o *CassandraSourceRegistrationParamsAllOf) HasKerberosPrincipal() bool {
	if o != nil && o.KerberosPrincipal.IsSet() {
		return true
	}

	return false
}

// SetKerberosPrincipal gets a reference to the given NullableString and assigns it to the KerberosPrincipal field.
func (o *CassandraSourceRegistrationParamsAllOf) SetKerberosPrincipal(v string) {
	o.KerberosPrincipal.Set(&v)
}
// SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil
func (o *CassandraSourceRegistrationParamsAllOf) SetKerberosPrincipalNil() {
	o.KerberosPrincipal.Set(nil)
}

// UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil
func (o *CassandraSourceRegistrationParamsAllOf) UnsetKerberosPrincipal() {
	o.KerberosPrincipal.Unset()
}

// GetDseSolrInfo returns the DseSolrInfo field value if set, zero value otherwise.
func (o *CassandraSourceRegistrationParamsAllOf) GetDseSolrInfo() DSESolrInfo {
	if o == nil || o.DseSolrInfo == nil {
		var ret DSESolrInfo
		return ret
	}
	return *o.DseSolrInfo
}

// GetDseSolrInfoOk returns a tuple with the DseSolrInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CassandraSourceRegistrationParamsAllOf) GetDseSolrInfoOk() (*DSESolrInfo, bool) {
	if o == nil || o.DseSolrInfo == nil {
		return nil, false
	}
	return o.DseSolrInfo, true
}

// HasDseSolrInfo returns a boolean if a field has been set.
func (o *CassandraSourceRegistrationParamsAllOf) HasDseSolrInfo() bool {
	if o != nil && o.DseSolrInfo != nil {
		return true
	}

	return false
}

// SetDseSolrInfo gets a reference to the given DSESolrInfo and assigns it to the DseSolrInfo field.
func (o *CassandraSourceRegistrationParamsAllOf) SetDseSolrInfo(v DSESolrInfo) {
	o.DseSolrInfo = &v
}

func (o CassandraSourceRegistrationParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JmxCredentials.IsSet() {
		toSerialize["jmxCredentials"] = o.JmxCredentials.Get()
	}
	if o.CassandraCredentials.IsSet() {
		toSerialize["cassandraCredentials"] = o.CassandraCredentials.Get()
	}
	if o.DataCenterNames != nil {
		toSerialize["dataCenterNames"] = o.DataCenterNames
	}
	if o.KerberosPrincipal.IsSet() {
		toSerialize["kerberosPrincipal"] = o.KerberosPrincipal.Get()
	}
	if o.DseSolrInfo != nil {
		toSerialize["dseSolrInfo"] = o.DseSolrInfo
	}
	return json.Marshal(toSerialize)
}

type NullableCassandraSourceRegistrationParamsAllOf struct {
	value *CassandraSourceRegistrationParamsAllOf
	isSet bool
}

func (v NullableCassandraSourceRegistrationParamsAllOf) Get() *CassandraSourceRegistrationParamsAllOf {
	return v.value
}

func (v *NullableCassandraSourceRegistrationParamsAllOf) Set(val *CassandraSourceRegistrationParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCassandraSourceRegistrationParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCassandraSourceRegistrationParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCassandraSourceRegistrationParamsAllOf(val *CassandraSourceRegistrationParamsAllOf) *NullableCassandraSourceRegistrationParamsAllOf {
	return &NullableCassandraSourceRegistrationParamsAllOf{value: val, isSet: true}
}

func (v NullableCassandraSourceRegistrationParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCassandraSourceRegistrationParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o CassandraSourceRegistrationParamsAllOf) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}