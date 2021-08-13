# ObjectWithChildren

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
**Objects** | Pointer to [**[]ObjectWithChildren**](ObjectWithChildren.md) | Specifies a list of child nodes for this specific node. | [optional] 

## Methods

### NewObjectWithChildren

`func NewObjectWithChildren() *ObjectWithChildren`

NewObjectWithChildren instantiates a new ObjectWithChildren object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectWithChildrenWithDefaults

`func NewObjectWithChildrenWithDefaults() *ObjectWithChildren`

NewObjectWithChildrenWithDefaults instantiates a new ObjectWithChildren object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectWithChildren) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectWithChildren) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectWithChildren) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectWithChildren) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ObjectWithChildren) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ObjectWithChildren) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ObjectWithChildren) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectWithChildren) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectWithChildren) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectWithChildren) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ObjectWithChildren) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ObjectWithChildren) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSourceId

`func (o *ObjectWithChildren) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ObjectWithChildren) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ObjectWithChildren) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ObjectWithChildren) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *ObjectWithChildren) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *ObjectWithChildren) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *ObjectWithChildren) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *ObjectWithChildren) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *ObjectWithChildren) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *ObjectWithChildren) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *ObjectWithChildren) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *ObjectWithChildren) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetEnvironment

`func (o *ObjectWithChildren) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ObjectWithChildren) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ObjectWithChildren) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ObjectWithChildren) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ObjectWithChildren) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ObjectWithChildren) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetObjectHash

`func (o *ObjectWithChildren) GetObjectHash() string`

GetObjectHash returns the ObjectHash field if non-nil, zero value otherwise.

### GetObjectHashOk

`func (o *ObjectWithChildren) GetObjectHashOk() (*string, bool)`

GetObjectHashOk returns a tuple with the ObjectHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectHash

`func (o *ObjectWithChildren) SetObjectHash(v string)`

SetObjectHash sets ObjectHash field to given value.

### HasObjectHash

`func (o *ObjectWithChildren) HasObjectHash() bool`

HasObjectHash returns a boolean if a field has been set.

### SetObjectHashNil

`func (o *ObjectWithChildren) SetObjectHashNil(b bool)`

 SetObjectHashNil sets the value for ObjectHash to be an explicit nil

### UnsetObjectHash
`func (o *ObjectWithChildren) UnsetObjectHash()`

UnsetObjectHash ensures that no value is present for ObjectHash, not even an explicit nil
### GetObjectType

`func (o *ObjectWithChildren) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ObjectWithChildren) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ObjectWithChildren) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *ObjectWithChildren) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectTypeNil

`func (o *ObjectWithChildren) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *ObjectWithChildren) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetLogicalSizeBytes

`func (o *ObjectWithChildren) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *ObjectWithChildren) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *ObjectWithChildren) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *ObjectWithChildren) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *ObjectWithChildren) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *ObjectWithChildren) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
### GetUuid

`func (o *ObjectWithChildren) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ObjectWithChildren) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ObjectWithChildren) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ObjectWithChildren) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ObjectWithChildren) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ObjectWithChildren) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetGlobalId

`func (o *ObjectWithChildren) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *ObjectWithChildren) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *ObjectWithChildren) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *ObjectWithChildren) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### SetGlobalIdNil

`func (o *ObjectWithChildren) SetGlobalIdNil(b bool)`

 SetGlobalIdNil sets the value for GlobalId to be an explicit nil

### UnsetGlobalId
`func (o *ObjectWithChildren) UnsetGlobalId()`

UnsetGlobalId ensures that no value is present for GlobalId, not even an explicit nil
### GetProtectionType

`func (o *ObjectWithChildren) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *ObjectWithChildren) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *ObjectWithChildren) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *ObjectWithChildren) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### SetProtectionTypeNil

`func (o *ObjectWithChildren) SetProtectionTypeNil(b bool)`

 SetProtectionTypeNil sets the value for ProtectionType to be an explicit nil

### UnsetProtectionType
`func (o *ObjectWithChildren) UnsetProtectionType()`

