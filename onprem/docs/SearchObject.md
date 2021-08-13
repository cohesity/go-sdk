# SearchObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies object id. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the object. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies registered source id to which object belongs. | [optional] 
**SourceName** | Pointer to **NullableString** | Specifies registered source name to which object belongs. | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the environment of the object. | [optional] 
**ObjectHash** | Pointer to **NullableString** | Specifies the hash identifier of the object. | [optional] 
**ObjectType** | Pointer to **NullableString** | Specifies the type of the object. | [optional] 
**LogicalSizeBytes** | Pointer to **NullableInt64** | Specifies the logical size of object in bytes. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the uuid which is a unique identifier of the object. | [optional] 
**GlobalId** | Pointer to **NullableString** | Specifies the global id which is a unique identifier of the object. | [optional] 
**ProtectionType** | Pointer to **NullableString** | Specifies the protection type of the object if any. | [optional] 
**OsType** | Pointer to **NullableString** | Specifies the operating system type of the object. | [optional] 
**VCenterSummary** | Pointer to [**ObjectTypeVCenterParams**](ObjectTypeVCenterParams.md) |  | [optional] 
**SharepointSiteSummary** | Pointer to [**SharepointObjectParams**](SharepointObjectParams.md) |  | [optional] 
**ProtectionStats** | Pointer to [**[]ObjectProtectionStatsSummary**](ObjectProtectionStatsSummary.md) | Specifies the count and size of protected and unprotected objects for the size. | [optional] 
**Permissions** | Pointer to [**PermissionInfo**](PermissionInfo.md) |  | [optional] 
**VmwareParams** | Pointer to [**VmwareObjectEntityParams**](VmwareObjectEntityParams.md) |  | [optional] 
**IsilonParams** | Pointer to [**IsilonObjectParams**](IsilonObjectParams.md) | Specifies the parameters for Isilon object. | [optional] 
**NetappParams** | Pointer to [**NetappObjectParams**](NetappObjectParams.md) | Specifies the parameters for NetApp object. | [optional] 
**GenericNasParams** | Pointer to [**CommonNasObjectParams**](CommonNasObjectParams.md) | Specifies the parameters for GenericNas object. | [optional] 
**FlashbladeParams** | Pointer to [**FlashbladeObjectParams**](FlashbladeObjectParams.md) | Specifies the parameters for Flashblade object. | [optional] 
**ElastifileParams** | Pointer to [**CommonNasObjectParams**](CommonNasObjectParams.md) | Specifies the parameters for Elastifile object. | [optional] 
**GpfsParams** | Pointer to [**CommonNasObjectParams**](CommonNasObjectParams.md) | Specifies the parameters for GPFS object. | [optional] 
**MssqlParams** | Pointer to [**MssqlObjectEntityParams**](MssqlObjectEntityParams.md) | Specifies the parameters for Msssql object. | [optional] 
**OracleParams** | Pointer to [**OracleObjectEntityParams**](OracleObjectEntityParams.md) | Specifies the parameters for Oracle object. | [optional] 
**PhysicalParams** | Pointer to [**PhysicalObjectEntityParams**](PhysicalObjectEntityParams.md) | Specifies the parameters for Physical object. | [optional] 
**SharepointParams** | Pointer to [**SharepointObjectEntityParams**](SharepointObjectEntityParams.md) | Specifies the parameters for Sharepoint object. | [optional] 
**Tags** | Pointer to [**[]TagInfo**](TagInfo.md) | Specifies tag applied to the object. | [optional] 
**SnapshotTags** | Pointer to [**[]SnapshotTagInfo**](SnapshotTagInfo.md) | Specifies snapshot tags applied to the object. | [optional] 
**IsDeleted** | Pointer to **NullableBool** | Specifies whether the object is deleted. Deleted objects can&#39;t be protected but can be recovered. | [optional] 
**SourceInfo** | Pointer to [**Object**](Object.md) | Specifies the Source Object information. | [optional] 
**ObjectProtectionInfos** | Pointer to [**[]ObjectProtectionInfo**](ObjectProtectionInfo.md) | Specifies the object info on each cluster. | [optional] 

## Methods

### NewSearchObject

`func NewSearchObject() *SearchObject`

NewSearchObject instantiates a new SearchObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchObjectWithDefaults

`func NewSearchObjectWithDefaults() *SearchObject`

NewSearchObjectWithDefaults instantiates a new SearchObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchObject) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchObject) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchObject) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SearchObject) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SearchObject) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SearchObject) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *SearchObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchObject) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SearchObject) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SearchObject) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSourceId

`func (o *SearchObject) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *SearchObject) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *SearchObject) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *SearchObject) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *SearchObject) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *SearchObject) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *SearchObject) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *SearchObject) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *SearchObject) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *SearchObject) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *SearchObject) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *SearchObject) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetEnvironment

`func (o *SearchObject) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SearchObject) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SearchObject) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SearchObject) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *SearchObject) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *SearchObject) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetObjectHash

`func (o *SearchObject) GetObjectHash() string`

GetObjectHash returns the ObjectHash field if non-nil, zero value otherwise.

### GetObjectHashOk

`func (o *SearchObject) GetObjectHashOk() (*string, bool)`

GetObjectHashOk returns a tuple with the ObjectHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectHash

`func (o *SearchObject) SetObjectHash(v string)`

SetObjectHash sets ObjectHash field to given value.

### HasObjectHash

`func (o *SearchObject) HasObjectHash() bool`

HasObjectHash returns a boolean if a field has been set.

### SetObjectHashNil

`func (o *SearchObject) SetObjectHashNil(b bool)`

 SetObjectHashNil sets the value for ObjectHash to be an explicit nil

### UnsetObjectHash
`func (o *SearchObject) UnsetObjectHash()`

UnsetObjectHash ensures that no value is present for ObjectHash, not even an explicit nil
### GetObjectType

`func (o *SearchObject) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SearchObject) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SearchObject) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *SearchObject) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectTypeNil

`func (o *SearchObject) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *SearchObject) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetLogicalSizeBytes

`func (o *SearchObject) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *SearchObject) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *SearchObject) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *SearchObject) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *SearchObject) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *SearchObject) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
### GetUuid

`func (o *SearchObject) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SearchObject) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SearchObject) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SearchObject) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *SearchObject) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *SearchObject) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetGlobalId

`func (o *SearchObject) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *SearchObject) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *SearchObject) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *SearchObject) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### SetGlobalIdNil

`func (o *SearchObject) SetGlobalIdNil(b bool)`

 SetGlobalIdNil sets the value for GlobalId to be an explicit nil

### UnsetGlobalId
`func (o *SearchObject) UnsetGlobalId()`

UnsetGlobalId ensures that no value is present for GlobalId, not even an explicit nil
### GetProtectionType

`func (o *SearchObject) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *SearchObject) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *SearchObject) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *SearchObject) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### SetProtectionTypeNil

`func (o *SearchObject) SetProtectionTypeNil(b bool)`

 SetProtectionTypeNil sets the value for ProtectionType to be an explicit nil

### UnsetProtectionType
`func (o *SearchObject) UnsetProtectionType()`

UnsetProtectionType ensures that no value is present for ProtectionType, not even an explicit nil
### GetOsType

