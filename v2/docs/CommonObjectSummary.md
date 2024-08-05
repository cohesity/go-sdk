# CommonObjectSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Specifies the environment of the object. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies object id. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the object. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies registered source id to which object belongs. | [optional] 
**SourceName** | Pointer to **NullableString** | Specifies registered source name to which object belongs. | [optional] 
**ChildObjects** | Pointer to [**[]ObjectSummary**](ObjectSummary.md) | Specifies child object details. | [optional] 
**GlobalId** | Pointer to **NullableString** | Specifies the global id which is a unique identifier of the object. | [optional] 
**LogicalSizeBytes** | Pointer to **NullableInt64** | Specifies the logical size of object in bytes. | [optional] 
**ObjectHash** | Pointer to **NullableString** | Specifies the hash identifier of the object. | [optional] 
**ObjectType** | Pointer to **NullableString** | Specifies the type of the object. | [optional] 
**OsType** | Pointer to **NullableString** | Specifies the operating system type of the object. | [optional] 
**ProtectionType** | Pointer to **NullableString** | Specifies the protection type of the object if any. | [optional] 
**SharepointSiteSummary** | Pointer to [**SharepointObjectParams**](SharepointObjectParams.md) |  | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the uuid which is a unique identifier of the object. | [optional] 
**VCenterSummary** | Pointer to [**ObjectTypeVCenterParams**](ObjectTypeVCenterParams.md) |  | [optional] 
**WindowsClusterSummary** | Pointer to [**ObjectTypeWindowsClusterParams**](ObjectTypeWindowsClusterParams.md) |  | [optional] 
**Permissions** | Pointer to [**PermissionInfo**](PermissionInfo.md) |  | [optional] 
**ProtectionStats** | Pointer to [**[]ObjectProtectionStatsSummary**](ObjectProtectionStatsSummary.md) | Specifies the count and size of protected and unprotected objects for the size. | [optional] 

## Methods

### NewCommonObjectSummary

`func NewCommonObjectSummary() *CommonObjectSummary`

NewCommonObjectSummary instantiates a new CommonObjectSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonObjectSummaryWithDefaults

`func NewCommonObjectSummaryWithDefaults() *CommonObjectSummary`

NewCommonObjectSummaryWithDefaults instantiates a new CommonObjectSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *CommonObjectSummary) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CommonObjectSummary) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CommonObjectSummary) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CommonObjectSummary) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *CommonObjectSummary) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *CommonObjectSummary) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetId

`func (o *CommonObjectSummary) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonObjectSummary) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonObjectSummary) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CommonObjectSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CommonObjectSummary) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CommonObjectSummary) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *CommonObjectSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonObjectSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonObjectSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommonObjectSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CommonObjectSummary) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommonObjectSummary) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSourceId

`func (o *CommonObjectSummary) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *CommonObjectSummary) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *CommonObjectSummary) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *CommonObjectSummary) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *CommonObjectSummary) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *CommonObjectSummary) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *CommonObjectSummary) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *CommonObjectSummary) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *CommonObjectSummary) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *CommonObjectSummary) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *CommonObjectSummary) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *CommonObjectSummary) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetChildObjects

`func (o *CommonObjectSummary) GetChildObjects() []ObjectSummary`

GetChildObjects returns the ChildObjects field if non-nil, zero value otherwise.

### GetChildObjectsOk

`func (o *CommonObjectSummary) GetChildObjectsOk() (*[]ObjectSummary, bool)`

GetChildObjectsOk returns a tuple with the ChildObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildObjects

`func (o *CommonObjectSummary) SetChildObjects(v []ObjectSummary)`

SetChildObjects sets ChildObjects field to given value.

### HasChildObjects

`func (o *CommonObjectSummary) HasChildObjects() bool`

HasChildObjects returns a boolean if a field has been set.

### SetChildObjectsNil

`func (o *CommonObjectSummary) SetChildObjectsNil(b bool)`

 SetChildObjectsNil sets the value for ChildObjects to be an explicit nil

### UnsetChildObjects
`func (o *CommonObjectSummary) UnsetChildObjects()`

UnsetChildObjects ensures that no value is present for ChildObjects, not even an explicit nil
### GetGlobalId

`func (o *CommonObjectSummary) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *CommonObjectSummary) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *CommonObjectSummary) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *CommonObjectSummary) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### SetGlobalIdNil

`func (o *CommonObjectSummary) SetGlobalIdNil(b bool)`

 SetGlobalIdNil sets the value for GlobalId to be an explicit nil

### UnsetGlobalId
`func (o *CommonObjectSummary) UnsetGlobalId()`

UnsetGlobalId ensures that no value is present for GlobalId, not even an explicit nil
### GetLogicalSizeBytes

`func (o *CommonObjectSummary) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *CommonObjectSummary) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *CommonObjectSummary) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *CommonObjectSummary) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *CommonObjectSummary) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *CommonObjectSummary) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
### GetObjectHash

`func (o *CommonObjectSummary) GetObjectHash() string`

GetObjectHash returns the ObjectHash field if non-nil, zero value otherwise.

### GetObjectHashOk

`func (o *CommonObjectSummary) GetObjectHashOk() (*string, bool)`

GetObjectHashOk returns a tuple with the ObjectHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectHash

`func (o *CommonObjectSummary) SetObjectHash(v string)`

SetObjectHash sets ObjectHash field to given value.

### HasObjectHash

`func (o *CommonObjectSummary) HasObjectHash() bool`

HasObjectHash returns a boolean if a field has been set.

### SetObjectHashNil

`func (o *CommonObjectSummary) SetObjectHashNil(b bool)`

 SetObjectHashNil sets the value for ObjectHash to be an explicit nil

