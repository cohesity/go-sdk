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

// checks if the CassandraSourceConfigParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CassandraSourceConfigParams{}

// CassandraSourceConfigParams Specifies the parameters fetched by reading cassandra configuration on the seed node.
type CassandraSourceConfigParams struct {
	// Cassandra partitioner required in compaction.
	CassandraPartitioner NullableString `json:"cassandraPartitioner,omitempty"`
	CassandraPortInfo *CassandraPortInfo `json:"cassandraPortInfo,omitempty"`
	CassandraSecurityInfo *CassandraSecurityInfo `json:"cassandraSecurityInfo,omitempty"`
	// Cassandra Version.
	CassandraVersion NullableString `json:"cassandraVersion,omitempty"`
	// Commit Logs backup location on cassandra nodes
	CommitLogBackupLocation NullableString `json:"commitLogBackupLocation,omitempty"`
	// Data centers for this cluster.
	DataCenterNames []string `json:"dataCenterNames,omitempty"`
	// DSE Version
	DseVersion NullableString `json:"dseVersion,omitempty"`
	// Endpoint snitch used for this cluster.
	EndpointSnitch NullableString `json:"endpointSnitch,omitempty"`
	// Is JMX Authentication enabled in this cluster ?
	IsJmxAuthEnable NullableBool `json:"isJmxAuthEnable,omitempty"`
	// Populated if cassandraAuthType is Kerberos.
	KerberosSaslProtocol NullableString `json:"kerberosSaslProtocol,omitempty"`
	// Seed nodes of this cluster.
	Seeds []string `json:"seeds,omitempty"`
}

// NewCassandraSourceConfigParams instantiates a new CassandraSourceConfigParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCassandraSourceConfigParams() *CassandraSourceConfigParams {
	this := CassandraSourceConfigParams{}
	return &this
}

// NewCassandraSourceConfigParamsWithDefaults instantiates a new CassandraSourceConfigParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCassandraSourceConfigParamsWithDefaults() *CassandraSourceConfigParams {
	this := CassandraSourceConfigParams{}
	return &this
}

// GetCassandraPartitioner returns the CassandraPartitioner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceConfigParams) GetCassandraPartitioner() string {
	if o == nil || IsNil(o.CassandraPartitioner.Get()) {
		var ret string
		return ret
	}
	return *o.CassandraPartitioner.Get()
}

// GetCassandraPartitionerOk returns a tuple with the CassandraPartitioner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceConfigParams) GetCassandraPartitionerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CassandraPartitioner.Get(), o.CassandraPartitioner.IsSet()
}

// HasCassandraPartitioner returns a boolean if a field has been set.
func (o *CassandraSourceConfigParams) HasCassandraPartitioner() bool {
	if o != nil && o.CassandraPartitioner.IsSet() {
		return true
	}

	return false
}

// SetCassandraPartitioner gets a reference to the given NullableString and assigns it to the CassandraPartitioner field.
func (o *CassandraSourceConfigParams) SetCassandraPartitioner(v string) {
	o.CassandraPartitioner.Set(&v)
}
// SetCassandraPartitionerNil sets the value for CassandraPartitioner to be an explicit nil
func (o *CassandraSourceConfigParams) SetCassandraPartitionerNil() {
	o.CassandraPartitioner.Set(nil)
}

// UnsetCassandraPartitioner ensures that no value is present for CassandraPartitioner, not even an explicit nil
func (o *CassandraSourceConfigParams) UnsetCassandraPartitioner() {
	o.CassandraPartitioner.Unset()
}

// GetCassandraPortInfo returns the CassandraPortInfo field value if set, zero value otherwise.
func (o *CassandraSourceConfigParams) GetCassandraPortInfo() CassandraPortInfo {
	if o == nil || IsNil(o.CassandraPortInfo) {
		var ret CassandraPortInfo
		return ret
	}
	return *o.CassandraPortInfo
}

// GetCassandraPortInfoOk returns a tuple with the CassandraPortInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CassandraSourceConfigParams) GetCassandraPortInfoOk() (*CassandraPortInfo, bool) {
	if o == nil || IsNil(o.CassandraPortInfo) {
		return nil, false
	}
	return o.CassandraPortInfo, true
}

// HasCassandraPortInfo returns a boolean if a field has been set.
func (o *CassandraSourceConfigParams) HasCassandraPortInfo() bool {
	if o != nil && !IsNil(o.CassandraPortInfo) {
		return true
	}

	return false
}

// SetCassandraPortInfo gets a reference to the given CassandraPortInfo and assigns it to the CassandraPortInfo field.
func (o *CassandraSourceConfigParams) SetCassandraPortInfo(v CassandraPortInfo) {
	o.CassandraPortInfo = &v
}

