# SelfServiceSnapshotConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **NullableBool** | Specifies if self service snapshot feature is enabled. If this is set to true, the feature will also be enabled for NFS protocol. This field is deprecated. | [optional] 
**PreviousVersionsEnabled** | Pointer to **NullableBool** | Specifies if previouse versions feature is enabled with SMB protocol. | [optional] 
**NfsAccessEnabled** | Pointer to **NullableBool** | Specifies if self service snapshot feature is enabled for NFS protocol. | [optional] 
**SmbAccessEnabled** | Pointer to **NullableBool** | Specifies if self service snapshot feature is enabled for SMB protocol. | [optional] 
**SnapshotDirectoryName** | Pointer to **NullableString** | Specifies the directory name for the snapshots. | [optional] 
**AlternateSnapshotDirectoryName** | Pointer to **NullableString** | Specifies the alternate directory name for the snapshots. If it is not set, this feature for SMB protocol will not work. | [optional] 
**AllowAccessSids** | Pointer to **[]string** | Specifies a list of sids who has access to the snapshots. | [optional] 
**DenyAccessSids** | Pointer to **[]string** | Specifies a list of sids who does not have access to the snapshots. This field overrides &#39;allowAccessSids&#39;. | [optional] 

## Methods

### NewSelfServiceSnapshotConfig

`func NewSelfServiceSnapshotConfig() *SelfServiceSnapshotConfig`

NewSelfServiceSnapshotConfig instantiates a new SelfServiceSnapshotConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfServiceSnapshotConfigWithDefaults

`func NewSelfServiceSnapshotConfigWithDefaults() *SelfServiceSnapshotConfig`

NewSelfServiceSnapshotConfigWithDefaults instantiates a new SelfServiceSnapshotConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SelfServiceSnapshotConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SelfServiceSnapshotConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SelfServiceSnapshotConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SelfServiceSnapshotConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *SelfServiceSnapshotConfig) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *SelfServiceSnapshotConfig) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetPreviousVersionsEnabled

`func (o *SelfServiceSnapshotConfig) GetPreviousVersionsEnabled() bool`

GetPreviousVersionsEnabled returns the PreviousVersionsEnabled field if non-nil, zero value otherwise.

### GetPreviousVersionsEnabledOk

`func (o *SelfServiceSnapshotConfig) GetPreviousVersionsEnabledOk() (*bool, bool)`

GetPreviousVersionsEnabledOk returns a tuple with the PreviousVersionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVersionsEnabled

`func (o *SelfServiceSnapshotConfig) SetPreviousVersionsEnabled(v bool)`

SetPreviousVersionsEnabled sets PreviousVersionsEnabled field to given value.

### HasPreviousVersionsEnabled

`func (o *SelfServiceSnapshotConfig) HasPreviousVersionsEnabled() bool`

HasPreviousVersionsEnabled returns a boolean if a field has been set.

### SetPreviousVersionsEnabledNil

`func (o *SelfServiceSnapshotConfig) SetPreviousVersionsEnabledNil(b bool)`

 SetPreviousVersionsEnabledNil sets the value for PreviousVersionsEnabled to be an explicit nil

### UnsetPreviousVersionsEnabled
`func (o *SelfServiceSnapshotConfig) UnsetPreviousVersionsEnabled()`

UnsetPreviousVersionsEnabled ensures that no value is present for PreviousVersionsEnabled, not even an explicit nil
### GetNfsAccessEnabled

`func (o *SelfServiceSnapshotConfig) GetNfsAccessEnabled() bool`

GetNfsAccessEnabled returns the NfsAccessEnabled field if non-nil, zero value otherwise.

### GetNfsAccessEnabledOk

`func (o *SelfServiceSnapshotConfig) GetNfsAccessEnabledOk() (*bool, bool)`

GetNfsAccessEnabledOk returns a tuple with the NfsAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsAccessEnabled

`func (o *SelfServiceSnapshotConfig) SetNfsAccessEnabled(v bool)`

SetNfsAccessEnabled sets NfsAccessEnabled field to given value.

### HasNfsAccessEnabled

`func (o *SelfServiceSnapshotConfig) HasNfsAccessEnabled() bool`

HasNfsAccessEnabled returns a boolean if a field has been set.

### SetNfsAccessEnabledNil

`func (o *SelfServiceSnapshotConfig) SetNfsAccessEnabledNil(b bool)`

 SetNfsAccessEnabledNil sets the value for NfsAccessEnabled to be an explicit nil

### UnsetNfsAccessEnabled
`func (o *SelfServiceSnapshotConfig) UnsetNfsAccessEnabled()`

UnsetNfsAccessEnabled ensures that no value is present for NfsAccessEnabled, not even an explicit nil
### GetSmbAccessEnabled

`func (o *SelfServiceSnapshotConfig) GetSmbAccessEnabled() bool`

GetSmbAccessEnabled returns the SmbAccessEnabled field if non-nil, zero value otherwise.

### GetSmbAccessEnabledOk

`func (o *SelfServiceSnapshotConfig) GetSmbAccessEnabledOk() (*bool, bool)`

GetSmbAccessEnabledOk returns a tuple with the SmbAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbAccessEnabled

`func (o *SelfServiceSnapshotConfig) SetSmbAccessEnabled(v bool)`

