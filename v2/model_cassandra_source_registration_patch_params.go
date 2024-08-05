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

// checks if the CassandraSourceRegistrationPatchParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CassandraSourceRegistrationPatchParams{}

// CassandraSourceRegistrationPatchParams Specifies parameters to update cassandra source.
type CassandraSourceRegistrationPatchParams struct {
	CassandraCredentials NullableCassandraSourceRegistrationPatchParamsCassandraCredentials `json:"cassandraCredentials,omitempty"`
	// Commit Logs backup location on cassandra nodes
	CommitLogBackupLocation NullableString `json:"commitLogBackupLocation,omitempty"`
	// Directory path containing Cassandra configuration YAML file.
	ConfigDirectory NullableString `json:"configDirectory,omitempty"`
	// Data centers for this cluster.
	DataCenterNames []string `json:"dataCenterNames,omitempty"`
	// Directory from where DSE specific configuration can be read. This should be set only when you are using the DSE distribution of Cassandra.
	DseConfigurationDirectory NullableString `json:"dseConfigurationDirectory,omitempty"`
	DseSolrInfo *DSESolrInfo `json:"dseSolrInfo,omitempty"`
	// Set to true if this cluster has DSE Authenticator.
	IsDseAuthenticator NullableBool `json:"isDseAuthenticator,omitempty"`
	// Set to true if this cluster has DSE tiered storage.
	IsDseTieredStorage NullableBool `json:"isDseTieredStorage,omitempty"`
	JmxCredentials NullableCassandraSourceRegistrationPatchParamsJmxCredentials `json:"jmxCredentials,omitempty"`
	// Principal for the kerberos connection. (This is required only if your Cassandra has Kerberos authentication. Please refer to the user guide.)
	KerberosPrincipal NullableString `json:"kerberosPrincipal,omitempty"`
	// Any one seed node of the Cassandra cluster.
	SeedNode NullableString `json:"seedNode,omitempty"`
	SshPasswordCredentials NullableSshPasswordCredentials `json:"sshPasswordCredentials,omitempty"`
	SshPrivateKeyCredentials NullableSshPrivateKeyCredentials `json:"sshPrivateKeyCredentials,omitempty"`
}

// NewCassandraSourceRegistrationPatchParams instantiates a new CassandraSourceRegistrationPatchParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCassandraSourceRegistrationPatchParams() *CassandraSourceRegistrationPatchParams {
	this := CassandraSourceRegistrationPatchParams{}
	return &this
}

// NewCassandraSourceRegistrationPatchParamsWithDefaults instantiates a new CassandraSourceRegistrationPatchParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCassandraSourceRegistrationPatchParamsWithDefaults() *CassandraSourceRegistrationPatchParams {
	this := CassandraSourceRegistrationPatchParams{}
	return &this
}

// GetCassandraCredentials returns the CassandraCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceRegistrationPatchParams) GetCassandraCredentials() CassandraSourceRegistrationPatchParamsCassandraCredentials {
	if o == nil || IsNil(o.CassandraCredentials.Get()) {
		var ret CassandraSourceRegistrationPatchParamsCassandraCredentials
		return ret
	}
	return *o.CassandraCredentials.Get()
}

// GetCassandraCredentialsOk returns a tuple with the CassandraCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceRegistrationPatchParams) GetCassandraCredentialsOk() (*CassandraSourceRegistrationPatchParamsCassandraCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return o.CassandraCredentials.Get(), o.CassandraCredentials.IsSet()
}

// HasCassandraCredentials returns a boolean if a field has been set.
func (o *CassandraSourceRegistrationPatchParams) HasCassandraCredentials() bool {
	if o != nil && o.CassandraCredentials.IsSet() {
		return true
	}

	return false
}

// SetCassandraCredentials gets a reference to the given NullableCassandraSourceRegistrationPatchParamsCassandraCredentials and assigns it to the CassandraCredentials field.
func (o *CassandraSourceRegistrationPatchParams) SetCassandraCredentials(v CassandraSourceRegistrationPatchParamsCassandraCredentials) {
	o.CassandraCredentials.Set(&v)
}
// SetCassandraCredentialsNil sets the value for CassandraCredentials to be an explicit nil
func (o *CassandraSourceRegistrationPatchParams) SetCassandraCredentialsNil() {
	o.CassandraCredentials.Set(nil)
}