// GetCassandraSecurityInfo returns the CassandraSecurityInfo field value if set, zero value otherwise.
func (o *CassandraSourceConfigParams) GetCassandraSecurityInfo() CassandraSecurityInfo {
	if o == nil || IsNil(o.CassandraSecurityInfo) {
		var ret CassandraSecurityInfo
		return ret
	}
	return *o.CassandraSecurityInfo
}

// GetCassandraSecurityInfoOk returns a tuple with the CassandraSecurityInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CassandraSourceConfigParams) GetCassandraSecurityInfoOk() (*CassandraSecurityInfo, bool) {
	if o == nil || IsNil(o.CassandraSecurityInfo) {
		return nil, false
	}
	return o.CassandraSecurityInfo, true
}

// HasCassandraSecurityInfo returns a boolean if a field has been set.
func (o *CassandraSourceConfigParams) HasCassandraSecurityInfo() bool {
	if o != nil && !IsNil(o.CassandraSecurityInfo) {
		return true
	}

	return false
}

// SetCassandraSecurityInfo gets a reference to the given CassandraSecurityInfo and assigns it to the CassandraSecurityInfo field.
func (o *CassandraSourceConfigParams) SetCassandraSecurityInfo(v CassandraSecurityInfo) {
	o.CassandraSecurityInfo = &v
}

// GetCassandraVersion returns the CassandraVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceConfigParams) GetCassandraVersion() string {
	if o == nil || IsNil(o.CassandraVersion.Get()) {
		var ret string
		return ret
	}
	return *o.CassandraVersion.Get()
}

// GetCassandraVersionOk returns a tuple with the CassandraVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceConfigParams) GetCassandraVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CassandraVersion.Get(), o.CassandraVersion.IsSet()
}

// HasCassandraVersion returns a boolean if a field has been set.
func (o *CassandraSourceConfigParams) HasCassandraVersion() bool {
	if o != nil && o.CassandraVersion.IsSet() {
		return true
	}

	return false
}

// SetCassandraVersion gets a reference to the given NullableString and assigns it to the CassandraVersion field.
func (o *CassandraSourceConfigParams) SetCassandraVersion(v string) {
	o.CassandraVersion.Set(&v)
}
// SetCassandraVersionNil sets the value for CassandraVersion to be an explicit nil
func (o *CassandraSourceConfigParams) SetCassandraVersionNil() {
	o.CassandraVersion.Set(nil)
}

// UnsetCassandraVersion ensures that no value is present for CassandraVersion, not even an explicit nil
func (o *CassandraSourceConfigParams) UnsetCassandraVersion() {
	o.CassandraVersion.Unset()
}

// GetCommitLogBackupLocation returns the CommitLogBackupLocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceConfigParams) GetCommitLogBackupLocation() string {
	if o == nil || IsNil(o.CommitLogBackupLocation.Get()) {
		var ret string
		return ret
	}
	return *o.CommitLogBackupLocation.Get()
}

// GetCommitLogBackupLocationOk returns a tuple with the CommitLogBackupLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceConfigParams) GetCommitLogBackupLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitLogBackupLocation.Get(), o.CommitLogBackupLocation.IsSet()
}

// HasCommitLogBackupLocation returns a boolean if a field has been set.
func (o *CassandraSourceConfigParams) HasCommitLogBackupLocation() bool {
	if o != nil && o.CommitLogBackupLocation.IsSet() {
		return true
	}

	return false
}

// SetCommitLogBackupLocation gets a reference to the given NullableString and assigns it to the CommitLogBackupLocation field.
func (o *CassandraSourceConfigParams) SetCommitLogBackupLocation(v string) {
	o.CommitLogBackupLocation.Set(&v)
}
// SetCommitLogBackupLocationNil sets the value for CommitLogBackupLocation to be an explicit nil
func (o *CassandraSourceConfigParams) SetCommitLogBackupLocationNil() {
	o.CommitLogBackupLocation.Set(nil)
}

// UnsetCommitLogBackupLocation ensures that no value is present for CommitLogBackupLocation, not even an explicit nil
func (o *CassandraSourceConfigParams) UnsetCommitLogBackupLocation() {
	o.CommitLogBackupLocation.Unset()
}

// GetDataCenterNames returns the DataCenterNames field value if set, zero value otherwise.
func (o *CassandraSourceConfigParams) GetDataCenterNames() []string {
	if o == nil || IsNil(o.DataCenterNames) {
		var ret []string
		return ret
	}
	return o.DataCenterNames
}

// GetDataCenterNamesOk returns a tuple with the DataCenterNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CassandraSourceConfigParams) GetDataCenterNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.DataCenterNames) {
		return nil, false
	}
	return o.DataCenterNames, true
}

