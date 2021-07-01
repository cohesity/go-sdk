/*
 * Cohesity REST API
 *
 * This API list provides operations for interfacing with the Cohesity Cluster.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesitysdk

import (
	"encoding/json"
)

// HostEntry Specifies the parameters of a host entry that can be stored in the cluster's /etc/hosts file.
type HostEntry struct {
	// Description the host entry.
	Description NullableString `json:"description,omitempty"`
	// Specifies the domain names of the host.
	DomainNames []string `json:"domainNames"`
	// Specifies the IP address of the host.
	Ip NullableString `json:"ip"`
}

// NewHostEntry instantiates a new HostEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostEntry(domainNames []string, ip NullableString) *HostEntry {
	this := HostEntry{}
	this.DomainNames = domainNames
	this.Ip = ip
	return &this
}

// NewHostEntryWithDefaults instantiates a new HostEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostEntryWithDefaults() *HostEntry {
	this := HostEntry{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostEntry) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostEntry) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *HostEntry) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *HostEntry) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *HostEntry) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *HostEntry) UnsetDescription() {
	o.Description.Unset()
}

// GetDomainNames returns the DomainNames field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *HostEntry) GetDomainNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DomainNames
}

// GetDomainNamesOk returns a tuple with the DomainNames field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostEntry) GetDomainNamesOk() (*[]string, bool) {
	if o == nil || o.DomainNames == nil {
		return nil, false
	}
	return &o.DomainNames, true
}

// SetDomainNames sets field value
func (o *HostEntry) SetDomainNames(v []string) {
	o.DomainNames = v
}

// GetIp returns the Ip field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HostEntry) GetIp() string {
	if o == nil || o.Ip.Get() == nil {
		var ret string
		return ret
	}

	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostEntry) GetIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// SetIp sets field value
func (o *HostEntry) SetIp(v string) {
	o.Ip.Set(&v)
}

func (o HostEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DomainNames != nil {
		toSerialize["domainNames"] = o.DomainNames
	}
	if true {
		toSerialize["ip"] = o.Ip.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableHostEntry struct {
	value *HostEntry
	isSet bool
}

func (v NullableHostEntry) Get() *HostEntry {
	return v.value
}

func (v *NullableHostEntry) Set(val *HostEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableHostEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableHostEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostEntry(val *HostEntry) *NullableHostEntry {
	return &NullableHostEntry{value: val, isSet: true}
}

func (v NullableHostEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


