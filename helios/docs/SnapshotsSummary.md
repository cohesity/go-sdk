# SnapshotsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | Specifies the cluster id where the snapshots is stored. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the cluster incarnation id where the snapshots is stored. | [optional] 
**RegionId** | Pointer to **NullableString** | Specifies the cluster indentifier where the snapshots is stored. | [optional] 
**SnapshotTargetType** | Pointer to **NullableString** | Specifies the target type where the Object&#39;s snapshot resides. | [optional] 
**ExternalTargetInfo** | Pointer to [**NullableArchivalTargetSummaryInfo**](ArchivalTargetSummaryInfo.md) | Specifies the external target information if this is an archival snapshot. | [optional] 
**SnapshotCount** | Pointer to **NullableInt64** | Specifies the number of snapshots of this type and target. | [optional] 
**LatestSnapshotTimestampUsecs** | Pointer to **NullableInt64** | Specifies the timestamp in Unix time epoch in microseconds when the latest snapshot is taken. | [optional] 
**LatestRunStartTimeUsecs** | Pointer to **NullableInt64** | Specifies the timestamp in Unix time epoch in microseconds when the latest run started. | [optional] 
**LatestRunStatus** | Pointer to **NullableString** | Specifies the status of latest run. | [optional] 

## Methods

### NewSnapshotsSummary

`func NewSnapshotsSummary() *SnapshotsSummary`

NewSnapshotsSummary instantiates a new SnapshotsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotsSummaryWithDefaults

`func NewSnapshotsSummaryWithDefaults() *SnapshotsSummary`

NewSnapshotsSummaryWithDefaults instantiates a new SnapshotsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *SnapshotsSummary) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *SnapshotsSummary) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *SnapshotsSummary) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *SnapshotsSummary) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *SnapshotsSummary) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *SnapshotsSummary) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *SnapshotsSummary) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *SnapshotsSummary) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *SnapshotsSummary) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *SnapshotsSummary) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *SnapshotsSummary) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *SnapshotsSummary) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetRegionId

`func (o *SnapshotsSummary) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *SnapshotsSummary) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *SnapshotsSummary) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *SnapshotsSummary) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *SnapshotsSummary) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *SnapshotsSummary) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetSnapshotTargetType

`func (o *SnapshotsSummary) GetSnapshotTargetType() string`

GetSnapshotTargetType returns the SnapshotTargetType field if non-nil, zero value otherwise.

### GetSnapshotTargetTypeOk

`func (o *SnapshotsSummary) GetSnapshotTargetTypeOk() (*string, bool)`

GetSnapshotTargetTypeOk returns a tuple with the SnapshotTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTargetType

`func (o *SnapshotsSummary) SetSnapshotTargetType(v string)`

SetSnapshotTargetType sets SnapshotTargetType field to given value.

### HasSnapshotTargetType

`func (o *SnapshotsSummary) HasSnapshotTargetType() bool`

HasSnapshotTargetType returns a boolean if a field has been set.

### SetSnapshotTargetTypeNil

`func (o *SnapshotsSummary) SetSnapshotTargetTypeNil(b bool)`

 SetSnapshotTargetTypeNil sets the value for SnapshotTargetType to be an explicit nil

### UnsetSnapshotTargetType
`func (o *SnapshotsSummary) UnsetSnapshotTargetType()`

UnsetSnapshotTargetType ensures that no value is present for SnapshotTargetType, not even an explicit nil
### GetExternalTargetInfo

`func (o *SnapshotsSummary) GetExternalTargetInfo() ArchivalTargetSummaryInfo`

GetExternalTargetInfo returns the ExternalTargetInfo field if non-nil, zero value otherwise.

### GetExternalTargetInfoOk

`func (o *SnapshotsSummary) GetExternalTargetInfoOk() (*ArchivalTargetSummaryInfo, bool)`

GetExternalTargetInfoOk returns a tuple with the ExternalTargetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTargetInfo

`func (o *SnapshotsSummary) SetExternalTargetInfo(v ArchivalTargetSummaryInfo)`

SetExternalTargetInfo sets ExternalTargetInfo field to given value.

### HasExternalTargetInfo

`func (o *SnapshotsSummary) HasExternalTargetInfo() bool`

HasExternalTargetInfo returns a boolean if a field has been set.

### SetExternalTargetInfoNil

`func (o *SnapshotsSummary) SetExternalTargetInfoNil(b bool)`

 SetExternalTargetInfoNil sets the value for ExternalTargetInfo to be an explicit nil

### UnsetExternalTargetInfo
`func (o *SnapshotsSummary) UnsetExternalTargetInfo()`

UnsetExternalTargetInfo ensures that no value is present for ExternalTargetInfo, not even an explicit nil
### GetSnapshotCount

`func (o *SnapshotsSummary) GetSnapshotCount() int64`

GetSnapshotCount returns the SnapshotCount field if non-nil, zero value otherwise.

### GetSnapshotCountOk

`func (o *SnapshotsSummary) GetSnapshotCountOk() (*int64, bool)`

GetSnapshotCountOk returns a tuple with the SnapshotCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotCount

`func (o *SnapshotsSummary) SetSnapshotCount(v int64)`

SetSnapshotCount sets SnapshotCount field to given value.

### HasSnapshotCount

`func (o *SnapshotsSummary) HasSnapshotCount() bool`

HasSnapshotCount returns a boolean if a field has been set.

### SetSnapshotCountNil

`func (o *SnapshotsSummary) SetSnapshotCountNil(b bool)`

 SetSnapshotCountNil sets the value for SnapshotCount to be an explicit nil

### UnsetSnapshotCount
`func (o *SnapshotsSummary) UnsetSnapshotCount()`

UnsetSnapshotCount ensures that no value is present for SnapshotCount, not even an explicit nil
### GetLatestSnapshotTimestampUsecs

`func (o *SnapshotsSummary) GetLatestSnapshotTimestampUsecs() int64`

GetLatestSnapshotTimestampUsecs returns the LatestSnapshotTimestampUsecs field if non-nil, zero value otherwise.

### GetLatestSnapshotTimestampUsecsOk

`func (o *SnapshotsSummary) GetLatestSnapshotTimestampUsecsOk() (*int64, bool)`

GetLatestSnapshotTimestampUsecsOk returns a tuple with the LatestSnapshotTimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSnapshotTimestampUsecs

`func (o *SnapshotsSummary) SetLatestSnapshotTimestampUsecs(v int64)`

SetLatestSnapshotTimestampUsecs sets LatestSnapshotTimestampUsecs field to given value.

### HasLatestSnapshotTimestampUsecs

`func (o *SnapshotsSummary) HasLatestSnapshotTimestampUsecs() bool`

HasLatestSnapshotTimestampUsecs returns a boolean if a field has been set.

### SetLatestSnapshotTimestampUsecsNil

`func (o *SnapshotsSummary) SetLatestSnapshotTimestampUsecsNil(b bool)`

 SetLatestSnapshotTimestampUsecsNil sets the value for LatestSnapshotTimestampUsecs to be an explicit nil

### UnsetLatestSnapshotTimestampUsecs
`func (o *SnapshotsSummary) UnsetLatestSnapshotTimestampUsecs()`

UnsetLatestSnapshotTimestampUsecs ensures that no value is present for LatestSnapshotTimestampUsecs, not even an explicit nil
### GetLatestRunStartTimeUsecs

`func (o *SnapshotsSummary) GetLatestRunStartTimeUsecs() int64`

GetLatestRunStartTimeUsecs returns the LatestRunStartTimeUsecs field if non-nil, zero value otherwise.

### GetLatestRunStartTimeUsecsOk

`func (o *SnapshotsSummary) GetLatestRunStartTimeUsecsOk() (*int64, bool)`

GetLatestRunStartTimeUsecsOk returns a tuple with the LatestRunStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRunStartTimeUsecs

`func (o *SnapshotsSummary) SetLatestRunStartTimeUsecs(v int64)`

SetLatestRunStartTimeUsecs sets LatestRunStartTimeUsecs field to given value.

### HasLatestRunStartTimeUsecs

`func (o *SnapshotsSummary) HasLatestRunStartTimeUsecs() bool`

HasLatestRunStartTimeUsecs returns a boolean if a field has been set.

### SetLatestRunStartTimeUsecsNil

`func (o *SnapshotsSummary) SetLatestRunStartTimeUsecsNil(b bool)`

 SetLatestRunStartTimeUsecsNil sets the value for LatestRunStartTimeUsecs to be an explicit nil

### UnsetLatestRunStartTimeUsecs
`func (o *SnapshotsSummary) UnsetLatestRunStartTimeUsecs()`

UnsetLatestRunStartTimeUsecs ensures that no value is present for LatestRunStartTimeUsecs, not even an explicit nil
### GetLatestRunStatus

`func (o *SnapshotsSummary) GetLatestRunStatus() string`

GetLatestRunStatus returns the LatestRunStatus field if non-nil, zero value otherwise.

### GetLatestRunStatusOk

`func (o *SnapshotsSummary) GetLatestRunStatusOk() (*string, bool)`

GetLatestRunStatusOk returns a tuple with the LatestRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRunStatus

`func (o *SnapshotsSummary) SetLatestRunStatus(v string)`

SetLatestRunStatus sets LatestRunStatus field to given value.

### HasLatestRunStatus

`func (o *SnapshotsSummary) HasLatestRunStatus() bool`

HasLatestRunStatus returns a boolean if a field has been set.

### SetLatestRunStatusNil

`func (o *SnapshotsSummary) SetLatestRunStatusNil(b bool)`

 SetLatestRunStatusNil sets the value for LatestRunStatus to be an explicit nil

### UnsetLatestRunStatus
`func (o *SnapshotsSummary) UnsetLatestRunStatus()`

UnsetLatestRunStatus ensures that no value is present for LatestRunStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