UnsetProtectionType ensures that no value is present for ProtectionType, not even an explicit nil
### GetOsType

`func (o *ObjectWithChildren) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ObjectWithChildren) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ObjectWithChildren) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *ObjectWithChildren) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### SetOsTypeNil

`func (o *ObjectWithChildren) SetOsTypeNil(b bool)`

 SetOsTypeNil sets the value for OsType to be an explicit nil

### UnsetOsType
`func (o *ObjectWithChildren) UnsetOsType()`

UnsetOsType ensures that no value is present for OsType, not even an explicit nil
### GetVCenterSummary

`func (o *ObjectWithChildren) GetVCenterSummary() ObjectTypeVCenterParams`

GetVCenterSummary returns the VCenterSummary field if non-nil, zero value otherwise.

### GetVCenterSummaryOk

`func (o *ObjectWithChildren) GetVCenterSummaryOk() (*ObjectTypeVCenterParams, bool)`

GetVCenterSummaryOk returns a tuple with the VCenterSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCenterSummary

`func (o *ObjectWithChildren) SetVCenterSummary(v ObjectTypeVCenterParams)`

SetVCenterSummary sets VCenterSummary field to given value.

### HasVCenterSummary

`func (o *ObjectWithChildren) HasVCenterSummary() bool`

HasVCenterSummary returns a boolean if a field has been set.

### GetSharepointSiteSummary

`func (o *ObjectWithChildren) GetSharepointSiteSummary() SharepointObjectParams`

GetSharepointSiteSummary returns the SharepointSiteSummary field if non-nil, zero value otherwise.

### GetSharepointSiteSummaryOk

`func (o *ObjectWithChildren) GetSharepointSiteSummaryOk() (*SharepointObjectParams, bool)`

GetSharepointSiteSummaryOk returns a tuple with the SharepointSiteSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointSiteSummary

`func (o *ObjectWithChildren) SetSharepointSiteSummary(v SharepointObjectParams)`

SetSharepointSiteSummary sets SharepointSiteSummary field to given value.

### HasSharepointSiteSummary

`func (o *ObjectWithChildren) HasSharepointSiteSummary() bool`

HasSharepointSiteSummary returns a boolean if a field has been set.

### GetProtectionStats

`func (o *ObjectWithChildren) GetProtectionStats() []ObjectProtectionStatsSummary`

GetProtectionStats returns the ProtectionStats field if non-nil, zero value otherwise.

### GetProtectionStatsOk

`func (o *ObjectWithChildren) GetProtectionStatsOk() (*[]ObjectProtectionStatsSummary, bool)`

GetProtectionStatsOk returns a tuple with the ProtectionStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionStats

`func (o *ObjectWithChildren) SetProtectionStats(v []ObjectProtectionStatsSummary)`

SetProtectionStats sets ProtectionStats field to given value.

### HasProtectionStats

`func (o *ObjectWithChildren) HasProtectionStats() bool`

HasProtectionStats returns a boolean if a field has been set.

### SetProtectionStatsNil

`func (o *ObjectWithChildren) SetProtectionStatsNil(b bool)`

 SetProtectionStatsNil sets the value for ProtectionStats to be an explicit nil

### UnsetProtectionStats
`func (o *ObjectWithChildren) UnsetProtectionStats()`

UnsetProtectionStats ensures that no value is present for ProtectionStats, not even an explicit nil
### GetPermissions

`func (o *ObjectWithChildren) GetPermissions() PermissionInfo`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ObjectWithChildren) GetPermissionsOk() (*PermissionInfo, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ObjectWithChildren) SetPermissions(v PermissionInfo)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ObjectWithChildren) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetVmwareParams

`func (o *ObjectWithChildren) GetVmwareParams() VmwareObjectEntityParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *ObjectWithChildren) GetVmwareParamsOk() (*VmwareObjectEntityParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *ObjectWithChildren) SetVmwareParams(v VmwareObjectEntityParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *ObjectWithChildren) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *ObjectWithChildren) GetIsilonParams() IsilonObjectParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *ObjectWithChildren) GetIsilonParamsOk() (*IsilonObjectParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *ObjectWithChildren) SetIsilonParams(v IsilonObjectParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *ObjectWithChildren) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *ObjectWithChildren) GetNetappParams() NetappObjectParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *ObjectWithChildren) GetNetappParamsOk() (*NetappObjectParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *ObjectWithChildren) SetNetappParams(v NetappObjectParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *ObjectWithChildren) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *ObjectWithChildren) GetGenericNasParams() CommonNasObjectParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *ObjectWithChildren) GetGenericNasParamsOk() (*CommonNasObjectParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *ObjectWithChildren) SetGenericNasParams(v CommonNasObjectParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *ObjectWithChildren) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *ObjectWithChildren) GetFlashbladeParams() FlashbladeObjectParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *ObjectWithChildren) GetFlashbladeParamsOk() (*FlashbladeObjectParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *ObjectWithChildren) SetFlashbladeParams(v FlashbladeObjectParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *ObjectWithChildren) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *ObjectWithChildren) GetElastifileParams() CommonNasObjectParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *ObjectWithChildren) GetElastifileParamsOk() (*CommonNasObjectParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *ObjectWithChildren) SetElastifileParams(v CommonNasObjectParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *ObjectWithChildren) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *ObjectWithChildren) GetGpfsParams() CommonNasObjectParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *ObjectWithChildren) GetGpfsParamsOk() (*CommonNasObjectParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *ObjectWithChildren) SetGpfsParams(v CommonNasObjectParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *ObjectWithChildren) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### GetMssqlParams

