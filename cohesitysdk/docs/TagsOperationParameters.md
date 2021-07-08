# TagsOperationParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | ClusterId is the Id of the cluster used for constructing JobUid. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | ClusterIncarnationId is the incarnation Id of the cluster used for constructing JobUid. | [optional] 
**DocumentIds** | Pointer to **[]string** | DocumentIds are list of documents to be tagged. | [optional] 
**EntityId** | Pointer to **NullableInt64** | EntityId is the Id of the entity where the file resides. | [optional] 
**JobId** | Pointer to **NullableInt64** | JobId is the Id of the job that took the snapshot. | [optional] 
**JobInstanceIds** | Pointer to **[]int64** | JobInstanceIds to tag corresponding snapshots. | [optional] 
**Tags** | Pointer to **[]string** | Tags are list of tags that will be operated on to corresponding objects. | [optional] 

## Methods

### NewTagsOperationParameters

`func NewTagsOperationParameters() *TagsOperationParameters`

NewTagsOperationParameters instantiates a new TagsOperationParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagsOperationParametersWithDefaults

`func NewTagsOperationParametersWithDefaults() *TagsOperationParameters`

NewTagsOperationParametersWithDefaults instantiates a new TagsOperationParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *TagsOperationParameters) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *TagsOperationParameters) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *TagsOperationParameters) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *TagsOperationParameters) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *TagsOperationParameters) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *TagsOperationParameters) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *TagsOperationParameters) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *TagsOperationParameters) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *TagsOperationParameters) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *TagsOperationParameters) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *TagsOperationParameters) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *TagsOperationParameters) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetDocumentIds

`func (o *TagsOperationParameters) GetDocumentIds() []string`

GetDocumentIds returns the DocumentIds field if non-nil, zero value otherwise.

### GetDocumentIdsOk

`func (o *TagsOperationParameters) GetDocumentIdsOk() (*[]string, bool)`

GetDocumentIdsOk returns a tuple with the DocumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIds

`func (o *TagsOperationParameters) SetDocumentIds(v []string)`

SetDocumentIds sets DocumentIds field to given value.

### HasDocumentIds

`func (o *TagsOperationParameters) HasDocumentIds() bool`

HasDocumentIds returns a boolean if a field has been set.

### SetDocumentIdsNil

`func (o *TagsOperationParameters) SetDocumentIdsNil(b bool)`

 SetDocumentIdsNil sets the value for DocumentIds to be an explicit nil

### UnsetDocumentIds
`func (o *TagsOperationParameters) UnsetDocumentIds()`

UnsetDocumentIds ensures that no value is present for DocumentIds, not even an explicit nil
### GetEntityId

`func (o *TagsOperationParameters) GetEntityId() int64`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *TagsOperationParameters) GetEntityIdOk() (*int64, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *TagsOperationParameters) SetEntityId(v int64)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *TagsOperationParameters) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *TagsOperationParameters) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *TagsOperationParameters) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetJobId

`func (o *TagsOperationParameters) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *TagsOperationParameters) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *TagsOperationParameters) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *TagsOperationParameters) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *TagsOperationParameters) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *TagsOperationParameters) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetJobInstanceIds

`func (o *TagsOperationParameters) GetJobInstanceIds() []int64`

GetJobInstanceIds returns the JobInstanceIds field if non-nil, zero value otherwise.

### GetJobInstanceIdsOk

`func (o *TagsOperationParameters) GetJobInstanceIdsOk() (*[]int64, bool)`

GetJobInstanceIdsOk returns a tuple with the JobInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobInstanceIds

`func (o *TagsOperationParameters) SetJobInstanceIds(v []int64)`

SetJobInstanceIds sets JobInstanceIds field to given value.

### HasJobInstanceIds

`func (o *TagsOperationParameters) HasJobInstanceIds() bool`

HasJobInstanceIds returns a boolean if a field has been set.

### SetJobInstanceIdsNil

`func (o *TagsOperationParameters) SetJobInstanceIdsNil(b bool)`

 SetJobInstanceIdsNil sets the value for JobInstanceIds to be an explicit nil

### UnsetJobInstanceIds
`func (o *TagsOperationParameters) UnsetJobInstanceIds()`

UnsetJobInstanceIds ensures that no value is present for JobInstanceIds, not even an explicit nil
### GetTags

`func (o *TagsOperationParameters) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TagsOperationParameters) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TagsOperationParameters) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TagsOperationParameters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *TagsOperationParameters) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *TagsOperationParameters) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


