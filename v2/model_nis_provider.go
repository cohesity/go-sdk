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

// checks if the NisProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NisProvider{}

// NisProvider Specifies an NIS Provider.
type NisProvider struct {
	// Specifies the Domain Name of NIS Provider.
	Domain NullableString `json:"domain"`
	// Specifies the hostname of Master Server.
	MasterServerHostname NullableString `json:"masterServerHostname"`
	// Specifies a list of slave servers in the NIS Domain.
	SlaveServers []string `json:"slaveServers,omitempty"`
	// Specifies the list of tenant Ids for NIS Provider.
	TenantIds []string `json:"tenantIds,omitempty"`
}

type _NisProvider NisProvider

// NewNisProvider instantiates a new NisProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNisProvider(domain NullableString, masterServerHostname NullableString) *NisProvider {
	this := NisProvider{}
	this.Domain = domain
	this.MasterServerHostname = masterServerHostname
	return &this
}

// NewNisProviderWithDefaults instantiates a new NisProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNisProviderWithDefaults() *NisProvider {
	this := NisProvider{}
	return &this
}

// GetDomain returns the Domain field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NisProvider) GetDomain() string {
	if o == nil || o.Domain.Get() == nil {
		var ret string
		return ret
	}

	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NisProvider) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// SetDomain sets field value
func (o *NisProvider) SetDomain(v string) {
	o.Domain.Set(&v)
}

// GetMasterServerHostname returns the MasterServerHostname field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NisProvider) GetMasterServerHostname() string {
	if o == nil || o.MasterServerHostname.Get() == nil {
		var ret string
		return ret
	}

	return *o.MasterServerHostname.Get()
}

// GetMasterServerHostnameOk returns a tuple with the MasterServerHostname field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NisProvider) GetMasterServerHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MasterServerHostname.Get(), o.MasterServerHostname.IsSet()
}

// SetMasterServerHostname sets field value
func (o *NisProvider) SetMasterServerHostname(v string) {
	o.MasterServerHostname.Set(&v)
}

// GetSlaveServers returns the SlaveServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NisProvider) GetSlaveServers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SlaveServers
}

// GetSlaveServersOk returns a tuple with the SlaveServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NisProvider) GetSlaveServersOk() ([]string, bool) {
	if o == nil || IsNil(o.SlaveServers) {
		return nil, false
	}
	return o.SlaveServers, true
}

// HasSlaveServers returns a boolean if a field has been set.
func (o *NisProvider) HasSlaveServers() bool {
	if o != nil && !IsNil(o.SlaveServers) {
		return true
	}

	return false
}

// SetSlaveServers gets a reference to the given []string and assigns it to the SlaveServers field.
func (o *NisProvider) SetSlaveServers(v []string) {
	o.SlaveServers = v
}

// GetTenantIds returns the TenantIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NisProvider) GetTenantIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TenantIds
}

// GetTenantIdsOk returns a tuple with the TenantIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NisProvider) GetTenantIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TenantIds) {
		return nil, false
	}
	return o.TenantIds, true
}

// HasTenantIds returns a boolean if a field has been set.
func (o *NisProvider) HasTenantIds() bool {
	if o != nil && !IsNil(o.TenantIds) {
		return true
	}

	return false
}

// SetTenantIds gets a reference to the given []string and assigns it to the TenantIds field.
func (o *NisProvider) SetTenantIds(v []string) {
	o.TenantIds = v
}

func (o NisProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NisProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domain"] = o.Domain.Get()
	toSerialize["masterServerHostname"] = o.MasterServerHostname.Get()
	if o.SlaveServers != nil {
		toSerialize["slaveServers"] = o.SlaveServers
	}
	if o.TenantIds != nil {
		toSerialize["tenantIds"] = o.TenantIds
	}
	return toSerialize, nil
}

func (o *NisProvider) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"domain",
		"masterServerHostname",
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

	varNisProvider := _NisProvider{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNisProvider)

	if err != nil {
		return err
	}

	*o = NisProvider(varNisProvider)

	return err
}

type NullableNisProvider struct {
	value *NisProvider
	isSet bool
}

func (v NullableNisProvider) Get() *NisProvider {
	return v.value
}

func (v *NullableNisProvider) Set(val *NisProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableNisProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableNisProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNisProvider(val *NisProvider) *NullableNisProvider {
	return &NullableNisProvider{value: val, isSet: true}
}

func (v NullableNisProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNisProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