`func (o *ObjectWithChildren) GetMssqlParams() MssqlObjectEntityParams`

GetMssqlParams returns the MssqlParams field if non-nil, zero value otherwise.

### GetMssqlParamsOk

`func (o *ObjectWithChildren) GetMssqlParamsOk() (*MssqlObjectEntityParams, bool)`

GetMssqlParamsOk returns a tuple with the MssqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlParams

`func (o *ObjectWithChildren) SetMssqlParams(v MssqlObjectEntityParams)`

SetMssqlParams sets MssqlParams field to given value.

### HasMssqlParams

`func (o *ObjectWithChildren) HasMssqlParams() bool`

HasMssqlParams returns a boolean if a field has been set.

### GetOracleParams

`func (o *ObjectWithChildren) GetOracleParams() OracleObjectEntityParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *ObjectWithChildren) GetOracleParamsOk() (*OracleObjectEntityParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *ObjectWithChildren) SetOracleParams(v OracleObjectEntityParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *ObjectWithChildren) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *ObjectWithChildren) GetPhysicalParams() PhysicalObjectEntityParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *ObjectWithChildren) GetPhysicalParamsOk() (*PhysicalObjectEntityParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *ObjectWithChildren) SetPhysicalParams(v PhysicalObjectEntityParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *ObjectWithChildren) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetSharepointParams

`func (o *ObjectWithChildren) GetSharepointParams() SharepointObjectEntityParams`

GetSharepointParams returns the SharepointParams field if non-nil, zero value otherwise.

### GetSharepointParamsOk

`func (o *ObjectWithChildren) GetSharepointParamsOk() (*SharepointObjectEntityParams, bool)`

GetSharepointParamsOk returns a tuple with the SharepointParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointParams

`func (o *ObjectWithChildren) SetSharepointParams(v SharepointObjectEntityParams)`

SetSharepointParams sets SharepointParams field to given value.

### HasSharepointParams

`func (o *ObjectWithChildren) HasSharepointParams() bool`

HasSharepointParams returns a boolean if a field has been set.

### GetObjects

`func (o *ObjectWithChildren) GetObjects() []ObjectWithChildren`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ObjectWithChildren) GetObjectsOk() (*[]ObjectWithChildren, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ObjectWithChildren) SetObjects(v []ObjectWithChildren)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ObjectWithChildren) HasObjects() bool`

HasObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