// UnsetCassandraCredentials ensures that no value is present for CassandraCredentials, not even an explicit nil
func (o *CassandraSourceRegistrationPatchParams) UnsetCassandraCredentials() {
	o.CassandraCredentials.Unset()
}

// GetCommitLogBackupLocation returns the CommitLogBackupLocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceRegistrationPatchParams) GetCommitLogBackupLocation() string {
	if o == nil || IsNil(o.CommitLogBackupLocation.Get()) {
		var ret string
		return ret
	}
	return *o.CommitLogBackupLocation.Get()
}

// GetCommitLogBackupLocationOk returns a tuple with the CommitLogBackupLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceRegistrationPatchParams) GetCommitLogBackupLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitLogBackupLocation.Get(), o.CommitLogBackupLocation.IsSet()
}

// HasCommitLogBackupLocation returns a boolean if a field has been set.
func (o *CassandraSourceRegistrationPatchParams) HasCommitLogBackupLocation() bool {
	if o != nil && o.CommitLogBackupLocation.IsSet() {
		return true
	}

	return false
}

// SetCommitLogBackupLocation gets a reference to the given NullableString and assigns it to the CommitLogBackupLocation field.
func (o *CassandraSourceRegistrationPatchParams) SetCommitLogBackupLocation(v string) {
	o.CommitLogBackupLocation.Set(&v)
}
// SetCommitLogBackupLocationNil sets the value for CommitLogBackupLocation to be an explicit nil
func (o *CassandraSourceRegistrationPatchParams) SetCommitLogBackupLocationNil() {
	o.CommitLogBackupLocation.Set(nil)
}

// UnsetCommitLogBackupLocation ensures that no value is present for CommitLogBackupLocation, not even an explicit nil
func (o *CassandraSourceRegistrationPatchParams) UnsetCommitLogBackupLocation() {
	o.CommitLogBackupLocation.Unset()
}

// GetConfigDirectory returns the ConfigDirectory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceRegistrationPatchParams) GetConfigDirectory() string {
	if o == nil || IsNil(o.ConfigDirectory.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigDirectory.Get()
}

// GetConfigDirectoryOk returns a tuple with the ConfigDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceRegistrationPatchParams) GetConfigDirectoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigDirectory.Get(), o.ConfigDirectory.IsSet()
}

// HasConfigDirectory returns a boolean if a field has been set.
func (o *CassandraSourceRegistrationPatchParams) HasConfigDirectory() bool {
	if o != nil && o.ConfigDirectory.IsSet() {
		return true
	}

	return false
}

// SetConfigDirectory gets a reference to the given NullableString and assigns it to the ConfigDirectory field.
func (o *CassandraSourceRegistrationPatchParams) SetConfigDirectory(v string) {
	o.ConfigDirectory.Set(&v)
}
// SetConfigDirectoryNil sets the value for ConfigDirectory to be an explicit nil
func (o *CassandraSourceRegistrationPatchParams) SetConfigDirectoryNil() {
	o.ConfigDirectory.Set(nil)
}

// UnsetConfigDirectory ensures that no value is present for ConfigDirectory, not even an explicit nil
func (o *CassandraSourceRegistrationPatchParams) UnsetConfigDirectory() {
	o.ConfigDirectory.Unset()
}

// GetDataCenterNames returns the DataCenterNames field value if set, zero value otherwise.
func (o *CassandraSourceRegistrationPatchParams) GetDataCenterNames() []string {
	if o == nil || IsNil(o.DataCenterNames) {
		var ret []string
		return ret
	}
	return o.DataCenterNames
}

// GetDataCenterNamesOk returns a tuple with the DataCenterNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CassandraSourceRegistrationPatchParams) GetDataCenterNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.DataCenterNames) {
		return nil, false
	}
	return o.DataCenterNames, true
}

// HasDataCenterNames returns a boolean if a field has been set.
func (o *CassandraSourceRegistrationPatchParams) HasDataCenterNames() bool {
	if o != nil && !IsNil(o.DataCenterNames) {
		return true
	}

	return false
}

