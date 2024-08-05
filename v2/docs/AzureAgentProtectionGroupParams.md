# AzureAgentProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppConsistentSnapshot** | Pointer to **NullableBool** | Specifies whether or not to quiesce apps and the file system in order to take app consistent snapshots. | [optional] 
**CloudMigration** | Pointer to **NullableBool** | Specifies whether or not to move the workload to the cloud. | [optional] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the objects to be excluded in the Protection Group. | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**Objects** | [**[]AzureAgentProtectionGroupObjectParams**](AzureAgentProtectionGroupObjectParams.md) | Specifies the objects to be included in the Protection Group. | 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 

## Methods

### NewAzureAgentProtectionGroupParams

`func NewAzureAgentProtectionGroupParams(objects []AzureAgentProtectionGroupObjectParams, ) *AzureAgentProtectionGroupParams`

NewAzureAgentProtectionGroupParams instantiates a new AzureAgentProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureAgentProtectionGroupParamsWithDefaults

`func NewAzureAgentProtectionGroupParamsWithDefaults() *AzureAgentProtectionGroupParams`

NewAzureAgentProtectionGroupParamsWithDefaults instantiates a new AzureAgentProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppConsistentSnapshot

`func (o *AzureAgentProtectionGroupParams) GetAppConsistentSnapshot() bool`

GetAppConsistentSnapshot returns the AppConsistentSnapshot field if non-nil, zero value otherwise.

### GetAppConsistentSnapshotOk

`func (o *AzureAgentProtectionGroupParams) GetAppConsistentSnapshotOk() (*bool, bool)`

GetAppConsistentSnapshotOk returns a tuple with the AppConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppConsistentSnapshot

`func (o *AzureAgentProtectionGroupParams) SetAppConsistentSnapshot(v bool)`

SetAppConsistentSnapshot sets AppConsistentSnapshot field to given value.

### HasAppConsistentSnapshot

`func (o *AzureAgentProtectionGroupParams) HasAppConsistentSnapshot() bool`

HasAppConsistentSnapshot returns a boolean if a field has been set.

### SetAppConsistentSnapshotNil

`func (o *AzureAgentProtectionGroupParams) SetAppConsistentSnapshotNil(b bool)`

 SetAppConsistentSnapshotNil sets the value for AppConsistentSnapshot to be an explicit nil

### UnsetAppConsistentSnapshot
`func (o *AzureAgentProtectionGroupParams) UnsetAppConsistentSnapshot()`

UnsetAppConsistentSnapshot ensures that no value is present for AppConsistentSnapshot, not even an explicit nil
### GetCloudMigration

`func (o *AzureAgentProtectionGroupParams) GetCloudMigration() bool`

GetCloudMigration returns the CloudMigration field if non-nil, zero value otherwise.

### GetCloudMigrationOk

`func (o *AzureAgentProtectionGroupParams) GetCloudMigrationOk() (*bool, bool)`

GetCloudMigrationOk returns a tuple with the CloudMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudMigration

`func (o *AzureAgentProtectionGroupParams) SetCloudMigration(v bool)`

SetCloudMigration sets CloudMigration field to given value.

### HasCloudMigration

`func (o *AzureAgentProtectionGroupParams) HasCloudMigration() bool`

HasCloudMigration returns a boolean if a field has been set.

### SetCloudMigrationNil

`func (o *AzureAgentProtectionGroupParams) SetCloudMigrationNil(b bool)`

 SetCloudMigrationNil sets the value for CloudMigration to be an explicit nil

### UnsetCloudMigration
`func (o *AzureAgentProtectionGroupParams) UnsetCloudMigration()`

UnsetCloudMigration ensures that no value is present for CloudMigration, not even an explicit nil
### GetExcludeObjectIds

`func (o *AzureAgentProtectionGroupParams) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *AzureAgentProtectionGroupParams) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *AzureAgentProtectionGroupParams) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *AzureAgentProtectionGroupParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### GetIndexingPolicy

`func (o *AzureAgentProtectionGroupParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *AzureAgentProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *AzureAgentProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *AzureAgentProtectionGroupParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetObjects

`func (o *AzureAgentProtectionGroupParams) GetObjects() []AzureAgentProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *AzureAgentProtectionGroupParams) GetObjectsOk() (*[]AzureAgentProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *AzureAgentProtectionGroupParams) SetObjects(v []AzureAgentProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.


### GetSourceId

`func (o *AzureAgentProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AzureAgentProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AzureAgentProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AzureAgentProtectionGroupParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *AzureAgentProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *AzureAgentProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *AzureAgentProtectionGroupParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *AzureAgentProtectionGroupParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *AzureAgentProtectionGroupParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *AzureAgentProtectionGroupParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *AzureAgentProtectionGroupParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *AzureAgentProtectionGroupParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


