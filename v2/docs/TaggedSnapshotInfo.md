# TaggedSnapshotInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | Specifies the cluster Id of the tagged snapshot. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the clusterIncarnationId of the tagged snapshot. | [optional] 
**JobId** | Pointer to **NullableInt64** | Specifies the jobId of the tagged snapshot. | [optional] 
**ObjectUuid** | Pointer to **NullableString** | Specifies the object uuid of the tagged snapshot. | [optional] 
**RunStartTimeUsecs** | Pointer to **NullableInt64** | Specifies the runStartTimeUsecs of the tagged snapshot. | [optional] 
**Tags** | Pointer to [**[]HeliosTagInfo**](HeliosTagInfo.md) | Specifies tag applied to the object. | [optional] 

## Methods

### NewTaggedSnapshotInfo

`func NewTaggedSnapshotInfo() *TaggedSnapshotInfo`

NewTaggedSnapshotInfo instantiates a new TaggedSnapshotInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaggedSnapshotInfoWithDefaults

`func NewTaggedSnapshotInfoWithDefaults() *TaggedSnapshotInfo`

NewTaggedSnapshotInfoWithDefaults instantiates a new TaggedSnapshotInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *TaggedSnapshotInfo) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *TaggedSnapshotInfo) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *TaggedSnapshotInfo) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *TaggedSnapshotInfo) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *TaggedSnapshotInfo) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *TaggedSnapshotInfo) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *TaggedSnapshotInfo) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *TaggedSnapshotInfo) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *TaggedSnapshotInfo) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *TaggedSnapshotInfo) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *TaggedSnapshotInfo) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *TaggedSnapshotInfo) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetJobId

`func (o *TaggedSnapshotInfo) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *TaggedSnapshotInfo) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *TaggedSnapshotInfo) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *TaggedSnapshotInfo) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *TaggedSnapshotInfo) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *TaggedSnapshotInfo) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetObjectUuid

`func (o *TaggedSnapshotInfo) GetObjectUuid() string`

GetObjectUuid returns the ObjectUuid field if non-nil, zero value otherwise.

### GetObjectUuidOk

`func (o *TaggedSnapshotInfo) GetObjectUuidOk() (*string, bool)`

GetObjectUuidOk returns a tuple with the ObjectUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectUuid

`func (o *TaggedSnapshotInfo) SetObjectUuid(v string)`

SetObjectUuid sets ObjectUuid field to given value.

### HasObjectUuid

`func (o *TaggedSnapshotInfo) HasObjectUuid() bool`

HasObjectUuid returns a boolean if a field has been set.

### SetObjectUuidNil

`func (o *TaggedSnapshotInfo) SetObjectUuidNil(b bool)`

 SetObjectUuidNil sets the value for ObjectUuid to be an explicit nil

### UnsetObjectUuid
`func (o *TaggedSnapshotInfo) UnsetObjectUuid()`

UnsetObjectUuid ensures that no value is present for ObjectUuid, not even an explicit nil
### GetRunStartTimeUsecs

`func (o *TaggedSnapshotInfo) GetRunStartTimeUsecs() int64`

GetRunStartTimeUsecs returns the RunStartTimeUsecs field if non-nil, zero value otherwise.

### GetRunStartTimeUsecsOk

`func (o *TaggedSnapshotInfo) GetRunStartTimeUsecsOk() (*int64, bool)`

GetRunStartTimeUsecsOk returns a tuple with the RunStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStartTimeUsecs

`func (o *TaggedSnapshotInfo) SetRunStartTimeUsecs(v int64)`

SetRunStartTimeUsecs sets RunStartTimeUsecs field to given value.

### HasRunStartTimeUsecs

`func (o *TaggedSnapshotInfo) HasRunStartTimeUsecs() bool`

HasRunStartTimeUsecs returns a boolean if a field has been set.

### SetRunStartTimeUsecsNil

`func (o *TaggedSnapshotInfo) SetRunStartTimeUsecsNil(b bool)`

 SetRunStartTimeUsecsNil sets the value for RunStartTimeUsecs to be an explicit nil

### UnsetRunStartTimeUsecs
`func (o *TaggedSnapshotInfo) UnsetRunStartTimeUsecs()`

UnsetRunStartTimeUsecs ensures that no value is present for RunStartTimeUsecs, not even an explicit nil
### GetTags

`func (o *TaggedSnapshotInfo) GetTags() []HeliosTagInfo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TaggedSnapshotInfo) GetTagsOk() (*[]HeliosTagInfo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TaggedSnapshotInfo) SetTags(v []HeliosTagInfo)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TaggedSnapshotInfo) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *TaggedSnapshotInfo) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *TaggedSnapshotInfo) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