// SetDataCenterNames gets a reference to the given []string and assigns it to the DataCenterNames field.
func (o *CassandraSourceRegistrationPatchParams) SetDataCenterNames(v []string) {
	o.DataCenterNames = v
}

// GetDseConfigurationDirectory returns the DseConfigurationDirectory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceRegistrationPatchParams) GetDseConfigurationDirectory() string {
	if o == nil || IsNil(o.DseConfigurationDirectory.Get()) {
		var ret string
		return ret
	}
	return *o.DseConfigurationDirectory.Get()
}

// GetDseConfigurationDirectoryOk returns a tuple with the DseConfigurationDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceRegistrationPatchParams) GetDseConfigurationDirectoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DseConfigurationDirectory.Get(), o.DseConfigurationDirectory.IsSet()
}

// HasDseConfigurationDirectory returns a boolean if a field has been set.
func (o *CassandraSourceRegistrationPatchParams) HasDseConfigurationDirectory() bool {
	if o != nil && o.DseConfigurationDirectory.IsSet() {
		return true
	}

	return false
}

// SetDseConfigurationDirectory gets a reference to the given NullableString and assigns it to the DseConfigurationDirectory field.
func (o *CassandraSourceRegistrationPatchParams) SetDseConfigurationDirectory(v string) {
	o.DseConfigurationDirectory.Set(&v)
}
// SetDseConfigurationDirectoryNil sets the value for DseConfigurationDirectory to be an explicit nil
func (o *CassandraSourceRegistrationPatchParams) SetDseConfigurationDirectoryNil() {
	o.DseConfigurationDirectory.Set(nil)
}

// UnsetDseConfigurationDirectory ensures that no value is present for DseConfigurationDirectory, not even an explicit nil
func (o *CassandraSourceRegistrationPatchParams) UnsetDseConfigurationDirectory() {
	o.DseConfigurationDirectory.Unset()
}

// GetDseSolrInfo returns the DseSolrInfo field value if set, zero value otherwise.
func (o *CassandraSourceRegistrationPatchParams) GetDseSolrInfo() DSESolrInfo {
	if o == nil || IsNil(o.DseSolrInfo) {
		var ret DSESolrInfo
		return ret
	}
	return *o.DseSolrInfo
}

// GetDseSolrInfoOk returns a tuple with the DseSolrInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CassandraSourceRegistrationPatchParams) GetDseSolrInfoOk() (*DSESolrInfo, bool) {
	if o == nil || IsNil(o.DseSolrInfo) {
		return nil, false
	}
	return o.DseSolrInfo, true
}

// HasDseSolrInfo returns a boolean if a field has been set.
func (o *CassandraSourceRegistrationPatchParams) HasDseSolrInfo() bool {
	if o != nil && !IsNil(o.DseSolrInfo) {
		return true
	}

	return false
}

// SetDseSolrInfo gets a reference to the given DSESolrInfo and assigns it to the DseSolrInfo field.
func (o *CassandraSourceRegistrationPatchParams) SetDseSolrInfo(v DSESolrInfo) {
	o.DseSolrInfo = &v
}

// GetIsDseAuthenticator returns the IsDseAuthenticator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceRegistrationPatchParams) GetIsDseAuthenticator() bool {
	if o == nil || IsNil(o.IsDseAuthenticator.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDseAuthenticator.Get()
}

// GetIsDseAuthenticatorOk returns a tuple with the IsDseAuthenticator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceRegistrationPatchParams) GetIsDseAuthenticatorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDseAuthenticator.Get(), o.IsDseAuthenticator.IsSet()
}

// HasIsDseAuthenticator returns a boolean if a field has been set.
func (o *CassandraSourceRegistrationPatchParams) HasIsDseAuthenticator() bool {
	if o != nil && o.IsDseAuthenticator.IsSet() {
		return true
	}

	return false
}

// SetIsDseAuthenticator gets a reference to the given NullableBool and assigns it to the IsDseAuthenticator field.
func (o *CassandraSourceRegistrationPatchParams) SetIsDseAuthenticator(v bool) {
	o.IsDseAuthenticator.Set(&v)
}
// SetIsDseAuthenticatorNil sets the value for IsDseAuthenticator to be an explicit nil
func (o *CassandraSourceRegistrationPatchParams) SetIsDseAuthenticatorNil() {
	o.IsDseAuthenticator.Set(nil)
}

