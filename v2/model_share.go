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

// checks if the Share type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Share{}

// Share Specifies the details of a Share.
type Share struct {
	// List of external client subnet IPs that are allowed to access the share.
	ClientSubnetWhitelist []Subnet `json:"clientSubnetWhitelist,omitempty"`
	// This field is currently deprecated. Specifies if Filer Audit Logging is enabled for this Share.
	EnableFilerAuditLogging NullableBool `json:"enableFilerAuditLogging,omitempty"`
	// Specifies the state of File Audit logging for this Share. Inherited: Audit log setting is inherited from the  View. Enabled: Audit log is enabled for this Share. Disabled: Audit log is disabled for this Share.
	FileAuditLoggingState NullableString `json:"fileAuditLoggingState,omitempty"`
	SmbConfig *UpdateShareParamSmbConfig `json:"smbConfig,omitempty"`
	// Specifies the Share name.
	Name NullableString `json:"name"`
	// Specifies the path for mounting this Share as an NFS share. If Kerberos Provider has multiple hostaliases, each host alias has its own path.
	NfsMountPaths []string `json:"nfsMountPaths,omitempty"`
	// Specifies the path to access this Share as an S3 share.
	S3AccessPath NullableString `json:"s3AccessPath,omitempty"`
	// Specifies the possible paths that can be used to mount this Share as a SMB share. If Active Directory has multiple account names, each machine account has its own path.
	SmbMountPaths []string `json:"smbMountPaths,omitempty"`
	// Specifies the tenant id who has access to this Share.
	TenantId NullableString `json:"tenantId,omitempty"`
	// Specifies the View name of this Share.
	ViewName NullableString `json:"viewName"`
	// Specifies the View path of this Share.
	ViewPath NullableString `json:"viewPath"`
}

type _Share Share

// NewShare instantiates a new Share object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShare(name NullableString, viewName NullableString, viewPath NullableString) *Share {
	this := Share{}
	this.Name = name
	this.ViewName = viewName
	this.ViewPath = viewPath
	return &this
}

// NewShareWithDefaults instantiates a new Share object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShareWithDefaults() *Share {
	this := Share{}
	return &this
}

// GetClientSubnetWhitelist returns the ClientSubnetWhitelist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Share) GetClientSubnetWhitelist() []Subnet {
	if o == nil {
		var ret []Subnet
		return ret
	}
	return o.ClientSubnetWhitelist
}

// GetClientSubnetWhitelistOk returns a tuple with the ClientSubnetWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Share) GetClientSubnetWhitelistOk() ([]Subnet, bool) {
	if o == nil || IsNil(o.ClientSubnetWhitelist) {
		return nil, false
	}
	return o.ClientSubnetWhitelist, true
}

// HasClientSubnetWhitelist returns a boolean if a field has been set.
func (o *Share) HasClientSubnetWhitelist() bool {
	if o != nil && !IsNil(o.ClientSubnetWhitelist) {
		return true
	}

	return false
}

// SetClientSubnetWhitelist gets a reference to the given []Subnet and assigns it to the ClientSubnetWhitelist field.
func (o *Share) SetClientSubnetWhitelist(v []Subnet) {
	o.ClientSubnetWhitelist = v
}

// GetEnableFilerAuditLogging returns the EnableFilerAuditLogging field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Share) GetEnableFilerAuditLogging() bool {
	if o == nil || IsNil(o.EnableFilerAuditLogging.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableFilerAuditLogging.Get()
}

// GetEnableFilerAuditLoggingOk returns a tuple with the EnableFilerAuditLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Share) GetEnableFilerAuditLoggingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableFilerAuditLogging.Get(), o.EnableFilerAuditLogging.IsSet()
}

// HasEnableFilerAuditLogging returns a boolean if a field has been set.
func (o *Share) HasEnableFilerAuditLogging() bool {
	if o != nil && o.EnableFilerAuditLogging.IsSet() {
		return true
	}

	return false
}

// SetEnableFilerAuditLogging gets a reference to the given NullableBool and assigns it to the EnableFilerAuditLogging field.
func (o *Share) SetEnableFilerAuditLogging(v bool) {
	o.EnableFilerAuditLogging.Set(&v)
}
// SetEnableFilerAuditLoggingNil sets the value for EnableFilerAuditLogging to be an explicit nil
func (o *Share) SetEnableFilerAuditLoggingNil() {
	o.EnableFilerAuditLogging.Set(nil)
}

