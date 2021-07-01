# MagnetoObjectId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**JobId** | Pointer to **NullableInt64** | The id of the local job that the object belongs to, which may or may not match the object_id field in job_uid below depending on whether the object originally belonged to this local job or to a different remote job. | [optional] 
**JobUid** | Pointer to [**UniversalIdProto**](UniversalIdProto.md) |  | [optional] 

## Methods

### NewMagnetoObjectId

`func NewMagnetoObjectId() *MagnetoObjectId`

NewMagnetoObjectId instantiates a new MagnetoObjectId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMagnetoObjectIdWithDefaults

`func NewMagnetoObjectIdWithDefaults() *MagnetoObjectId`

NewMagnetoObjectIdWithDefaults instantiates a new MagnetoObjectId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *MagnetoObjectId) GetEntity() EntityProto`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *MagnetoObjectId) GetEntityOk() (*EntityProto, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *MagnetoObjectId) SetEntity(v EntityProto)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *MagnetoObjectId) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetJobId

`func (o *MagnetoObjectId) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *MagnetoObjectId) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *MagnetoObjectId) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *MagnetoObjectId) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *MagnetoObjectId) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *MagnetoObjectId) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetJobUid

`func (o *MagnetoObjectId) GetJobUid() UniversalIdProto`

GetJobUid returns the JobUid field if non-nil, zero value otherwise.

### GetJobUidOk

`func (o *MagnetoObjectId) GetJobUidOk() (*UniversalIdProto, bool)`

GetJobUidOk returns a tuple with the JobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobUid

`func (o *MagnetoObjectId) SetJobUid(v UniversalIdProto)`

SetJobUid sets JobUid field to given value.

### HasJobUid

`func (o *MagnetoObjectId) HasJobUid() bool`

HasJobUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


