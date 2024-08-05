# InfectedFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntivirusServiceGroupName** | Pointer to **NullableString** | Specifies the Antivirus Service group which detected the threats. | [optional] 
**AntivirusServiceIcapUri** | Pointer to **NullableString** | Specifies the ICAP Uri of the Antivirus Service which detected the threats. | [optional] 
**DetectedTimeUsecs** | Pointer to **NullableInt64** | Specifies the timestamp in microseconds when the threats were detected. | [optional] 
**EntityId** | **NullableInt64** | Specifies the entity id of the infected entity. | 
**EntityType** | Pointer to **NullableString** | Specifies the type of the infected entity. | [optional] 
**LastModifiedTimeUsecs** | Pointer to **NullableInt64** | Specifies the timestamp in microseconds when this entity was last modified. | [optional] 
**Path** | Pointer to **NullableString** | Specifies the infected entity path. | [optional] 
**RootInodeId** | **NullableInt64** | Specifies the root inode id of the file system which the infected entity belongs to. | 
**ScannedTimeUsecs** | Pointer to **NullableInt64** | Specifies the timestamp in microseconds when inode was scanned for viruses. | [optional] 
**State** | Pointer to **NullableString** | Specifies the state of the infected entity. | [optional] 
**ThreatDescriptions** | Pointer to **[]string** | Specifies a list of virus threat descriptions found in the entity. | [optional] 
**ViewId** | **NullableInt64** | Specifies the view id which the infected entity belongs to. | 
**ViewName** | Pointer to **NullableString** | Specifies the View name to which the infected entity belongs to. | [optional] 

## Methods

### NewInfectedFile

`func NewInfectedFile(entityId NullableInt64, rootInodeId NullableInt64, viewId NullableInt64, ) *InfectedFile`

NewInfectedFile instantiates a new InfectedFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfectedFileWithDefaults

`func NewInfectedFileWithDefaults() *InfectedFile`

NewInfectedFileWithDefaults instantiates a new InfectedFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntivirusServiceGroupName

`func (o *InfectedFile) GetAntivirusServiceGroupName() string`

GetAntivirusServiceGroupName returns the AntivirusServiceGroupName field if non-nil, zero value otherwise.

### GetAntivirusServiceGroupNameOk

`func (o *InfectedFile) GetAntivirusServiceGroupNameOk() (*string, bool)`

GetAntivirusServiceGroupNameOk returns a tuple with the AntivirusServiceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntivirusServiceGroupName

`func (o *InfectedFile) SetAntivirusServiceGroupName(v string)`

SetAntivirusServiceGroupName sets AntivirusServiceGroupName field to given value.

### HasAntivirusServiceGroupName

`func (o *InfectedFile) HasAntivirusServiceGroupName() bool`

HasAntivirusServiceGroupName returns a boolean if a field has been set.

### SetAntivirusServiceGroupNameNil

`func (o *InfectedFile) SetAntivirusServiceGroupNameNil(b bool)`

 SetAntivirusServiceGroupNameNil sets the value for AntivirusServiceGroupName to be an explicit nil

### UnsetAntivirusServiceGroupName
`func (o *InfectedFile) UnsetAntivirusServiceGroupName()`

UnsetAntivirusServiceGroupName ensures that no value is present for AntivirusServiceGroupName, not even an explicit nil
### GetAntivirusServiceIcapUri

`func (o *InfectedFile) GetAntivirusServiceIcapUri() string`

GetAntivirusServiceIcapUri returns the AntivirusServiceIcapUri field if non-nil, zero value otherwise.

### GetAntivirusServiceIcapUriOk

`func (o *InfectedFile) GetAntivirusServiceIcapUriOk() (*string, bool)`

GetAntivirusServiceIcapUriOk returns a tuple with the AntivirusServiceIcapUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntivirusServiceIcapUri

`func (o *InfectedFile) SetAntivirusServiceIcapUri(v string)`

SetAntivirusServiceIcapUri sets AntivirusServiceIcapUri field to given value.

### HasAntivirusServiceIcapUri

`func (o *InfectedFile) HasAntivirusServiceIcapUri() bool`

HasAntivirusServiceIcapUri returns a boolean if a field has been set.

### SetAntivirusServiceIcapUriNil

`func (o *InfectedFile) SetAntivirusServiceIcapUriNil(b bool)`

 SetAntivirusServiceIcapUriNil sets the value for AntivirusServiceIcapUri to be an explicit nil

### UnsetAntivirusServiceIcapUri
`func (o *InfectedFile) UnsetAntivirusServiceIcapUri()`

UnsetAntivirusServiceIcapUri ensures that no value is present for AntivirusServiceIcapUri, not even an explicit nil
### GetDetectedTimeUsecs

`func (o *InfectedFile) GetDetectedTimeUsecs() int64`

