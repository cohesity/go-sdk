# ArchivalNasExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **NullableString** | Specifies the host of the NAS external target. | 
**IsForeverIncrementalArchivalEnabled** | Pointer to **NullableBool** | Specifies if Forever Incremental Archival setting is enabled or not. | [optional] 
**IsIncrementalArchivalEnabled** | Pointer to **NullableBool** | Specifies if Incremental Archival setting is enabled or not. | [optional] 
**KerberosRealmName** | Pointer to **NullableString** | Specifies the Kerberos realm name for a Kerberos-secured target. | [optional] 
**MountPath** | **NullableString** | Specifies the mount path of the NAS external target. | 
**NfsSecurityType** | Pointer to **NullableString** | Specifies the NFS security type of the target. | [optional] 
**NfsVersionNumber** | Pointer to **NullableString** | Specifies the NFS version number of the target. | [optional] 
**ShareType** | Pointer to **NullableString** | Specifies the share type of the NAS external target. | [optional] [readonly] 
**SourceSideDeduplication** | Pointer to **NullableBool** | Specifies the Source Side Deduplication setting for the Nas external target | [optional] 

## Methods

### NewArchivalNasExternalTargetParams

`func NewArchivalNasExternalTargetParams(host NullableString, mountPath NullableString, ) *ArchivalNasExternalTargetParams`

NewArchivalNasExternalTargetParams instantiates a new ArchivalNasExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivalNasExternalTargetParamsWithDefaults

`func NewArchivalNasExternalTargetParamsWithDefaults() *ArchivalNasExternalTargetParams`

NewArchivalNasExternalTargetParamsWithDefaults instantiates a new ArchivalNasExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *ArchivalNasExternalTargetParams) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ArchivalNasExternalTargetParams) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ArchivalNasExternalTargetParams) SetHost(v string)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *ArchivalNasExternalTargetParams) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *ArchivalNasExternalTargetParams) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetIsForeverIncrementalArchivalEnabled

`func (o *ArchivalNasExternalTargetParams) GetIsForeverIncrementalArchivalEnabled() bool`

GetIsForeverIncrementalArchivalEnabled returns the IsForeverIncrementalArchivalEnabled field if non-nil, zero value otherwise.

### GetIsForeverIncrementalArchivalEnabledOk

`func (o *ArchivalNasExternalTargetParams) GetIsForeverIncrementalArchivalEnabledOk() (*bool, bool)`

GetIsForeverIncrementalArchivalEnabledOk returns a tuple with the IsForeverIncrementalArchivalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForeverIncrementalArchivalEnabled

`func (o *ArchivalNasExternalTargetParams) SetIsForeverIncrementalArchivalEnabled(v bool)`

SetIsForeverIncrementalArchivalEnabled sets IsForeverIncrementalArchivalEnabled field to given value.

### HasIsForeverIncrementalArchivalEnabled

`func (o *ArchivalNasExternalTargetParams) HasIsForeverIncrementalArchivalEnabled() bool`

HasIsForeverIncrementalArchivalEnabled returns a boolean if a field has been set.

### SetIsForeverIncrementalArchivalEnabledNil

`func (o *ArchivalNasExternalTargetParams) SetIsForeverIncrementalArchivalEnabledNil(b bool)`

 SetIsForeverIncrementalArchivalEnabledNil sets the value for IsForeverIncrementalArchivalEnabled to be an explicit nil

### UnsetIsForeverIncrementalArchivalEnabled
`func (o *ArchivalNasExternalTargetParams) UnsetIsForeverIncrementalArchivalEnabled()`

UnsetIsForeverIncrementalArchivalEnabled ensures that no value is present for IsForeverIncrementalArchivalEnabled, not even an explicit nil
### GetIsIncrementalArchivalEnabled

`func (o *ArchivalNasExternalTargetParams) GetIsIncrementalArchivalEnabled() bool`

GetIsIncrementalArchivalEnabled returns the IsIncrementalArchivalEnabled field if non-nil, zero value otherwise.

