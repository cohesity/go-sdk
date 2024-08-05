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

// checks if the HbaseSourceRegistrationParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HbaseSourceRegistrationParams{}

// HbaseSourceRegistrationParams Specifies parameters to register an HBase source.
type HbaseSourceRegistrationParams struct {
	// Authentication type.
	AuthType NullableString `json:"authType,omitempty"`
	// The 'Data root directory' for this HBase.
	DataRootDirectory NullableString `json:"dataRootDirectory,omitempty"`
	// The 'Zookeeper Quorum' for this HBase.
	ZookeeperQuorum []string `json:"zookeeperQuorum,omitempty"`
	// The directory containing the hbase-site.xml.
	ConfigurationDirectory string `json:"configurationDirectory"`
	// Protection Source registration id of the HDFS on which this HBase is running.
	HdfsSourceRegistrationID int64 `json:"hdfsSourceRegistrationID"`
	// IP or hostname of any host from which the HBase configuration file hbase-site.xml can be read.
	Host string `json:"host"`
	// The kerberos principal to be used to connect to this Hbase source.
	KerberosPrincipal NullableString `json:"kerberosPrincipal,omitempty"`
	SshPasswordCredentials NullableHbaseSourceRegistrationParamsAllOfSshPasswordCredentials `json:"sshPasswordCredentials,omitempty"`
	SshPrivateKeyCredentials NullableHbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials `json:"sshPrivateKeyCredentials,omitempty"`
}

type _HbaseSourceRegistrationParams HbaseSourceRegistrationParams

// NewHbaseSourceRegistrationParams instantiates a new HbaseSourceRegistrationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHbaseSourceRegistrationParams(configurationDirectory string, hdfsSourceRegistrationID int64, host string) *HbaseSourceRegistrationParams {
	this := HbaseSourceRegistrationParams{}
	this.ConfigurationDirectory = configurationDirectory
	this.HdfsSourceRegistrationID = hdfsSourceRegistrationID
	this.Host = host
	return &this
}

// NewHbaseSourceRegistrationParamsWithDefaults instantiates a new HbaseSourceRegistrationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHbaseSourceRegistrationParamsWithDefaults() *HbaseSourceRegistrationParams {
	this := HbaseSourceRegistrationParams{}
	return &this
}

// GetAuthType returns the AuthType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HbaseSourceRegistrationParams) GetAuthType() string {
	if o == nil || IsNil(o.AuthType.Get()) {
		var ret string
		return ret
	}
	return *o.AuthType.Get()
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HbaseSourceRegistrationParams) GetAuthTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthType.Get(), o.AuthType.IsSet()
}

// HasAuthType returns a boolean if a field has been set.
func (o *HbaseSourceRegistrationParams) HasAuthType() bool {
	if o != nil && o.AuthType.IsSet() {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given NullableString and assigns it to the AuthType field.
func (o *HbaseSourceRegistrationParams) SetAuthType(v string) {
	o.AuthType.Set(&v)
}
// SetAuthTypeNil sets the value for AuthType to be an explicit nil
func (o *HbaseSourceRegistrationParams) SetAuthTypeNil() {
	o.AuthType.Set(nil)
}

// UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil
func (o *HbaseSourceRegistrationParams) UnsetAuthType() {
	o.AuthType.Unset()
}

// GetDataRootDirectory returns the DataRootDirectory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HbaseSourceRegistrationParams) GetDataRootDirectory() string {
	if o == nil || IsNil(o.DataRootDirectory.Get()) {
		var ret string
		return ret
	}
	return *o.DataRootDirectory.Get()
}

// GetDataRootDirectoryOk returns a tuple with the DataRootDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HbaseSourceRegistrationParams) GetDataRootDirectoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataRootDirectory.Get(), o.DataRootDirectory.IsSet()
}

// HasDataRootDirectory returns a boolean if a field has been set.
func (o *HbaseSourceRegistrationParams) HasDataRootDirectory() bool {
	if o != nil && o.DataRootDirectory.IsSet() {
		return true
	}

	return false
}

// SetDataRootDirectory gets a reference to the given NullableString and assigns it to the DataRootDirectory field.
func (o *HbaseSourceRegistrationParams) SetDataRootDirectory(v string) {
	o.DataRootDirectory.Set(&v)
}
// SetDataRootDirectoryNil sets the value for DataRootDirectory to be an explicit nil
func (o *HbaseSourceRegistrationParams) SetDataRootDirectoryNil() {
	o.DataRootDirectory.Set(nil)
}

// UnsetDataRootDirectory ensures that no value is present for DataRootDirectory, not even an explicit nil
func (o *HbaseSourceRegistrationParams) UnsetDataRootDirectory() {
	o.DataRootDirectory.Unset()
}

