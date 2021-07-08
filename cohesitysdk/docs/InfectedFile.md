# InfectedFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntivirusProviderName** | Pointer to **NullableString** | Specifies the name of antivirus service provider. | [optional] 
**EntityId** | Pointer to **NullableInt64** | Specifies the entity id of the infected file. | [optional] 
**FilePath** | Pointer to **NullableString** | Specifies file path of the infected file. | [optional] 
**InfectionDetectionTimestamp** | Pointer to **NullableInt64** | Specifies unix epoch timestamp (in microseconds) at which these threats were detected. | [optional] 
**ModifiedTimestampUsecs** | Pointer to **NullableInt64** | Specifies unix epoch timestamp (in microseconds) at which this file is modified. | [optional] 
**RemediationState** | Pointer to **NullableString** | Specifies the remediation state of the file. Remediation State. &#39;kQuarantine&#39; indicates &#39;Quarantine&#39; state of the file. This state blocks the client access. The administrator will have to manually delete, rescan or unquarantine the file. &#39;kUnquarantine&#39; indicates &#39;Unquarantine&#39; state of the file. The administrator has manually moved files from quarantined to the unquarantined state to allow client access. Unquarantined files are not scanned for virus until manually reset. | [optional] 
**RootInodeId** | Pointer to **NullableInt64** | Specifies the root inode id of the file system that infected file belongs to. | [optional] 
**ScanTimestampUsecs** | Pointer to **NullableInt64** | Specifies unix epoch timestamp (in microseconds) at which inode was scanned for viruses. | [optional] 
**ServiceIcapUri** | Pointer to **NullableString** | Specifies the instance of an antivirus ICAP server in the cluster config that detected these threats. | [optional] 
**ThreatDescriptions** | Pointer to **[]string** | Specifies the list of virus threat descriptions found in the file. | [optional] 
**ViewId** | Pointer to **NullableInt64** | Specifies the id of the View the infected file belongs to. | [optional] 
**ViewName** | Pointer to **NullableString** | Specifies the View name corresponding to above view id. | [optional] 

## Methods

### NewInfectedFile

`func NewInfectedFile() *InfectedFile`

NewInfectedFile instantiates a new InfectedFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfectedFileWithDefaults

`func NewInfectedFileWithDefaults() *InfectedFile`

NewInfectedFileWithDefaults instantiates a new InfectedFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntivirusProviderName

`func (o *InfectedFile) GetAntivirusProviderName() string`

GetAntivirusProviderName returns the AntivirusProviderName field if non-nil, zero value otherwise.

### GetAntivirusProviderNameOk

`func (o *InfectedFile) GetAntivirusProviderNameOk() (*string, bool)`

GetAntivirusProviderNameOk returns a tuple with the AntivirusProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntivirusProviderName

`func (o *InfectedFile) SetAntivirusProviderName(v string)`

SetAntivirusProviderName sets AntivirusProviderName field to given value.

### HasAntivirusProviderName

`func (o *InfectedFile) HasAntivirusProviderName() bool`

HasAntivirusProviderName returns a boolean if a field has been set.

### SetAntivirusProviderNameNil

`func (o *InfectedFile) SetAntivirusProviderNameNil(b bool)`

 SetAntivirusProviderNameNil sets the value for AntivirusProviderName to be an explicit nil

### UnsetAntivirusProviderName
`func (o *InfectedFile) UnsetAntivirusProviderName()`

UnsetAntivirusProviderName ensures that no value is present for AntivirusProviderName, not even an explicit nil
### GetEntityId

`func (o *InfectedFile) GetEntityId() int64`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *InfectedFile) GetEntityIdOk() (*int64, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *InfectedFile) SetEntityId(v int64)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *InfectedFile) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *InfectedFile) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *InfectedFile) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetFilePath

`func (o *InfectedFile) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *InfectedFile) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *InfectedFile) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *InfectedFile) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *InfectedFile) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *InfectedFile) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetInfectionDetectionTimestamp

`func (o *InfectedFile) GetInfectionDetectionTimestamp() int64`

GetInfectionDetectionTimestamp returns the InfectionDetectionTimestamp field if non-nil, zero value otherwise.

### GetInfectionDetectionTimestampOk

`func (o *InfectedFile) GetInfectionDetectionTimestampOk() (*int64, bool)`

GetInfectionDetectionTimestampOk returns a tuple with the InfectionDetectionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfectionDetectionTimestamp

`func (o *InfectedFile) SetInfectionDetectionTimestamp(v int64)`

SetInfectionDetectionTimestamp sets InfectionDetectionTimestamp field to given value.

### HasInfectionDetectionTimestamp

`func (o *InfectedFile) HasInfectionDetectionTimestamp() bool`

HasInfectionDetectionTimestamp returns a boolean if a field has been set.

### SetInfectionDetectionTimestampNil

`func (o *InfectedFile) SetInfectionDetectionTimestampNil(b bool)`

 SetInfectionDetectionTimestampNil sets the value for InfectionDetectionTimestamp to be an explicit nil

### UnsetInfectionDetectionTimestamp
`func (o *InfectedFile) UnsetInfectionDetectionTimestamp()`

UnsetInfectionDetectionTimestamp ensures that no value is present for InfectionDetectionTimestamp, not even an explicit nil
### GetModifiedTimestampUsecs

`func (o *InfectedFile) GetModifiedTimestampUsecs() int64`

GetModifiedTimestampUsecs returns the ModifiedTimestampUsecs field if non-nil, zero value otherwise.

### GetModifiedTimestampUsecsOk

`func (o *InfectedFile) GetModifiedTimestampUsecsOk() (*int64, bool)`

GetModifiedTimestampUsecsOk returns a tuple with the ModifiedTimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTimestampUsecs

`func (o *InfectedFile) SetModifiedTimestampUsecs(v int64)`

SetModifiedTimestampUsecs sets ModifiedTimestampUsecs field to given value.

### HasModifiedTimestampUsecs

`func (o *InfectedFile) HasModifiedTimestampUsecs() bool`

HasModifiedTimestampUsecs returns a boolean if a field has been set.

### SetModifiedTimestampUsecsNil

`func (o *InfectedFile) SetModifiedTimestampUsecsNil(b bool)`

 SetModifiedTimestampUsecsNil sets the value for ModifiedTimestampUsecs to be an explicit nil

### UnsetModifiedTimestampUsecs
`func (o *InfectedFile) UnsetModifiedTimestampUsecs()`

UnsetModifiedTimestampUsecs ensures that no value is present for ModifiedTimestampUsecs, not even an explicit nil
### GetRemediationState

`func (o *InfectedFile) GetRemediationState() string`

GetRemediationState returns the RemediationState field if non-nil, zero value otherwise.

### GetRemediationStateOk

`func (o *InfectedFile) GetRemediationStateOk() (*string, bool)`

GetRemediationStateOk returns a tuple with the RemediationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationState

`func (o *InfectedFile) SetRemediationState(v string)`

SetRemediationState sets RemediationState field to given value.

### HasRemediationState

`func (o *InfectedFile) HasRemediationState() bool`

HasRemediationState returns a boolean if a field has been set.

### SetRemediationStateNil

`func (o *InfectedFile) SetRemediationStateNil(b bool)`

 SetRemediationStateNil sets the value for RemediationState to be an explicit nil

### UnsetRemediationState
`func (o *InfectedFile) UnsetRemediationState()`

UnsetRemediationState ensures that no value is present for RemediationState, not even an explicit nil
### GetRootInodeId

`func (o *InfectedFile) GetRootInodeId() int64`

GetRootInodeId returns the RootInodeId field if non-nil, zero value otherwise.

### GetRootInodeIdOk

`func (o *InfectedFile) GetRootInodeIdOk() (*int64, bool)`

GetRootInodeIdOk returns a tuple with the RootInodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootInodeId

`func (o *InfectedFile) SetRootInodeId(v int64)`

SetRootInodeId sets RootInodeId field to given value.

### HasRootInodeId

`func (o *InfectedFile) HasRootInodeId() bool`

HasRootInodeId returns a boolean if a field has been set.

### SetRootInodeIdNil

`func (o *InfectedFile) SetRootInodeIdNil(b bool)`

 SetRootInodeIdNil sets the value for RootInodeId to be an explicit nil

### UnsetRootInodeId
`func (o *InfectedFile) UnsetRootInodeId()`

UnsetRootInodeId ensures that no value is present for RootInodeId, not even an explicit nil
### GetScanTimestampUsecs

`func (o *InfectedFile) GetScanTimestampUsecs() int64`

GetScanTimestampUsecs returns the ScanTimestampUsecs field if non-nil, zero value otherwise.

### GetScanTimestampUsecsOk

`func (o *InfectedFile) GetScanTimestampUsecsOk() (*int64, bool)`

GetScanTimestampUsecsOk returns a tuple with the ScanTimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanTimestampUsecs

`func (o *InfectedFile) SetScanTimestampUsecs(v int64)`

SetScanTimestampUsecs sets ScanTimestampUsecs field to given value.

### HasScanTimestampUsecs

`func (o *InfectedFile) HasScanTimestampUsecs() bool`

HasScanTimestampUsecs returns a boolean if a field has been set.

### SetScanTimestampUsecsNil

`func (o *InfectedFile) SetScanTimestampUsecsNil(b bool)`

 SetScanTimestampUsecsNil sets the value for ScanTimestampUsecs to be an explicit nil

### UnsetScanTimestampUsecs
`func (o *InfectedFile) UnsetScanTimestampUsecs()`

UnsetScanTimestampUsecs ensures that no value is present for ScanTimestampUsecs, not even an explicit nil
### GetServiceIcapUri

`func (o *InfectedFile) GetServiceIcapUri() string`

GetServiceIcapUri returns the ServiceIcapUri field if non-nil, zero value otherwise.

### GetServiceIcapUriOk

`func (o *InfectedFile) GetServiceIcapUriOk() (*string, bool)`

GetServiceIcapUriOk returns a tuple with the ServiceIcapUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIcapUri

`func (o *InfectedFile) SetServiceIcapUri(v string)`

SetServiceIcapUri sets ServiceIcapUri field to given value.

### HasServiceIcapUri

`func (o *InfectedFile) HasServiceIcapUri() bool`

HasServiceIcapUri returns a boolean if a field has been set.

### SetServiceIcapUriNil

`func (o *InfectedFile) SetServiceIcapUriNil(b bool)`

 SetServiceIcapUriNil sets the value for ServiceIcapUri to be an explicit nil

### UnsetServiceIcapUri
`func (o *InfectedFile) UnsetServiceIcapUri()`

UnsetServiceIcapUri ensures that no value is present for ServiceIcapUri, not even an explicit nil
### GetThreatDescriptions

`func (o *InfectedFile) GetThreatDescriptions() []string`

GetThreatDescriptions returns the ThreatDescriptions field if non-nil, zero value otherwise.

### GetThreatDescriptionsOk

`func (o *InfectedFile) GetThreatDescriptionsOk() (*[]string, bool)`

GetThreatDescriptionsOk returns a tuple with the ThreatDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatDescriptions

`func (o *InfectedFile) SetThreatDescriptions(v []string)`

SetThreatDescriptions sets ThreatDescriptions field to given value.

### HasThreatDescriptions

`func (o *InfectedFile) HasThreatDescriptions() bool`

HasThreatDescriptions returns a boolean if a field has been set.

### SetThreatDescriptionsNil

`func (o *InfectedFile) SetThreatDescriptionsNil(b bool)`

 SetThreatDescriptionsNil sets the value for ThreatDescriptions to be an explicit nil

### UnsetThreatDescriptions
`func (o *InfectedFile) UnsetThreatDescriptions()`

UnsetThreatDescriptions ensures that no value is present for ThreatDescriptions, not even an explicit nil
### GetViewId

`func (o *InfectedFile) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *InfectedFile) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *InfectedFile) SetViewId(v int64)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *InfectedFile) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *InfectedFile) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *InfectedFile) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil
### GetViewName

`func (o *InfectedFile) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *InfectedFile) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *InfectedFile) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *InfectedFile) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *InfectedFile) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *InfectedFile) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