### GetIsIncrementalArchivalEnabledOk

`func (o *ArchivalNasExternalTargetParams) GetIsIncrementalArchivalEnabledOk() (*bool, bool)`

GetIsIncrementalArchivalEnabledOk returns a tuple with the IsIncrementalArchivalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncrementalArchivalEnabled

`func (o *ArchivalNasExternalTargetParams) SetIsIncrementalArchivalEnabled(v bool)`

SetIsIncrementalArchivalEnabled sets IsIncrementalArchivalEnabled field to given value.

### HasIsIncrementalArchivalEnabled

`func (o *ArchivalNasExternalTargetParams) HasIsIncrementalArchivalEnabled() bool`

HasIsIncrementalArchivalEnabled returns a boolean if a field has been set.

### SetIsIncrementalArchivalEnabledNil

`func (o *ArchivalNasExternalTargetParams) SetIsIncrementalArchivalEnabledNil(b bool)`

 SetIsIncrementalArchivalEnabledNil sets the value for IsIncrementalArchivalEnabled to be an explicit nil

### UnsetIsIncrementalArchivalEnabled
`func (o *ArchivalNasExternalTargetParams) UnsetIsIncrementalArchivalEnabled()`

UnsetIsIncrementalArchivalEnabled ensures that no value is present for IsIncrementalArchivalEnabled, not even an explicit nil
### GetKerberosRealmName

`func (o *ArchivalNasExternalTargetParams) GetKerberosRealmName() string`

GetKerberosRealmName returns the KerberosRealmName field if non-nil, zero value otherwise.

### GetKerberosRealmNameOk

`func (o *ArchivalNasExternalTargetParams) GetKerberosRealmNameOk() (*string, bool)`

GetKerberosRealmNameOk returns a tuple with the KerberosRealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosRealmName

`func (o *ArchivalNasExternalTargetParams) SetKerberosRealmName(v string)`

SetKerberosRealmName sets KerberosRealmName field to given value.

### HasKerberosRealmName

`func (o *ArchivalNasExternalTargetParams) HasKerberosRealmName() bool`

HasKerberosRealmName returns a boolean if a field has been set.

### SetKerberosRealmNameNil

`func (o *ArchivalNasExternalTargetParams) SetKerberosRealmNameNil(b bool)`

 SetKerberosRealmNameNil sets the value for KerberosRealmName to be an explicit nil

### UnsetKerberosRealmName
`func (o *ArchivalNasExternalTargetParams) UnsetKerberosRealmName()`

UnsetKerberosRealmName ensures that no value is present for KerberosRealmName, not even an explicit nil
### GetMountPath