// UnsetEnableFilerAuditLogging ensures that no value is present for EnableFilerAuditLogging, not even an explicit nil
func (o *Share) UnsetEnableFilerAuditLogging() {
	o.EnableFilerAuditLogging.Unset()
}

// GetFileAuditLoggingState returns the FileAuditLoggingState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Share) GetFileAuditLoggingState() string {
	if o == nil || IsNil(o.FileAuditLoggingState.Get()) {
		var ret string
		return ret
	}
	return *o.FileAuditLoggingState.Get()
}

// GetFileAuditLoggingStateOk returns a tuple with the FileAuditLoggingState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Share) GetFileAuditLoggingStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileAuditLoggingState.Get(), o.FileAuditLoggingState.IsSet()
}

// HasFileAuditLoggingState returns a boolean if a field has been set.
func (o *Share) HasFileAuditLoggingState() bool {
	if o != nil && o.FileAuditLoggingState.IsSet() {
		return true
	}

	return false
}

// SetFileAuditLoggingState gets a reference to the given NullableString and assigns it to the FileAuditLoggingState field.
func (o *Share) SetFileAuditLoggingState(v string) {
	o.FileAuditLoggingState.Set(&v)
}
// SetFileAuditLoggingStateNil sets the value for FileAuditLoggingState to be an explicit nil
func (o *Share) SetFileAuditLoggingStateNil() {
	o.FileAuditLoggingState.Set(nil)
}

// UnsetFileAuditLoggingState ensures that no value is present for FileAuditLoggingState, not even an explicit nil
func (o *Share) UnsetFileAuditLoggingState() {
	o.FileAuditLoggingState.Unset()
}

// GetSmbConfig returns the SmbConfig field value if set, zero value otherwise.
func (o *Share) GetSmbConfig() UpdateShareParamSmbConfig {
	if o == nil || IsNil(o.SmbConfig) {
		var ret UpdateShareParamSmbConfig
		return ret
	}
	return *o.SmbConfig
}

// GetSmbConfigOk returns a tuple with the SmbConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetSmbConfigOk() (*UpdateShareParamSmbConfig, bool) {
	if o == nil || IsNil(o.SmbConfig) {
		return nil, false
	}
	return o.SmbConfig, true
}

// HasSmbConfig returns a boolean if a field has been set.
func (o *Share) HasSmbConfig() bool {
	if o != nil && !IsNil(o.SmbConfig) {
		return true
	}

	return false
}

// SetSmbConfig gets a reference to the given UpdateShareParamSmbConfig and assigns it to the SmbConfig field.
func (o *Share) SetSmbConfig(v UpdateShareParamSmbConfig) {
	o.SmbConfig = &v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Share) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Share) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *Share) SetName(v string) {
	o.Name.Set(&v)
}

// GetNfsMountPaths returns the NfsMountPaths field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Share) GetNfsMountPaths() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NfsMountPaths
}

// GetNfsMountPathsOk returns a tuple with the NfsMountPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Share) GetNfsMountPathsOk() ([]string, bool) {
	if o == nil || IsNil(o.NfsMountPaths) {
		return nil, false
	}
	return o.NfsMountPaths, true
}

// HasNfsMountPaths returns a boolean if a field has been set.
func (o *Share) HasNfsMountPaths() bool {
	if o != nil && !IsNil(o.NfsMountPaths) {
		return true
	}

	return false
}

// SetNfsMountPaths gets a reference to the given []string and assigns it to the NfsMountPaths field.
func (o *Share) SetNfsMountPaths(v []string) {
	o.NfsMountPaths = v
}

// GetS3AccessPath returns the S3AccessPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Share) GetS3AccessPath() string {
	if o == nil || IsNil(o.S3AccessPath.Get()) {
		var ret string
		return ret
	}
	return *o.S3AccessPath.Get()
}

// GetS3AccessPathOk returns a tuple with the S3AccessPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Share) GetS3AccessPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.S3AccessPath.Get(), o.S3AccessPath.IsSet()
}