// HasDataCenterNames returns a boolean if a field has been set.
func (o *CassandraSourceConfigParams) HasDataCenterNames() bool {
	if o != nil && !IsNil(o.DataCenterNames) {
		return true
	}

	return false
}

// SetDataCenterNames gets a reference to the given []string and assigns it to the DataCenterNames field.
func (o *CassandraSourceConfigParams) SetDataCenterNames(v []string) {
	o.DataCenterNames = v
}

// GetDseVersion returns the DseVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceConfigParams) GetDseVersion() string {
	if o == nil || IsNil(o.DseVersion.Get()) {
		var ret string
		return ret
	}
	return *o.DseVersion.Get()
}

// GetDseVersionOk returns a tuple with the DseVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceConfigParams) GetDseVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DseVersion.Get(), o.DseVersion.IsSet()
}

// HasDseVersion returns a boolean if a field has been set.
func (o *CassandraSourceConfigParams) HasDseVersion() bool {
	if o != nil && o.DseVersion.IsSet() {
		return true
	}

	return false
}

// SetDseVersion gets a reference to the given NullableString and assigns it to the DseVersion field.
func (o *CassandraSourceConfigParams) SetDseVersion(v string) {
	o.DseVersion.Set(&v)
}
// SetDseVersionNil sets the value for DseVersion to be an explicit nil
func (o *CassandraSourceConfigParams) SetDseVersionNil() {
	o.DseVersion.Set(nil)
}

// UnsetDseVersion ensures that no value is present for DseVersion, not even an explicit nil
func (o *CassandraSourceConfigParams) UnsetDseVersion() {
	o.DseVersion.Unset()
}

// GetEndpointSnitch returns the EndpointSnitch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceConfigParams) GetEndpointSnitch() string {
	if o == nil || IsNil(o.EndpointSnitch.Get()) {
		var ret string
		return ret
	}
	return *o.EndpointSnitch.Get()
}

// GetEndpointSnitchOk returns a tuple with the EndpointSnitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceConfigParams) GetEndpointSnitchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndpointSnitch.Get(), o.EndpointSnitch.IsSet()
}

// HasEndpointSnitch returns a boolean if a field has been set.
func (o *CassandraSourceConfigParams) HasEndpointSnitch() bool {
	if o != nil && o.EndpointSnitch.IsSet() {
		return true
	}

	return false
}

// SetEndpointSnitch gets a reference to the given NullableString and assigns it to the EndpointSnitch field.
func (o *CassandraSourceConfigParams) SetEndpointSnitch(v string) {
	o.EndpointSnitch.Set(&v)
}
// SetEndpointSnitchNil sets the value for EndpointSnitch to be an explicit nil
func (o *CassandraSourceConfigParams) SetEndpointSnitchNil() {
	o.EndpointSnitch.Set(nil)
}

// UnsetEndpointSnitch ensures that no value is present for EndpointSnitch, not even an explicit nil
func (o *CassandraSourceConfigParams) UnsetEndpointSnitch() {
	o.EndpointSnitch.Unset()
}

// GetIsJmxAuthEnable returns the IsJmxAuthEnable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceConfigParams) GetIsJmxAuthEnable() bool {
	if o == nil || IsNil(o.IsJmxAuthEnable.Get()) {
		var ret bool
		return ret
	}
	return *o.IsJmxAuthEnable.Get()
}

// GetIsJmxAuthEnableOk returns a tuple with the IsJmxAuthEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceConfigParams) GetIsJmxAuthEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsJmxAuthEnable.Get(), o.IsJmxAuthEnable.IsSet()
}

// HasIsJmxAuthEnable returns a boolean if a field has been set.
func (o *CassandraSourceConfigParams) HasIsJmxAuthEnable() bool {
	if o != nil && o.IsJmxAuthEnable.IsSet() {
		return true
	}

	return false
}

// SetIsJmxAuthEnable gets a reference to the given NullableBool and assigns it to the IsJmxAuthEnable field.
func (o *CassandraSourceConfigParams) SetIsJmxAuthEnable(v bool) {
	o.IsJmxAuthEnable.Set(&v)
}
// SetIsJmxAuthEnableNil sets the value for IsJmxAuthEnable to be an explicit nil
func (o *CassandraSourceConfigParams) SetIsJmxAuthEnableNil() {
	o.IsJmxAuthEnable.Set(nil)
}

// UnsetIsJmxAuthEnable ensures that no value is present for IsJmxAuthEnable, not even an explicit nil
func (o *CassandraSourceConfigParams) UnsetIsJmxAuthEnable() {
	o.IsJmxAuthEnable.Unset()
}

