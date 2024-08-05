# CreateDataAccessSessionRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseSnapshotInfo** | Pointer to [**NullableDataAccessSnapshotInfo**](DataAccessSnapshotInfo.md) | Specifies information about the base snapshot of an object. | [optional] 
**CurrentSnapshotInfo** | [**DataAccessSnapshotInfo**](DataAccessSnapshotInfo.md) | Specifies information about the backup snapshot and object whose data/metadata needs to be accessed. | 
**SourceId** | **int64** | Specifies the entity id of the source. | 

## Methods

### NewCreateDataAccessSessionRequestParams

`func NewCreateDataAccessSessionRequestParams(currentSnapshotInfo DataAccessSnapshotInfo, sourceId int64, ) *CreateDataAccessSessionRequestParams`

NewCreateDataAccessSessionRequestParams instantiates a new CreateDataAccessSessionRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDataAccessSessionRequestParamsWithDefaults

`func NewCreateDataAccessSessionRequestParamsWithDefaults() *CreateDataAccessSessionRequestParams`

NewCreateDataAccessSessionRequestParamsWithDefaults instantiates a new CreateDataAccessSessionRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseSnapshotInfo

`func (o *CreateDataAccessSessionRequestParams) GetBaseSnapshotInfo() DataAccessSnapshotInfo`

GetBaseSnapshotInfo returns the BaseSnapshotInfo field if non-nil, zero value otherwise.

### GetBaseSnapshotInfoOk

`func (o *CreateDataAccessSessionRequestParams) GetBaseSnapshotInfoOk() (*DataAccessSnapshotInfo, bool)`

GetBaseSnapshotInfoOk returns a tuple with the BaseSnapshotInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSnapshotInfo

`func (o *CreateDataAccessSessionRequestParams) SetBaseSnapshotInfo(v DataAccessSnapshotInfo)`

SetBaseSnapshotInfo sets BaseSnapshotInfo field to given value.

### HasBaseSnapshotInfo

`func (o *CreateDataAccessSessionRequestParams) HasBaseSnapshotInfo() bool`

HasBaseSnapshotInfo returns a boolean if a field has been set.

### SetBaseSnapshotInfoNil

`func (o *CreateDataAccessSessionRequestParams) SetBaseSnapshotInfoNil(b bool)`

 SetBaseSnapshotInfoNil sets the value for BaseSnapshotInfo to be an explicit nil

### UnsetBaseSnapshotInfo
`func (o *CreateDataAccessSessionRequestParams) UnsetBaseSnapshotInfo()`

UnsetBaseSnapshotInfo ensures that no value is present for BaseSnapshotInfo, not even an explicit nil
### GetCurrentSnapshotInfo

`func (o *CreateDataAccessSessionRequestParams) GetCurrentSnapshotInfo() DataAccessSnapshotInfo`

GetCurrentSnapshotInfo returns the CurrentSnapshotInfo field if non-nil, zero value otherwise.

### GetCurrentSnapshotInfoOk

`func (o *CreateDataAccessSessionRequestParams) GetCurrentSnapshotInfoOk() (*DataAccessSnapshotInfo, bool)`

GetCurrentSnapshotInfoOk returns a tuple with the CurrentSnapshotInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSnapshotInfo

`func (o *CreateDataAccessSessionRequestParams) SetCurrentSnapshotInfo(v DataAccessSnapshotInfo)`

SetCurrentSnapshotInfo sets CurrentSnapshotInfo field to given value.


### GetSourceId

`func (o *CreateDataAccessSessionRequestParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *CreateDataAccessSessionRequestParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *CreateDataAccessSessionRequestParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