// UnsetIsDseAuthenticator ensures that no value is present for IsDseAuthenticator, not even an explicit nil
func (o *CassandraSourceRegistrationPatchParams) UnsetIsDseAuthenticator() {
	o.IsDseAuthenticator.Unset()
}

// GetIsDseTieredStorage returns the IsDseTieredStorage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceRegistrationPatchParams) GetIsDseTieredStorage() bool {
	if o == nil || IsNil(o.IsDseTieredStorage.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDseTieredStorage.Get()
}

// GetIsDseTieredStorageOk returns a tuple with the IsDseTieredStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceRegistrationPatchParams) GetIsDseTieredStorageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDseTieredStorage.Get(), o.IsDseTieredStorage.IsSet()
}

// HasIsDseTieredStorage returns a boolean if a field has been set.
func (o *CassandraSourceRegistrationPatchParams) HasIsDseTieredStorage() bool {
	if o != nil && o.IsDseTieredStorage.IsSet() {
		return true
	}

	return false
}

// SetIsDseTieredStorage gets a reference to the given NullableBool and assigns it to the IsDseTieredStorage field.
func (o *CassandraSourceRegistrationPatchParams) SetIsDseTieredStorage(v bool) {
	o.IsDseTieredStorage.Set(&v)
}
// SetIsDseTieredStorageNil sets the value for IsDseTieredStorage to be an explicit nil
func (o *CassandraSourceRegistrationPatchParams) SetIsDseTieredStorageNil() {
	o.IsDseTieredStorage.Set(nil)
}

// UnsetIsDseTieredStorage ensures that no value is present for IsDseTieredStorage, not even an explicit nil
func (o *CassandraSourceRegistrationPatchParams) UnsetIsDseTieredStorage() {
	o.IsDseTieredStorage.Unset()
}

// GetJmxCredentials returns the JmxCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceRegistrationPatchParams) GetJmxCredentials() CassandraSourceRegistrationPatchParamsJmxCredentials {
	if o == nil || IsNil(o.JmxCredentials.Get()) {
		var ret CassandraSourceRegistrationPatchParamsJmxCredentials
		return ret
	}
	return *o.JmxCredentials.Get()
}

// GetJmxCredentialsOk returns a tuple with the JmxCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceRegistrationPatchParams) GetJmxCredentialsOk() (*CassandraSourceRegistrationPatchParamsJmxCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return o.JmxCredentials.Get(), o.JmxCredentials.IsSet()
}

// HasJmxCredentials returns a boolean if a field has been set.
func (o *CassandraSourceRegistrationPatchParams) HasJmxCredentials() bool {
	if o != nil && o.JmxCredentials.IsSet() {
		return true
	}

	return false
}

// SetJmxCredentials gets a reference to the given NullableCassandraSourceRegistrationPatchParamsJmxCredentials and assigns it to the JmxCredentials field.
func (o *CassandraSourceRegistrationPatchParams) SetJmxCredentials(v CassandraSourceRegistrationPatchParamsJmxCredentials) {
	o.JmxCredentials.Set(&v)
}
// SetJmxCredentialsNil sets the value for JmxCredentials to be an explicit nil
func (o *CassandraSourceRegistrationPatchParams) SetJmxCredentialsNil() {
	o.JmxCredentials.Set(nil)
}

// UnsetJmxCredentials ensures that no value is present for JmxCredentials, not even an explicit nil
func (o *CassandraSourceRegistrationPatchParams) UnsetJmxCredentials() {
	o.JmxCredentials.Unset()
}

// GetKerberosPrincipal returns the KerberosPrincipal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceRegistrationPatchParams) GetKerberosPrincipal() string {
	if o == nil || IsNil(o.KerberosPrincipal.Get()) {
		var ret string
		return ret
	}
	return *o.KerberosPrincipal.Get()
}

// GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceRegistrationPatchParams) GetKerberosPrincipalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KerberosPrincipal.Get(), o.KerberosPrincipal.IsSet()
}

// HasKerberosPrincipal returns a boolean if a field has been set.
func (o *CassandraSourceRegistrationPatchParams) HasKerberosPrincipal() bool {
	if o != nil && o.KerberosPrincipal.IsSet() {
		return true
	}

	return false
}

