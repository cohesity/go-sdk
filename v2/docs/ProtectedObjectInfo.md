# ProtectedObjectInfo

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
**LastRun** | Pointer to [**ObjectProtectionRunSummary**](ObjectProtectionRunSummary.md) |  | [optional] 
**ObjectBackupConfiguration** | Pointer to [**ProtectedObjectBackupConfig**](ProtectedObjectBackupConfig.md) |  | [optional] 
**Permissions** | Pointer to [**[]Tenant**](Tenant.md) | Specifies the list of tenants that have permissions for this accessing given protected object. | [optional] 
**ProtectionGroupConfigurations** | Pointer to [**[]ProtectedObjectGroupBackupConfig**](ProtectedObjectGroupBackupConfig.md) | Specifies the protection info associated with every object. There can be multiple instances of protection info since the same object can be protected in multiple protection groups. | [optional] 

## Methods

### NewProtectedObjectInfo

`func NewProtectedObjectInfo() *ProtectedObjectInfo`

NewProtectedObjectInfo instantiates a new ProtectedObjectInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedObjectInfoWithDefaults

`func NewProtectedObjectInfoWithDefaults() *ProtectedObjectInfo`

NewProtectedObjectInfoWithDefaults instantiates a new ProtectedObjectInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ProtectedObjectInfo) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ProtectedObjectInfo) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ProtectedObjectInfo) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ProtectedObjectInfo) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ProtectedObjectInfo) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ProtectedObjectInfo) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetId

`func (o *ProtectedObjectInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProtectedObjectInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProtectedObjectInfo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProtectedObjectInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProtectedObjectInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProtectedObjectInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ProtectedObjectInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProtectedObjectInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProtectedObjectInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProtectedObjectInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProtectedObjectInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProtectedObjectInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSourceId

`func (o *ProtectedObjectInfo) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ProtectedObjectInfo) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ProtectedObjectInfo) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ProtectedObjectInfo) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *ProtectedObjectInfo) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *ProtectedObjectInfo) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *ProtectedObjectInfo) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *ProtectedObjectInfo) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *ProtectedObjectInfo) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *ProtectedObjectInfo) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *ProtectedObjectInfo) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *ProtectedObjectInfo) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetChildObjects

`func (o *ProtectedObjectInfo) GetChildObjects() []ObjectSummary`

GetChildObjects returns the ChildObjects field if non-nil, zero value otherwise.

### GetChildObjectsOk

`func (o *ProtectedObjectInfo) GetChildObjectsOk() (*[]ObjectSummary, bool)`

GetChildObjectsOk returns a tuple with the ChildObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildObjects

`func (o *ProtectedObjectInfo) SetChildObjects(v []ObjectSummary)`

SetChildObjects sets ChildObjects field to given value.

### HasChildObjects

`func (o *ProtectedObjectInfo) HasChildObjects() bool`

HasChildObjects returns a boolean if a field has been set.

### SetChildObjectsNil

`func (o *ProtectedObjectInfo) SetChildObjectsNil(b bool)`

 SetChildObjectsNil sets the value for ChildObjects to be an explicit nil

### UnsetChildObjects
`func (o *ProtectedObjectInfo) UnsetChildObjects()`

UnsetChildObjects ensures that no value is present for ChildObjects, not even an explicit nil
### GetGlobalId

`func (o *ProtectedObjectInfo) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *ProtectedObjectInfo) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *ProtectedObjectInfo) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *ProtectedObjectInfo) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### SetGlobalIdNil

`func (o *ProtectedObjectInfo) SetGlobalIdNil(b bool)`

 SetGlobalIdNil sets the value for GlobalId to be an explicit nil

### UnsetGlobalId
`func (o *ProtectedObjectInfo) UnsetGlobalId()`

UnsetGlobalId ensures that no value is present for GlobalId, not even an explicit nil
### GetLogicalSizeBytes

`func (o *ProtectedObjectInfo) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *ProtectedObjectInfo) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *ProtectedObjectInfo) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *ProtectedObjectInfo) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *ProtectedObjectInfo) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *ProtectedObjectInfo) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
### GetObjectHash