`func (o *SearchObject) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *SearchObject) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *SearchObject) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *SearchObject) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### SetOsTypeNil

`func (o *SearchObject) SetOsTypeNil(b bool)`

 SetOsTypeNil sets the value for OsType to be an explicit nil

### UnsetOsType
`func (o *SearchObject) UnsetOsType()`

UnsetOsType ensures that no value is present for OsType, not even an explicit nil
### GetVCenterSummary

`func (o *SearchObject) GetVCenterSummary() ObjectTypeVCenterParams`

GetVCenterSummary returns the VCenterSummary field if non-nil, zero value otherwise.

### GetVCenterSummaryOk

`func (o *SearchObject) GetVCenterSummaryOk() (*ObjectTypeVCenterParams, bool)`

GetVCenterSummaryOk returns a tuple with the VCenterSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCenterSummary

`func (o *SearchObject) SetVCenterSummary(v ObjectTypeVCenterParams)`

SetVCenterSummary sets VCenterSummary field to given value.

### HasVCenterSummary

`func (o *SearchObject) HasVCenterSummary() bool`

HasVCenterSummary returns a boolean if a field has been set.

### GetSharepointSiteSummary

`func (o *SearchObject) GetSharepointSiteSummary() SharepointObjectParams`

GetSharepointSiteSummary returns the SharepointSiteSummary field if non-nil, zero value otherwise.

### GetSharepointSiteSummaryOk

`func (o *SearchObject) GetSharepointSiteSummaryOk() (*SharepointObjectParams, bool)`

GetSharepointSiteSummaryOk returns a tuple with the SharepointSiteSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointSiteSummary

`func (o *SearchObject) SetSharepointSiteSummary(v SharepointObjectParams)`

SetSharepointSiteSummary sets SharepointSiteSummary field to given value.

### HasSharepointSiteSummary

`func (o *SearchObject) HasSharepointSiteSummary() bool`

HasSharepointSiteSummary returns a boolean if a field has been set.

### GetProtectionStats

`func (o *SearchObject) GetProtectionStats() []ObjectProtectionStatsSummary`

GetProtectionStats returns the ProtectionStats field if non-nil, zero value otherwise.

### GetProtectionStatsOk

`func (o *SearchObject) GetProtectionStatsOk() (*[]ObjectProtectionStatsSummary, bool)`

GetProtectionStatsOk returns a tuple with the ProtectionStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionStats

`func (o *SearchObject) SetProtectionStats(v []ObjectProtectionStatsSummary)`

SetProtectionStats sets ProtectionStats field to given value.

### HasProtectionStats

`func (o *SearchObject) HasProtectionStats() bool`

HasProtectionStats returns a boolean if a field has been set.

### SetProtectionStatsNil

`func (o *SearchObject) SetProtectionStatsNil(b bool)`

 SetProtectionStatsNil sets the value for ProtectionStats to be an explicit nil

### UnsetProtectionStats
`func (o *SearchObject) UnsetProtectionStats()`

UnsetProtectionStats ensures that no value is present for ProtectionStats, not even an explicit nil
### GetPermissions

`func (o *SearchObject) GetPermissions() PermissionInfo`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SearchObject) GetPermissionsOk() (*PermissionInfo, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SearchObject) SetPermissions(v PermissionInfo)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SearchObject) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetVmwareParams

`func (o *SearchObject) GetVmwareParams() VmwareObjectEntityParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *SearchObject) GetVmwareParamsOk() (*VmwareObjectEntityParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *SearchObject) SetVmwareParams(v VmwareObjectEntityParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *SearchObject) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *SearchObject) GetIsilonParams() IsilonObjectParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *SearchObject) GetIsilonParamsOk() (*IsilonObjectParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *SearchObject) SetIsilonParams(v IsilonObjectParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *SearchObject) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *SearchObject) GetNetappParams() NetappObjectParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *SearchObject) GetNetappParamsOk() (*NetappObjectParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *SearchObject) SetNetappParams(v NetappObjectParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *SearchObject) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *SearchObject) GetGenericNasParams() CommonNasObjectParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *SearchObject) GetGenericNasParamsOk() (*CommonNasObjectParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *SearchObject) SetGenericNasParams(v CommonNasObjectParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *SearchObject) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *SearchObject) GetFlashbladeParams() FlashbladeObjectParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *SearchObject) GetFlashbladeParamsOk() (*FlashbladeObjectParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *SearchObject) SetFlashbladeParams(v FlashbladeObjectParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *SearchObject) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *SearchObject) GetElastifileParams() CommonNasObjectParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *SearchObject) GetElastifileParamsOk() (*CommonNasObjectParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *SearchObject) SetElastifileParams(v CommonNasObjectParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *SearchObject) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *SearchObject) GetGpfsParams() CommonNasObjectParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *SearchObject) GetGpfsParamsOk() (*CommonNasObjectParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *SearchObject) SetGpfsParams(v CommonNasObjectParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *SearchObject) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### GetMssqlParams

`func (o *SearchObject) GetMssqlParams() MssqlObjectEntityParams`

GetMssqlParams returns the MssqlParams field if non-nil, zero value otherwise.

### GetMssqlParamsOk

`func (o *SearchObject) GetMssqlParamsOk() (*MssqlObjectEntityParams, bool)`

GetMssqlParamsOk returns a tuple with the MssqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlParams

`func (o *SearchObject) SetMssqlParams(v MssqlObjectEntityParams)`

SetMssqlParams sets MssqlParams field to given value.

### HasMssqlParams

`func (o *SearchObject) HasMssqlParams() bool`

