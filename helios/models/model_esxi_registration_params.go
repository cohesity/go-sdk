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

// EsxiRegistrationParams Specifies parameters to register VMware ESXi host.
type EsxiRegistrationParams struct {
	// Specifies the username to access target entity.
	Username string `json:"username"`
	// Specifies the password to access target entity.
	Password string `json:"password"`
	// Specifies the endpoint IPaddress, URL or hostname of the host.
	Endpoint string `json:"endpoint"`
	// Specifies the description of the source being registered.
	Description NullableString `json:"description,omitempty"`
	// Specifies the minimum free space (in GB) expected to be available in the datastore where the virtual disks of the VM being backed up reside. If the space available is lower than the specified value, backup will be aborted.
	MinFreeDatastoreSpaceForBackupGb NullableInt64 `json:"minFreeDatastoreSpaceForBackupGb,omitempty"`
	// If this value is > 0 and the number of streams concurrently active on a datastore is equal to it, then any further requests to access the datastore would be denied until the number of active streams reduces. This applies for all the datastores in the specified host.
	MaxConcurrentStreams NullableInt32 `json:"maxConcurrentStreams,omitempty"`
	// Specifies the datastore specific params.
	DataStoreParams []DatastoreParams `json:"dataStoreParams,omitempty"`
}

// NewEsxiRegistrationParams instantiates a new EsxiRegistrationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsxiRegistrationParams(username string, password string, endpoint string) *EsxiRegistrationParams {
	this := EsxiRegistrationParams{}
	this.Username = username
	this.Password = password
	this.Endpoint = endpoint
	return &this
}

// NewEsxiRegistrationParamsWithDefaults instantiates a new EsxiRegistrationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsxiRegistrationParamsWithDefaults() *EsxiRegistrationParams {
	this := EsxiRegistrationParams{}
	return &this
}

// GetUsername returns the Username field value
func (o *EsxiRegistrationParams) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *EsxiRegistrationParams) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *EsxiRegistrationParams) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *EsxiRegistrationParams) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *EsxiRegistrationParams) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *EsxiRegistrationParams) SetPassword(v string) {
	o.Password = v
}

// GetEndpoint returns the Endpoint field value
func (o *EsxiRegistrationParams) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *EsxiRegistrationParams) GetEndpointOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *EsxiRegistrationParams) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EsxiRegistrationParams) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EsxiRegistrationParams) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *EsxiRegistrationParams) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *EsxiRegistrationParams) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *EsxiRegistrationParams) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *EsxiRegistrationParams) UnsetDescription() {
	o.Description.Unset()
}

// GetMinFreeDatastoreSpaceForBackupGb returns the MinFreeDatastoreSpaceForBackupGb field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EsxiRegistrationParams) GetMinFreeDatastoreSpaceForBackupGb() int64 {
	if o == nil || o.MinFreeDatastoreSpaceForBackupGb.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MinFreeDatastoreSpaceForBackupGb.Get()
}

// GetMinFreeDatastoreSpaceForBackupGbOk returns a tuple with the MinFreeDatastoreSpaceForBackupGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EsxiRegistrationParams) GetMinFreeDatastoreSpaceForBackupGbOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MinFreeDatastoreSpaceForBackupGb.Get(), o.MinFreeDatastoreSpaceForBackupGb.IsSet()
}

// HasMinFreeDatastoreSpaceForBackupGb returns a boolean if a field has been set.
func (o *EsxiRegistrationParams) HasMinFreeDatastoreSpaceForBackupGb() bool {
	if o != nil && o.MinFreeDatastoreSpaceForBackupGb.IsSet() {
		return true
	}

	return false
}

// SetMinFreeDatastoreSpaceForBackupGb gets a reference to the given NullableInt64 and assigns it to the MinFreeDatastoreSpaceForBackupGb field.
func (o *EsxiRegistrationParams) SetMinFreeDatastoreSpaceForBackupGb(v int64) {
	o.MinFreeDatastoreSpaceForBackupGb.Set(&v)
}
// SetMinFreeDatastoreSpaceForBackupGbNil sets the value for MinFreeDatastoreSpaceForBackupGb to be an explicit nil
func (o *EsxiRegistrationParams) SetMinFreeDatastoreSpaceForBackupGbNil() {
	o.MinFreeDatastoreSpaceForBackupGb.Set(nil)
}