// SetKerberosPrincipal gets a reference to the given NullableString and assigns it to the KerberosPrincipal field.
func (o *CassandraSourceRegistrationPatchParams) SetKerberosPrincipal(v string) {
	o.KerberosPrincipal.Set(&v)
}
// SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil
func (o *CassandraSourceRegistrationPatchParams) SetKerberosPrincipalNil() {
	o.KerberosPrincipal.Set(nil)
}

// UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil
func (o *CassandraSourceRegistrationPatchParams) UnsetKerberosPrincipal() {
	o.KerberosPrincipal.Unset()
}

// GetSeedNode returns the SeedNode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceRegistrationPatchParams) GetSeedNode() string {
	if o == nil || IsNil(o.SeedNode.Get()) {
		var ret string
		return ret
	}
	return *o.SeedNode.Get()
}

// GetSeedNodeOk returns a tuple with the SeedNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceRegistrationPatchParams) GetSeedNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeedNode.Get(), o.SeedNode.IsSet()
}

// HasSeedNode returns a boolean if a field has been set.
func (o *CassandraSourceRegistrationPatchParams) HasSeedNode() bool {
	if o != nil && o.SeedNode.IsSet() {
		return true
	}

	return false
}

// SetSeedNode gets a reference to the given NullableString and assigns it to the SeedNode field.
func (o *CassandraSourceRegistrationPatchParams) SetSeedNode(v string) {
	o.SeedNode.Set(&v)
}
// SetSeedNodeNil sets the value for SeedNode to be an explicit nil
func (o *CassandraSourceRegistrationPatchParams) SetSeedNodeNil() {
	o.SeedNode.Set(nil)
}

// UnsetSeedNode ensures that no value is present for SeedNode, not even an explicit nil
func (o *CassandraSourceRegistrationPatchParams) UnsetSeedNode() {
	o.SeedNode.Unset()
}

// GetSshPasswordCredentials returns the SshPasswordCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceRegistrationPatchParams) GetSshPasswordCredentials() SshPasswordCredentials {
	if o == nil || IsNil(o.SshPasswordCredentials.Get()) {
		var ret SshPasswordCredentials
		return ret
	}
	return *o.SshPasswordCredentials.Get()
}

// GetSshPasswordCredentialsOk returns a tuple with the SshPasswordCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceRegistrationPatchParams) GetSshPasswordCredentialsOk() (*SshPasswordCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return o.SshPasswordCredentials.Get(), o.SshPasswordCredentials.IsSet()
}

// HasSshPasswordCredentials returns a boolean if a field has been set.
func (o *CassandraSourceRegistrationPatchParams) HasSshPasswordCredentials() bool {
	if o != nil && o.SshPasswordCredentials.IsSet() {
		return true
	}

	return false
}

// SetSshPasswordCredentials gets a reference to the given NullableSshPasswordCredentials and assigns it to the SshPasswordCredentials field.
func (o *CassandraSourceRegistrationPatchParams) SetSshPasswordCredentials(v SshPasswordCredentials) {
	o.SshPasswordCredentials.Set(&v)
}
// SetSshPasswordCredentialsNil sets the value for SshPasswordCredentials to be an explicit nil
func (o *CassandraSourceRegistrationPatchParams) SetSshPasswordCredentialsNil() {
	o.SshPasswordCredentials.Set(nil)
}

// UnsetSshPasswordCredentials ensures that no value is present for SshPasswordCredentials, not even an explicit nil
func (o *CassandraSourceRegistrationPatchParams) UnsetSshPasswordCredentials() {
	o.SshPasswordCredentials.Unset()
}

// GetSshPrivateKeyCredentials returns the SshPrivateKeyCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CassandraSourceRegistrationPatchParams) GetSshPrivateKeyCredentials() SshPrivateKeyCredentials {
	if o == nil || IsNil(o.SshPrivateKeyCredentials.Get()) {
		var ret SshPrivateKeyCredentials
		return ret
	}
	return *o.SshPrivateKeyCredentials.Get()
}

// GetSshPrivateKeyCredentialsOk returns a tuple with the SshPrivateKeyCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CassandraSourceRegistrationPatchParams) GetSshPrivateKeyCredentialsOk() (*SshPrivateKeyCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return o.SshPrivateKeyCredentials.Get(), o.SshPrivateKeyCredentials.IsSet()
}

