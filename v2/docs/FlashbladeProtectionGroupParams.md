# FlashbladeProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether or not the Protection Group should continue regardless of whether or not an error was encountered. | [optional] 
**DirectCloudArchive** | Pointer to **NullableBool** | Specifies whether or not to store the snapshots in this run directly in an Archive Target instead of on the Cluster. If this is set to true, the associated policy must have exactly one Archive Target associated with it and the policy must be set up to archive after every run. Also, a Storage Domain cannot be specified. Default behavior is &#39;false&#39;. | [optional] 
**EncryptionEnabled** | Pointer to **NullableBool** | Specifies whether the protection group should use encryption while backup or not. | [optional] 
**FileFilters** | Pointer to [**FileFilteringPolicy**](FileFilteringPolicy.md) |  | [optional] 
**FileLockConfig** | Pointer to [**FileLevelDataLockConfig**](FileLevelDataLockConfig.md) |  | [optional] 
**FilterIpConfig** | Pointer to [**FilterIpConfig**](FilterIpConfig.md) |  | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**NativeFormat** | Pointer to **NullableBool** | Specifies whether or not to enable native format for direct archive job. This field is set to true if native format should be used for archiving. | [optional] 
**Objects** | [**[]FlashbladeProtectionGroupObjectParams**](FlashbladeProtectionGroupObjectParams.md) | Specifies the objects to be included in the Protection Group. | 
**PrePostScript** | Pointer to [**HostBasedBackupScriptParams**](HostBasedBackupScriptParams.md) |  | [optional] 
**Protocol** | Pointer to **NullableString** | Specifies the preferred protocol to use if this device supports multiple protocols. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 
**ThrottlingConfig** | Pointer to [**NasThrottlingConfig**](NasThrottlingConfig.md) |  | [optional] 

## Methods

### NewFlashbladeProtectionGroupParams

`func NewFlashbladeProtectionGroupParams(objects []FlashbladeProtectionGroupObjectParams, ) *FlashbladeProtectionGroupParams`

NewFlashbladeProtectionGroupParams instantiates a new FlashbladeProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlashbladeProtectionGroupParamsWithDefaults

`func NewFlashbladeProtectionGroupParamsWithDefaults() *FlashbladeProtectionGroupParams`

NewFlashbladeProtectionGroupParamsWithDefaults instantiates a new FlashbladeProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *FlashbladeProtectionGroupParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *FlashbladeProtectionGroupParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *FlashbladeProtectionGroupParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *FlashbladeProtectionGroupParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *FlashbladeProtectionGroupParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *FlashbladeProtectionGroupParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetDirectCloudArchive

`func (o *FlashbladeProtectionGroupParams) GetDirectCloudArchive() bool`

GetDirectCloudArchive returns the DirectCloudArchive field if non-nil, zero value otherwise.

### GetDirectCloudArchiveOk

`func (o *FlashbladeProtectionGroupParams) GetDirectCloudArchiveOk() (*bool, bool)`

GetDirectCloudArchiveOk returns a tuple with the DirectCloudArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectCloudArchive

`func (o *FlashbladeProtectionGroupParams) SetDirectCloudArchive(v bool)`

SetDirectCloudArchive sets DirectCloudArchive field to given value.

### HasDirectCloudArchive

`func (o *FlashbladeProtectionGroupParams) HasDirectCloudArchive() bool`

HasDirectCloudArchive returns a boolean if a field has been set.

### SetDirectCloudArchiveNil

`func (o *FlashbladeProtectionGroupParams) SetDirectCloudArchiveNil(b bool)`

 SetDirectCloudArchiveNil sets the value for DirectCloudArchive to be an explicit nil

### UnsetDirectCloudArchive
`func (o *FlashbladeProtectionGroupParams) UnsetDirectCloudArchive()`

UnsetDirectCloudArchive ensures that no value is present for DirectCloudArchive, not even an explicit nil
### GetEncryptionEnabled