GetDetectedTimeUsecs returns the DetectedTimeUsecs field if non-nil, zero value otherwise.

### GetDetectedTimeUsecsOk

`func (o *InfectedFile) GetDetectedTimeUsecsOk() (*int64, bool)`

GetDetectedTimeUsecsOk returns a tuple with the DetectedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedTimeUsecs

`func (o *InfectedFile) SetDetectedTimeUsecs(v int64)`

SetDetectedTimeUsecs sets DetectedTimeUsecs field to given value.

### HasDetectedTimeUsecs

`func (o *InfectedFile) HasDetectedTimeUsecs() bool`

HasDetectedTimeUsecs returns a boolean if a field has been set.

### SetDetectedTimeUsecsNil

`func (o *InfectedFile) SetDetectedTimeUsecsNil(b bool)`

 SetDetectedTimeUsecsNil sets the value for DetectedTimeUsecs to be an explicit nil

### UnsetDetectedTimeUsecs
`func (o *InfectedFile) UnsetDetectedTimeUsecs()`

UnsetDetectedTimeUsecs ensures that no value is present for DetectedTimeUsecs, not even an explicit nil
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


### SetEntityIdNil

`func (o *InfectedFile) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *InfectedFile) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetEntityType

`func (o *InfectedFile) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *InfectedFile) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *InfectedFile) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *InfectedFile) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### SetEntityTypeNil

`func (o *InfectedFile) SetEntityTypeNil(b bool)`

 SetEntityTypeNil sets the value for EntityType to be an explicit nil

### UnsetEntityType
`func (o *InfectedFile) UnsetEntityType()`

UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
### GetLastModifiedTimeUsecs

`func (o *InfectedFile) GetLastModifiedTimeUsecs() int64`

GetLastModifiedTimeUsecs returns the LastModifiedTimeUsecs field if non-nil, zero value otherwise.

### GetLastModifiedTimeUsecsOk

`func (o *InfectedFile) GetLastModifiedTimeUsecsOk() (*int64, bool)`

GetLastModifiedTimeUsecsOk returns a tuple with the LastModifiedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTimeUsecs

`func (o *InfectedFile) SetLastModifiedTimeUsecs(v int64)`

SetLastModifiedTimeUsecs sets LastModifiedTimeUsecs field to given value.

### HasLastModifiedTimeUsecs

`func (o *InfectedFile) HasLastModifiedTimeUsecs() bool`

HasLastModifiedTimeUsecs returns a boolean if a field has been set.

### SetLastModifiedTimeUsecsNil

`func (o *InfectedFile) SetLastModifiedTimeUsecsNil(b bool)`

 SetLastModifiedTimeUsecsNil sets the value for LastModifiedTimeUsecs to be an explicit nil

### UnsetLastModifiedTimeUsecs
`func (o *InfectedFile) UnsetLastModifiedTimeUsecs()`

UnsetLastModifiedTimeUsecs ensures that no value is present for LastModifiedTimeUsecs, not even an explicit nil
### GetPath

`func (o *InfectedFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *InfectedFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *InfectedFile) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *InfectedFile) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *InfectedFile) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *InfectedFile) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
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


### SetRootInodeIdNil

`func (o *InfectedFile) SetRootInodeIdNil(b bool)`

 SetRootInodeIdNil sets the value for RootInodeId to be an explicit nil

### UnsetRootInodeId
`func (o *InfectedFile) UnsetRootInodeId()`

UnsetRootInodeId ensures that no value is present for RootInodeId, not even an explicit nil
### GetScannedTimeUsecs

`func (o *InfectedFile) GetScannedTimeUsecs() int64`

GetScannedTimeUsecs returns the ScannedTimeUsecs field if non-nil, zero value otherwise.

### GetScannedTimeUsecsOk

`func (o *InfectedFile) GetScannedTimeUsecsOk() (*int64, bool)`

GetScannedTimeUsecsOk returns a tuple with the ScannedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScannedTimeUsecs

`func (o *InfectedFile) SetScannedTimeUsecs(v int64)`

SetScannedTimeUsecs sets ScannedTimeUsecs field to given value.

### HasScannedTimeUsecs

`func (o *InfectedFile) HasScannedTimeUsecs() bool`

HasScannedTimeUsecs returns a boolean if a field has been set.

### SetScannedTimeUsecsNil

`func (o *InfectedFile) SetScannedTimeUsecsNil(b bool)`

 SetScannedTimeUsecsNil sets the value for ScannedTimeUsecs to be an explicit nil

### UnsetScannedTimeUsecs
`func (o *InfectedFile) UnsetScannedTimeUsecs()`

UnsetScannedTimeUsecs ensures that no value is present for ScannedTimeUsecs, not even an explicit nil
### GetState

`func (o *InfectedFile) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InfectedFile) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InfectedFile) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *InfectedFile) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *InfectedFile) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *InfectedFile) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
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


