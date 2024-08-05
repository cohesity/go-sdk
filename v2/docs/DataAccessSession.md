# DataAccessSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTimeUsecs** | Pointer to **int64** | Specifies the time at which the session was created. | [optional] [readonly] 
**DataWorkerEndpoints** | Pointer to [**[]WorkerEndpoint**](WorkerEndpoint.md) | Specifies the list of data worker endpoints. For load balancing client can make rpc call to different data worker for different nodes. Client should make the rpc call to same worker for data for a given node. | [optional] 
**LastModificationTimeUsecs** | Pointer to **int64** | Specifies the time at which the session was last modified. | [optional] [readonly] 
**MetadataWorkerEndpoints** | Pointer to [**[]WorkerEndpoint**](WorkerEndpoint.md) | Specifies the list of metadata worker endpoints. In case of more than one metadata point client can contact any metadata worker. | [optional] 
**Name** | Pointer to **NullableString** | The name of the data access session. | [optional] 
**SessionId** | Pointer to **NullableString** | Specifies the id of the data access session. | [optional] 
**Status** | Pointer to **string** | Specifies the status of the Data Access Session. Machine status such as Admitted/WaitingForArchiveDownload/ WaitingForResource/SetupInProgress/Ready/Finished | [optional] 
**BaseSnapshotInfo** | Pointer to **map[string]interface{}** | Specifies information about the base snapshot of an object. | [optional] 
**CurrentSnapshotInfo** | Pointer to **map[string]interface{}** | Specifies information about the backup snapshot and object whose data/metadata needs to be accessed. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies registered source id. | [optional] 

## Methods

### NewDataAccessSession

`func NewDataAccessSession() *DataAccessSession`

NewDataAccessSession instantiates a new DataAccessSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataAccessSessionWithDefaults

`func NewDataAccessSessionWithDefaults() *DataAccessSession`

NewDataAccessSessionWithDefaults instantiates a new DataAccessSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTimeUsecs

`func (o *DataAccessSession) GetCreationTimeUsecs() int64`

GetCreationTimeUsecs returns the CreationTimeUsecs field if non-nil, zero value otherwise.

### GetCreationTimeUsecsOk

`func (o *DataAccessSession) GetCreationTimeUsecsOk() (*int64, bool)`

GetCreationTimeUsecsOk returns a tuple with the CreationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimeUsecs

`func (o *DataAccessSession) SetCreationTimeUsecs(v int64)`

SetCreationTimeUsecs sets CreationTimeUsecs field to given value.

### HasCreationTimeUsecs

`func (o *DataAccessSession) HasCreationTimeUsecs() bool`

HasCreationTimeUsecs returns a boolean if a field has been set.

### GetDataWorkerEndpoints

`func (o *DataAccessSession) GetDataWorkerEndpoints() []WorkerEndpoint`

GetDataWorkerEndpoints returns the DataWorkerEndpoints field if non-nil, zero value otherwise.

### GetDataWorkerEndpointsOk

`func (o *DataAccessSession) GetDataWorkerEndpointsOk() (*[]WorkerEndpoint, bool)`

GetDataWorkerEndpointsOk returns a tuple with the DataWorkerEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWorkerEndpoints

`func (o *DataAccessSession) SetDataWorkerEndpoints(v []WorkerEndpoint)`

SetDataWorkerEndpoints sets DataWorkerEndpoints field to given value.

### HasDataWorkerEndpoints

`func (o *DataAccessSession) HasDataWorkerEndpoints() bool`

HasDataWorkerEndpoints returns a boolean if a field has been set.

### SetDataWorkerEndpointsNil

`func (o *DataAccessSession) SetDataWorkerEndpointsNil(b bool)`

 SetDataWorkerEndpointsNil sets the value for DataWorkerEndpoints to be an explicit nil

### UnsetDataWorkerEndpoints
`func (o *DataAccessSession) UnsetDataWorkerEndpoints()`

UnsetDataWorkerEndpoints ensures that no value is present for DataWorkerEndpoints, not even an explicit nil
### GetLastModificationTimeUsecs

`func (o *DataAccessSession) GetLastModificationTimeUsecs() int64`

GetLastModificationTimeUsecs returns the LastModificationTimeUsecs field if non-nil, zero value otherwise.

### GetLastModificationTimeUsecsOk

`func (o *DataAccessSession) GetLastModificationTimeUsecsOk() (*int64, bool)`

GetLastModificationTimeUsecsOk returns a tuple with the LastModificationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTimeUsecs

`func (o *DataAccessSession) SetLastModificationTimeUsecs(v int64)`

SetLastModificationTimeUsecs sets LastModificationTimeUsecs field to given value.

### HasLastModificationTimeUsecs

`func (o *DataAccessSession) HasLastModificationTimeUsecs() bool`

HasLastModificationTimeUsecs returns a boolean if a field has been set.

