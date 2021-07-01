# CloneRefreshRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloneTaskId** | Pointer to **NullableInt64** | Specifies the ID of the clone task. This is required to determine the details of the clone to be refreshed as clone task contains the details of the clone. | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies if the Restore Task should continue when some operations on some objects fail. If true, the Cohesity Cluster ignores intermittent errors and restores as many objects as possible. | [optional] 
**Name** | **NullableString** | Specifies the name of the Restore Task. This field must be set and must be a unique name. | 
**NewParentId** | Pointer to **NullableInt64** | Specify a new registered parent Protection Source. If specified the selected objects are cloned or recovered to this new Protection Source. If not specified, objects are cloned or recovered to the original Protection Source that was managing them. | [optional] 
**Objects** | Pointer to [**[]RestoreObjectDetails**](RestoreObjectDetails.md) | Array of Objects.  Specifies a list of Protection Source objects or Protection Job objects (with specified Protection Source objects). | [optional] 
**RefreshTimeSecs** | Pointer to **NullableInt64** | Specifies a point in time (unix epoch) to which the database needs to be refreshed. This helps granular refresh of the database. If this is set, relevant archive logs (redo logs) will also be re-played to match with the specified time. For this, the log backup should be enabled in the backup policy. If this is not set, then only the incremental backup data will be used to refresh the target database. | [optional] 
**SourceDatabaseId** | Pointer to **NullableInt64** | Specifies the ID of the source database in the backup job snapshot. This is the entity ID of the database, which needs to be used as a source during the refresh process. | [optional] 
**VlanParameters** | Pointer to [**VlanParameters**](VlanParameters.md) |  | [optional] 

## Methods

### NewCloneRefreshRequest

`func NewCloneRefreshRequest(name NullableString, ) *CloneRefreshRequest`

NewCloneRefreshRequest instantiates a new CloneRefreshRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneRefreshRequestWithDefaults

`func NewCloneRefreshRequestWithDefaults() *CloneRefreshRequest`

NewCloneRefreshRequestWithDefaults instantiates a new CloneRefreshRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloneTaskId

`func (o *CloneRefreshRequest) GetCloneTaskId() int64`

GetCloneTaskId returns the CloneTaskId field if non-nil, zero value otherwise.

### GetCloneTaskIdOk

`func (o *CloneRefreshRequest) GetCloneTaskIdOk() (*int64, bool)`

GetCloneTaskIdOk returns a tuple with the CloneTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneTaskId

`func (o *CloneRefreshRequest) SetCloneTaskId(v int64)`

SetCloneTaskId sets CloneTaskId field to given value.

### HasCloneTaskId

`func (o *CloneRefreshRequest) HasCloneTaskId() bool`

HasCloneTaskId returns a boolean if a field has been set.

### SetCloneTaskIdNil

`func (o *CloneRefreshRequest) SetCloneTaskIdNil(b bool)`

 SetCloneTaskIdNil sets the value for CloneTaskId to be an explicit nil

### UnsetCloneTaskId
`func (o *CloneRefreshRequest) UnsetCloneTaskId()`

UnsetCloneTaskId ensures that no value is present for CloneTaskId, not even an explicit nil
### GetContinueOnError

