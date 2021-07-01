# MapReduceInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | System generated ID of map reduce instance. | [optional] 
**InputParams** | Pointer to [**[]MapReduceInstanceInputParam**](MapReduceInstanceInputParam.md) |  | [optional] 
**InputSpec** | Pointer to [**InputSpec**](InputSpec.md) |  | [optional] 
**MapReduceInfoId** | Pointer to **NullableInt64** | ID of Map reduce info. | [optional] 
**OutputSpec** | Pointer to [**OutputSpec**](OutputSpec.md) |  | [optional] 
**RunInfo** | Pointer to [**MapReduceInstanceRunInfo**](MapReduceInstanceRunInfo.md) |  | [optional] 

## Methods

### NewMapReduceInstance

`func NewMapReduceInstance() *MapReduceInstance`

NewMapReduceInstance instantiates a new MapReduceInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapReduceInstanceWithDefaults

`func NewMapReduceInstanceWithDefaults() *MapReduceInstance`

NewMapReduceInstanceWithDefaults instantiates a new MapReduceInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MapReduceInstance) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MapReduceInstance) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MapReduceInstance) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MapReduceInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MapReduceInstance) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MapReduceInstance) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetInputParams

`func (o *MapReduceInstance) GetInputParams() []MapReduceInstanceInputParam`

GetInputParams returns the InputParams field if non-nil, zero value otherwise.

### GetInputParamsOk

`func (o *MapReduceInstance) GetInputParamsOk() (*[]MapReduceInstanceInputParam, bool)`

GetInputParamsOk returns a tuple with the InputParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParams

`func (o *MapReduceInstance) SetInputParams(v []MapReduceInstanceInputParam)`

SetInputParams sets InputParams field to given value.

### HasInputParams

`func (o *MapReduceInstance) HasInputParams() bool`

HasInputParams returns a boolean if a field has been set.

### SetInputParamsNil

`func (o *MapReduceInstance) SetInputParamsNil(b bool)`

 SetInputParamsNil sets the value for InputParams to be an explicit nil

### UnsetInputParams
`func (o *MapReduceInstance) UnsetInputParams()`

UnsetInputParams ensures that no value is present for InputParams, not even an explicit nil
### GetInputSpec

`func (o *MapReduceInstance) GetInputSpec() InputSpec`

GetInputSpec returns the InputSpec field if non-nil, zero value otherwise.

### GetInputSpecOk

`func (o *MapReduceInstance) GetInputSpecOk() (*InputSpec, bool)`

GetInputSpecOk returns a tuple with the InputSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSpec

`func (o *MapReduceInstance) SetInputSpec(v InputSpec)`

SetInputSpec sets InputSpec field to given value.

### HasInputSpec

`func (o *MapReduceInstance) HasInputSpec() bool`

HasInputSpec returns a boolean if a field has been set.

### GetMapReduceInfoId

`func (o *MapReduceInstance) GetMapReduceInfoId() int64`

GetMapReduceInfoId returns the MapReduceInfoId field if non-nil, zero value otherwise.

### GetMapReduceInfoIdOk

`func (o *MapReduceInstance) GetMapReduceInfoIdOk() (*int64, bool)`

GetMapReduceInfoIdOk returns a tuple with the MapReduceInfoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapReduceInfoId

`func (o *MapReduceInstance) SetMapReduceInfoId(v int64)`

SetMapReduceInfoId sets MapReduceInfoId field to given value.

### HasMapReduceInfoId

`func (o *MapReduceInstance) HasMapReduceInfoId() bool`

HasMapReduceInfoId returns a boolean if a field has been set.

### SetMapReduceInfoIdNil

`func (o *MapReduceInstance) SetMapReduceInfoIdNil(b bool)`

 SetMapReduceInfoIdNil sets the value for MapReduceInfoId to be an explicit nil

### UnsetMapReduceInfoId
`func (o *MapReduceInstance) UnsetMapReduceInfoId()`

UnsetMapReduceInfoId ensures that no value is present for MapReduceInfoId, not even an explicit nil
### GetOutputSpec

`func (o *MapReduceInstance) GetOutputSpec() OutputSpec`

GetOutputSpec returns the OutputSpec field if non-nil, zero value otherwise.

### GetOutputSpecOk

`func (o *MapReduceInstance) GetOutputSpecOk() (*OutputSpec, bool)`

GetOutputSpecOk returns a tuple with the OutputSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSpec

`func (o *MapReduceInstance) SetOutputSpec(v OutputSpec)`

SetOutputSpec sets OutputSpec field to given value.

### HasOutputSpec

`func (o *MapReduceInstance) HasOutputSpec() bool`

HasOutputSpec returns a boolean if a field has been set.

### GetRunInfo

`func (o *MapReduceInstance) GetRunInfo() MapReduceInstanceRunInfo`

GetRunInfo returns the RunInfo field if non-nil, zero value otherwise.

### GetRunInfoOk

`func (o *MapReduceInstance) GetRunInfoOk() (*MapReduceInstanceRunInfo, bool)`

GetRunInfoOk returns a tuple with the RunInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunInfo

`func (o *MapReduceInstance) SetRunInfo(v MapReduceInstanceRunInfo)`

SetRunInfo sets RunInfo field to given value.

### HasRunInfo

`func (o *MapReduceInstance) HasRunInfo() bool`

HasRunInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