`func (o *ArchivalNasExternalTargetParams) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *ArchivalNasExternalTargetParams) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *ArchivalNasExternalTargetParams) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.


### SetMountPathNil

`func (o *ArchivalNasExternalTargetParams) SetMountPathNil(b bool)`

 SetMountPathNil sets the value for MountPath to be an explicit nil

### UnsetMountPath
`func (o *ArchivalNasExternalTargetParams) UnsetMountPath()`

UnsetMountPath ensures that no value is present for MountPath, not even an explicit nil
### GetNfsSecurityType

`func (o *ArchivalNasExternalTargetParams) GetNfsSecurityType() string`

GetNfsSecurityType returns the NfsSecurityType field if non-nil, zero value otherwise.

### GetNfsSecurityTypeOk

`func (o *ArchivalNasExternalTargetParams) GetNfsSecurityTypeOk() (*string, bool)`

GetNfsSecurityTypeOk returns a tuple with the NfsSecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsSecurityType

`func (o *ArchivalNasExternalTargetParams) SetNfsSecurityType(v string)`

SetNfsSecurityType sets NfsSecurityType field to given value.

### HasNfsSecurityType

`func (o *ArchivalNasExternalTargetParams) HasNfsSecurityType() bool`

HasNfsSecurityType returns a boolean if a field has been set.

### SetNfsSecurityTypeNil

`func (o *ArchivalNasExternalTargetParams) SetNfsSecurityTypeNil(b bool)`

 SetNfsSecurityTypeNil sets the value for NfsSecurityType to be an explicit nil

### UnsetNfsSecurityType
`func (o *ArchivalNasExternalTargetParams) UnsetNfsSecurityType()`

UnsetNfsSecurityType ensures that no value is present for NfsSecurityType, not even an explicit nil
### GetNfsVersionNumber

`func (o *ArchivalNasExternalTargetParams) GetNfsVersionNumber() string`

GetNfsVersionNumber returns the NfsVersionNumber field if non-nil, zero value otherwise.

### GetNfsVersionNumberOk

`func (o *ArchivalNasExternalTargetParams) GetNfsVersionNumberOk() (*string, bool)`

GetNfsVersionNumberOk returns a tuple with the NfsVersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsVersionNumber

`func (o *ArchivalNasExternalTargetParams) SetNfsVersionNumber(v string)`

SetNfsVersionNumber sets NfsVersionNumber field to given value.

### HasNfsVersionNumber

`func (o *ArchivalNasExternalTargetParams) HasNfsVersionNumber() bool`

HasNfsVersionNumber returns a boolean if a field has been set.

### SetNfsVersionNumberNil

`func (o *ArchivalNasExternalTargetParams) SetNfsVersionNumberNil(b bool)`

 SetNfsVersionNumberNil sets the value for NfsVersionNumber to be an explicit nil

### UnsetNfsVersionNumber
`func (o *ArchivalNasExternalTargetParams) UnsetNfsVersionNumber()`

UnsetNfsVersionNumber ensures that no value is present for NfsVersionNumber, not even an explicit nil
### GetShareType

`func (o *ArchivalNasExternalTargetParams) GetShareType() string`

GetShareType returns the ShareType field if non-nil, zero value otherwise.

### GetShareTypeOk

`func (o *ArchivalNasExternalTargetParams) GetShareTypeOk() (*string, bool)`

GetShareTypeOk returns a tuple with the ShareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareType

`func (o *ArchivalNasExternalTargetParams) SetShareType(v string)`

SetShareType sets ShareType field to given value.

### HasShareType

`func (o *ArchivalNasExternalTargetParams) HasShareType() bool`

HasShareType returns a boolean if a field has been set.

### SetShareTypeNil

`func (o *ArchivalNasExternalTargetParams) SetShareTypeNil(b bool)`

 SetShareTypeNil sets the value for ShareType to be an explicit nil

### UnsetShareType
`func (o *ArchivalNasExternalTargetParams) UnsetShareType()`

UnsetShareType ensures that no value is present for ShareType, not even an explicit nil
### GetSourceSideDeduplication

`func (o *ArchivalNasExternalTargetParams) GetSourceSideDeduplication() bool`

GetSourceSideDeduplication returns the SourceSideDeduplication field if non-nil, zero value otherwise.

### GetSourceSideDeduplicationOk

`func (o *ArchivalNasExternalTargetParams) GetSourceSideDeduplicationOk() (*bool, bool)`

GetSourceSideDeduplicationOk returns a tuple with the SourceSideDeduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSideDeduplication

`func (o *ArchivalNasExternalTargetParams) SetSourceSideDeduplication(v bool)`

SetSourceSideDeduplication sets SourceSideDeduplication field to given value.

### HasSourceSideDeduplication

`func (o *ArchivalNasExternalTargetParams) HasSourceSideDeduplication() bool`

HasSourceSideDeduplication returns a boolean if a field has been set.

### SetSourceSideDeduplicationNil

`func (o *ArchivalNasExternalTargetParams) SetSourceSideDeduplicationNil(b bool)`

 SetSourceSideDeduplicationNil sets the value for SourceSideDeduplication to be an explicit nil

### UnsetSourceSideDeduplication
`func (o *ArchivalNasExternalTargetParams) UnsetSourceSideDeduplication()`

UnsetSourceSideDeduplication ensures that no value is present for SourceSideDeduplication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