// GetKerberosSaslProtocol returns the KerberosSaslProtocol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceConfigParams) GetKerberosSaslProtocol() string {
	if o == nil || IsNil(o.KerberosSaslProtocol.Get()) {
		var ret string
		return ret
	}
	return *o.KerberosSaslProtocol.Get()
}

// GetKerberosSaslProtocolOk returns a tuple with the KerberosSaslProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceConfigParams) GetKerberosSaslProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KerberosSaslProtocol.Get(), o.KerberosSaslProtocol.IsSet()
}

// HasKerberosSaslProtocol returns a boolean if a field has been set.
func (o *CassandraSourceConfigParams) HasKerberosSaslProtocol() bool {
	if o != nil && o.KerberosSaslProtocol.IsSet() {
		return true
	}

	return false
}

// SetKerberosSaslProtocol gets a reference to the given NullableString and assigns it to the KerberosSaslProtocol field.
func (o *CassandraSourceConfigParams) SetKerberosSaslProtocol(v string) {
	o.KerberosSaslProtocol.Set(&v)
}
// SetKerberosSaslProtocolNil sets the value for KerberosSaslProtocol to be an explicit nil
func (o *CassandraSourceConfigParams) SetKerberosSaslProtocolNil() {
	o.KerberosSaslProtocol.Set(nil)
}

// UnsetKerberosSaslProtocol ensures that no value is present for KerberosSaslProtocol, not even an explicit nil
func (o *CassandraSourceConfigParams) UnsetKerberosSaslProtocol() {
	o.KerberosSaslProtocol.Unset()
}

// GetSeeds returns the Seeds field value if set, zero value otherwise.
func (o *CassandraSourceConfigParams) GetSeeds() []string {
	if o == nil || IsNil(o.Seeds) {
		var ret []string
		return ret
	}
	return o.Seeds
}

// GetSeedsOk returns a tuple with the Seeds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CassandraSourceConfigParams) GetSeedsOk() ([]string, bool) {
	if o == nil || IsNil(o.Seeds) {
		return nil, false
	}
	return o.Seeds, true
}

// HasSeeds returns a boolean if a field has been set.
func (o *CassandraSourceConfigParams) HasSeeds() bool {
	if o != nil && !IsNil(o.Seeds) {
		return true
	}

	return false
}

// SetSeeds gets a reference to the given []string and assigns it to the Seeds field.
func (o *CassandraSourceConfigParams) SetSeeds(v []string) {
	o.Seeds = v
}

func (o CassandraSourceConfigParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CassandraSourceConfigParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CassandraPartitioner.IsSet() {
		toSerialize["cassandraPartitioner"] = o.CassandraPartitioner.Get()
	}
	if !IsNil(o.CassandraPortInfo) {
		toSerialize["cassandraPortInfo"] = o.CassandraPortInfo
	}
	if !IsNil(o.CassandraSecurityInfo) {
		toSerialize["cassandraSecurityInfo"] = o.CassandraSecurityInfo
	}
	if o.CassandraVersion.IsSet() {
		toSerialize["cassandraVersion"] = o.CassandraVersion.Get()
	}
	if o.CommitLogBackupLocation.IsSet() {
		toSerialize["commitLogBackupLocation"] = o.CommitLogBackupLocation.Get()
	}
	if !IsNil(o.DataCenterNames) {
		toSerialize["dataCenterNames"] = o.DataCenterNames
	}
	if o.DseVersion.IsSet() {
		toSerialize["dseVersion"] = o.DseVersion.Get()
	}
	if o.EndpointSnitch.IsSet() {
		toSerialize["endpointSnitch"] = o.EndpointSnitch.Get()
	}
	if o.IsJmxAuthEnable.IsSet() {
		toSerialize["isJmxAuthEnable"] = o.IsJmxAuthEnable.Get()
	}
	if o.KerberosSaslProtocol.IsSet() {
		toSerialize["kerberosSaslProtocol"] = o.KerberosSaslProtocol.Get()
	}
	if !IsNil(o.Seeds) {
		toSerialize["seeds"] = o.Seeds
	}
	return toSerialize, nil
}

type NullableCassandraSourceConfigParams struct {
	value *CassandraSourceConfigParams
	isSet bool
}

func (v NullableCassandraSourceConfigParams) Get() *CassandraSourceConfigParams {
	return v.value
}

func (v *NullableCassandraSourceConfigParams) Set(val *CassandraSourceConfigParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCassandraSourceConfigParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCassandraSourceConfigParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCassandraSourceConfigParams(val *CassandraSourceConfigParams) *NullableCassandraSourceConfigParams {
	return &NullableCassandraSourceConfigParams{value: val, isSet: true}
}

func (v NullableCassandraSourceConfigParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCassandraSourceConfigParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


