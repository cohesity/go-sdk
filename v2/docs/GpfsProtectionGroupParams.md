# GpfsProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether or not the Protection Group should continue regardless of whether or not an error was encountered during protection group run. | [optional] 
**DirectCloudArchive** | Pointer to **NullableBool** | Specifies whether or not to store the snapshots in this run directly in an Archive Target instead of on the Cluster. If this is set to true, the associated policy must have exactly one Archive Target associated with it and the policy must be set up to archive after every run. Also, a Storage Domain cannot be specified. Default behavior is &#39;false&#39;. | [optional] 
**EncryptionEnabled** | Pointer to **NullableBool** | Specifies whether the protection group should use encryption while backup or not. | [optional] 
**FileFilters** | Pointer to [**FileFilteringPolicy**](FileFilteringPolicy.md) |  | [optional] 
**FileLockConfig** | Pointer to [**FileLevelDataLockConfig**](FileLevelDataLockConfig.md) |  | [optional] 
**FilterIpConfig** | Pointer to [**FilterIpConfig**](FilterIpConfig.md) |  | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**NativeFormat** | Pointer to **NullableBool** | Specifies whether or not to enable native format for direct archive job. This field is set to true if native format should be used for archiving. | [optional] 
**Objects** | [**[]GpfsProtectionGroupObjectParams**](GpfsProtectionGroupObjectParams.md) | Specifies the objects to be included in the Protection Group. | 
**PrePostScript** | Pointer to [**HostBasedBackupScriptParams**](HostBasedBackupScriptParams.md) |  | [optional] 
**Protocol** | Pointer to **NullableString** | Specifies the preferred protocol to use if this device supports multiple protocols. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 
**ThrottlingConfig** | Pointer to [**NasThrottlingConfig**](NasThrottlingConfig.md) |  | [optional] 

## Methods

### NewGpfsProtectionGroupParams

`func NewGpfsProtectionGroupParams(objects []GpfsProtectionGroupObjectParams, ) *GpfsProtectionGroupParams`

NewGpfsProtectionGroupParams instantiates a new GpfsProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpfsProtectionGroupParamsWithDefaults

`func NewGpfsProtectionGroupParamsWithDefaults() *GpfsProtectionGroupParams`

NewGpfsProtectionGroupParamsWithDefaults instantiates a new GpfsProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *GpfsProtectionGroupParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *GpfsProtectionGroupParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *GpfsProtectionGroupParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *GpfsProtectionGroupParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *GpfsProtectionGroupParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *GpfsProtectionGroupParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetDirectCloudArchive

`func (o *GpfsProtectionGroupParams) GetDirectCloudArchive() bool`

GetDirectCloudArchive returns the DirectCloudArchive field if non-nil, zero value otherwise.

### GetDirectCloudArchiveOk

`func (o *GpfsProtectionGroupParams) GetDirectCloudArchiveOk() (*bool, bool)`

GetDirectCloudArchiveOk returns a tuple with the DirectCloudArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectCloudArchive

`func (o *GpfsProtectionGroupParams) SetDirectCloudArchive(v bool)`

SetDirectCloudArchive sets DirectCloudArchive field to given value.

### HasDirectCloudArchive

`func (o *GpfsProtectionGroupParams) HasDirectCloudArchive() bool`

HasDirectCloudArchive returns a boolean if a field has been set.

### SetDirectCloudArchiveNil

`func (o *GpfsProtectionGroupParams) SetDirectCloudArchiveNil(b bool)`

 SetDirectCloudArchiveNil sets the value for DirectCloudArchive to be an explicit nil

### UnsetDirectCloudArchive
`func (o *GpfsProtectionGroupParams) UnsetDirectCloudArchive()`

UnsetDirectCloudArchive ensures that no value is present for DirectCloudArchive, not even an explicit nil
### GetEncryptionEnabled