HasMssqlParams returns a boolean if a field has been set.

### GetOracleParams

`func (o *SearchObject) GetOracleParams() OracleObjectEntityParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *SearchObject) GetOracleParamsOk() (*OracleObjectEntityParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *SearchObject) SetOracleParams(v OracleObjectEntityParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *SearchObject) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *SearchObject) GetPhysicalParams() PhysicalObjectEntityParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *SearchObject) GetPhysicalParamsOk() (*PhysicalObjectEntityParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *SearchObject) SetPhysicalParams(v PhysicalObjectEntityParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *SearchObject) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetSharepointParams

`func (o *SearchObject) GetSharepointParams() SharepointObjectEntityParams`

GetSharepointParams returns the SharepointParams field if non-nil, zero value otherwise.

### GetSharepointParamsOk

`func (o *SearchObject) GetSharepointParamsOk() (*SharepointObjectEntityParams, bool)`

GetSharepointParamsOk returns a tuple with the SharepointParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointParams

`func (o *SearchObject) SetSharepointParams(v SharepointObjectEntityParams)`

SetSharepointParams sets SharepointParams field to given value.

### HasSharepointParams

`func (o *SearchObject) HasSharepointParams() bool`

HasSharepointParams returns a boolean if a field has been set.

### GetTags

`func (o *SearchObject) GetTags() []TagInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SearchObject) GetTagsOk() (*[]TagInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SearchObject) SetTags(v []TagInfo)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SearchObject) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *SearchObject) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *SearchObject) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetSnapshotTags

`func (o *SearchObject) GetSnapshotTags() []SnapshotTagInfo`

GetSnapshotTags returns the SnapshotTags field if non-nil, zero value otherwise.

### GetSnapshotTagsOk

`func (o *SearchObject) GetSnapshotTagsOk() (*[]SnapshotTagInfo, bool)`

GetSnapshotTagsOk returns a tuple with the SnapshotTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTags

`func (o *SearchObject) SetSnapshotTags(v []SnapshotTagInfo)`

SetSnapshotTags sets SnapshotTags field to given value.

### HasSnapshotTags

`func (o *SearchObject) HasSnapshotTags() bool`

HasSnapshotTags returns a boolean if a field has been set.

### SetSnapshotTagsNil

`func (o *SearchObject) SetSnapshotTagsNil(b bool)`

 SetSnapshotTagsNil sets the value for SnapshotTags to be an explicit nil

### UnsetSnapshotTags
`func (o *SearchObject) UnsetSnapshotTags()`

UnsetSnapshotTags ensures that no value is present for SnapshotTags, not even an explicit nil
### GetIsDeleted

`func (o *SearchObject) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *SearchObject) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *SearchObject) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *SearchObject) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *SearchObject) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *SearchObject) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetSourceInfo

`func (o *SearchObject) GetSourceInfo() Object`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *SearchObject) GetSourceInfoOk() (*Object, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *SearchObject) SetSourceInfo(v Object)`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *SearchObject) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.

### GetObjectProtectionInfos

`func (o *SearchObject) GetObjectProtectionInfos() []ObjectProtectionInfo`

GetObjectProtectionInfos returns the ObjectProtectionInfos field if non-nil, zero value otherwise.

### GetObjectProtectionInfosOk

`func (o *SearchObject) GetObjectProtectionInfosOk() (*[]ObjectProtectionInfo, bool)`

GetObjectProtectionInfosOk returns a tuple with the ObjectProtectionInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectProtectionInfos

`func (o *SearchObject) SetObjectProtectionInfos(v []ObjectProtectionInfo)`

SetObjectProtectionInfos sets ObjectProtectionInfos field to given value.

### HasObjectProtectionInfos

`func (o *SearchObject) HasObjectProtectionInfos() bool`

HasObjectProtectionInfos returns a boolean if a field has been set.

### SetObjectProtectionInfosNil

`func (o *SearchObject) SetObjectProtectionInfosNil(b bool)`

 SetObjectProtectionInfosNil sets the value for ObjectProtectionInfos to be an explicit nil

### UnsetObjectProtectionInfos
`func (o *SearchObject) UnsetObjectProtectionInfos()`

UnsetObjectProtectionInfos ensures that no value is present for ObjectProtectionInfos, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


