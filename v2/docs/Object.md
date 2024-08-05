# Object

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

## Methods

### NewObject

`func NewObject() *Object`

NewObject instantiates a new Object object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectWithDefaults

`func NewObjectWithDefaults() *Object`

NewObjectWithDefaults instantiates a new Object object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *Object) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Object) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Object) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *Object) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *Object) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *Object) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetId

`func (o *Object) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Object) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Object) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Object) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Object) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Object) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Object) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Object) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Object) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Object) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Object) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Object) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSourceId

`func (o *Object) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Object) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Object) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *Object) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *Object) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *Object) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *Object) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *Object) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *Object) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *Object) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *Object) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *Object) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetChildObjects

`func (o *Object) GetChildObjects() []ObjectSummary`

GetChildObjects returns the ChildObjects field if non-nil, zero value otherwise.

### GetChildObjectsOk

`func (o *Object) GetChildObjectsOk() (*[]ObjectSummary, bool)`

GetChildObjectsOk returns a tuple with the ChildObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildObjects

`func (o *Object) SetChildObjects(v []ObjectSummary)`

SetChildObjects sets ChildObjects field to given value.

### HasChildObjects

`func (o *Object) HasChildObjects() bool`

HasChildObjects returns a boolean if a field has been set.

### SetChildObjectsNil

`func (o *Object) SetChildObjectsNil(b bool)`

 SetChildObjectsNil sets the value for ChildObjects to be an explicit nil

### UnsetChildObjects
`func (o *Object) UnsetChildObjects()`

UnsetChildObjects ensures that no value is present for ChildObjects, not even an explicit nil
### GetGlobalId

`func (o *Object) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *Object) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *Object) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *Object) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### SetGlobalIdNil

`func (o *Object) SetGlobalIdNil(b bool)`

 SetGlobalIdNil sets the value for GlobalId to be an explicit nil

### UnsetGlobalId
`func (o *Object) UnsetGlobalId()`

UnsetGlobalId ensures that no value is present for GlobalId, not even an explicit nil
### GetLogicalSizeBytes

`func (o *Object) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *Object) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *Object) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *Object) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *Object) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *Object) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
### GetObjectHash

`func (o *Object) GetObjectHash() string`

GetObjectHash returns the ObjectHash field if non-nil, zero value otherwise.

### GetObjectHashOk

`func (o *Object) GetObjectHashOk() (*string, bool)`

GetObjectHashOk returns a tuple with the ObjectHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectHash

`func (o *Object) SetObjectHash(v string)`

SetObjectHash sets ObjectHash field to given value.

### HasObjectHash

`func (o *Object) HasObjectHash() bool`

HasObjectHash returns a boolean if a field has been set.

### SetObjectHashNil

`func (o *Object) SetObjectHashNil(b bool)`

 SetObjectHashNil sets the value for ObjectHash to be an explicit nil

### UnsetObjectHash
`func (o *Object) UnsetObjectHash()`

UnsetObjectHash ensures that no value is present for ObjectHash, not even an explicit nil
### GetObjectType

`func (o *Object) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *Object) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *Object) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *Object) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectTypeNil

`func (o *Object) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *Object) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetOsType

`func (o *Object) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *Object) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *Object) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *Object) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### SetOsTypeNil

`func (o *Object) SetOsTypeNil(b bool)`

 SetOsTypeNil sets the value for OsType to be an explicit nil

### UnsetOsType
`func (o *Object) UnsetOsType()`

UnsetOsType ensures that no value is present for OsType, not even an explicit nil
### GetProtectionType

`func (o *Object) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *Object) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *Object) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *Object) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### SetProtectionTypeNil

`func (o *Object) SetProtectionTypeNil(b bool)`

 SetProtectionTypeNil sets the value for ProtectionType to be an explicit nil

### UnsetProtectionType
`func (o *Object) UnsetProtectionType()`

UnsetProtectionType ensures that no value is present for ProtectionType, not even an explicit nil
### GetSharepointSiteSummary

`func (o *Object) GetSharepointSiteSummary() SharepointObjectParams`

GetSharepointSiteSummary returns the SharepointSiteSummary field if non-nil, zero value otherwise.