`func (o *FlashbladeProtectionGroupParams) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *FlashbladeProtectionGroupParams) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *FlashbladeProtectionGroupParams) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *FlashbladeProtectionGroupParams) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *FlashbladeProtectionGroupParams) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *FlashbladeProtectionGroupParams) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetFileFilters

`func (o *FlashbladeProtectionGroupParams) GetFileFilters() FileFilteringPolicy`

GetFileFilters returns the FileFilters field if non-nil, zero value otherwise.

### GetFileFiltersOk

`func (o *FlashbladeProtectionGroupParams) GetFileFiltersOk() (*FileFilteringPolicy, bool)`

GetFileFiltersOk returns a tuple with the FileFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFilters

`func (o *FlashbladeProtectionGroupParams) SetFileFilters(v FileFilteringPolicy)`

SetFileFilters sets FileFilters field to given value.

### HasFileFilters

`func (o *FlashbladeProtectionGroupParams) HasFileFilters() bool`

HasFileFilters returns a boolean if a field has been set.

### GetFileLockConfig

`func (o *FlashbladeProtectionGroupParams) GetFileLockConfig() FileLevelDataLockConfig`

GetFileLockConfig returns the FileLockConfig field if non-nil, zero value otherwise.

### GetFileLockConfigOk

`func (o *FlashbladeProtectionGroupParams) GetFileLockConfigOk() (*FileLevelDataLockConfig, bool)`

GetFileLockConfigOk returns a tuple with the FileLockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLockConfig

`func (o *FlashbladeProtectionGroupParams) SetFileLockConfig(v FileLevelDataLockConfig)`

SetFileLockConfig sets FileLockConfig field to given value.

### HasFileLockConfig

`func (o *FlashbladeProtectionGroupParams) HasFileLockConfig() bool`

HasFileLockConfig returns a boolean if a field has been set.

### GetFilterIpConfig

`func (o *FlashbladeProtectionGroupParams) GetFilterIpConfig() FilterIpConfig`

GetFilterIpConfig returns the FilterIpConfig field if non-nil, zero value otherwise.

### GetFilterIpConfigOk

`func (o *FlashbladeProtectionGroupParams) GetFilterIpConfigOk() (*FilterIpConfig, bool)`

GetFilterIpConfigOk returns a tuple with the FilterIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterIpConfig

`func (o *FlashbladeProtectionGroupParams) SetFilterIpConfig(v FilterIpConfig)`

SetFilterIpConfig sets FilterIpConfig field to given value.

### HasFilterIpConfig

`func (o *FlashbladeProtectionGroupParams) HasFilterIpConfig() bool`

HasFilterIpConfig returns a boolean if a field has been set.

### GetIndexingPolicy

`func (o *FlashbladeProtectionGroupParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *FlashbladeProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *FlashbladeProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *FlashbladeProtectionGroupParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetNativeFormat

`func (o *FlashbladeProtectionGroupParams) GetNativeFormat() bool`

GetNativeFormat returns the NativeFormat field if non-nil, zero value otherwise.

### GetNativeFormatOk

`func (o *FlashbladeProtectionGroupParams) GetNativeFormatOk() (*bool, bool)`

GetNativeFormatOk returns a tuple with the NativeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeFormat

`func (o *FlashbladeProtectionGroupParams) SetNativeFormat(v bool)`

SetNativeFormat sets NativeFormat field to given value.

### HasNativeFormat

`func (o *FlashbladeProtectionGroupParams) HasNativeFormat() bool`

HasNativeFormat returns a boolean if a field has been set.

### SetNativeFormatNil

`func (o *FlashbladeProtectionGroupParams) SetNativeFormatNil(b bool)`

 SetNativeFormatNil sets the value for NativeFormat to be an explicit nil

### UnsetNativeFormat
`func (o *FlashbladeProtectionGroupParams) UnsetNativeFormat()`

UnsetNativeFormat ensures that no value is present for NativeFormat, not even an explicit nil
### GetObjects

`func (o *FlashbladeProtectionGroupParams) GetObjects() []FlashbladeProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *FlashbladeProtectionGroupParams) GetObjectsOk() (*[]FlashbladeProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *FlashbladeProtectionGroupParams) SetObjects(v []FlashbladeProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.


### GetPrePostScript

`func (o *FlashbladeProtectionGroupParams) GetPrePostScript() HostBasedBackupScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *FlashbladeProtectionGroupParams) GetPrePostScriptOk() (*HostBasedBackupScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *FlashbladeProtectionGroupParams) SetPrePostScript(v HostBasedBackupScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *FlashbladeProtectionGroupParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetProtocol

`func (o *FlashbladeProtectionGroupParams) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FlashbladeProtectionGroupParams) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FlashbladeProtectionGroupParams) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FlashbladeProtectionGroupParams) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *FlashbladeProtectionGroupParams) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *FlashbladeProtectionGroupParams) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetSourceId

`func (o *FlashbladeProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *FlashbladeProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *FlashbladeProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *FlashbladeProtectionGroupParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *FlashbladeProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *FlashbladeProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *FlashbladeProtectionGroupParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *FlashbladeProtectionGroupParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *FlashbladeProtectionGroupParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *FlashbladeProtectionGroupParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *FlashbladeProtectionGroupParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *FlashbladeProtectionGroupParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetThrottlingConfig

`func (o *FlashbladeProtectionGroupParams) GetThrottlingConfig() NasThrottlingConfig`

GetThrottlingConfig returns the ThrottlingConfig field if non-nil, zero value otherwise.

### GetThrottlingConfigOk

`func (o *FlashbladeProtectionGroupParams) GetThrottlingConfigOk() (*NasThrottlingConfig, bool)`

GetThrottlingConfigOk returns a tuple with the ThrottlingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingConfig

`func (o *FlashbladeProtectionGroupParams) SetThrottlingConfig(v NasThrottlingConfig)`

SetThrottlingConfig sets ThrottlingConfig field to given value.

### HasThrottlingConfig

`func (o *FlashbladeProtectionGroupParams) HasThrottlingConfig() bool`

HasThrottlingConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