// UnsetMinFreeDatastoreSpaceForBackupGb ensures that no value is present for MinFreeDatastoreSpaceForBackupGb, not even an explicit nil
func (o *EsxiRegistrationParams) UnsetMinFreeDatastoreSpaceForBackupGb() {
	o.MinFreeDatastoreSpaceForBackupGb.Unset()
}

// GetMaxConcurrentStreams returns the MaxConcurrentStreams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EsxiRegistrationParams) GetMaxConcurrentStreams() int32 {
	if o == nil || o.MaxConcurrentStreams.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxConcurrentStreams.Get()
}

// GetMaxConcurrentStreamsOk returns a tuple with the MaxConcurrentStreams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EsxiRegistrationParams) GetMaxConcurrentStreamsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxConcurrentStreams.Get(), o.MaxConcurrentStreams.IsSet()
}

// HasMaxConcurrentStreams returns a boolean if a field has been set.
func (o *EsxiRegistrationParams) HasMaxConcurrentStreams() bool {
	if o != nil && o.MaxConcurrentStreams.IsSet() {
		return true
	}

	return false
}

// SetMaxConcurrentStreams gets a reference to the given NullableInt32 and assigns it to the MaxConcurrentStreams field.
func (o *EsxiRegistrationParams) SetMaxConcurrentStreams(v int32) {
	o.MaxConcurrentStreams.Set(&v)
}
// SetMaxConcurrentStreamsNil sets the value for MaxConcurrentStreams to be an explicit nil
func (o *EsxiRegistrationParams) SetMaxConcurrentStreamsNil() {
	o.MaxConcurrentStreams.Set(nil)
}

// UnsetMaxConcurrentStreams ensures that no value is present for MaxConcurrentStreams, not even an explicit nil
func (o *EsxiRegistrationParams) UnsetMaxConcurrentStreams() {
	o.MaxConcurrentStreams.Unset()
}

// GetDataStoreParams returns the DataStoreParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EsxiRegistrationParams) GetDataStoreParams() []DatastoreParams {
	if o == nil  {
		var ret []DatastoreParams
		return ret
	}
	return o.DataStoreParams
}

// GetDataStoreParamsOk returns a tuple with the DataStoreParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EsxiRegistrationParams) GetDataStoreParamsOk() (*[]DatastoreParams, bool) {
	if o == nil || o.DataStoreParams == nil {
		return nil, false
	}
	return &o.DataStoreParams, true
}

// HasDataStoreParams returns a boolean if a field has been set.
func (o *EsxiRegistrationParams) HasDataStoreParams() bool {
	if o != nil && o.DataStoreParams != nil {
		return true
	}

	return false
}

// SetDataStoreParams gets a reference to the given []DatastoreParams and assigns it to the DataStoreParams field.
func (o *EsxiRegistrationParams) SetDataStoreParams(v []DatastoreParams) {
	o.DataStoreParams = v
}

func (o EsxiRegistrationParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.MinFreeDatastoreSpaceForBackupGb.IsSet() {
		toSerialize["minFreeDatastoreSpaceForBackupGb"] = o.MinFreeDatastoreSpaceForBackupGb.Get()
	}
	if o.MaxConcurrentStreams.IsSet() {
		toSerialize["maxConcurrentStreams"] = o.MaxConcurrentStreams.Get()
	}
	if o.DataStoreParams != nil {
		toSerialize["dataStoreParams"] = o.DataStoreParams
	}
	return json.Marshal(toSerialize)
}

type NullableEsxiRegistrationParams struct {
	value *EsxiRegistrationParams
	isSet bool
}

func (v NullableEsxiRegistrationParams) Get() *EsxiRegistrationParams {
	return v.value
}

func (v *NullableEsxiRegistrationParams) Set(val *EsxiRegistrationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableEsxiRegistrationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableEsxiRegistrationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsxiRegistrationParams(val *EsxiRegistrationParams) *NullableEsxiRegistrationParams {
	return &NullableEsxiRegistrationParams{value: val, isSet: true}
}

func (v NullableEsxiRegistrationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsxiRegistrationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