### GetSharepointSiteSummaryOk

`func (o *Object) GetSharepointSiteSummaryOk() (*SharepointObjectParams, bool)`

GetSharepointSiteSummaryOk returns a tuple with the SharepointSiteSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointSiteSummary

`func (o *Object) SetSharepointSiteSummary(v SharepointObjectParams)`

SetSharepointSiteSummary sets SharepointSiteSummary field to given value.

### HasSharepointSiteSummary

`func (o *Object) HasSharepointSiteSummary() bool`

HasSharepointSiteSummary returns a boolean if a field has been set.

### GetUuid

`func (o *Object) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Object) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Object) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Object) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *Object) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *Object) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetVCenterSummary

`func (o *Object) GetVCenterSummary() ObjectTypeVCenterParams`

GetVCenterSummary returns the VCenterSummary field if non-nil, zero value otherwise.

### GetVCenterSummaryOk

`func (o *Object) GetVCenterSummaryOk() (*ObjectTypeVCenterParams, bool)`

GetVCenterSummaryOk returns a tuple with the VCenterSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCenterSummary

`func (o *Object) SetVCenterSummary(v ObjectTypeVCenterParams)`

SetVCenterSummary sets VCenterSummary field to given value.

### HasVCenterSummary

`func (o *Object) HasVCenterSummary() bool`

HasVCenterSummary returns a boolean if a field has been set.

### GetWindowsClusterSummary

`func (o *Object) GetWindowsClusterSummary() ObjectTypeWindowsClusterParams`

GetWindowsClusterSummary returns the WindowsClusterSummary field if non-nil, zero value otherwise.

### GetWindowsClusterSummaryOk

`func (o *Object) GetWindowsClusterSummaryOk() (*ObjectTypeWindowsClusterParams, bool)`

GetWindowsClusterSummaryOk returns a tuple with the WindowsClusterSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsClusterSummary

`func (o *Object) SetWindowsClusterSummary(v ObjectTypeWindowsClusterParams)`

SetWindowsClusterSummary sets WindowsClusterSummary field to given value.

### HasWindowsClusterSummary

`func (o *Object) HasWindowsClusterSummary() bool`

HasWindowsClusterSummary returns a boolean if a field has been set.

### GetPermissions

`func (o *Object) GetPermissions() PermissionInfo`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Object) GetPermissionsOk() (*PermissionInfo, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Object) SetPermissions(v PermissionInfo)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Object) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetProtectionStats

`func (o *Object) GetProtectionStats() []ObjectProtectionStatsSummary`

GetProtectionStats returns the ProtectionStats field if non-nil, zero value otherwise.

### GetProtectionStatsOk

`func (o *Object) GetProtectionStatsOk() (*[]ObjectProtectionStatsSummary, bool)`

GetProtectionStatsOk returns a tuple with the ProtectionStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionStats

`func (o *Object) SetProtectionStats(v []ObjectProtectionStatsSummary)`

SetProtectionStats sets ProtectionStats field to given value.

### HasProtectionStats

`func (o *Object) HasProtectionStats() bool`

HasProtectionStats returns a boolean if a field has been set.

### SetProtectionStatsNil

`func (o *Object) SetProtectionStatsNil(b bool)`

 SetProtectionStatsNil sets the value for ProtectionStats to be an explicit nil

### UnsetProtectionStats
`func (o *Object) UnsetProtectionStats()`

UnsetProtectionStats ensures that no value is present for ProtectionStats, not even an explicit nil
### GetElastifileParams

`func (o *Object) GetElastifileParams() map[string]interface{}`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *Object) GetElastifileParamsOk() (*map[string]interface{}, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *Object) SetElastifileParams(v map[string]interface{})`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *Object) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *Object) GetFlashbladeParams() map[string]interface{}`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *Object) GetFlashbladeParamsOk() (*map[string]interface{}, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *Object) SetFlashbladeParams(v map[string]interface{})`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *Object) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *Object) GetGenericNasParams() map[string]interface{}`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *Object) GetGenericNasParamsOk() (*map[string]interface{}, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *Object) SetGenericNasParams(v map[string]interface{})`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *Object) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *Object) GetGpfsParams() map[string]interface{}`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *Object) GetGpfsParamsOk() (*map[string]interface{}, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *Object) SetGpfsParams(v map[string]interface{})`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *Object) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### GetGroupParams