`func (o *GpfsProtectionGroupParams) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *GpfsProtectionGroupParams) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *GpfsProtectionGroupParams) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *GpfsProtectionGroupParams) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *GpfsProtectionGroupParams) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *GpfsProtectionGroupParams) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetFileFilters

`func (o *GpfsProtectionGroupParams) GetFileFilters() FileFilteringPolicy`

GetFileFilters returns the FileFilters field if non-nil, zero value otherwise.

### GetFileFiltersOk

`func (o *GpfsProtectionGroupParams) GetFileFiltersOk() (*FileFilteringPolicy, bool)`

GetFileFiltersOk returns a tuple with the FileFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFilters

`func (o *GpfsProtectionGroupParams) SetFileFilters(v FileFilteringPolicy)`

SetFileFilters sets FileFilters field to given value.

### HasFileFilters

`func (o *GpfsProtectionGroupParams) HasFileFilters() bool`

HasFileFilters returns a boolean if a field has been set.

### GetFileLockConfig

`func (o *GpfsProtectionGroupParams) GetFileLockConfig() FileLevelDataLockConfig`

GetFileLockConfig returns the FileLockConfig field if non-nil, zero value otherwise.

### GetFileLockConfigOk

`func (o *GpfsProtectionGroupParams) GetFileLockConfigOk() (*FileLevelDataLockConfig, bool)`

GetFileLockConfigOk returns a tuple with the FileLockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLockConfig

`func (o *GpfsProtectionGroupParams) SetFileLockConfig(v FileLevelDataLockConfig)`

SetFileLockConfig sets FileLockConfig field to given value.

### HasFileLockConfig

`func (o *GpfsProtectionGroupParams) HasFileLockConfig() bool`

HasFileLockConfig returns a boolean if a field has been set.

### GetFilterIpConfig

`func (o *GpfsProtectionGroupParams) GetFilterIpConfig() FilterIpConfig`

GetFilterIpConfig returns the FilterIpConfig field if non-nil, zero value otherwise.

### GetFilterIpConfigOk

`func (o *GpfsProtectionGroupParams) GetFilterIpConfigOk() (*FilterIpConfig, bool)`

GetFilterIpConfigOk returns a tuple with the FilterIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterIpConfig

`func (o *GpfsProtectionGroupParams) SetFilterIpConfig(v FilterIpConfig)`

SetFilterIpConfig sets FilterIpConfig field to given value.

### HasFilterIpConfig

`func (o *GpfsProtectionGroupParams) HasFilterIpConfig() bool`

HasFilterIpConfig returns a boolean if a field has been set.

### GetIndexingPolicy

`func (o *GpfsProtectionGroupParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *GpfsProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *GpfsProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *GpfsProtectionGroupParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetNativeFormat

`func (o *GpfsProtectionGroupParams) GetNativeFormat() bool`

GetNativeFormat returns the NativeFormat field if non-nil, zero value otherwise.

### GetNativeFormatOk

`func (o *GpfsProtectionGroupParams) GetNativeFormatOk() (*bool, bool)`

GetNativeFormatOk returns a tuple with the NativeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeFormat

`func (o *GpfsProtectionGroupParams) SetNativeFormat(v bool)`

SetNativeFormat sets NativeFormat field to given value.

### HasNativeFormat

`func (o *GpfsProtectionGroupParams) HasNativeFormat() bool`

HasNativeFormat returns a boolean if a field has been set.

### SetNativeFormatNil

`func (o *GpfsProtectionGroupParams) SetNativeFormatNil(b bool)`

 SetNativeFormatNil sets the value for NativeFormat to be an explicit nil

### UnsetNativeFormat
`func (o *GpfsProtectionGroupParams) UnsetNativeFormat()`

UnsetNativeFormat ensures that no value is present for NativeFormat, not even an explicit nil
### GetObjects

`func (o *GpfsProtectionGroupParams) GetObjects() []GpfsProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *GpfsProtectionGroupParams) GetObjectsOk() (*[]GpfsProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *GpfsProtectionGroupParams) SetObjects(v []GpfsProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.


### GetPrePostScript

`func (o *GpfsProtectionGroupParams) GetPrePostScript() HostBasedBackupScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *GpfsProtectionGroupParams) GetPrePostScriptOk() (*HostBasedBackupScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *GpfsProtectionGroupParams) SetPrePostScript(v HostBasedBackupScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *GpfsProtectionGroupParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetProtocol

`func (o *GpfsProtectionGroupParams) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GpfsProtectionGroupParams) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GpfsProtectionGroupParams) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GpfsProtectionGroupParams) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *GpfsProtectionGroupParams) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *GpfsProtectionGroupParams) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetSourceId

`func (o *GpfsProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *GpfsProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *GpfsProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *GpfsProtectionGroupParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *GpfsProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *GpfsProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *GpfsProtectionGroupParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *GpfsProtectionGroupParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *GpfsProtectionGroupParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *GpfsProtectionGroupParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *GpfsProtectionGroupParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *GpfsProtectionGroupParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetThrottlingConfig

`func (o *GpfsProtectionGroupParams) GetThrottlingConfig() NasThrottlingConfig`

GetThrottlingConfig returns the ThrottlingConfig field if non-nil, zero value otherwise.

### GetThrottlingConfigOk

`func (o *GpfsProtectionGroupParams) GetThrottlingConfigOk() (*NasThrottlingConfig, bool)`

GetThrottlingConfigOk returns a tuple with the ThrottlingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingConfig

`func (o *GpfsProtectionGroupParams) SetThrottlingConfig(v NasThrottlingConfig)`

SetThrottlingConfig sets ThrottlingConfig field to given value.

### HasThrottlingConfig

`func (o *GpfsProtectionGroupParams) HasThrottlingConfig() bool`

HasThrottlingConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