SetSmbAccessEnabled sets SmbAccessEnabled field to given value.

### HasSmbAccessEnabled

`func (o *SelfServiceSnapshotConfig) HasSmbAccessEnabled() bool`

HasSmbAccessEnabled returns a boolean if a field has been set.

### SetSmbAccessEnabledNil

`func (o *SelfServiceSnapshotConfig) SetSmbAccessEnabledNil(b bool)`

 SetSmbAccessEnabledNil sets the value for SmbAccessEnabled to be an explicit nil

### UnsetSmbAccessEnabled
`func (o *SelfServiceSnapshotConfig) UnsetSmbAccessEnabled()`

UnsetSmbAccessEnabled ensures that no value is present for SmbAccessEnabled, not even an explicit nil
### GetSnapshotDirectoryName

`func (o *SelfServiceSnapshotConfig) GetSnapshotDirectoryName() string`

GetSnapshotDirectoryName returns the SnapshotDirectoryName field if non-nil, zero value otherwise.

### GetSnapshotDirectoryNameOk

`func (o *SelfServiceSnapshotConfig) GetSnapshotDirectoryNameOk() (*string, bool)`

GetSnapshotDirectoryNameOk returns a tuple with the SnapshotDirectoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotDirectoryName

`func (o *SelfServiceSnapshotConfig) SetSnapshotDirectoryName(v string)`

SetSnapshotDirectoryName sets SnapshotDirectoryName field to given value.

### HasSnapshotDirectoryName

`func (o *SelfServiceSnapshotConfig) HasSnapshotDirectoryName() bool`

HasSnapshotDirectoryName returns a boolean if a field has been set.

### SetSnapshotDirectoryNameNil

`func (o *SelfServiceSnapshotConfig) SetSnapshotDirectoryNameNil(b bool)`

 SetSnapshotDirectoryNameNil sets the value for SnapshotDirectoryName to be an explicit nil

### UnsetSnapshotDirectoryName
`func (o *SelfServiceSnapshotConfig) UnsetSnapshotDirectoryName()`

UnsetSnapshotDirectoryName ensures that no value is present for SnapshotDirectoryName, not even an explicit nil
### GetAlternateSnapshotDirectoryName

`func (o *SelfServiceSnapshotConfig) GetAlternateSnapshotDirectoryName() string`

GetAlternateSnapshotDirectoryName returns the AlternateSnapshotDirectoryName field if non-nil, zero value otherwise.

### GetAlternateSnapshotDirectoryNameOk

`func (o *SelfServiceSnapshotConfig) GetAlternateSnapshotDirectoryNameOk() (*string, bool)`

GetAlternateSnapshotDirectoryNameOk returns a tuple with the AlternateSnapshotDirectoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateSnapshotDirectoryName

`func (o *SelfServiceSnapshotConfig) SetAlternateSnapshotDirectoryName(v string)`

SetAlternateSnapshotDirectoryName sets AlternateSnapshotDirectoryName field to given value.

### HasAlternateSnapshotDirectoryName

`func (o *SelfServiceSnapshotConfig) HasAlternateSnapshotDirectoryName() bool`

HasAlternateSnapshotDirectoryName returns a boolean if a field has been set.

### SetAlternateSnapshotDirectoryNameNil

`func (o *SelfServiceSnapshotConfig) SetAlternateSnapshotDirectoryNameNil(b bool)`

 SetAlternateSnapshotDirectoryNameNil sets the value for AlternateSnapshotDirectoryName to be an explicit nil

### UnsetAlternateSnapshotDirectoryName
`func (o *SelfServiceSnapshotConfig) UnsetAlternateSnapshotDirectoryName()`

UnsetAlternateSnapshotDirectoryName ensures that no value is present for AlternateSnapshotDirectoryName, not even an explicit nil
### GetAllowAccessSids

`func (o *SelfServiceSnapshotConfig) GetAllowAccessSids() []string`

GetAllowAccessSids returns the AllowAccessSids field if non-nil, zero value otherwise.

### GetAllowAccessSidsOk

`func (o *SelfServiceSnapshotConfig) GetAllowAccessSidsOk() (*[]string, bool)`

GetAllowAccessSidsOk returns a tuple with the AllowAccessSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAccessSids

`func (o *SelfServiceSnapshotConfig) SetAllowAccessSids(v []string)`

SetAllowAccessSids sets AllowAccessSids field to given value.

### HasAllowAccessSids

`func (o *SelfServiceSnapshotConfig) HasAllowAccessSids() bool`

HasAllowAccessSids returns a boolean if a field has been set.

### GetDenyAccessSids

`func (o *SelfServiceSnapshotConfig) GetDenyAccessSids() []string`

GetDenyAccessSids returns the DenyAccessSids field if non-nil, zero value otherwise.

### GetDenyAccessSidsOk

`func (o *SelfServiceSnapshotConfig) GetDenyAccessSidsOk() (*[]string, bool)`

GetDenyAccessSidsOk returns a tuple with the DenyAccessSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyAccessSids

`func (o *SelfServiceSnapshotConfig) SetDenyAccessSids(v []string)`

SetDenyAccessSids sets DenyAccessSids field to given value.

### HasDenyAccessSids

`func (o *SelfServiceSnapshotConfig) HasDenyAccessSids() bool`

HasDenyAccessSids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


