# ProtectedObject

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
**ElastifileParams** | Pointer to **map[string]interface{}** | Specifies the parameters for Elastifile object. | [optional] 
**FlashbladeParams** | Pointer to **map[string]interface{}** | Specifies the parameters for Flashblade object. | [optional] 
**GenericNasParams** | Pointer to **map[string]interface{}** | Specifies the parameters for GenericNas object. | [optional] 
**GpfsParams** | Pointer to **map[string]interface{}** | Specifies the parameters for GPFS object. | [optional] 
**GroupParams** | Pointer to **map[string]interface{}** | Specifies the parameters for M365 Group object. | [optional] 
**IsilonParams** | Pointer to **map[string]interface{}** | Specifies the parameters for Isilon object. | [optional] 
**MssqlParams** | Pointer to **map[string]interface{}** | Specifies the parameters for Msssql object. | [optional] 
**NetappParams** | Pointer to **map[string]interface{}** | Specifies the parameters for NetApp object. | [optional] 
**OracleParams** | Pointer to **map[string]interface{}** | Specifies the parameters for Oracle object. | [optional] 
**PhysicalParams** | Pointer to **map[string]interface{}** | Specifies the parameters for Physical object. | [optional] 
**SharepointParams** | Pointer to **map[string]interface{}** | Specifies the parameters for Sharepoint object. | [optional] 
**UdaParams** | Pointer to **map[string]interface{}** | Specifies the parameters for UDA object. | [optional] 
**ViewParams** | Pointer to **map[string]interface{}** | Specifies the parameters for a View. | [optional] 
**VmwareParams** | Pointer to [**VmwareObjectEntityParams**](VmwareObjectEntityParams.md) |  | [optional] 
**LatestSnapshotsInfo** | Pointer to [**[]ObjectSnapshotsInfo**](ObjectSnapshotsInfo.md) | Specifies the latest snapshot information for every Protection Group for a given object. | [optional] 
**SourceInfo** | Pointer to **map[string]interface{}** | Specifies the Source Object information. | [optional] 

## Methods

### NewProtectedObject

`func NewProtectedObject() *ProtectedObject`

NewProtectedObject instantiates a new ProtectedObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedObjectWithDefaults

`func NewProtectedObjectWithDefaults() *ProtectedObject`

NewProtectedObjectWithDefaults instantiates a new ProtectedObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ProtectedObject) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ProtectedObject) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ProtectedObject) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ProtectedObject) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ProtectedObject) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ProtectedObject) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetId

`func (o *ProtectedObject) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProtectedObject) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProtectedObject) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProtectedObject) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProtectedObject) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProtectedObject) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ProtectedObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProtectedObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProtectedObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProtectedObject) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProtectedObject) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProtectedObject) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSourceId

`func (o *ProtectedObject) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ProtectedObject) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ProtectedObject) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ProtectedObject) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *ProtectedObject) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *ProtectedObject) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *ProtectedObject) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *ProtectedObject) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *ProtectedObject) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *ProtectedObject) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *ProtectedObject) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *ProtectedObject) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetChildObjects

`func (o *ProtectedObject) GetChildObjects() []ObjectSummary`

GetChildObjects returns the ChildObjects field if non-nil, zero value otherwise.

### GetChildObjectsOk

`func (o *ProtectedObject) GetChildObjectsOk() (*[]ObjectSummary, bool)`

GetChildObjectsOk returns a tuple with the ChildObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildObjects

`func (o *ProtectedObject) SetChildObjects(v []ObjectSummary)`

SetChildObjects sets ChildObjects field to given value.

### HasChildObjects

`func (o *ProtectedObject) HasChildObjects() bool`

HasChildObjects returns a boolean if a field has been set.

### SetChildObjectsNil

`func (o *ProtectedObject) SetChildObjectsNil(b bool)`

 SetChildObjectsNil sets the value for ChildObjects to be an explicit nil

### UnsetChildObjects
`func (o *ProtectedObject) UnsetChildObjects()`

UnsetChildObjects ensures that no value is present for ChildObjects, not even an explicit nil
### GetGlobalId

`func (o *ProtectedObject) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *ProtectedObject) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *ProtectedObject) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *ProtectedObject) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### SetGlobalIdNil

`func (o *ProtectedObject) SetGlobalIdNil(b bool)`

 SetGlobalIdNil sets the value for GlobalId to be an explicit nil

### UnsetGlobalId
`func (o *ProtectedObject) UnsetGlobalId()`

UnsetGlobalId ensures that no value is present for GlobalId, not even an explicit nil
### GetLogicalSizeBytes

`func (o *ProtectedObject) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *ProtectedObject) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *ProtectedObject) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *ProtectedObject) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *ProtectedObject) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *ProtectedObject) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
### GetObjectHash