// HasS3AccessPath returns a boolean if a field has been set.
func (o *Share) HasS3AccessPath() bool {
	if o != nil && o.S3AccessPath.IsSet() {
		return true
	}

	return false
}

// SetS3AccessPath gets a reference to the given NullableString and assigns it to the S3AccessPath field.
func (o *Share) SetS3AccessPath(v string) {
	o.S3AccessPath.Set(&v)
}
// SetS3AccessPathNil sets the value for S3AccessPath to be an explicit nil
func (o *Share) SetS3AccessPathNil() {
	o.S3AccessPath.Set(nil)
}

// UnsetS3AccessPath ensures that no value is present for S3AccessPath, not even an explicit nil
func (o *Share) UnsetS3AccessPath() {
	o.S3AccessPath.Unset()
}

// GetSmbMountPaths returns the SmbMountPaths field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Share) GetSmbMountPaths() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SmbMountPaths
}

// GetSmbMountPathsOk returns a tuple with the SmbMountPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Share) GetSmbMountPathsOk() ([]string, bool) {
	if o == nil || IsNil(o.SmbMountPaths) {
		return nil, false
	}
	return o.SmbMountPaths, true
}

// HasSmbMountPaths returns a boolean if a field has been set.
func (o *Share) HasSmbMountPaths() bool {
	if o != nil && !IsNil(o.SmbMountPaths) {
		return true
	}

	return false
}

// SetSmbMountPaths gets a reference to the given []string and assigns it to the SmbMountPaths field.
func (o *Share) SetSmbMountPaths(v []string) {
	o.SmbMountPaths = v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Share) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Share) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *Share) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *Share) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *Share) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *Share) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetViewName returns the ViewName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Share) GetViewName() string {
	if o == nil || o.ViewName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ViewName.Get()
}

// GetViewNameOk returns a tuple with the ViewName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Share) GetViewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewName.Get(), o.ViewName.IsSet()
}

// SetViewName sets field value
func (o *Share) SetViewName(v string) {
	o.ViewName.Set(&v)
}

// GetViewPath returns the ViewPath field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Share) GetViewPath() string {
	if o == nil || o.ViewPath.Get() == nil {
		var ret string
		return ret
	}

	return *o.ViewPath.Get()
}

// GetViewPathOk returns a tuple with the ViewPath field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Share) GetViewPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewPath.Get(), o.ViewPath.IsSet()
}

// SetViewPath sets field value
func (o *Share) SetViewPath(v string) {
	o.ViewPath.Set(&v)
}

func (o Share) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Share) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientSubnetWhitelist != nil {
		toSerialize["clientSubnetWhitelist"] = o.ClientSubnetWhitelist
	}
	if o.EnableFilerAuditLogging.IsSet() {
		toSerialize["enableFilerAuditLogging"] = o.EnableFilerAuditLogging.Get()
	}
	if o.FileAuditLoggingState.IsSet() {
		toSerialize["fileAuditLoggingState"] = o.FileAuditLoggingState.Get()
	}
	if !IsNil(o.SmbConfig) {
		toSerialize["smbConfig"] = o.SmbConfig
	}
	toSerialize["name"] = o.Name.Get()
	if o.NfsMountPaths != nil {
		toSerialize["nfsMountPaths"] = o.NfsMountPaths
	}
	if o.S3AccessPath.IsSet() {
		toSerialize["s3AccessPath"] = o.S3AccessPath.Get()
	}
	if o.SmbMountPaths != nil {
		toSerialize["smbMountPaths"] = o.SmbMountPaths
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	toSerialize["viewName"] = o.ViewName.Get()
	toSerialize["viewPath"] = o.ViewPath.Get()
	return toSerialize, nil
}

func (o *Share) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"viewName",
		"viewPath",
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

	varShare := _Share{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShare)

	if err != nil {
		return err
	}

	*o = Share(varShare)

	return err
}

type NullableShare struct {
	value *Share
	isSet bool
}

func (v NullableShare) Get() *Share {
	return v.value
}

func (v *NullableShare) Set(val *Share) {
	v.value = val
	v.isSet = true
}

func (v NullableShare) IsSet() bool {
	return v.isSet
}

func (v *NullableShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShare(val *Share) *NullableShare {
	return &NullableShare{value: val, isSet: true}
}

func (v NullableShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