`func (o *ProtectedObjectInfo) GetObjectHash() string`

GetObjectHash returns the ObjectHash field if non-nil, zero value otherwise.

### GetObjectHashOk

`func (o *ProtectedObjectInfo) GetObjectHashOk() (*string, bool)`

GetObjectHashOk returns a tuple with the ObjectHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectHash

`func (o *ProtectedObjectInfo) SetObjectHash(v string)`

SetObjectHash sets ObjectHash field to given value.

### HasObjectHash

`func (o *ProtectedObjectInfo) HasObjectHash() bool`

HasObjectHash returns a boolean if a field has been set.

### SetObjectHashNil

`func (o *ProtectedObjectInfo) SetObjectHashNil(b bool)`

 SetObjectHashNil sets the value for ObjectHash to be an explicit nil

### UnsetObjectHash
`func (o *ProtectedObjectInfo) UnsetObjectHash()`

UnsetObjectHash ensures that no value is present for ObjectHash, not even an explicit nil
### GetObjectType

`func (o *ProtectedObjectInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ProtectedObjectInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ProtectedObjectInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *ProtectedObjectInfo) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectTypeNil

`func (o *ProtectedObjectInfo) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *ProtectedObjectInfo) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetOsType

`func (o *ProtectedObjectInfo) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ProtectedObjectInfo) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ProtectedObjectInfo) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *ProtectedObjectInfo) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### SetOsTypeNil

`func (o *ProtectedObjectInfo) SetOsTypeNil(b bool)`

 SetOsTypeNil sets the value for OsType to be an explicit nil

### UnsetOsType
`func (o *ProtectedObjectInfo) UnsetOsType()`

UnsetOsType ensures that no value is present for OsType, not even an explicit nil
### GetProtectionType

`func (o *ProtectedObjectInfo) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *ProtectedObjectInfo) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *ProtectedObjectInfo) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *ProtectedObjectInfo) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### SetProtectionTypeNil

`func (o *ProtectedObjectInfo) SetProtectionTypeNil(b bool)`

 SetProtectionTypeNil sets the value for ProtectionType to be an explicit nil

### UnsetProtectionType
`func (o *ProtectedObjectInfo) UnsetProtectionType()`

UnsetProtectionType ensures that no value is present for ProtectionType, not even an explicit nil
### GetSharepointSiteSummary

`func (o *ProtectedObjectInfo) GetSharepointSiteSummary() SharepointObjectParams`

GetSharepointSiteSummary returns the SharepointSiteSummary field if non-nil, zero value otherwise.

### GetSharepointSiteSummaryOk

`func (o *ProtectedObjectInfo) GetSharepointSiteSummaryOk() (*SharepointObjectParams, bool)`

GetSharepointSiteSummaryOk returns a tuple with the SharepointSiteSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointSiteSummary

`func (o *ProtectedObjectInfo) SetSharepointSiteSummary(v SharepointObjectParams)`

SetSharepointSiteSummary sets SharepointSiteSummary field to given value.

### HasSharepointSiteSummary

`func (o *ProtectedObjectInfo) HasSharepointSiteSummary() bool`

HasSharepointSiteSummary returns a boolean if a field has been set.

### GetUuid