`func (o *Object) GetGroupParams() map[string]interface{}`

GetGroupParams returns the GroupParams field if non-nil, zero value otherwise.

### GetGroupParamsOk

`func (o *Object) GetGroupParamsOk() (*map[string]interface{}, bool)`

GetGroupParamsOk returns a tuple with the GroupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupParams

`func (o *Object) SetGroupParams(v map[string]interface{})`

SetGroupParams sets GroupParams field to given value.

### HasGroupParams

`func (o *Object) HasGroupParams() bool`

HasGroupParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *Object) GetIsilonParams() map[string]interface{}`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *Object) GetIsilonParamsOk() (*map[string]interface{}, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *Object) SetIsilonParams(v map[string]interface{})`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *Object) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetMssqlParams

`func (o *Object) GetMssqlParams() map[string]interface{}`

GetMssqlParams returns the MssqlParams field if non-nil, zero value otherwise.

### GetMssqlParamsOk

`func (o *Object) GetMssqlParamsOk() (*map[string]interface{}, bool)`

GetMssqlParamsOk returns a tuple with the MssqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlParams

`func (o *Object) SetMssqlParams(v map[string]interface{})`

SetMssqlParams sets MssqlParams field to given value.

### HasMssqlParams

`func (o *Object) HasMssqlParams() bool`

HasMssqlParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *Object) GetNetappParams() map[string]interface{}`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *Object) GetNetappParamsOk() (*map[string]interface{}, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *Object) SetNetappParams(v map[string]interface{})`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *Object) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetOracleParams

`func (o *Object) GetOracleParams() map[string]interface{}`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *Object) GetOracleParamsOk() (*map[string]interface{}, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *Object) SetOracleParams(v map[string]interface{})`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *Object) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *Object) GetPhysicalParams() map[string]interface{}`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *Object) GetPhysicalParamsOk() (*map[string]interface{}, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *Object) SetPhysicalParams(v map[string]interface{})`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *Object) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetSharepointParams

`func (o *Object) GetSharepointParams() map[string]interface{}`

GetSharepointParams returns the SharepointParams field if non-nil, zero value otherwise.

### GetSharepointParamsOk

`func (o *Object) GetSharepointParamsOk() (*map[string]interface{}, bool)`

GetSharepointParamsOk returns a tuple with the SharepointParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointParams

`func (o *Object) SetSharepointParams(v map[string]interface{})`

SetSharepointParams sets SharepointParams field to given value.

### HasSharepointParams

`func (o *Object) HasSharepointParams() bool`

HasSharepointParams returns a boolean if a field has been set.

### GetUdaParams

`func (o *Object) GetUdaParams() map[string]interface{}`

GetUdaParams returns the UdaParams field if non-nil, zero value otherwise.

### GetUdaParamsOk

`func (o *Object) GetUdaParamsOk() (*map[string]interface{}, bool)`

GetUdaParamsOk returns a tuple with the UdaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdaParams

`func (o *Object) SetUdaParams(v map[string]interface{})`

SetUdaParams sets UdaParams field to given value.

### HasUdaParams

`func (o *Object) HasUdaParams() bool`

HasUdaParams returns a boolean if a field has been set.

### GetViewParams

`func (o *Object) GetViewParams() map[string]interface{}`

GetViewParams returns the ViewParams field if non-nil, zero value otherwise.

### GetViewParamsOk

`func (o *Object) GetViewParamsOk() (*map[string]interface{}, bool)`

GetViewParamsOk returns a tuple with the ViewParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewParams

`func (o *Object) SetViewParams(v map[string]interface{})`

SetViewParams sets ViewParams field to given value.

### HasViewParams

`func (o *Object) HasViewParams() bool`

HasViewParams returns a boolean if a field has been set.

### GetVmwareParams

`func (o *Object) GetVmwareParams() VmwareObjectEntityParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *Object) GetVmwareParamsOk() (*VmwareObjectEntityParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *Object) SetVmwareParams(v VmwareObjectEntityParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *Object) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