### UnsetObjectHash
`func (o *CommonObjectSummary) UnsetObjectHash()`

UnsetObjectHash ensures that no value is present for ObjectHash, not even an explicit nil
### GetObjectType

`func (o *CommonObjectSummary) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CommonObjectSummary) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CommonObjectSummary) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *CommonObjectSummary) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectTypeNil

`func (o *CommonObjectSummary) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *CommonObjectSummary) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetOsType

`func (o *CommonObjectSummary) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *CommonObjectSummary) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *CommonObjectSummary) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *CommonObjectSummary) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### SetOsTypeNil

`func (o *CommonObjectSummary) SetOsTypeNil(b bool)`

 SetOsTypeNil sets the value for OsType to be an explicit nil

### UnsetOsType
`func (o *CommonObjectSummary) UnsetOsType()`

UnsetOsType ensures that no value is present for OsType, not even an explicit nil
### GetProtectionType

`func (o *CommonObjectSummary) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *CommonObjectSummary) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *CommonObjectSummary) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *CommonObjectSummary) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### SetProtectionTypeNil

`func (o *CommonObjectSummary) SetProtectionTypeNil(b bool)`

 SetProtectionTypeNil sets the value for ProtectionType to be an explicit nil

### UnsetProtectionType
`func (o *CommonObjectSummary) UnsetProtectionType()`

UnsetProtectionType ensures that no value is present for ProtectionType, not even an explicit nil
### GetSharepointSiteSummary

`func (o *CommonObjectSummary) GetSharepointSiteSummary() SharepointObjectParams`

GetSharepointSiteSummary returns the SharepointSiteSummary field if non-nil, zero value otherwise.

### GetSharepointSiteSummaryOk

`func (o *CommonObjectSummary) GetSharepointSiteSummaryOk() (*SharepointObjectParams, bool)`

GetSharepointSiteSummaryOk returns a tuple with the SharepointSiteSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointSiteSummary

`func (o *CommonObjectSummary) SetSharepointSiteSummary(v SharepointObjectParams)`

SetSharepointSiteSummary sets SharepointSiteSummary field to given value.

### HasSharepointSiteSummary

`func (o *CommonObjectSummary) HasSharepointSiteSummary() bool`

HasSharepointSiteSummary returns a boolean if a field has been set.

### GetUuid

`func (o *CommonObjectSummary) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CommonObjectSummary) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CommonObjectSummary) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CommonObjectSummary) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *CommonObjectSummary) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *CommonObjectSummary) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetVCenterSummary

`func (o *CommonObjectSummary) GetVCenterSummary() ObjectTypeVCenterParams`

GetVCenterSummary returns the VCenterSummary field if non-nil, zero value otherwise.

### GetVCenterSummaryOk

`func (o *CommonObjectSummary) GetVCenterSummaryOk() (*ObjectTypeVCenterParams, bool)`

GetVCenterSummaryOk returns a tuple with the VCenterSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCenterSummary

`func (o *CommonObjectSummary) SetVCenterSummary(v ObjectTypeVCenterParams)`

SetVCenterSummary sets VCenterSummary field to given value.

### HasVCenterSummary

`func (o *CommonObjectSummary) HasVCenterSummary() bool`

HasVCenterSummary returns a boolean if a field has been set.

### GetWindowsClusterSummary

`func (o *CommonObjectSummary) GetWindowsClusterSummary() ObjectTypeWindowsClusterParams`

GetWindowsClusterSummary returns the WindowsClusterSummary field if non-nil, zero value otherwise.

### GetWindowsClusterSummaryOk

`func (o *CommonObjectSummary) GetWindowsClusterSummaryOk() (*ObjectTypeWindowsClusterParams, bool)`

GetWindowsClusterSummaryOk returns a tuple with the WindowsClusterSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsClusterSummary

`func (o *CommonObjectSummary) SetWindowsClusterSummary(v ObjectTypeWindowsClusterParams)`

SetWindowsClusterSummary sets WindowsClusterSummary field to given value.

### HasWindowsClusterSummary

`func (o *CommonObjectSummary) HasWindowsClusterSummary() bool`

HasWindowsClusterSummary returns a boolean if a field has been set.

### GetPermissions

`func (o *CommonObjectSummary) GetPermissions() PermissionInfo`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CommonObjectSummary) GetPermissionsOk() (*PermissionInfo, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CommonObjectSummary) SetPermissions(v PermissionInfo)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CommonObjectSummary) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetProtectionStats

`func (o *CommonObjectSummary) GetProtectionStats() []ObjectProtectionStatsSummary`

GetProtectionStats returns the ProtectionStats field if non-nil, zero value otherwise.

### GetProtectionStatsOk

`func (o *CommonObjectSummary) GetProtectionStatsOk() (*[]ObjectProtectionStatsSummary, bool)`

GetProtectionStatsOk returns a tuple with the ProtectionStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionStats

`func (o *CommonObjectSummary) SetProtectionStats(v []ObjectProtectionStatsSummary)`

SetProtectionStats sets ProtectionStats field to given value.

### HasProtectionStats

`func (o *CommonObjectSummary) HasProtectionStats() bool`

HasProtectionStats returns a boolean if a field has been set.

### SetProtectionStatsNil

`func (o *CommonObjectSummary) SetProtectionStatsNil(b bool)`

 SetProtectionStatsNil sets the value for ProtectionStats to be an explicit nil

### UnsetProtectionStats
`func (o *CommonObjectSummary) UnsetProtectionStats()`

UnsetProtectionStats ensures that no value is present for ProtectionStats, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


