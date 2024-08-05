# SnapshotDiffParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseSnapshotJobInstanceId** | **int64** |  | 
**BaseSnapshotTimeUsecs** | **int64** |  | 
**ClusterId** | **int64** |  | 
**EntityType** | **string** |  | 
**IncarnationId** | Pointer to **int64** |  | [optional] 
**JobId** | **int64** |  | 
**PageNumber** | **int64** |  | 
**PageSize** | **int64** |  | 
**PartitionId** | **int64** |  | 
**SnapshotJobInstanceId** | **int64** |  | 
**SnapshotTimeUsecs** | **int64** |  | 

## Methods

### NewSnapshotDiffParams

`func NewSnapshotDiffParams(baseSnapshotJobInstanceId int64, baseSnapshotTimeUsecs int64, clusterId int64, entityType string, jobId int64, pageNumber int64, pageSize int64, partitionId int64, snapshotJobInstanceId int64, snapshotTimeUsecs int64, ) *SnapshotDiffParams`

NewSnapshotDiffParams instantiates a new SnapshotDiffParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotDiffParamsWithDefaults

`func NewSnapshotDiffParamsWithDefaults() *SnapshotDiffParams`

NewSnapshotDiffParamsWithDefaults instantiates a new SnapshotDiffParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseSnapshotJobInstanceId

`func (o *SnapshotDiffParams) GetBaseSnapshotJobInstanceId() int64`

GetBaseSnapshotJobInstanceId returns the BaseSnapshotJobInstanceId field if non-nil, zero value otherwise.

### GetBaseSnapshotJobInstanceIdOk

`func (o *SnapshotDiffParams) GetBaseSnapshotJobInstanceIdOk() (*int64, bool)`

GetBaseSnapshotJobInstanceIdOk returns a tuple with the BaseSnapshotJobInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSnapshotJobInstanceId

`func (o *SnapshotDiffParams) SetBaseSnapshotJobInstanceId(v int64)`

SetBaseSnapshotJobInstanceId sets BaseSnapshotJobInstanceId field to given value.


### GetBaseSnapshotTimeUsecs

`func (o *SnapshotDiffParams) GetBaseSnapshotTimeUsecs() int64`

GetBaseSnapshotTimeUsecs returns the BaseSnapshotTimeUsecs field if non-nil, zero value otherwise.

### GetBaseSnapshotTimeUsecsOk

`func (o *SnapshotDiffParams) GetBaseSnapshotTimeUsecsOk() (*int64, bool)`

GetBaseSnapshotTimeUsecsOk returns a tuple with the BaseSnapshotTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSnapshotTimeUsecs

`func (o *SnapshotDiffParams) SetBaseSnapshotTimeUsecs(v int64)`

SetBaseSnapshotTimeUsecs sets BaseSnapshotTimeUsecs field to given value.


### GetClusterId

`func (o *SnapshotDiffParams) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *SnapshotDiffParams) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *SnapshotDiffParams) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.


### GetEntityType

`func (o *SnapshotDiffParams) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *SnapshotDiffParams) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *SnapshotDiffParams) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetIncarnationId

`func (o *SnapshotDiffParams) GetIncarnationId() int64`

GetIncarnationId returns the IncarnationId field if non-nil, zero value otherwise.

### GetIncarnationIdOk

`func (o *SnapshotDiffParams) GetIncarnationIdOk() (*int64, bool)`

GetIncarnationIdOk returns a tuple with the IncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncarnationId

`func (o *SnapshotDiffParams) SetIncarnationId(v int64)`

SetIncarnationId sets IncarnationId field to given value.

### HasIncarnationId

`func (o *SnapshotDiffParams) HasIncarnationId() bool`

HasIncarnationId returns a boolean if a field has been set.

### GetJobId

`func (o *SnapshotDiffParams) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *SnapshotDiffParams) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *SnapshotDiffParams) SetJobId(v int64)`

SetJobId sets JobId field to given value.


### GetPageNumber

`func (o *SnapshotDiffParams) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *SnapshotDiffParams) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *SnapshotDiffParams) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *SnapshotDiffParams) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *SnapshotDiffParams) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *SnapshotDiffParams) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetPartitionId

`func (o *SnapshotDiffParams) GetPartitionId() int64`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *SnapshotDiffParams) GetPartitionIdOk() (*int64, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *SnapshotDiffParams) SetPartitionId(v int64)`

SetPartitionId sets PartitionId field to given value.


### GetSnapshotJobInstanceId

`func (o *SnapshotDiffParams) GetSnapshotJobInstanceId() int64`

GetSnapshotJobInstanceId returns the SnapshotJobInstanceId field if non-nil, zero value otherwise.

### GetSnapshotJobInstanceIdOk

`func (o *SnapshotDiffParams) GetSnapshotJobInstanceIdOk() (*int64, bool)`

GetSnapshotJobInstanceIdOk returns a tuple with the SnapshotJobInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotJobInstanceId

`func (o *SnapshotDiffParams) SetSnapshotJobInstanceId(v int64)`

SetSnapshotJobInstanceId sets SnapshotJobInstanceId field to given value.


### GetSnapshotTimeUsecs

`func (o *SnapshotDiffParams) GetSnapshotTimeUsecs() int64`

GetSnapshotTimeUsecs returns the SnapshotTimeUsecs field if non-nil, zero value otherwise.

### GetSnapshotTimeUsecsOk

`func (o *SnapshotDiffParams) GetSnapshotTimeUsecsOk() (*int64, bool)`

GetSnapshotTimeUsecsOk returns a tuple with the SnapshotTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTimeUsecs

`func (o *SnapshotDiffParams) SetSnapshotTimeUsecs(v int64)`

SetSnapshotTimeUsecs sets SnapshotTimeUsecs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