### GetMetadataWorkerEndpoints

`func (o *DataAccessSession) GetMetadataWorkerEndpoints() []WorkerEndpoint`

GetMetadataWorkerEndpoints returns the MetadataWorkerEndpoints field if non-nil, zero value otherwise.

### GetMetadataWorkerEndpointsOk

`func (o *DataAccessSession) GetMetadataWorkerEndpointsOk() (*[]WorkerEndpoint, bool)`

GetMetadataWorkerEndpointsOk returns a tuple with the MetadataWorkerEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataWorkerEndpoints

`func (o *DataAccessSession) SetMetadataWorkerEndpoints(v []WorkerEndpoint)`

SetMetadataWorkerEndpoints sets MetadataWorkerEndpoints field to given value.

### HasMetadataWorkerEndpoints

`func (o *DataAccessSession) HasMetadataWorkerEndpoints() bool`

HasMetadataWorkerEndpoints returns a boolean if a field has been set.

### SetMetadataWorkerEndpointsNil

`func (o *DataAccessSession) SetMetadataWorkerEndpointsNil(b bool)`

 SetMetadataWorkerEndpointsNil sets the value for MetadataWorkerEndpoints to be an explicit nil

### UnsetMetadataWorkerEndpoints
`func (o *DataAccessSession) UnsetMetadataWorkerEndpoints()`

UnsetMetadataWorkerEndpoints ensures that no value is present for MetadataWorkerEndpoints, not even an explicit nil
### GetName

`func (o *DataAccessSession) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataAccessSession) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataAccessSession) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataAccessSession) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DataAccessSession) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DataAccessSession) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSessionId

`func (o *DataAccessSession) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *DataAccessSession) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *DataAccessSession) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *DataAccessSession) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *DataAccessSession) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *DataAccessSession) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetStatus

`func (o *DataAccessSession) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DataAccessSession) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DataAccessSession) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DataAccessSession) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBaseSnapshotInfo

`func (o *DataAccessSession) GetBaseSnapshotInfo() map[string]interface{}`

GetBaseSnapshotInfo returns the BaseSnapshotInfo field if non-nil, zero value otherwise.

### GetBaseSnapshotInfoOk

`func (o *DataAccessSession) GetBaseSnapshotInfoOk() (*map[string]interface{}, bool)`

GetBaseSnapshotInfoOk returns a tuple with the BaseSnapshotInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSnapshotInfo

`func (o *DataAccessSession) SetBaseSnapshotInfo(v map[string]interface{})`

SetBaseSnapshotInfo sets BaseSnapshotInfo field to given value.

### HasBaseSnapshotInfo

`func (o *DataAccessSession) HasBaseSnapshotInfo() bool`

HasBaseSnapshotInfo returns a boolean if a field has been set.

### SetBaseSnapshotInfoNil

`func (o *DataAccessSession) SetBaseSnapshotInfoNil(b bool)`

 SetBaseSnapshotInfoNil sets the value for BaseSnapshotInfo to be an explicit nil

### UnsetBaseSnapshotInfo
`func (o *DataAccessSession) UnsetBaseSnapshotInfo()`

UnsetBaseSnapshotInfo ensures that no value is present for BaseSnapshotInfo, not even an explicit nil
### GetCurrentSnapshotInfo

`func (o *DataAccessSession) GetCurrentSnapshotInfo() map[string]interface{}`

GetCurrentSnapshotInfo returns the CurrentSnapshotInfo field if non-nil, zero value otherwise.

### GetCurrentSnapshotInfoOk

`func (o *DataAccessSession) GetCurrentSnapshotInfoOk() (*map[string]interface{}, bool)`

GetCurrentSnapshotInfoOk returns a tuple with the CurrentSnapshotInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSnapshotInfo

`func (o *DataAccessSession) SetCurrentSnapshotInfo(v map[string]interface{})`

SetCurrentSnapshotInfo sets CurrentSnapshotInfo field to given value.

### HasCurrentSnapshotInfo

`func (o *DataAccessSession) HasCurrentSnapshotInfo() bool`

HasCurrentSnapshotInfo returns a boolean if a field has been set.

### SetCurrentSnapshotInfoNil

`func (o *DataAccessSession) SetCurrentSnapshotInfoNil(b bool)`

 SetCurrentSnapshotInfoNil sets the value for CurrentSnapshotInfo to be an explicit nil

### UnsetCurrentSnapshotInfo
`func (o *DataAccessSession) UnsetCurrentSnapshotInfo()`

UnsetCurrentSnapshotInfo ensures that no value is present for CurrentSnapshotInfo, not even an explicit nil
### GetSourceId

`func (o *DataAccessSession) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *DataAccessSession) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *DataAccessSession) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *DataAccessSession) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *DataAccessSession) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *DataAccessSession) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


