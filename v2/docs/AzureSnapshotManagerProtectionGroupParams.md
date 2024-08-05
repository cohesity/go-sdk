# AzureSnapshotManagerProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudPrePostScript** | Pointer to [**CloudBackupScriptParams**](CloudBackupScriptParams.md) |  | [optional] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the objects to be excluded in the Protection Group. | [optional] 
**ExcludeVmTagIds** | Pointer to **[][]int64** | Array of arrays of VM Tag Ids that Specify VMs to Exclude. | [optional] 
**Objects** | Pointer to [**[]AzureSnapshotManagerProtectionGroupObjectParams**](AzureSnapshotManagerProtectionGroupObjectParams.md) | Specifies the objects to be included in the Protection Group. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 
**VmTagIds** | Pointer to **[][]int64** | Array of arrays of VM Tag Ids that Specify VMs to Protect. | [optional] 

## Methods

### NewAzureSnapshotManagerProtectionGroupParams

`func NewAzureSnapshotManagerProtectionGroupParams() *AzureSnapshotManagerProtectionGroupParams`

NewAzureSnapshotManagerProtectionGroupParams instantiates a new AzureSnapshotManagerProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSnapshotManagerProtectionGroupParamsWithDefaults

`func NewAzureSnapshotManagerProtectionGroupParamsWithDefaults() *AzureSnapshotManagerProtectionGroupParams`

NewAzureSnapshotManagerProtectionGroupParamsWithDefaults instantiates a new AzureSnapshotManagerProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudPrePostScript

`func (o *AzureSnapshotManagerProtectionGroupParams) GetCloudPrePostScript() CloudBackupScriptParams`

GetCloudPrePostScript returns the CloudPrePostScript field if non-nil, zero value otherwise.

### GetCloudPrePostScriptOk

`func (o *AzureSnapshotManagerProtectionGroupParams) GetCloudPrePostScriptOk() (*CloudBackupScriptParams, bool)`

GetCloudPrePostScriptOk returns a tuple with the CloudPrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPrePostScript

`func (o *AzureSnapshotManagerProtectionGroupParams) SetCloudPrePostScript(v CloudBackupScriptParams)`

SetCloudPrePostScript sets CloudPrePostScript field to given value.

### HasCloudPrePostScript

`func (o *AzureSnapshotManagerProtectionGroupParams) HasCloudPrePostScript() bool`

HasCloudPrePostScript returns a boolean if a field has been set.

### GetExcludeObjectIds

`func (o *AzureSnapshotManagerProtectionGroupParams) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *AzureSnapshotManagerProtectionGroupParams) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *AzureSnapshotManagerProtectionGroupParams) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *AzureSnapshotManagerProtectionGroupParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### SetExcludeObjectIdsNil

`func (o *AzureSnapshotManagerProtectionGroupParams) SetExcludeObjectIdsNil(b bool)`

 SetExcludeObjectIdsNil sets the value for ExcludeObjectIds to be an explicit nil

### UnsetExcludeObjectIds
`func (o *AzureSnapshotManagerProtectionGroupParams) UnsetExcludeObjectIds()`

UnsetExcludeObjectIds ensures that no value is present for ExcludeObjectIds, not even an explicit nil
### GetExcludeVmTagIds

`func (o *AzureSnapshotManagerProtectionGroupParams) GetExcludeVmTagIds() [][]int64`

GetExcludeVmTagIds returns the ExcludeVmTagIds field if non-nil, zero value otherwise.

### GetExcludeVmTagIdsOk

`func (o *AzureSnapshotManagerProtectionGroupParams) GetExcludeVmTagIdsOk() (*[][]int64, bool)`

GetExcludeVmTagIdsOk returns a tuple with the ExcludeVmTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeVmTagIds

`func (o *AzureSnapshotManagerProtectionGroupParams) SetExcludeVmTagIds(v [][]int64)`

SetExcludeVmTagIds sets ExcludeVmTagIds field to given value.

### HasExcludeVmTagIds

`func (o *AzureSnapshotManagerProtectionGroupParams) HasExcludeVmTagIds() bool`

HasExcludeVmTagIds returns a boolean if a field has been set.

### SetExcludeVmTagIdsNil

`func (o *AzureSnapshotManagerProtectionGroupParams) SetExcludeVmTagIdsNil(b bool)`

 SetExcludeVmTagIdsNil sets the value for ExcludeVmTagIds to be an explicit nil

### UnsetExcludeVmTagIds
`func (o *AzureSnapshotManagerProtectionGroupParams) UnsetExcludeVmTagIds()`

UnsetExcludeVmTagIds ensures that no value is present for ExcludeVmTagIds, not even an explicit nil
### GetObjects

`func (o *AzureSnapshotManagerProtectionGroupParams) GetObjects() []AzureSnapshotManagerProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *AzureSnapshotManagerProtectionGroupParams) GetObjectsOk() (*[]AzureSnapshotManagerProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *AzureSnapshotManagerProtectionGroupParams) SetObjects(v []AzureSnapshotManagerProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *AzureSnapshotManagerProtectionGroupParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetSourceId

`func (o *AzureSnapshotManagerProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AzureSnapshotManagerProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AzureSnapshotManagerProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AzureSnapshotManagerProtectionGroupParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *AzureSnapshotManagerProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *AzureSnapshotManagerProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *AzureSnapshotManagerProtectionGroupParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *AzureSnapshotManagerProtectionGroupParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *AzureSnapshotManagerProtectionGroupParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *AzureSnapshotManagerProtectionGroupParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *AzureSnapshotManagerProtectionGroupParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *AzureSnapshotManagerProtectionGroupParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetVmTagIds

`func (o *AzureSnapshotManagerProtectionGroupParams) GetVmTagIds() [][]int64`

GetVmTagIds returns the VmTagIds field if non-nil, zero value otherwise.

### GetVmTagIdsOk

`func (o *AzureSnapshotManagerProtectionGroupParams) GetVmTagIdsOk() (*[][]int64, bool)`

GetVmTagIdsOk returns a tuple with the VmTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTagIds

`func (o *AzureSnapshotManagerProtectionGroupParams) SetVmTagIds(v [][]int64)`

SetVmTagIds sets VmTagIds field to given value.

### HasVmTagIds

`func (o *AzureSnapshotManagerProtectionGroupParams) HasVmTagIds() bool`

HasVmTagIds returns a boolean if a field has been set.

### SetVmTagIdsNil

`func (o *AzureSnapshotManagerProtectionGroupParams) SetVmTagIdsNil(b bool)`

 SetVmTagIdsNil sets the value for VmTagIds to be an explicit nil

### UnsetVmTagIds
`func (o *AzureSnapshotManagerProtectionGroupParams) UnsetVmTagIds()`

UnsetVmTagIds ensures that no value is present for VmTagIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