`func (o *CloneRefreshRequest) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *CloneRefreshRequest) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *CloneRefreshRequest) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *CloneRefreshRequest) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *CloneRefreshRequest) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *CloneRefreshRequest) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetName

`func (o *CloneRefreshRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloneRefreshRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloneRefreshRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CloneRefreshRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CloneRefreshRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNewParentId

`func (o *CloneRefreshRequest) GetNewParentId() int64`

GetNewParentId returns the NewParentId field if non-nil, zero value otherwise.

### GetNewParentIdOk

`func (o *CloneRefreshRequest) GetNewParentIdOk() (*int64, bool)`

GetNewParentIdOk returns a tuple with the NewParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewParentId

`func (o *CloneRefreshRequest) SetNewParentId(v int64)`

SetNewParentId sets NewParentId field to given value.

### HasNewParentId

`func (o *CloneRefreshRequest) HasNewParentId() bool`

HasNewParentId returns a boolean if a field has been set.

### SetNewParentIdNil

`func (o *CloneRefreshRequest) SetNewParentIdNil(b bool)`

 SetNewParentIdNil sets the value for NewParentId to be an explicit nil

### UnsetNewParentId
`func (o *CloneRefreshRequest) UnsetNewParentId()`

UnsetNewParentId ensures that no value is present for NewParentId, not even an explicit nil
### GetObjects

`func (o *CloneRefreshRequest) GetObjects() []RestoreObjectDetails`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *CloneRefreshRequest) GetObjectsOk() (*[]RestoreObjectDetails, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *CloneRefreshRequest) SetObjects(v []RestoreObjectDetails)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *CloneRefreshRequest) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *CloneRefreshRequest) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *CloneRefreshRequest) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRefreshTimeSecs

`func (o *CloneRefreshRequest) GetRefreshTimeSecs() int64`

GetRefreshTimeSecs returns the RefreshTimeSecs field if non-nil, zero value otherwise.

### GetRefreshTimeSecsOk

`func (o *CloneRefreshRequest) GetRefreshTimeSecsOk() (*int64, bool)`

GetRefreshTimeSecsOk returns a tuple with the RefreshTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTimeSecs

`func (o *CloneRefreshRequest) SetRefreshTimeSecs(v int64)`

SetRefreshTimeSecs sets RefreshTimeSecs field to given value.

### HasRefreshTimeSecs

`func (o *CloneRefreshRequest) HasRefreshTimeSecs() bool`

HasRefreshTimeSecs returns a boolean if a field has been set.

### SetRefreshTimeSecsNil

`func (o *CloneRefreshRequest) SetRefreshTimeSecsNil(b bool)`

 SetRefreshTimeSecsNil sets the value for RefreshTimeSecs to be an explicit nil

### UnsetRefreshTimeSecs
`func (o *CloneRefreshRequest) UnsetRefreshTimeSecs()`

UnsetRefreshTimeSecs ensures that no value is present for RefreshTimeSecs, not even an explicit nil
### GetSourceDatabaseId

`func (o *CloneRefreshRequest) GetSourceDatabaseId() int64`

GetSourceDatabaseId returns the SourceDatabaseId field if non-nil, zero value otherwise.

### GetSourceDatabaseIdOk

`func (o *CloneRefreshRequest) GetSourceDatabaseIdOk() (*int64, bool)`

GetSourceDatabaseIdOk returns a tuple with the SourceDatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDatabaseId

`func (o *CloneRefreshRequest) SetSourceDatabaseId(v int64)`

SetSourceDatabaseId sets SourceDatabaseId field to given value.

### HasSourceDatabaseId

`func (o *CloneRefreshRequest) HasSourceDatabaseId() bool`

HasSourceDatabaseId returns a boolean if a field has been set.

### SetSourceDatabaseIdNil

`func (o *CloneRefreshRequest) SetSourceDatabaseIdNil(b bool)`

 SetSourceDatabaseIdNil sets the value for SourceDatabaseId to be an explicit nil

### UnsetSourceDatabaseId
`func (o *CloneRefreshRequest) UnsetSourceDatabaseId()`

UnsetSourceDatabaseId ensures that no value is present for SourceDatabaseId, not even an explicit nil
### GetVlanParameters

`func (o *CloneRefreshRequest) GetVlanParameters() VlanParameters`

GetVlanParameters returns the VlanParameters field if non-nil, zero value otherwise.

### GetVlanParametersOk

`func (o *CloneRefreshRequest) GetVlanParametersOk() (*VlanParameters, bool)`

GetVlanParametersOk returns a tuple with the VlanParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanParameters

`func (o *CloneRefreshRequest) SetVlanParameters(v VlanParameters)`

SetVlanParameters sets VlanParameters field to given value.

### HasVlanParameters

`func (o *CloneRefreshRequest) HasVlanParameters() bool`

HasVlanParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