`func (o *ProtectedObject) GetObjectHash() string`

GetObjectHash returns the ObjectHash field if non-nil, zero value otherwise.

### GetObjectHashOk

`func (o *ProtectedObject) GetObjectHashOk() (*string, bool)`

GetObjectHashOk returns a tuple with the ObjectHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectHash

`func (o *ProtectedObject) SetObjectHash(v string)`

SetObjectHash sets ObjectHash field to given value.

### HasObjectHash

`func (o *ProtectedObject) HasObjectHash() bool`

HasObjectHash returns a boolean if a field has been set.

### SetObjectHashNil

`func (o *ProtectedObject) SetObjectHashNil(b bool)`

 SetObjectHashNil sets the value for ObjectHash to be an explicit nil

### UnsetObjectHash
`func (o *ProtectedObject) UnsetObjectHash()`

UnsetObjectHash ensures that no value is present for ObjectHash, not even an explicit nil
### GetObjectType

`func (o *ProtectedObject) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ProtectedObject) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ProtectedObject) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *ProtectedObject) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectTypeNil

`func (o *ProtectedObject) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *ProtectedObject) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetOsType

`func (o *ProtectedObject) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ProtectedObject) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ProtectedObject) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *ProtectedObject) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### SetOsTypeNil

`func (o *ProtectedObject) SetOsTypeNil(b bool)`

 SetOsTypeNil sets the value for OsType to be an explicit nil

### UnsetOsType
`func (o *ProtectedObject) UnsetOsType()`

UnsetOsType ensures that no value is present for OsType, not even an explicit nil
### GetProtectionType

`func (o *ProtectedObject) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *ProtectedObject) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *ProtectedObject) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *ProtectedObject) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### SetProtectionTypeNil

`func (o *ProtectedObject) SetProtectionTypeNil(b bool)`

 SetProtectionTypeNil sets the value for ProtectionType to be an explicit nil

### UnsetProtectionType
`func (o *ProtectedObject) UnsetProtectionType()`

UnsetProtectionType ensures that no value is present for ProtectionType, not even an explicit nil
### GetSharepointSiteSummary

`func (o *ProtectedObject) GetSharepointSiteSummary() SharepointObjectParams`

GetSharepointSiteSummary returns the SharepointSiteSummary field if non-nil, zero value otherwise.

### GetSharepointSiteSummaryOk

`func (o *ProtectedObject) GetSharepointSiteSummaryOk() (*SharepointObjectParams, bool)`

GetSharepointSiteSummaryOk returns a tuple with the SharepointSiteSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointSiteSummary

`func (o *ProtectedObject) SetSharepointSiteSummary(v SharepointObjectParams)`

SetSharepointSiteSummary sets SharepointSiteSummary field to given value.

### HasSharepointSiteSummary

`func (o *ProtectedObject) HasSharepointSiteSummary() bool`

HasSharepointSiteSummary returns a boolean if a field has been set.

### GetUuid