`func (o *ProtectedObjectInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProtectedObjectInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProtectedObjectInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ProtectedObjectInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ProtectedObjectInfo) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ProtectedObjectInfo) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetVCenterSummary

`func (o *ProtectedObjectInfo) GetVCenterSummary() ObjectTypeVCenterParams`

GetVCenterSummary returns the VCenterSummary field if non-nil, zero value otherwise.

### GetVCenterSummaryOk

`func (o *ProtectedObjectInfo) GetVCenterSummaryOk() (*ObjectTypeVCenterParams, bool)`

GetVCenterSummaryOk returns a tuple with the VCenterSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCenterSummary

`func (o *ProtectedObjectInfo) SetVCenterSummary(v ObjectTypeVCenterParams)`

SetVCenterSummary sets VCenterSummary field to given value.

### HasVCenterSummary

`func (o *ProtectedObjectInfo) HasVCenterSummary() bool`

HasVCenterSummary returns a boolean if a field has been set.

### GetWindowsClusterSummary

`func (o *ProtectedObjectInfo) GetWindowsClusterSummary() ObjectTypeWindowsClusterParams`

GetWindowsClusterSummary returns the WindowsClusterSummary field if non-nil, zero value otherwise.

### GetWindowsClusterSummaryOk

`func (o *ProtectedObjectInfo) GetWindowsClusterSummaryOk() (*ObjectTypeWindowsClusterParams, bool)`

GetWindowsClusterSummaryOk returns a tuple with the WindowsClusterSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsClusterSummary

`func (o *ProtectedObjectInfo) SetWindowsClusterSummary(v ObjectTypeWindowsClusterParams)`

SetWindowsClusterSummary sets WindowsClusterSummary field to given value.

### HasWindowsClusterSummary

`func (o *ProtectedObjectInfo) HasWindowsClusterSummary() bool`

HasWindowsClusterSummary returns a boolean if a field has been set.

### GetLastRun

`func (o *ProtectedObjectInfo) GetLastRun() ObjectProtectionRunSummary`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *ProtectedObjectInfo) GetLastRunOk() (*ObjectProtectionRunSummary, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *ProtectedObjectInfo) SetLastRun(v ObjectProtectionRunSummary)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *ProtectedObjectInfo) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetObjectBackupConfiguration

`func (o *ProtectedObjectInfo) GetObjectBackupConfiguration() ProtectedObjectBackupConfig`

GetObjectBackupConfiguration returns the ObjectBackupConfiguration field if non-nil, zero value otherwise.

### GetObjectBackupConfigurationOk

`func (o *ProtectedObjectInfo) GetObjectBackupConfigurationOk() (*ProtectedObjectBackupConfig, bool)`

GetObjectBackupConfigurationOk returns a tuple with the ObjectBackupConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectBackupConfiguration

`func (o *ProtectedObjectInfo) SetObjectBackupConfiguration(v ProtectedObjectBackupConfig)`

SetObjectBackupConfiguration sets ObjectBackupConfiguration field to given value.

### HasObjectBackupConfiguration

`func (o *ProtectedObjectInfo) HasObjectBackupConfiguration() bool`

HasObjectBackupConfiguration returns a boolean if a field has been set.

### GetPermissions

`func (o *ProtectedObjectInfo) GetPermissions() []Tenant`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ProtectedObjectInfo) GetPermissionsOk() (*[]Tenant, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ProtectedObjectInfo) SetPermissions(v []Tenant)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ProtectedObjectInfo) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *ProtectedObjectInfo) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *ProtectedObjectInfo) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetProtectionGroupConfigurations

`func (o *ProtectedObjectInfo) GetProtectionGroupConfigurations() []ProtectedObjectGroupBackupConfig`

GetProtectionGroupConfigurations returns the ProtectionGroupConfigurations field if non-nil, zero value otherwise.

### GetProtectionGroupConfigurationsOk

`func (o *ProtectedObjectInfo) GetProtectionGroupConfigurationsOk() (*[]ProtectedObjectGroupBackupConfig, bool)`

GetProtectionGroupConfigurationsOk returns a tuple with the ProtectionGroupConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupConfigurations

`func (o *ProtectedObjectInfo) SetProtectionGroupConfigurations(v []ProtectedObjectGroupBackupConfig)`

SetProtectionGroupConfigurations sets ProtectionGroupConfigurations field to given value.

### HasProtectionGroupConfigurations

`func (o *ProtectedObjectInfo) HasProtectionGroupConfigurations() bool`

HasProtectionGroupConfigurations returns a boolean if a field has been set.

### SetProtectionGroupConfigurationsNil

`func (o *ProtectedObjectInfo) SetProtectionGroupConfigurationsNil(b bool)`

 SetProtectionGroupConfigurationsNil sets the value for ProtectionGroupConfigurations to be an explicit nil

### UnsetProtectionGroupConfigurations
`func (o *ProtectedObjectInfo) UnsetProtectionGroupConfigurations()`

UnsetProtectionGroupConfigurations ensures that no value is present for ProtectionGroupConfigurations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


