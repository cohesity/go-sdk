# AzureNativeProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudMigration** | Pointer to **NullableBool** | Specifies whether or not to move the workload to the cloud. | [optional] 
**CloudPrePostScript** | Pointer to [**CloudBackupScriptParams**](CloudBackupScriptParams.md) |  | [optional] 
**DataTransferInfo** | Pointer to [**DataTransferInfo**](DataTransferInfo.md) |  | [optional] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the objects to be excluded in the Protection Group. | [optional] 
**ExcludeVmTagIds** | Pointer to **[][]int64** | Array of arrays of VM Tag Ids that Specify VMs to Exclude. | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**Objects** | Pointer to [**[]AzureNativeProtectionGroupObjectParams**](AzureNativeProtectionGroupObjectParams.md) | Specifies the objects to be included in the Protection Group. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 
**VmTagIds** | Pointer to **[][]int64** | Array of arrays of VM Tag Ids that Specify VMs to Protect. | [optional] 

## Methods

### NewAzureNativeProtectionGroupParams

`func NewAzureNativeProtectionGroupParams() *AzureNativeProtectionGroupParams`

NewAzureNativeProtectionGroupParams instantiates a new AzureNativeProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureNativeProtectionGroupParamsWithDefaults

`func NewAzureNativeProtectionGroupParamsWithDefaults() *AzureNativeProtectionGroupParams`

NewAzureNativeProtectionGroupParamsWithDefaults instantiates a new AzureNativeProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudMigration

`func (o *AzureNativeProtectionGroupParams) GetCloudMigration() bool`

GetCloudMigration returns the CloudMigration field if non-nil, zero value otherwise.

### GetCloudMigrationOk

`func (o *AzureNativeProtectionGroupParams) GetCloudMigrationOk() (*bool, bool)`

GetCloudMigrationOk returns a tuple with the CloudMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudMigration

`func (o *AzureNativeProtectionGroupParams) SetCloudMigration(v bool)`

SetCloudMigration sets CloudMigration field to given value.

### HasCloudMigration

`func (o *AzureNativeProtectionGroupParams) HasCloudMigration() bool`

HasCloudMigration returns a boolean if a field has been set.

### SetCloudMigrationNil

`func (o *AzureNativeProtectionGroupParams) SetCloudMigrationNil(b bool)`

 SetCloudMigrationNil sets the value for CloudMigration to be an explicit nil

### UnsetCloudMigration
`func (o *AzureNativeProtectionGroupParams) UnsetCloudMigration()`

UnsetCloudMigration ensures that no value is present for CloudMigration, not even an explicit nil
### GetCloudPrePostScript

`func (o *AzureNativeProtectionGroupParams) GetCloudPrePostScript() CloudBackupScriptParams`

GetCloudPrePostScript returns the CloudPrePostScript field if non-nil, zero value otherwise.

### GetCloudPrePostScriptOk

`func (o *AzureNativeProtectionGroupParams) GetCloudPrePostScriptOk() (*CloudBackupScriptParams, bool)`

GetCloudPrePostScriptOk returns a tuple with the CloudPrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPrePostScript

`func (o *AzureNativeProtectionGroupParams) SetCloudPrePostScript(v CloudBackupScriptParams)`

SetCloudPrePostScript sets CloudPrePostScript field to given value.

### HasCloudPrePostScript

`func (o *AzureNativeProtectionGroupParams) HasCloudPrePostScript() bool`

HasCloudPrePostScript returns a boolean if a field has been set.

### GetDataTransferInfo

`func (o *AzureNativeProtectionGroupParams) GetDataTransferInfo() DataTransferInfo`

GetDataTransferInfo returns the DataTransferInfo field if non-nil, zero value otherwise.

### GetDataTransferInfoOk

`func (o *AzureNativeProtectionGroupParams) GetDataTransferInfoOk() (*DataTransferInfo, bool)`

GetDataTransferInfoOk returns a tuple with the DataTransferInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferInfo

`func (o *AzureNativeProtectionGroupParams) SetDataTransferInfo(v DataTransferInfo)`

SetDataTransferInfo sets DataTransferInfo field to given value.

### HasDataTransferInfo

`func (o *AzureNativeProtectionGroupParams) HasDataTransferInfo() bool`

HasDataTransferInfo returns a boolean if a field has been set.

### GetExcludeObjectIds

`func (o *AzureNativeProtectionGroupParams) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *AzureNativeProtectionGroupParams) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *AzureNativeProtectionGroupParams) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *AzureNativeProtectionGroupParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### GetExcludeVmTagIds

`func (o *AzureNativeProtectionGroupParams) GetExcludeVmTagIds() [][]int64`

GetExcludeVmTagIds returns the ExcludeVmTagIds field if non-nil, zero value otherwise.

### GetExcludeVmTagIdsOk

`func (o *AzureNativeProtectionGroupParams) GetExcludeVmTagIdsOk() (*[][]int64, bool)`

GetExcludeVmTagIdsOk returns a tuple with the ExcludeVmTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeVmTagIds

`func (o *AzureNativeProtectionGroupParams) SetExcludeVmTagIds(v [][]int64)`

SetExcludeVmTagIds sets ExcludeVmTagIds field to given value.

### HasExcludeVmTagIds

`func (o *AzureNativeProtectionGroupParams) HasExcludeVmTagIds() bool`

HasExcludeVmTagIds returns a boolean if a field has been set.

### SetExcludeVmTagIdsNil

`func (o *AzureNativeProtectionGroupParams) SetExcludeVmTagIdsNil(b bool)`

 SetExcludeVmTagIdsNil sets the value for ExcludeVmTagIds to be an explicit nil

### UnsetExcludeVmTagIds
`func (o *AzureNativeProtectionGroupParams) UnsetExcludeVmTagIds()`

UnsetExcludeVmTagIds ensures that no value is present for ExcludeVmTagIds, not even an explicit nil
### GetIndexingPolicy

`func (o *AzureNativeProtectionGroupParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *AzureNativeProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *AzureNativeProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *AzureNativeProtectionGroupParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetObjects

`func (o *AzureNativeProtectionGroupParams) GetObjects() []AzureNativeProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *AzureNativeProtectionGroupParams) GetObjectsOk() (*[]AzureNativeProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *AzureNativeProtectionGroupParams) SetObjects(v []AzureNativeProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *AzureNativeProtectionGroupParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetSourceId

`func (o *AzureNativeProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AzureNativeProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AzureNativeProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AzureNativeProtectionGroupParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *AzureNativeProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *AzureNativeProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *AzureNativeProtectionGroupParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *AzureNativeProtectionGroupParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *AzureNativeProtectionGroupParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *AzureNativeProtectionGroupParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *AzureNativeProtectionGroupParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *AzureNativeProtectionGroupParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetVmTagIds

`func (o *AzureNativeProtectionGroupParams) GetVmTagIds() [][]int64`

GetVmTagIds returns the VmTagIds field if non-nil, zero value otherwise.

### GetVmTagIdsOk

`func (o *AzureNativeProtectionGroupParams) GetVmTagIdsOk() (*[][]int64, bool)`

GetVmTagIdsOk returns a tuple with the VmTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTagIds

`func (o *AzureNativeProtectionGroupParams) SetVmTagIds(v [][]int64)`

SetVmTagIds sets VmTagIds field to given value.

### HasVmTagIds

`func (o *AzureNativeProtectionGroupParams) HasVmTagIds() bool`

HasVmTagIds returns a boolean if a field has been set.

### SetVmTagIdsNil

`func (o *AzureNativeProtectionGroupParams) SetVmTagIdsNil(b bool)`

 SetVmTagIdsNil sets the value for VmTagIds to be an explicit nil

### UnsetVmTagIds
`func (o *AzureNativeProtectionGroupParams) UnsetVmTagIds()`

UnsetVmTagIds ensures that no value is present for VmTagIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