`func (o *ProtectedObject) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProtectedObject) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProtectedObject) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ProtectedObject) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ProtectedObject) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ProtectedObject) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetVCenterSummary

`func (o *ProtectedObject) GetVCenterSummary() ObjectTypeVCenterParams`

GetVCenterSummary returns the VCenterSummary field if non-nil, zero value otherwise.

### GetVCenterSummaryOk

`func (o *ProtectedObject) GetVCenterSummaryOk() (*ObjectTypeVCenterParams, bool)`

GetVCenterSummaryOk returns a tuple with the VCenterSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCenterSummary

`func (o *ProtectedObject) SetVCenterSummary(v ObjectTypeVCenterParams)`

SetVCenterSummary sets VCenterSummary field to given value.

### HasVCenterSummary

`func (o *ProtectedObject) HasVCenterSummary() bool`

HasVCenterSummary returns a boolean if a field has been set.

### GetWindowsClusterSummary

`func (o *ProtectedObject) GetWindowsClusterSummary() ObjectTypeWindowsClusterParams`

GetWindowsClusterSummary returns the WindowsClusterSummary field if non-nil, zero value otherwise.

### GetWindowsClusterSummaryOk

`func (o *ProtectedObject) GetWindowsClusterSummaryOk() (*ObjectTypeWindowsClusterParams, bool)`

GetWindowsClusterSummaryOk returns a tuple with the WindowsClusterSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsClusterSummary

`func (o *ProtectedObject) SetWindowsClusterSummary(v ObjectTypeWindowsClusterParams)`

SetWindowsClusterSummary sets WindowsClusterSummary field to given value.

### HasWindowsClusterSummary

`func (o *ProtectedObject) HasWindowsClusterSummary() bool`

HasWindowsClusterSummary returns a boolean if a field has been set.

### GetPermissions

`func (o *ProtectedObject) GetPermissions() PermissionInfo`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ProtectedObject) GetPermissionsOk() (*PermissionInfo, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ProtectedObject) SetPermissions(v PermissionInfo)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ProtectedObject) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetProtectionStats

`func (o *ProtectedObject) GetProtectionStats() []ObjectProtectionStatsSummary`

GetProtectionStats returns the ProtectionStats field if non-nil, zero value otherwise.

### GetProtectionStatsOk

`func (o *ProtectedObject) GetProtectionStatsOk() (*[]ObjectProtectionStatsSummary, bool)`

GetProtectionStatsOk returns a tuple with the ProtectionStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionStats

`func (o *ProtectedObject) SetProtectionStats(v []ObjectProtectionStatsSummary)`

SetProtectionStats sets ProtectionStats field to given value.

### HasProtectionStats

`func (o *ProtectedObject) HasProtectionStats() bool`

HasProtectionStats returns a boolean if a field has been set.

### SetProtectionStatsNil

`func (o *ProtectedObject) SetProtectionStatsNil(b bool)`

 SetProtectionStatsNil sets the value for ProtectionStats to be an explicit nil

### UnsetProtectionStats
`func (o *ProtectedObject) UnsetProtectionStats()`

UnsetProtectionStats ensures that no value is present for ProtectionStats, not even an explicit nil
### GetElastifileParams

`func (o *ProtectedObject) GetElastifileParams() map[string]interface{}`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *ProtectedObject) GetElastifileParamsOk() (*map[string]interface{}, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *ProtectedObject) SetElastifileParams(v map[string]interface{})`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *ProtectedObject) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *ProtectedObject) GetFlashbladeParams() map[string]interface{}`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *ProtectedObject) GetFlashbladeParamsOk() (*map[string]interface{}, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *ProtectedObject) SetFlashbladeParams(v map[string]interface{})`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *ProtectedObject) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *ProtectedObject) GetGenericNasParams() map[string]interface{}`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *ProtectedObject) GetGenericNasParamsOk() (*map[string]interface{}, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *ProtectedObject) SetGenericNasParams(v map[string]interface{})`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *ProtectedObject) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *ProtectedObject) GetGpfsParams() map[string]interface{}`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *ProtectedObject) GetGpfsParamsOk() (*map[string]interface{}, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *ProtectedObject) SetGpfsParams(v map[string]interface{})`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *ProtectedObject) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### GetGroupParams

`func (o *ProtectedObject) GetGroupParams() map[string]interface{}`

GetGroupParams returns the GroupParams field if non-nil, zero value otherwise.

### GetGroupParamsOk

`func (o *ProtectedObject) GetGroupParamsOk() (*map[string]interface{}, bool)`

GetGroupParamsOk returns a tuple with the GroupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupParams

`func (o *ProtectedObject) SetGroupParams(v map[string]interface{})`

SetGroupParams sets GroupParams field to given value.

### HasGroupParams

`func (o *ProtectedObject) HasGroupParams() bool`

HasGroupParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *ProtectedObject) GetIsilonParams() map[string]interface{}`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *ProtectedObject) GetIsilonParamsOk() (*map[string]interface{}, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *ProtectedObject) SetIsilonParams(v map[string]interface{})`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *ProtectedObject) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetMssqlParams

`func (o *ProtectedObject) GetMssqlParams() map[string]interface{}`

GetMssqlParams returns the MssqlParams field if non-nil, zero value otherwise.

### GetMssqlParamsOk

`func (o *ProtectedObject) GetMssqlParamsOk() (*map[string]interface{}, bool)`

GetMssqlParamsOk returns a tuple with the MssqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlParams

`func (o *ProtectedObject) SetMssqlParams(v map[string]interface{})`

SetMssqlParams sets MssqlParams field to given value.

### HasMssqlParams

`func (o *ProtectedObject) HasMssqlParams() bool`

HasMssqlParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *ProtectedObject) GetNetappParams() map[string]interface{}`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *ProtectedObject) GetNetappParamsOk() (*map[string]interface{}, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *ProtectedObject) SetNetappParams(v map[string]interface{})`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *ProtectedObject) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetOracleParams

`func (o *ProtectedObject) GetOracleParams() map[string]interface{}`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *ProtectedObject) GetOracleParamsOk() (*map[string]interface{}, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *ProtectedObject) SetOracleParams(v map[string]interface{})`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *ProtectedObject) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *ProtectedObject) GetPhysicalParams() map[string]interface{}`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *ProtectedObject) GetPhysicalParamsOk() (*map[string]interface{}, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *ProtectedObject) SetPhysicalParams(v map[string]interface{})`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *ProtectedObject) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetSharepointParams

`func (o *ProtectedObject) GetSharepointParams() map[string]interface{}`

GetSharepointParams returns the SharepointParams field if non-nil, zero value otherwise.

### GetSharepointParamsOk

`func (o *ProtectedObject) GetSharepointParamsOk() (*map[string]interface{}, bool)`

GetSharepointParamsOk returns a tuple with the SharepointParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointParams

`func (o *ProtectedObject) SetSharepointParams(v map[string]interface{})`

SetSharepointParams sets SharepointParams field to given value.

### HasSharepointParams

`func (o *ProtectedObject) HasSharepointParams() bool`

HasSharepointParams returns a boolean if a field has been set.

### GetUdaParams

`func (o *ProtectedObject) GetUdaParams() map[string]interface{}`

GetUdaParams returns the UdaParams field if non-nil, zero value otherwise.

### GetUdaParamsOk

`func (o *ProtectedObject) GetUdaParamsOk() (*map[string]interface{}, bool)`

GetUdaParamsOk returns a tuple with the UdaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdaParams

`func (o *ProtectedObject) SetUdaParams(v map[string]interface{})`

SetUdaParams sets UdaParams field to given value.

### HasUdaParams

`func (o *ProtectedObject) HasUdaParams() bool`

HasUdaParams returns a boolean if a field has been set.

### GetViewParams

`func (o *ProtectedObject) GetViewParams() map[string]interface{}`

GetViewParams returns the ViewParams field if non-nil, zero value otherwise.

### GetViewParamsOk

`func (o *ProtectedObject) GetViewParamsOk() (*map[string]interface{}, bool)`

GetViewParamsOk returns a tuple with the ViewParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewParams

`func (o *ProtectedObject) SetViewParams(v map[string]interface{})`

SetViewParams sets ViewParams field to given value.

### HasViewParams

`func (o *ProtectedObject) HasViewParams() bool`

HasViewParams returns a boolean if a field has been set.

### GetVmwareParams

`func (o *ProtectedObject) GetVmwareParams() VmwareObjectEntityParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *ProtectedObject) GetVmwareParamsOk() (*VmwareObjectEntityParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *ProtectedObject) SetVmwareParams(v VmwareObjectEntityParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *ProtectedObject) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.

### GetLatestSnapshotsInfo

`func (o *ProtectedObject) GetLatestSnapshotsInfo() []ObjectSnapshotsInfo`

GetLatestSnapshotsInfo returns the LatestSnapshotsInfo field if non-nil, zero value otherwise.

### GetLatestSnapshotsInfoOk

`func (o *ProtectedObject) GetLatestSnapshotsInfoOk() (*[]ObjectSnapshotsInfo, bool)`

GetLatestSnapshotsInfoOk returns a tuple with the LatestSnapshotsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSnapshotsInfo

`func (o *ProtectedObject) SetLatestSnapshotsInfo(v []ObjectSnapshotsInfo)`

SetLatestSnapshotsInfo sets LatestSnapshotsInfo field to given value.

### HasLatestSnapshotsInfo

`func (o *ProtectedObject) HasLatestSnapshotsInfo() bool`

HasLatestSnapshotsInfo returns a boolean if a field has been set.

### SetLatestSnapshotsInfoNil

`func (o *ProtectedObject) SetLatestSnapshotsInfoNil(b bool)`

 SetLatestSnapshotsInfoNil sets the value for LatestSnapshotsInfo to be an explicit nil

### UnsetLatestSnapshotsInfo
`func (o *ProtectedObject) UnsetLatestSnapshotsInfo()`

UnsetLatestSnapshotsInfo ensures that no value is present for LatestSnapshotsInfo, not even an explicit nil
### GetSourceInfo

`func (o *ProtectedObject) GetSourceInfo() map[string]interface{}`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *ProtectedObject) GetSourceInfoOk() (*map[string]interface{}, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *ProtectedObject) SetSourceInfo(v map[string]interface{})`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *ProtectedObject) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.

### SetSourceInfoNil

`func (o *ProtectedObject) SetSourceInfoNil(b bool)`

 SetSourceInfoNil sets the value for SourceInfo to be an explicit nil

### UnsetSourceInfo
`func (o *ProtectedObject) UnsetSourceInfo()`

UnsetSourceInfo ensures that no value is present for SourceInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


