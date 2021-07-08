# RunMapReduceInstanceResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**MapReduceInstanceId** | Pointer to **NullableInt64** | Return the ID of instance. | [optional] 

## Methods

### NewRunMapReduceInstanceResult

`func NewRunMapReduceInstanceResult() *RunMapReduceInstanceResult`

NewRunMapReduceInstanceResult instantiates a new RunMapReduceInstanceResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunMapReduceInstanceResultWithDefaults

`func NewRunMapReduceInstanceResultWithDefaults() *RunMapReduceInstanceResult`

NewRunMapReduceInstanceResultWithDefaults instantiates a new RunMapReduceInstanceResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *RunMapReduceInstanceResult) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RunMapReduceInstanceResult) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RunMapReduceInstanceResult) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *RunMapReduceInstanceResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMapReduceInstanceId

`func (o *RunMapReduceInstanceResult) GetMapReduceInstanceId() int64`

GetMapReduceInstanceId returns the MapReduceInstanceId field if non-nil, zero value otherwise.

### GetMapReduceInstanceIdOk

`func (o *RunMapReduceInstanceResult) GetMapReduceInstanceIdOk() (*int64, bool)`

GetMapReduceInstanceIdOk returns a tuple with the MapReduceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapReduceInstanceId

`func (o *RunMapReduceInstanceResult) SetMapReduceInstanceId(v int64)`

SetMapReduceInstanceId sets MapReduceInstanceId field to given value.

### HasMapReduceInstanceId

`func (o *RunMapReduceInstanceResult) HasMapReduceInstanceId() bool`

HasMapReduceInstanceId returns a boolean if a field has been set.

### SetMapReduceInstanceIdNil

`func (o *RunMapReduceInstanceResult) SetMapReduceInstanceIdNil(b bool)`

 SetMapReduceInstanceIdNil sets the value for MapReduceInstanceId to be an explicit nil

### UnsetMapReduceInstanceId
`func (o *RunMapReduceInstanceResult) UnsetMapReduceInstanceId()`

UnsetMapReduceInstanceId ensures that no value is present for MapReduceInstanceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