// GetZookeeperQuorum returns the ZookeeperQuorum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HbaseSourceRegistrationParams) GetZookeeperQuorum() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ZookeeperQuorum
}

// GetZookeeperQuorumOk returns a tuple with the ZookeeperQuorum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HbaseSourceRegistrationParams) GetZookeeperQuorumOk() ([]string, bool) {
	if o == nil || IsNil(o.ZookeeperQuorum) {
		return nil, false
	}
	return o.ZookeeperQuorum, true
}

// HasZookeeperQuorum returns a boolean if a field has been set.
func (o *HbaseSourceRegistrationParams) HasZookeeperQuorum() bool {
	if o != nil && !IsNil(o.ZookeeperQuorum) {
		return true
	}

	return false
}

// SetZookeeperQuorum gets a reference to the given []string and assigns it to the ZookeeperQuorum field.
func (o *HbaseSourceRegistrationParams) SetZookeeperQuorum(v []string) {
	o.ZookeeperQuorum = v
}

// GetConfigurationDirectory returns the ConfigurationDirectory field value
func (o *HbaseSourceRegistrationParams) GetConfigurationDirectory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigurationDirectory
}

// GetConfigurationDirectoryOk returns a tuple with the ConfigurationDirectory field value
// and a boolean to check if the value has been set.
func (o *HbaseSourceRegistrationParams) GetConfigurationDirectoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationDirectory, true
}

// SetConfigurationDirectory sets field value
func (o *HbaseSourceRegistrationParams) SetConfigurationDirectory(v string) {
	o.ConfigurationDirectory = v
}

// GetHdfsSourceRegistrationID returns the HdfsSourceRegistrationID field value
func (o *HbaseSourceRegistrationParams) GetHdfsSourceRegistrationID() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.HdfsSourceRegistrationID
}

// GetHdfsSourceRegistrationIDOk returns a tuple with the HdfsSourceRegistrationID field value
// and a boolean to check if the value has been set.
func (o *HbaseSourceRegistrationParams) GetHdfsSourceRegistrationIDOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HdfsSourceRegistrationID, true
}

// SetHdfsSourceRegistrationID sets field value
func (o *HbaseSourceRegistrationParams) SetHdfsSourceRegistrationID(v int64) {
	o.HdfsSourceRegistrationID = v
}

// GetHost returns the Host field value
func (o *HbaseSourceRegistrationParams) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *HbaseSourceRegistrationParams) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *HbaseSourceRegistrationParams) SetHost(v string) {
	o.Host = v
}

// GetKerberosPrincipal returns the KerberosPrincipal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HbaseSourceRegistrationParams) GetKerberosPrincipal() string {
	if o == nil || IsNil(o.KerberosPrincipal.Get()) {
		var ret string
		return ret
	}
	return *o.KerberosPrincipal.Get()
}

// GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HbaseSourceRegistrationParams) GetKerberosPrincipalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KerberosPrincipal.Get(), o.KerberosPrincipal.IsSet()
}

// HasKerberosPrincipal returns a boolean if a field has been set.
func (o *HbaseSourceRegistrationParams) HasKerberosPrincipal() bool {
	if o != nil && o.KerberosPrincipal.IsSet() {
		return true
	}

	return false
}

// SetKerberosPrincipal gets a reference to the given NullableString and assigns it to the KerberosPrincipal field.
func (o *HbaseSourceRegistrationParams) SetKerberosPrincipal(v string) {
	o.KerberosPrincipal.Set(&v)
}
// SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil
func (o *HbaseSourceRegistrationParams) SetKerberosPrincipalNil() {
	o.KerberosPrincipal.Set(nil)
}

// UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil
func (o *HbaseSourceRegistrationParams) UnsetKerberosPrincipal() {
	o.KerberosPrincipal.Unset()
}

// GetSshPasswordCredentials returns the SshPasswordCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HbaseSourceRegistrationParams) GetSshPasswordCredentials() HbaseSourceRegistrationParamsAllOfSshPasswordCredentials {
	if o == nil || IsNil(o.SshPasswordCredentials.Get()) {
		var ret HbaseSourceRegistrationParamsAllOfSshPasswordCredentials
		return ret
	}
	return *o.SshPasswordCredentials.Get()
}

// GetSshPasswordCredentialsOk returns a tuple with the SshPasswordCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HbaseSourceRegistrationParams) GetSshPasswordCredentialsOk() (*HbaseSourceRegistrationParamsAllOfSshPasswordCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return o.SshPasswordCredentials.Get(), o.SshPasswordCredentials.IsSet()
}

// HasSshPasswordCredentials returns a boolean if a field has been set.
func (o *HbaseSourceRegistrationParams) HasSshPasswordCredentials() bool {
	if o != nil && o.SshPasswordCredentials.IsSet() {
		return true
	}

	return false
}

