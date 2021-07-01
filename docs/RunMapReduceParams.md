# RunMapReduceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **NullableInt64** | ApplicationId is the Id of the map reduce application to run. | [optional] 
**InputParams** | Pointer to [**[]MapReduceInstanceInputParam**](MapReduceInstanceInputParam.md) | InputParams specifies optional list of key&#x3D;value input params specified for running the map reduce instance. | [optional] 
**MrInput** | Pointer to [**InputSpec**](InputSpec.md) |  | [optional] 
**MrOutput** | Pointer to [**OutputSpec**](OutputSpec.md) |  | [optional] 

## Methods

### NewRunMapReduceParams

`func NewRunMapReduceParams() *RunMapReduceParams`

NewRunMapReduceParams instantiates a new RunMapReduceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunMapReduceParamsWithDefaults

`func NewRunMapReduceParamsWithDefaults() *RunMapReduceParams`

NewRunMapReduceParamsWithDefaults instantiates a new RunMapReduceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *RunMapReduceParams) GetAppId() int64`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *RunMapReduceParams) GetAppIdOk() (*int64, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *RunMapReduceParams) SetAppId(v int64)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *RunMapReduceParams) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *RunMapReduceParams) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *RunMapReduceParams) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetInputParams

`func (o *RunMapReduceParams) GetInputParams() []MapReduceInstanceInputParam`

GetInputParams returns the InputParams field if non-nil, zero value otherwise.

### GetInputParamsOk

`func (o *RunMapReduceParams) GetInputParamsOk() (*[]MapReduceInstanceInputParam, bool)`

GetInputParamsOk returns a tuple with the InputParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParams

`func (o *RunMapReduceParams) SetInputParams(v []MapReduceInstanceInputParam)`

SetInputParams sets InputParams field to given value.

### HasInputParams

`func (o *RunMapReduceParams) HasInputParams() bool`

HasInputParams returns a boolean if a field has been set.

### SetInputParamsNil

`func (o *RunMapReduceParams) SetInputParamsNil(b bool)`

 SetInputParamsNil sets the value for InputParams to be an explicit nil

### UnsetInputParams
`func (o *RunMapReduceParams) UnsetInputParams()`

UnsetInputParams ensures that no value is present for InputParams, not even an explicit nil
### GetMrInput

`func (o *RunMapReduceParams) GetMrInput() InputSpec`

GetMrInput returns the MrInput field if non-nil, zero value otherwise.

### GetMrInputOk

`func (o *RunMapReduceParams) GetMrInputOk() (*InputSpec, bool)`

GetMrInputOk returns a tuple with the MrInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrInput

`func (o *RunMapReduceParams) SetMrInput(v InputSpec)`

SetMrInput sets MrInput field to given value.

### HasMrInput

`func (o *RunMapReduceParams) HasMrInput() bool`

HasMrInput returns a boolean if a field has been set.

### GetMrOutput

`func (o *RunMapReduceParams) GetMrOutput() OutputSpec`

GetMrOutput returns the MrOutput field if non-nil, zero value otherwise.

### GetMrOutputOk

`func (o *RunMapReduceParams) GetMrOutputOk() (*OutputSpec, bool)`

GetMrOutputOk returns a tuple with the MrOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrOutput

`func (o *RunMapReduceParams) SetMrOutput(v OutputSpec)`

SetMrOutput sets MrOutput field to given value.

### HasMrOutput

`func (o *RunMapReduceParams) HasMrOutput() bool`

HasMrOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