// HasSshPrivateKeyCredentials returns a boolean if a field has been set.
func (o *CassandraSourceRegistrationPatchParams) HasSshPrivateKeyCredentials() bool {
	if o != nil && o.SshPrivateKeyCredentials.IsSet() {
		return true
	}

	return false
}

// SetSshPrivateKeyCredentials gets a reference to the given NullableSshPrivateKeyCredentials and assigns it to the SshPrivateKeyCredentials field.
func (o *CassandraSourceRegistrationPatchParams) SetSshPrivateKeyCredentials(v SshPrivateKeyCredentials) {
	o.SshPrivateKeyCredentials.Set(&v)
}
// SetSshPrivateKeyCredentialsNil sets the value for SshPrivateKeyCredentials to be an explicit nil
func (o *CassandraSourceRegistrationPatchParams) SetSshPrivateKeyCredentialsNil() {
	o.SshPrivateKeyCredentials.Set(nil)
}

// UnsetSshPrivateKeyCredentials ensures that no value is present for SshPrivateKeyCredentials, not even an explicit nil
func (o *CassandraSourceRegistrationPatchParams) UnsetSshPrivateKeyCredentials() {
	o.SshPrivateKeyCredentials.Unset()
}

func (o CassandraSourceRegistrationPatchParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CassandraSourceRegistrationPatchParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CassandraCredentials.IsSet() {
		toSerialize["cassandraCredentials"] = o.CassandraCredentials.Get()
	}
	if o.CommitLogBackupLocation.IsSet() {
		toSerialize["commitLogBackupLocation"] = o.CommitLogBackupLocation.Get()
	}
	if o.ConfigDirectory.IsSet() {
		toSerialize["configDirectory"] = o.ConfigDirectory.Get()
	}
	if !IsNil(o.DataCenterNames) {
		toSerialize["dataCenterNames"] = o.DataCenterNames
	}
	if o.DseConfigurationDirectory.IsSet() {
		toSerialize["dseConfigurationDirectory"] = o.DseConfigurationDirectory.Get()
	}
	if !IsNil(o.DseSolrInfo) {
		toSerialize["dseSolrInfo"] = o.DseSolrInfo
	}
	if o.IsDseAuthenticator.IsSet() {
		toSerialize["isDseAuthenticator"] = o.IsDseAuthenticator.Get()
	}
	if o.IsDseTieredStorage.IsSet() {
		toSerialize["isDseTieredStorage"] = o.IsDseTieredStorage.Get()
	}
	if o.JmxCredentials.IsSet() {
		toSerialize["jmxCredentials"] = o.JmxCredentials.Get()
	}
	if o.KerberosPrincipal.IsSet() {
		toSerialize["kerberosPrincipal"] = o.KerberosPrincipal.Get()
	}
	if o.SeedNode.IsSet() {
		toSerialize["seedNode"] = o.SeedNode.Get()
	}
	if o.SshPasswordCredentials.IsSet() {
		toSerialize["sshPasswordCredentials"] = o.SshPasswordCredentials.Get()
	}
	if o.SshPrivateKeyCredentials.IsSet() {
		toSerialize["sshPrivateKeyCredentials"] = o.SshPrivateKeyCredentials.Get()
	}
	return toSerialize, nil
}

type NullableCassandraSourceRegistrationPatchParams struct {
	value *CassandraSourceRegistrationPatchParams
	isSet bool
}

func (v NullableCassandraSourceRegistrationPatchParams) Get() *CassandraSourceRegistrationPatchParams {
	return v.value
}

func (v *NullableCassandraSourceRegistrationPatchParams) Set(val *CassandraSourceRegistrationPatchParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCassandraSourceRegistrationPatchParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCassandraSourceRegistrationPatchParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCassandraSourceRegistrationPatchParams(val *CassandraSourceRegistrationPatchParams) *NullableCassandraSourceRegistrationPatchParams {
	return &NullableCassandraSourceRegistrationPatchParams{value: val, isSet: true}
}

func (v NullableCassandraSourceRegistrationPatchParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCassandraSourceRegistrationPatchParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


