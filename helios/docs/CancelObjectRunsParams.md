# CancelObjectRunsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectId** | **NullableInt64** | Specifies object id | 
**RunsConfig** | Pointer to [**[]CancelObjectRunParams**](CancelObjectRunParams.md) | Specifies a list of runs to cancel. If no runs are specified, then all the outstanding runs will be canceled. | [optional] 

## Methods

### NewCancelObjectRunsParams

`func NewCancelObjectRunsParams(objectId NullableInt64, ) *CancelObjectRunsParams`

NewCancelObjectRunsParams instantiates a new CancelObjectRunsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelObjectRunsParamsWithDefaults

`func NewCancelObjectRunsParamsWithDefaults() *CancelObjectRunsParams`

NewCancelObjectRunsParamsWithDefaults instantiates a new CancelObjectRunsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectId

`func (o *CancelObjectRunsParams) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *CancelObjectRunsParams) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *CancelObjectRunsParams) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### SetObjectIdNil

`func (o *CancelObjectRunsParams) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *CancelObjectRunsParams) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
### GetRunsConfig

`func (o *CancelObjectRunsParams) GetRunsConfig() []CancelObjectRunParams`

GetRunsConfig returns the RunsConfig field if non-nil, zero value otherwise.

### GetRunsConfigOk

`func (o *CancelObjectRunsParams) GetRunsConfigOk() (*[]CancelObjectRunParams, bool)`

GetRunsConfigOk returns a tuple with the RunsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunsConfig

`func (o *CancelObjectRunsParams) SetRunsConfig(v []CancelObjectRunParams)`

SetRunsConfig sets RunsConfig field to given value.

### HasRunsConfig

`func (o *CancelObjectRunsParams) HasRunsConfig() bool`

HasRunsConfig returns a boolean if a field has been set.

### SetRunsConfigNil

`func (o *CancelObjectRunsParams) SetRunsConfigNil(b bool)`

 SetRunsConfigNil sets the value for RunsConfig to be an explicit nil

### UnsetRunsConfig
`func (o *CancelObjectRunsParams) UnsetRunsConfig()`

UnsetRunsConfig ensures that no value is present for RunsConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


