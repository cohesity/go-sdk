/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the HBaseAdditionalParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HBaseAdditionalParams{}

// HBaseAdditionalParams Additional params for HBase protection source.
type HBaseAdditionalParams struct {
	// Authentication type.
	AuthType NullableString `json:"authType,omitempty"`
	// The 'Data root directory' for this HBase.
	DataRootDirectory NullableString `json:"dataRootDirectory,omitempty"`
	// The 'Zookeeper Quorum' for this HBase.
	ZookeeperQuorum []string `json:"zookeeperQuorum,omitempty"`
}

// NewHBaseAdditionalParams instantiates a new HBaseAdditionalParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHBaseAdditionalParams() *HBaseAdditionalParams {
	this := HBaseAdditionalParams{}
	return &this
}

// NewHBaseAdditionalParamsWithDefaults instantiates a new HBaseAdditionalParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHBaseAdditionalParamsWithDefaults() *HBaseAdditionalParams {
	this := HBaseAdditionalParams{}
	return &this
}

// GetAuthType returns the AuthType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HBaseAdditionalParams) GetAuthType() string {
	if o == nil || IsNil(o.AuthType.Get()) {
		var ret string
		return ret
	}
	return *o.AuthType.Get()
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HBaseAdditionalParams) GetAuthTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthType.Get(), o.AuthType.IsSet()
}

// HasAuthType returns a boolean if a field has been set.
func (o *HBaseAdditionalParams) HasAuthType() bool {
	if o != nil && o.AuthType.IsSet() {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given NullableString and assigns it to the AuthType field.
func (o *HBaseAdditionalParams) SetAuthType(v string) {
	o.AuthType.Set(&v)
}
// SetAuthTypeNil sets the value for AuthType to be an explicit nil
func (o *HBaseAdditionalParams) SetAuthTypeNil() {
	o.AuthType.Set(nil)
}

// UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil
func (o *HBaseAdditionalParams) UnsetAuthType() {
	o.AuthType.Unset()
}

// GetDataRootDirectory returns the DataRootDirectory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HBaseAdditionalParams) GetDataRootDirectory() string {
	if o == nil || IsNil(o.DataRootDirectory.Get()) {
		var ret string
		return ret
	}
	return *o.DataRootDirectory.Get()
}

// GetDataRootDirectoryOk returns a tuple with the DataRootDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HBaseAdditionalParams) GetDataRootDirectoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataRootDirectory.Get(), o.DataRootDirectory.IsSet()
}

// HasDataRootDirectory returns a boolean if a field has been set.
func (o *HBaseAdditionalParams) HasDataRootDirectory() bool {
	if o != nil && o.DataRootDirectory.IsSet() {
		return true
	}

	return false
}

// SetDataRootDirectory gets a reference to the given NullableString and assigns it to the DataRootDirectory field.
func (o *HBaseAdditionalParams) SetDataRootDirectory(v string) {
	o.DataRootDirectory.Set(&v)
}
// SetDataRootDirectoryNil sets the value for DataRootDirectory to be an explicit nil
func (o *HBaseAdditionalParams) SetDataRootDirectoryNil() {
	o.DataRootDirectory.Set(nil)
}

// UnsetDataRootDirectory ensures that no value is present for DataRootDirectory, not even an explicit nil
func (o *HBaseAdditionalParams) UnsetDataRootDirectory() {
	o.DataRootDirectory.Unset()
}

// GetZookeeperQuorum returns the ZookeeperQuorum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HBaseAdditionalParams) GetZookeeperQuorum() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ZookeeperQuorum
}

// GetZookeeperQuorumOk returns a tuple with the ZookeeperQuorum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HBaseAdditionalParams) GetZookeeperQuorumOk() ([]string, bool) {
	if o == nil || IsNil(o.ZookeeperQuorum) {
		return nil, false
	}
	return o.ZookeeperQuorum, true
}

// HasZookeeperQuorum returns a boolean if a field has been set.
func (o *HBaseAdditionalParams) HasZookeeperQuorum() bool {
	if o != nil && !IsNil(o.ZookeeperQuorum) {
		return true
	}

	return false
}

// SetZookeeperQuorum gets a reference to the given []string and assigns it to the ZookeeperQuorum field.
func (o *HBaseAdditionalParams) SetZookeeperQuorum(v []string) {
	o.ZookeeperQuorum = v
}

func (o HBaseAdditionalParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HBaseAdditionalParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthType.IsSet() {
		toSerialize["authType"] = o.AuthType.Get()
	}
	if o.DataRootDirectory.IsSet() {
		toSerialize["dataRootDirectory"] = o.DataRootDirectory.Get()
	}
	if o.ZookeeperQuorum != nil {
		toSerialize["zookeeperQuorum"] = o.ZookeeperQuorum
	}
	return toSerialize, nil
}

type NullableHBaseAdditionalParams struct {
	value *HBaseAdditionalParams
	isSet bool
}

func (v NullableHBaseAdditionalParams) Get() *HBaseAdditionalParams {
	return v.value
}

func (v *NullableHBaseAdditionalParams) Set(val *HBaseAdditionalParams) {
	v.value = val
	v.isSet = true
}

func (v NullableHBaseAdditionalParams) IsSet() bool {
	return v.isSet
}

func (v *NullableHBaseAdditionalParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHBaseAdditionalParams(val *HBaseAdditionalParams) *NullableHBaseAdditionalParams {
	return &NullableHBaseAdditionalParams{value: val, isSet: true}
}

func (v NullableHBaseAdditionalParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHBaseAdditionalParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