// SetSshPasswordCredentials gets a reference to the given NullableHbaseSourceRegistrationParamsAllOfSshPasswordCredentials and assigns it to the SshPasswordCredentials field.
func (o *HbaseSourceRegistrationParams) SetSshPasswordCredentials(v HbaseSourceRegistrationParamsAllOfSshPasswordCredentials) {
	o.SshPasswordCredentials.Set(&v)
}
// SetSshPasswordCredentialsNil sets the value for SshPasswordCredentials to be an explicit nil
func (o *HbaseSourceRegistrationParams) SetSshPasswordCredentialsNil() {
	o.SshPasswordCredentials.Set(nil)
}

// UnsetSshPasswordCredentials ensures that no value is present for SshPasswordCredentials, not even an explicit nil
func (o *HbaseSourceRegistrationParams) UnsetSshPasswordCredentials() {
	o.SshPasswordCredentials.Unset()
}

// GetSshPrivateKeyCredentials returns the SshPrivateKeyCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HbaseSourceRegistrationParams) GetSshPrivateKeyCredentials() HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials {
	if o == nil || IsNil(o.SshPrivateKeyCredentials.Get()) {
		var ret HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials
		return ret
	}
	return *o.SshPrivateKeyCredentials.Get()
}

// GetSshPrivateKeyCredentialsOk returns a tuple with the SshPrivateKeyCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HbaseSourceRegistrationParams) GetSshPrivateKeyCredentialsOk() (*HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return o.SshPrivateKeyCredentials.Get(), o.SshPrivateKeyCredentials.IsSet()
}

// HasSshPrivateKeyCredentials returns a boolean if a field has been set.
func (o *HbaseSourceRegistrationParams) HasSshPrivateKeyCredentials() bool {
	if o != nil && o.SshPrivateKeyCredentials.IsSet() {
		return true
	}

	return false
}

// SetSshPrivateKeyCredentials gets a reference to the given NullableHbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials and assigns it to the SshPrivateKeyCredentials field.
func (o *HbaseSourceRegistrationParams) SetSshPrivateKeyCredentials(v HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials) {
	o.SshPrivateKeyCredentials.Set(&v)
}
// SetSshPrivateKeyCredentialsNil sets the value for SshPrivateKeyCredentials to be an explicit nil
func (o *HbaseSourceRegistrationParams) SetSshPrivateKeyCredentialsNil() {
	o.SshPrivateKeyCredentials.Set(nil)
}

// UnsetSshPrivateKeyCredentials ensures that no value is present for SshPrivateKeyCredentials, not even an explicit nil
func (o *HbaseSourceRegistrationParams) UnsetSshPrivateKeyCredentials() {
	o.SshPrivateKeyCredentials.Unset()
}

func (o HbaseSourceRegistrationParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HbaseSourceRegistrationParams) ToMap() (map[string]interface{}, error) {
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
	toSerialize["configurationDirectory"] = o.ConfigurationDirectory
	toSerialize["hdfsSourceRegistrationID"] = o.HdfsSourceRegistrationID
	toSerialize["host"] = o.Host
	if o.KerberosPrincipal.IsSet() {
		toSerialize["kerberosPrincipal"] = o.KerberosPrincipal.Get()
	}
	if o.SshPasswordCredentials.IsSet() {
		toSerialize["sshPasswordCredentials"] = o.SshPasswordCredentials.Get()
	}
	if o.SshPrivateKeyCredentials.IsSet() {
		toSerialize["sshPrivateKeyCredentials"] = o.SshPrivateKeyCredentials.Get()
	}
	return toSerialize, nil
}

func (o *HbaseSourceRegistrationParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"configurationDirectory",
		"hdfsSourceRegistrationID",
		"host",
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

	varHbaseSourceRegistrationParams := _HbaseSourceRegistrationParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHbaseSourceRegistrationParams)

	if err != nil {
		return err
	}

	*o = HbaseSourceRegistrationParams(varHbaseSourceRegistrationParams)

	return err
}

type NullableHbaseSourceRegistrationParams struct {
	value *HbaseSourceRegistrationParams
	isSet bool
}

func (v NullableHbaseSourceRegistrationParams) Get() *HbaseSourceRegistrationParams {
	return v.value
}

func (v *NullableHbaseSourceRegistrationParams) Set(val *HbaseSourceRegistrationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableHbaseSourceRegistrationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableHbaseSourceRegistrationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHbaseSourceRegistrationParams(val *HbaseSourceRegistrationParams) *NullableHbaseSourceRegistrationParams {
	return &NullableHbaseSourceRegistrationParams{value: val, isSet: true}
}

func (v NullableHbaseSourceRegistrationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHbaseSourceRegistrationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


