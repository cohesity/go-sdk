# CommonDataAccessSessionInformation

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

## Methods

### NewCommonDataAccessSessionInformation

`func NewCommonDataAccessSessionInformation() *CommonDataAccessSessionInformation`

NewCommonDataAccessSessionInformation instantiates a new CommonDataAccessSessionInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonDataAccessSessionInformationWithDefaults

`func NewCommonDataAccessSessionInformationWithDefaults() *CommonDataAccessSessionInformation`

NewCommonDataAccessSessionInformationWithDefaults instantiates a new CommonDataAccessSessionInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTimeUsecs

`func (o *CommonDataAccessSessionInformation) GetCreationTimeUsecs() int64`

GetCreationTimeUsecs returns the CreationTimeUsecs field if non-nil, zero value otherwise.

### GetCreationTimeUsecsOk

`func (o *CommonDataAccessSessionInformation) GetCreationTimeUsecsOk() (*int64, bool)`

GetCreationTimeUsecsOk returns a tuple with the CreationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimeUsecs

`func (o *CommonDataAccessSessionInformation) SetCreationTimeUsecs(v int64)`

SetCreationTimeUsecs sets CreationTimeUsecs field to given value.

### HasCreationTimeUsecs

`func (o *CommonDataAccessSessionInformation) HasCreationTimeUsecs() bool`

HasCreationTimeUsecs returns a boolean if a field has been set.

### GetDataWorkerEndpoints

`func (o *CommonDataAccessSessionInformation) GetDataWorkerEndpoints() []WorkerEndpoint`

GetDataWorkerEndpoints returns the DataWorkerEndpoints field if non-nil, zero value otherwise.

### GetDataWorkerEndpointsOk

`func (o *CommonDataAccessSessionInformation) GetDataWorkerEndpointsOk() (*[]WorkerEndpoint, bool)`

GetDataWorkerEndpointsOk returns a tuple with the DataWorkerEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWorkerEndpoints

`func (o *CommonDataAccessSessionInformation) SetDataWorkerEndpoints(v []WorkerEndpoint)`

SetDataWorkerEndpoints sets DataWorkerEndpoints field to given value.

### HasDataWorkerEndpoints

`func (o *CommonDataAccessSessionInformation) HasDataWorkerEndpoints() bool`

HasDataWorkerEndpoints returns a boolean if a field has been set.

### SetDataWorkerEndpointsNil

`func (o *CommonDataAccessSessionInformation) SetDataWorkerEndpointsNil(b bool)`

 SetDataWorkerEndpointsNil sets the value for DataWorkerEndpoints to be an explicit nil

### UnsetDataWorkerEndpoints
`func (o *CommonDataAccessSessionInformation) UnsetDataWorkerEndpoints()`

UnsetDataWorkerEndpoints ensures that no value is present for DataWorkerEndpoints, not even an explicit nil
### GetLastModificationTimeUsecs

`func (o *CommonDataAccessSessionInformation) GetLastModificationTimeUsecs() int64`

GetLastModificationTimeUsecs returns the LastModificationTimeUsecs field if non-nil, zero value otherwise.

### GetLastModificationTimeUsecsOk

`func (o *CommonDataAccessSessionInformation) GetLastModificationTimeUsecsOk() (*int64, bool)`

GetLastModificationTimeUsecsOk returns a tuple with the LastModificationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTimeUsecs

`func (o *CommonDataAccessSessionInformation) SetLastModificationTimeUsecs(v int64)`

SetLastModificationTimeUsecs sets LastModificationTimeUsecs field to given value.

### HasLastModificationTimeUsecs

`func (o *CommonDataAccessSessionInformation) HasLastModificationTimeUsecs() bool`

HasLastModificationTimeUsecs returns a boolean if a field has been set.

### GetMetadataWorkerEndpoints

`func (o *CommonDataAccessSessionInformation) GetMetadataWorkerEndpoints() []WorkerEndpoint`

GetMetadataWorkerEndpoints returns the MetadataWorkerEndpoints field if non-nil, zero value otherwise.

### GetMetadataWorkerEndpointsOk

`func (o *CommonDataAccessSessionInformation) GetMetadataWorkerEndpointsOk() (*[]WorkerEndpoint, bool)`

GetMetadataWorkerEndpointsOk returns a tuple with the MetadataWorkerEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataWorkerEndpoints

`func (o *CommonDataAccessSessionInformation) SetMetadataWorkerEndpoints(v []WorkerEndpoint)`

SetMetadataWorkerEndpoints sets MetadataWorkerEndpoints field to given value.

### HasMetadataWorkerEndpoints

`func (o *CommonDataAccessSessionInformation) HasMetadataWorkerEndpoints() bool`

HasMetadataWorkerEndpoints returns a boolean if a field has been set.

### SetMetadataWorkerEndpointsNil

`func (o *CommonDataAccessSessionInformation) SetMetadataWorkerEndpointsNil(b bool)`

 SetMetadataWorkerEndpointsNil sets the value for MetadataWorkerEndpoints to be an explicit nil

### UnsetMetadataWorkerEndpoints
`func (o *CommonDataAccessSessionInformation) UnsetMetadataWorkerEndpoints()`

UnsetMetadataWorkerEndpoints ensures that no value is present for MetadataWorkerEndpoints, not even an explicit nil
### GetName

`func (o *CommonDataAccessSessionInformation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonDataAccessSessionInformation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonDataAccessSessionInformation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommonDataAccessSessionInformation) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CommonDataAccessSessionInformation) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommonDataAccessSessionInformation) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSessionId

`func (o *CommonDataAccessSessionInformation) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CommonDataAccessSessionInformation) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CommonDataAccessSessionInformation) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *CommonDataAccessSessionInformation) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *CommonDataAccessSessionInformation) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *CommonDataAccessSessionInformation) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetStatus

`func (o *CommonDataAccessSessionInformation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommonDataAccessSessionInformation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommonDataAccessSessionInformation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommonDataAccessSessionInformation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


